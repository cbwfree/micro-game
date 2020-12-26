package web

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	merr "github.com/micro/go-micro/v2/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Time int64       `json:"time"`
}

func NewResult(code int, msg string, data interface{}) *Result {
	return &Result{
		Code: code,
		Msg:  msg,
		Data: data,
		Time: time.Now().Unix(),
	}
}

func CtxError(ctx echo.Context, code int, msg ...string) error {
	var mg string
	if len(msg) > 0 {
		mg = msg[0]
	} else {
		mg = http.StatusText(code)
		if mg == "" {
			mg = "unknown error"
		}
	}
	return CtxResult(ctx, NewResult(code, mg, ""))
}

func CtxSuccess(ctx echo.Context, data interface{}, msg ...string) error {
	if len(msg) > 0 {
		return CtxResult(ctx, NewResult(0, msg[0], data))
	}
	return CtxResult(ctx, NewResult(0, "success", data))
}

func CtxResult(ctx echo.Context, res *Result) error {
	return ctx.JSON(http.StatusOK, res)
}

func ParseError(err error) *Result {
	switch err {
	case mongo.ErrNilDocument:
		err = echo.NewHTTPError(http.StatusNotFound, "mongodb nil document")
	case mongo.ErrNoDocuments:
		err = echo.NewHTTPError(http.StatusNotFound, "mongodb no documents")
	case redis.Nil:
		err = echo.NewHTTPError(http.StatusNotFound, "redis nil")
	}

	var code int
	var msg string
	if he, ok := err.(*echo.HTTPError); ok { // HTTP 错误
		code = he.Code
		if he.Internal != nil {
			msg = fmt.Sprintf("%v, %v", err, he.Internal)
		} else {
			if m, ok := he.Message.(string); ok {
				msg = m
			} else {
				msg = he.Error()
			}
		}
	} else if ee := merr.Parse(err.Error()); ee.Code > 0 { // Micro RPC 返回的错误
		code = int(ee.Code)
		msg = ee.Detail
	} else {
		code = http.StatusInternalServerError
		msg = err.Error()
	}

	return NewResult(code, msg, nil)
}
