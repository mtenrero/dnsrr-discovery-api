package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("resolve", func() {
	BasePath("/")
	Action("resolve", func() {
		Routing(GET("/:hostname"))
		Description("Resolves IPs from a specified hostname")
		Params(func() {
			Param("hostname", String, "Hostname to resolve")
		})
		Response("OK", ArrayOf(String))
		Response("Not Discoverable", func() {
			Description("The given hostname couldn't be resolved")
			Status(204)
			Media("text/plain")
		})
	})
})

var Resolution = MediaType("application/atq.dnsresolution+json", func() {
	Description("Contents of DNS Resolution")
	Attributes(func() {
		Attribute("ip", String, "IP")
	})
	View("default", func() {
		Attribute("ip")
	})
})
