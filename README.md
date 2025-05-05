# uniqueslice
The uniqueslice package provides facilities for canonicalizing ("interning") slices of comparable values.

## Documentation

[GoDoc](https://pkg.go.dev/github.com/Snawoot/uniqueslice)

## Example

```go
package main

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/Snawoot/uniqueslice"
)

func main() {
	r := strings.NewReader(`123
1234
123`)
	scanner := bufio.NewScanner(r)
	var handles []uniqueslice.Handle[[]byte, byte]
	for scanner.Scan() {
		handles = append(handles, uniqueslice.Make(scanner.Bytes()))
	}
	if err := scanner.Err(); err != nil {
		panic(fmt.Errorf("scanner error: %v", err))
	}
	for i, h := range handles {
		fmt.Printf("handles[%d] = %v (%q)\n", i, h.Value(), h.Value())
	}
	if handles[0] != handles[2] {
		panic("handles[0] != handles[2]")
	}
	if handles[0] == handles[1] {
		panic("handles[0] == handles[1]")
	}
}
```
