package mysql

import (
	"fmt"
	"log"
	"time"

	"github.com/chanudinho/go-api-example/pkg/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg configs.EnvConfigs) *gorm.DB {
	var database *gorm.DB
	var err error

	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	for i := 1; i <= 3; i++ {
		database, err = gorm.Open(mysql.New(mysql.Config{
			DSN: dbURL,
		}), &gorm.Config{})
		if err == nil {
			break
		} else {
			log.Printf("Attempt %d: Failed to initialize database. Retrying...", i)
			time.Sleep(3 * time.Second)
		}
	}

	return database
}
