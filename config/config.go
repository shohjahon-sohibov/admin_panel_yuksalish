package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

const (
	// DebugMode indicates service mode is debug.
	DebugMode = "debug"
	// TestMode indicates service mode is test.
	TestMode = "test"
	// ReleaseMode indicates service mode is release.
	ReleaseMode = "release"
)

type Config struct {
	ServiceName string
	Environment string // debug, test, release
	Version     string

	HTTPPort   string
	HTTPScheme string

	MongoHost     string
	MongoPort     int
	MongoUser     string
	MongoPassword string
	MongoDatabase string

	RPCPort string

	SecretKey string

	PasscodePool   string
	PasscodeLength int

	DefaultOffset string
	DefaultLimit  string

	SMSUserLogin    string
	SMSUserPassword string
	SMSSender       string

	BotToken string

	MinioEndpoint        string
	MinioAccessKeyID     string
	MinioSecretAccessKey string
	MinioSSL             bool
}

// Load ...
func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.ServiceName = cast.ToString(getOrReturnDefaultValue("SERVICE_NAME", "admin_panel"))
	config.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", DebugMode))
	config.Version = cast.ToString(getOrReturnDefaultValue("VERSION", "1.0"))

	config.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":3000"))
	config.HTTPScheme = cast.ToString(getOrReturnDefaultValue("HTTP_SCHEME", "http")) // brat ci cd togrilab qoydim change qilish kerak emas uje

	config.RPCPort = cast.ToString(getOrReturnDefaultValue("RPC_PORT", ":5004"))

	config.SecretKey = cast.ToString(getOrReturnDefaultValue("SECRET_KEY", ""))

	config.MongoHost = cast.ToString(getOrReturnDefaultValue("MONGO_HOST", "localhost")) 
	config.MongoPort = cast.ToInt(getOrReturnDefaultValue("MONGO_PORT", 27017)) 
	config.MongoUser = cast.ToString(getOrReturnDefaultValue("MONGO_USER", "shohjahon"))
	config.MongoPassword = cast.ToString(getOrReturnDefaultValue("MONGO_PASSWORD", "1"))
	config.MongoDatabase = cast.ToString(getOrReturnDefaultValue("MONGO_DATABASE", "admin_panel"))

	// config.MinioAccessKeyID = cast.ToString(getOrReturnDefaultValue("MINIO_ACCESS_KEY", "access_key"))
	// config.MinioSecretAccessKey = cast.ToString(getOrReturnDefaultValue("MINIO_SECRET_KEY", "secret_key"))
	// config.MinioEndpoint = cast.ToString(getOrReturnDefaultValue("MINIO_ENDPOINT", "url"))
	// config.MinioSSL = cast.ToBool(getOrReturnDefaultValue("MINIO_SSL", true))

	config.DefaultOffset = cast.ToString(getOrReturnDefaultValue("DEFAULT_OFFSET", "0")) // 64a92d75a4135b099e1679e3
	config.DefaultLimit = cast.ToString(getOrReturnDefaultValue("DEFAULT_LIMIT", "10")) // 64a85eaea4135b099e167977
	// config.SMSUserLogin = cast.ToString(getOrReturnDefaultValue("SMS_USER_LOGIN", "admin@mediapark.uz"))
	// config.SMSUserPassword = cast.ToString(getOrReturnDefaultValue("SMS_USER_PASSWORD", "cok1S4MSLsmxRYfPOuTlkXsaxQOSoRAamhsmEK9o"))
	// config.SMSSender = cast.ToString(getOrReturnDefaultValue("SMS_SENDER", "Mediapark"))
	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}

//
