package constants

import "os"

var (
	DATABASE_URL = lookup("DATABASE_URL", "")
)

func lookup(key string, defaultValue string) string {
	v, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return v
}
