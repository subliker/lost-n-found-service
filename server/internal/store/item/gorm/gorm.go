package gorm

import (
	"fmt"
	"time"

	"github.com/subliker/server/internal/config"
	"github.com/subliker/server/internal/logger"
	"github.com/subliker/server/internal/models"
	"github.com/subliker/server/internal/store/item"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type gormStore struct {
	db *gorm.DB
}

// NewMySQL creates new instance of item store with MySQL database
func NewMySQL(cfg config.ItemStore) (item.Store, error) {
	// my sql connection string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)

	// opening sql connection
	var db *gorm.DB
	var err error

	maxRetries := 5
	delay := 3 * time.Second

	for i := range maxRetries {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}

		logger.Zap.Infof("Item store connection retry %d of %d", i+1, maxRetries)

		// fix ex delay
		if i == maxRetries-1 {
			break
		}
		time.Sleep(delay)
	}

	if err != nil {
		logger.Zap.Fatalf("Item store connection fail after %d retries", maxRetries)
	} else {
		logger.Zap.Info("Item store connected")
	}

	// migration if needed
	if cfg.Migration {
		if err := db.AutoMigrate(&models.Item{}); err != nil {
			logger.Zap.Fatalf("my sql migration error: %s", err)
		}
		logger.Zap.Info("Item store migrated")
	}

	return &gormStore{
		db: db,
	}, nil
}

// Find fills items array by items from database
func (s *gormStore) Find(items *[]models.Item) error {
	return s.db.Find(items).Error
}

// Find fills items array by items from database with pagination
func (s *gormStore) FindWithPagination(items *[]models.Item, page int, pageSize int) error {
	return s.db.Offset((page - 1) * pageSize).Limit(pageSize).Find(items).Error
}

// Create creates new item row in database
func (s *gormStore) Create(item *models.Item) error {
	return s.db.Create(item).Error
}

// Close closes sql connection
func (s *gormStore) Close() {
	sqlDB, err := s.db.DB()
	if err != nil {
		logger.Zap.Fatalf("getting sql db error: %s", err)
	}
	sqlDB.Close()
}
