package foundation

// Abs for int64
func Abs(value int64) int64 {
	if value < 0 {
		return value * -1
	}
	return value

}

// ClampInt [min, max)
func ClampInt(target, min, max int) int {
	if max < min {
		panic("ClampInt Error")
	} else if max == min {
		return min
	}

	max--
	if min > target {
		target = min
	} else if max <= target {
		target = max
	}
	return target
}
