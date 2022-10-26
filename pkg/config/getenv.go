package config

import (
	"errors"
	"os"
	"strconv"
)

func getEnv() (int, int, int, error) {
	messagePerSecond := 100
	messagePerSecondEnv, messagePerSecondHas := os.LookupEnv("MESSAGE_PER_SECOND")
	if messagePerSecondHas {
		messagePerSecondInt, err := strconv.Atoi(messagePerSecondEnv)
		if err != nil {
			return 0, 0, 0, errors.New("MESSAGE_PER_SECOND should be INT")
		}
		messagePerSecond = messagePerSecondInt
	}
	if messagePerSecond > 1000 {
		return 0, 0, 0, errors.New("Environment 'MESSAGE_SIZE' can't be more then 1000")
	}

	messageSize := -1
	messageSizeEnv, messageSizeHas := os.LookupEnv("MESSAGE_SIZE")
	if messageSizeHas {
		messageSizeInt, err := strconv.Atoi(messageSizeEnv)
		if err != nil {
			return 0, 0, 0, errors.New("MESSAGE_SIZE should be INT")
		}
		messageSize = messageSizeInt
	}

	messageSizeFrom := -1
	messageSizeFromEnv, messageSizeFromHas := os.LookupEnv("MESSAGE_SIZE_FROM")
	if messageSizeFromHas {
		messageSizeFromInt, err := strconv.Atoi(messageSizeFromEnv)
		if err != nil {
			return 0, 0, 0, errors.New("MESSAGE_SIZE_FROM should be INT")
		}
		messageSizeFrom = messageSizeFromInt
	}

	messageSizeTo := -1
	messageSizeToEnv, messageSizeToHas := os.LookupEnv("MESSAGE_SIZE_TO")
	if messageSizeToHas {
		messageSizeToInt, err := strconv.Atoi(messageSizeToEnv)
		if err != nil {
			return 0, 0, 0, errors.New("MESSAGE_SIZE should be INT")
		}
		messageSizeTo = messageSizeToInt
	}

	from, to, err := getSizeRange(messageSizeHas, messageSizeFromHas, messageSizeToHas, messageSize, messageSizeFrom, messageSizeTo)
	if err != nil {
		return 0, 0, 0, err
	}

	return messagePerSecond, from, to, nil

}

// Get sizeFrom and sizeTo from environment
func getSizeRange(sizeHas, sizeFromHas, sizeToHas bool, size, sizeFrom, sizeTo int) (int, int, error) {
	if !sizeHas && !sizeFromHas && !sizeToHas {
		return 0, 0, errors.New("Environment 'MESSAGE_SIZE' or 'MESSAGE_SIZE_FROM' and 'MESSAGE_SIZE_TO' must be specified")
	} else if sizeHas && (sizeFromHas || sizeToHas) {
		return 0, 0, errors.New("Environment 'MESSAGE_SIZE' can't work with environment 'MESSAGE_SIZE_FROM' or 'MESSAGE_SIZE_TO'")
	} else if sizeFrom > sizeTo {
		return 0, 0, errors.New("Environment 'MESSAGE_SIZE_FROM' can't be more or equals then 'MESSAGE_SIZE_TO'")
	} else if sizeHas && !sizeFromHas && !sizeToHas {
		return size, size, nil
	} else if !sizeHas && sizeFromHas && sizeToHas {
		return sizeFrom, sizeTo, nil
	}
	return 0, 0, errors.New("Both environment 'MESSAGE_SIZE_FROM' or 'MESSAGE_SIZE_TO' must be specified, or use 'MESSAGE_SIZE'")
}
