package main

import (
	"fmt"
	"waysbeans/database"
	"waysbeans/pkg/mysql"
	"waysbeans/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	mysql.DatabaseInit()
	database.RunMigration()

	routes.RouteInit(e.Group("/api/v1"))
	e.Static("/uploads", "./uploads")

	fmt.Println("Succeeded Run On Server")
	e.Logger.Fatal(e.Start("localhost:5000"))
}
