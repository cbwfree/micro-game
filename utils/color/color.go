package color

import (
	"fmt"
	"regexp"
	"runtime"
	"strings"
)

const (
	TplShowColor = "\x1b[%sm%s\x1b[0m" // 颜色模板
	TplColorSet  = "\x1b[%sm"
	TplResetSet  = "\x1b[0m"        // 重置/正常 关闭所有属性。
	CodeExpr     = `\033\[[\d;?]+m` // 清理颜色代码正则, 例如 "\033[1;36mText\x1b[0m"
)

var (
	isSupport = runtime.GOOS != "windows" // 不支持Windows
	codeRegex = regexp.MustCompile(CodeExpr)
)

// RenderStr 渲染字符串输出颜色
func RenderStr(code string, str string) string {
	if len(code) == 0 || str == "" {
		return str
	}
	if !isSupport {
		return ClearCode(str)
	}
	return fmt.Sprintf(TplShowColor, code, str)
}

// RenderCode render message by color code.
func RenderCode(code string, args ...interface{}) string {
	var message string
	if ln := len(args); ln == 0 {
		return ""
	}

	message = fmt.Sprint(args...)
	if len(code) == 0 {
		return message
	}

	if !isSupport {
		return ClearCode(message)
	}

	return fmt.Sprintf(TplShowColor, code, message)
}

// ClearCode 清除色彩.
// eg: "\033[36;1mText\x1b[0m" -> "Text"
func ClearCode(str string) string {
	return codeRegex.ReplaceAllString(str, "")
}

// Set set console color attributes
func Set(colors ...Color) (int, error) {
	if !isSupport { // not enable
		return 0, nil
	}
	return fmt.Printf(TplColorSet, colors2code(colors...))
}

// Reset reset console color attributes
func Reset() (int, error) {
	if !isSupport { // not enable
		return 0, nil
	}
	return fmt.Print(TplResetSet)
}

// convert colors to code. return like "32;45;3"
func colors2code(colors ...Color) string {
	if len(colors) == 0 {
		return ""
	}

	var codes []string
	for _, color := range colors {
		codes = append(codes, color.Code())
	}

	return strings.Join(codes, ";")
}
