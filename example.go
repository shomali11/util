package main

import (
	"fmt"
	"github.com/shomali11/util/calculations"
	"github.com/shomali11/util/compressions"
	"github.com/shomali11/util/concurrency"
	"github.com/shomali11/util/conditions"
	"github.com/shomali11/util/conversions"
	"github.com/shomali11/util/errors"
	"github.com/shomali11/util/hashes"
	"github.com/shomali11/util/manipulations"
	"github.com/shomali11/util/strings"
	"math/rand"
	"time"
)

func main() {
	Hashes()
	Errors()
	Strings()
	Conditions()
	Conversions()
	Calculations()
	Manipulations()
	Compressions()
	Concurrency()
}

func Conversions() {
	x := map[string]interface{}{"number": 1, "string": "cool", "bool": true, "float": 1.5}
	fmt.Println(conversions.PrettyJson(x))
	fmt.Println(conversions.Stringify(x))

	data := "{\"bool\":true,\"float\":1.5,\"number\":1,\"string\":\"cool\"}"
	var results map[string]interface{}
	fmt.Println(conversions.Structify(data, &results))
	fmt.Println(results)
}

func Errors() {
	fmt.Println(errors.DefaultErrorIfNil(nil, "Cool"))
	fmt.Println(errors.DefaultErrorIfNil(errors.New("Oops"), "Cool"))
}

func Strings() {
	fmt.Println(strings.IsEmpty(""))
	fmt.Println(strings.IsEmpty("text"))
	fmt.Println(strings.IsEmpty("	"))

	fmt.Println(strings.IsNotEmpty(""))
	fmt.Println(strings.IsNotEmpty("text"))
	fmt.Println(strings.IsNotEmpty("	"))

	fmt.Println(strings.IsBlank(""))
	fmt.Println(strings.IsBlank("	"))
	fmt.Println(strings.IsBlank("text"))

	fmt.Println(strings.IsNotBlank(""))
	fmt.Println(strings.IsNotBlank("	"))
	fmt.Println(strings.IsNotBlank("text"))

	fmt.Println(strings.Length(""))
	fmt.Println(strings.Length("X"))
	fmt.Println(strings.Length("b\u0301"))
	fmt.Println(strings.Length("😎⚽"))
	fmt.Println(strings.Length("Les Mise\u0301rables"))
	fmt.Println(strings.Length("ab\u0301cde"))
	fmt.Println(strings.Length("This `\xc5` is an invalid UTF8 character"))
	fmt.Println(strings.Length("The quick bròwn 狐 jumped over the lazy 犬"))

	fmt.Println(strings.Reverse(""))
	fmt.Println(strings.Reverse("X"))
	fmt.Println(strings.Reverse("😎⚽"))
	fmt.Println(strings.Reverse("Les Mise\u0301rables"))
	fmt.Println(strings.Reverse("This `\xc5` is an invalid UTF8 character"))
	fmt.Println(strings.Reverse("The quick bròwn 狐 jumped over the lazy 犬"))
	fmt.Println(strings.Reverse("رائد شوملي"))
}

func Conditions() {
	fmt.Println(conditions.IfThen(1 == 1, "Yes"))
	fmt.Println(conditions.IfThen(1 != 1, "Woo"))
	fmt.Println(conditions.IfThen(1 < 2, "Less"))

	fmt.Println(conditions.IfThenElse(1 == 1, "Yes", false))
	fmt.Println(conditions.IfThenElse(1 != 1, nil, 1))
	fmt.Println(conditions.IfThenElse(1 < 2, nil, "No"))

	fmt.Println(conditions.DefaultIfNil(nil, nil))
	fmt.Println(conditions.DefaultIfNil(nil, ""))
	fmt.Println(conditions.DefaultIfNil("A", "B"))
	fmt.Println(conditions.DefaultIfNil(true, "B"))
	fmt.Println(conditions.DefaultIfNil(1, false))

	fmt.Println(conditions.FirstNonNil(nil, nil))
	fmt.Println(conditions.FirstNonNil(nil, ""))
	fmt.Println(conditions.FirstNonNil("A", "B"))
	fmt.Println(conditions.FirstNonNil(true, "B"))
	fmt.Println(conditions.FirstNonNil(1, false))
	fmt.Println(conditions.FirstNonNil(nil, nil, nil, 10))
	fmt.Println(conditions.FirstNonNil(nil, nil, nil, nil, nil))
	fmt.Println(conditions.FirstNonNil())
}

func Calculations() {
	fmt.Println(calculations.Divide(0, 0))
	fmt.Println(calculations.Divide(1, 2))
	fmt.Println(calculations.Divide(0, 3))
	fmt.Println(calculations.Divide(10, 3))
	fmt.Println(calculations.Divide(22, 7))
	fmt.Println(calculations.Divide(100, 145))
}

func Compressions() {
	fmt.Println(compressions.Compress("Raed Shomali"))
}

func Hashes() {
	fmt.Println(hashes.MD5("Raed Shomali"))
	fmt.Println(hashes.SHA1("Raed Shomali"))
	fmt.Println(hashes.SHA256("Raed Shomali"))
	fmt.Println(hashes.SHA512("Raed Shomali"))
}

func Manipulations() {
	source := rand.NewSource(time.Now().UnixNano())

	array := []interface{}{"a", "b", "c"}
	manipulations.Shuffle(array, source)

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

	concurrency.Parallelize(func1, func2)
}
