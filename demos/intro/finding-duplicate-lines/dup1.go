// Print the text of eaach line that appears more
// once in the standard input, preceded by its count.
// You can pipe data.txt file: ./dup1 < data.txt
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
		// NOTE: Ignoring potential errors from input.Err()
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
