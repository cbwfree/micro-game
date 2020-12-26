package meta

import (
	"encoding/json"
	"strconv"
)

// 转换Meta值
func ToMetaValue(value interface{}) string {
	switch val := value.(type) {
	case string:
		return val
	case []byte:
		return string(val)
	case int:
		return strconv.Itoa(val)
	case int32:
		return strconv.Itoa(int(val))
	case int64:
		return strconv.FormatInt(val, 10)
	case bool:
		return strconv.FormatBool(val)
	default:
		b, _ := json.Marshal(val)
		return string(b)
	}
}

// 解析Meta值
func ParseMetaValue(val string, res interface{}) error {
	return json.Unmarshal([]byte(val), res)
}
