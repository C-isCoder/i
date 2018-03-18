package main

import (
	"fmt"
	"bytes"
	"io"
	"os"
)
func main() {
	var b bytes.Buffer

	b.Write([]byte("hello"))

	fmt.Fprint(&b,"World")

	io.Copy(os.Stdout,&b)
}
