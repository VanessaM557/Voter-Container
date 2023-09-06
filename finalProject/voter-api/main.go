package main

import (
	"log"
	"voter-api/api"
)

func main() {
	r := api.SetupRouter().
	log.Fatal(r.Run(":8080")) 
}

