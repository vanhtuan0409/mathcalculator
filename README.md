# Math calculator

A simple math expression evaluation written in Go (ultilize
[rtokenizer](https://github.com/vanhtuan0409/rtokenize))

### Installation

```
go get https://github.com/vanhtuan0409/mathcalculator
```

### API

```go
func Evaluate(expression string) (float64, error) {}
```

Currently only support basic math operator includes:

* Add (`+`)
* Substract (`-`)
* Multiply (`*`)
* Divide (`/`)
* Bracket (`()`)

### Example

```go
package main

import (
	"fmt"

	"github.com/vanhtuan0409/mathcalculator"
)

func main() {
	expression := "1+2    * (3 + 4) - 5"
	result, err := mathcalculator.Evaluate(expression)
	if err != nil {
		fmt.Println(err)
	} else {
        // Expected result == 10
		fmt.Println(expression, "=", result)
	}
}
```
