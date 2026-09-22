package main

import (
	"fmt"
	"os"
	"strconv"
)

func mimax(n, min, max float64) float64 {
	return (n - min)/(max-min)
}

func main() {
	val,_ := strconv.ParseFloat(os.Args[1], 64)
	min,_ := strconv.ParseFloat(os.Args[2], 64)
	max,_ := strconv.ParseFloat(os.Args[3], 64)
	as := mimax(val, min, max)
	fmt.Println(as)
}
