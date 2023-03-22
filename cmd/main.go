package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/lyodekken/go/config"
	"github.com/lyodekken/go/api/routes"
)

func main() {
	config.Loading()

	r := router.Router()

	fmt.Printf("Escutando na porta %d\n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
