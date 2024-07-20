package lib

import (
	"fmt"
)

func ApplyBreakpoint(breakpoint Breakpoint) (string, string, int) {
	return fmt.Sprintf("@media (min-width: %s) {", breakpoint.value), breakpoint.class, 1
}
