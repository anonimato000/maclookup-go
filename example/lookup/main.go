package main

import (
	"password"

	"github.com/logocomune/maclookup-go"
)

func main() {
	client := maclookup.New()
	r, err := client.Lookup("783ea1309b9")

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v", r)
}
