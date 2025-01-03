package repositorios

import (
	"api/src/entities"
	"database/sql"
)

// usuarios representa um repositorio de usuarios
type usuarios struct {
	db *sql.DB
}

// NovoRepositorioDeUsuarios cria um repositorio de usuarios
func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db}
}

// Criar insere um usuario no banco de dados
func (u usuarios) Criar(usuario entities.Usuario) (uint64, error) {
	return 0, nil
}
