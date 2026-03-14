package model

import (
	"time"

	"gorm.io/gorm"
)

type Reward struct {
	gorm.Model
	RedeemPoint      int
	StartDate        time.Time
	EndDate          time.Time
	State            int
	ConditionBy      string
	IsLimitRedeem    bool
	CountLimitRedeem int
	RewardInfo       RewardInfo   `gorm:"foreignKey:RewardID"`
	Tiers            []Tier       `gorm:"foreignKey:RewardID"`
	RewardUsers      []RewardUser `gorm:"foreignKey:RewardID"`
}

type RewardInfo struct {
	gorm.Model
	RewardID           uint
	TitleTH            string
	SubtitleTH         string
	DescriptionTH      string
	TermConditionTH    string
	TitleEN            string
	SubtitleEN         string
	DescriptionEN      string
	TermConditionEN    string
	BannerURL          string
	LogoURL            string
	IsCodeDisplayed    bool
	IsQrDisplayed      bool
	IsBarcodeDisplayed bool
}
