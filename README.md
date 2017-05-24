# util [![GoDoc](https://godoc.org/github.com/shomali11/util?status.svg)](https://godoc.org/github.com/shomali11/util) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A group of generic useful utility functions

## Divide

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
