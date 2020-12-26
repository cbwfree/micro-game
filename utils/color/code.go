package color

import (
	"fmt"
)

// Color Color16, 16 color value type
// 3(2^3=8) OR 4(2^4=16) bite color.
type Color uint8

/*************************************************************
 * Basic 16 color definition
 *************************************************************/

// Foreground colors. basic foreground colors 30 - 37
const (
	FgBlack Color = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta // 品红
	FgCyan    // 青色
	FgWhite
	// FgDefault revert default FG
	FgDefault Color = 39
)

// Extra foreground color 90 - 97(非标准)
const (
	FgDarkGray Color = iota + 90 // 亮黑（灰）
	FgLightRed
	FgLightGreen
	FgLightYellow
	FgLightBlue
	FgLightMagenta
	FgLightCyan
	FgLightWhite
	// FgGray is alias of FgDarkGray
	FgGray Color = 90 // 亮黑（灰）
)

// Background colors. basic background colors 40 - 47
const (
	BgBlack Color = iota + 40
	BgRed
	BgGreen
	BgYellow // BgBrown like yellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
	// BgDefault revert default BG
	BgDefault Color = 49
)

// Extra background color 100 - 107(非标准)
const (
	BgDarkGray Color = iota + 99
	BgLightRed
	BgLightGreen
	BgLightYellow
	BgLightBlue
	BgLightMagenta
	BgLightCyan
	BgLightWhite
	// BgGray is alias of BgDarkGray
	BgGray Color = 100
)

// Option settings
const (
	OpReset         Color = iota // 0 重置所有设置
	OpBold                       // 1 加粗
	OpFuzzy                      // 2 模糊(不是所有的终端仿真器都支持)
	OpItalic                     // 3 斜体(不是所有的终端仿真器都支持)
	OpUnderscore                 // 4 下划线
	OpBlink                      // 5 闪烁
	OpFastBlink                  // 5 快速闪烁(未广泛支持)
	OpReverse                    // 7 颠倒的 交换背景色与前景色
	OpConcealed                  // 8 隐匿的
	OpStrikethrough              // 9 删除的，删除线(未广泛支持)
)

// There are basic and light foreground color aliases
const (
	Red     = FgRed
	Cyan    = FgCyan
	Gray    = FgDarkGray // is light Black
	Blue    = FgBlue
	Black   = FgBlack
	Green   = FgGreen
	White   = FgWhite
	Yellow  = FgYellow
	Magenta = FgMagenta

	// special

	Bold   = OpBold
	Normal = FgDefault

	// extra light

	LightRed     = FgLightRed
	LightCyan    = FgLightCyan
	LightBlue    = FgLightBlue
	LightGreen   = FgLightGreen
	LightWhite   = FgLightWhite
	LightYellow  = FgLightYellow
	LightMagenta = FgLightMagenta
)

// IsValid 颜色值是否有效
func (c Color) IsValid() bool {
	return c < 107
}

// Light 更改为亮色
func (c Color) Light() Color {
	val := int(c)
	if val >= 30 && val <= 47 {
		return Color(uint8(c) + 60)
	}
	return c
}

// Darken 更改为暗色
func (c Color) Darken() Color {
	val := int(c)
	if val >= 90 && val <= 107 {
		return Color(uint8(c) - 60)
	}

	// don't change
	return c
}

// Code 转换颜色编码为字符串
func (c Color) Code() string {
	return fmt.Sprintf("%d", c)
}

// Text render a text message
func (c Color) Text(message string) string {
	return RenderStr(c.Code(), message)
}

// Render messages by color setting
func (c Color) Render(a ...interface{}) string {
	return RenderCode(c.Code(), a...)
}

/*************************************************************
 * basic color maps
 *************************************************************/

// FgColors foreground colors map
var FgColors = map[string]Color{
	"black":   FgBlack,
	"red":     FgRed,
	"green":   FgGreen,
	"yellow":  FgYellow,
	"blue":    FgBlue,
	"magenta": FgMagenta,
	"cyan":    FgCyan,
	"white":   FgWhite,
	"default": FgDefault,
}

// BgColors background colors map
var BgColors = map[string]Color{
	"black":   BgBlack,
	"red":     BgRed,
	"green":   BgGreen,
	"yellow":  BgYellow,
	"blue":    BgBlue,
	"magenta": BgMagenta,
	"cyan":    BgCyan,
	"white":   BgWhite,
	"default": BgDefault,
}

// ExFgColors extra foreground colors map
var ExFgColors = map[string]Color{
	"darkGray":     FgDarkGray,
	"lightRed":     FgLightRed,
	"lightGreen":   FgLightGreen,
	"lightYellow":  FgLightYellow,
	"lightBlue":    FgLightBlue,
	"lightMagenta": FgLightMagenta,
	"lightCyan":    FgLightCyan,
	"lightWhite":   FgLightWhite,
}

// ExBgColors extra background colors map
var ExBgColors = map[string]Color{
	"darkGray":     BgDarkGray,
	"lightRed":     BgLightRed,
	"lightGreen":   BgLightGreen,
	"lightYellow":  BgLightYellow,
	"lightBlue":    BgLightBlue,
	"lightMagenta": BgLightMagenta,
	"lightCyan":    BgLightCyan,
	"lightWhite":   BgLightWhite,
}

// Options color options map
var Options = map[string]Color{
	"reset":      OpReset,
	"bold":       OpBold,
	"fuzzy":      OpFuzzy,
	"italic":     OpItalic,
	"underscore": OpUnderscore,
	"blink":      OpBlink,
	"reverse":    OpReverse,
	"concealed":  OpConcealed,
}
