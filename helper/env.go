package helper

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnv() error {

	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("failed to load env : %v", err)
	}

	return nil

}
