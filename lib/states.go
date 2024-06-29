package lib

type State struct {
	name         string
	apply        func(value string) string
	exscapedName string
}

func States() []State {
	return []State{
		{"hover", func(value string) string { return "test-hover" }, "hover"},
		{"focus", func(value string) string { return "test-focus" }, "focus"},
	}
}
