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

	config.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":8080"))
	config.HTTPScheme = cast.ToString(getOrReturnDefaultValue("HTTP_SCHEME", "http"))

	config.MongoHost = cast.ToString(getOrReturnDefaultValue("MONGO_HOST", "cluster0.q1xijde")) 
	config.MongoPort = cast.ToInt(getOrReturnDefaultValue("MONGO_PORT", 27017)) 
	config.MongoUser = cast.ToString(getOrReturnDefaultValue("MONGO_USER", "shohjahon"))
	config.MongoPassword = cast.ToString(getOrReturnDefaultValue("MONGO_PASSWORD", "9suQpK8sFAQPOMLF"))
	config.MongoDatabase = cast.ToString(getOrReturnDefaultValue("MONGO_DATABASE", "admin_panel"))


	config.DefaultOffset = cast.ToString(getOrReturnDefaultValue("DEFAULT_OFFSET", "0")) // 64a92d75a4135b099e1679e3
	config.DefaultLimit = cast.ToString(getOrReturnDefaultValue("DEFAULT_LIMIT", "10")) // 64a85eaea4135b099e167977
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
