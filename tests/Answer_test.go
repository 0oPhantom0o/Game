package tests

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Mock the Logic interface

func (m *MockLogic) CheckAnswer(id, answer string) (bool, error) {
	args := m.Called(id, answer)
	return args.Bool(0), args.Error(1)
}

// Test function for Answer
func TestAnswer(t *testing.T) {
	gin.SetMode(gin.TestMode) // Set Gin to test mode
	router := gin.Default()
	mockLogic := new(MockLogic)
	ctrl := &GameController{Logic: mockLogic}

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
			token:          "valid_token",
			mockReturn:     true,
			mockErr:        nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"message":"right_answer"}`,
		},
		{
			name:           "Wrong Answer",
			input:          `{"answer":"23"}`,
			token:          "valid_token",
			mockReturn:     false,
			mockErr:        nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"message":"wronged_answer"}`,
		},
		{
			name:           "Error Checking Answer",
			input:          `{"answer":"42"}`,
			token:          "valid_token",
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
			// Setup the mock
			mockLogic.On("CheckAnswer", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(tt.mockReturn, tt.mockErr)

			// Create a new HTTP request
			req, _ := http.NewRequest(http.MethodPost, "/answer", strings.NewReader(tt.input))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", tt.token)

			// Create a ResponseRecorder to record the response
			w := httptest.NewRecorder()

			// Call the Answer handler
			router.POST("/answer", ctrl.Answer)
			router.ServeHTTP(w, req)

			// Assert the status code and response body
			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())

			// Assert that the mock was called
			mockLogic.AssertExpectations(t)
		})
	}
}
