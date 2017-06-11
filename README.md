# util [![Go Report Card](https://goreportcard.com/badge/github.com/shomali11/util)](https://goreportcard.com/report/github.com/shomali11/util) [![GoDoc](https://godoc.org/github.com/shomali11/util?status.svg)](https://godoc.org/github.com/shomali11/util) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A group of generic useful utility functions

## Usage

Using `govendor` [github.com/kardianos/govendor](https://github.com/kardianos/govendor):

```
govendor fetch github.com/shomali11/util
```

# Examples

## Concurrency

```go
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

Parallelize(func1, func2)  // a 1 b 2 c 3
```

## Strings

```go
IsEmpty("")     // true
IsEmpty("text") // false
IsEmpty("	")  // false

IsNotEmpty("")     // false
IsNotEmpty("text") // true
IsNotEmpty("	") // true

IsBlank("")     // true
IsBlank("	")  // true
IsBlank("text") // false

IsNotBlank("")     // false
IsNotBlank("	") // false
IsNotBlank("text") // true

Length("")                                          // 0
Length("X")                                         // 1
Length("b\u0301")                                   // 1
Length("😎⚽")                                      // 2
Length("Les Mise\u0301rables")                      // 14
Length("ab\u0301cde")                               // 5
Length("This `\xc5` is an invalid UTF8 character")  // 37
Length("The quick bròwn 狐 jumped over the lazy 犬") // 40

Reverse("")                                            // ""
Reverse("X")                                           // "X"
Reverse("😎⚽")                                        // "⚽😎"
Reverse("Les Mise\u0301rables")                        // "selbare\u0301siM seL"
Reverse("This `\xc5` is an invalid UTF8 character")    // "retcarahc 8FTU dilavni na si `�` sihT"
Reverse("The quick bròwn 狐 jumped over the lazy 犬")  // "犬 yzal eht revo depmuj 狐 nwòrb kciuq ehT"
```

## Conditions

```go
IfThen(1 == 1, "Yes") // "Yes"
IfThen(1 != 1, "Woo") // nil
IfThen(1 < 2, "Less") // "Less"

IfThenElse(1 == 1, "Yes", false) // "Yes"
IfThenElse(1 != 1, nil, 1)       // 1
IfThenElse(1 < 2, nil, "No")     // nil

DefaultIfNil(nil, nil)  // nil
DefaultIfNil(nil, "")   // ""
DefaultIfNil("A", "B")  // "A"
DefaultIfNil(true, "B") // true
DefaultIfNil(1, false)  // 1

FirstNonNil(nil, nil)                // nil
FirstNonNil(nil, "")                 // ""
FirstNonNil("A", "B")                // "A"
FirstNonNil(true, "B")               // true
FirstNonNil(1, false)                // 1
FirstNonNil(nil, nil, nil, 10)       // 10
FirstNonNil(nil, nil, nil, nil, nil) // nil
FirstNonNil()                        // nil
```

## Errors

```go
DefaultErrorIfNil(nil, "Cool")                // "Cool"
DefaultErrorIfNil(errors.New("Oops"), "Cool") // "Oops"
```

## Manipulations

```go
source := rand.NewSource(time.Now().UnixNano())
array := []interface{}{"a", "b", "c"}

Shuffle(array, source) // [c b a]
```

## Calculations

Given two integers representing the numerator and denominator of a fraction, return the fraction in string format with the repeating part enclosed in parentheses

```go
Divide(0, 0)     // "ERROR"
Divide(1, 2)     // "0.5(0)"
Divide(0, 3)     // "0.(0)"
Divide(10, 3)    // "3.(3)"
Divide(22, 7)    // "3.(142857)"
Divide(100, 145) // "0.(6896551724137931034482758620)"
```

## PrettyJson

Pretty JSON Prints!

```go
x := map[string]interface{}{"number": 1, "string": "cool", "bool": true, "float": 1.5}
results, _ := PrettyJson(x)
fmt.Println(results)
```

```json
{
	"bool": true,
	"float": 1.5,
	"number": 1,
	"string": "cool"
}
```
