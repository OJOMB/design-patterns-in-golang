package main

import "fmt"

type Mac struct {
	printer printer
}

func (m *Mac) print() {
	fmt.Println("Print request for Mac")
	m.printer.printFile()
}

func (m *Mac) setPrinter(p printer) {
	m.printer = p
}
