package main

import (
	"log"
)

func main() {

	store, err := NewPostgressStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.init(); err != nil {
		log.Fatal(err)
	}

	server := newAPIserver(":5000", store)

	server.Run()
}
