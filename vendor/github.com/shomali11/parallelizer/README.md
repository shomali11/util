# parallelizer [![Go Report Card](https://goreportcard.com/badge/github.com/shomali11/parallelizer)](https://goreportcard.com/report/github.com/shomali11/parallelizer) [![GoDoc](https://godoc.org/github.com/shomali11/parallelizer?status.svg)](https://godoc.org/github.com/shomali11/parallelizer) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Simplifies the parallelization of function calls with an optional timeout

## Usage

Using `govendor` [github.com/kardianos/govendor](https://github.com/kardianos/govendor):

```
govendor fetch github.com/shomali11/parallelizer
```

# Examples

## Example 1

Running multiple function calls in parallel without a timeout.

```go
package main

import (
	"fmt"
	"github.com/shomali11/parallelizer"
)

func main() {
	func1 := func() {
		for char := 'a'; char < 'a'+3; char++ {
			fmt.Printf("%c ", char)
		}
	}

	func2 := func() {
		for number := 1; number < 4; number++ {
			fmt.Printf("%d ", number)
		}
	}

	runner := parallelizer.Runner{}
	err := runner.Run(func1, func2)

	fmt.Println()
	fmt.Println("Done")
	fmt.Printf("Error: %v", err)
}
```

Output:

```text
a 1 b 2 c 3 
Done
Error: <nil>
```

## Example 2

Running multiple slow function calls in parallel with a short timeout.

```go
package main

import (
	"fmt"
	"github.com/shomali11/parallelizer"
	"time"
)

func main() {
	func1 := func() {
		time.Sleep(time.Minute)

		for char := 'a'; char < 'a'+3; char++ {
			fmt.Printf("%c ", char)
		}
	}

	func2 := func() {
		time.Sleep(time.Minute)

		for number := 1; number < 4; number++ {
			fmt.Printf("%d ", number)
		}
	}

	runner := parallelizer.Runner{Timeout: time.Second}
	err := runner.Run(func1, func2)

	fmt.Println()
	fmt.Println("Done")
	fmt.Printf("Error: %v", err)
}
```

Output:

```text

Done
Error: timeout
```
