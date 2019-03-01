package params

import (
	"log"
	"os"
)

func getEnv(key string, def string) string {
	rez := os.Getenv(key)
	if rez == "" {
		rez = def
	}
	return rez
}
func getEnvFatal(key string) string {
	rez := os.Getenv(key)
	if rez == "" {
		log.Fatal(key + " must be set in ENV variables.")
	}
	return rez
}
func HttpPort() string {
	return getEnv("HTTP_PORT", http_port)
}
func MongoUrl() string {
	return getEnvFatal("MONGO_URL")
}
