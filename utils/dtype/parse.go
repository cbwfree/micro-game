// 数据转换
package dtype

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// ParseStr 转为字符串
func ParseStr(v interface{}) string {
	if v == nil {
		return ""
	}

	switch result := v.(type) {
	case string:
		return result
	case []byte:
		return string(result)
	case int:
		return strconv.Itoa(result)
	case int32:
		return strconv.Itoa(int(result))
	case int64:
		return strconv.FormatInt(result, 10)
	default:
		return fmt.Sprint(result)
	}
}

// ParseInt 转为int
func ParseInt(v interface{}) int {
	switch result := v.(type) {
	case json.Number:
		i, _ := result.Int64()
		return int(i)
	case int:
		return result
	case int32:
		return int(result)
	case int64:
		return int(result)
	case float32:
		return int(result)
	case float64:
		return int(result)
	default:
		if d := ParseStr(v); d != "" {
			value, _ := strconv.Atoi(d)
			return value
		}
	}
	return 0
}

// ParseInt32 convert interface to int32.
func ParseInt32(v interface{}) int32 {
	return int32(ParseInt(v))
}

// ParseInt64 convert interface to int64.
func ParseInt64(v interface{}) int64 {
	switch result := v.(type) {
	case json.Number:
		i, _ := result.Int64()
		return i
	case int:
		return int64(result)
	case int32:
		return int64(result)
	case int64:
		return result
	case float32:
		return int64(result)
	case float64:
		return int64(result)
	default:
		if d := ParseStr(v); d != "" {
			value, _ := strconv.ParseInt(d, 10, 64)
			return value
		}
	}
	return 0
}

// ParseUint convert interface to int.
func ParseUint(v interface{}) uint {
	switch result := v.(type) {
	case json.Number:
		i, _ := result.Int64()
		return uint(i)
	case int:
		return uint(result)
	case int32:
		return uint(result)
	case int64:
		return uint(result)
	case uint:
		return result
	case uint32:
		return uint(result)
	case uint64:
		return uint(result)
	case float32:
		return uint(result)
	case float64:
		return uint(result)
	default:
		if d := ParseStr(v); d != "" {
			value, _ := strconv.ParseUint(d, 10, 32)
			return uint(value)
		}
	}
	return 0
}

// ParseUint32 convert interface to int32.
func ParseUint32(v interface{}) uint32 {
	return uint32(ParseUint(v))
}

// ParseUint64 convert interface to int64.
func ParseUint64(v interface{}) uint64 {
	switch result := v.(type) {
	case json.Number:
		i, _ := result.Int64()
		return uint64(i)
	case int:
		return uint64(result)
	case int32:
		return uint64(result)
	case int64:
		return uint64(result)
	case uint:
		return uint64(result)
	case uint32:
		return uint64(result)
	case uint64:
		return result
	case float32:
		return uint64(result)
	case float64:
		return uint64(result)
	default:
		if d := ParseStr(v); d != "" {
			value, _ := strconv.ParseUint(d, 10, 64)
			return value
		}
	}
	return 0
}

// ParseFloat32 convert interface to float32.
func ParseFloat32(v interface{}) float32 {
	switch result := v.(type) {
	case json.Number:
		i, _ := result.Float64()
		return float32(i)
	case float32:
		return result
	case float64:
		return float32(result)
	case int:
		return float32(result)
	case int32:
		return float32(result)
	case int64:
		return float32(result)
	default:
		if d := ParseStr(v); d != "" {
			value, _ := strconv.ParseFloat(d, 32)
			return float32(value)
		}
	}
	return 0
}

// ParseFloat64 convert interface to float64.
func ParseFloat64(v interface{}) float64 {
	switch result := v.(type) {
	case json.Number:
		i, _ := result.Float64()
		return i
	case float64:
		return result
	case float32:
		return float64(result)
	case int:
		return float64(result)
	case int32:
		return float64(result)
	case int64:
		return float64(result)
	default:
		if d := ParseStr(v); d != "" {
			value, _ := strconv.ParseFloat(d, 64)
			return value
		}
	}
	return 0
}

// ParseBool convert interface to bool.
func ParseBool(v interface{}) bool {
	switch result := v.(type) {
	case bool:
		return result
	default:
		if d := ParseStr(v); d != "" {
			value, _ := strconv.ParseBool(d)
			return value
		}
	}
	return false
}

// ParseStrSlice convert []interface to []string
func ParseStrSlice(v []interface{}) []string {
	var result []string
	for _, t := range v {
		result = append(result, ParseStr(t))
	}
	return result
}

func ParseSlice(v interface{}) []interface{} {
	switch res := v.(type) {
	case []interface{}:
		return res
	}
	return nil
}
