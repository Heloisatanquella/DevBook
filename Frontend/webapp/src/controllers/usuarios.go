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
		respostas.JSON(w, http.StatusInternalServerError, map[string]string{"erro": "Erro ao converter JSON"})
		return
	}

	req, err := http.NewRequest("POST", "http://localhost:5000/usuarios", bytes.NewBuffer(usuario))
	if err != nil {
		respostas.JSON(w, http.StatusInternalServerError, map[string]string{"erro": "Erro ao criar requisição"})
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		respostas.JSON(w, http.StatusServiceUnavailable, map[string]string{"erro": "Erro ao conectar com o servidor"})
		return
	}
	defer resp.Body.Close()

	respostas.JSON(w, resp.StatusCode, nil)
}
