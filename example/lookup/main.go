package main

import (
	"log"

	"github.com/logocomune/maclookup-go"
)

func main() {
	client := maclookup.New()
	r, err := client.Lookup("bcf88bd84560")

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v", r)
}
