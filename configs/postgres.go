package configs

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//function create gorm database recive *viper.Viper  return *gorm.DB
func CreatePostgres(config *viper.Viper) *gorm.DB {
	//get env config
	host := config.GetString("POSTGRES_HOST")
	port := config.GetString("POSTGRES_PORT")
	user := config.GetString("POSTGRES_USER")
	password := config.GetString("POSTGRES_PASSWORD")
	dbname := config.GetString("POSTGRES_DBNAME")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok", host, user, password, dbname, port)
	//create gorm db
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
