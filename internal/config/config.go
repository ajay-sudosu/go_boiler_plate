package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI   string
	DBName     string
	Port       string
	LibvirtURI string
}

func Load() *Config {
	_ = godotenv.Load(".env") // silently load if exists

	return &Config{
		MongoURI:   getEnv("MONGO_URI", "mongodb://localhost:27017"),
		DBName:     getEnv("DB_NAME", "users"),
		Port:       getEnv("PORT", "8000"),
		LibvirtURI: getEnv("LIBVIRTURI", "qemu:///system"),
	}
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	log.Printf("Warning: Using default for %s\n", key)
	return fallback
}
