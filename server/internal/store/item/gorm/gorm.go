package gorm

import (
	"flag"
	"fmt"
	"time"

	"github.com/subliker/server/internal/config"
	"github.com/subliker/server/internal/logger"
	"github.com/subliker/server/internal/models"
	"github.com/subliker/server/internal/store/item"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var migrateMode bool

func init() {
	flag.BoolVar(&migrateMode, "migrate", true, "use to run migrations")
}

type gormStore struct {
	db *gorm.DB
}

func NewMySQL(cfg config.ItemStore) (item.Store, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)

	var db *gorm.DB
	attempts := 3
	for i := range attempts {
		var err error
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		logger.Zap.Infof("Item store connection retry %d", i+1)
		if i != attempts-1 {
			time.Sleep(5 * time.Second)
		}
	}
	logger.Zap.Info("Item store connected")

	if err := db.AutoMigrate(&models.Item{}); err != nil {
		logger.Zap.Fatalf("my sql migration error: %s", err)
	}
	logger.Zap.Info("Item store migrated")

	return &gormStore{
		db: db,
	}, nil
}

func (s *gormStore) Find(items *[]models.Item) error {
	return s.db.Find(items).Error
}

func (s *gormStore) Create(item *models.Item) error {
	return s.db.Create(item).Error
}

func (s *gormStore) Close() {
	sqlDB, err := s.db.DB()
	if err != nil {
		logger.Zap.Fatalf("getting sql db error: %s", err)
	}
	sqlDB.Close()
}
