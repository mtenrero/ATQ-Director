package dockerMiddleware

// Service holds the different types available
type Service int

const (
	Worker Service = 0
	Master Service = 1
)

// Returns the Service Type Name
func (service Service) Name() string {
	switch service {
	case 0:
		return "WORKER"
	case 1:
		return "MASTER"
	}
	return ""
}
