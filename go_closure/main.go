package main

import "fmt"

func numFunction() func() int {
	num := 1

	return func() int { //閉包與匿名函數 //不允許別的方式改變狀態 //golang私有變數? 小寫
		num++
		return num
	}

	// return f() int {
	// 	num++
	// };
}
func main() {
	f := numFunction()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
