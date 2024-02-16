package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"small-app/middleware"
	"small-app/models"
	"small-app/pkg/ctxmanage"
)

func API() *gin.Engine {

	r := gin.New()
	c := models.NewConn()
	s := models.NewStore(&c)
	h := handler{Store: s}
	//apply middleware to all the endpoints using r.Use
	r.Use(middleware.Logger())
	r.GET("/check", check)
	r.POST("/signup", h.Signup)

	return r
}

func check(c *gin.Context) {
	traceId := ctxmanage.GetTraceIdOfRequest(c)
	u := struct {
		Status string
	}{
		Status: "Ok",
	}
	fmt.Println("check handler ", traceId)
	//JSON serializes the given struct as JSON into the response body. It also sets the Content-Type as "application/json".
	c.JSON(http.StatusOK, u)

}
