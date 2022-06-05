package interfaces

import (
	"github.com/fsvxavier/unico/models"
)

//HealthCheckUseCase ...
type HealthCheckUseCase interface {
	Check() (*models.HealthCheck, error)
}

//feiraLivreUseCase ...
type FeiraLivreUseCase interface {
	GetByID(id int64) ([]*models.FeiraLivre, error)
	GetAllByIds(ids string) ([]*models.FeiraLivre, error)
	CreateFeiraLivre(iFeiras *models.FeiraLivre) (*models.Response, error)
	UpdateFeiraLivre(iFeiras *models.FeiraLivre) (*models.Response, error)
	SearchFeiraLivre(sFeiras *models.SearchFeiraLivre) ([]*models.FeiraLivre, error)
	DeleteFeiraLivre(id int64) (*models.Response, error)
}
