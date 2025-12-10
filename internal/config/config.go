package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Port    int
	Env     string
	Version string

	Database DatabaseConfig

	Redis RedisConfig

	Storage StorageConfig

	Kafka KafkaConfig

	JWT JWTConfig

	LogLevel  string
	LogFormat string

	RateLimit RateLimitConfig
}

type DatabaseConfig struct {
	Host            string
	Port            int
	Name            string
	User            string
	Password        string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
	SSLMode         string
}

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

type StorageConfig struct {
	Type      string
	Endpoint  string
	AccessKey string
	SecretKey string
	UseSSL    bool
	Bucket    string
	Region    string
}

type KafkaConfig struct {
	Brokers       []string
	TopicPrefix   string
	ConsumerGroup string
}

type JWTConfig struct {
	Secret             string
	ExpiryHours        int
	RefreshExpiryHours int
}

type RateLimitConfig struct {
	Enabled         bool
	RequestsPerHour int
}

func Load() *Config {
	_ = godotenv.Load()

	return &Config{
		Port:    getEnvInt("PORT", 8080),
		Env:     getEnvString("ENV", "development"),
		Version: getEnvString("API_VERSION", "v1"),

		Database: DatabaseConfig{
			Host:            getEnvString("DB_HOST", "localhost"),
			Port:            getEnvInt("DB_PORT", 5432),
			Name:            getEnvString("DB_NAME", "spotify_db"),
			User:            getEnvString("DB_USER", "postgres"),
			Password:        getEnvString("DB_PASSWORD", "postgres"),
			MaxIdleConns:    getEnvInt("DB_MAX_IDLE_CONNS", 10),
			MaxOpenConns:    getEnvInt("DB_MAX_OPEN_CONNS", 100),
			ConnMaxLifetime: time.Duration(getEnvInt("DB_CONN_MAX_LIFETIME", 3600)) * time.Second,
			SSLMode:         getEnvString("DB_SSL_MODE", "disable"),
		},

		Redis: RedisConfig{
			Host:     getEnvString("REDIS_HOST", "localhost"),
			Port:     getEnvInt("REDIS_PORT", 6379),
			Password: getEnvString("REDIS_PASSWORD", ""),
			DB:       getEnvInt("REDIS_DB", 0),
		},

		Storage: StorageConfig{
			Type:      getEnvString("STORAGE_TYPE", "minio"),
			Endpoint:  getEnvString("MINIO_ENDPOINT", "localhost:9000"),
			AccessKey: getEnvString("MINIO_ACCESS_KEY", "minioadmin"),
			SecretKey: getEnvString("MINIO_SECRET_KEY", "minioadmin"),
			UseSSL:    getEnvBool("MINIO_USE_SSL", false),
			Bucket:    getEnvString("MINIO_BUCKET", "spotify"),
			Region:    getEnvString("MINIO_REGION", "us-east-1"),
		},

		Kafka: KafkaConfig{
			Brokers:       []string{getEnvString("KAFKA_BROKERS", "localhost:9092")},
			TopicPrefix:   getEnvString("KAFKA_TOPIC_PREFIX", "spotify"),
			ConsumerGroup: getEnvString("KAFKA_CONSUMER_GROUP", "spotify-service"),
		},

		JWT: JWTConfig{
			Secret:             getEnvString("JWT_SECRET", "change-me-in-production"),
			ExpiryHours:        getEnvInt("JWT_EXPIRY_HOURS", 24),
			RefreshExpiryHours: getEnvInt("JWT_REFRESH_EXPIRY_HOURS", 720),
		},

		LogLevel:  getEnvString("LOG_LEVEL", "info"),
		LogFormat: getEnvString("LOG_FORMAT", "json"),

		RateLimit: RateLimitConfig{
			Enabled:         getEnvBool("RATE_LIMIT_ENABLED", true),
			RequestsPerHour: getEnvInt("RATE_LIMIT_REQUESTS", 1000),
		},
	}
}

func getEnvString(key, defaultVal string) string {
	if val, exists := os.LookupEnv(key); exists {
		return val
	}
	return defaultVal
}

func getEnvInt(key string, defaultVal int) int {
	if val, exists := os.LookupEnv(key); exists {
		if intVal, err := strconv.Atoi(val); err == nil {
			return intVal
		}
	}

	return defaultVal
}

func getEnvBool(key string, defaultVal bool) bool {
	if val, exists := os.LookupEnv(key); exists {
		return val == "true" || val == "1" || val == "yes"
	}
	return defaultVal
}
