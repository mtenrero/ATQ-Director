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

// globalService returns a Global Service Mode
func globalService() *swarm.ServiceMode {
	global := swarm.GlobalService{}
	serviceMode := swarm.ServiceMode{
		Replicated: nil,
		Global:     &global,
	}

	return &serviceMode
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

// ServiceVIPS return the list of containers Virtual IPs of the given Service
func ServiceVIPS(serviceID string) (*[]swarm.EndpointVirtualIP, error) {
	client := getClient()

	service, _, err := client.ServiceInspectWithRaw(context.Background(), serviceID, types.ServiceInspectOptions{})
	if err != nil {
		return nil, err
	}

	vips := service.Endpoint.VirtualIPs

	return &vips, err
}

// ServiceModeDefault returns a Replicated Service Mode with 1 replica by default
func ServiceModeDefault() *swarm.ServiceMode {

	var amountReplicas = uint64(1)

	replicated := swarm.ReplicatedService{
		Replicas: &amountReplicas,
	}

	serviceMode := swarm.ServiceMode{
		Replicated: &replicated,
	}
	return &serviceMode
}

// ServiceDetails return the Service details
func ServiceDetails(serviceID string) (*swarm.Service, error) {
	client := getClient()

	serviceInspect, _, err := client.ServiceInspectWithRaw(context.Background(), serviceID, types.ServiceInspectOptions{})
	if err != nil {
		return nil, err
	}
	return &serviceInspect, nil
}

// ServiceAttachedNetworkID finds the network attached to a Service and return its attached netwrok
func ServiceAttachedNetworkID(serviceID string) (*string, error) {

	endpoint, err := ServiceVIPS(serviceID)
	if err != nil {
		return nil, err
	}

	if len(*endpoint) != 1 {
		return nil, err
	}

	var networkID string

	for _, vip := range *endpoint {
		networkID = vip.NetworkID
	}

	return &networkID, nil
}
