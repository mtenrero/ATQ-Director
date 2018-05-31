package dockerMiddleware

// Service holds the different types available
type Service int

const (
	Worker     Service = 0
	Master     Service = 1
	Discoverer Service = 3
)

// Returns the Service Type Name
func (service Service) Name() string {
	switch service {
	case 0:
		return "WORKER"
	case 1:
		return "MASTER"
	case 3:
		return "DISCOVERY"
	}
	return ""
}
