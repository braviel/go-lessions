package main

const repeatCount = 5

// Repeat function repeat a string for a number of n time(s)
func Repeat(character string, n int) string {
	var repeated string
	for i := 0; i < n; i++ {
		repeated += character
	}
	return repeated
}
