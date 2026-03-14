package dto

import "github.com/pramot5525/reward-management/internal/core/domain"

type RewardResponse struct {
	ID               uint        `json:"id"`
	RedeemPoint      int         `json:"redeemPoint"`
	StartDate        string      `json:"startDate"`
	EndDate          string      `json:"endDate"`
	State            int         `json:"state"`
	ConditionBy      string      `json:"conditionBy"`
	IsLimitRedeem    bool        `json:"isLimitRedeem"`
	CountLimitRedeem int         `json:"countLimitRedeem"`
	Info             RewardInfoResponse `json:"info"`
}

type RewardInfoResponse struct {
	Title          string `json:"title"`
	Subtitle       string `json:"subtitle"`
	Description    string `json:"description"`
	TermCondition  string `json:"termCondition"`
	BannerURL      string `json:"bannerUrl"`
	LogoURL        string `json:"logoUrl"`
	IsCodeDisplayed    bool `json:"isCodeDisplayed"`
	IsQrDisplayed      bool `json:"isQrDisplayed"`
	IsBarcodeDisplayed bool `json:"isBarcodeDisplayed"`
}

type PaginatedRewardResponse struct {
	Items []RewardResponse `json:"items"`
	Total int64            `json:"total"`
	Page  int              `json:"page"`
	Limit int              `json:"limit"`
}

func ToRewardResponse(r domain.Reward, lang string) RewardResponse {
	info := RewardInfoResponse{
		BannerURL:          r.RewardInfo.BannerURL,
		LogoURL:            r.RewardInfo.LogoURL,
		IsCodeDisplayed:    r.RewardInfo.IsCodeDisplayed,
		IsQrDisplayed:      r.RewardInfo.IsQrDisplayed,
		IsBarcodeDisplayed: r.RewardInfo.IsBarcodeDisplayed,
	}
	if lang == "en" {
		info.Title = r.RewardInfo.TitleEN
		info.Subtitle = r.RewardInfo.SubtitleEN
		info.Description = r.RewardInfo.DescriptionEN
		info.TermCondition = r.RewardInfo.TermConditionEN
	} else {
		info.Title = r.RewardInfo.TitleTH
		info.Subtitle = r.RewardInfo.SubtitleTH
		info.Description = r.RewardInfo.DescriptionTH
		info.TermCondition = r.RewardInfo.TermConditionTH
	}
	return RewardResponse{
		ID:               r.ID,
		RedeemPoint:      r.RedeemPoint,
		StartDate:        r.StartDate.Format("2006-01-02"),
		EndDate:          r.EndDate.Format("2006-01-02"),
		State:            int(r.State),
		ConditionBy:      string(r.ConditionBy),
		IsLimitRedeem:    r.IsLimitRedeem,
		CountLimitRedeem: r.CountLimitRedeem,
		Info:             info,
	}
}
