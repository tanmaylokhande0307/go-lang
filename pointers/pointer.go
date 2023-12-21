package main

import "fmt"

// func zeroval(ival int) {
// 	ival = 0
// }

// func zeroptr(iptr *int) {
// 	*iptr = 0
// }

// func swap(x *int, y *int) {
// 	temp := *x
// 	*x = *y
// 	*y = temp
// }

// func double(x *int) *int {
// 	result := *x * 2
// 	return &result
// }

func main() {
	// x := 1
	// y := 2
	// fmt.Println(x, y) // output: 1 2
	// swap(&x, &y)
	// fmt.Println(x, y) // output: 2 1
	// ptr := double(&x)
	// fmt.Println(*ptr) // output: 4

	// fmt.Println(x)

	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	var ptr *int = &arr[0]
	for i := 0; i < len(arr); i++ {
		fmt.Println(*ptr)
		if(i == len(arr) - 1){
			return
		}
		ptr = &arr[i+1]
	}
	fmt.Println(arr)
}
