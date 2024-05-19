package main

import (
	"find-ip/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("-----------------------")

	appication := app.Gerar()
	if err := appication.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
