package lib

import (
	"fmt"
	"regexp"
)

type Style struct {
	name  regexp.Regexp
	build func(args []string) []string
}

func Styles() []Style {
	styles := make([]Style, 5)
	styles = []Style{
		{
			*regexp.MustCompile("w-[[:digit:]]"),
			func(args []string) []string {
				return []string{fmt.Sprintf("width: %srem", args[0])}
			},
		},
		{
			*regexp.MustCompile("h-[[:digit:]]"),
			func(args []string) []string {
				return []string{fmt.Sprintf("height: %srem", args[0])}
			},
		},
		{
			*regexp.MustCompile("m-[[:digit:]]"),
			func(args []string) []string {
				return []string{fmt.Sprintf("margin: %srem", args[0])}
			},
		},
		{
			*regexp.MustCompile("p-[[:digit:]]"),
			func(args []string) []string {
				return []string{fmt.Sprintf("padding: %srem", args[0])}
			},
		},
		{
			*regexp.MustCompile("cursor-(auto|default|pointer|wait|text|move|help|not-allowed|none|context-menu|progress|cell|crosshair|vertical-text|alias|copy|no-drop|grab|grabbing|all-scroll|col-resize|row-resize|n-resize|e-resize|s-resize|w-resize|ne-resize|nw-resize|se-resize|sw-resize|ew-resize|ns-resize|nesw-resize|nwse-resize|zoom-in|zoom-out)"),
			func(args []string) []string {
				return []string{fmt.Sprintf("cursor: %s", args[0])}
			},
		},
	}

	return styles
}
