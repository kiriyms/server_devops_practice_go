package common

import (
	"fmt"
	"math/rand"
	"time"
)

func GetUserId() string {
	timestamp := time.Now().Format("20060102150405")
	rnd := rand.Intn(1000)
	return fmt.Sprintf("%s_%v", timestamp, rnd)
}
