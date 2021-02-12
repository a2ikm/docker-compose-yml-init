package main

import (
	"log"
	"os"
)

func main() {
	yml := `version: 3.7
services:
	mysql:
		image: mysql:8:0
		environment:
			MYSQL_DATABASE: app_development
			MYSQL_USER: app
			MYSQL_PASSWORD: password
	`

	f, err := os.OpenFile("docker-compose.yml", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	f.WriteString(yml)
}
