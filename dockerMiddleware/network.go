package dockerMiddleware

import (
	"context"

	"docker.io/go-docker/api/types"
)

// CreateOverlayNetwork creates a new Overlay Network in the Swarm
func CreateOverlayNetwork(id string) (types.NetworkCreateResponse, error) {
	client := getClient()

	networkCreateResponse, err := client.NetworkCreate(
		context.Background(),
		id,
		overlayNetworkSpec(id),
	)

	return networkCreateResponse, err
}

// RemoveNetwork removes specified network from the cluster
func RemoveNetwork(id string) error {
	client := getClient()

	return client.NetworkRemove(context.Background(), id)
}

// overlayNetworkSpec returns a predefined Spec for a Docker Overlay Network
func overlayNetworkSpec(id string) types.NetworkCreate {

	options := make(map[string]string)
	options["com.docker.network.driver.overlay.vxlanid_list"] = "4096"

	labels := make(map[string]string)
	labels["com.docker.network.driver.overlay.vxlanid_list"] = id

	return types.NetworkCreate{
		Driver: "overlay",
		Scope:  "swarm",
		Labels: labels,
	}
}
