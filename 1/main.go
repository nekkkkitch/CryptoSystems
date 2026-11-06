package main

import "fmt"

func main() {
	fmt.Println(modPow(29, 7, 33))
	x := new(int)
	y := new(int)
	fmt.Println(GCD(3, 55, x, y), *x, *y)
	fmt.Println(inverse(11, 40))
}

func modPow(num, pow, mod int) int {
	mul := num
	for range pow - 1 {
		num = (num * mul) % mod
	}
	return num
}

func GCD(first, second int, x, y *int) int {
	if first == 0 {
		*x = 0
		*y = 1
		return second
	}
	x1 := new(int)
	y1 := new(int)
	d := GCD(second%first, first, x1, y1)
	*x = *y1 - (second/first)**x1
	*y = *x1
	return d
}

func inverse(e, z int) int {
	x := new(int)
	y := new(int)
	if GCD(e, z, x, y) != 1 {
		return -1
	}
	return *x
}

func getPrimeDividers(num int) map[int]int {
	dividers := map[int]int{}
	for num > 1 {
		for i := 2; i <= num; i++ {
			if num%i == 0 {
				dividers[i] += 1
				num /= i
				break
			}
		}
	}
	return dividers
}
