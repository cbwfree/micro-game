package tool

import (
	"math/rand"
	"reflect"
	"strings"
	"time"
)

type reducetype func(interface{}) interface{}
type filtertype func(interface{}) bool

// InSlice checks given interface in interface slice.
func InSlice(v interface{}, slice interface{}) bool {
	switch slice.(type) {
	case nil:
		return false
	case []interface{}:
		for _, vv := range slice.([]interface{}) {
			if vv == v {
				return true
			}
		}
	case []string:
		if _v, ok := v.(string); ok {
			return InStrSlice(_v, slice.([]string))
		}
		return false
	case []int:
		if _v, ok := v.(int); ok {
			return InIntSlice(_v, slice.([]int))
		}
		return false
	case []int64:
		if _v, ok := v.(int64); ok {
			return InInt64Slice(_v, slice.([]int64))
		}
		return false
	case []float32:
		if _v, ok := v.(float32); ok {
			return InFloat32Slice(_v, slice.([]float32))
		}
		return false
	case []float64:
		if _v, ok := v.(float64); ok {
			return InFloat64Slice(_v, slice.([]float64))
		}
		return false
	}

	s := reflect.ValueOf(slice)
	kind := s.Kind()
	if kind != reflect.Slice && kind != reflect.Array {
		return false
	}

	slen := s.Len()
	for i := 0; i < slen; i++ {
		if reflect.DeepEqual(v, s.Index(i).Interface()) {
			return true
		}
	}
	return false
}

// InStrSlice checks given string in string slice or not.
func InStrSlice(v string, sl []string) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

// InIntSlice checks given int in int slice or not.
func InIntSlice(v int, sl []int) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

// InInt32Slice checks given int in int slice or not.
func InInt32Slice(v int32, sl []int32) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

// InInt64Slice checks given int64 in int64 slice or not.
func InInt64Slice(v int64, sl []int64) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

// InInt64Slice checks given int64 in int64 slice or not.
func InFloat32Slice(v float32, sl []float32) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

// InInt64Slice checks given int64 in int64 slice or not.
func InFloat64Slice(v float64, sl []float64) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

// SliceRandList generate an int slice from min to max.
func SliceRandList(min, max int) []int {
	if max < min {
		min, max = max, min
	}
	length := max - min + 1
	t0 := time.Now()
	rand.Seed(int64(t0.Nanosecond()))
	list := rand.Perm(length)
	for index := range list {
		list[index] += min
	}
	return list
}

// SliceMerge merges interface slices to one slice.
func SliceMerge(slice1, slice2 []interface{}) (c []interface{}) {
	c = append(slice1, slice2...)
	return
}

// SliceReduce generates a new slice after parsing every value by reduce function
func SliceReduce(slice []interface{}, a reducetype) (dslice []interface{}) {
	for _, v := range slice {
		dslice = append(dslice, a(v))
	}
	return
}

// SliceRand returns random one from slice.
func SliceRand(a []interface{}) (b interface{}) {
	randnum := rand.Intn(len(a))
	b = a[randnum]
	return
}

// SliceRand returns random one from slice.
func SliceRandStr(a []string) (b string) {
	randnum := rand.Intn(len(a))
	b = a[randnum]
	return
}

// SliceSum sums all values in int64 slice.
func SliceSum(intslice []int64) (sum int64) {
	for _, v := range intslice {
		sum += v
	}
	return
}

// SliceFilter generates a new slice after filter function.
func SliceFilter(slice []interface{}, a filtertype) (ftslice []interface{}) {
	for _, v := range slice {
		if a(v) {
			ftslice = append(ftslice, v)
		}
	}
	return
}

// SliceDiff returns diff slice of slice1 - slice2.
func SliceDiff(slice1, slice2 []interface{}) (diffslice []interface{}) {
	for _, v := range slice1 {
		if !InSlice(v, slice2) {
			diffslice = append(diffslice, v)
		}
	}
	return
}

// SliceIntersect returns slice that are present in all the slice1 and slice2.
func SliceIntersect(slice1, slice2 []interface{}) (diffslice []interface{}) {
	for _, v := range slice1 {
		if InSlice(v, slice2) {
			diffslice = append(diffslice, v)
		}
	}
	return
}

