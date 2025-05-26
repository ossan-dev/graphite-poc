package config

import "os"

func GetEnvOrDefault(name, value string) string {
	key, isFound := os.LookupEnv(name)
	if isFound {
		return key
	}
	return value
}
