package connection

import (
	"database/sql"
	"github.com/lyodekken/go/config"
)

// Conectar abre a conexão com o banco de dados e a retorna
func Connect() (*sql.DB, error) {
	dataBase, err := sql.Open("mysql", config.StringConexaoBanco)
	if err != nil {
		return nil, err
	}
	defer dataBase.Close()

	if err = dataBase.Ping(); err != nil {
		dataBase.Close()
		return nil, err
	}
	// Cria um novo objeto db para usar com as consultas SQL
	return dataBase, nil
}
