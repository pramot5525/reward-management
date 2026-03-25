package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pramot5525/reward-management/internal/core/ports/input"
)

type TierHandler struct {
	usecase input.TierUsecase
}

func NewTierHandler(usecase input.TierUsecase) *TierHandler {
	return &TierHandler{usecase: usecase}
}

type tierDetailResponse struct {
	ID       uint   `json:"id"`
	Code     string `json:"code"`
	Name     string `json:"name"`
	ImageURL string `json:"imageUrl"`
}

func (h *TierHandler) GetTiers(c *fiber.Ctx) error {
	tiers, total, err := h.usecase.GetTiers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	data := make([]tierDetailResponse, len(tiers))
	for i, t := range tiers {
		data[i] = tierDetailResponse{
			ID:       t.ID,
			Code:     t.Code,
			Name:     t.Name,
			ImageURL: t.ImageURL,
		}
	}

	return c.JSON(fiber.Map{
		"meta": fiber.Map{"totalCount": total},
		"data": data,
	})
}
