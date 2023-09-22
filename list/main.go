package main

import (
	"fmt"
	list "list/storage/list"
)

func main() {
	list := list.NewList()
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)
	list.Add(6)
	list.Add(7)
	list.Add(8)
	list.Add(9)
	list.Add(10)
	list.Add(11)
	list.Add(12)
	list.Add(13)
	list.Add(14)
	list.Add(15)
	list.Add(16)
	list.Add(17)
	list.Add(18)
	list.Add(19)
	list.Add(20)
	list.Add(21)
	list.Add(22)

	fmt.Println("список:")
	list.Print()
}
