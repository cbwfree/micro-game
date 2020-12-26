package log

import (
	"fmt"
	"github.com/cbwfree/micro-game/utils/color"
	"github.com/cbwfree/micro-game/utils/debug"
	"github.com/micro/go-micro/v2/logger"
	"os"
	"strings"
	"time"
)

var (
	Logger       *logger.Helper
	disableColor = false
)

func init() {
	lvl, err := logger.GetLevel(os.Getenv("GAME_LOG_LEVEL"))
	if err != nil {
		lvl = logger.InfoLevel
	}

	Logger = logger.NewHelper(logger.NewLogger(logger.WithLevel(lvl)))
}

// 禁用颜色
func SetDisableColor() {
	disableColor = true
}

// 显示颜色
func SetEnableColor() {
	disableColor = false
}

// 获取当前日志级别
func GetLevel() logger.Level {
	return Logger.Options().Level
}

// 检查当前日志输出级别
func CheckLevel(level logger.Level) bool {
	return logger.V(level, Logger)
}

// 是否Trace模式
func IsTrace() bool {
	return CheckLevel(logger.TraceLevel)
}

// 是否开发模式
func IsDev() bool {
	return CheckLevel(logger.DebugLevel)
}

func Trace(tpl string, a ...interface{}) {
	frames := debug.GetCallers(3, 10)
	var lines []string
	for _, frame := range frames {
		lines = append(lines, format(color.Trace, " > %s:%d %s", frame.File, frame.Line, frame.Func))
	}
	Logger.Logf(logger.TraceLevel, format(color.Trace, "%s\n%s", tpl, strings.Join(lines, "\n")), a...)
}

func Debug(tpl string, a ...interface{}) {
	Logger.Log(logger.DebugLevel, format(color.Debug, tpl, a...))
}

func Info(tpl string, a ...interface{}) {
	Logger.Log(logger.InfoLevel, format(color.Info, tpl, a...))
}

func Warn(tpl string, a ...interface{}) {
	Logger.Log(logger.WarnLevel, format(color.Warn, tpl, a...))
}

func Error(tpl string, a ...interface{}) {
	Logger.Log(logger.ErrorLevel, format(color.Error, tpl, a...))
}

func Fatal(tpl string, a ...interface{}) {
	Logger.Log(logger.FatalLevel, format(color.Fatal, tpl, a...))
}

func Printf(tpl string, a ...interface{}) {
	t := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("%s %v", t, fmt.Sprintf(tpl, a...))
}

func Println(a ...interface{}) {
	t := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("%s %v\n", t, fmt.Sprint(a...))
}

// 格式化字符串
func format(theme *color.Theme, tpl string, a ...interface{}) string {
	if disableColor || theme == nil {
		return fmt.Sprintf(tpl, a...)
	}
	return theme.Text(tpl, a...)
}
