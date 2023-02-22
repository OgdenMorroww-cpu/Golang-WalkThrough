package exported

import (
	"fmt"
	"math"
	"net/http"
)

func Exported() {
	fmt.Printf("HTTP status OK uses code: %d\n", http.StatusOK)

	getValueOfPi := math.Pi

	fmt.Printf("The value of PI is: %v\n", getValueOfPi)
}
