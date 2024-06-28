package lib

type Breakpoint struct {
	name        string
	value       string
	escapedName string
}

func Breakpoints() [5]Breakpoint {
	return [5]Breakpoint{
		Breakpoint{"sm", "640px", "sm"},
		Breakpoint{"md", "768px", "md"},
		Breakpoint{"lg", "1024px", "lg"},
		Breakpoint{"xl", "1280px", "xl"},
		Breakpoint{"2xl", "1536px", "\\32xl"},
	}
}
