package JSN

import (
	"bytes" // standard byte slice manipulation package
	"compress/zlib"
	"fmt"
	"io" //standard input output package
	"os" //standard operating system package
)

func Inflate() {
	// io.ReadAll expects a type implementing io.Reader interface and reads contents from it till EOF
	// os.Stdin is type defined in os package implementing io.Reader interface
	// Read method on os.Stdin returns Standard Input received by programme
	// In our use case, we are pipelining ouput of cat command to our programme
	deflated, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		return
	}

	// deflated is of type []byte
	// bytes.Newbuffer takes a []byte and returns a Reader which implements io.Reader
	b := bytes.NewReader(deflated)

	// NewReader expects a type implementing io.Reader interface
	// It reads contents of the input, decompresses the data and return a type implementing io.ReadCloser
	// io.ReadCloser is embedded interface implementing both io.Reader and io.Closer interface
	z, err := zlib.NewReader(b)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Closing z using method defined in Closer interface to release resources and avoid memory leaks
	z.Close()

	// Finally we read decompressed data from z using io.ReadAll as z implements io.Reader interface
	inflated, err := io.ReadAll(z)
	if err != nil {
		fmt.Println(err)
		return
	}

	//TypeCasting inflated to string or else a array containing Ascii code of string characters will be printed
	fmt.Println(string(inflated))
}
