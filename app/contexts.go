// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "ATQ - Director": Application Contexts
//
// Command:
// $ goagen
// --design=github.com/mtenrero/ATQ-Director/http/design
// --out=$(GOPATH)\src\github.com\mtenrero\ATQ-Director
// --version=v1.3.1

package app

import (
	"context"
	"github.com/goadesign/goa"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

// ListDatabindContext provides the databind list action context.
type ListDatabindContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewListDatabindContext parses the incoming request URL and body, performs validations and creates the
// context used by the databind controller list action.
func NewListDatabindContext(ctx context.Context, r *http.Request, service *goa.Service) (*ListDatabindContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ListDatabindContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListDatabindContext) OK(r AtqDatabindUploadCollection) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/atq.databind.upload+json; type=collection")
	}
	if r == nil {
		r = AtqDatabindUploadCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKError sends a HTTP response with status code 200.
func (ctx *ListDatabindContext) OKError(r AtqDatabindUploadErrorCollection) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/atq.databind.upload+json; type=collection")
	}
	if r == nil {
		r = AtqDatabindUploadErrorCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NoContent sends a HTTP response with status code 204.
func (ctx *ListDatabindContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// UploadDatabindContext provides the databind upload action context.
type UploadDatabindContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *UploadPayload
}

// NewUploadDatabindContext parses the incoming request URL and body, performs validations and creates the
// context used by the databind controller upload action.
func NewUploadDatabindContext(ctx context.Context, r *http.Request, service *goa.Service) (*UploadDatabindContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := UploadDatabindContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *UploadDatabindContext) OK(r *AtqDatabindUpload) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/atq.databind.upload+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKError sends a HTTP response with status code 200.
func (ctx *UploadDatabindContext) OKError(r *AtqDatabindUploadError) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/atq.databind.upload+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// TheFileDoesnTHaveAnAcceptedCompressionError sends a HTTP response with status code 415.
func (ctx *UploadDatabindContext) TheFileDoesnTHaveAnAcceptedCompressionError(r *AtqDatabindUploadError) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/atq.databind.upload+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 415, r)
}

// UploadErrorError sends a HTTP response with status code 500.
func (ctx *UploadDatabindContext) UploadErrorError(r *AtqDatabindUploadError) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/atq.databind.upload+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// PingMonitoringContext provides the monitoring ping action context.
type PingMonitoringContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewPingMonitoringContext parses the incoming request URL and body, performs validations and creates the
// context used by the monitoring controller ping action.
func NewPingMonitoringContext(ctx context.Context, r *http.Request, service *goa.Service) (*PingMonitoringContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := PingMonitoringContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// Alive sends a HTTP response with status code 200.
func (ctx *PingMonitoringContext) Alive(resp []byte) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	}
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// StatusSwarmContext provides the swarm status action context.
type StatusSwarmContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewStatusSwarmContext parses the incoming request URL and body, performs validations and creates the
// context used by the swarm controller status action.
func NewStatusSwarmContext(ctx context.Context, r *http.Request, service *goa.Service) (*StatusSwarmContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := StatusSwarmContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// Status sends a HTTP response with status code 200.
func (ctx *StatusSwarmContext) Status(resp []byte) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	}
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// CreateTaskContext provides the task create action context.
type CreateTaskContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *TaskPayload
}

// NewCreateTaskContext parses the incoming request URL and body, performs validations and creates the
// context used by the task controller create action.
func NewCreateTaskContext(ctx context.Context, r *http.Request, service *goa.Service) (*CreateTaskContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := CreateTaskContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *CreateTaskContext) OK(r *AtqTask) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/atq.task+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKFull sends a HTTP response with status code 200.
func (ctx *CreateTaskContext) OKFull(r *AtqTaskFull) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/atq.task+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// DefinitionError sends a HTTP response with status code 417.
func (ctx *CreateTaskContext) DefinitionError(r *AtqTask) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/atq.task+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 417, r)
}

// DefinitionErrorFull sends a HTTP response with status code 417.
func (ctx *CreateTaskContext) DefinitionErrorFull(r *AtqTaskFull) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/atq.task+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 417, r)
}

// DeleteTaskContext provides the task delete action context.
type DeleteTaskContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID string
}

// NewDeleteTaskContext parses the incoming request URL and body, performs validations and creates the
// context used by the task controller delete action.
func NewDeleteTaskContext(ctx context.Context, r *http.Request, service *goa.Service) (*DeleteTaskContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := DeleteTaskContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		rctx.ID = rawID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 204.
func (ctx *DeleteTaskContext) OK() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// TaskNotIdentified sends a HTTP response with status code 404.
func (ctx *DeleteTaskContext) TaskNotIdentified() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// TaskCouldNotBeDeleted sends a HTTP response with status code 500.
func (ctx *DeleteTaskContext) TaskCouldNotBeDeleted() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// InspectTaskContext provides the task inspect action context.
type InspectTaskContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID uuid.UUID
}

// NewInspectTaskContext parses the incoming request URL and body, performs validations and creates the
// context used by the task controller inspect action.
func NewInspectTaskContext(ctx context.Context, r *http.Request, service *goa.Service) (*InspectTaskContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := InspectTaskContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		if id, err2 := uuid.FromString(rawID); err2 == nil {
			rctx.ID = id
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("id", rawID, "uuid"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *InspectTaskContext) OK(r *AtqTask) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/atq.task+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKFull sends a HTTP response with status code 200.
func (ctx *InspectTaskContext) OKFull(r *AtqTaskFull) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/atq.task+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// TaskNotIdentified sends a HTTP response with status code 404.
func (ctx *InspectTaskContext) TaskNotIdentified() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ErrorCreatingTheTask sends a HTTP response with status code 500.
func (ctx *InspectTaskContext) ErrorCreatingTheTask(r *AtqTask) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/atq.task+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// ErrorCreatingTheTaskFull sends a HTTP response with status code 500.
func (ctx *InspectTaskContext) ErrorCreatingTheTaskFull(r *AtqTaskFull) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/atq.task+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}