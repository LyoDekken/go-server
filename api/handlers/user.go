package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/lyodekken/go/infra/db"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request, dt *db.Queries) {
	// Extrai as informações do novo usuário do corpo da solicitação
	var newUser db.CreateUserParams
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Erro ao decodificar o corpo da solicitação", http.StatusBadRequest)
		return
	}

	// Executa a consulta para criar o novo usuário
	err = dt.CreateUser(r.Context(), newUser)
	if err != nil {
		http.Error(w, "Erro ao criar o novo usuário", http.StatusInternalServerError)
		return
	}

	// Retorna uma resposta HTTP indicando que o usuário foi criado com sucesso
	w.WriteHeader(http.StatusCreated)
}
