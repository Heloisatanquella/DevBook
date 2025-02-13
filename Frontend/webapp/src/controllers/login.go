package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/entities"
	"webapp/src/respostas"
)

// FazerLogin utiliza o e-mail e senha do usuário para autenticar na aplicação
func FazerLogin(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		respostas.JSON(w, http.StatusBadRequest, map[string]string{"erro": "Erro ao processar formulário"})
		return
	}

	usuario, erro := json.Marshal(map[string]string{
		"email": r.FormValue("email"),
		"senha": r.FormValue("senha"),
	})

	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/login", config.APIURL)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(usuario))
	if err != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: err.Error()})
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: err.Error()})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, resp)
		return
	}

	var dadosAutenticacao entities.DadosAutenticacao
	if erro := json.NewDecoder(resp.Body).Decode(&dadosAutenticacao); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

}
