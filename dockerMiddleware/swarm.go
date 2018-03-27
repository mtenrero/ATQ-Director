package dockerMiddleware

import (
	"context"

	"docker.io/go-docker/api/types/swarm"
)

// SwarmInspect gives the information about the Swarm
func SwarmInspect() (swarm.Swarm, error) {
	client := getClient()

	return client.SwarmInspect(context.Background())
}

// IsSwarmNode returns an error if the node is not a swarm node member
func IsSwarmNode() bool {
	client := getClient()
	_, error := client.SwarmInspect(context.Background())

	if error == nil {
		return true
	}
	return false
}
