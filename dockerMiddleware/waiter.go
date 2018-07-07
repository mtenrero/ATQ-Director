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
func VIPSWaiter(globalAlias, serviceName string, replicas int, timeout int, service Service, host string) error {

	var vipsExpected = replicas

	tick := time.Tick(500 * time.Millisecond)

	timer := time.NewTimer(time.Duration(timeout) * time.Second)
	defer timer.Stop()

	for {
		select {
		case <-tick:
			vips, _ := dnsdiscovery.Discovery("http://"+host+":9090/api/", globalAlias+"_"+serviceName+service.Name())

			vipsAmount := len(*vips)

			if vipsAmount == vipsExpected {
				return nil
			}

		case <-timer.C:
			return errors.New("Timed out waiting for all VIPS. Containers are not ready")
		}
	}
}

func injectVIPsIntoService(globalAlias, serviceName string, service *app.ServicePayload) (*app.ServicePayload, error) {
	vips, err := dnsdiscovery.Discovery("http://localhost:9090/api/", globalAlias+"_"+serviceName+"WORKER")
	if err != nil {
		return nil, err
	}

	newService := app.ServicePayload{
		Alias:       service.Alias,
		Args:        service.Args,
		Fileid:      service.Fileid,
		Image:       service.Image,
		Replicas:    service.Replicas,
		Tty:         service.Tty,
		Environment: service.Environment,
	}

	csvVips := strings.Join(*vips, ",")

	newService.Environment = append(newService.Environment, "WORKER_CSV_VIPS="+csvVips)
	newService.Environment = append(newService.Environment, "REMOTES="+csvVips)

	return &newService, nil
}
