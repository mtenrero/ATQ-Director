package persistance

import "time"

// UID generates a timestamp to uint64
func UID() int64 {
	return time.Now().UnixNano()
}
