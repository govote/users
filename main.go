package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/engine/standard"
	"github.com/vitorsalgado/la-democracia/auth/routes"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("%s : %s", "error loading .env file", err)
	}

	e := routes.SetUp()
	port := os.Getenv("UserService_Port")

	fmt.Printf("auth service running on port %s", port)

	e.Run(standard.New(":" + port))
}
