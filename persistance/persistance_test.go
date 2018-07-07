package persistance

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

const TestingDBPath = "./storage/dev_testing.atq"

func TestBasicDatastore(t *testing.T) {
	absolutePath, _ := filepath.Abs(TestingDBPath)
	p, err := InitPersistance(absolutePath, "", "localhost")
	if err != nil {
		t.Error(err)
	}

	p.delete("TEST")

	err = p.store("TEST", "TOST")
	if err != nil {
		t.Error(err)
	}

	value, err := p.read("TEST")
	if err != nil {
		t.Error(err)
	}

	assert.EqualValues(t, "TOST", value, "Value Mismatch!")

	defer p.delete("TEST")
	defer p.ClosePersistance()
}
