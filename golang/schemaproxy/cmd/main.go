package main

import "fmt"

const year = 31536000
const year2 = 30240000

func main() {
	a := make([]int, 0)
	a = append(a, 1)
	a = append(a, 2)
	fmt.Println(a)
}
