package util

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

// InterfaceToString  转换interface到字符串
func InterfaceToString(val interface{}) string {
	if val == nil {
		return ""
	}
	switch ret := val.(type) {
	case string:
		return ret
	case int8, uint8, int16, uint16, int, uint, int64, uint64, float32, float64:
		return fmt.Sprintf("%v", ret)
	default:
		bytes, err := json.Marshal(val)
		if err != nil {
			fmt.Println(err)
			return ""
		}
		return string(bytes)
	}
}

// EncryptForMd5  md5 加密
func EncryptForMd5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
