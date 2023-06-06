package main

import (
	"TopLearn/demo/demowebportal"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type configuration struct {
	ServerAddress string `json:"webserver"`
}

func main() {
	fmt.Println("Openning config File...")
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Decode config File...")
	config := new(configuration)
	err = json.NewDecoder(file).Decode(config)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Starting Server in Address ", config.ServerAddress)
	demowebportal.RunWebPortal(config.ServerAddress)
}
