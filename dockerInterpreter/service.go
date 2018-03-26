package dockerInterpreter

import (
	"context"

	"docker.io/go-docker/api/types"
	"docker.io/go-docker/api/types/swarm"
)

// CreateService creates a new service in the cluster with the given parameters
func CreateService(serviceSpec swarm.ServiceSpec) (types.ServiceCreateResponse, error) {
	client := getClient()

	serviceCreateResponse, err := client.ServiceCreate(context.Background(), serviceSpec, types.ServiceCreateOptions{})

	return serviceCreateResponse, err
}
