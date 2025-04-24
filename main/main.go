package main

import (
	"fmt"
	"log"

	"a3aan.cat/api"
)

func main() {
	log.SetPrefix("Api response: ")
	log.SetFlags(0)

	messages, error := api.Hellos([]string{"ImRock", "Michaili (a broken train)", "leahhh!"})

	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(messages)
}
