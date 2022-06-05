package usecases_test

import (
	"errors"
	"testing"

	"github.com/fsvxavier/unico/models"
	mocksRepo "github.com/fsvxavier/unico/repositories/mocks"
	"github.com/fsvxavier/unico/usecases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

//TestFeiraLivreGetByID ...
func TestFeiraLivreGetByID(t *testing.T) {
	mockFeiraLivre := []*models.FeiraLivre{
		{
			ID:         1,
			Longi:      0,
			Lat:        0,
			Setcens:    0,
			Areap:      0,
			Coddist:    0,
			Distrito:   "",
			Codsubpref: 0,
			Subprefe:   "",
			Regiao5:    "",
			Regiao8:    "",
			NomeFeira:  "",
			Registro:   "",
			Logradouro: "",
			Numero:     "",
			Bairro:     "",
			Referencia: "",
		},
	}
	mockFeiraLivreRepo := new(mocksRepo.NewFeiraLivreRepository)

	t.Run("sucess", func(t *testing.T) {

		mockFeiraLivreRepo.On("GetByID", int64(1)).Return(mockFeiraLivre, nil)
		u := usecases.NewFeiraLivreUseCase(mockFeiraLivreRepo)
		p, err := u.GetByID(int64(1))

		assert.NoError(t, err)
		assert.NotNil(t, p)
		mockFeiraLivreRepo.AssertExpectations(t)
	})

	t.Run("error-getbyid", func(t *testing.T) {
		mockFeiraLivreRepo.On("GetByID", mock.AnythingOfType("int64")).Return(nil, errors.New("Unexpected Error"))
		mockFeiraLivreRepo.On("GetItemByID", mock.AnythingOfType("int64")).Return(nil, errors.New("Usnexpected Error"))
		u := usecases.NewFeiraLivreUseCase(mockFeiraLivreRepo)
		p, err := u.GetByID(0)
		assert.Error(t, err)
		assert.Nil(t, p)
	})
}

//TestFeiraLivreGetAllByIds...
func TestFeiraLivreGetAllByIds(t *testing.T) {

	mockFeiraLivreRepo := new(mocksRepo.NewFeiraLivreRepository)

	mockListFeiraLivre := []*models.FeiraLivre{
		{
			ID:         1,
			Longi:      0,
			Lat:        0,
			Setcens:    0,
			Areap:      0,
			Coddist:    0,
			Distrito:   "",
			Codsubpref: 0,
			Subprefe:   "",
			Regiao5:    "",
			Regiao8:    "",
			NomeFeira:  "",
			Registro:   "",
			Logradouro: "",
			Numero:     "",
			Bairro:     "",
			Referencia: "",
		},
		{
			ID:         2,
			Longi:      0,
			Lat:        0,
			Setcens:    0,
			Areap:      0,
			Coddist:    0,
			Distrito:   "",
			Codsubpref: 0,
			Subprefe:   "",
			Regiao5:    "",
			Regiao8:    "",
			NomeFeira:  "",
			Registro:   "",
			Logradouro: "",
			Numero:     "",
			Bairro:     "",
			Referencia: "",
		},
		{
			ID:         3,
			Longi:      0,
			Lat:        0,
			Setcens:    0,
			Areap:      0,
			Coddist:    0,
			Distrito:   "",
			Codsubpref: 0,
			Subprefe:   "",
			Regiao5:    "",
			Regiao8:    "",
			NomeFeira:  "",
			Registro:   "",
			Logradouro: "",
			Numero:     "",
			Bairro:     "",
			Referencia: "",
		},
	}

	t.Run("sucess", func(t *testing.T) {
		mockFeiraLivreRepo.On("GetAllByIds", "1,2,3").Return(mockListFeiraLivre, nil)

		u := usecases.NewFeiraLivreUseCase(mockFeiraLivreRepo)
		list, err := u.GetAllByIds("1,2,3")

		assert.NoError(t, err)
		assert.Len(t, list, len(mockListFeiraLivre))
		mockFeiraLivreRepo.AssertExpectations(t)
	})
}

//TestFeiraLivreGetAllByIdsError...
func TestFeiraLivreGetAllByIdsError(t *testing.T) {

	mockFeiraLivreRepo := new(mocksRepo.NewFeiraLivreRepository)

	t.Run("error", func(t *testing.T) {
		mockFeiraLivreRepo.On("GetAllByIds", "-1").Return(nil, errors.New("error"))

		u := usecases.NewFeiraLivreUseCase(mockFeiraLivreRepo)
		_, err := u.GetAllByIds("-1")

		assert.Error(t, err)
	})
}

//TestFeiraLivreGetByID ...
func TestCreateFeiraLivre(t *testing.T) {

	mockResponse := models.Response{}

	mockFeiraLivre := models.FeiraLivre{}
	mockFeiraLivreRepo := new(mocksRepo.NewFeiraLivreRepository)

	t.Run("sucess", func(t *testing.T) {

		mockFeiraLivreRepo.On("CreateFeiraLivre", &mockFeiraLivre).Return(&mockResponse, nil)
		u := usecases.NewFeiraLivreUseCase(mockFeiraLivreRepo)
		p, err := u.CreateFeiraLivre(&mockFeiraLivre)

		assert.NoError(t, err)
		assert.NotNil(t, p)
		mockFeiraLivreRepo.AssertExpectations(t)
	})
}

//UpdateFeiraLivre...
func TestUpdateFeiraLivre(t *testing.T) {

	mockResponse := models.Response{}

	mockFeiraLivre := models.FeiraLivre{}
	mockFeiraLivreRepo := new(mocksRepo.NewFeiraLivreRepository)

	t.Run("sucess", func(t *testing.T) {

		mockFeiraLivreRepo.On("UpdateFeiraLivre", &mockFeiraLivre).Return(&mockResponse, nil)
		u := usecases.NewFeiraLivreUseCase(mockFeiraLivreRepo)
		p, err := u.UpdateFeiraLivre(&mockFeiraLivre)

		assert.NoError(t, err)
		assert.NotNil(t, p)
		mockFeiraLivreRepo.AssertExpectations(t)
	})
}

//SearchFeiraLivre...
func TestSearchFeiraLivre(t *testing.T) {

	var mockResponse []*models.FeiraLivre

	mockFeiraLivre := models.SearchFeiraLivre{}
	mockFeiraLivreRepo := new(mocksRepo.NewFeiraLivreRepository)

	t.Run("sucess", func(t *testing.T) {

		mockFeiraLivreRepo.On("SearchFeiraLivre", &mockFeiraLivre).Return(mockResponse, nil)
		u := usecases.NewFeiraLivreUseCase(mockFeiraLivreRepo)
		_, err := u.SearchFeiraLivre(&mockFeiraLivre)

		assert.NoError(t, err)
		mockFeiraLivreRepo.AssertExpectations(t)
	})
}

//UpdateFeiraLivre...
func TestDeleteFeiraLivre(t *testing.T) {

	mockResponse := models.Response{}

	mockFeiraLivreRepo := new(mocksRepo.NewFeiraLivreRepository)

	id := int64(1)

	t.Run("sucess", func(t *testing.T) {

		mockFeiraLivreRepo.On("DeleteFeiraLivre", id).Return(&mockResponse, nil)
		u := usecases.NewFeiraLivreUseCase(mockFeiraLivreRepo)
		p, err := u.DeleteFeiraLivre(id)

		assert.NoError(t, err)
		assert.NotNil(t, p)
		mockFeiraLivreRepo.AssertExpectations(t)
	})
}
