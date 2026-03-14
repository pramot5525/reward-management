package domain

type TransactionStatus string

const (
	TransactionStatusSuccess TransactionStatus = "SUCCESS"
	TransactionStatusFailed  TransactionStatus = "FAILED"
)

type RewardTransaction struct {
	ID           uint
	RewardCodeID uint
	UserID       string
	Status       TransactionStatus
	ErrorMsg     string
}
