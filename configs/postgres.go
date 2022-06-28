package configs

import (
	"github.com/spf13/viper"
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
	//create gorm db
	db, err := gorm.Open("postgres", "host="+host+" port="+port+" user="+user+" dbname="+dbname+" password="+password)
	if err != nil {
		panic(err)
	}
	return db
}
