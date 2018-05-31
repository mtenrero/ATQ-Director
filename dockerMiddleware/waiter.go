package dockerMiddleware

import (
	"errors"
	"log"
	"strings"
	"time"

	"github.com/mtenrero/ATQ-Director/app"

	"github.com/mtenrero/ATQ-Director/dnsdiscovery"
)

// ServiceHostWaiter await for a given amount of VIPS specified in the parameter.
// Timeout in seconds
// Only for service deployed on the current host
func ServiceHostWaiter(serviceID string, replicas int, timeout int) error {

	var vipsExpected = replicas

	timeoutchan := time.After(time.Duration(timeout) * time.Second)
	tick := time.Tick(500 * time.Millisecond)

	for {
		select {
		case <-timeoutchan:
			return errors.New("Timed out waiting for all VIPS. Containers are not ready")
		case <-tick:
			log.Println("tick")
			// Obtain Service Attached Network
			networkID, err := ServiceAttachedNetworkID(serviceID)
			if err != nil {
				return err
			}

			log.Println(*networkID)

			if networkID != nil {
				vips, err := NetworkVIPs(*networkID)
				if err != nil {
					return err
				}

				vipsAmount := len(*vips)
				log.Println(vipsAmount)

				if vipsAmount == vipsExpected {
					return nil
				}
			}
		}
	}
}

// VIPSWaiter await for a given amount of VIPS specified in the parameter.
// Timeout in seconds
func VIPSWaiter(globalAlias, serviceName string, replicas int, timeout int) error {

	var vipsExpected = replicas

	timeoutchan := time.After(time.Duration(timeout) * time.Second)
	tick := time.Tick(500 * time.Millisecond)

	for {
		select {
		case <-timeoutchan:
			return errors.New("Timed out waiting for all VIPS. Containers are not ready")
		case <-tick:
			log.Println("tick")
			// Obtain Service Attached Network
			vips, err := dnsdiscovery.Discovery("http://localhost:9090/api/", globalAlias+"_"+serviceName+"WORKER")
			if err != nil {
				return err
			}

			vipsAmount := len(*vips)
			log.Println(vipsAmount)

			if vipsAmount == vipsExpected {
				return nil
			}
		}
	}
}

func injectVIPsIntoService(globalAlias, serviceName string, service *app.ServicePayload) (*app.ServicePayload, error) {
	vips, err := dnsdiscovery.Discovery("http://localhost:9090/api/", globalAlias+"_"+serviceName+"WORKER")
	if err != nil {
		return nil, err
	}

	newService := *service

	csvVips := strings.Join(*vips, ",")

	newService.Environment = append(newService.Environment, "WORKER_CSV_VIPS="+csvVips)

	return &newService, nil
}
