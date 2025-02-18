package controllers

import (
	"net/http"
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
	utils.ExecutarTemplate(w, "home.html", nil)
}
