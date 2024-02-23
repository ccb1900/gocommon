package env

import (
	"fmt"

	"github.com/joho/godotenv"
)

// stage: .env <production < test < development<local
func InitEnv(stage string) error {
	if err := godotenv.Overload(); err != nil {
		return fmt.Errorf("Error loading .env file,err=%w", err)
	}
	stages := []string{"production", "test", "development"}
	for i := 0; i < len(stages); i++ {
		if stage == stages[i] {
			if err := godotenv.Overload(fmt.Sprintf(".env.%s", stage)); err != nil {
				return fmt.Errorf("Error loading .env file,stage=%s,err=%w", stage, err)
			}
		}
	}

	if err := godotenv.Overload(fmt.Sprintf(".env.%s", "local")); err != nil {
		return fmt.Errorf("Error loading .env file,stage=%s,err=%w", stage, err)
	}

	return nil
}
