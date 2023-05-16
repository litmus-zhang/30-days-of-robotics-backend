package main

import (
	"30-days-of-robotics-backend/src/controller"
	"30-days-of-robotics-backend/src/setup"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	return router
}
func TestAppSetup(t *testing.T) {
	router := SetupRouter()
	router.GET("/health", setup.HealthChecker)
	req, err := http.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()

	response := `{"message":"System is up and running"}`

	router.ServeHTTP(w, req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	responseMsg, _ := io.ReadAll(w.Body)
	assert.Equal(t, response, string(responseMsg))

}

func TestRegisterRoute(t *testing.T) {
	url := "/api/v1/users/register"
	router := SetupRouter()
	router.POST(url, controller.Register)
	user := map[string]string{
		"first_name":       "Litmus",
		"last_name":        "Zhang",
		"email":            "zhang@gmail.com",
		"password":         "litmus1234",
		"confirm_password": "litmus1234",
		"track":            "1",
	}

	t.Run("Status 201 Created", func(t *testing.T) {
		parsedUser, _ := json.Marshal(user)
		req, _ := http.NewRequest("POST", url, bytes.NewBuffer(parsedUser))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("Send success message", func(t *testing.T) {
		message := `{"message":"User registration successful"}`
		parsedUser, _ := json.Marshal(user)
		req, _ := http.NewRequest("POST", url, bytes.NewBuffer(parsedUser))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		formattedMsg, _ := io.ReadAll(w.Body)
		assert.Equal(t, message, string(formattedMsg))
	})
	t.Run("Send password mismatch", func(t *testing.T) {
		message := `{"message":"Password Mismatch"}`
		newUser := user
		newUser["confirm_password"] = "litmus12"
		parsedUser, _ := json.Marshal(newUser)
		req, _ := http.NewRequest("POST", url, bytes.NewBuffer(parsedUser))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		formattedMsg, _ := io.ReadAll(w.Body)
		assert.Equal(t, message, string(formattedMsg))
	})

}
