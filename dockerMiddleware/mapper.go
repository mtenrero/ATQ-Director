package dockerMiddleware

import (
	"docker.io/go-docker/api/types"
	"docker.io/go-docker/api/types/mount"
	"docker.io/go-docker/api/types/swarm"
	atqTypes "github.com/mtenrero/ATQ-Director/types"
)

// ContainerSpecMapper initializes a new Docker ContainerSpec type with the predefined serviceImage
func ContainerSpecMapper(serviceImage *atqTypes.ServiceImage, alias string, mounts []mount.Mount) *swarm.ContainerSpec {

	aliasMap := make(map[string]string)
	aliasMap["atq_alias"] = alias

	spec := swarm.ContainerSpec{
		Image:  serviceImage.ImageName,
		TTY:    serviceImage.TTY,
		Mounts: mounts,
		Labels: aliasMap,
	}

	return &spec
}

// TaskSpecMapper maps the service configuration
func TaskSpecMapper(containerSpec *swarm.ContainerSpec, networkAttachConfig []swarm.NetworkAttachmentConfig) *swarm.TaskSpec {

	return &swarm.TaskSpec{
		ContainerSpec: containerSpec,
		Networks:      networkAttachConfig,
	}

}

// CreateMounts configures the mounts given an alias and a path
func CreateMounts(path *string, alias string) []mount.Mount {
	var mounts []mount.Mount

	if path != nil {

		mount := mount.Mount{
			Type:   "bind",
			Source: *path,
			Target: "/atq/data",
		}

		return append(mounts, mount)
	}

	return nil
}

// CreateNetworkMap creates a new specific network for the service
func CreateNetworkMap(alias string) (*swarm.NetworkAttachmentConfig, error) {

	aliases := []string{alias}

	network, netError := CreateOverlayNetwork(alias)

	networkAttachConfig := swarm.NetworkAttachmentConfig{
		Target:  network.ID,
		Aliases: aliases,
	}

	return &networkAttachConfig, netError
}

// AttachNetworkMap creates a networkAttachmentConfig with the existing specified network
func AttachNetworkMap(alias string, networkID string) *swarm.NetworkAttachmentConfig {
	aliases := []string{alias}

	networkAttachConfig := swarm.NetworkAttachmentConfig{
		Target:  networkID,
		Aliases: aliases,
	}

	return &networkAttachConfig
}

// ComposeService Maps the values to a new service
func ComposeService(serviceImage *atqTypes.ServiceImage, globalAlias, alias string, path *string, mode *swarm.ServiceMode, networkID *string) (*types.ServiceCreateResponse, error) {

	mounts := CreateMounts(path, alias)

	containerSpec := ContainerSpecMapper(serviceImage, alias, mounts)

	var networkSpec *swarm.NetworkAttachmentConfig

	if networkID != nil {
		networkSpec = AttachNetworkMap(globalAlias, *networkID)
	} else {
		var netError error
		networkSpec, netError = CreateNetworkMap(globalAlias)

		if netError != nil {
			return nil, netError
		}
	}

	task := TaskSpecMapper(containerSpec, []swarm.NetworkAttachmentConfig{*networkSpec})

	annotations := swarm.Annotations{
		Name: alias,
	}

	var serviceMode swarm.ServiceMode

	if mode == nil {
		serviceMode = *ServiceModeDefault()
	} else {
		serviceMode = *mode
	}

	serviceSpec := swarm.ServiceSpec{
		Annotations:  annotations,
		TaskTemplate: *task,
		Mode:         serviceMode,
	}

	service, serviceErr := CreateService(serviceSpec)

	if serviceErr != nil {
		return nil, serviceErr
	}

	return &service, nil
}
