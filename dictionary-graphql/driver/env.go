package driver

import (
	"errors"
	"fmt"
	"os"
)

func Env(key string) (string, error) {
	message := "environment variables not set, key = %s"

	value, ok := os.LookupEnv(key)
	if !ok {
		return "", errors.New(fmt.Sprintf(message, key))
	}

	return value, nil
}
