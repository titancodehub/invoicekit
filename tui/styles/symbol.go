package styles

import "fmt"

var Symbol = map[string]string{
	"check": "\u2713",
	"cross": "\u2A2F",
}

func WithErrorSymbol(v string) string {
	return ErrorTextStyle.Render(fmt.Sprintf("%s %s", Symbol["cross"], v))
}

func WithSuccessSymbol(v string) string {
	return SuccessTextStyle.Render(fmt.Sprintf("%s %s", Symbol["check"], v))
}
