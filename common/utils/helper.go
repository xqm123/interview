package utils

import "encoding/json"

// JsonString json格式化
func JsonString(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(b)
}
