package handlers_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	handlers "github.com/fsvxavier/unico/api/handlers"
	"github.com/fsvxavier/unico/models"
	"github.com/fsvxavier/unico/usecases"
	mockUse "github.com/fsvxavier/unico/usecases/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func TestFeiraLivre(t *testing.T) {
	var mockFeiraLivre models.FeiraLivre
	mockUCase := new(mockUse.NewFeiraLivreUseCase)
	mockUCase.On("Check").Return(&mockFeiraLivre, nil)
	req, err := http.NewRequest(http.MethodGet, "/health", nil)
	if err != nil {
		t.Fatal(err)
	}
	assert.NoError(t, err)
	rr := httptest.NewRecorder()

	r := gin.New()
	v1 := r.Group("")
	FeiraLivreUseCase := usecases.NewFeiraLivreUseCase(mockUCase)
	handlers.NewFeiraLivreHTTPHandler(v1, FeiraLivreUseCase)
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	mockUCase.AssertExpectations(t)
}

func TestFeiraLivreError(t *testing.T) {
	var mockFeiraLivre models.FeiraLivre
	mockUCase := new(mockUse.NewFeiraLivreUseCase)
	mockUCase.On("Check").Return(&mockFeiraLivre, errors.New(""))
	req, err := http.NewRequest(http.MethodGet, "/health", nil)
	if err != nil {
		t.Fatal(err)
	}
	assert.NoError(t, err)
	rr := httptest.NewRecorder()

	r := gin.New()
	v1 := r.Group("")
	FeiraLivreUseCase := usecases.NewFeiraLivreUseCase(mockUCase)
	handlers.NewFeiraLivreHTTPHandler(v1, FeiraLivreUseCase)
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	mockUCase.AssertExpectations(t)
}
