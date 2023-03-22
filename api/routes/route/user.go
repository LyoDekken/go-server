package route

import (
	"github.com/lyodekken/go/api/controllers"
	"net/http"
)

var userRoute = []Route{
	{
		URI:               "/user",
		Metodo:            http.MethodPost,
		Function:          controllers.CreateUser,
		NeedAuthorization: false,
	},
}
