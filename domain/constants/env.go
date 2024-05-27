package constants

import "os"

var (
	DATABASE_URL            = lookup("DATABASE_URL", "")
	MICROSOFT_CLIENT_ID     = lookup("MICROSOFT_CLIENT_ID", "")
	MICROSOFT_CLIENT_SECRET = lookup("MICROSOFT_CLIENT_SECRET", "")
	MICROSOFT_CALLBACK_URL  = lookup("MICROSOFT_CALLBACK_URL", "")
	MICROSOFT_SCOPE         = lookup("MICROSOFT_SCOPE", "")
	MICROSOFT_TENANT        = lookup("MICROSOFT_TENANT", "")
)

func lookup(key string, defaultValue string) string {
	v, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return v
}
