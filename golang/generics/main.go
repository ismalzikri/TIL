package main

import "fmt"

func Sum[V int](numbers []V) V {
	var total V
	for _, e := range numbers {
		total += e
	}
	return total
}

func main() {
	total1 := Sum([]int{1, 2, 3, 4, 5})
	fmt.Println("total:", total1)
}
