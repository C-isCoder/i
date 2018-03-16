package i

import (
	"fmt"
	"bytes"
	"io"
	"os"
)
func main() {
	var b bytes.Buffer

	b.Write([]byte("hello"))

	fmt.Fprintf(&b," World")

	io.Copy(os.Stdout,&b)
}
