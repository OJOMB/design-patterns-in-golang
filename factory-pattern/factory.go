package main

import (
	"fmt"

	"github.com/OJOMB/design-patterns-in-golang/factory-pattern/employee"
)

func main() {
	devFactory := employee.FactoryGenerator("developer", 48e3)

	oscar := devFactory("oscar")

	fmt.Println(oscar)
}
