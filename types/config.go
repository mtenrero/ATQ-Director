package types

// ControllerConfig contains all the configuration available for Controller Mode
type ControllerConfig struct {
	Port   string          `json:"port" yaml:"port"`
	Images []*ServiceImage `json:"images" yaml:"images"`
}
