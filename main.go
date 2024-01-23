package main

import (
	"account/config"
	ph "account/features/permissions/handler"
	pr "account/features/permissions/repository"
	ps "account/features/permissions/services"
	"account/routes"
	"account/utils/database"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	cfg := config.InitConfig()

	mongoClient, err := database.InitMongoDB(*cfg)
	if err != nil {
		e.Logger.Fatal("tidak bisa start bro", err.Error())
	}

	permissionsRepo := pr.New(mongoClient.Client().Database(cfg.DBNAME), "permissions")
	permissionsServices := ps.New(permissionsRepo)
	permissionsHandler := ph.New(permissionsServices)

	routes.InitRoute(e, permissionsHandler)

	e.Logger.Fatal(e.Start(":8000"))
}
