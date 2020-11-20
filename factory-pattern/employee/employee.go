package employee

import "fmt"

// Employee - worker droid
type Employee interface {
	DoWork()
}

type employee struct {
	Name, Position string
	Salary         float64
}

func (e *employee) DoWork() {
	fmt.Printf("Employee: %s does some work\n", e.Name)
}

// FactoryGenerator - returns a factory for different employee positions
func FactoryGenerator(position string, salary float64) func(string) Employee {
	return func(name string) Employee {
		return &employee{
			Name:     name,
			Position: position,
			Salary:   salary,
		}

	}
}
