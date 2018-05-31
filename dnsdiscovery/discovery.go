package dnsdiscovery

import "github.com/dghubble/sling"

// Discovery function returns a string with the IPs discovered given a service name
// inside the same Docker Network
func Discovery(endpoint, serviceName string) (*[]string, error) {
	return request(endpoint, serviceName)
}

func request(endpoint, serviceName string) (*[]string, error) {
	ips := new([]string)
	var errr string

	_, err := sling.New().Get(endpoint+serviceName).Receive(ips, errr)

	return ips, err
}
