package main

import (
	"log"

	"github.com/turtacn/ngcloud-installer/pkg/console"
)

func main() {
	if err := console.RunConsole(); err != nil {
		log.Panicln(err)
	}
}
