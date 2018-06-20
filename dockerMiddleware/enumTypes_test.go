package dockerMiddleware

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnums(t *testing.T) {
	var service0 Service
	var service1 Service
	var service2 Service

	service0 = 0
	assert.Equal(t, "WORKER", service0.Name())

	service1 = 1
	assert.Equal(t, "MASTER", service1.Name())

	service2 = 3
	assert.Equal(t, "DISCOVERY", service2.Name())
}
