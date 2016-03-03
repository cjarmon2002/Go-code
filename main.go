package main

import (
	"fmt"
	// "os"
)

	
func main() {
	haystack := "The spice must flow"
	cut := strings.index(haystack[5:], " ")
	fmt.Println(cut)
}
