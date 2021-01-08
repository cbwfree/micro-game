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

func ParseStrSlice(val string) ([]string, error) {
	var slice []string
	if err := ParseMetaValue(val, &slice); err != nil {
		return nil, err
	}
	return slice, nil
}

func ParseIntSlice(val string) ([]int, error) {
	var slice []int
	if err := ParseMetaValue(val, &slice); err != nil {
		return nil, err
	}
	return slice, nil
}

func ParseInt32Slice(val string) ([]int32, error) {
	var slice []int32
	if err := ParseMetaValue(val, &slice); err != nil {
		return nil, err
	}
	return slice, nil
}

func ParseInt64Slice(val string) ([]int64, error) {
	var slice []int64
	if err := ParseMetaValue(val, &slice); err != nil {
		return nil, err
	}
	return slice, nil
}

func ParseFloat32Slice(val string) ([]float32, error) {
	var slice []float32
	if err := ParseMetaValue(val, &slice); err != nil {
		return nil, err
	}
	return slice, nil
}

func ParseFloat64Slice(val string) ([]float64, error) {
	var slice []float64
	if err := ParseMetaValue(val, &slice); err != nil {
		return nil, err
	}
	return slice, nil
}
