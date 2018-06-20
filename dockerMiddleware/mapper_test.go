package dockerMiddleware

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateMountsMapper(t *testing.T) {
	mountSrc := "/mount"
	mounts := CreateMounts(&mountSrc, "ALIAS")

	assert.Len(t, mounts, 1, "Returned mounts are more than expected")
	assert.Equal(t, "/mount", mounts[0].Source)
}

func TestAttachNetworkMap(t *testing.T) {
	networkAttachment := AttachNetworkMap("ALIAS", "netID")

	assert.Equal(t, "ALIAS", networkAttachment.Aliases[0])
	assert.Equal(t, "netID", networkAttachment.Target)
}

func TestComposeServiceMapper(t *testing.T) {

}
