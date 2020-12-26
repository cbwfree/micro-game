package errors

import (
	"fmt"
	"github.com/cbwfree/micro-game/app"
	"github.com/cbwfree/micro-game/utils/dtype"
	"github.com/micro/go-micro/v2/errors"
	"net/http"
)

// New 创建系统错误
func New(code int32, a ...interface{}) error {
	var detail string
	switch len(a) {
	case 0:
		detail = ""
	case 1:
		detail = dtype.ParseStr(a[0])
	default:
		detail = fmt.Sprintf(dtype.ParseStr(a[0]), a[1:]...)
	}
	return &errors.Error{
		Id:     app.NameId(),
		Code:   code,
		Detail: detail,
		Status: StatusText(code),
	}
}

// Parse 普通错误转系统错误
func Parse(err error) *errors.Error {
	merr, ok := err.(*errors.Error)
	if !ok {
		merr = errors.Parse(err.Error())
	}

	if merr.Code == 0 {
		merr.Code = http.StatusInternalServerError
	}
	if merr.Id == "" {
		merr.Id = app.NameId()
	}

	if merr.Status == "" {
		merr.Status = StatusText(merr.Code)
	}

	return merr
}

func Wrap(err error, code int32, a ...interface{}) error {
	merr := Parse(err)
	merr.Code = code
	merr.Status = StatusText(code)
	if len(a) > 0 {
		switch len(a) {
		case 1:
			merr.Detail = dtype.ParseStr(a[0])
		default:
			merr.Detail = fmt.Sprintf(dtype.ParseStr(a[0]), a[1:]...)
		}
	}
	return merr
}

// IsCode 检查错误码
func IsCode(err error, code int32) bool {
	return Parse(err).Code == code
}

// Is 比较两个错误
func Is(a, b error) bool {
	ae, be := Parse(a), Parse(b)
	if ae.Code == be.Code {
		return true
	}
	return false
}
