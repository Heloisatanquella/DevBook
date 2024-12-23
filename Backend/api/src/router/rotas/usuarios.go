package rotas

import "net/http"

var rotasUsuarios = []Rota{
	{
		URI:    "/usuarios",
		Metodo: http.MethodPost,
		Funcao: func(w http.ResponseWriter, r *http.Request) {

		},
		RequerAutenticacao: false,
	}, {
		URI:    "/usuarios",
		Metodo: http.MethodGet,
		Funcao: func(w http.ResponseWriter, r *http.Request) {

		},
		RequerAutenticacao: false,
	}, {
		URI:    "/usuarios/{id}",
		Metodo: http.MethodGet,
		Funcao: func(w http.ResponseWriter, r *http.Request) {

		},
		RequerAutenticacao: false,
	}, {
		URI:    "/usuarios/{id}",
		Metodo: http.MethodPut,
		Funcao: func(w http.ResponseWriter, r *http.Request) {

		},
		RequerAutenticacao: false,
	}, {
		URI:    "/usuarios/{id}",
		Metodo: http.MethodDelete,
		Funcao: func(w http.ResponseWriter, r *http.Request) {

		},
		RequerAutenticacao: false,
	},
}
