package main

import (
	"fmt"

	"github.com/OJOMB/design-patterns-in-golang/factory-pattern/employee"
)

func main() {
	fullTimeDevFactory := employee.FactoryGenerator("developer", 48e3, true)
	oscar := fullTimeDevFactory("oscar")
	fmt.Println(oscar)

	partTimeDevFactory := employee.FactoryGenerator("developer", 24e3, false)
	barry := partTimeDevFactory("barry")
	fmt.Println(barry)

	partTimeManagerGreg := employee.NewEmployee(1, true)
	partTimeManagerGreg.SetName("partTimeManagerGreg")
	fmt.Println(partTimeManagerGreg)
}
