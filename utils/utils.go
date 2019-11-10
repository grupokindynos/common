package utils

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func RandomString() string {
	charset := "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 50)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func UnitToSat(u float64) int64 {
	return int64(u / math.Pow10(0))
}

func SatToUnit(s int64) float64 {
	return float64(s) / math.Pow10(8)
}

func UnitToString(u float64) string {
	return fmt.Sprintf("%f", u)
}