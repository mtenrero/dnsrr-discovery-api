// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "DNS Resolver": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/mtenrero/dockerDnsrr-discovery/design
// --out=$(GOPATH)/src/github.com/mtenrero/dockerDnsrr-discovery
// --version=v1.3.1

package client

import (
	"net/http"
)

// Contents of DNS Resolution (default view)
//
// Identifier: application/atq.dnsresolution+json; view=default
type AtqDnsresolution struct {
	// IP
	IP *string `form:"ip,omitempty" json:"ip,omitempty" xml:"ip,omitempty"`
}

// DecodeAtqDnsresolution decodes the AtqDnsresolution instance encoded in resp body.
func (c *Client) DecodeAtqDnsresolution(resp *http.Response) (*AtqDnsresolution, error) {
	var decoded AtqDnsresolution
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}