package api

import (
	"database/sql"
	"fmt"
	"log"

	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/fsvxavier/unico/api/handlers"
	"github.com/fsvxavier/unico/docs"
	"github.com/fsvxavier/unico/middlewares"
	"github.com/fsvxavier/unico/repositories"
	"github.com/fsvxavier/unico/usecases"
	"github.com/fsvxavier/unico/utils"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//Server ...
type Server struct {
	Logger       utils.GenericLogger
	DBConnMySQL  *sql.DB
	Route        *gin.Engine
	RouteGroupV0 *gin.RouterGroup
	RouteGroupV1 *gin.RouterGroup
}

//Initialize the server
func (s *Server) Initialize() {

	GROUPREQUEST := "/api/" + os.Getenv("VERSION_API")

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Unico API"
	docs.SwaggerInfo.Description = "Aplicação Entrevista."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.BasePath = GROUPREQUEST

	s.Logger.Module = "server"
	s.Logger.GetLogger()
	s.Route.Use(gin.Recovery())

	s.Route.Use(gzip.Gzip(gzip.DefaultCompression))
	s.Route.Use(middlewares.CORS())
	s.Route.Use(middlewares.Logger(s.Logger.Log.Logger))

	//Initial interfaces.Repository
	healthCheckRepository := repositories.NewHealthCheckRepository()
	feiraLivreRepository := repositories.NewMySQLFeiraLivreRepository(s.DBConnMySQL)

	// Initial interfaces.UseCase and inject dependency of repository
	healthCheckUseCase := usecases.NewHealthCheckUseCase(healthCheckRepository)
	feiraLivreUseCase := usecases.NewFeiraLivreUseCase(feiraLivreRepository)

	s.RouteGroupV0 = s.Route.Group("/")
	s.RouteGroupV1 = s.Route.Group(GROUPREQUEST)

	handlers.NewHealthCheckHTTPHandler(s.RouteGroupV0, healthCheckUseCase)
	handlers.NewFeiraLivreHTTPHandler(s.RouteGroupV1, feiraLivreUseCase)

	s.RouteGroupV0.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//s.Route.NoRoute(utils.EndpointNotFound)
}

// StartServer start of server
func (s *Server) StartServer() {
	s.Initialize()

	// Start Server
	port, _ := strconv.Atoi(fmt.Sprintf("%v", os.Getenv("PORT")))
	httpReadTimeout, _ := strconv.Atoi(os.Getenv("HTTP_READ_TIMEOUT"))
	httpWriteTimeout, _ := strconv.Atoi(os.Getenv("HTTP_WRITE_TIMEOUT"))
	httpServer := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		ReadTimeout:  time.Duration(httpReadTimeout) * time.Second,
		WriteTimeout: time.Duration(httpWriteTimeout) * time.Second,
		Handler:      s.Route,
	}

	s.Logger.LogIt("INFO", fmt.Sprintf("Starting Server on port %d", port), nil)
	if err := httpServer.ListenAndServe(); err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}
