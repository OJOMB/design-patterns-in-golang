package main

import "fmt"

type HP struct{}

func (p *HP) printFile() {
	fmt.Println("Printing by a HP Printer")
}
