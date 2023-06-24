package main

import (
	"alqinsidev/auth-service/cmd"
	"alqinsidev/auth-service/config"
	"alqinsidev/auth-service/utils"

	"github.com/spf13/viper"
)

func main() {
	cfg := config.EnvConfig()
	db := utils.NewDBConnection(cfg)

	r := cmd.SetupRouter(db.Gorm)

	port := ":" + viper.GetString("APP_PORT")
	r.Run(port)

}
