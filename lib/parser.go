package lib

import (
	"os"
	"strings"
)

type Writer interface {
	writeFile(*os.File) error
}

type ExtractClass struct {
	// Class as extracted.
	class string
	// Values of the class as property => value. Should mostly only be one.
	values map[string]string
	// States applied to the class.
	states []StyleState
}

type StyleState struct {
	prefix *string
	suffix *string
}

type MediaTree struct {
	// Classes contained in the media query.
	classes []*ExtractClass
	// Values of the media.
	mediaValues []MediaValue
	// Parent if exists.
	parent *MediaTree
	// Childs if multiple media queries are used.
	// Example: breakpoint query -> print query
	childs []*MediaTree
}

type MediaValue struct {
	key   *string
	value string
}

func Parse(classes []string) ([]*ExtractClass, map[string]MediaTree) {
	// TODO: This implementation currently expects that the given classes only have one breakpoint.
	// But the best thing would be for them to be unchanged, since we need the full extracted class for later.

	var extractClasses []*ExtractClass
	var mediaTrees map[string]*MediaTree

	for _, class := range classes {
		// First element could be a breakpoint.
		// After that, an unknown amount of states could be.
		// At the end is the name with the value.
		splitted := strings.Split(class, ":")

		id, mediaTree, mediaTrees := applyBreakpoint(mediaTrees, splitted)

		id, mediaTree, extractClass := applyState()
	}

	return extractClasses, mediaTrees
}

func applyBreakpoint(mediaTrees map[string]*MediaTree, splitted []string) (int, *MediaTree, map[string]*MediaTree) {
	for _, breakpoint := range Breakpoints() {
		if breakpoint.name == splitted[0] {
			if mediaTrees[splitted[0]] == nil {
				mediaTrees[splitted[0]] = buildMediaTree()
			}

			return 1, mediaTrees[splitted[0]], mediaTrees
		}
	}

	return 0, nil, mediaTrees
}

func applyStates(mediaTree *MediaTree, splitted []string) (int, *MediaTree, ExtractClass) {

}

func buildMediaTree() *MediaTree {

}

func buildStyle() *ExtractClass {

}
