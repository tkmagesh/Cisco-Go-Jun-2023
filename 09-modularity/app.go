package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/tkmagesh/Cisco-Go-Jun-2023/09-modularity/calculator"
	"github.com/tkmagesh/Cisco-Go-Jun-2023/09-modularity/calculator/utils"
)

func main() {
	color.Red("app invoked")
	fmt.Println(greet())
	fmt.Println(calculator.Add(100, 200))
	fmt.Println(calculator.Subtract(100, 200))
	fmt.Println(calculator.OpCount())
	fmt.Println(utils.IsEven(21))
}
