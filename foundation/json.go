package foundation

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ToJSONByte Convert to json byte
func ToJSONByte(data interface{}) []byte {
	jsonByte, err := json.Marshal(data)
	if err != nil {
		return nil
	}

	return jsonByte
}

// JSONToString conver JsonStruct to JsonString
func JSONToString(v interface{}) string {
	data, _ := json.MarshalIndent(v, "", " ")
	STR := string(data)
	STR = strings.ReplaceAll(STR, string(10), ``)
	return STR
}

// StringToJSON ...
func StringToJSON(jsStr string) map[string]interface{} {
	return ByteToJSON([]byte(jsStr))
}

// ByteToStruct ...
func ByteToStruct(jsByte []byte, structInfo interface{}) error {
	if errMsg := json.Unmarshal(jsByte, structInfo); errMsg != nil {
		fmt.Println(string(jsByte))
		return errMsg
	}
	return nil
}

// ByteToJSON ...
func ByteToJSON(jsByte []byte) map[string]interface{} {
	var data map[string]interface{}
	if errMsg := json.Unmarshal(jsByte, &data); errMsg != nil {
		panic(errMsg)
	}

	return data
}
