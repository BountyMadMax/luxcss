package lib

type Breakpoint struct {
	name        string
	value       string
	escapedName string
}

func Breakpoints() [5]Breakpoint {
	return [5]Breakpoint{
		{"sm", "640px", "sm"},
		{"md", "768px", "md"},
		{"lg", "1024px", "lg"},
		{"xl", "1280px", "xl"},
		{"2xl", "1536px", "\\32xl"},
	}
}
