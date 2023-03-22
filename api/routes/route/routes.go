package route

import (
	"github.com/gorilla/mux"
	"github.com/lyodekken/go/api/middlewares"
	"net/http"
)

// Rota representa todas as rotas da API
type Route struct {
	URI               string
	Metodo            string
	Function          func(http.ResponseWriter, *http.Request)
	NeedAuthorization bool
}

// SetUp coloca todas as rotas dentro do router
func SetUp(r *mux.Router) *mux.Router {
	routes := userRoute

	for _, route := range routes {

		if route.NeedAuthorization {
			r.HandleFunc(route.URI,
				middlewares.Logger(middlewares.Authenticate(route.Function)),
			).Methods(route.Metodo)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Metodo)
		}

	}

	return r
}
