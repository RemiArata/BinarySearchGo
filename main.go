package main

import (
	"BinarySearchGo/bs"
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	var data = generateData()
	fmt.Println(data)
	var findVal int = 90

	loc, found := bs.BinarySearch(data, findVal)
	if found {
		fmt.Println("Value", findVal, "found! At Location:", loc)
	} else {
		fmt.Print("Print value not found.")
	}
}

func generateData() []int {
	rand.Seed(5)
	var d []int
	for i := 0; i < 10; i++ {
		d = append(d, rand.Intn(100))
	}
	sort.Slice(d, func(i, j int) bool {
		return d[i] < d[j]
	})
	return d
}
