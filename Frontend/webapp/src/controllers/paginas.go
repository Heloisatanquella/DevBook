package controllers

import (
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/requisicoes"
	"webapp/src/utils"
)

// CarregarTelaLogin vai renderizar a tela de login
func CarregarTelaLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "login.html", nil)
}

// CarregarTelaLogin vai renderizar a tela de cadastro
func CarregarTelaCadastro(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "register.html", nil)
}

// CarregarPaginaPrincipal vai renderizar a tela home com as publicacoes
func CarregarPaginaPrincipal(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/publicacoes", config.APIURL)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	fmt.Println(response.StatusCode, erro)
	fmt.Println(url)

	utils.ExecutarTemplate(w, "home.html", nil)
}
