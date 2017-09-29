package stringutil

func ReverseTwo(s string) string {
	r := []rune(s)  // create and assign variable, array of string?

	for i,j := 0, len(r)-1;i < len(r)/2;i,j = i+1,j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r)
}

// within a package, stuff from one file will be accessible inside another file -- eg. reverseTwo will be accessible in names.go
