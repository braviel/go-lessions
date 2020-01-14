package main

const greetingPrefix = "Hello, "

// Hello helloworld
func Hello(who string) string {
	if who == "" {
		return greetingPrefix + "World"
	}
	return greetingPrefix + who
	//return who
}
