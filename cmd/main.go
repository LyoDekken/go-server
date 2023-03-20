package main

import (
    "net/http"
    "database/sql"
    "go-server/internal/db"
    "go-server/api/routes"
    "go-server/api/handlers"
)

func main() {
    // Cria uma conexão com o banco de dados
	dbcon, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/dbname")
	if err != nil {
		panic(err)
	}
	defer dbcon.Close()

	// Cria um novo objeto db para usar com as consultas SQL
	dt := db.New(dbcon)

    // Configura as rotas da API
	http.HandleFunc("/hello", routes.HelloHandler)
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateUserHandler(w, r, dt)
	})

    // Inicia o servidor HTTP
    http.ListenAndServe(":8080", nil)
}
