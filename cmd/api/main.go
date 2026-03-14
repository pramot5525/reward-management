package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"

	adapthttp "github.com/nocnoc-thailand/reward-management/internal/adapters/primary/http"
	"github.com/nocnoc-thailand/reward-management/internal/adapters/primary/http/handler"
	"github.com/nocnoc-thailand/reward-management/internal/adapters/secondary/cache"
	"github.com/nocnoc-thailand/reward-management/internal/adapters/secondary/mysql"
	s3adapter "github.com/nocnoc-thailand/reward-management/internal/adapters/secondary/s3"
	"github.com/nocnoc-thailand/reward-management/internal/core/services"
	"github.com/nocnoc-thailand/reward-management/pkg/config"
	"github.com/nocnoc-thailand/reward-management/pkg/datasource"
	"github.com/nocnoc-thailand/reward-management/pkg/logger"
)

func main() {
	// Config
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("failed to load config:", err)
	}

	// Logger
	if err := logger.Init(cfg.App.Env); err != nil {
		log.Fatal("failed to init logger:", err)
	}

	// Database
	db, err := datasource.NewMySQL(cfg.Database)
	if err != nil {
		log.Fatal("failed to connect mysql:", err)
	}

	// S3 (LocalstackEndpoint is empty on non-local envs → uses real AWS)
	s3Client, err := datasource.NewS3Client(cfg.AWS.LocalstackEndpoint)
	if err != nil {
		log.Fatal("failed to init s3:", err)
	}

	// Secondary adapters (output)
	rewardRepo := mysql.NewRewardRepository(db)
	rewardCodeRepo := mysql.NewRewardCodeRepository(db)
	txRepo := mysql.NewRewardTransactionRepository(db)
	redisClient := datasource.NewRedisClient(cfg.Cache)
	cacheAdapter := cache.NewRedisCache(redisClient)
	storageAdapter := s3adapter.NewS3Storage(s3Client, cfg.AWS.S3Bucket, cfg.AWS.CDN)

	// Services (use cases)
	rewardSvc := services.NewRewardService(rewardRepo, cacheAdapter)
	rewardCodeSvc := services.NewRewardCodeService(rewardCodeRepo, cacheAdapter)
	redeemSvc := services.NewRedeemService(rewardRepo, rewardCodeRepo, txRepo, cacheAdapter)
	bucketSvc := services.NewBucketService(storageAdapter)

	// Primary adapters (input)
	rewardHandler := handler.NewRewardHandler(rewardSvc)
	rewardCodeHandler := handler.NewRewardCodeHandler(rewardCodeSvc)
	redeemHandler := handler.NewRedeemHandler(redeemSvc)
	bucketHandler := handler.NewBucketHandler(bucketSvc)

	// Fiber app
	app := fiber.New()
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{AllowOrigins: cfg.App.OriginDomain}))

	// Routes
	router := adapthttp.NewRouter(cfg.App.JWTSecret, rewardHandler, rewardCodeHandler, redeemHandler, bucketHandler)
	router.Register(app)

	log.Fatal(app.Listen(":" + cfg.App.Port))
}
