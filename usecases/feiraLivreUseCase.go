package usecases

import (
	"errors"

	"github.com/fsvxavier/unico/interfaces"
	"github.com/fsvxavier/unico/models"
)

type FeiraLivreUseCase struct {
	FeiraLivreRepo interfaces.FeiraLivreRepository
}

//NewFeiraLivreUseCase will create a new a arquivoPostgresUseCase object representation of usecase.ArquivoPostgresUseCase interfaces
func NewFeiraLivreUseCase(i interfaces.FeiraLivreRepository) interfaces.FeiraLivreUseCase {
	return &FeiraLivreUseCase{
		FeiraLivreRepo: i,
	}
}

// GetAllByIds ...
func (u *FeiraLivreUseCase) GetAllByIds(ids string) ([]*models.FeiraLivre, error) {
	record, err := u.FeiraLivreRepo.GetAllByIds(ids)
	if err != nil {
		return nil, err
	}
	return record, err
}

// GetByID ...
func (u *FeiraLivreUseCase) GetByID(id int64) ([]*models.FeiraLivre, error) {
	record, err := u.FeiraLivreRepo.GetByID(id)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return record, err
}

// CreateFeiraLivre ...
func (u *FeiraLivreUseCase) CreateFeiraLivre(iFeiras *models.FeiraLivre) (*models.Response, error) {
	ret, err := u.FeiraLivreRepo.CreateFeiraLivre(iFeiras)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return ret, nil
}

// UpdateFeiraLivre ...
func (u *FeiraLivreUseCase) UpdateFeiraLivre(iFeiras *models.FeiraLivre) (*models.Response, error) {
	ret, err := u.FeiraLivreRepo.UpdateFeiraLivre(iFeiras)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return ret, nil
}

// SearchFeiraLivre ...
func (u *FeiraLivreUseCase) SearchFeiraLivre(sFeiras *models.SearchFeiraLivre) ([]*models.FeiraLivre, error) {
	ret, err := u.FeiraLivreRepo.SearchFeiraLivre(sFeiras)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return ret, nil
}

// DeleteFeiraLivre ...
func (u *FeiraLivreUseCase) DeleteFeiraLivre(id int64) (*models.Response, error) {
	ret, err := u.FeiraLivreRepo.DeleteFeiraLivre(id)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return ret, nil
}
