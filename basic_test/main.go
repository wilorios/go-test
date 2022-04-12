package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("sum ", sum(1, 2))
	c, err := sumString("1", "2")
	if err != nil {
		fmt.Println("sumString ", c)
	}
}

func sum(a int, b int) int {
	return a + b
}

func sumString(a string, b string) (int, error) {
	c, err := strconv.Atoi(a)
	if err != nil {
		return 0, err
	}

	d, err := strconv.Atoi(b)
	if err != nil {
		return 0, err
	}
	return c + d, nil
}
