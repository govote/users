package main

import (
	"fmt"
	"os"

	"github.com/deputadosemfoco/users/network"
	"github.com/deputadosemfoco/users/repositories"
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

	network.FacebookAPP()
	repositories.Bootstrap()

	port := os.Getenv("PORT")
	e := routes.SetUp()

	fmt.Printf("auth service wiil run on port %s", port)
	fmt.Println("")

	e.Run(standard.New(":" + port))
}
