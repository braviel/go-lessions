package di

import (
	"fmt"
	"io"
)

//Greet there
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
