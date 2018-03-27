package dockerMiddleware

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateOverlayNetwork(t *testing.T) {
	RemoveNetwork("testOverlay")
	response, err := CreateOverlayNetwork("testOverlay")

	RemoveNetwork("testOverlay")
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "", response.Warning)
}

func TestRemoveNetwork(t *testing.T) {
	RemoveNetwork("testRemove")

	_, err := CreateOverlayNetwork("testRemove")
	if err != nil {
		t.Errorf("Error creating test Network : %s", err)
	}

	errCreate := RemoveNetwork("testRemove")
	if err != nil {
		t.Errorf("Error deleting network : %s", errCreate)
	}
}
