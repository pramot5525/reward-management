package datasource

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/pramot5525/reward-management/internal/adapters/secondary/mysql/model"
	"github.com/pramot5525/reward-management/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewMySQL(cfg config.DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Name,
	)
	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: gormLogger})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(
		&model.Reward{},
		&model.RewardInfo{},
		&model.RewardCode{},
		&model.RewardTransaction{},
		&model.Tier{},
		&model.RewardUser{},
	); err != nil {
		return nil, err
	}

	return db, nil
}
