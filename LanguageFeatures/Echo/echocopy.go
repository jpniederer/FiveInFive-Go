// Echocopy prints the command-line arguments.
// build with: "go build"
// to run: "./Echo arg1 arg2 arg3 ..."
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Variables are implicitly defined to the zero value. 0, "", etc.
	var s string
	var separator = ". "
	var nl = "\n"
	// Check it out, no parens. {} required.
	fmt.Printf("Echo output from %s: \n", os.Args[0])
	for i := 1; i < len(os.Args); i++ {
		s += strconv.Itoa(i) + separator + os.Args[i] + nl
		separator = ". "
	}
	fmt.Println(s)

	// One liner: fmt.Println(strings.Join(os.Args[1:], " "))
}
