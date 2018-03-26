package types

type ServiceCreateResponse struct {
	// ID is the ID of the created service.
	ID string
	// Warnings is a set of non-fatal warning messages to pass on to the user.
	Warnings []string `json:",omitempty"`
}
