package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configuration *Config

type DBConfig struct {
	Host          string
	Port          int
	Name          string
	User          string
	Password      string
	EnableSSLMODE bool
}

type Config struct {
	Version      string
	ServiceName  string
	HttpPort     int
	JwtSecretkey string
	DB           *DBConfig
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Faild to load env variables:", err)
		os.Exit(1)

	}
	version := os.Getenv("VERSION")

	if version == "" {
		fmt.Println("Version is required")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")

	if serviceName == "" {
		fmt.Println("Service name is required")
		os.Exit(1)

	}

	httpPort := os.Getenv("HTTP_PORT")

	if httpPort == "" {
		fmt.Println("Http Port is required")
		os.Exit(1)

	}

	port, err := strconv.ParseInt(httpPort, 10, 64)

	if err != nil {

		fmt.Println("port must be Number")
		os.Exit(1)

	}

	jwtSecretkey := os.Getenv("JWT_SECRET_KEY")

	if jwtSecretkey == "" {
		fmt.Println("Jwt secret key is required")
		os.Exit(1)

	}

	dbHost := os.Getenv("DB_HOST")

	if dbHost == "" {
		fmt.Println("Db Host is required")
		os.Exit(1)

	}

	dbPort := os.Getenv("DB_PORT")

	if dbPort == "" {
		fmt.Println("Db Port is required")
		os.Exit(1)

	}

	dbprt, err := strconv.ParseInt(dbPort, 10, 64)

	if err != nil {

		fmt.Println("port must be Number")
		os.Exit(1)

	}

	dbName := os.Getenv("DB_NAME")

	if dbName == "" {
		fmt.Println("Db Name is required")
		os.Exit(1)

	}

	dbUser := os.Getenv("DB_USER")

	if dbUser == "" {
		fmt.Println("Db User is required")
		os.Exit(1)

	}

	dbPasword := os.Getenv("DB_PASSWORD")

	if dbPasword == "" {
		fmt.Println("Db Password is required")
		os.Exit(1)

	}

	enableSslMode := os.Getenv("DB_ENABLE_SSL_MODE")

	enablrSSLModel, err := strconv.ParseBool(enableSslMode)
	if err != nil {
		fmt.Println("Invalid enable ssl mode value", err)
		os.Exit(1)

	}

	dbConfig := &DBConfig{
		Host:          dbHost,
		Port:          int(dbprt),
		Name:          dbName,
		User:          dbUser,
		Password:      dbPasword,
		EnableSSLMODE: enablrSSLModel,
	}

	configuration = &Config{
		Version:      version,
		ServiceName:  serviceName,
		HttpPort:     int(port),
		JwtSecretkey: jwtSecretkey,
		DB:           dbConfig,
	}

}

func GetConfig() *Config {
	if configuration == nil {
		//first time
		loadConfig()
	}
	return configuration
}
