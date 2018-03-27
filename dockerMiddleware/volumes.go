package dockerMiddleware

import (
	"context"

	"docker.io/go-docker/api/types"
	"docker.io/go-docker/api/types/volume"
	"github.com/mtenrero/ATQ-Director/uuid"
)

// CreateVolumeWithData declares a new volume mapped to a given directory
func CreateVolumeWithData(alias, path string) (types.Volume, error) {
	client := getClient()

	driveOptions := make(map[string]string)

	driveOptions["o"] = "bind"
	driveOptions["device"] = path

	volumeOptions := volume.VolumesCreateBody{
		Driver:     "local",
		DriverOpts: driveOptions,
		Labels:     make(map[string]string),
		Name:       uuid.AppendAlias(alias),
	}

	volume, err := client.VolumeCreate(context.Background(), volumeOptions)

	return volume, err
}

// RemoveVolume removes a given Volume
func RemoveVolume(volumeID string, force bool) error {
	client := getClient()

	return client.VolumeRemove(context.Background(), volumeID, force)
}
