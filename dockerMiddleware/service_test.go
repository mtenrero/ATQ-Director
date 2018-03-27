package dockerMiddleware

import (
	"testing"

	"github.com/mtenrero/ATQ-Director/types"
)

func TestCreateService(t *testing.T) {
	ensureSwarm()

	RemoveNetwork("TEST_SERVICE")

	serviceImage := types.ServiceImage{
		ImageName: "hello-world",
		TTY:       true,
	}

	service, err := ComposeService(&serviceImage, "TEST_SERVICECREATION", "/tmp")
	if err != nil {
		t.Errorf("Error creating service : %s", err)
	}

	defer RemoveService(service.ID)
	defer RemoveNetwork("TEST_SERVICECREATION")
}

func TestRemoveService(t *testing.T) {
	ensureSwarm()

	RemoveNetwork("TEST_SERVICEDELETE")

	serviceImage := types.ServiceImage{
		ImageName: "hello-world",
		TTY:       true,
	}

	service, err := ComposeService(&serviceImage, "TEST_SERVICEDELETE", "/tmp")
	if err != nil {
		t.Errorf("Error creating service : %s", err)
	}

	errRemove := RemoveService(service.ID)
	if errRemove != nil {
		t.Error(errRemove)
	}

	defer RemoveService(service.ID)
	defer RemoveNetwork("TEST_SERVICEDELETE")

}
