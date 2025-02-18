package main

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

func main() {
	env := os.Getenv("ENV")
	fmt.Println("async api", env)

}
