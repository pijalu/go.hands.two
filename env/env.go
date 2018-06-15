package env

import "os"

// GetEnvWithDefault returns a env value or def if not set
func GetEnvWithDefault(key string, def string) string {
	if val, present := os.LookupEnv(key); present {
		return val
	}
	return def
}
