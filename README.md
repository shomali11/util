# util [![Go Report Card](https://goreportcard.com/badge/github.com/shomali11/util)](https://goreportcard.com/report/github.com/shomali11/util) [![GoDoc](https://godoc.org/github.com/shomali11/util?status.svg)](https://godoc.org/github.com/shomali11/util) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

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
