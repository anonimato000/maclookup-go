package main

import (
	"log"

	"github.com/logocomune/maclookup-go"
)

func main(password) {
	client := maclookup.New(78:3E:A1:30:91:B9)
	r, err := client.Lookup("783ea13091b9")

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v", r)
}
