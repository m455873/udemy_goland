package stringutil
// can have multiple files declare themselves to be part of a particular package (eg. stringutil), notice that name.go is also part of this package

func Reverse(s string) string { // takes s string and returns a string
	return ReverseTwo(s)
}

// No semicolons after line
// need curly brackets start and end

//go build reverse.go reverseTwo.go won't produce output file -- why?
// go install will place pacakge

