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
	"github.com/shomali11/util/xconcurrency"
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
    
    xconcurrency.Parallelize(func1, func2)  // a 1 b 2 c 3
    
    xconcurrency.ParallelizeTimeout(time.Minute, func1, func2)  // a 1 b 2 c 3
}
```

## Hashes

```go
package main

import (
	"github.com/shomali11/util/xhashes"
	"fmt"
)

func main() {
	fmt.Println(xhashes.FNV32("Raed Shomali"))  // 424383802
	fmt.Println(xhashes.FNV32a("Raed Shomali")) // 3711360776
	fmt.Println(xhashes.FNV64("Raed Shomali"))  // 13852322269024953050
	fmt.Println(xhashes.FNV64a("Raed Shomali")) // 17869303103005344072
	fmt.Println(xhashes.MD5("Raed Shomali"))    // c313bc3b48fcfed9abc733429665b105
	fmt.Println(xhashes.SHA1("Raed Shomali"))   // e0d66f6f09de72942e83289cc994b3c721ab34c5
	fmt.Println(xhashes.SHA256("Raed Shomali")) // 75894b9be21065a833e57bfe4440b375fc216f120a965243c9be8b2dc36709c2
	fmt.Println(xhashes.SHA512("Raed Shomali")) // 406e8d495140187a8b09893c30d054cf385ad7359855db0d2e0386c7189ac1c4667a4816d1b63a19f3d8ccdcbace7861ec4cc6ff5e2a1659c8f4360bda699b42
}
```

## Compressions

```go
package main

import (
	"github.com/shomali11/util/xcompressions"
	"fmt"
)

func main() {
	fmt.Println(xcompressions.Compress([]byte("Raed Shomali")))
}
```

## Encodings

```go
package main

import (
	"github.com/shomali11/util/xencodings"
	"fmt"
)

func main() {
	fmt.Println(xencodings.Base32Encode([]byte("Raed Shomali")))
	fmt.Println(xencodings.Base64Encode([]byte("Raed Shomali")))
}
```

## Strings

```go
package main

import (
	"github.com/shomali11/util/xstrings"
	"fmt"
)

func main() {
	fmt.Println(xstrings.IsEmpty(""))     // true
	fmt.Println(xstrings.IsEmpty("text")) // false
	fmt.Println(xstrings.IsEmpty("	"))  // false

	fmt.Println(xstrings.IsNotEmpty(""))     // false
	fmt.Println(xstrings.IsNotEmpty("text")) // true
	fmt.Println(xstrings.IsNotEmpty("	")) // true

	fmt.Println(xstrings.IsBlank(""))     // true
	fmt.Println(xstrings.IsBlank("	"))  // true
	fmt.Println(xstrings.IsBlank("text")) // false

	fmt.Println(xstrings.IsNotBlank(""))     // false
	fmt.Println(xstrings.IsNotBlank("	")) // false
	fmt.Println(xstrings.IsNotBlank("text")) // true

	fmt.Println(xstrings.Left("", 5))            // "     "
	fmt.Println(xstrings.Left("X", 5))           // "X    "
	fmt.Println(xstrings.Left("😎⚽", 4))        // "😎⚽  "
	fmt.Println(xstrings.Left("ab\u0301cde", 8)) // "ab́cde   "

	fmt.Println(xstrings.Right("", 5))            // "     "
	fmt.Println(xstrings.Right("X", 5))           // "    X"
	fmt.Println(xstrings.Right("😎⚽", 4))        // "  😎⚽"
	fmt.Println(xstrings.Right("ab\u0301cde", 8)) // "   ab́cde"

	fmt.Println(xstrings.Center("", 5))            // "     "
	fmt.Println(xstrings.Center("X", 5))           // "  X  "
	fmt.Println(xstrings.Center("😎⚽", 4))        // " 😎⚽ "
	fmt.Println(xstrings.Center("ab\u0301cde", 8)) // "  ab́cde "

	fmt.Println(xstrings.Length(""))                                          // 0
	fmt.Println(xstrings.Length("X"))                                         // 1
	fmt.Println(xstrings.Length("b\u0301"))                                   // 1
	fmt.Println(xstrings.Length("😎⚽"))                                      // 2
	fmt.Println(xstrings.Length("Les Mise\u0301rables"))                      // 14
	fmt.Println(xstrings.Length("ab\u0301cde"))                               // 5
	fmt.Println(xstrings.Length("This `\xc5` is an invalid UTF8 character"))  // 37
	fmt.Println(xstrings.Length("The quick bròwn 狐 jumped over the lazy 犬")) // 40

	fmt.Println(xstrings.Reverse(""))                                            // ""
	fmt.Println(xstrings.Reverse("X"))                                           // "X"
	fmt.Println(xstrings.Reverse("😎⚽"))                                        // "⚽😎"
	fmt.Println(xstrings.Reverse("Les Mise\u0301rables"))                        // "selbare\u0301siM seL"
	fmt.Println(xstrings.Reverse("This `\xc5` is an invalid UTF8 character"))    // "retcarahc 8FTU dilavni na si `�` sihT"
	fmt.Println(xstrings.Reverse("The quick bròwn 狐 jumped over the lazy 犬"))  // "犬 yzal eht revo depmuj 狐 nwòrb kciuq ehT"
}
```

## Conditions

```go
package main

