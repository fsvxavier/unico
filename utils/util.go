package utils

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/gin-gonic/gin"
)

type key int

// Const...
const (
	KeyNrID key = iota
)

var (
	// Err ...
	err error
	// Hostname ...
	Hostname string
	// Version ...
	Version string

	// App ...
	App string
	// Env ...
	Env string
	// Squad ...
	Squad string
	// Tribe ...
	Tribe string
	//LogLevel ...
	LogLevel string

	//MaxPoolSize
	MaxPoolSize int
)

type header struct {
	Key   string
	Value string
}

func init() {
	Hostname, err = os.Hostname()
	if err != nil {
		Hostname = "unknown"
	}
	Version = os.Getenv("VERSION_APP")

	//log envs
	App = os.Getenv("APP")
	Env = os.Getenv("ENV")
	Squad = os.Getenv("SQUAD")
	Tribe = os.Getenv("TRIBE")
	LogLevel = os.Getenv("LOGRUS_LOG_LEVEL")

}

// PerformRequestV2 ...
func PerformRequestV2(r http.Handler, method, path string, headers ...header) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, nil)
	for _, h := range headers {
		req.Header.Add(h.Key, h.Value)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func PerformRequest(r *gin.Engine, method, origin string) *httptest.ResponseRecorder {
	return performRequestWithHeaders(r, method, origin, http.Header{})
}

func performRequestWithHeaders(r *gin.Engine, method, origin string, header http.Header) *httptest.ResponseRecorder {
	httptest.NewRecorder()
	req, _ := http.NewRequest(method, "/", nil)
	req.Host = header.Get("Host")
	header.Del("Host")
	if len(origin) > 0 {
		header.Set("Origin", origin)
	}
	req.Header = header
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

// EndpointNotFound ...
func EndpointNotFound(c *gin.Context) {
	c.Writer.WriteString("there's no endpoint for that!")
}

// NewNullString ...
func NewNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}
