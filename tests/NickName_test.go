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

func (m *MockLogic) UpdateNickName(nickName, id string) error {
	args := m.Called(nickName, id)
	return args.Error(0)
}

// Mock token verification
func MockVerifyToken(token string) (string, error) {
	if token == "valid_token" {
		return "user_id_123", nil
	}
	return "", errors.New("invalid token")
}

// Test function for NickName
func TestNickName(t *testing.T) {
	gin.SetMode(gin.TestMode) // Set Gin to test mode
	router := gin.Default()
	mockLogic := new(MockLogic)
	ctrl := &GameController{Logic: mockLogic}

	// Test cases
	tests := []struct {
		name           string
		input          string
		token          string
		mockReturn     error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid Nickname Change",
			input:          `{"nickName":"new_nick"}`,
			token:          "valid_token",
			mockReturn:     nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"nickname_changed_to":{"nickName":"new_nick"}}`,
		},
		{
			name:           "Invalid Token",
			input:          `{"nickName":"new_nick"}`,
			token:          "invalid_token",
			mockReturn:     nil,
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   `{"error":"invalid token"}`,
		},
		{
			name:           "Error Changing Nickname",
			input:          `{"nickName":"new_nick"}`,
			token:          "valid_token",
			mockReturn:     errors.New("you changed nickname before"),
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"you changed nickname before"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup the mock
			mockLogic.On("UpdateNickName", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(tt.mockReturn)

			// Create a new HTTP request
			req, _ := http.NewRequest(http.MethodPost, "/nickname", strings.NewReader(tt.input))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", tt.token)

			// Create a ResponseRecorder to record the response
			w := httptest.NewRecorder()

			// Call the NickName handler
			router.POST("/nickname", ctrl.NickName)
			router.ServeHTTP(w, req)

			// Assert the status code and response body
			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())

			// Assert that the mock was called
			mockLogic.AssertExpectations(t)
		})
	}
}
