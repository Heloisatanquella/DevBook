package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/entities"
	"webapp/src/respostas"
)

// FazerLogin utiliza o e-mail e senha do usuário para autenticar na aplicação
func FazerLogin(w http.ResponseWriter, r *http.Request) {
	var usuario map[string]string

	if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: "Erro ao ler JSON da requisição"})
		return
	}

	if usuario["email"] == "" || usuario["senha"] == "" {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: "Todos os campos são obrigatórios"})
		return
	}

	usuarioJSON, erro := json.Marshal(usuario)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: "Erro ao converter JSON"})
		return
	}

	url := fmt.Sprintf("%s/login", config.APIURL)
	response, erro := http.Post(url, "application/json", bytes.NewBuffer(usuarioJSON))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	var dadosAutenticacao entities.DadosAutenticacao
	if erro = json.NewDecoder(response.Body).Decode(&dadosAutenticacao); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	if erro = cookies.Salvar(w, dadosAutenticacao.ID, dadosAutenticacao.Token); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	respostas.JSON(w, http.StatusOK, nil)

}
