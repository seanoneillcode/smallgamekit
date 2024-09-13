package smallgamekit

import "math"

func Normalize(x float64, y float64) (float64, float64) {
	var nx, ny float64
	norm2 := x*x + y*y
	if norm2 == 0 {
		return nx, ny
	}
	ratio := (1 / math.Sqrt(norm2))
	nx = x * ratio
	ny = y * ratio
	return nx, ny
}
