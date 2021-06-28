// Prints the content found at each specified url
// Test with ./fetch http://gopl.io
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, requestErr := http.Get(url)

		if requestErr != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", requestErr)
			os.Exit(1)
		}

		bytes, ioErr := ioutil.ReadAll(resp.Body)

		resp.Body.Close()

		if ioErr != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, ioErr)
			os.Exit(1)
		}

		fmt.Printf("%s", bytes)
	}
}
