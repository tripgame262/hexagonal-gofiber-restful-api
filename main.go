package main

import (
	"fmt"
	"mansion/database"
	"mansion/routes"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main() {
	//Read config
	initConfig()

	db, err := database.ConnectMySQL()

	if err != nil {
		panic(err)
	}

	app := fiber.New()

	routes.InitRoute(app, db)

	err = app.Listen(fmt.Sprintf(":%v", viper.GetString("app.port")))
	if err != nil {
		panic(err)
	}
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}
}
