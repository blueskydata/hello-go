package main

import "fmt"

func main() {
	fmt.Println("hello git 第一次提交")
	fmt.Println("新增：第二次修改代码")

	ret := add(3, 4)
	fmt.Printf("The result is %d\n", ret)

	ret = sub(3, 4)
	fmt.Printf("The result is %d\n", ret)

	fmt.Println("finishh")
}
