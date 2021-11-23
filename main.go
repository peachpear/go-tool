package main

import (
	"github.com/peachpear/go-tool/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal("cmd.Execute err: %v", err)
	}
}
