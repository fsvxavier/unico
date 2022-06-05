package config

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/fsvxavier/unico/models"
)

//Config o aplicativo
type Config struct {
	Env string
}

type key int

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	file, err := ioutil.ReadFile(basepath + "/env.json")
	if err != nil {
		return
	}

	jsonMap := make(map[string]interface{})
	json.Unmarshal(file, &jsonMap)

	env := "production"
	if jsonMap[env] == nil {
		return
	}
	database := jsonMap[env].(map[string]interface{})

	for key, value := range database {
		switch value.(type) {
		case string:
			os.Setenv(key, value.(string))
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
			os.Setenv(key, fmt.Sprintf("%d", value.(int)))
		case float32, float64:
			val := fmt.Sprintf("%.2f", value.(float64))
			strings := strings.Split(val, ".")
			if strings[1] != "00" {
				os.Setenv(key, val)
			} else {
				os.Setenv(key, strings[0])
			}
		case bool:
			os.Setenv(key, fmt.Sprintf("%v", value.(bool)))
		}
	}
}

//ResponseWithJSON ...
func (c *Config) ResponseWithJSON(g *gin.Context, code int, payload interface{}) {
	var r models.ResponseSuccess
	r.Records = payload
	lenPayload := reflect.ValueOf(payload)
	r.Meta.RecordCount = 1
	r.Meta.Limit = 1
	if lenPayload.Kind() == reflect.Slice {
		r.Meta.Limit = lenPayload.Len()
		r.Meta.RecordCount = lenPayload.Len()
	}
	g.JSON(code, r)
}

//RespondWithError corresponde a funcao que restorna erro
func (c *Config) RespondWithError(g *gin.Context, code int, message string, moreInfo string) {
	var m models.ResponseError
	m.DeveloperMessage = message
	m.UserMessage = "Erro"
	m.ErrorCode = code
	m.MoreInfo = moreInfo
	e := getMessageError(code)
	g.JSON(code, e)
}

// ResponseWithError ...
func (c *Config) ResponseWithError(g *gin.Context, code int, message string, moreInfo string) {
	var m models.ResponseError
	m.DeveloperMessage = message
	m.UserMessage = "Erro"
	m.ErrorCode = code
	m.MoreInfo = moreInfo
	e := models.ResponseError{
		DeveloperMessage: message,
		UserMessage:      moreInfo,
		MoreInfo:         "http://www.fabricioxavier.br",
		ErrorCode:        code,
	}
	g.JSON(code, e)
}

func respondWithError(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

//GetStatusCode ...
func (c *Config) GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	switch err {
	case models.ErrInternalServerError:
		return http.StatusInternalServerError
	case models.ErrNotFound:
		return http.StatusNotFound
	case sql.ErrNoRows:
		return http.StatusNotFound
	case models.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

//GetMessageError ...
func getMessageError(errorCode int) *models.ResponseError {
	switch errorCode {
	case http.StatusInternalServerError:
		return &models.ResponseError{
			DeveloperMessage: "Internal server error",
			UserMessage:      "Was encountered an error when processing your request. We apologize for the inconvenience",
			MoreInfo:         "http://www.fabricioxavier.com.br",
			ErrorCode:        errorCode,
		}
	default:
		return &models.ResponseError{
			DeveloperMessage: "Resource not found",
			UserMessage:      "Resource not found",
			MoreInfo:         "http://www.fabricioxavier.com.br",
			ErrorCode:        404,
		}
	}
}
