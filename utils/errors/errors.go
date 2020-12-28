package errors

import (
	"encoding/json"
	"fmt"
	"github.com/cbwfree/micro-game/utils/dtype"
	"github.com/micro/go-micro/v2/errors"
	"io"
	"os"
	"strconv"
)

var isStack bool

func init() {
	isStack, _ = strconv.ParseBool(os.Getenv("GAME_ERROR_STACK"))
}

func IsStack() bool {
	return isStack
}

func message(args ...interface{}) string {
	var msg string
	switch len(args) {
	case 0:
		msg = ""
	case 1:
		msg = dtype.ParseStr(args[0])
	default:
		msg = fmt.Sprintf(dtype.ParseStr(args[0]), args[1:]...)
	}
	return msg
}

// New returns an error with the supplied message.
// New also records the stack trace at the point it was called.
func New(code int32, a ...interface{}) *Error {
	e := &Error{
		Code:   code,
		Detail: message(a...),
		Status: StatusText(code),
	}
	if isStack {
		e.stack = callers()
	}
	return e
}

type Error struct {
	Code   int32  `json:"code"`
	Status string `json:"status"`
	Detail string `json:"detail"`
	*stack `json:"-"`
}

func (e *Error) Error() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func (e *Error) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			io.WriteString(s, e.Error())
			if e.stack != nil {
				e.stack.Format(s, verb)
			}
			return
		}
		fallthrough
	case 's':
		io.WriteString(s, e.Error())
	case 'q':
		fmt.Fprintf(s, "%q", e.Error())
	}
}

// WithStack annotates err with a stack trace at the point WithStack was called.
// If err is nil, WithStack returns nil.
func Wrap(err error, args ...interface{}) error {
	if err == nil {
		return nil
	}
	e := &wrapError{
		detail: message(args...),
		error:  err,
	}
	if isStack {
		e.stack = callers()
	}
	return e
}

type wrapError struct {
	detail string
	error
	*stack
}

func (w *wrapError) Error() string { return fmt.Sprintf("%s : %s", w.detail, w.error) }

func (w *wrapError) Cause() error { return w.error }

func (w *wrapError) Unwrap() error { return w.error }

func (w *wrapError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%+v", w.Cause())
			if w.stack != nil {
				w.stack.Format(s, verb)
			}
			return
		}
		fallthrough
	case 's':
		io.WriteString(s, w.Error())
	case 'q':
		fmt.Fprintf(s, "%q", w.Error())
	}
}

type causer interface {
	Cause() error
}

// Cause returns the underlying cause of the error, if possible.
// An error value has a cause if it implements the following
// interface:
//
//     type causer interface {
//            Cause() error
//     }
//
// If the error does not implement Cause, the original error will
// be returned. If the error is nil, nil will be returned without further
// investigation.
func Cause(err error) error {
	for err != nil {
		cause, ok := err.(causer)
		if !ok {
			break
		}
		err = cause.Cause()
	}
	return err
}

func Parse(err error) *Error {
	if err == nil {
		return nil
	}

	switch ee := err.(type) {
	case causer:
		return Parse(ee.Cause())
	case *Error:
		return ee
	case *errors.Error:
		e := &Error{
			Code:   ee.Code,
			Status: ee.Status,
			Detail: ee.Detail,
		}
		if isStack {
			e.stack = callers()
		}
		return e
	default:
		es := ee.Error()
		e := new(Error)
		if er := json.Unmarshal([]byte(es), e); er != nil {
			e = Unknown(es)
		}
		return e
	}
}

func MicroError(id string, err error) *errors.Error {
	if err == nil {
		return nil
	}

	switch ee := Cause(err).(type) {
	case *errors.Error:
		return ee
	case *Error:
		return &errors.Error{
			Id:     id,
			Code:   ee.Code,
			Detail: ee.Detail,
			Status: ee.Status,
		}
	default:
		e := new(errors.Error)
		if er := json.Unmarshal([]byte(ee.Error()), e); er != nil {
			e.Detail = ee.Error()
		}
		if e.Id == "" {
			e.Id = id
		}
		if e.Code == 0 {
			e.Code = CodeUnknown
			e.Status = StatusText(CodeUnknown)
		}
		return e
	}
}

// IsCode 检查错误码
func IsCode(err error, code int32) bool {
	return Parse(Cause(err)).Code == code
}

// Is 比较两个错误
func Is(a, b error) bool {
	if Parse(Cause(a)).Code == Parse(Cause(b)).Code {
		return true
	}
	return false
}
