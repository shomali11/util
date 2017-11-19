# parallelizer [![Go Report Card](https://goreportcard.com/badge/github.com/shomali11/parallelizer)](https://goreportcard.com/report/github.com/shomali11/parallelizer) [![GoDoc](https://godoc.org/github.com/shomali11/parallelizer?status.svg)](https://godoc.org/github.com/shomali11/parallelizer) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Simplifies creating a pool of workers that execute jobs in parallel

## Features

* Easy to use
* Context Support
* Customizable Pool Size
    * Default number of workers is 10
* Customizable Job Queue Size
    * Default size is 100

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

Running multiple slow function calls in parallel with a context with a short timeout.
_Note: The timeout will not kill the routines. It will just stop waiting for them to finish_

```go
package main

import (
	"context"
	"fmt"
	"github.com/shomali11/parallelizer"
	"time"
)

func main() {
	group := parallelizer.NewGroup()
	defer group.Close()

	group.Add(func() {
		time.Sleep(2 * time.Second)

		fmt.Println("Finished work 1")
	})

	group.Add(func() {
		time.Sleep(2 * time.Second)

		fmt.Println("Finished work 2")
	})

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err := group.Wait(parallelizer.WithContext(ctx))

	fmt.Println("Done")
	fmt.Printf("Error: %v", err)
	fmt.Println()

	time.Sleep(2 * time.Second)
}
```

Output:

```text
Done
Error: context deadline exceeded
Finished work 2
Finished work 1
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

Running multiple function calls in parallel with a small worker pool and job queue size. _Note: the `Add` call blocks until there is space to push into the Job Queue_

```go
package main

import (
	"fmt"
	"github.com/shomali11/parallelizer"
	"time"
)

func main() {
	group := parallelizer.NewGroup(parallelizer.WithPoolSize(1), parallelizer.WithJobQueueSize(1))
	defer group.Close()

	for i := 1; i <= 10; i++ {
		group.Add(func() {
			time.Sleep(time.Second)
		})

		fmt.Println("Job added at", time.Now().Format("04:05"))
	}

	err := group.Wait()

	fmt.Println()
	fmt.Println("Done")
	fmt.Printf("Error: %v", err)
}
```

Output:

```text
Job added at 00:12
Job added at 00:13
Job added at 00:14
Job added at 00:15
Job added at 00:16
Job added at 00:17
Job added at 00:18
Job added at 00:19
Job added at 00:20
Job added at 00:21

Done
Error: <nil>
```

## Example 6

Running multiple function calls in parallel with a large enough worker pool and job queue size. _Note: In here the `Add` calls did not block because there was plenty of space in the Job Queue_

```go
package main

import (
	"fmt"
	"github.com/shomali11/parallelizer"
	"time"
)

func main() {
	group := parallelizer.NewGroup(parallelizer.WithPoolSize(10), parallelizer.WithJobQueueSize(10))
	defer group.Close()

	for i := 1; i <= 10; i++ {
		group.Add(func() {
			time.Sleep(time.Second)
		})

		fmt.Println("Job added at", time.Now().Format("04:05"))
	}

	err := group.Wait()

	fmt.Println()
	fmt.Println("Done")
	fmt.Printf("Error: %v", err)
}
```

Output:

```text
Job added at 00:30
Job added at 00:30
Job added at 00:30
Job added at 00:30
Job added at 00:30
Job added at 00:30
Job added at 00:30
Job added at 00:30
Job added at 00:30
Job added at 00:30

Done
Error: <nil>
```

## Example 7

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
		fmt.Println("Finished work")
	})

	fmt.Println("We did not wait!")

	time.Sleep(time.Second)
}
```

Output:

```text
We did not wait!
Finished work
```

## Example 8

Showing an example with a mixture of `Add` and `Wait` calls.

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

	fmt.Println("Waiting for workers 1 and 2 to finish")

	group.Wait()

	fmt.Println("Workers 1 and 2 have finished")

	group.Add(func() {
		fmt.Println("Worker 3")
	})

	fmt.Println("Waiting for worker 3 to finish")

	group.Wait()

	fmt.Println("Worker 3 has finished")
}

```

Output:

```text
Waiting for workers 1 and 2 to finish
Worker 1
Worker 2
Workers 1 and 2 have finished
Waiting for worker 3 to finish
Worker 3
Worker 3 has finished
```