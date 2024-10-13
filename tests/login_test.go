package tests

import (
	"context"
	"errors"
	"game/app"
	"game/controller"
	_ "game/controller" // Import the package where GameController is defined
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

func (m *MockLogic) Login(phone string) (string, error) {
	args := m.Called(phone)
	return args.String(0), args.Error(1)
}

// Test function for Login
func TestLogin(t *testing.T) {
	if err := app.InitDb(); err != nil {
		log.Panicf("DataBase is not running:%v", err)
	}

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
			name:           "Successful Login",
			input:          `{"phone":"1234567890"}`,
			mockReturn:     "valid_token",
			mockErr:        nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"message":{"token":"valid_token"}}`,
		},
		{
			name:           "User Not Found",
			input:          `{"phone":"non_existent_phone"}`,
			mockReturn:     "",
			mockErr:        errors.New("user not found"),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"user not found"}`,
		},
		{
			name:           "Error Generating Token",
			input:          `{"phone":"1234567890"}`,
			mockReturn:     "",
			mockErr:        errors.New("failed to generate token"),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"failed to generate token"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup the mock
			mockLogic.On("Login", mock.AnythingOfType("string")).Return(tt.mockReturn, tt.mockErr)

			// Create a new HTTP request
			req, _ := http.NewRequest(http.MethodPost, "/login", strings.NewReader(tt.input))
			req.Header.Set("Content-Type", "application/json")

			// Create a ResponseRecorder to record the response
			w := httptest.NewRecorder()

			// Call the Login handler
			router.POST("/login", ctrl.Login)
			router.ServeHTTP(w, req)

			// Assert the status code and response body
			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())

			// Assert that the mock was called
			mockLogic.AssertExpectations(t)
		})
	}
}
