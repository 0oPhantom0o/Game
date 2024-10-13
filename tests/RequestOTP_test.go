package tests

import (
	"context"
	"errors"
	"game/app"
	"game/controller"
	"game/logic"
	"game/repository"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock the Logic interface
type MockLogic struct {
	mock.Mock
}

func (m *MockLogic) RequestOtp(phone string) error {
	args := m.Called(phone)
	return args.Error(0)
}

func TestRequestOtp(t *testing.T) {
	gin.SetMode(gin.TestMode) // Set Gin to test mode
	router := gin.Default()
	mongodb, err := app.Collection()
	if err != nil {
		log.Fatalf("Error initializing databases: %v", err)
	}
	ctx := context.Background()
	redisdb := app.RedisDB
	repo := repository.NewMongoRepository(redisdb, mongodb, ctx)
	svc := logic.NewRestaurantService(repo)

	mockLogic := new(MockLogic)
	ctrl := &controller.GameController{Logic: svc}

	// Test cases
	tests := []struct {
		name           string
		input          string
		mockReturn     error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid OTP Request",
			input:          `{"phone":"1234567890"}`,
			mockReturn:     nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"message":"OTP requested successfully"}`,
		},
		{
			name:           "Invalid Phone Number Format",
			input:          `{"phone":""}`,
			mockReturn:     errors.New("invalid phone number"),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"invalid request: invalid phone number"}`,
		},
		{
			name:           "OTP Request Limit Exceeded",
			input:          `{"phone":"1234567890"}`,
			mockReturn:     errors.New("user has reached the maximum limit of OTP requests"),
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"user has reached the maximum limit of OTP requests"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup the mock
			mockLogic.On("RequestOtp", mock.AnythingOfType("string")).Return(tt.mockReturn)

			// Create a new HTTP request
			req, _ := http.NewRequest(http.MethodPost, "/request-otp", strings.NewReader(tt.input))
			req.Header.Set("Content-Type", "application/json")

			// Create a ResponseRecorder to record the response
			w := httptest.NewRecorder()

			// Call the RequestOtp handler
			router.POST("/request-otp", ctrl.RequestOtp)
			router.ServeHTTP(w, req)

			// Assert the status code and response body
			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())

			// Assert that the mock was called
			mockLogic.AssertExpectations(t)
		})
	}
}
