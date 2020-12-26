package tool

import (
	"bytes"
	"fmt"
	"strconv"
)

// GetCacheName 生成缓存名称
func GetCacheName(prefix string, tags ...interface{}) string {
	var buf = new(bytes.Buffer)
	buf.WriteString(prefix)
	for _, tag := range tags {
		buf.WriteString(":")
		switch tag.(type) {
		case string:
			buf.WriteString(tag.(string))
		case int:
			buf.WriteString(strconv.FormatInt(int64(tag.(int)), 10))
		case int32:
			buf.WriteString(strconv.FormatInt(int64(tag.(int32)), 10))
		case int64:
			buf.WriteString(strconv.FormatInt(tag.(int64), 10))
		case uint:
			buf.WriteString(strconv.FormatUint(uint64(tag.(uint)), 10))
		case uint32:
			buf.WriteString(strconv.FormatUint(uint64(tag.(uint32)), 10))
		case uint64:
			buf.WriteString(strconv.FormatUint(tag.(uint64), 10))
		default:
			buf.WriteString(fmt.Sprint(tag))
		}
	}
	return buf.String()
}
