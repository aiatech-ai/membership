package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/shinystarvn/common/pkg/config"
	"github.com/shinystarvn/membership/cmd/restapi/api"
	gmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	configFile := flag.String("config", "./", "config file")

	log.Println("starting Start ShinyTest...")
	configApp, err := config.LoadConfig(*configFile)
	if err != nil {
		log.Fatalln(err)
	}

	connectDB(configApp)

	engine := gin.Default()
	engine.Use(api.AllowAllOrigins())

	err = engine.Run(fmt.Sprintf(":%d", configApp.SeverPort))
	if err != nil {
		log.Fatalln(err)
	}
}

func connectDB(cfg *config.Config) *gorm.DB {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.MYSQLUser, cfg.MYSQLPass, cfg.MYSQLHostName, cfg.MYSQLPort, cfg.MYSQLDBName)
	db, err := gorm.Open(gmysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		}})

	if err != nil {
		fmt.Println("Error connecting to database : error=", err)
		log.Fatal(err)
		return nil
	}

	return db
}
