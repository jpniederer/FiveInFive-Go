package main

import (
	"fmt"
)

func main() {
	averageTemp := getAverageTemp(10, 15, 22, 17, 7, 8)
	fmt.Println(averageTemp)
}

func getAverageTemp(temps ...int) float32 {
	sum := 0
	for _, i := range temps {
		sum += i
	}

	return float32(sum) / float32(len(temps))
}
