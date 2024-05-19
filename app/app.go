package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func main() {

}

// Gerar ira retornar uma aplicação cli
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Find IP"
	app.Usage = "Search for IPs and hostnames on the internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "www.mypomolist.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "ip",
			Usage:   "Search for IPs on the internet",
			Flags:   flags,
			Aliases: []string{"i"},
			Action:  searchForIP,
		},
		{
			Name:    "server",
			Usage:   "Search for servers on the internet",
			Flags:   flags,
			Aliases: []string{"s"},
			Action:  buscarServidor,
		},
	}

	return app
}

func searchForIP(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("IP address found for the host: %s \n", host)
	fmt.Println("-----------------------")
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidor(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Servers found for the host: %s \n", host)
	fmt.Println("-----------------------")
	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
