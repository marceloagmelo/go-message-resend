package utils

import (
	"fmt"
	"log"
)

// CheckErrFatal checar o erro
func CheckErrFatal(err error, msg string) {
	if err != nil {
		mensagem := fmt.Sprintf("%s: %s", msg, err)
		log.Fatalln(mensagem)
	}
}
