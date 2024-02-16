package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log/slog"
	"net/http"
	"small-app/models"
	"small-app/pkg/ctxmanage"
)

type handler struct {
	// Conn is a dependency for handlers package,
	//adding it in the struct so handler package method can call method using conn struct
	//models.Conn
	models.Store // using a struct that wraps interface instead of using conn type directly
}

/*
signup

	{
	  "name": "John Doe",
	  "email": "d@email.com",
	  "age": 19,
	  "password": "password123"
	}
*/
// Function to handle Signup
func (h *handler) Signup(c *gin.Context) {

	// Get the traceId from the request. Useful for tracking the request in logs
	traceId := ctxmanage.GetTraceIdOfRequest(c)

	// Check if the size of the request body is more than 5KB
	if c.Request.ContentLength > 5*1024 {
		// Log error for payload exceeding size limit
		slog.Error("request body limit breached", slog.String("TRACE ID", traceId), slog.Int64("Size Received", c.Request.ContentLength))

		// Return a 400 Bad Request status code along with an error message
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Request body too large. Limit is 5KB"})

		// Gracefully terminate this function
		return
	}

	// Declare a variable to hold decoded data from request body
	var newUser models.NewUser

	// Bind JSON request body to newUser struct. c.ShouldBindJSON() internally uses json.Unmarshal and returns error if any
	err := c.ShouldBindJSON(&newUser)

	// Check if JSON bind resulted in errors
	if err != nil {

		// Log error and associate it with a trace id for easy correlation
		slog.Error("json validation error", slog.String("TRACE ID", traceId),
			slog.String("Error", err.Error()))

		// Respond with a 400 Bad Request status code and error message
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})

		// Stop further execution
		return
	}

	// Create a new validator instance
	validate := validator.New()

	// Validate the newUser struct using the validate instance
	err = validate.Struct(newUser)

	// Check if validation encountered errors
	if err != nil {

		// Declare a variable to hold validation errors
		var vErrs validator.ValidationErrors

		// Check if the error can be asserted to validation errors
		if errors.As(err, &vErrs) {

			// Loop through each validation error
			for _, vErr := range vErrs {

				// Process error based on error tag
				switch vErr.Tag() {

				case "required":

					// Log the validation error
					slog.Error("validation failed", slog.String("TRACE ID", traceId),
						slog.String("Error", err.Error()))

					// Respond with 400 Bad Request and error message identified by vErr.Field + " value missing"
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": vErr.Field() + " value missing"})

					// Stop further execution
					return
				case "min":
					slog.Error("validation failed", slog.String("TRACE ID", traceId),
						slog.String("Error", err.Error()))
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": vErr.Field() + " value is less than " + vErr.Param()})
					return

				default:
					slog.Error("validation failed", slog.String("TRACE ID", traceId),
						slog.String("Error", err.Error()))
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
					return
				}

			}
		}
		// Log the validation error along with its context

		slog.Error("validation failed", slog.String("TRACE ID", traceId),
			slog.String("ERROR", err.Error()))

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return

	}
	uData, err := h.CreateUser(newUser)
	// If user fetch operation fails, respond with an error
	if err != nil {

		slog.Error("error in creating the user", slog.String("Trace ID", traceId),
			slog.String("Error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Problem in creating user"})
		return
	}

	c.JSON(http.StatusOK, uData)

}

// Rest of your logic here...
