package types

// ServiceImage contains the basic information for services declaration
type ServiceImage struct {
	ImageName   string   `json:"imageName" yaml:"imageName"`
	Args        []string `json:"args" yaml:"args"`
	Environment []string `json:"environment" yaml:"environment"`
	TTY         bool     `json:"tty" yaml:"tty"`
}
