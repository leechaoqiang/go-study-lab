package main

import (
 	"fmt"
	"github.com/leechaoqiang/go-study-lab/greetings"
)

func main() {
	fmt.Println("main start...")
	var result = greetings.Hello("vincent");
	fmt.Println(result)
	fmt.Println("main end...")
}