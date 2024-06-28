package lib

import (
	"fmt"
)

func WrapWithBreakpoint(content []string, breakpoint Breakpoint) []string {
	wrapped := make([]string, len(content)+2)
	wrapped = append(wrapped, fmt.Sprintf("@media (min-width: %d) {", breakpoint.value))
	wrapped = append(wrapped, content...)
	wrapped = append(wrapped, "}")
	return wrapped
}
