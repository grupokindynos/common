package conversion

import (
	"fmt"
	"math"
)

func UnitToSat(u float64) int64 {
	return int64(u / math.Pow10(0))
}

func SatToUnit(s int64) float64 {
	return float64(s) / math.Pow10(8)
}

func UnitToString(u float64) string {
	return fmt.Sprintf("%f", u)
}