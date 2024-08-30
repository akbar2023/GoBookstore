package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	// Adjust import path
	"example/web-service-gin/handlers" // Adjust import path
	"example/web-service-gin/models"   // Adjust import path

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetBooks(t *testing.T) {
	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Initialize test database
	db := database_test.InitTestDB()
	defer database_test.CloseTestDB(db) // Ensure the database is closed after the test

	// Setup Router
	router := gin.Default()

	// Insert mock data into the test database
	mockBooks := []models.Book{
		{Title: "Book 1", Author: "Author 1", Price: 10.99},
		{Title: "Book 2", Author: "Author 2", Price: 15.49},
	}
	db.Create(&mockBooks)

	// Register routes and handlers
	router.GET("/books", handlers.GetBooks)

	// Create a request to the router
	req, err := http.NewRequest(http.MethodGet, "/books", nil)
	assert.NoError(t, err)

	// Create a ResponseRecorder to record the response
	recorder := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(recorder, req)

	// Check if the status code is 200
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Check if the response body is as expected
	expectedBody := `[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"title":"Book 1","author":"Author 1","price":10.99},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"title":"Book 2","author":"Author 2","price":15.49}]`

	// Use JSONEq to compare JSON strings without strict ordering
	assert.JSONEq(t, expectedBody, recorder.Body.String())
}
