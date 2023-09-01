package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// cli(pacote do github) - app(tipo dentro do pacote)
// Gerar() vai retornar a aplicacao de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "App de linha de comando"
	app.Usage = "Busca Ip's e nomes de servidores"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca ip's de endereco na internet",
			Flags:  flags,
			Action: buscaIps,
		},
		{
			Name:   "servidores",
			Usage:  "busca o nome dos servidores",
			Flags:  flags,
			Action: buscaServer,
		},
	}

	return app
}

func buscaIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscaServer(c *cli.Context) {
	host := c.String("host")

	servers, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
