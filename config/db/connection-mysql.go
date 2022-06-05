package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	//factory
	_ "github.com/go-sql-driver/mysql"
)

//ServerMysql ...
type ServerMysql struct {
	DB  *sql.DB
	Env string
}

//factory
var (
	MysqlDB *sql.DB
)

//InitMysqlDB represent a factory of database
func InitMysqlDB() {
	a := ServerMysql{}
	a.Env = os.Getenv("ENV")
	connectionString := fmt.Sprintf("%s", a.GetDNS())
	var err error

	MysqlDB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Printf("[db/init] - Erro ao tentar abrir conexão (%s). Erro: %s", a.Env, err.Error())
	}

	maxLifeTimeInt, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFE_TIME"))
	maxIdleConns, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNS"))
	maxOpenConns, _ := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNS"))

	maxLifeTime := time.Duration(maxLifeTimeInt)

	MysqlDB.SetConnMaxLifetime(time.Minute * maxLifeTime)
	MysqlDB.SetMaxIdleConns(maxIdleConns)
	MysqlDB.SetMaxOpenConns(maxOpenConns)

	err = MysqlDB.Ping()
	if err != nil {
		log.Printf("[db/init] - Erro ao tentar abrir conexão (%s). Erro: %s", a.Env, err.Error())
	}
}

//GetDNS representa a recuperação do acesso ao banco
func (a *ServerMysql) GetDNS() string {
	var (
		user     string
		password string
		dbname   string
		host     string
		dbPort   int
	)
	file, err := ioutil.ReadFile("./config/env.json")
	if err == nil {
		jsonMap := make(map[string]interface{})
		json.Unmarshal(file, &jsonMap)

		env := os.Getenv("ENV")
		if env == "" {
			env = "production"
		}

		database := jsonMap[env].(map[string]interface{})

		user = fmt.Sprintf("%v", database["MYSQLDBUSER"])
		password = fmt.Sprintf("%v", database["MYSQLDBPASSWORD"])
		dbname = fmt.Sprintf("%v", database["MYSQLDBNAME"])
		host = fmt.Sprintf("%v", database["MYSQLDBHOST"])
		dbPort, _ = strconv.Atoi(fmt.Sprintf("%v", database["MYSQLDBPORT"]))
	} else {
		user = os.Getenv("MYSQLDBUSER")
		password = os.Getenv("MYSQLDBPASSWORD")
		dbname = os.Getenv("MYSQLDBNAME")
		host = os.Getenv("MYSQLDBHOST")
		dbPort, _ = strconv.Atoi(os.Getenv("MYSQLDBPORT"))
	}
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, dbPort, dbname)
	return connectionString
}
