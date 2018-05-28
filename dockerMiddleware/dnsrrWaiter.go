package dockerMiddleware

import (
	"errors"
	"log"
	"time"
)

// WaitDNS waits for a given amount of IPs discovered in a hostname
// Requieres tenrero/dnsrr-discovery-api deployed and reachable in global mode
func WaitDNS(hostname string, replicas int) error {
	timeoutchan := time.After(time.Duration(GlobalTimeoutSeconds) * time.Second)
	tick := time.Tick(500 * time.Millisecond)

	for {
		select {
		case <-timeoutchan:
			return errors.New("Timed out waiting for all VIPS. Containers are not ready")
		case <-tick:
			log.Println("tick")

		}
	}

	return nil
}
