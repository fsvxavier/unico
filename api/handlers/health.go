package handlers

import (
	"net/http"

	"github.com/fsvxavier/unico/config"
	"github.com/fsvxavier/unico/interfaces"
	"github.com/gin-gonic/gin"
)

// httpHealthCheckHandler represent the httphandler for healthcheck
type httpHealthCheckHandler struct {
	HUsecase interfaces.HealthCheckUseCase
}

// NewHealthCheckHTTPHandler ...
func NewHealthCheckHTTPHandler(r *gin.RouterGroup, us interfaces.HealthCheckUseCase) {
	handler := &httpHealthCheckHandler{
		HUsecase: us,
	}
	r.GET("/health", handler.HealthCheck)
	r.GET("/", handler.HealthCheck)
}

// @Summary HealthCheck
// @Description HealthCheck API
// @Failure 400 {object} models.ResponseError
// @Success 200 {object} models.HealthCheck
// @Router /health [get]
func (h *httpHealthCheckHandler) HealthCheck(c *gin.Context) {
	var t config.Config
	hc, err := h.HUsecase.Check()

	if err != nil {
		t.RespondWithError(c, http.StatusBadRequest, "Error no healthcheck. %s", err.Error())
		return
	}

	hc.Status = "up"
	t.ResponseWithJSON(c, http.StatusOK, hc)
}
