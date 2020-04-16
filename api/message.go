package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/marceloagmelo/go-message-resend/logger"
	"github.com/marceloagmelo/go-message-resend/models"
	"github.com/marceloagmelo/go-message-resend/variaveis"
)

var api = "go-message/api/v1"

//ListaMensagensPendentes listar mensagens
func ListaMensagensPendentes() (mensagens models.Mensagens, erro error) {
	endpoint := variaveis.ApiURL + "/" + api + "/mensagem/status/1"

	resposta, err := GetRequest(endpoint)
	if err != nil {
		return nil, err
	}
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			mensagem := fmt.Sprintf("%s: %s", "Erro ao ler o conteudo recebido", err.Error())
			logger.Erro.Println(mensagem)
			return nil, err
		}
		mensagens = models.Mensagens{}
		err = json.Unmarshal(corpo, &mensagens)
		if err != nil {
			mensagem := fmt.Sprintf("%s: %s", "Erro ao converter o retorno JSON", err.Error())
			logger.Erro.Println(mensagem)
			return nil, err
		}
	}
	return mensagens, nil
}

//ReenviarMensagem reenviar a mensagem
func ReenviarMensagem(novaMensagem models.Mensagem) (mensagemRetorno models.Mensagem, erro error) {
	endpoint := variaveis.ApiURL + "/" + api + "/mensagem/reenviar"

	resposta, err := PutRequest(endpoint, novaMensagem)
	if err != nil {
		return mensagemRetorno, err
	}
	defer resposta.Body.Close()
	if resposta.StatusCode == http.StatusOK {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			mensagem := fmt.Sprintf("%s: %s", "Erro ao ler o conteudo recebido", err.Error())
			logger.Erro.Println(mensagem)
			return mensagemRetorno, err
		}
		mensagemRetorno = models.Mensagem{}
		err = json.Unmarshal(corpo, &mensagemRetorno)
		if err != nil {
			mensagem := fmt.Sprintf("%s: %s", "Erro ao converter o retorno JSON", err.Error())
			logger.Erro.Println(mensagem)
			return mensagemRetorno, err
		}
	}
	return mensagemRetorno, nil
}
