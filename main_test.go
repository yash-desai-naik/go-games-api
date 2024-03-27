package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// Start the server
	go func() {
		main()
	}()

	// Wait for the server to start
	time.Sleep(1 * time.Second)

	// Run tests
	code := m.Run()

	// Exit
	shutdownServer()

	// Return exit code
	os.Exit(code)
}

func TestListGames(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "status code should be OK")
	// t.Logf("\033[32mTestListGames: Passed\033[0m")
}

func TestGetGameByID(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/game?id=94c1d94dcfbd493eb0fa8b65072e9c03", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "status code should be OK")
	// t.Logf("\033[32mTestGetGameByID: Passed\033[0m")
}

func TestNonExistentID(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/game?id=9999", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code, "status code should be NotFound")
	// t.Logf("\033[32mTestNonExistentID: Passed\033[0m")
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", listGames)
	r.GET("/game", getGameByID)
	return r
}

func shutdownServer() {
	// Add shutdown logic here if needed
}
