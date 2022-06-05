package middlewares_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/fsvxavier/unico/middlewares"
	util "github.com/fsvxavier/unico/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func setupRouterLogger() *gin.Engine {
	var logger util.GenericLogger
	logger.Module = "teste"
	logger.GetLogger()

	router := gin.New()
	router.Use(middlewares.Logger(logger.Log.Logger))
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

func TestLogger200(t *testing.T) {
	var logger util.GenericLogger
	logger.Module = "teste"
	logger.GetLogger()

	router := gin.New()
	router.Use(middlewares.Logger(logger.Log.Logger))
	router.GET("/example", func(c *gin.Context) {})

	w := util.PerformRequestV2(router, "GET", "/example?a=100")
	response := w.Result()
	assert.Equal(t, 200, response.StatusCode)
}

func TestLogger404(t *testing.T) {
	var logger util.GenericLogger
	logger.Module = "teste"
	logger.GetLogger()

	router := gin.New()
	router.Use(middlewares.Logger(logger.Log.Logger))

	router.GET("/example", func(c *gin.Context) {})

	w := util.PerformRequestV2(router, "GET", "/notfound")
	response := w.Result()
	assert.Equal(t, 404, response.StatusCode)
}

func TestLogger500(t *testing.T) {
	var logger util.GenericLogger
	logger.Module = "teste"
	logger.GetLogger()

	router := gin.New()
	router.Use(middlewares.Logger(logger.Log.Logger))
	router.GET("/err", func(c *gin.Context) {
		c.Status(500)
	})

	w := util.PerformRequestV2(router, "GET", "/err")
	response := w.Result()
	assert.Equal(t, 500, response.StatusCode)
}

func TestLogger500Abort(t *testing.T) {
	var logger util.GenericLogger
	logger.Module = "teste"
	logger.GetLogger()

	router := gin.New()
	router.Use(middlewares.Logger(logger.Log.Logger))
	wantErr := errors.New("oh no")
	router.GET("/err", func(c *gin.Context) {
		c.AbortWithError(500, wantErr)
	})

	w := util.PerformRequestV2(router, "GET", "/err")
	response := w.Result()
	assert.Equal(t, 500, response.StatusCode)
}
