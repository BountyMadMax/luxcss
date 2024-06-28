package lib

import "fmt"

type Style struct {
	name    string
	pattern string
	build   func(args []string) []string
}

func Styles() []Style {
	styles := make([]Style, 5)
	styles = append(styles, Style{"width", "w-\\d", func(args []string) []string {
		output := make([]string, 1)
		output = append(output, fmt.Sprintf("width: %drem", args[0]))
		return output
	}})

	return styles
}
