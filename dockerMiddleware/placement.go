package dockerMiddleware

import "docker.io/go-docker/api/types/swarm"

func placementManager() *swarm.Placement {

	constraints := []string{"node.role == manager"}

	placement := swarm.Placement{
		Constraints: constraints,
	}

	return &placement
}
