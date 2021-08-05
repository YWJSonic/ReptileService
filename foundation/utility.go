package foundation

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"reflect"
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
		panic(fmt.Sprintf("Conver Error: %v, %v", v, reflect.TypeOf(v)))
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
		panic(fmt.Sprintf("Conver Error: %v, %v", v, reflect.TypeOf(v)))
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
		panic(fmt.Sprintf("Conver Error: %v, %v", v, reflect.TypeOf(v)))
	}
}

// InterfaceToString ...
func InterfaceToString(v interface{}) string {
	if v == nil {
		return ""
	} else {
		switch v.(type) {
		case string:
			return v.(string)
		case float64:
			return fmt.Sprint(v.(float64))
		case int64:
			return fmt.Sprint(v.(int64))
		default:
			panic(fmt.Sprintf("Conver Error: %v, %v", v, reflect.TypeOf(v)))
		}
	}
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

// StringColor red and green
func StringColor(Str, Color string) string {
	if "red" == Color || "Red" == Color {
		return fmt.Sprintf("\033[1;31m%s\033[0m", Str)
	} else if "green" == Color || "Green" == Color {
		return fmt.Sprintf("\033[1;32m%s\033[0m", Str)
	}
	return Str
}

// IStringArrayColor red and green
func IStringArrayColor(IValues interface{}, Color string) []interface{} {
	var result []interface{}

	Values := reflect.ValueOf(IValues)
	switch Values.Kind() {
	case reflect.Slice:
		if "red" == Color || "Red" == Color {
			for i := 0; i < Values.Len(); i++ {
				result = append(result, fmt.Sprintf("\033[1;31m%s\033[0m", Values.Index(i)))
			}
		} else if "green" == Color || "Green" == Color {
			for i := 0; i < Values.Len(); i++ {
				result = append(result, fmt.Sprintf("\033[1;32m%s\033[0m", Values.Index(i)))
			}
		} else {
			for i := 0; i < Values.Len(); i++ {
				result = append(result, fmt.Sprint(Values.Index(i)))
			}
		}
	}

	return result
}

// Float32ArrayColor red and green
func Float32ArrayColor(Color string, Values ...float32) []interface{} {
	var result []interface{}
	if "red" == Color || "Red" == Color {
		for _, value := range Values {
			result = append(result, fmt.Sprintf("\033[1;31m%f\033[0m", value))
		}
	} else if "green" == Color || "Green" == Color {
		for _, value := range Values {
			result = append(result, fmt.Sprintf("\033[1;32m%f\033[0m", value))
		}
	} else {
		for _, value := range Values {
			result = append(result, fmt.Sprint(value))
		}
	}

	return result
}

// IsAfterNowTime ...
func IsAfterNowTime(Year, Month, Day int) bool {
	return time.Date(Year, time.Month(Month), Day, 0, 0, 0, 0, time.UTC).After(time.Now())
}

// IsBeforeNowTime ...
func IsBeforeNowTime(Year, Month, Day int) bool {
	return time.Date(Year, time.Month(Month), Day, 0, 0, 0, 0, time.UTC).Before(time.Now())
}
