package sgk

func Overlap(x1, y1, s1, x2, y2, s2 float64) bool {
	if x2 > x1+s1 || x2+s2 < x1 {
		return false
	}
	if y2 > y1+s1 || y2+s2 < y1 {
		return false
	}
	return true
}
