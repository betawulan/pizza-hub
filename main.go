package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"

	"github.com/betawulan/pizza-hub/delivery"
	"github.com/betawulan/pizza-hub/repository"
	"github.com/betawulan/pizza-hub/service"
)

func main() {
	viper.AutomaticEnv()
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("failed running because file .env")
	}

	port := viper.GetString("port")
	dsn := viper.GetString("mysql_dsn")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("can't connect database")
	}

	chefRepo := repository.NewChefRepository(db)
	menuRepo := repository.NewMenuRepository(db)
	orderRepo := repository.NewOrderRepository(db)

	chefService := service.NewChefService(chefRepo)
	menuService := service.NewMenuService(menuRepo)
	orderService := service.NewOrderService(orderRepo)

	e := echo.New()
	delivery.AddChefRoute(chefService, e)
	delivery.AddMenuRoute(menuService, e)
	delivery.AddOrderRoute(orderService, e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
