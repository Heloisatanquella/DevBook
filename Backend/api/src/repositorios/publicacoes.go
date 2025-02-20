package repositorios

import (
	"api/src/entities"
	"database/sql"
)

// Publicacoes representa um repositorio de publicacoes
type Publicacoes struct {
	db *sql.DB
}

// NovoRepositorioDePublicacoes cria um repositorio de publicacoes
func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

// Criar insere uma publicacao no banco de dados
func (repositorio Publicacoes) Criar(publicacao entities.Publicacao) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into publicacoes (titulo, conteudo, autor_id) values (?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil

}

// BuscarPorID traz uma única publicacao do benco de dados
func (repositorio Publicacoes) BuscarPorID(publicacaoID uint64) (entities.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
		select p.*, u.nick from
		publicacoes p inner join usuarios u
		on u.id = p.autor_id where p.id = ?`, publicacaoID,
	)
	if erro != nil {
		return entities.Publicacao{}, erro
	}
	defer linhas.Close()

	var publicacao entities.Publicacao

	if linhas.Next() {
		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return entities.Publicacao{}, erro
		}
	}

	return publicacao, nil
}

// Buscar traz as publicações dos usuários seguidos e também do próprio usuário que fez a requisição
func (repositorio Publicacoes) Buscar(usuarioID uint64) ([]entities.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
    SELECT DISTINCT p.*, u.nick 
    FROM publicacoes p 
    INNER JOIN usuarios u ON u.id = p.autor_id 
    LEFT JOIN seguidores s ON p.autor_id = s.usuario_id 
    AND s.seguidor_id = ?
    WHERE u.id = ? OR s.seguidor_id IS NOT NULL -- Adicionando a condição para incluir as próprias publicações
    ORDER BY p.id DESC;
	`, usuarioID, usuarioID) // Agora os dois "?" são usados corretamente

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []entities.Publicacao

	for linhas.Next() {
		var publicacao entities.Publicacao

		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

// Atualizar altera os dados de uma publicacao no banco de dados
func (repositorio Publicacoes) Atualizar(publicacaoID uint64, publicacao entities.Publicacao) error {
	statement, erro := repositorio.db.Prepare(
		"update publicacoes set titulo = ?, conteudo = ? where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoID); erro != nil {
		return erro
	}
	return nil
}

// Deletar exclui uma publicacao do banco de dados
func (repositorio Publicacoes) Deletar(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare(
		"delete from publicacoes where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}
	return nil
}

// BuscarPublicacoesPorUsuario traz todas as publicacoes de um usuario especifico
func (repositorio Publicacoes) BuscarPorUsuario(usuarioID uint64) ([]entities.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
		select p.*, u.nick from publicacoes p
		join usuarios u on u.id = p.autor_id
		where p.autor_id = ?
	`, usuarioID,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []entities.Publicacao

	for linhas.Next() {
		var publicacao entities.Publicacao

		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

// Curtir adiciona uma curtida na publicacao
func (repositorio Publicacoes) Curtir(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare(
		`update publicacoes set curtidas = curtidas + 1 where id = ?`)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}
	return nil
}

// Descurtir remove uma curtida da publicacao
func (repositorio Publicacoes) Descurtir(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare(
		`update publicacoes set curtidas = 
		CASE 
			WHEN curtidas > 0 THEN curtidas - 1
			ELSE 0 
		END
		where id = ?
		`)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}
	return nil
}
