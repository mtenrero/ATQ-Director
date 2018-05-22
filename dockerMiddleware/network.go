package dockerMiddleware

import (
	"context"

	"docker.io/go-docker/api/types"
	"docker.io/go-docker/api/types/network"
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

// NetworkContainerPeers returns a list containing all Containers using the network and its Virtual IP
func NetworkContainerPeers(networkID string) (*[]network.PeerInfo, error) {

	client := getClient()

	network, err := client.NetworkInspect(context.Background(), networkID, types.NetworkInspectOptions{})
	if err != nil {
		return nil, err
	}

	var peers = network.Peers
	return &peers, nil
}

// NetworkVIPs returns the string list containing all Virtual IPs of a given Network
func NetworkVIPs(containerID string) (*[]string, error) {
	ips := make([]string, 0)

	peers, err := NetworkContainerPeers(containerID)
	if err != nil {
		return nil, err
	}

	for _, ip := range *peers {
		ips = append(ips, ip.IP)
	}

	return &ips, nil
}
