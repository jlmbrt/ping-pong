package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Address string
	Port    string
	Name    string
	Message string
}

func (c *Config) Bind() string {
	return fmt.Sprintf("%s:%s", c.Address, c.Port)
}

var config Config

func main() {
	config = loadConfig()

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/ready", readyHandler)
	http.HandleFunc("/sleep/", sleepHandler)

	fmt.Printf("%s listening on %s\n", config.Name, config.Bind())
	err := http.ListenAndServe(config.Bind(), nil)

	if err != nil {
		log.Fatalln("There's an error with the server", err)
	}
}

func loadConfig() Config {
	return Config{
		Address: getEnv("ADDR", "0.0.0.0"),
		Port:    getEnv("PORT", "8080"),
		Name:    getEnv("NAME", generateName()),
		Message: getEnv("MESSAGE", "Happy to meet you 💻"),
	}
}

func getEnv(key string, defaultValue string) string {
	if value := os.Getenv(key); len(value) > 0 {
		return value
	} else {
		return defaultValue
	}
}
