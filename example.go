package main

import (
	"fmt"
	"github.com/shomali11/util/xcalculations"
	"github.com/shomali11/util/xcompressions"
	"github.com/shomali11/util/xconcurrency"
	"github.com/shomali11/util/xconditions"
	"github.com/shomali11/util/xconversions"
	"github.com/shomali11/util/xencodings"
	"github.com/shomali11/util/xerrors"
	"github.com/shomali11/util/xhashes"
	"github.com/shomali11/util/xmanipulations"
	"github.com/shomali11/util/xstrings"
	"math/rand"
	"time"
)

func main() {
	Hashes()
	Errors()
	Strings()
	Encodings()
	Conditions()
	Conversions()
	Calculations()
	Manipulations()
	Compressions()
	Concurrency()
}

func Conversions() {
	x := map[string]interface{}{"number": 1, "string": "cool", "bool": true, "float": 1.5}
	fmt.Println(xconversions.PrettyJson(x))
	fmt.Println(xconversions.Stringify(x))

	data := "{\"bool\":true,\"float\":1.5,\"number\":1,\"string\":\"cool\"}"
	var results map[string]interface{}
	fmt.Println(xconversions.Structify(data, &results))
	fmt.Println(results)
}

func Errors() {
	fmt.Println(xerrors.DefaultErrorIfNil(nil, "Cool"))
	fmt.Println(xerrors.DefaultErrorIfNil(xerrors.New("Oops"), "Cool"))
}

func Strings() {
	fmt.Println(xstrings.IsEmpty(""))
	fmt.Println(xstrings.IsEmpty("text"))
	fmt.Println(xstrings.IsEmpty("	"))

	fmt.Println(xstrings.IsNotEmpty(""))
	fmt.Println(xstrings.IsNotEmpty("text"))
	fmt.Println(xstrings.IsNotEmpty("	"))

	fmt.Println(xstrings.IsBlank(""))
	fmt.Println(xstrings.IsBlank("	"))
	fmt.Println(xstrings.IsBlank("text"))

	fmt.Println(xstrings.IsNotBlank(""))
	fmt.Println(xstrings.IsNotBlank("	"))
	fmt.Println(xstrings.IsNotBlank("text"))

	fmt.Println(xstrings.Left("", 5))
	fmt.Println(xstrings.Left("X", 5))
	fmt.Println(xstrings.Left("😎⚽", 4))
	fmt.Println(xstrings.Left("ab\u0301cde", 8))

	fmt.Println(xstrings.Right("", 5))
	fmt.Println(xstrings.Right("X", 5))
	fmt.Println(xstrings.Right("😎⚽", 4))
	fmt.Println(xstrings.Right("ab\u0301cde", 8))

	fmt.Println(xstrings.Center("", 5))
	fmt.Println(xstrings.Center("X", 5))
	fmt.Println(xstrings.Center("😎⚽", 4))
	fmt.Println(xstrings.Center("ab\u0301cde", 8))

	fmt.Println(xstrings.Length(""))
	fmt.Println(xstrings.Length("X"))
	fmt.Println(xstrings.Length("b\u0301"))
	fmt.Println(xstrings.Length("😎⚽"))
	fmt.Println(xstrings.Length("Les Mise\u0301rables"))
	fmt.Println(xstrings.Length("ab\u0301cde"))
	fmt.Println(xstrings.Length("This `\xc5` is an invalid UTF8 character"))
	fmt.Println(xstrings.Length("The quick bròwn 狐 jumped over the lazy 犬"))

	fmt.Println(xstrings.Reverse(""))
	fmt.Println(xstrings.Reverse("X"))
	fmt.Println(xstrings.Reverse("😎⚽"))
	fmt.Println(xstrings.Reverse("Les Mise\u0301rables"))
	fmt.Println(xstrings.Reverse("This `\xc5` is an invalid UTF8 character"))
	fmt.Println(xstrings.Reverse("The quick bròwn 狐 jumped over the lazy 犬"))
	fmt.Println(xstrings.Reverse("رائد شوملي"))
}

func Conditions() {
	fmt.Println(xconditions.IfThen(1 == 1, "Yes"))
	fmt.Println(xconditions.IfThen(1 != 1, "Woo"))
	fmt.Println(xconditions.IfThen(1 < 2, "Less"))

	fmt.Println(xconditions.IfThenElse(1 == 1, "Yes", false))
	fmt.Println(xconditions.IfThenElse(1 != 1, nil, 1))
	fmt.Println(xconditions.IfThenElse(1 < 2, nil, "No"))

	fmt.Println(xconditions.DefaultIfNil(nil, nil))
	fmt.Println(xconditions.DefaultIfNil(nil, ""))
	fmt.Println(xconditions.DefaultIfNil("A", "B"))
	fmt.Println(xconditions.DefaultIfNil(true, "B"))
	fmt.Println(xconditions.DefaultIfNil(1, false))

	fmt.Println(xconditions.FirstNonNil(nil, nil))
	fmt.Println(xconditions.FirstNonNil(nil, ""))
	fmt.Println(xconditions.FirstNonNil("A", "B"))
	fmt.Println(xconditions.FirstNonNil(true, "B"))
	fmt.Println(xconditions.FirstNonNil(1, false))
	fmt.Println(xconditions.FirstNonNil(nil, nil, nil, 10))
	fmt.Println(xconditions.FirstNonNil(nil, nil, nil, nil, nil))
	fmt.Println(xconditions.FirstNonNil())
}

func Calculations() {
	fmt.Println(xcalculations.Divide(0, 0))
	fmt.Println(xcalculations.Divide(1, 2))
	fmt.Println(xcalculations.Divide(0, 3))
	fmt.Println(xcalculations.Divide(10, 3))
	fmt.Println(xcalculations.Divide(22, 7))
	fmt.Println(xcalculations.Divide(100, 145))
}

func Compressions() {
	fmt.Println(xcompressions.Compress([]byte("Raed Shomali")))
}

func Encodings() {
	fmt.Println(xencodings.Base32Encode([]byte("Raed Shomali")))
	fmt.Println(xencodings.Base64Encode([]byte("Raed Shomali")))
}

func Hashes() {
	fmt.Println(xhashes.FNV32("Raed Shomali"))
	fmt.Println(xhashes.FNV32a("Raed Shomali"))
	fmt.Println(xhashes.FNV64("Raed Shomali"))
	fmt.Println(xhashes.FNV64a("Raed Shomali"))
	fmt.Println(xhashes.MD5("Raed Shomali"))
	fmt.Println(xhashes.SHA1("Raed Shomali"))
	fmt.Println(xhashes.SHA256("Raed Shomali"))
	fmt.Println(xhashes.SHA512("Raed Shomali"))
}

func Manipulations() {
	source := rand.NewSource(time.Now().UnixNano())

	array := []interface{}{"a", "b", "c"}
	xmanipulations.Shuffle(array, source)

	fmt.Println(array)
}

func Concurrency() {
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

	xconcurrency.Parallelize(func1, func2)
}
