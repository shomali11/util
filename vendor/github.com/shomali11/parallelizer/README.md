# parallelizer [![Go Report Card](https://goreportcard.com/badge/github.com/shomali11/parallelizer)](https://goreportcard.com/report/github.com/shomali11/parallelizer) [![GoDoc](https://godoc.org/github.com/shomali11/parallelizer?status.svg)](https://godoc.org/github.com/shomali11/parallelizer) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Simplifies creating a pool of workers that execute jobs in parallel

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
	group := parallelizer.NewGroup()
	defer group.Close()

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

	err := group.Wait()

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
	group := parallelizer.NewGroup()
	defer group.Close()

	group.Add(func() {
		time.Sleep(2 * time.Second)
	})

	err := group.Wait(parallelizer.WithTimeout(time.Second))

	fmt.Println()
	fmt.Println("Done")
	fmt.Printf("Error: %v", err)
}
```

Output:

```text

Error: timeout
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
	group := parallelizer.NewGroup(parallelizer.WithPoolSize(10))
	defer group.Close()

	for i := 1; i <= 10; i++ {
		i := i
		group.Add(func() {
			fmt.Print(i, " ")
		})
	}

	err := group.Wait()

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
	group := parallelizer.NewGroup(parallelizer.WithPoolSize(1))
	defer group.Close()

	for i := 1; i <= 10; i++ {
		i := i
		group.Add(func() {
			fmt.Print(i, " ")
		})
	}

	err := group.Wait()

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

## Example 5

Showing an example without calling `Wait`

```go
package main

import (
	"fmt"
	"github.com/shomali11/parallelizer"
	"time"
)

func main() {
	group := parallelizer.NewGroup()
	defer group.Close()

	group.Add(func() {
		fmt.Print("Worker 1")
	})

	fmt.Println()
	fmt.Println("We did not wait!")

	time.Sleep(time.Second)
}
```

Output:

```text

We did not wait!
Worker 1
```

## Example 6

Showing an example with a mixuture of `Add` and `Wait` calls.

```go
package main

import (
	"fmt"
	"github.com/shomali11/parallelizer"
)

func main() {
	group := parallelizer.NewGroup()
	defer group.Close()

	group.Add(func() {
		fmt.Println("Worker 1")
	})

	group.Add(func() {
		fmt.Println("Worker 2")
	})

	group.Wait()

	fmt.Println("Workers 1 and 2 have finished")

	group.Add(func() {
		fmt.Println("Worker 3")
	})

	group.Wait()

	fmt.Println("Worker 3 has finished")
}
```

Output:

```text
Worker 2
Worker 1
Workers 1 and 2 have finished
Worker 3
Worker 3 has finished

```