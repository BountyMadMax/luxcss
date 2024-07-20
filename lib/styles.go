package lib

import (
	"fmt"
)

type Style struct {
	name  string
	build func(args []string) []string
}

func Styles() []Style {
	styles := make([]Style, 5)
	styles = []Style{
		{
			"w-[[:digit:]]",
			func(args []string) []string {
				return []string{fmt.Sprintf("width: %srem;", args[0])}
			},
		},
		{
			"h-[[:digit:]]",
			func(args []string) []string {
				return []string{fmt.Sprintf("height: %srem;", args[0])}
			},
		},
		{
			"m-[[:digit:]]",
			func(args []string) []string {
				return []string{fmt.Sprintf("margin: %srem;", args[0])}
			},
		},
		{
			"p-[[:digit:]]",
			func(args []string) []string {
				return []string{fmt.Sprintf("padding: %srem;", args[0])}
			},
		},
		{
			"cursor-(auto|default|pointer|wait|text|move|help|not-allowed|none|context-menu|progress|cell|crosshair|vertical-text|alias|copy|no-drop|grab|grabbing|all-scroll|col-resize|row-resize|n-resize|e-resize|s-resize|w-resize|ne-resize|nw-resize|se-resize|sw-resize|ew-resize|ns-resize|nesw-resize|nwse-resize|zoom-in|zoom-out)",
			func(args []string) []string {
				return []string{fmt.Sprintf("cursor: %s;", args[0])}
			},
		},
		{
			"(block|inline-block|inline|flex|inline-flex|table|inline-table|table-caption|table-cell|table-column|table-column-group|table-footer-group|table-header-group|table-row-group|table-row|flow-root|grid|inline-grid|contents|list-item|hidden)",
			func(args []string) []string {
				return []string{fmt.Sprintf("display: %s;", args[0])}
			},
		},
	}

	return styles
}
