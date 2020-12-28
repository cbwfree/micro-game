package errors

// 通用错误
func Universal(args ...interface{}) *Error {
	return New(CodeUniversal, args...)
}

// 未知错误
func Unknown(args ...interface{}) *Error {
	return New(CodeUnknown, args...)
}

// 无效
func Invalid(args ...interface{}) *Error {
	return New(CodeInvalid, args...)
}

// 已存在
func Exists(args ...interface{}) *Error {
	return New(CodeExists, args...)
}

// 已开始
func Started(args ...interface{}) *Error {
	return New(CodeStarted, args...)
}

// 已结束
func Finished(args ...interface{}) *Error {
	return New(CodeFinished, args...)
}

// 其它错误
func Other(args ...interface{}) *Error {
	return New(CodeOther, args...)
}

// 未授权
func Unauthorized(args ...interface{}) *Error {
	return New(CodeUnauthorized, args...)
}

// 拒绝访问
func Forbidden(args ...interface{}) *Error {
	return New(CodeForbidden, args...)
}

// 不存在
func NotFound(args ...interface{}) *Error {
	return New(CodeNotFound, args...)
}

// 请求超时
func Timeout(args ...interface{}) *Error {
	return New(CodeTimeout, args...)
}

// 服务器错误
func Server(args ...interface{}) *Error {
	return New(CodeServerError, args...)
}

// 不可用
func Unavailable(args ...interface{}) *Error {
	return New(CodeServiceUnavailable, args...)
}
