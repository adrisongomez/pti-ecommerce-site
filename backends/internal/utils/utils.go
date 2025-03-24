package utils

import "os"

func IsProduction() bool {
	return os.Getenv("APP_ENV") == "production"
}

func MinInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}
