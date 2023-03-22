package router

import (
	"github.com/gorilla/mux"
	"github.com/lyodekken/go/api/routes/route"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	return route.SetUp(r)
}
