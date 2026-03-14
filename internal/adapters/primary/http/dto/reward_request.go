package dto

type CreateRewardRequest struct {
	RedeemPoint      int         `json:"redeemPoint" validate:"required,min=1"`
	StartDate        string      `json:"startDate" validate:"required"`
	EndDate          string      `json:"endDate" validate:"required"`
	ConditionBy      string      `json:"conditionBy" validate:"required,oneof=tier userId"`
	IsLimitRedeem    bool        `json:"isLimitRedeem"`
	CountLimitRedeem int         `json:"countLimitRedeem" validate:"min=0,max=999"`
	Tiers            []string    `json:"tiers"`
	UserIDs          []string    `json:"userIds"`
	InfoTH           RewardInfoRequest `json:"infoTH" validate:"required"`
	InfoEN           RewardInfoRequest `json:"infoEN" validate:"required"`
}

type RewardInfoRequest struct {
	Title          string `json:"title" validate:"required"`
	Subtitle       string `json:"subtitle"`
	Description    string `json:"description"`
	TermCondition  string `json:"termCondition"`
	BannerURL      string `json:"bannerUrl"`
	LogoURL        string `json:"logoUrl"`
	IsCodeDisplayed    bool `json:"isCodeDisplayed"`
	IsQrDisplayed      bool `json:"isQrDisplayed"`
	IsBarcodeDisplayed bool `json:"isBarcodeDisplayed"`
}

type UpdateRewardRequest struct {
	ID               uint        `json:"id" validate:"required"`
	RedeemPoint      int         `json:"redeemPoint"`
	StartDate        string      `json:"startDate"`
	EndDate          string      `json:"endDate"`
	IsLimitRedeem    bool        `json:"isLimitRedeem"`
	CountLimitRedeem int         `json:"countLimitRedeem" validate:"min=0,max=999"`
	InfoTH           *RewardInfoRequest `json:"infoTH"`
	InfoEN           *RewardInfoRequest `json:"infoEN"`
}
