package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/simple-jwt-auth/servers"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	servers.Run()
	log.Println("Server exiting")
}
