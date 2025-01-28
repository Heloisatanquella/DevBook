package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotasUsuarios = []Rota{
	{
		URI:                "/register",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarTelaCadastro,
		RequerAutenticacao: false,
	},
}
