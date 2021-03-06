// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "ATQ - Director": swarm Resource Client
//
// Command:
// $ goagen
// --design=github.com/mtenrero/ATQ-Director/http/design
// --out=$(GOPATH)/src/github.com/mtenrero/ATQ-Director
// --version=v1.3.1

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// StatusSwarmPath computes a request path to the status action of swarm.
func StatusSwarmPath() string {

	return fmt.Sprintf("/api/swarm/")
}

// Response with the details of the swarm
func (c *Client) StatusSwarm(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewStatusSwarmRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewStatusSwarmRequest create the request corresponding to the status action endpoint of the swarm resource.
func (c *Client) NewStatusSwarmRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
