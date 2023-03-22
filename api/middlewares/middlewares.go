package middlewares

import (
	"github.com/lyodekken/go/api/auth"
	"github.com/lyodekken/go/api/json"
	"log"
	"net/http"
)

// Logger escreve informações da requisição no terminal
func Logger(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n%s %s %s", r.Method, r.RequestURI, r.Host)
		nextFunc(w, r)
	}
}

// Autenticar verifica se o usuário fazendo a requisição está autenticado
func Authenticate(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if erro := auth.TokenValidation(r); erro != nil {
			json.Erro(w, http.StatusUnauthorized, erro)
			return
		}
		nextFunc(w, r)
	}
}
