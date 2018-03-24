package types

// ServiceImage contains the basic information for services declaration
type ServiceImage struct {
	ImageName string      `json:"imageName" yaml:"imageName"`
	Args      []*Variable `json:"args" yaml:"args"`
	TTY       bool        `json:"tty" yaml:"tty"`
}

// Variable defines the variable and its requirement need
type Variable struct {
	Name     string `yaml:"name"`
	Required bool   `yaml:"required"`
}
