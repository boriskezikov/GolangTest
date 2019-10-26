package main

import "fmt"

func main() {
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	//fmt.Println(nextEven())
	//fmt.Println(nextEven())
	fmt.Println(factorial(4))

}

func makeEvenGenerator () func() uint{
	i:= uint(0)
	return func()(ret uint){
		ret = i
		i +=2
		return
	}
}


func average(xs [] float64) float64 {
	total := 0.0
	for _, val := range xs {
		total += val
	}
	return total / float64(len(xs))
}

func figureCounter(args ...int) (int, bool) {
	total := 0
	for i, _ := range args {
		total += i
	}
	if total > 0 {
		return total, true
	}
	return total, false
}

func factorial (x uint) uint {
	if x == 0 {return 1 }
	return x * factorial(x-1)
}
