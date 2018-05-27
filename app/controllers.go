// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "DNS Resolver": Application Controllers
//
// Command:
// $ goagen
// --design=github.com/mtenrero/dockerDnsrr-discovery/design
// --out=$(GOPATH)/src/github.com/mtenrero/dockerDnsrr-discovery
// --version=v1.3.1

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// ResolveController is the controller interface for the Resolve actions.
type ResolveController interface {
	goa.Muxer
	Resolve(*ResolveResolveContext) error
}

// MountResolveController "mounts" a Resolve resource controller on the given service.
func MountResolveController(service *goa.Service, ctrl ResolveController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewResolveResolveContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Resolve(rctx)
	}
	service.Mux.Handle("GET", "/api/:hostname", ctrl.MuxHandler("resolve", h, nil))
	service.LogInfo("mount", "ctrl", "Resolve", "action", "Resolve", "route", "GET /api/:hostname")
}
