package main

import (
	"dummy-app/models"
	"dummy-app/utils"
	"fmt"
)

type MyStr string

func (s MyStr) Length() int {
	return len(s)
}

func main() {
	pen := utils.MyProduct(models.Product{100, "Pen", 10})
	fmt.Println(pen.FormatProduct())
	pen.ApplyDiscount(10)
	fmt.Println(pen.FormatProduct())

	str := MyStr("Tempor exercitation non in exercitation nostrud sit elit qui. Consequat eu do aliquip nulla et velit incididunt quis ea pariatur. Aute nulla cupidatat eu aliquip labore dolore.")
	fmt.Println(str.Length())
}
