package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nocnoc-thailand/reward-management/internal/adapters/primary/http/handler"
	"github.com/nocnoc-thailand/reward-management/internal/adapters/primary/http/middleware"
)

type Router struct {
	jwtSecret  string
	reward     *handler.RewardHandler
	rewardCode *handler.RewardCodeHandler
	redeem     *handler.RedeemHandler
	bucket     *handler.BucketHandler
}

func NewRouter(
	jwtSecret string,
	reward *handler.RewardHandler,
	rewardCode *handler.RewardCodeHandler,
	redeem *handler.RedeemHandler,
	bucket *handler.BucketHandler,
) *Router {
	return &Router{jwtSecret: jwtSecret, reward: reward, rewardCode: rewardCode, redeem: redeem, bucket: bucket}
}

func (r *Router) Register(app *fiber.App) {
	app.Get("/actuator/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "UP"})
	})

	bucket := app.Group("/bucket")
	bucket.Post("/image", r.bucket.UploadImage)
	bucket.Get("/ping", r.bucket.Ping)

	api := app.Group("/api/v1", middleware.JWTAuth(r.jwtSecret))

	reward := api.Group("/reward")
	reward.Post("", r.reward.Create)
	reward.Put("", r.reward.Update)
	reward.Get("", r.reward.GetAvailable)
	reward.Get("/list", r.reward.GetList)
	reward.Get("/:rewardId", r.reward.GetByID)
	reward.Delete("/:rewardId", r.reward.Delete)

	reward.Post("/code", r.rewardCode.Create)
	reward.Put("/code", r.rewardCode.Update)
	reward.Get("/code", r.rewardCode.GetByRewardID)

	reward.Post("/redeem", r.redeem.Redeem)
	reward.Get("/redeem", r.redeem.GetUserRedeemed)
}
