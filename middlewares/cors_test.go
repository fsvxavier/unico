package middlewares_test

import (
	"net/http"
	"testing"

	"github.com/bmizerany/assert"
	"github.com/fsvxavier/unico/middlewares"
	util "github.com/fsvxavier/unico/utils"
	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func setupRouterCors() *gin.Engine {
	router := gin.New()
	router.Use(middlewares.CORS())
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "get")
	})
	router.POST("/", func(c *gin.Context) {
		c.String(http.StatusOK, "post")
	})
	router.PATCH("/", func(c *gin.Context) {
		c.String(http.StatusOK, "patch")
	})
	router.OPTIONS("/", func(c *gin.Context) {
		c.String(http.StatusOK, "options")
	})
	return router
}

func TestCors(t *testing.T) {
	router := setupRouterCors()

	w := util.PerformRequest(router, "GET", "https://gist.github.com")
	assert.Equal(t, 200, w.Code)

	w = util.PerformRequest(router, "GET", "https://github.com")
	assert.Equal(t, 200, w.Code)

	w = util.PerformRequest(router, "OPTIONS", "https://github.com")
	assert.Equal(t, 204, w.Code)
}
