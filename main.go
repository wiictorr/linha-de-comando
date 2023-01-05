package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

// OBS: go mod "nome do arquivo .mod" faz o arquivo .mod
// Local do pacote que será usado: github.com/urfave/cli

func main() {
	fmt.Println("Iniciando")
	aplicacao := app.Gerar()
	erro := aplicacao.Run(os.Args)
	// tratamento do erro (caso ocorra)
	if erro != nil {
		log.Fatal(erro) // função Fatal para cortar a função
	}
	
}
