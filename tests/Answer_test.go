package tests

import (
	"context"
	"errors"
	"game/app"
	"game/controller"
	"game/logic"
	"game/repository"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func (m *MockLogic) CheckAnswer(id, answer string) (bool, error) {
	args := m.Called(id, answer)
	return args.Bool(0), args.Error(1)
}

// Test function for Answer
func TestAnswer(t *testing.T) {
	gin.SetMode(gin.TestMode) // Set Gin to test mode
	if err := app.InitDb(); err != nil {
		log.Panicf("DataBase is not running:%v", err)
	}

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
		token          string
		mockReturn     bool
		mockErr        error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Correct Answer",
			input:          `{"answer":"42"}`,
			token:          "Bearer valid_token",
			mockReturn:     true,
			mockErr:        nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"message":"right_answer"}`,
		},
		{
			name:           "Wrong Answer",
			input:          `{"answer":"23"}`,
			token:          "Bearer valid_token",
			mockReturn:     false,
			mockErr:        nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"message":"wronged_answer"}`,
		},
		{
			name:           "Error Checking Answer",
			input:          `{"answer":"42"}`,
			token:          "Bearer valid_token", // Ensure this includes "Bearer "
			mockReturn:     false,
			mockErr:        errors.New("wrong answer. search for a new question"),
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"message":"wrong answer. search for a new question"}`,
		},
		{
			name:           "Unauthorized",
			input:          `{"answer":"42"}`,
			token:          "invalid_token",
			mockReturn:     false,
			mockErr:        nil,
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   `{"error":"invalid token"}`, // Assuming you handle this in the controller
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new router for each test

			// Register the route
			router.POST("/answer", ctrl.Answer)

			// Setup the mock
			mockLogic.On("CheckAnswer", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(tt.mockReturn, tt.mockErr)

			// Create a new HTTP request
			req, _ := http.NewRequest(http.MethodPost, "/answer", strings.NewReader(tt.input))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", tt.token)

			// Create a ResponseRecorder to record the response
			w := httptest.NewRecorder()

			// Call the Answer handler
			router.ServeHTTP(w, req)

			// Assert the status code and response body
			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())

			// Assert that the mock was called
			mockLogic.AssertExpectations(t)
		})
	}

}
