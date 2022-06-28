package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	//create gin server
	r := gin.Default()
	//gin cors middlewares
	r.Use(cors.Default())
	//viper get env config form .env file in root dir
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	//get env config
	port := viper.GetString("PORT")
	//run server
	r.Run(":" + port)
}
