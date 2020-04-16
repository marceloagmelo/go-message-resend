package main

import (
	"fmt"

	"github.com/marceloagmelo/go-message-resend/api"
	"github.com/marceloagmelo/go-message-resend/logger"
	"github.com/marceloagmelo/go-message-resend/utils"
)

func processaReenvio() {
	mensagens, err := api.ListaMensagensPendentes()
	utils.CheckErrFatal(err, "Falha ao ler as mensagens pendentes")

	// Ler todas as mensagens
	for _, mensagem := range mensagens {
		novaMensagem, err := api.ReenviarMensagem(mensagem)
		if err != nil {
			return
		}

		msg := fmt.Sprintf("Mensagem %v enviada com sucesso!", novaMensagem.ID)
		logger.Info.Println(msg)
	}

}

func main() {
	logger.Info.Println("Inicio...")

	processaReenvio()

	logger.Info.Println("Fim...")
}
