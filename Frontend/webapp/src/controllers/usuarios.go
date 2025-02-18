package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/respostas"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario map[string]string

	if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: "Erro ao ler JSON da requisição"})
		return
	}

	if usuario["nome"] == "" || usuario["email"] == "" || usuario["nick"] == "" || usuario["senha"] == "" {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: "Todos os campos são obrigatórios"})
		return
	}

	usuarioJSON, erro := json.Marshal(usuario)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: "Erro ao converter JSON"})
		return
	}

	url := fmt.Sprintf("%s/usuarios", config.APIURL)
	response, erro := http.Post(url, "application/json", bytes.NewBuffer(usuarioJSON))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: "Erro ao enviar requisição"})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)
}
