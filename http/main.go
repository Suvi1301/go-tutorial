package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	/*
		Initialise a byte size of a significant size.
		Read doesn't resize the byte slice.
		Then we can pass the bs to Read which will inject the response body into bs.
		But this is tedious. We want to simplify this.

		bs := make([]byte, 99999)
		resp.Body.Read(bs)
		fmt.Println(string(bs))
	*/

	/*
		We can simplify the above by using something that implements the Writer interface.
		Writer interface essentially reverses the Reader interface.
		It takes a byte slice and output to a source of output. e.g. to terminal etc.
		Writer interface should implement a Write(p []byte) (n int, err error) function.
	*/

	/*
		Copy takes a Reader Interface as a source and outputs to a Writer Interface.
		resp.Body implements Reader Interface which is our source.
		We need a output that implements a Writer interface. e.g. os.Stdout
	*/
	// io.Copy(os.Stdout, resp.Body)

	lw := logWriter{}
	io.Copy(lw, resp.Body) // here we are using our own custom implementation of a Writer in lw.
}

/*
	http.Get returns a *Response
	Response is a struct with properties:
	Status: string
	StatusCode: int
	Body: io.ReadCloser

	The struct can have an interface property which just means that the property can
	have any value assigned to it as long as it fulfills the interface.

	io.ReadCloser is an interface with two properties:
	- io.Reader
	- io.Closer

	The ReadCloser interface is just a meta interface which says you must fulfil both
	Reader and Closer interfaces to be considered a ReadCloser interface.

	io.Reader is an interface with method:
	- Read([]byte) (int, error)

	io.Closer is an interface with method:
	- Close() (error)
*/

// Our custom Write function implementation which means that logWriter now implements the io.Writer interface.
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
