package mysql

import (
	"fmt"
	"log"
	"time"

	"github.com/chanudinho/go-api-example/pkg/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(cfg configs.EnvConfigs) {
	var database *gorm.DB
	var err error

	dbURl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4", cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName)

	for i := 1; i <= 3; i++ {
		database, err = gorm.Open(mysql.New(mysql.Config{
			DSN: dbURl,
		}), &gorm.Config{})
		if err == nil {
			break
		} else {
			log.Printf("Attempt %d: Failed to initialize database. Retrying...", i)
			time.Sleep(3 * time.Second)
		}
	}

	DB = database
}
