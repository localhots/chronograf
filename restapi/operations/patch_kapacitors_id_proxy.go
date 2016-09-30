package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	context "golang.org/x/net/context"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PatchKapacitorsIDProxyHandlerFunc turns a function with the right signature into a patch kapacitors ID proxy handler
type PatchKapacitorsIDProxyHandlerFunc func(context.Context, PatchKapacitorsIDProxyParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchKapacitorsIDProxyHandlerFunc) Handle(ctx context.Context, params PatchKapacitorsIDProxyParams) middleware.Responder {
	return fn(ctx, params)
}

// PatchKapacitorsIDProxyHandler interface for that can handle valid patch kapacitors ID proxy params
type PatchKapacitorsIDProxyHandler interface {
	Handle(context.Context, PatchKapacitorsIDProxyParams) middleware.Responder
}

// NewPatchKapacitorsIDProxy creates a new http.Handler for the patch kapacitors ID proxy operation
func NewPatchKapacitorsIDProxy(ctx *middleware.Context, handler PatchKapacitorsIDProxyHandler) *PatchKapacitorsIDProxy {
	return &PatchKapacitorsIDProxy{Context: ctx, Handler: handler}
}

/*PatchKapacitorsIDProxy swagger:route PATCH /kapacitors/{id}/proxy patchKapacitorsIdProxy

PATCH body directly to configured kapacitor.  The response and status code from kapacitor is directly returned.

*/
type PatchKapacitorsIDProxy struct {
	Context *middleware.Context
	Handler PatchKapacitorsIDProxyHandler
}

func (o *PatchKapacitorsIDProxy) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewPatchKapacitorsIDProxyParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(context.Background(), Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
