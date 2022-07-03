package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	v = ReadInConfig()
)

func main() {
	//create gin server
	r := gin.Default()
	//gin cors middlewares
	r.Use(cors.Default())

	//get env config
	port := v.GetString("PORT")
	//run server
	r.Run(":" + port)
}

func ReadInConfig() *viper.Viper {
	v := viper.New()
	//viper get env config form .env file in root dir
	v.SetConfigFile(".env")
	v.AutomaticEnv()
	v.AddConfigPath("/etc/appname/")  // path to look for the config file in
	v.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
	v.AddConfigPath(".")              // optionally look for config in the working directory
	err := v.ReadInConfig()
	if err != nil {
		log.Fatal("Error reading config file", err)
	}
	return v
}
