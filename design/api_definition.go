package design

import (
	a "github.com/goadesign/goa/design/apidsl"
)

var _ = a.API("build-tool-detector", func() {
	a.Origin("*", func() {
		a.Methods("GET", "POST", "PUT", "PATCH", "DELETE")
		a.Headers("Accept", "Content-Type")
		a.Expose("Content-Type", "Origin")
		a.Credentials()
	})
	a.Title("Build Tool Detector")
	a.Description("Detects the build tool for a specific repository and branch")
	a.Scheme("http")
	a.Host("localhost:8080")
	a.BasePath("/build-tool-detector")
})