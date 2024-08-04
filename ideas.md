```scss
// main.scss

@use "fmt";

@func main() {
  $greeting = "Hello world";
  @import fmt.println($greeting);
}
```

Difference between `@mixin` and `@func` is that `@func` only returns `return` statement

```scss
// fmt.scss
@func print-ln($format) {
  // return something
}
```

## Go source

### fmt.Printf

```go
// fmt
...
package fmt
...
import (
	"internal/fmtsort"
	"io"
	"os"
	"reflect"
	"strconv"
	"sync"
	"unicode/utf8"
)
...
// Fprintf formats according to a format specifier and writes to w.
// It returns the number of bytes written and any write error encountered.
func Fprintf(w io.Writer, format string, a ...any) (n int, err error) {
	p := newPrinter()
	p.doPrintf(format, a)
	n, err = w.Write(p.buf)
	p.free()
	return
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
func Printf(format string, a ...any) (n int, err error) {
	return Fprintf(os.Stdout, format, a...)
}
...
```
