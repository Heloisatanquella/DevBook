package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotasUsuarios = []Rota{
	{
		URI:                "/cadastro",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarTelaCadastro,
		RequerAutenticacao: false,
	}, {
		URI:                "/usuarios",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarUsuario,
		RequerAutenticacao: false,
	}, {
		URI:                "/buscar-usuarios",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeUsuarios,
		RequerAutenticacao: true,
	},
}
