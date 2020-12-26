package web

import (
	"github.com/cbwfree/micro-game/utils/dtype"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/steambap/captcha"
	"image/color"
)

type Context struct {
	ctx echo.Context
}

func (c *Context) Echo() echo.Context {
	return c.ctx
}

func (c *Context) Error(err error) error {
	return CtxResult(c.ctx, ParseError(err))
}

func (c *Context) JsonError(code int, msg ...string) error {
	return CtxError(c.ctx, code, msg...)
}

func (c *Context) JsonSuccess(data interface{}, msg ...string) error {
	return CtxSuccess(c.ctx, data, msg...)
}

func (c *Context) Bind(req interface{}) error {
	return c.ctx.Bind(req)
}

func (c *Context) BindValid(req interface{}) error {
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.ctx.Validate(req); err != nil {
		return err
	}
	return nil
}

func (c *Context) Session() *sessions.Session {
	s, _ := session.Get("SESSION", c.Echo())
	return s
}

func (c *Context) SessionDo(closure func(*sessions.Session) interface{}) interface{} {
	s := c.Session()
	res := closure(s)
	_ = s.Save(c.ctx.Request(), c.ctx.Response())
	return res
}

func (c *Context) SessId() string {
	res := c.SessionDo(func(ss *sessions.Session) interface{} {
		return ss.ID
	})
	return res.(string)
}

func (c *Context) SessOpts() *sessions.Options {
	res := c.SessionDo(func(ss *sessions.Session) interface{} {
		return ss.Options
	})
	return res.(*sessions.Options)
}

func (c *Context) SessOptsSet(opts *sessions.Options) {
	_ = c.SessionDo(func(ss *sessions.Session) interface{} {
		ss.Options = opts
		return nil
	})
}

func (c *Context) SessFlashAdd(val interface{}, key ...string) {
	_ = c.SessionDo(func(ss *sessions.Session) interface{} {
		ss.AddFlash(val, key...)
		return nil
	})
}

func (c *Context) SessFlash(key ...string) []interface{} {
	res := c.SessionDo(func(ss *sessions.Session) interface{} {
		return ss.Flashes(key...)
	})
	return res.([]interface{})
}

func (c *Context) SessGetValues() map[interface{}]interface{} {
	res := c.SessionDo(func(ss *sessions.Session) interface{} {
		return ss.Values
	})
	return res.(map[interface{}]interface{})
}

func (c *Context) SessSetValues(values map[interface{}]interface{}) {
	_ = c.SessionDo(func(ss *sessions.Session) interface{} {
		for k, v := range values {
			ss.Values[k] = v
		}
		return nil
	})
}

func (c *Context) SessHas(key interface{}) bool {
	_, b := c.SessGetValues()[key]
	return b
}

func (c *Context) SessGet(key interface{}) interface{} {
	return c.SessGetValues()[key]
}

func (c *Context) SessSet(key, val interface{}) {
	c.SessSetValues(map[interface{}]interface{}{
		key: val,
	})
}

func (c *Context) SessDel(key ...interface{}) {
	c.SessionDo(func(s *sessions.Session) interface{} {
		for _, k := range key {
			if _, ok := s.Values[k]; ok {
				delete(s.Values, k)
			}
		}
		return nil
	})
}

func (c *Context) SessClean() {
	c.SessionDo(func(s *sessions.Session) interface{} {
		s.Values = make(map[interface{}]interface{})
		return nil
	})
}

// 创建验证码
func (c *Context) CaptchaNew(key string, width, height int, setOpt ...captcha.SetOption) error {
	var opt captcha.SetOption
	if len(setOpt) > 0 && setOpt[0] != nil {
		opt = setOpt[0]
	} else {
		opt = func(opt *captcha.Options) {
			opt.BackgroundColor = color.White
			opt.CharPreset = "0123456789"
			opt.CurveNumber = 1
			opt.FontDPI = 80
		}
	}

	data, err := captcha.New(width, height, opt)
	if err != nil {
		return c.Error(err)
	}

	c.SessSet(key, data.Text)

	return data.WriteImage(c.ctx.Response().Writer)
}

// 验证验证码
func (c *Context) CaptchaCheck(key string, captcha string) bool {
	codeText := c.SessionDo(func(s *sessions.Session) interface{} {
		if c, ok := s.Values[key]; ok {
			return c
		}
		return ""
	})
	return dtype.ParseStr(codeText) == captcha
}

func NewContext(ctx echo.Context) *Context {
	c := &Context{
		ctx: ctx,
	}
	return c
}
