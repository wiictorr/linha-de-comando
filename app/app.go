package app

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net"
)

// Função já em formato .exe em que é possível utilizá-la da seguinte forma:

// Para acessar o IP de um site, no terminal, escreve-se:
// ./linda-de-comando ip --host endereco_do_site
// Para acessar o servidor onde o IP está alocado, escreve-se:
// ./linda-de-comando servidores --host endereco_do_site

// -------------------------- função de procura de IP e alocação do site ---------------------------------------------

// função gerar vai retornar a linha de aplicação a ser executada

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Linha de comando"
	app.Usage = "Buscar IP's e nomes de servidores"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br", // Endereço de um site qualquer
		},
	}
	app.Commands = []cli.Command{
		{Name: "ip",
			Usage:  "Buscar IP's e nomes de servidores",
			Flags:  flags,
			Action: AcharIPs,
		},
		{Name: "servidores",
			Usage:  "Busca o nome de servidores na internet",
			Flags:  flags,
			Action: encontrarServidores},
	}
	// ---------------------------------
	return app

}

// ------------------- função que procura os IP's e cancela a procura caso ocorra falha ---------------------

func AcharIPs(c *cli.Context) {
	host := c.String("host")
	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

// ------------------ função que procura os servidores na, qual, os IP's estão alocados ------------------------

func encontrarServidores(c *cli.Context) {
	host := c.String("host")
	// tratamento de erro, caso ocorra
	servidores, erro := net.LookupNS(host) //name server
	if erro != nil {
		log.Fatal(erro)

	}
	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}
