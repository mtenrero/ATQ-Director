package dockerMiddleware

import (
	docker "docker.io/go-docker"
)

// getClient returns the Docker Client
func getClient() *docker.Client {
	cli, err := docker.NewEnvClient()
	if err != nil {
		panic(err)
	}

	return cli
}
