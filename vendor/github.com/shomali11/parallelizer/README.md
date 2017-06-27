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
	group := parallelizer.DefaultGroup()

	group.Add(func() {
		for char := 'a'; char < 'a'+3; char++ {
			fmt.Printf("%c ", char)
		}
	})

	group.Add(func() {
		for number := 1; number < 4; number++ {
			fmt.Printf("%d ", number)
		}
	})

	err := group.Run()

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
_Note: The timeout will not kill the routines. It will just stop waiting for them to finish_

```go
package main

import (
	"fmt"
	"github.com/shomali11/parallelizer"
	"time"
)

func main() {
	options := &parallelizer.Options{Timeout: time.Second}
	group := parallelizer.NewGroup(options)

	group.Add(func() {
		time.Sleep(2 * time.Second)

		for char := 'a'; char < 'a'+3; char++ {
			fmt.Printf("%c ", char)
		}
	})

	group.Add(func() {
		time.Sleep(2 * time.Second)

		for number := 1; number < 4; number++ {
			fmt.Printf("%d ", number)
		}
	})

	err := group.Run()

	fmt.Println()
	fmt.Println("Done")
	fmt.Printf("Error: %v", err)
	fmt.Println()

	time.Sleep(3 * time.Second)
}
```

Output:

```text

Done
Error: timeout
a 1 b 2 c 3
```

## Example 3

Running multiple function calls in parallel with a large enough worker pool.

```go
package main

import (
	"fmt"
	"github.com/shomali11/parallelizer"
)

func main() {
	options := &parallelizer.Options{WorkerPoolSize: 10}
	group := parallelizer.NewGroup(options)

	for i := 1; i <= 10; i++ {
		i := i
		group.Add(func() {
			fmt.Print(i, " ")
		})
	}

	err := group.Run()

	fmt.Println()
	fmt.Println("Done")
	fmt.Printf("Error: %v", err)
}
```

Output:

```text
7 6 3 2 8 9 5 10 1 4  
Done
Error: <nil>
```

## Example 4

Running multiple function calls with 1 worker. _Note: the functions are no longer executed in parallel but sequentially_

```go
package main

import (
	"fmt"
	"github.com/shomali11/parallelizer"
)

func main() {
	options := &parallelizer.Options{WorkerPoolSize: 1}
	group := parallelizer.NewGroup(options)

	for i := 1; i <= 10; i++ {
		i := i
		group.Add(func() {
			fmt.Print(i, " ")
		})
	}

	err := group.Run()

	fmt.Println()
	fmt.Println("Done")
	fmt.Printf("Error: %v", err)
}
```

Output:

```text
1 2 3 4 5 6 7 8 9 10 
Done
Error: <nil>
```