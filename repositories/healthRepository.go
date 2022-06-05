package repositories

import (
	"github.com/fsvxavier/unico/interfaces"
	"github.com/fsvxavier/unico/models"
)

type healthCheckRepository struct {
}

// NewHealthCheckRepository will create an object that represent the healthCheck.Repository interface
func NewHealthCheckRepository() interfaces.HealthCheckRepository {
	return &healthCheckRepository{}
}

func (m *healthCheckRepository) Check() (*models.HealthCheck, error) {
	a := &models.HealthCheck{}
	dbUp := "up"
	a.DbUP = dbUp
	return a, nil
}
