package main

import (
	"os"

	"github.com/mtenrero/ATQ-Director/dockerMiddleware"
	"github.com/sirupsen/logrus"
)

var LOGGER *logrus.Logger

func init() {
	logrus.SetOutput(os.Stdout)
	LOGGER = logrus.New()

}

func main() {
	swarmChecks()
}

func swarmChecks() {
	if !dockerMiddleware.IsSwarmNode() {
		_, error := dockerMiddleware.SwarmInspect()
		LOGGER.Error(error)
	}
}
