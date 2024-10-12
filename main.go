package main

import (
	"github.com/ducconit/gobase/cmd"
	"log"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
