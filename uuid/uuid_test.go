package uuid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUUIDAlias(t *testing.T) {
	generated := AppendAlias("TEST")

	assert.Contains(t, generated, "TEST_")
}

func TestUUIDAliasString(t *testing.T) {
	generated := AppendAliasString("TEST", "VOLUME")

	assert.Contains(t, generated, "TEST_VOLUME_")
}
