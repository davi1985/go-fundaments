package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

const IP string = "ip"
const SERVER string = "server"

func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "CLI App"
	app.Usage = "Busca IPs e nomes de Servidor na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPs na internet",
			Flags:  flags,
			Action: getIPs,
		},
		{
			Name:   "servers",
			Usage:  "Busca o nome dos servidores na internet",
			Flags:  flags,
			Action: getServerName,
		},
	}

	return app
}

func getIPs(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func getServerName(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
