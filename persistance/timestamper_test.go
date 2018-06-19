package persistance

import (
	"testing"
	"time"
)

func TestUIDTimestamper(t *testing.T) {
	timestamp := time.Now().UnixNano()
	uid := UID()

	if timestamp > uid {
		t.Error("Error in UID generation")
	}

}
