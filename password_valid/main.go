package main

import (
	"fmt"
	"sort"
)

func main() {
	var size int
	var str string
	fmt.Scan(&size)
	fmt.Scan(&str)
	arr := make([]string, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}
	runes := []rune(str)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	for i, value := range arr {
		runes1 := []rune(value)
		sort.Slice(runes1, func(i, j int) bool {
			return runes1[i] < runes1[j]
		})
		arr[i] = string(runes1)
	}
	for _, value := range arr {
		if value == string(runes) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
