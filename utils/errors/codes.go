// 系统错误
package errors

import "net/http"

// App System Error Code (1 ~ 99)
const (
	CodeOk        int32 = 0
	CodeUniversal       = 1
	CodeUnknown         = 2
	CodeInvalid         = 3
	CodeExists          = 4

	CodeStarted  = 90
	CodeFinished = 91
	CodeOther    = 99
)

// HTTP Status Code (100 ~ 600)
const (
	_                int32 = 0
	CodeUnauthorized       = 401
	CodeForbidden          = 403
	CodeNotFound           = 404
	CodeTimeout            = 408

	CodeServerError        = 500
	CodeServiceUnavailable = 503
)

func init() {
	AddStatusText(defaultLang, CodeUniversal, "universal")
	AddStatusText(defaultLang, CodeUnknown, "unknown")
	AddStatusText(defaultLang, CodeInvalid, "invalid")
	AddStatusText(defaultLang, CodeExists, "exists")
	AddStatusText(defaultLang, CodeStarted, "started")
	AddStatusText(defaultLang, CodeFinished, "finished")
	AddStatusText(defaultLang, CodeOther, "other")

	AddStatusText(defaultLang, CodeUnauthorized, http.StatusText(CodeUnauthorized))
	AddStatusText(defaultLang, CodeForbidden, http.StatusText(CodeForbidden))
	AddStatusText(defaultLang, CodeNotFound, http.StatusText(CodeNotFound))
	AddStatusText(defaultLang, CodeTimeout, http.StatusText(CodeTimeout))
	AddStatusText(defaultLang, CodeServerError, http.StatusText(CodeServerError))
	AddStatusText(defaultLang, CodeServiceUnavailable, http.StatusText(CodeServiceUnavailable))
}
