# util [![Go Report Card](https://goreportcard.com/badge/github.com/shomali11/util)](https://goreportcard.com/report/github.com/shomali11/util) [![GoDoc](https://godoc.org/github.com/shomali11/util?status.svg)](https://godoc.org/github.com/shomali11/util) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A group of generic useful utility functions

## Usage

Using `govendor` [github.com/kardianos/govendor](https://github.com/kardianos/govendor):

```
govendor fetch github.com/shomali11/util
```

## Dependencies

* `parallelizer` [github.com/shomali11/parallelizer](https://github.com/shomali11/parallelizer):


# Examples

## Concurrency

```go
package main

import (
	"github.com/shomali11/util/concurrency"
	"time"
	"fmt"
)

func main() {
    func1 := func() {
            for char := 'a'; char < 'a' + 3; char++ {
                fmt.Printf("%c ", char)
            }
    }
    
    func2 := func() {
            for number := 1; number < 4; number++ {
                fmt.Printf("%d ", number)
            }
    }
    
    concurrency.Parallelize(func1, func2)  // a 1 b 2 c 3
    
    concurrency.ParallelizeTimeout(time.Minute, func1, func2)  // a 1 b 2 c 3
}
```

## Strings

```go
package main

import (
	"github.com/shomali11/util/strings"
	"fmt"
)

func main() {
	fmt.Println(strings.IsEmpty(""))     // true
	fmt.Println(strings.IsEmpty("text")) // false
	fmt.Println(strings.IsEmpty("	"))  // false

	fmt.Println(strings.IsNotEmpty(""))     // false
	fmt.Println(strings.IsNotEmpty("text")) // true
	fmt.Println(strings.IsNotEmpty("	")) // true

	fmt.Println(strings.IsBlank(""))     // true
	fmt.Println(strings.IsBlank("	"))  // true
	fmt.Println(strings.IsBlank("text")) // false

	fmt.Println(strings.IsNotBlank(""))     // false
	fmt.Println(strings.IsNotBlank("	")) // false
	fmt.Println(strings.IsNotBlank("text")) // true

	fmt.Println(strings.Length(""))                                          // 0
	fmt.Println(strings.Length("X"))                                         // 1
	fmt.Println(strings.Length("b\u0301"))                                   // 1
	fmt.Println(strings.Length("😎⚽"))                                      // 2
	fmt.Println(strings.Length("Les Mise\u0301rables"))                      // 14
	fmt.Println(strings.Length("ab\u0301cde"))                               // 5
	fmt.Println(strings.Length("This `\xc5` is an invalid UTF8 character"))  // 37
	fmt.Println(strings.Length("The quick bròwn 狐 jumped over the lazy 犬")) // 40

	fmt.Println(strings.Reverse(""))                                            // ""
	fmt.Println(strings.Reverse("X"))                                           // "X"
	fmt.Println(strings.Reverse("😎⚽"))                                        // "⚽😎"
	fmt.Println(strings.Reverse("Les Mise\u0301rables"))                        // "selbare\u0301siM seL"
	fmt.Println(strings.Reverse("This `\xc5` is an invalid UTF8 character"))    // "retcarahc 8FTU dilavni na si `�` sihT"
	fmt.Println(strings.Reverse("The quick bròwn 狐 jumped over the lazy 犬"))  // "犬 yzal eht revo depmuj 狐 nwòrb kciuq ehT"
}
```

## Conditions

```go
package main

import (
	"github.com/shomali11/util/conditions"
	"fmt"
)

func main() {
	fmt.Println(conditions.IfThen(1 == 1, "Yes")) // "Yes"
	fmt.Println(conditions.IfThen(1 != 1, "Woo")) // nil
	fmt.Println(conditions.IfThen(1 < 2, "Less")) // "Less"

	fmt.Println(conditions.IfThenElse(1 == 1, "Yes", false)) // "Yes"
	fmt.Println(conditions.IfThenElse(1 != 1, nil, 1))       // 1
	fmt.Println(conditions.IfThenElse(1 < 2, nil, "No"))     // nil

	fmt.Println(conditions.DefaultIfNil(nil, nil))  // nil
	fmt.Println(conditions.DefaultIfNil(nil, ""))   // ""
	fmt.Println(conditions.DefaultIfNil("A", "B"))  // "A"
	fmt.Println(conditions.DefaultIfNil(true, "B")) // true
	fmt.Println(conditions.DefaultIfNil(1, false))  // 1

	fmt.Println(conditions.FirstNonNil(nil, nil))                // nil
	fmt.Println(conditions.FirstNonNil(nil, ""))                 // ""
	fmt.Println(conditions.FirstNonNil("A", "B"))                // "A"
	fmt.Println(conditions.FirstNonNil(true, "B"))               // true
	fmt.Println(conditions.FirstNonNil(1, false))                // 1
	fmt.Println(conditions.FirstNonNil(nil, nil, nil, 10))       // 10
	fmt.Println(conditions.FirstNonNil(nil, nil, nil, nil, nil)) // nil
	fmt.Println(conditions.FirstNonNil())                        // nil
}
```

## Errors

```go
package main

import (
	"github.com/shomali11/util/errors"
	"fmt"
)

func main() {
	fmt.Println(errors.DefaultErrorIfNil(nil, "Cool"))                // "Cool"
	fmt.Println(errors.DefaultErrorIfNil(errors.New("Oops"), "Cool")) // "Oops"
}
```

## Manipulations

```go
package main

import (
	"github.com/shomali11/util/manipulations"
	"math/rand"
	"time"
	"fmt"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())

	array := []interface{}{"a", "b", "c"}
	manipulations.Shuffle(array, source)

	fmt.Println(array) // [c b a]
}
```

## Calculations

Given two integers representing the numerator and denominator of a fraction, return the fraction in string format with the repeating part enclosed in parentheses

```go
package main

import (
	"github.com/shomali11/util/calculations"
	"fmt"
)

func main() {
    fmt.Println(calculations.Divide(0, 0))     // "ERROR"
    fmt.Println(calculations.Divide(1, 2))     // "0.5(0)"
    fmt.Println(calculations.Divide(0, 3))     // "0.(0)"
    fmt.Println(calculations.Divide(10, 3))    // "3.(3)"
    fmt.Println(calculations.Divide(22, 7))    // "3.(142857)"
    fmt.Println(calculations.Divide(100, 145)) // "0.(6896551724137931034482758620)"
}
```

## PrettyJson

Pretty JSON Prints!

```go
package main

import (
	"github.com/shomali11/util/json"
	"fmt"
)

func main() {
    x := map[string]interface{}{"number": 1, "string": "cool", "bool": true, "float": 1.5}    
    fmt.Println(json.PrettyJson(x))
}
```

```json
{
	"bool": true,
	"float": 1.5,
	"number": 1,
	"string": "cool"
}
```
