// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "ATQ - Director": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/mtenrero/ATQ-Director/http/design
// --out=$(GOPATH)/src/github.com/mtenrero/ATQ-Director
// --version=v1.3.1

package client

import (
	"github.com/goadesign/goa"
	uuid "github.com/goadesign/goa/uuid"
	"net/http"
)

// User upload files response (default view)
//
// Identifier: application/atq.databind.upload+json; view=default
type AtqDatabindUpload struct {
	// Upload ID
	ID *uuid.UUID `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
}

// User upload files response (error view)
//
// Identifier: application/atq.databind.upload+json; view=error
type AtqDatabindUploadError struct {
	// Error message if errored
	Error *string `form:"error,omitempty" json:"error,omitempty" xml:"error,omitempty"`
}

// DecodeAtqDatabindUpload decodes the AtqDatabindUpload instance encoded in resp body.
func (c *Client) DecodeAtqDatabindUpload(resp *http.Response) (*AtqDatabindUpload, error) {
	var decoded AtqDatabindUpload
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeAtqDatabindUploadError decodes the AtqDatabindUploadError instance encoded in resp body.
func (c *Client) DecodeAtqDatabindUploadError(resp *http.Response) (*AtqDatabindUploadError, error) {
	var decoded AtqDatabindUploadError
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// AtqDatabindUploadCollection is the media type for an array of AtqDatabindUpload (default view)
//
// Identifier: application/atq.databind.upload+json; type=collection; view=default
type AtqDatabindUploadCollection []*AtqDatabindUpload

// AtqDatabindUploadCollection is the media type for an array of AtqDatabindUpload (error view)
//
// Identifier: application/atq.databind.upload+json; type=collection; view=error
type AtqDatabindUploadErrorCollection []*AtqDatabindUploadError

// DecodeAtqDatabindUploadCollection decodes the AtqDatabindUploadCollection instance encoded in resp body.
func (c *Client) DecodeAtqDatabindUploadCollection(resp *http.Response) (AtqDatabindUploadCollection, error) {
	var decoded AtqDatabindUploadCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeAtqDatabindUploadErrorCollection decodes the AtqDatabindUploadErrorCollection instance encoded in resp body.
func (c *Client) DecodeAtqDatabindUploadErrorCollection(resp *http.Response) (AtqDatabindUploadErrorCollection, error) {
	var decoded AtqDatabindUploadErrorCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// Swarm Details (default view)
//
// Identifier: application/atq.swarm+json; view=default
type AtqSwarm struct {
	JoinTokens *JoinTokens `form:"joinTokens,omitempty" json:"joinTokens,omitempty" xml:"joinTokens,omitempty"`
}

// Swarm Details (error view)
//
// Identifier: application/atq.swarm+json; view=error
type AtqSwarmError struct {
	// Swarm Error Message
	Error *string `form:"error,omitempty" json:"error,omitempty" xml:"error,omitempty"`
}

// DecodeAtqSwarm decodes the AtqSwarm instance encoded in resp body.
func (c *Client) DecodeAtqSwarm(resp *http.Response) (*AtqSwarm, error) {
	var decoded AtqSwarm
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeAtqSwarmError decodes the AtqSwarmError instance encoded in resp body.
func (c *Client) DecodeAtqSwarmError(resp *http.Response) (*AtqSwarmError, error) {
	var decoded AtqSwarmError
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Task description (default view)
//
// Identifier: application/atq.task+json; view=default
type AtqTask struct {
	// Task ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Status of the Task
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
}

// Validate validates the AtqTask media type instance.
func (mt *AtqTask) Validate() (err error) {
	if mt.Status != nil {
		if !(*mt.Status == "initializing" || *mt.Status == "started" || *mt.Status == "stopped" || *mt.Status == "finished" || *mt.Status == "errored") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.status`, *mt.Status, []interface{}{"initializing", "started", "stopped", "finished", "errored"}))
		}
	}
	return
}

// Task description (full view)
//
// Identifier: application/atq.task+json; view=full
type AtqTaskFull struct {
	Delay *int `form:"delay,omitempty" json:"delay,omitempty" xml:"delay,omitempty"`
	// Task ID
	ID     *string         `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Master *ServicePayload `form:"master,omitempty" json:"master,omitempty" xml:"master,omitempty"`
	// Status of the Task
	Status      *string         `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	WaitCommand *WaitCommand    `form:"waitCommand,omitempty" json:"waitCommand,omitempty" xml:"waitCommand,omitempty"`
	Worker      *ServicePayload `form:"worker,omitempty" json:"worker,omitempty" xml:"worker,omitempty"`
}

// Validate validates the AtqTaskFull media type instance.
func (mt *AtqTaskFull) Validate() (err error) {
	if mt.Delay != nil {
		if *mt.Delay < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.delay`, *mt.Delay, 0, true))
		}
	}
	if mt.Master != nil {
		if err2 := mt.Master.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if mt.Status != nil {
		if !(*mt.Status == "initializing" || *mt.Status == "started" || *mt.Status == "stopped" || *mt.Status == "finished" || *mt.Status == "errored") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.status`, *mt.Status, []interface{}{"initializing", "started", "stopped", "finished", "errored"}))
		}
	}
	if mt.Worker != nil {
		if err2 := mt.Worker.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// DecodeAtqTask decodes the AtqTask instance encoded in resp body.
func (c *Client) DecodeAtqTask(resp *http.Response) (*AtqTask, error) {
	var decoded AtqTask
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeAtqTaskFull decodes the AtqTaskFull instance encoded in resp body.
func (c *Client) DecodeAtqTaskFull(resp *http.Response) (*AtqTaskFull, error) {
	var decoded AtqTaskFull
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
