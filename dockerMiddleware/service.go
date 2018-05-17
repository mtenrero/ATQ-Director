package dockerMiddleware

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

// RemoveService deletes the service specified
func RemoveService(serviceID string) error {
	client := getClient()

	return client.ServiceRemove(context.Background(), serviceID)
}

// replicatedService returns a Service Mode configurated with the given replicas amount
func replicatedService(replicas int) *swarm.ServiceMode {
	ureplicas := uint64(replicas)
	replicatedService := swarm.ReplicatedService{
		Replicas: &ureplicas,
	}

	serviceMode := swarm.ServiceMode{
		Replicated: &replicatedService,
		Global:     nil,
	}

	return &serviceMode
}
