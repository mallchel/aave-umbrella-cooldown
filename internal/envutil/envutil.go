package envutil

import "os"

// Get returns environment variable value or fallback when empty.
func Get(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
