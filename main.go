package main

import (
	"log"
	"os"

	yaml "github.com/goccy/go-yaml"
)

// Document wraps whole YAML
type Document struct {
	Version  string             `yaml:"version"`
	Services map[string]Service `yaml:"services"`
	Volumes  map[string]Volume  `yaml:"volumes"`
}

// Service wraps each service
type Service struct {
	Image       string            `yaml:"image"`
	Environment map[string]string `yaml:"environment"`
	WorkingDir  string            `yaml:"working_dir"`
	Command     []string          `yaml:"command"`
	Volumes     []string          `yaml:"volumes"`
	Ports       []string          `yaml:"ports"`
	DependsOn   []string          `yaml:"depends_on"`
	EnvFile     []string          `yaml:"env_file"`
}

// Volume wraps each volume
type Volume struct {
	Driver string `yaml:"driver"`
}

func main() {
	doc := Document{
		Version: "3.8",
		Services: map[string]Service{
			"app": {
				Image:      "ruby:3.0.0",
				WorkingDir: "/app",
				Command: []string{
					"bundle",
					"exec",
					"rails",
					"server",
					"--binding",
					"0.0.0.0",
				},
				Volumes: []string{
					".:/app:cached",
					"bundle:/usr/local/bundle:delegated",
				},
				Ports: []string{
					"3000:3000",
				},
				DependsOn: []string{
					"postgres",
				},
				EnvFile: []string{
					"app.env",
				},
			},
			"mysql": {
				Image: "mysql:8.0",
				Environment: map[string]string{
					"MYSQL_DATABASE": "app_development",
					"MYSQL_USER":     "app",
					"MYSQL_PASSWORD": "password",
				},
			},
			"postgres": {
				Image: "postgres:13.2-alpine",
				Environment: map[string]string{
					"POSTGRES_USER":     "app",
					"POSTGRES_PASSWORD": "password",
				},
			},
		},
		Volumes: map[string]Volume{
			"bundle": {
				Driver: "local",
			},
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
