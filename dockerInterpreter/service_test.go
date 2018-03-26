package dockerInterpreter

import (
	"testing"

	"github.com/mtenrero/ATQ-Director/types"
)

func TestCreation(t *testing.T) {

	RemoveNetwork("TEST_SERVICE")

	serviceImage := types.ServiceImage{
		ImageName: "mtenrero/jmeter",
		TTY:       true,
	}

	_, err := ComposeService(&serviceImage, "TEST_SERVICECREATION", "/tmp")
	if err != nil {
		t.Errorf("Error creating service : %s", err)
	}

	defer RemoveNetwork("TEST_SERVICECREATION")
}
