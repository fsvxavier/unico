package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/fsvxavier/unico/config"
	"github.com/fsvxavier/unico/interfaces"
	"github.com/fsvxavier/unico/models"
	"github.com/fsvxavier/unico/utils"
	"github.com/gin-gonic/gin"
)

// httpFeiraLivreHandler represent the httphandler for FeiraLivre
type httpFeiraLivreHandler struct {
	FLUsecase interfaces.FeiraLivreUseCase
}

// NewFeiraLivreHTTPHandler ...
func NewFeiraLivreHTTPHandler(r *gin.RouterGroup, us interfaces.FeiraLivreUseCase) {
	handler := &httpFeiraLivreHandler{
		FLUsecase: us,
	}
	r.PUT("/feiralivre/create", handler.CreateFeiraLivre)
	r.PUT("/feiralivre/update", handler.UpdateFeiraLivre)
	r.POST("/feiralivre/search", handler.SearchFeiraLivre)

	r.DELETE("/feiralivre/delete/", handler.DeleteFeiraLivre)
	r.GET("/feiralivre/getbyid/", handler.GetByID)
	r.GET("/feiralivre/getdiversbyids/", handler.GetDiversByiIds)
}

// @Summary FeiraLivre
// @Description FeiraLivre API
// @Param data body models.SearchFeiraLivre true "body request"
// @Failure 400 {object} models.ResponseError
// @Success 200 {object} models.FeiraLivre
// @Router /feiralivre/search [post]
func (h *httpFeiraLivreHandler) SearchFeiraLivre(c *gin.Context) {
	var t config.Config
	logger := new(utils.GenericLogger)
	logger.Module = "api"
	logger.GetLogger()

	payload, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logger.LogIt("ERROR", fmt.Sprintf("[handlers/Send] - Error on read parameters in request. Erro: %s", err.Error()), nil)
		t.RespondWithError(c, http.StatusInternalServerError, "", "")
		return
	}

	var outValuesDefault models.SearchFeiraLivre

	if err := json.Unmarshal([]byte(payload), &outValuesDefault); err != nil {
		logger.LogIt("ERROR", fmt.Sprintf("[handlers/Send] - Error on parse parameters of request. Erro: %s", err.Error()), nil)
		t.ResponseWithError(c, http.StatusBadRequest, "Error performing analysis of the submitted body: ", err.Error())
		return
	}

	hc, err := h.FLUsecase.SearchFeiraLivre(&outValuesDefault)

	if err != nil {
		logger.LogIt("ERROR", fmt.Sprintf("[handlers/Send] - Error on execute Search Handler. Erro: %s", err.Error()), nil)
		t.RespondWithError(c, http.StatusBadRequest, "Error no FeiraLivre. %s", err.Error())
		return
	}

	t.ResponseWithJSON(c, http.StatusOK, hc)
}

// @Summary FeiraLivre
// @Description FeiraLivre API
// @Accept  json
// @Produce  json
// @Param data body models.FeiraLivre true "body request"
// @Success 200 {object} models.ResponseSuccess
// @Failure 400 {object} models.ResponseError
// @Router /feiralivre/create [put]
func (h *httpFeiraLivreHandler) CreateFeiraLivre(c *gin.Context) {
	var t config.Config
	logger := new(utils.GenericLogger)
	logger.Module = "api"
	logger.GetLogger()

	payload, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logger.LogIt("ERROR", fmt.Sprintf("[handlers/Send] - Error on read parameters in request. Erro: %s", err.Error()), nil)
		t.RespondWithError(c, http.StatusInternalServerError, "", "")
		return
	}

	var outValuesDefault models.FeiraLivre

	if err := json.Unmarshal([]byte(payload), &outValuesDefault); err != nil {
		logger.LogIt("ERROR", fmt.Sprintf("[handlers/Send] - Error on parse parameters of request. Erro: %s", err.Error()), nil)
		t.ResponseWithError(c, http.StatusBadRequest, "Error performing analysis of the submitted body: ", err.Error())
		return
	}

	ret, err := h.FLUsecase.CreateFeiraLivre(&outValuesDefault)

	if err != nil {
		t.RespondWithError(c, http.StatusBadRequest, "Error no FeiraLivre. %s", err.Error())
		return
	}

	t.ResponseWithJSON(c, http.StatusOK, ret)
}

