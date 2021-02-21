package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4}
	m := make(map[int]*int, 0)
	for k, v := range slice {
		fmt.Println(k, " ", v, " ", &v)
		m[k] = &v
	}
	for k, v := range m {
		fmt.Println(k, " ", v, " ", *v)
	}
}
