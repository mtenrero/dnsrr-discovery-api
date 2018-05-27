//go:generate goagen bootstrap -d github.com/mtenrero/dockerDnsrr-discovery/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/mtenrero/dockerDnsrr-discovery/app"
)

func main() {
	// Create service
	service := goa.New("DNS Resolver")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "resolve" controller
	c := NewResolveController(service)
	app.MountResolveController(service, c)

	// Start service
	if err := service.ListenAndServe(":9090"); err != nil {
		service.LogError("startup", "err", err)
	}

}
