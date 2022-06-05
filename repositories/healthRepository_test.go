package repositories_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/fsvxavier/unico/repositories"
)

func TestHealthCheck(t *testing.T) {

	a := repositories.NewHealthCheckRepository()
	h, err := a.Check()
	assert.NoError(t, err)
	assert.NotNil(t, h)
}