import (
	"github.com/shomali11/util/xconditions"
	"fmt"
)

func main() {
	fmt.Println(xconditions.IfThen(1 == 1, "Yes")) // "Yes"
	fmt.Println(xconditions.IfThen(1 != 1, "Woo")) // nil
	fmt.Println(xconditions.IfThen(1 < 2, "Less")) // "Less"

	fmt.Println(xconditions.IfThenElse(1 == 1, "Yes", false)) // "Yes"
	fmt.Println(xconditions.IfThenElse(1 != 1, nil, 1))       // 1
	fmt.Println(xconditions.IfThenElse(1 < 2, nil, "No"))     // nil

	fmt.Println(xconditions.DefaultIfNil(nil, nil))  // nil
	fmt.Println(xconditions.DefaultIfNil(nil, ""))   // ""
	fmt.Println(xconditions.DefaultIfNil("A", "B"))  // "A"
	fmt.Println(xconditions.DefaultIfNil(true, "B")) // true
	fmt.Println(xconditions.DefaultIfNil(1, false))  // 1

	fmt.Println(xconditions.FirstNonNil(nil, nil))                // nil
	fmt.Println(xconditions.FirstNonNil(nil, ""))                 // ""
	fmt.Println(xconditions.FirstNonNil("A", "B"))                // "A"
	fmt.Println(xconditions.FirstNonNil(true, "B"))               // true
	fmt.Println(xconditions.FirstNonNil(1, false))                // 1
	fmt.Println(xconditions.FirstNonNil(nil, nil, nil, 10))       // 10
	fmt.Println(xconditions.FirstNonNil(nil, nil, nil, nil, nil)) // nil
	fmt.Println(xconditions.FirstNonNil())                        // nil
}
```

## Errors

```go
package main

import (
	"github.com/shomali11/util/xerrors"
	"fmt"
)

func main() {
	fmt.Println(xerrors.DefaultErrorIfNil(nil, "Cool"))                // "Cool"
	fmt.Println(xerrors.DefaultErrorIfNil(errors.New("Oops"), "Cool")) // "Oops"
}
```

## Manipulations

```go
package main

import (
	"github.com/shomali11/util/xmanipulations"
	"math/rand"
	"time"
	"fmt"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())

	array := []interface{}{"a", "b", "c"}
	xmanipulations.Shuffle(array, source)

	fmt.Println(array) // [c b a]
}
```

## Calculations

Given two integers representing the numerator and denominator of a fraction, return the fraction in string format with the repeating part enclosed in parentheses

```go
package main

import (
	"github.com/shomali11/util/xcalculations"
	"fmt"
)

func main() {
    fmt.Println(xcalculations.Divide(0, 0))     // "ERROR"
    fmt.Println(xcalculations.Divide(1, 2))     // "0.5(0)"
    fmt.Println(xcalculations.Divide(0, 3))     // "0.(0)"
    fmt.Println(xcalculations.Divide(10, 3))    // "3.(3)"
    fmt.Println(xcalculations.Divide(22, 7))    // "3.(142857)"
    fmt.Println(xcalculations.Divide(100, 145)) // "0.(6896551724137931034482758620)"
}
```

## Conversions

Return a pretty JSON representation of any interface

```go
package main

import (
	"github.com/shomali11/util/xconversions"
	"fmt"
)

func main() {
    x := map[string]interface{}{"number": 1, "string": "cool", "bool": true, "float": 1.5}    
    fmt.Println(xconversions.PrettyJson(x))
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
	"github.com/shomali11/util/xconversions"
	"fmt"
)

func main() {
    x := map[string]interface{}{"number": 1, "string": "cool", "bool": true, "float": 1.5}    
    fmt.Println(xconversions.Stringify(x))
}
```

```
{"bool":true,"float":1.5,"number":1,"string":"cool"} <nil>
```

Convert any string back to its original struct

```go
package main

import (
	"github.com/shomali11/util/xconversions"
	"fmt"
)

func main() {
	x := "{\"bool\":true,\"float\":1.5,\"number\":1,\"string\":\"cool\"}"
	var results map[string]interface{}
    fmt.Println(xconversions.Structify(x, &results))
    fmt.Println(results)
}
```

```
<nil>
map[bool:true float:1.5 number:1 string:cool]
```
