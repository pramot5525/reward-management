package config

import "github.com/spf13/viper"

type Config struct {
	App      AppConfig
	Database DatabaseConfig
	AWS      AWSConfig
	Cache    CacheConfig
}

type AppConfig struct {
	Port         string `mapstructure:"APP_PORT"`
	Env          string `mapstructure:"APP_ENV"`
	OriginDomain string `mapstructure:"APP_ORIGIN_DOMAIN"`
	JWTSecret    string `mapstructure:"APP_JWT_SECRET"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"DATABASE_MYSQL_HOST"`
	Port     string `mapstructure:"DATABASE_MYSQL_PORT"`
	Username string `mapstructure:"DATABASE_MYSQL_USERNAME"`
	Password string `mapstructure:"DATABASE_MYSQL_PASSWORD"`
	Name     string `mapstructure:"DATABASE_MYSQL_NAME"`
}

type AWSConfig struct {
	S3Bucket          string `mapstructure:"AWS_S3BUCKET"`
	CDN               string `mapstructure:"AWS_CDN"`
	LocalstackEndpoint string `mapstructure:"AWS_LOCALSTACK_ENDPOINT"` // set only on local env
}

type CacheConfig struct {
	Addr     string `mapstructure:"REDIS_ADDR"`
	Password string `mapstructure:"REDIS_PASSWORD"`
	DB       int    `mapstructure:"REDIS_DB"`
}

func Load() (*Config, error) {
	viper.AutomaticEnv()

	cfg := &Config{
		App: AppConfig{
			Port:         viper.GetString("APP_PORT"),
			Env:          viper.GetString("APP_ENV"),
			OriginDomain: viper.GetString("APP_ORIGIN_DOMAIN"),
			JWTSecret:    viper.GetString("APP_JWT_SECRET"),
		},
		Database: DatabaseConfig{
			Host:     viper.GetString("DATABASE_MYSQL_HOST"),
			Port:     viper.GetString("DATABASE_MYSQL_PORT"),
			Username: viper.GetString("DATABASE_MYSQL_USERNAME"),
			Password: viper.GetString("DATABASE_MYSQL_PASSWORD"),
			Name:     viper.GetString("DATABASE_MYSQL_NAME"),
		},
		AWS: AWSConfig{
			S3Bucket:           viper.GetString("AWS_S3BUCKET"),
			CDN:                viper.GetString("AWS_CDN"),
			LocalstackEndpoint: viper.GetString("AWS_LOCALSTACK_ENDPOINT"),
		},
		Cache: CacheConfig{
			Addr:     viper.GetString("REDIS_ADDR"),
			Password: viper.GetString("REDIS_PASSWORD"),
			DB:       viper.GetInt("REDIS_DB"),
		},
	}
	return cfg, nil
}
