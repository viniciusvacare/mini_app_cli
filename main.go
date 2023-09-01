package main

import (
	"fmt"
	"log"
	"mini-aplicacao/app"
	"os"
)

func main() {
	fmt.Println("inicio")

	aplicacao := app.Gerar()

	// aplicacao.Run(os.Args) - reconhece comandos do sistema no cmd
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro) // erro - para a execucao
	}
}
