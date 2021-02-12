package main

import (
	"log"
	"os"

	yaml "github.com/goccy/go-yaml"
)

type Service struct {
	Image       string
	Environment map[string]string
}

type Document struct {
	Version  string
	Services map[string]Service
}

func main() {
	doc := Document{
		Version:  "3.8",
		Services: map[string]Service{},
	}
	doc.Services["mysql"] = Service{
		Image: "mysql:8.0",
		Environment: map[string]string{
			"MYSQL_DATABASE": "app_development",
			"MYSQL_USER":     "app",
			"MYSQL_PASSWORD": "password",
		},
	}

	yml, err := yaml.Marshal(doc)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile("docker-compose.yml", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	f.Write(yml)
}
