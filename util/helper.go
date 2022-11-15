package util

import (
	"encoding/json"
)

// JsonMarshalIndent struct 转字符串
func JsonMarshalIndent(v interface{}) string {
	indent, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(indent)
}