// @Summary FeiraLivre
// @Description FeiraLivre API
// @Accept  json
// @Produce  json
// @Param data body models.FeiraLivre true "body request"
// @Success 200 {object} models.ResponseSuccess
// @Failure 400 {object} models.ResponseError
// @Router /feiralivre/update [put]
func (h *httpFeiraLivreHandler) UpdateFeiraLivre(c *gin.Context) {
	var t config.Config
	logger := new(utils.GenericLogger)
	logger.Module = "api"
	logger.GetLogger()

	payload, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logger.LogIt("ERROR", fmt.Sprintf("[handlers/Send] - Error on read parameters in request. Erro: %s", err.Error()), nil)
		t.RespondWithError(c, http.StatusInternalServerError, "", "")
		return
	}

	var outValuesDefault models.FeiraLivre

	if err := json.Unmarshal([]byte(payload), &outValuesDefault); err != nil {
		logger.LogIt("ERROR", fmt.Sprintf("[handlers/Send] - Error on parse parameters of request. Erro: %s", err.Error()), nil)
		t.ResponseWithError(c, http.StatusBadRequest, "Error performing analysis of the submitted body: ", err.Error())
		return
	}

	ret, err := h.FLUsecase.UpdateFeiraLivre(&outValuesDefault)

	if err != nil {
		t.RespondWithError(c, http.StatusBadRequest, "Error no FeiraLivre. %s", err.Error())
		return
	}

	t.ResponseWithJSON(c, http.StatusOK, ret)
}

// @Summary FeiraLivre
// @Description FeiraLivre API
// @Param id query integer true "id to delete"
// @Failure 400 {object} models.ResponseError
// @Success 200 {object} models.FeiraLivre
// @Router /feiralivre/delete [delete]
func (h *httpFeiraLivreHandler) DeleteFeiraLivre(c *gin.Context) {
	var t config.Config
	logger := new(utils.GenericLogger)
	logger.Module = "api"
	logger.GetLogger()

	id, _ := strconv.Atoi(c.Query("id"))

	hc, err := h.FLUsecase.DeleteFeiraLivre(int64(id))

	if err != nil {
		logger.LogIt("ERROR", fmt.Sprintf("[handlers/Send] - Error on execute DeleteFeiraLivre Handler. Erro: %s", err.Error()), nil)
		t.RespondWithError(c, http.StatusBadRequest, "Error no FeiraLivre. %s", err.Error())
		return
	}

	t.ResponseWithJSON(c, http.StatusOK, hc)
}

// @Summary FeiraLivre
// @Description FeiraLivre API
// @Param id query integer true "id to search"
// @Failure 400 {object} models.ResponseError
// @Success 200 {object} models.FeiraLivre
// @Router /feiralivre/getbyid [get]
func (h *httpFeiraLivreHandler) GetByID(c *gin.Context) {
	var t config.Config
	logger := new(utils.GenericLogger)
	logger.Module = "api"
	logger.GetLogger()

	id, _ := strconv.Atoi(c.Query("id"))

	hc, err := h.FLUsecase.GetByID(int64(id))

	if err != nil {
		logger.LogIt("ERROR", fmt.Sprintf("[handlers/Send] - Error on execute GetByID Handler. Erro: %s", err.Error()), nil)
		t.RespondWithError(c, http.StatusBadRequest, "Error no FeiraLivre. %s", err.Error())
		return
	}

	t.ResponseWithJSON(c, http.StatusOK, hc)
}

// @Summary FeiraLivre
// @Description FeiraLivre API
// @Param ids query string true "ids to search 1-2-3"
// @Failure 400 {object} models.ResponseError
// @Success 200 {object} models.FeiraLivre
// @Router /feiralivre/getdiversbyids [get]
func (h *httpFeiraLivreHandler) GetDiversByiIds(c *gin.Context) {
	var t config.Config
	logger := new(utils.GenericLogger)
	logger.Module = "api"
	logger.GetLogger()

	id := strings.Replace(c.Query("ids"), "-", ",", -1)

	hc, err := h.FLUsecase.GetAllByIds(id)

	if err != nil {
		logger.LogIt("ERROR", fmt.Sprintf("[handlers/Send] - Error on execute GetAllByIds Handler. Erro: %s", err.Error()), nil)
		t.RespondWithError(c, http.StatusBadRequest, "Error no FeiraLivre. %s", err.Error())
		return
	}

	t.ResponseWithJSON(c, http.StatusOK, hc)
}
