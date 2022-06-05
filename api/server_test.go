package api

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	middleware "github.com/fsvxavier/unico/middlewares"
	util "github.com/fsvxavier/unico/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

var (
	router *gin.Engine
)

func init() {
	gin.SetMode(gin.TestMode)

}

func TestInitializeSuccess(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("sucesso", func(mt *mtest.T) {

		server := Server{}

		var logger util.GenericLogger
		logger.Module = "server"
		logger.GetLogger()

		os.Setenv("NEW_RELIC_ACTIVE", "false")

		server.Route = gin.New()
		server.Route.Use(middleware.Logger(logger.Log.Logger))
		server.Route.Use(middleware.CORS())

		server.Initialize()

		server.Route = gin.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		g := httptest.NewRecorder()
		server.Route.ServeHTTP(g, req)
	})
}

func TestStartServer(t *testing.T) {

	router = gin.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

}
