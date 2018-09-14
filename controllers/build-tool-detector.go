/*

Package controllers is autogenerated
and containing scaffold outputs
as well as manually created sub-packages
and files.

*/
package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/goadesign/goa"
	"github.com/tinakurian/build-tool-detector/app"
	errs "github.com/tinakurian/build-tool-detector/controllers/error"
	"github.com/tinakurian/build-tool-detector/controllers/git"
	"github.com/tinakurian/build-tool-detector/controllers/git/buildtype"
	"github.com/tinakurian/build-tool-detector/controllers/git/github"
	"net/http"
)

// BuildToolDetectorController implements the build-tool-detector resource.
type BuildToolDetectorController struct {
	*goa.Controller
}

// NewBuildToolDetectorController creates a build-tool-detector controller.
func NewBuildToolDetectorController(service *goa.Service) *BuildToolDetectorController {
	return &BuildToolDetectorController{Controller: service.NewController("BuildToolDetectorController")}
}

// Show runs the show action.
func (c *BuildToolDetectorController) Show(ctx *app.ShowBuildToolDetectorContext) error {
	err, gitService := git.GetServiceType(ctx.URL)
	if err != nil {
		return handleRequest(ctx, err, nil)
	}

	if gitService.Service != git.GITHUB {
		return handleRequest(ctx, nil, buildtype.Unknown())
	}

	err, buildTool := github.DetectBuildTool(ctx, gitService.Segments)
	if err != nil {
		return handleRequest(ctx, err, nil)
	}

	return handleRequest(ctx, nil, buildTool)
}

func handleRequest(ctx *app.ShowBuildToolDetectorContext, httpTypeError *errs.HTTPTypeError, buildTool *app.GoaBuildToolDetector) error {
	ctx.ResponseWriter.Header().Set("Content-Type", "application/json")

	if httpTypeError != nil {
		ctx.WriteHeader(httpTypeError.StatusCode)
		fmt.Fprint(ctx.ResponseWriter, string(marshalJSON(httpTypeError)))
	}

	if httpTypeError == nil {
		return ctx.OK(buildTool)
	}

	return getErrResponse(ctx, httpTypeError)
}

func marshalJSON(httpTypeError *errs.HTTPTypeError) []byte {

	jsonHTTPTypeError, err := json.Marshal(httpTypeError)
	if err != nil {
		panic(err)
	}

	return jsonHTTPTypeError
}

func getErrResponse(ctx *app.ShowBuildToolDetectorContext, httpTypeError *errs.HTTPTypeError) error {
	var response error
	switch httpTypeError.StatusCode {
	case http.StatusBadRequest:
		response = ctx.BadRequest()
	case http.StatusInternalServerError:
		response = ctx.InternalServerError()
	}

	return response
}
