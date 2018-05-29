// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "ATQ - Director": Application User Types
//
// Command:
// $ goagen
// --design=github.com/mtenrero/ATQ-Director/http/design
// --out=$(GOPATH)/src/github.com/mtenrero/ATQ-Director
// --version=v1.3.1

package client

import (
	"github.com/goadesign/goa"
	"unicode/utf8"
)

// Docker Swarm Join Tokens
type joinTokens struct {
	Manager *string `form:"manager,omitempty" json:"manager,omitempty" xml:"manager,omitempty"`
	Worker  *string `form:"worker,omitempty" json:"worker,omitempty" xml:"worker,omitempty"`
}

// Publicize creates JoinTokens from joinTokens
func (ut *joinTokens) Publicize() *JoinTokens {
	var pub JoinTokens
	if ut.Manager != nil {
		pub.Manager = ut.Manager
	}
	if ut.Worker != nil {
		pub.Worker = ut.Worker
	}
	return &pub
}

// Docker Swarm Join Tokens
type JoinTokens struct {
	Manager *string `form:"manager,omitempty" json:"manager,omitempty" xml:"manager,omitempty"`
	Worker  *string `form:"worker,omitempty" json:"worker,omitempty" xml:"worker,omitempty"`
}

// servicePayload user type.
type servicePayload struct {
	// Service alias, this will identify the Service
	Alias *string `form:"alias,omitempty" json:"alias,omitempty" xml:"alias,omitempty"`
	// Arguments to be passed to the container
	Args []string `form:"args,omitempty" json:"args,omitempty" xml:"args,omitempty"`
	// ID of the Zipped contents that will be mounted and accesible inside the container, PREVIOUSLY UPLOADED
	Fileid *string `form:"fileid,omitempty" json:"fileid,omitempty" xml:"fileid,omitempty"`
	// Docker base image to attach to Service
	Image *string `form:"image,omitempty" json:"image,omitempty" xml:"image,omitempty"`
	// Amount of replicas to be deployed. (1 by default)
	Replicas *int `form:"replicas,omitempty" json:"replicas,omitempty" xml:"replicas,omitempty"`
	// Interactive shell requirement
	Tty *bool `form:"tty,omitempty" json:"tty,omitempty" xml:"tty,omitempty"`
}

// Validate validates the servicePayload type instance.
func (ut *servicePayload) Validate() (err error) {
	if ut.Image == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "image"))
	}
	if ut.Alias == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "alias"))
	}
	return
}

// Publicize creates ServicePayload from servicePayload
func (ut *servicePayload) Publicize() *ServicePayload {
	var pub ServicePayload
	if ut.Alias != nil {
		pub.Alias = *ut.Alias
	}
	if ut.Args != nil {
		pub.Args = ut.Args
	}
	if ut.Fileid != nil {
		pub.Fileid = ut.Fileid
	}
	if ut.Image != nil {
		pub.Image = *ut.Image
	}
	if ut.Replicas != nil {
		pub.Replicas = ut.Replicas
	}
	if ut.Tty != nil {
		pub.Tty = ut.Tty
	}
	return &pub
}

// ServicePayload user type.
type ServicePayload struct {
	// Service alias, this will identify the Service
	Alias string `form:"alias" json:"alias" xml:"alias"`
	// Arguments to be passed to the container
	Args []string `form:"args,omitempty" json:"args,omitempty" xml:"args,omitempty"`
	// ID of the Zipped contents that will be mounted and accesible inside the container, PREVIOUSLY UPLOADED
	Fileid *string `form:"fileid,omitempty" json:"fileid,omitempty" xml:"fileid,omitempty"`
	// Docker base image to attach to Service
	Image string `form:"image" json:"image" xml:"image"`
	// Amount of replicas to be deployed. (1 by default)
	Replicas *int `form:"replicas,omitempty" json:"replicas,omitempty" xml:"replicas,omitempty"`
	// Interactive shell requirement
	Tty *bool `form:"tty,omitempty" json:"tty,omitempty" xml:"tty,omitempty"`
}

// Validate validates the ServicePayload type instance.
func (ut *ServicePayload) Validate() (err error) {
	if ut.Image == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "image"))
	}
	if ut.Alias == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "alias"))
	}
	return
}

