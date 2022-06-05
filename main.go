package main

import (
	"log"
	"runtime"

	"github.com/fsvxavier/unico/api"
	_ "github.com/fsvxavier/unico/config"
	"github.com/fsvxavier/unico/config/db"
	util "github.com/fsvxavier/unico/utils"
	"github.com/gin-gonic/gin"
)

// @securityDefinitions.apikey Authorization
// @in header
// @name Authorization
func main() {

	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)

	var logger util.GenericLogger
	logger.GetLogger()

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	db.InitMysqlDB()

	server := api.Server{
		DBConnMySQL: db.MysqlDB,
	}
	server.Route = gin.New()

	server.StartServer()

}
