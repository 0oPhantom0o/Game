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

func (m *MockLogic) GenerateUser(phone, randomCode string) (string, error) {
	args := m.Called(phone, randomCode)
	return args.String(0), args.Error(1)
}

// Test function for GetOtp
func TestGetOtp(t *testing.T) {
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
		mockReturn     string
		mockErr        error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid OTP Answer",
			input:          `{"phone":"1234567890","randomCode":"123456"}`,
			mockReturn:     "generated_token",
			mockErr:        nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"token":"generated_token"}`,
		},
		{
			name:           "Invalid OTP Answer",
			input:          `{"phone":"1234567890","randomCode":"wrong_code"}`,
			mockReturn:     "",
			mockErr:        errors.New("wrong answer"),
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"wrong answer"}`,
		},
		{
			name:           "User Limited",
			input:          `{"phone":"1234567890","randomCode":"wrong_code"}`,
			mockReturn:     "",
			mockErr:        errors.New("user Limited. wait 10 minutes"),
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"user Limited. wait 10 minutes"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup the mock
			mockLogic.On("GenerateUser", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(tt.mockReturn, tt.mockErr)

			// Create a new HTTP request
			req, _ := http.NewRequest(http.MethodPost, "/get-otp", strings.NewReader(tt.input))
			req.Header.Set("Content-Type", "application/json")

			// Create a ResponseRecorder to record the response
			w := httptest.NewRecorder()

			// Call the GetOtp handler
			router.POST("/get-otp", ctrl.GetOtp)
			router.ServeHTTP(w, req)

			// Assert the status code and response body
			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())

			// Assert that the mock was called
			mockLogic.AssertExpectations(t)
		})
	}
}