// taskPayload user type.
type taskPayload struct {
	Delay  *int            `form:"delay,omitempty" json:"delay,omitempty" xml:"delay,omitempty"`
	Master *servicePayload `form:"master,omitempty" json:"master,omitempty" xml:"master,omitempty"`
	// Task Name Identifier
	Name        *string         `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	WaitCommand *waitCommand    `form:"waitCommand,omitempty" json:"waitCommand,omitempty" xml:"waitCommand,omitempty"`
	Worker      *servicePayload `form:"worker,omitempty" json:"worker,omitempty" xml:"worker,omitempty"`
}

// Validate validates the taskPayload type instance.
func (ut *taskPayload) Validate() (err error) {
	if ut.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "name"))
	}
	if ut.Worker == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "worker"))
	}
	if ut.Master == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "master"))
	}
	if ut.Delay != nil {
		if *ut.Delay < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`request.delay`, *ut.Delay, 0, true))
		}
	}
	if ut.Master != nil {
		if err2 := ut.Master.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if ut.Name != nil {
		if utf8.RuneCountInString(*ut.Name) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.name`, *ut.Name, utf8.RuneCountInString(*ut.Name), 3, true))
		}
	}
	if ut.Name != nil {
		if utf8.RuneCountInString(*ut.Name) > 10 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.name`, *ut.Name, utf8.RuneCountInString(*ut.Name), 10, false))
		}
	}
	if ut.Worker != nil {
		if err2 := ut.Worker.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// Publicize creates TaskPayload from taskPayload
func (ut *taskPayload) Publicize() *TaskPayload {
	var pub TaskPayload
	if ut.Delay != nil {
		pub.Delay = ut.Delay
	}
	if ut.Master != nil {
		pub.Master = ut.Master.Publicize()
	}
	if ut.Name != nil {
		pub.Name = *ut.Name
	}
	if ut.WaitCommand != nil {
		pub.WaitCommand = ut.WaitCommand.Publicize()
	}
	if ut.Worker != nil {
		pub.Worker = ut.Worker.Publicize()
	}
	return &pub
}

// TaskPayload user type.
type TaskPayload struct {
	Delay  *int            `form:"delay,omitempty" json:"delay,omitempty" xml:"delay,omitempty"`
	Master *ServicePayload `form:"master" json:"master" xml:"master"`
	// Task Name Identifier
	Name        string          `form:"name" json:"name" xml:"name"`
	WaitCommand *WaitCommand    `form:"waitCommand,omitempty" json:"waitCommand,omitempty" xml:"waitCommand,omitempty"`
	Worker      *ServicePayload `form:"worker" json:"worker" xml:"worker"`
}

// Validate validates the TaskPayload type instance.
func (ut *TaskPayload) Validate() (err error) {
	if ut.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "name"))
	}
	if ut.Worker == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "worker"))
	}
	if ut.Master == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "master"))
	}
	if ut.Delay != nil {
		if *ut.Delay < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`type.delay`, *ut.Delay, 0, true))
		}
	}
	if ut.Master != nil {
		if err2 := ut.Master.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if utf8.RuneCountInString(ut.Name) < 3 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.name`, ut.Name, utf8.RuneCountInString(ut.Name), 3, true))
	}
	if utf8.RuneCountInString(ut.Name) > 10 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.name`, ut.Name, utf8.RuneCountInString(ut.Name), 10, false))
	}
	if ut.Worker != nil {
		if err2 := ut.Worker.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// Definition of a command to be executed
type waitCommand struct {
	// Command to be executed
	Command *string `form:"command,omitempty" json:"command,omitempty" xml:"command,omitempty"`
	// Result expected
	ExpectedResult *string `form:"expectedResult,omitempty" json:"expectedResult,omitempty" xml:"expectedResult,omitempty"`
	// Maximum seconds to wait until succesfull response
	Timeout *int `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
}

// Publicize creates WaitCommand from waitCommand
func (ut *waitCommand) Publicize() *WaitCommand {
	var pub WaitCommand
	if ut.Command != nil {
		pub.Command = ut.Command
	}
	if ut.ExpectedResult != nil {
		pub.ExpectedResult = ut.ExpectedResult
	}
	if ut.Timeout != nil {
		pub.Timeout = ut.Timeout
	}
	return &pub
}

// Definition of a command to be executed
type WaitCommand struct {
	// Command to be executed
	Command *string `form:"command,omitempty" json:"command,omitempty" xml:"command,omitempty"`
	// Result expected
	ExpectedResult *string `form:"expectedResult,omitempty" json:"expectedResult,omitempty" xml:"expectedResult,omitempty"`
	// Maximum seconds to wait until succesfull response
	Timeout *int `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
}
