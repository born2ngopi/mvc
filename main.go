package main

import (
	"github.com/joho/godotenv"
	// "os"
	// "log"
)

func main(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error load .env file")
	}
}