package constants

import "fmt"

func Constants() {
	fmt.Println(StatusOk, StatusCreated, StatusAccepted, StatusNonAuthoritativeInfo, StatusNotContent, StatusResetContent)

	Pi := 3.14
	Truth := false
	Big := 1 << 6
	Small := Big >> 61

	fmt.Printf("PI %v\n", Pi)
	fmt.Printf("Truth %v\n", Truth)
	fmt.Printf("Small %v\n", Small)
}

const (
	StatusOk                   = 200
	StatusCreated              = 201
	StatusAccepted             = 202
	StatusNonAuthoritativeInfo = 203
	StatusNotContent           = 204
	StatusResetContent         = 205
	StatusPartialContent       = 206
)
