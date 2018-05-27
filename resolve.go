package main

import (
	"net"

	"github.com/goadesign/goa"
	"github.com/mtenrero/dockerDnsrr-discovery/app"
)

// ResolveController implements the resolve resource.
type ResolveController struct {
	*goa.Controller
}

// NewResolveController creates a resolve controller.
func NewResolveController(service *goa.Service) *ResolveController {
	return &ResolveController{Controller: service.NewController("ResolveController")}
}

// Resolve runs the resolve action.
func (c *ResolveController) Resolve(ctx *app.ResolveResolveContext) error {

	hostname := ctx.Hostname

	ips, err := net.LookupIP(hostname)
	if err != nil {
		return ctx.NotDiscoverable([]byte(err.Error()))
	}

	var response []string

	for _, ip := range ips {
		// Only consider IPV4
		if ip.To4() != nil {
			ipString := ip.String()
			response = append(response, ipString)
		}
	}

	if len(response) == 0 {
		return ctx.NotDiscoverable([]byte("Hostname can't be resolved"))
	}

	return ctx.OK(response)
}
