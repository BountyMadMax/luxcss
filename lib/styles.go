package lib

import (
	"fmt"
	"strconv"
)

type Style struct {
	prefix       string
	options      []string
	build        func(args []string) []string
	customStyles bool
}

func Styles() []Style {
	return []Style{
		{
			"w-",
			[]string{"[[:digit:]]"},
			func(args []string) []string {
				if _, err := strconv.Atoi(args[0]); err == nil {
					return []string{fmt.Sprintf("width: %srem;", args[0])}
				}
				return []string{fmt.Sprintf("width: %s;", args[0])}
			},
			true,
		},
		{
			"h-",
			[]string{"[[:digit:]]"},
			func(args []string) []string {
				if _, err := strconv.Atoi(args[0]); err == nil {
					return []string{fmt.Sprintf("height: %srem;", args[0])}
				}
				return []string{fmt.Sprintf("height: %s;", args[0])}
			},
			true,
		},
		{
			"m-",
			[]string{"[[:digit:]]"},
			func(args []string) []string {
				if _, err := strconv.Atoi(args[0]); err == nil {
					return []string{fmt.Sprintf("margin: %srem;", args[0])}
				}
				return []string{fmt.Sprintf("margin: %s;", args[0])}
			},
			true,
		},
		{
			"p-",
			[]string{"[[:digit:]]"},
			func(args []string) []string {
				if _, err := strconv.Atoi(args[0]); err == nil {
					return []string{fmt.Sprintf("padding: %srem;", args[0])}
				}
				return []string{fmt.Sprintf("padding: %s;", args[0])}
			},
			true,
		},
		{
			"cursor-",
			[]string{"auto", "default", "pointer", "wait", "text", "move", "help", "not-allowed", "none", "context-menu", "progress", "cell", "crosshair", "vertical-text", "alias", "copy", "no-drop", "grap", "grabbing", "all-scroll", "col-resize", "row-resize", "n-resize", "e-resize", "s-resize", "w-resize", "ne-resize", "nw-resize", "se-resize", "sw-resize", "ew-resize", "ns-resize", "nesw-resize", "nwse-resize", "zoom-in", "zoom-out"},
			func(args []string) []string {
				return []string{fmt.Sprintf("cursor: %s;", args[0])}
			},
			false,
		},
		{
			"",
			[]string{"block", "inline-block", "inline", "flex", "inline-flex", "table", "inline-table", "table-caption", "table-cell", "table-column", "table-column-group", "table-footer-group", "table-header-group", "table-row-group", "table-row", "flow-root", "grid", "inline-grid", "contents", "list-item", "hidden"},
			func(args []string) []string {
				return []string{fmt.Sprintf("display: %s;", args[0])}
			},
			false,
		},
		{
			"gap-",
			[]string{"[[:digit:]]"},
			func(args []string) []string {
				if _, err := strconv.Atoi(args[0]); err == nil {
					return []string{fmt.Sprintf("gap: %srem;", args[0])}
				}
				return []string{fmt.Sprintf("gap: %s", args[0])}
			},
			true,
		},
	}
}
