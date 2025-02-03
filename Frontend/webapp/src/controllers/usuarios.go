package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
	"webapp/src/respostas"
)

// Cliente HTTP reutilizável
var httpClient = &http.Client{
	Timeout: time.Second * 10,
	Transport: &http.Transport{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 10,
	},
}

// CriarUsuario chama a API para cadastrar um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		respostas.JSON(w, http.StatusBadRequest, map[string]string{"erro": "Erro ao processar formulário"})
		return
	}

	usuario, err := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"email": r.FormValue("email"),
		"nick":  r.FormValue("nick"),
		"senha": r.FormValue("senha"),
	})
	if err != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: err.Error()})
		return
	}

	req, err := http.NewRequest("POST", "http://localhost:5000/usuarios", bytes.NewBuffer(usuario))
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

	if resp.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, resp)
		return
	}

	defer resp.Body.Close()

	respostas.JSON(w, resp.StatusCode, nil)
}
