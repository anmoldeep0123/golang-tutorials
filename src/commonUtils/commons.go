package commonUtils

import (
	"math/rand"
	"strconv"
	"time"
)

func GetRandomNumberAsStr() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Intn(9999))
}
