# util [![Go Report Card](https://goreportcard.com/badge/github.com/shomali11/util)](https://goreportcard.com/report/github.com/shomali11/util) [![GoDoc](https://godoc.org/github.com/shomali11/util?status.svg)](https://godoc.org/github.com/shomali11/util) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A group of generic useful utility functions

## Usage

Using `govendor` [github.com/kardianos/govendor](https://github.com/kardianos/govendor):

```
govendor fetch github.com/shomali11/util
```

## Dependencies

* `parallelizer` [github.com/shomali11/parallelizer](https://github.com/shomali11/parallelizer)


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

## Hashes

```go
package main

import (
	"github.com/shomali11/util/hashes"
	"fmt"
)

func main() {
	fmt.Println(hashes.FNV32("Raed Shomali"))  // 424383802
	fmt.Println(hashes.FNV32a("Raed Shomali")) // 3711360776
	fmt.Println(hashes.FNV64("Raed Shomali"))  // 13852322269024953050
	fmt.Println(hashes.FNV64a("Raed Shomali")) // 17869303103005344072
	fmt.Println(hashes.MD5("Raed Shomali"))    // c313bc3b48fcfed9abc733429665b105
	fmt.Println(hashes.SHA1("Raed Shomali"))   // e0d66f6f09de72942e83289cc994b3c721ab34c5
	fmt.Println(hashes.SHA256("Raed Shomali")) // 75894b9be21065a833e57bfe4440b375fc216f120a965243c9be8b2dc36709c2
	fmt.Println(hashes.SHA512("Raed Shomali")) // 406e8d495140187a8b09893c30d054cf385ad7359855db0d2e0386c7189ac1c4667a4816d1b63a19f3d8ccdcbace7861ec4cc6ff5e2a1659c8f4360bda699b42
}
```

## Compressions

```go
package main

import (
	"github.com/shomali11/util/compressions"
	"fmt"
)

func main() {
	fmt.Println(compressions.Compress("Raed Shomali"))
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

	fmt.Println(strings.Left("", 5))            // "     "
	fmt.Println(strings.Left("X", 5))           // "X    "
	fmt.Println(strings.Left("😎⚽", 4))        // "😎⚽  "
	fmt.Println(strings.Left("ab\u0301cde", 8)) // "ab́cde   "

	fmt.Println(strings.Right("", 5))            // "     "
	fmt.Println(strings.Right("X", 5))           // "    X"
	fmt.Println(strings.Right("😎⚽", 4))        // "  😎⚽"
	fmt.Println(strings.Right("ab\u0301cde", 8)) // "   ab́cde"

	fmt.Println(strings.Center("", 5))            // "     "
	fmt.Println(strings.Center("X", 5))           // "  X  "
	fmt.Println(strings.Center("😎⚽", 4))        // " 😎⚽ "
	fmt.Println(strings.Center("ab\u0301cde", 8)) // "  ab́cde "

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

## Conversions

Return a pretty JSON representation of any interface

```go
package main

import (
	"github.com/shomali11/util/conversions"
	"fmt"
)

func main() {
    x := map[string]interface{}{"number": 1, "string": "cool", "bool": true, "float": 1.5}    
    fmt.Println(conversions.PrettyJson(x))
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

Convert any interface to a String

```go
package main

import (
	"github.com/shomali11/util/conversions"
	"fmt"
)

func main() {
    x := map[string]interface{}{"number": 1, "string": "cool", "bool": true, "float": 1.5}    
    fmt.Println(conversions.Stringify(x))
}
```

```
{"bool":true,"float":1.5,"number":1,"string":"cool"} <nil>
```

Convert any string back to its original struct

```go
package main

import (
	"github.com/shomali11/util/conversions"
	"fmt"
)

func main() {
	x := "{\"bool\":true,\"float\":1.5,\"number\":1,\"string\":\"cool\"}"
	var results map[string]interface{}
    fmt.Println(conversions.Structify(x, &results))
    fmt.Println(results)
}
```

```
<nil>
map[bool:true float:1.5 number:1 string:cool]
```
