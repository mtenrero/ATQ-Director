package dockerInterpreter

import (
	"context"
	"testing"

	"docker.io/go-docker/api/types/swarm"
)

func TestSwarmMember(t *testing.T) {

	memberBefore := IsSwarmNode()

	client := getClient()

	if !memberBefore {
		_, error := client.SwarmInit(context.Background(), swarm.InitRequest{
			ListenAddr: "0.0.0.0",
		})
		t.Log(error)
	}

	if !IsSwarmNode() {
		t.Error("NOT MEMBER OF A SWARM CLUSTER!")
	}

	if !memberBefore {
		client.SwarmLeave(context.Background(), false)
	}
}

func TestSwarmInspect(t *testing.T) {
	memberBefore := IsSwarmNode()

	client := getClient()

	if !memberBefore {
		_, error := client.SwarmInit(context.Background(), swarm.InitRequest{
			ListenAddr: "0.0.0.0",
		})
		t.Log(error)
	}

	_, error := SwarmInspect()

	if error != nil {
		t.Error(error)
	}

	if !memberBefore {
		client.SwarmLeave(context.Background(), false)
	}
}

func ensureSwarm() {
	memberBefore := IsSwarmNode()

	client := getClient()

	if !memberBefore {
		client.SwarmInit(context.Background(), swarm.InitRequest{
			ListenAddr: "0.0.0.0",
		})

	}
}
