package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	context "golang.org/x/net/context"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetKapacitorsHandlerFunc turns a function with the right signature into a get kapacitors handler
type GetKapacitorsHandlerFunc func(context.Context, GetKapacitorsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetKapacitorsHandlerFunc) Handle(ctx context.Context, params GetKapacitorsParams) middleware.Responder {
	return fn(ctx, params)
}

// GetKapacitorsHandler interface for that can handle valid get kapacitors params
type GetKapacitorsHandler interface {
	Handle(context.Context, GetKapacitorsParams) middleware.Responder
}

// NewGetKapacitors creates a new http.Handler for the get kapacitors operation
func NewGetKapacitors(ctx *middleware.Context, handler GetKapacitorsHandler) *GetKapacitors {
	return &GetKapacitors{Context: ctx, Handler: handler}
}

/*GetKapacitors swagger:route GET /kapacitors getKapacitors

Configured kapacitors

*/
type GetKapacitors struct {
	Context *middleware.Context
	Handler GetKapacitorsHandler
}

func (o *GetKapacitors) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetKapacitorsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(context.Background(), Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
