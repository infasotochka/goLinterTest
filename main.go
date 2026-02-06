package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("жесть. Деление на нуль")
	}
	return a / b, nil
}

func main() {
	fmt.Println(divide(10, 0))
}
