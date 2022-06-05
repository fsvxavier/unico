package interfaces

import (
	"github.com/fsvxavier/unico/models"
)

//HealthCheckRepository ...
type HealthCheckRepository interface {
	Check() (*models.HealthCheck, error)
}

//FeiraLivreRepository ...
type FeiraLivreRepository interface {
	GetByID(id int64) ([]*models.FeiraLivre, error)
	GetAllByIds(ids string) ([]*models.FeiraLivre, error)
	CreateFeiraLivre(iFeiras *models.FeiraLivre) (*models.Response, error)
	UpdateFeiraLivre(iFeiras *models.FeiraLivre) (*models.Response, error)
	SearchFeiraLivre(sFeiras *models.SearchFeiraLivre) ([]*models.FeiraLivre, error)
	DeleteFeiraLivre(id int64) (*models.Response, error)
}
