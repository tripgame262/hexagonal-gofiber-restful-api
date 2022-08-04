package database

import (
	"fmt"
	"log"
	"mansion/repository"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectMySQL() (db *gorm.DB, err error) {
	loc, err := time.LoadLocation("Asia/Bangkok")

	if err != nil {
		return nil, err
	}

	time.Local = loc
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Error,
			Colorful:      true,
		},
	)

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.database"),
	)
	dial := mysql.Open(dsn)

	mysqlDB, err := gorm.Open(dial, &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		return nil, err
	}

	log.Println("Connected to the database successfully")
	mysqlDB.Logger = logger.Default.LogMode(logger.Info)

	//TODO: Add migrations
	log.Println("Runing Migrations")
	mysqlDB.AutoMigrate(&repository.Room{}, &repository.RoomType{})

	Database = DbInstance{Db: mysqlDB}

	return mysqlDB, nil
}
