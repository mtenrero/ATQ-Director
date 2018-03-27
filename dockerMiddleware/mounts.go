package dockerMiddleware

import "docker.io/go-docker/api/types/mount"

// AddMount Add the given Mount to Slice
func AddMount(slice []mount.Mount, mount mount.Mount) []mount.Mount {
	return append(slice, mount)
}
