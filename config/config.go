package config

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

type AppConfig struct {
	Name            string
	Port            string
	NumberOfWorkers int
}

type DbConfig struct {
	Host            string
	Port            string
	User            string
	Pass            string
	Schema          string
	MaxIdleConn     int
	MaxOpenConn     int
	MaxConnLifetime time.Duration
	Debug           bool
}

type RedisConfig struct {
	Host               string
	Port               string
	Pass               string
	Db                 int
	MandatoryPrefix    string
	AccessUuidPrefix   string
	RefreshUuidPrefix  string
	UserPrefix         string
	PermissionPrefix   string
	UserCacheTTL       time.Duration
	PermissionCacheTTL time.Duration
}

type AsynqConfig struct {
	RedisAddr          string
	DB                 int
	Pass               string
	Concurrency        int
	Queue              string
	Retention          time.Duration // in hours
	RetryCount         int
	Delay              time.Duration // in seconds
	StockSyncTaskDelay time.Duration // in seconds
}

type JwtConfig struct {
	AccessTokenSecret  string
	RefreshTokenSecret string
	AccessTokenExpiry  time.Duration
	RefreshTokenExpiry time.Duration
}

type LoggerConfig struct {
	Level    string
	FilePath string
}

type EmailConfig struct {
	Url     string
	Timeout time.Duration
}

type Config struct {
	App    *AppConfig
	DB     *DbConfig
	Redis  *RedisConfig
	Asynq  *AsynqConfig
	Logger *LoggerConfig
	Jwt    *JwtConfig
	Email  *EmailConfig
}

var config Config

func GetAll() Config {
	return config
}

func App() *AppConfig {
	return config.App
}

func Db() *DbConfig {
	return config.DB
}

func Redis() *RedisConfig {
	return config.Redis
}

func Asynq() *AsynqConfig {
	return config.Asynq
}

func Logger() *LoggerConfig {
	return config.Logger
}

func Jwt() *JwtConfig {
	return config.Jwt
}

func Email() *EmailConfig {
	return config.Email
}

func LoadConfig() {
	setDefaultConfig()

	const (
		ENV_CONSUL_URL  = "CONSUL_URL"
		ENV_CONSUL_PATH = "CONSUL_PATH"
	)

	_ = viper.BindEnv(ENV_CONSUL_URL)
	_ = viper.BindEnv(ENV_CONSUL_PATH)

	consulURL := viper.GetString(ENV_CONSUL_URL)
	consulPath := viper.GetString(ENV_CONSUL_PATH)

	if consulURL == "" {
		consulURL = "http://localhost:8500"
	}
	if consulPath == "" {
		consulPath = "inventory-management"
	}

	if err := viper.AddRemoteProvider("consul", consulURL, consulPath); err != nil {
		log.Printf("[WARN] Could not add remote provider: %v. Using default config.", err)
		return
	}

	viper.SetConfigType("json")
	if err := viper.ReadRemoteConfig(); err != nil {
		log.Printf("[WARN] Failed to read config from Consul: %v. Using default config.", err)
		return
	}

	var remoteConfig Config
	if err := viper.Unmarshal(&remoteConfig); err != nil {
		log.Printf("[WARN] Failed to unmarshal remote config: %v. Using default config.", err)
		return
	}

	config = remoteConfig

	if r, err := json.MarshalIndent(&config, "", "  "); err == nil {
		log.Println("[INFO] Successfully loaded config from Consul:")
		fmt.Println(string(r))
	}
}

func setDefaultConfig() {
	config.App = &AppConfig{
		Name:            "inventory-management-service",
		Port:            "8080",
		NumberOfWorkers: 5,
	}

	config.DB = &DbConfig{
		Host:            "192.168.56.106",
		Port:            "3306",
		User:            "root",
		Pass:            "password",
		Schema:          "inventory_management",
		MaxIdleConn:     1,
		MaxOpenConn:     2,
		MaxConnLifetime: 30,
		Debug:           true,
	}

	config.Redis = &RedisConfig{
		Host:               "192.168.56.106",
		Port:               "6379",
		Pass:               "password",
		Db:                 2,
		MandatoryPrefix:    "inventory_management_",
		AccessUuidPrefix:   "access-uuid_",
		RefreshUuidPrefix:  "refresh-uuid_",
		UserPrefix:         "user_",
		PermissionPrefix:   "permissions_",
		UserCacheTTL:       3600,
		PermissionCacheTTL: 86400,
	}

	config.Asynq = &AsynqConfig{
		RedisAddr:          "192.168.56.106:6379",
		DB:                 15,
		Pass:               "password",
		Concurrency:        10,
		Queue:              "inventory_management",
		Retention:          168,
		RetryCount:         3,
		Delay:              0,
		StockSyncTaskDelay: 10,
	}

	config.Logger = &LoggerConfig{
		Level:    "debug",
		FilePath: "app.log",
	}

	config.Jwt = &JwtConfig{
		AccessTokenSecret:  "access_token",
		RefreshTokenSecret: "refresh_token",
		AccessTokenExpiry:  3600,
		RefreshTokenExpiry: 3600,
	}

	config.Email = &EmailConfig{
		Url:     "https://webhook.site/b2218104-ab43-4b22-949f-7796be528019/email",
		Timeout: 5,
	}
}
