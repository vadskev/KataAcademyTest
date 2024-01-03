package main

import (
	"fmt"
	"kata/internal"
)

func main() {
	fmt.Println("Введите операцию:")
	line := internal.ReadConsole()
	res, _ := internal.Calc(line)

	fmt.Printf("Output: %s\n", res)
}
