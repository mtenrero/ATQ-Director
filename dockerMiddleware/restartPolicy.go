package dockerMiddleware

import "docker.io/go-docker/api/types/swarm"

func restartPolicyNone() *swarm.RestartPolicy {

	restartPolicy := swarm.RestartPolicy{
		Condition: swarm.RestartPolicyConditionNone,
	}

	return &restartPolicy
}
