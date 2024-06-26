//- объявите переменную A как указатель на `int`;
//- переменную B — целочисленную с произвольным значением;
//- присвойте в переменную A указатель на целочисленную переменную B и выведите на экран значение путем разыменовывания указателя A;
//- присвойте целочисленной переменной B новое произвольное значение через указатель A и выведите на экран новое значение переменной B.

package main

import (
	"fmt"
)

func main() {
	var A *int
	B := 40

	A = &B
	fmt.Println("Значение через указатель A:", *A)  // 40

	*A = 58
	fmt.Println("Значение через указатель B:", B)  // 58

	fmt.Println("Значение через указатель A:", *A) // 58
}

