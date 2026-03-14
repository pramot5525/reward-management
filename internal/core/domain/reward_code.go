package domain

import "time"

type RewardCode struct {
	ID          uint
	RewardID    uint
	Code        string
	IsAvailable bool
	ExpiredDate *time.Time
	CreatedBy   string
	UpdatedBy   string
}
