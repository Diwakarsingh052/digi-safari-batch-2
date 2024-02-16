package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"small-app/middleware"
	"small-app/models"
	"small-app/models/mockmodels"
	"testing"
)

func TestHandlerSignup(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUser := models.User{
		Id:           "ab49a45c-ec2c-47a5-8675-9f072e2d9216",
		Email:        "d@email.com",
		Name:         "John Doe",
		Age:          30,
		PasswordHash: "2a$10$EimVQRw4YiKIoMqh3JMwOesA9ngPGZT.chFEmPSaHzYl.mlnhLr12",
	}
	newUser := models.NewUser{
		Name:     "John Doe",
		Email:    "d@email.com",
		Age:      30,
		Password: "abc",
	}
	tt := []struct {
		name                string
		body                any    // Body to send to request
		expectedStatus      int    // Expected status of the response
		expectedResponse    string // Expected response body
		mockUserServiceFunc func(m *mockmodels.MockService)
	}{
		{
			name: "OK",
			body: models.NewUser{
				Name:     "John Doe",
				Age:      30,
				Email:    "d@email.com",
				Password: "abc",
			},
			expectedStatus:   http.StatusOK,
			expectedResponse: `{"Id":"ab49a45c-ec2c-47a5-8675-9f072e2d9216","email":"d@email.com","name":"John Doe","age":30,"password_hash":"2a$10$EimVQRw4YiKIoMqh3JMwOesA9ngPGZT.chFEmPSaHzYl.mlnhLr12"}`,
			//a function setting up an expectation on a mock service
			mockUserServiceFunc: func(m *mockmodels.MockService) {
				m.EXPECT().CreateUser(gomock.Eq(newUser)).Times(1).Return(mockUser, nil)
			},
		},
	}
	router := gin.New()
	// Create a new Gomock controller.
	ctrl := gomock.NewController(t)
	// get the mock implementation of the Service Interface
	mockService := mockmodels.NewMockService(ctrl)
	s := models.NewStore(mockService)
	h := handler{Store: s}
	router.POST("/signup", h.Signup)

	for _, tc := range tt {

		// Apply the mock to the user service.
		tc.mockUserServiceFunc(mockService)
		ctx := context.Background()
		// Create a fake TraceID.
		traceID := "fake-trace-id"
		// Insert the TraceId into the context.
		ctx = context.WithValue(ctx, middleware.TraceIdKey, traceID)

		// Create a new HTTP Response Recorder.
		// ResponseRecorder is an implementation of http.ResponseWriter that
		// records its mutations for later inspection in tests.
		rec := httptest.NewRecorder()
		body, err := json.Marshal(tc.body)
		require.NoError(t, err)

		//request construction with the context which have a fake trace-id in it
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/signup", bytes.NewReader(body))
		require.NoError(t, err)

		//this would invoke our handler function signup
		router.ServeHTTP(rec, req)

		// rec.Body.String() gives us the response that was written over response writer
		require.Equal(t, tc.expectedResponse, rec.Body.String())
		require.Equal(t, tc.expectedStatus, rec.Code)
	}
}
