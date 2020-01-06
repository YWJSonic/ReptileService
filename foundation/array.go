package foundation

import "reflect"

// ReverseAny ...
func ReverseAny(s interface{}) {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

// ReverseArray Reverse string array
func ReverseArray(data []string) []string {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
	return data
}

// DeleteArrayElement ...
func DeleteArrayElement(elementIndex interface{}, array []interface{}) []interface{} {
	count := len(array)
	for index := 0; index < count; index++ {
		if elementIndex == array[index] {
			return append(array[:index], array[index+1:]...)
		}
	}
	return array
}

// MapIIKeys ...
func MapIIKeys(mapData map[int64]int64) (Keys []int64) {
	for key := range mapData {
		Keys = append(Keys, key)
	}
	return
}

// MapSSKeys ...
func MapSSKeys(mapData map[string]string) (Keys []string) {
	for key := range mapData {
		Keys = append(Keys, key)
	}
	return
}

// MapSIKeys ...
func MapSIKeys(mapData map[string]int64) (Keys []string) {
	for key := range mapData {
		Keys = append(Keys, key)
	}
	return
}

// MapSIAAKeys ...
func MapSIAAKeys(mapData map[string][][]int64) (Keys []string) {
	for key := range mapData {
		Keys = append(Keys, key)
	}
	return
}

// MapISAKeys ...
func MapISAKeys(mapData map[int64][]string) (Keys []int64) {
	for key := range mapData {
		Keys = append(Keys, key)
	}
	return
}

// MapIIToArray return map keys and values
func MapIIToArray(mapData map[int64]int64) (keys, values []int64) {
	for key, value := range mapData {
		keys = append(keys, key)
		values = append(values, value)
	}
	return
}

// MapSSToArray return map keys and values
func MapSSToArray(mapData map[string]string) (keys, values []string) {
	for key, value := range mapData {
		keys = append(keys, key)
		values = append(values, value)
	}
	return
}

// MapSIToArray return map keys and values
func MapSIToArray(mapData map[string]int64) (keys []string, values []int64) {
	for key, value := range mapData {
		keys = append(keys, key)
		values = append(values, value)
	}
	return
}

// AppendMap map append map
func AppendMap(Target map[string]interface{}, Source map[string]interface{}) map[string]interface{} {
	for Key, Value := range Source {
		Target[Key] = Value
	}
	return Target
}

// ArrayShift Array Type []map[string]interface{}
func ArrayShift(Target []map[string]interface{}) (map[string]interface{}, []map[string]interface{}) {

	var out map[string]interface{}
	out = Target[0]
	Target = Target[1:]

	return out, Target
}

// CopyArray new array memory array
func CopyArray(source []int) []int {
	result := make([]int, len(source))
	copy(result, source)
	return result
}
