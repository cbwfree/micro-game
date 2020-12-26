package dtype

import (
	"reflect"
)

var (
	RefTypeStr     = reflect.TypeOf("")
	RefTypeInt     = reflect.TypeOf(0)
	RefTypeInt32   = reflect.TypeOf(int32(0))
	RefTypeInt64   = reflect.TypeOf(int64(0))
	RefTypeFloat32 = reflect.TypeOf(float32(0))
	RefTypeFloat64 = reflect.TypeOf(float64(0))
	RefTypeBool    = reflect.TypeOf(false)
)

// 反射创建结构体 (以 .Addr.Interface() 获取指针数据)
func Ptr(v reflect.Type) reflect.Value {
	return reflect.New(v)
}

func Elem(v reflect.Type) reflect.Value {
	return reflect.New(v).Elem()
}

func StrElem() reflect.Value {
	return reflect.New(RefTypeStr).Elem()
}

func IntElem() reflect.Value {
	return reflect.New(RefTypeInt).Elem()
}

func Int32Elem() reflect.Value {
	return reflect.New(RefTypeInt32).Elem()
}

func Int64Elem() reflect.Value {
	return reflect.New(RefTypeInt64).Elem()
}

func Float32Elem() reflect.Value {
	return reflect.New(RefTypeFloat32).Elem()
}

func Float64Elem() reflect.Value {
	return reflect.New(RefTypeFloat64).Elem()
}

func BoolElem() reflect.Value {
	return reflect.New(RefTypeBool).Elem()
}

// 反射创建切片 (以 .Addr.Interface() 获取指针数据)
func SliceElem(v reflect.Type) reflect.Value {
	var t reflect.Type
	if v.Kind() == reflect.Struct {
		t = Ptr(v).Type()
	} else {
		t = v
	}
	return reflect.New(reflect.SliceOf(t)).Elem()
}

// 反射创建切片 (以 .Addr.Interface() 获取指针数据)
func SliceTwoElem(v reflect.Type) reflect.Value {
	var t reflect.Type
	if v.Kind() == reflect.Struct {
		t = Ptr(v).Type()
	} else {
		t = v
	}
	return reflect.New(reflect.SliceOf(reflect.SliceOf(t))).Elem()
}

func StrSliceElem() reflect.Value {
	return SliceElem(RefTypeStr)
}

func IntSliceElem() reflect.Value {
	return SliceElem(RefTypeInt)
}

func Int32SliceElem() reflect.Value {
	return SliceElem(RefTypeInt32)
}

func Int64SliceElem() reflect.Value {
	return SliceElem(RefTypeInt64)
}

func Float32SliceElem() reflect.Value {
	return SliceElem(RefTypeFloat32)
}

func Float64SliceElem() reflect.Value {
	return SliceElem(RefTypeFloat64)
}

func BoolSliceElem() reflect.Value {
	return SliceElem(RefTypeBool)
}

// 反射创建MAP (以 .Addr.Interface() 获取指针数据)
func MapElem(k, v reflect.Type) reflect.Value {
	var t reflect.Type
	if v.Kind() == reflect.Struct {
		t = Ptr(v).Type()
	} else {
		t = v
	}
	return reflect.New(reflect.MapOf(k, t)).Elem()
}

// Str Map
func StrMapElem(v reflect.Type) reflect.Value {
	return MapElem(RefTypeStr, v)
}

// Int Map
func IntMapElem(v reflect.Type) reflect.Value {
	return MapElem(RefTypeInt, v)
}

func Int32MapElem(v reflect.Type) reflect.Value {
	return MapElem(RefTypeInt32, v)
}

func Int64MapElem(v reflect.Type) reflect.Value {
	return MapElem(RefTypeInt64, v)
}
