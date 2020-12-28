package color

// Theme definition. extends from Style
type Theme struct {
	// Name theme name
	Name string
	// Style for the theme
	Style
}

// NewTheme instance
func NewTheme(name string, style Style) *Theme {
	return &Theme{name, style}
}

// Save to themes map
func (t *Theme) Save() {
	AddTheme(t.Name, t.Style)
}

/*************************************************************
 * Theme: internal themes
 *************************************************************/

// internal themes(like bootstrap style)
// Usage:
// 	color.Info.Print("message")
// 	color.Info.Printf("a %s message", "test")
// 	color.Warn.Println("message")
// 	color.Error.Println("message")
var (
	// Trace color style
	Trace = &Theme{"trace", Style{OpReset, FgBlue}}
	// Debug color style
	Debug = &Theme{"debug", Style{OpReset, FgDefault}}
	// Info color style
	Info = &Theme{"info", Style{OpReset, FgCyan}}
	// Warn color style
	Warn = &Theme{"warn", Style{OpBold, FgYellow}}
	// Error color style
	Error = &Theme{"error", Style{OpBold, FgRed}}
	// Question color style
	Question = &Theme{"question", Style{OpBold, FgMagenta}}
	// Success color style
	Success = &Theme{"success", Style{OpBold, FgGreen}}
	// Danger color style
	Fatal = &Theme{"danger", Style{FgLightWhite, BgRed}}

	// Primary color style
	Primary = &Theme{"primary", Style{OpReset, FgLightBlue}}
	// Secondary color style
	Comment = &Theme{"secondary", Style{OpReset, FgDarkGray}}
	// Note color style
	Note = &Theme{"note", Style{OpBold, FgLightCyan}}
)

// Themes internal defined themes.
// Usage:
// 	color.Themes["info"].Println("message")
var Themes = map[string]*Theme{
	"trace": Trace,
	"debug": Debug,
	"info":  Info,
	"warn":  Warn,
	"error": Error,
	"fatal": Fatal,

	"note":     Note,
	"success":  Success,
	"comment":  Comment,
	"primary":  Primary,
	"question": Question,
}

// AddTheme add a theme and style
func AddTheme(name string, style Style) {
	Themes[name] = NewTheme(name, style)
	Styles[name] = style
}

// GetTheme get defined theme by name
func GetTheme(name string) *Theme {
	return Themes[name]
}
