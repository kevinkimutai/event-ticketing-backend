package utils

import "time"

func GenerateUniqueNumber() int64 {
	return time.Now().UnixNano()
}
