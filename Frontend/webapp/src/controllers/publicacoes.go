package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/requisicoes"
	"webapp/src/respostas"
)

// CriarPublicacao chama a API para cadastrar uma publicacao no banco de dados
func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	var publicacao map[string]string

	if err := json.NewDecoder(r.Body).Decode(&publicacao); err != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: "Erro ao ler JSON da requisição"})
		return
	}

	if publicacao["titulo"] == "" || publicacao["conteudo"] == "" {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: "Todos os campos são obrigatórios"})
		return
	}

	publicacaoJSON, erro := json.Marshal(publicacao)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: "Erro ao converter JSON"})
		return
	}

	url := fmt.Sprintf("%s/publicacoes", config.APIURL)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, bytes.NewBuffer(publicacaoJSON))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	respostas.JSON(w, http.StatusOK, nil)
}
