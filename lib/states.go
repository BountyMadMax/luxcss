package lib

type State struct {
	name     string
	function func(value string) string
}

func States() []State {
	return []State{
		State{"hover", func(value string) string { return "test-state" }},
	}
}
