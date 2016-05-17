package main

import (
	"github.com/kirillDanshin/gcd"
	"fmt"
)

func main() {
	fmt.Printf("gcd(1071, 462) is %d\n", gcd.Calc(1071, 462))
	fmt.Printf("gcd(1000, 42)  is %d\n", gcd.Calc(1000, 42))
	fmt.Printf("gcd(1931, 522) is %d\n", gcd.Calc(1931, 522))
}
