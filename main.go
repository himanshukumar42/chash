/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"log"

	"github.com/himanshuk42/chash/cmd"
	"github.com/himanshuk42/chash/data"
)

func main() {
	if err := data.OpenDatabase(); err != nil {
		log.Fatal(err)
	}
	cmd.Execute()
}
