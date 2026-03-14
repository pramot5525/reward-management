package domain

import "time"

type RewardState int

const (
	RewardStateActive   RewardState = 100
	RewardStateInactive RewardState = 101
	RewardStateDeleted  RewardState = 102
)

type ConditionBy string

const (
	ConditionByTier   ConditionBy = "tier"
	ConditionByUserID ConditionBy = "userId"
)

type Reward struct {
	ID               uint
	RedeemPoint      int
	StartDate        time.Time
	EndDate          time.Time
	State            RewardState
	ConditionBy      ConditionBy
	IsLimitRedeem    bool
	CountLimitRedeem int
	RewardInfo       RewardInfo
	Tiers            []Tier
	RewardUsers      []RewardUser
}

type RewardInfo struct {
	ID               uint
	RewardID         uint
	TitleTH          string
	SubtitleTH       string
	DescriptionTH    string
	TermConditionTH  string
	TitleEN          string
	SubtitleEN       string
	DescriptionEN    string
	TermConditionEN  string
	BannerURL        string
	LogoURL          string
	IsCodeDisplayed  bool
	IsQrDisplayed    bool
	IsBarcodeDisplayed bool
}
