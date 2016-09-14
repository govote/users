package main

import (
	"fmt"
	"os"

	"github.com/deputadosemfoco/users/routes"
	"github.com/dimiro1/banner"
	_ "github.com/go-sql-driver/mysql"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/engine/standard"
)

func main() {
	godotenv.Load()

	in, _ := os.Open("banner.txt")
	defer in.Close()
	banner.Init(os.Stdout, true, false, in)

	e := routes.SetUp()

	if os.Getenv("GO_ENV") == "development" {
		e.SetDebug(true)
	}

	port := os.Getenv("PORT")

	fmt.Printf("auth service wiil run on port %s", port)

	e.Run(standard.New(":" + port))
}
