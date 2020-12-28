package color

import "fmt"

type Style []Color

func New(colors ...Color) Style {
	return colors
}

// IsEmpty style
func (s Style) IsEmpty() bool {
	return len(s) == 0
}

// Code convert to code string. returns like "32;45;3"
func (s Style) Code() string {
	return colors2code(s...)
}

// Save to styles map
func (s Style) Save(name string) {
	AddStyle(name, s)
}

// Text render a text message
func (s Style) Text(format string, a ...interface{}) string {
	return RenderStr(s.Code(), fmt.Sprintf(format, a...))
}

// Render messages by color setting
func (s Style) Render(a ...interface{}) string {
	return RenderCode(s.Code(), a...)
}

/*************************************************************
 * internal styles
 *************************************************************/

// Styles internal defined styles, like bootstrap styles.
// Usage:
// 	color.Styles["info"].Println("message")
var Styles = map[string]Style{
	"info":  {OpReset, FgGreen},
	"note":  {OpBold, FgLightCyan},
	"light": {FgLightWhite, BgRed},
	"error": {OpBold, FgRed},

	"danger":  {OpBold, FgRed},
	"notice":  {OpBold, FgCyan},
	"success": {OpBold, FgGreen},
	"comment": {OpReset, FgMagenta},
	"primary": {OpReset, FgBlue},
	"warning": {OpBold, FgYellow},

	"question":  {OpReset, FgMagenta},
	"secondary": {FgDarkGray},
}

// some style name alias
var styleAliases = map[string]string{
	"err":  "error",
	"suc":  "success",
	"warn": "warning",
}

// AddStyle add a style
func AddStyle(name string, s Style) {
	Styles[name] = s
}

// GetStyle get defined style by name
func GetStyle(name string) Style {
	if s, ok := Styles[name]; ok {
		return s
	}

	if realName, ok := styleAliases[name]; ok {
		return Styles[realName]
	}

	// empty style
	return New()
}
