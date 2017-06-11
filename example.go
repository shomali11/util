package util

import (
	"fmt"
	"github.com/shomali11/util/calculations"
	"github.com/shomali11/util/concurrency"
	"github.com/shomali11/util/conditions"
	"github.com/shomali11/util/json"
	"github.com/shomali11/util/manipulations"
	"github.com/shomali11/util/strings"
	"math/rand"
	"time"
)

func Concurrency() {
	concurrency.Parallelize(Strings, Conditions, Calculations, Manipulations, Json)
}

func Strings() {
	fmt.Println(strings.Reverse("رائد شوملي"))
}

func Conditions() {
	fmt.Println(conditions.FirstNonNil("A", "B"))
}

func Calculations() {
	fmt.Println(calculations.Divide(100, 145))
}

func Manipulations() {
	source := rand.NewSource(time.Now().UnixNano())

	array := []interface{}{"a", "b", "c"}
	manipulations.Shuffle(array, source)

	fmt.Println(array)
}

func Json() {
	x := map[string]interface{}{"number": 1, "string": "cool", "bool": true, "float": 1.5}
	fmt.Println(json.PrettyJson(x))
}
