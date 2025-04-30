// package config

// import (
// 	"log"
// 	"os"

// 	"fmt"
// 	"github.com/joho/godotenv"
// )

// type Config struct {
// 	Port string
// }

// var AppConfig Config

// func LoadEnv() {
// 	// .env File: This is a file that you typically keep in the root of your project (e.g., .env). It contains key-value pairs
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

//		AppConfig = Config{
//			Port: os.Getenv("PORT"),
//		}
//		fmt.Printf("Port: %v\n", AppConfig.Port)
//	}
package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port     string
	MongoURI string
	DBName   string
}

var AppConfig Config
var envLoaded bool // Flag to track if env variables are already loaded

// LoadEnv loads the environment variables from the .env file, but only once
func LoadEnv() {
	if envLoaded {
		return // Exit if already loaded
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	AppConfig = Config{
		Port:     os.Getenv("PORT"),
		MongoURI: os.Getenv("MONGODB_URI"),
		DBName:   os.Getenv("MONGODB_DBNAME"),
	}

	// // Debugging: Print all environment variables to see if MONGODB_URI is loaded
	// fmt.Println("🌍 Environment variables loaded:")
	// fmt.Println("MONGODB_URI:", os.Getenv("MONGODB_URI"))

	fmt.Printf("🌍 Loaded config: Port=%s, DB=%s\n", AppConfig.Port, AppConfig.DBName)
	// fmt.Printf("🌍 Mongo URI: %s\n", AppConfig.MongoURI)
	envLoaded = true
}

func GetPort() string {
	return AppConfig.Port
}

func GetMongoURI() string {
	return AppConfig.MongoURI
}

func GetDBName() string {
	return AppConfig.DBName
}
