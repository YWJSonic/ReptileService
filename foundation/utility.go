package foundation

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// InterfaceTofloat64 ...
func InterfaceTofloat64(v interface{}) float64 {
	return v.(float64)
}

// InterfaceToInt ...
func InterfaceToInt(v interface{}) int {
	switch v.(type) {
	case float64:
		return int(InterfaceTofloat64(v))
	case int:
		return v.(int)
	case int64:
		return int(v.(int64))
	default:
		panic("Conver Error")
	}
}

// InterfaceToInt64 ...
func InterfaceToInt64(v interface{}) int64 {
	switch v.(type) {
	case float64:
		return int64(v.(float64))
	case int:
		return int64(v.(int))
	case int64:
		return v.(int64)
	default:
		fmt.Print("Conver", v)
		panic("Conver Error")
	}
}

// InterfaceToBool ...
func InterfaceToBool(v interface{}) bool {
	switch v.(type) {
	case int:
		return v.(bool)
	case bool:
		return v.(bool)
	default:
		panic("Conver Error")
	}
}

// InterfaceToString ...
func InterfaceToString(v interface{}) string {
	return v.(string)
}

// MD5Code encode MD5
func MD5Code(data string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(data)))
}

// RandomMutily ...
func RandomMutily(rangeInt []int, pickCount int) []int {
	var result []int
	var targetPoint int

	if pickCount >= len(rangeInt) {
		return rangeInt
	}

	for i, imax := 0, pickCount; i < imax; i++ {
		targetPoint = rand.Intn(len(rangeInt))
		result = append(result, rangeInt[targetPoint])
		if len(rangeInt) > 1 {
			rangeInt = append(rangeInt[:targetPoint], rangeInt[targetPoint+1:]...)
		}
	}
	return result
}

// RangeRandomInt64 array random index
func RangeRandomInt64(rangeInt []int64) int {
	var Sum int64

	for _, value := range rangeInt {
		Sum += value
	}

	random := rand.Int63n(Sum)

	Sum = 0
	for i, value := range rangeInt {
		Sum += value
		if Sum > random {
			return i
		}
	}
	return -1
}

// RangeRandom array random index
func RangeRandom(rangeInt []int) int {
	Sum := 0

	for _, value := range rangeInt {
		Sum += value
	}

	random := rand.Intn(Sum)

	Sum = 0
	for i, value := range rangeInt {
		Sum += value
		if Sum > random {
			return i
		}
	}
	return -1
}

// ConevrToTimeInt64 Get time point
func ConevrToTimeInt64(year int, month time.Month, day, hour, min, sec, nsec int) int64 {
	return time.Date(year, month, day, hour, min, sec, nsec, time.Local).Unix()
}
