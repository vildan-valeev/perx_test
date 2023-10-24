package api

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/vildan-valeev/perx_test/internal/config"
	"github.com/vildan-valeev/perx_test/internal/service"
	"github.com/vildan-valeev/perx_test/internal/transport/http/api/handlers"
)

func NewAPI(ctx context.Context, cfg config.Config, s service.Item) *Router {
	r := NewRouter()
	r.Methods(http.MethodPost).Handler(`/item_to_queue`, handlers.ItemToQueueHandler(ctx, s))
	r.Methods(http.MethodGet).Handler(`/list`, handlers.ListItemHandler(ctx, s))

	return r
}

var (
	HTTP404Response = []byte("page not found")     //nolint:gochecknoglobals
	HTTP405Response = []byte("method not allowed") //nolint:gochecknoglobals
)

func Set404(content string) {
	HTTP404Response = []byte(content)
}

func Set405(content string) {
	HTTP404Response = []byte(content)
}

// Router represents the router which handles routing.
type Router struct {
	tree *tree
}

// route represents the route which has data for a routing.
type route struct {
	methods []string
	path    string
	handler http.Handler
}

var (
	//nolint:gochecknoglobals
	tmpRoute            = &route{}
	ErrNotFound         = errors.New("no matching route was found") // Error for not found.
	ErrMethodNotAllowed = errors.New("methods is not allowed")      // Error for method not allowed.
)

func NewRouter() *Router {
	return &Router{
		tree: NewTree(),
	}
}

func (r *Router) Methods(methods ...string) *Router {
	tmpRoute.methods = append(tmpRoute.methods, methods...)
	return r
}

// Handler sets a handler.
func (r *Router) Handler(path string, handler http.Handler) {
	tmpRoute.handler = handler
	tmpRoute.path = path

	r.Handle()
}

// Handle handles a route.
func (r *Router) Handle() {
	err := r.tree.Insert(tmpRoute.methods, tmpRoute.path, tmpRoute.handler)
	if err != nil {
		log.Fatal(err)
	}

	tmpRoute = &route{}
}

// ServeHTTP dispatches the request to the handler whose
// pattern most closely matches the request URL.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	method := req.Method
	path := req.URL.Path

	result, err := r.tree.Search(method, path)
	if err != nil {
		status, msg := handleErr(err)
		w.WriteHeader(status)

		_, e := w.Write(msg)
		if e != nil {
			return
		}

		return
	}

	h := result.actions.handler
	h.ServeHTTP(w, req)
}

func handleErr(err error) (int, []byte) {
	var status int
	var body []byte

	switch {
	case errors.Is(err, ErrMethodNotAllowed):
		status = http.StatusMethodNotAllowed
		body = HTTP405Response
	case errors.Is(err, ErrNotFound):
		status = http.StatusNotFound
		body = HTTP404Response
	}

	return status, body
}
