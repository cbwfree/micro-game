package dtype

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"reflect"
	"strings"
)

// ToJson 转换为JSON数据
func ToJson(data interface{}) []byte {
	buf, _ := json.Marshal(data)
	return buf
}

// ToDict 转换为字典数据(MAP)
func ToDict(data interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	_ = json.Unmarshal(ToJson(data), &res)
	return res
}

// ToStruct 转换结构体数据
func ToStruct(data interface{}, obj interface{}) error {
	if reflect.ValueOf(data).Kind() == reflect.String {
		return json.Unmarshal([]byte(data.(string)), obj)
	}
	return json.Unmarshal(ToJson(data), obj)
}

// StructToDict 结构体转字典 (不支持嵌套)
func StructToDict(data interface{}) map[string]interface{} {
	res := make(map[string]interface{}, 1)

	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		val := v.Field(i).Interface()
		key := strings.Split(v.Type().Field(i).Tag.Get("json"), ",")[0]
		if key == "-" {
			continue
		}
		if key == "" {
			key = v.Type().Field(i).Name
		}

		res[key] = val
	}

	return res
}

// 深度拷贝对象
func DeepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}
