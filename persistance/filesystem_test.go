package persistance

import (
	"os"
	"testing"
)

func TestFirstRun(t *testing.T) {
	os.RemoveAll("./storage")

	p, err := InitPersistance(TestingDBPath, ".", "localhost")

	if err != nil {
		t.Error(err)
	}

	defer p.ClosePersistance()
}
