package controllers

import (
	"net/http"
	"webapp/src/utils"
)

// CarregarTelaLogin vai renderizar a tela de login
func CarregarTelaLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "login.html", nil)
}
