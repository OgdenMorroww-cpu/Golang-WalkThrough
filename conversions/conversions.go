package conversions

import (
	"fmt"
	"time"
)

func Conversions() {
	var i int = 45
	var f float64 = float64(i)
	var u uint = uint(f)

	fmt.Printf("Result: %v\n", u)

	results := map[string]interface{}{
		"Matt Donovan": 42,
	}
	timeMap(results)
	fmt.Printf("Result: %v\n", results)

	s := &FakedString{"Hey my name is: wasser und brot danke"}
	printString(s)
	printString("Hello, Gophers")

}

func timeMap(y interface{}) {
	z, ok := y.(map[string]interface{})

	if ok {
		z["updated_at"] = time.Now()
	}
}

type Stranger interface {
	String() string
}

type FakedString struct {
	contents string
}

func (s *FakedString) String() string {
	return s.contents
}

func printString(value interface{}) {
	switch str := value.(type) {
	case string:
		fmt.Println(str)

	case Stranger:
		fmt.Println(str.String())
	}
}
