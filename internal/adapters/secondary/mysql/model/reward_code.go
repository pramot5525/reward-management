package model

import (
	"time"

	"gorm.io/gorm"
)

type RewardCode struct {
	gorm.Model
	RewardID    uint
	Code        string
	IsAvailable bool
	ExpiredDate *time.Time
	CreatedBy   string
	UpdatedBy   string
}

type RewardTransaction struct {
	gorm.Model
	RewardCodeID uint
	UserID       string
	Status       string
	ErrorMsg     string
}

type Tier struct {
	gorm.Model
	RewardID uint
	TierID   string
}

type RewardUser struct {
	gorm.Model
	RewardID uint
	UserID   string
}

type LoyaltyTier struct {
	gorm.Model
	Code     string `gorm:"uniqueIndex;size:50"`
	Name     string `gorm:"size:100"`
	ImageURL string `gorm:"size:500"`
}
