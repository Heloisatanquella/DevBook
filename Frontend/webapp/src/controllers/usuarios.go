package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/requisicoes"
	"webapp/src/respostas"

	"github.com/gorilla/mux"
)

// CriarUsuario chama a API para criar um usuario no banco de dados
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

// ParaDeSeguirUsuario chama a API para parar de seguir um usuario
func ParaDeSeguirUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/usuarios/%d/parar-de-seguir", config.APIURL, usuarioID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	if response.StatusCode != http.StatusNoContent {
		defer response.Body.Close()
	}

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	if response.StatusCode == http.StatusNoContent {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)
}

// SeguirUsuario chama a API para seguir um usuario
func SeguirUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/usuarios/%d/seguir", config.APIURL, usuarioID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	if response.StatusCode != http.StatusNoContent {
		defer response.Body.Close()
	}

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	if response.StatusCode == http.StatusNoContent {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)
}

// EditarUsuario chama a API para editar os dadosa de um usuario
func EditarUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario map[string]string

	if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: "Erro ao ler JSON da requisição"})
		return
	}

	if usuario["nome"] == "" || usuario["email"] == "" || usuario["nick"] == "" {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: "Todos os campos são obrigatórios"})
		return
	}

	usuarioJSON, erro := json.Marshal(usuario)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: "Erro ao converter JSON"})
		return
	}
	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/usuarios/%d", config.APIURL, usuarioID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPut, url, bytes.NewBuffer(usuarioJSON))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: "Erro ao enviar requisição"})
		return
	}

	if response.StatusCode != http.StatusNoContent {
		defer response.Body.Close()
	}

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	if response.StatusCode == http.StatusNoContent {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)

}

// AtualizarSenha chama a API para atualizar a senha do usuario
func AtualizarSenha(w http.ResponseWriter, r *http.Request) {
	var senha map[string]string

	if err := json.NewDecoder(r.Body).Decode(&senha); err != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: "Erro ao ler JSON da requisição"})
		return
	}

	if senha["atual"] == "" || senha["nova"] == "" {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: "Todos os campos são obrigatórios"})
		return
	}

	senhaJSON, erro := json.Marshal(senha)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: "Erro ao converter JSON"})
		return
	}
	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/usuarios/%d/atualizar-senha", config.APIURL, usuarioID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, bytes.NewBuffer(senhaJSON))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: "Erro ao enviar requisição"})
		return
	}

	if response.StatusCode != http.StatusNoContent {
		defer response.Body.Close()
	}

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	if response.StatusCode == http.StatusNoContent {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)

}