// SliceChunk separates one slice to some sized slice.
func SliceChunk(slice []interface{}, size int) (chunkslice [][]interface{}) {
	if size >= len(slice) {
		chunkslice = append(chunkslice, slice)
		return
	}
	end := size
	for i := 0; i <= (len(slice) - size); i += size {
		chunkslice = append(chunkslice, slice[i:end])
		end += size
	}
	return
}

// SliceRange generates a new slice from begin to end with step duration of int64 number.
func SliceRange(start, end, step int64) (intslice []int64) {
	for i := start; i <= end; i += step {
		intslice = append(intslice, i)
	}
	return
}

// SlicePad prepends size number of val into slice.
func SlicePad(slice []interface{}, size int, val interface{}) []interface{} {
	if size <= len(slice) {
		return slice
	}
	for i := 0; i < (size - len(slice)); i++ {
		slice = append(slice, val)
	}
	return slice
}

// SliceUnique cleans repeated values in slice.
func SliceUnique(slice []interface{}) (uniqueslice []interface{}) {
	for _, v := range slice {
		if !InSlice(v, uniqueslice) {
			uniqueslice = append(uniqueslice, v)
		}
	}
	return
}

// Int32SliceUnique cleans repeated values in slice.
func Int32SliceUnique(slice []int32) (uniqueslice []int32) {
	for _, v := range slice {
		if !InInt32Slice(v, uniqueslice) {
			uniqueslice = append(uniqueslice, v)
		}
	}
	return
}

// Int64SliceUnique cleans repeated values in slice.
func Int64SliceUnique(slice []int64) (uniqueslice []int64) {
	for _, v := range slice {
		if !InInt64Slice(v, uniqueslice) {
			uniqueslice = append(uniqueslice, v)
		}
	}
	return
}

// StrSliceUnique cleans repeated values in slice.
func StrSliceUnique(slice []string) (uniqueslice []string) {
	for _, v := range slice {
		if !InStrSlice(v, uniqueslice) {
			uniqueslice = append(uniqueslice, v)
		}
	}
	return
}

// SliceShuffle shuffles a slice.
func SliceShuffle(slice []interface{}) []interface{} {
	for i := 0; i < len(slice); i++ {
		a := rand.Intn(len(slice))
		b := rand.Intn(len(slice))
		slice[a], slice[b] = slice[b], slice[a]
	}
	return slice
}

// SliceShuffleInt shuffles a int slice.
func SliceShuffleInt(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		a := rand.Intn(len(slice))
		b := rand.Intn(len(slice))
		slice[a], slice[b] = slice[b], slice[a]
	}
	return slice
}

// StrSliceIface 字符串切片转 []interface{}
func StrSliceIface(slice []string) (result []interface{}) {
	result = make([]interface{}, len(slice))
	for i, v := range slice {
		result[i] = v
	}
	return
}

// IntSliceIface 数字切片转 []interface{}
func IntSliceIface(slice []int) (result []interface{}) {
	result = make([]interface{}, len(slice))
	for i, v := range slice {
		result[i] = v
	}
	return
}

// Int64SliceIface 数字切片转 []interface{}
func Int64SliceIface(slice []int64) (result []interface{}) {
	result = make([]interface{}, len(slice))
	for i, v := range slice {
		result[i] = v
	}
	return
}

// SplitStrSlice 获取字符串切片
func SplitStrSlice(str string, sep string) []interface{} {
	slice := StrSliceIface(strings.Split(str, sep))
	slice = SliceFilter(slice[:], func(v interface{}) bool { return v != "" })
	slice = SliceUnique(slice)
	return slice
}

// 切片乱序（在原对象的基础上）
func RandSlice(slice interface{}) {
	rv := reflect.ValueOf(slice)
	if rv.Type().Kind() != reflect.Slice {
		return
	}

	length := rv.Len()
	if length < 2 {
		return
	}

	swap := reflect.Swapper(slice)
	rand.Seed(time.Now().Unix())
	for i := length - 1; i >= 0; i-- {
		j := rand.Intn(length)
		swap(i, j)
	}
	return
}
