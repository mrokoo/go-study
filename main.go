package main

import "fmt"

func main() {
	a := []int{4, 5, 6}
	for i, v := range a {
		fmt.Println(i, v)
	}
}
