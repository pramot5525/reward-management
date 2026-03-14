package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/nocnoc-thailand/reward-management/internal/core/ports/input"
)

type BucketHandler struct {
	usecase input.BucketUsecase
}

func NewBucketHandler(usecase input.BucketUsecase) *BucketHandler {
	return &BucketHandler{usecase: usecase}
}

func (h *BucketHandler) UploadImage(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "file is required"})
	}

	f, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer f.Close()

	data := make([]byte, file.Size)
	if _, err := f.Read(data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	contentType := file.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	key := fmt.Sprintf("uploads/%s", file.Filename)
	url, err := h.usecase.UploadImage(key, data, contentType)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"url": url})
}

func (h *BucketHandler) Ping(c *fiber.Ctx) error {
	if err := h.usecase.PingStorage(); err != nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"status": "down", "error": err.Error()})
	}
	return c.JSON(fiber.Map{"status": "up"})
}
