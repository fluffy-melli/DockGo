package DockGo

import (
	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		Error(ERROR, "\033[41m\033[33m%v\033[0m", err)
	}
}
