package db

import (
	"flag"
	"fmt"
	"time"

	"github.com/subliker/server/internal/config"
	"github.com/subliker/server/internal/logger"
	"github.com/subliker/server/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MigrateMode bool

func init() {
	flag.BoolVar(&MigrateMode, "migrate", true, "use to run migrations")
}

func NewMySQL(cfg config.DBConfig, gcfg *gorm.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)

	var db *gorm.DB
	var err error
	attempts := 3
	for i := range 3 {
		db, err = gorm.Open(mysql.Open(dsn), gcfg)
		if err == nil {
			logger.Zap.Info("DB connected")
			break
		}
		logger.Zap.Infof("DB connection retry %d", i+1)
		if i != attempts-1 {
			time.Sleep(5 * time.Second)
		}
	}

	return db, err
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Item{})
}
