package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/lyodekken/go/api/server"
	"github.com/lyodekken/go/api/models"
	"github.com/lyodekken/go/infra/db"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Extrair informações do corpo da solicitação HTTP e criar um novo usuário
var user models.User
err := json.NewDecoder(r.Body).Decode(&user)
if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
}

// Criar um novo objeto CreateUserParams a partir das informações do usuário
params := db.CreateUserParams{
    Name:     user.Name,
    Email:    user.Email,
    Password: user.Password,
    Role:     db.UsersRole(user.Role),
}

// Obter uma conexão com o banco de dados
dbcon, err := connection.Connect()
if err != nil {
    http.Error(w, "Could not connect to database", http.StatusInternalServerError)
    return
}
defer dbcon.Close()

// Criar um novo objeto Queries usando a conexão do banco de dados
queries := db.New(dbcon)

// Chamar a função CreateUser com os parâmetros apropriados
if err := queries.CreateUser(r.Context(), params); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
}

// Retornar uma resposta HTTP de sucesso
w.WriteHeader(http.StatusOK)
}
