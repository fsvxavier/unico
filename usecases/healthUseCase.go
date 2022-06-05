package usecases

import (
	"github.com/fsvxavier/unico/interfaces"
	"github.com/fsvxavier/unico/models"
)

type healthcheckUseCase struct {
	healthcheckRepo interfaces.HealthCheckRepository
}

//NewHealthCheckUseCase will create new an healthcheckUsecase object representation of usecase.HealthCheckUsecase interface
func NewHealthCheckUseCase(h interfaces.HealthCheckRepository) interfaces.HealthCheckUseCase {
	return &healthcheckUseCase{
		healthcheckRepo: h,
	}
}

func (h *healthcheckUseCase) Check() (*models.HealthCheck, error) {
	res, err := h.healthcheckRepo.Check()
	if err != nil {
		return nil, err
	}
	return res, nil
}
