package employee

import "fmt"

const (
	// Developer - developer type employee
	Developer = iota
	// Manager - manager type employee
	Manager
)

// Employee - worker droid
type Employee interface {
	DoWork()
	SetName(name string)
}

type employee struct {
	Name, Position string
	Salary         float64
}

func (e *employee) SetName(name string) {
	e.Name = name
}

// FullTimeEmployee - employee that works 40hrs
type FullTimeEmployee struct {
	employee
}

// PartTimeEmployee - employee that works < 40hrs
type PartTimeEmployee struct {
	employee
	WeeklyHours int
}

func (e *employee) DoWork() {
	fmt.Printf("Employee: %s does some work\n", e.Name)
}

// NewEmployee - returns an employee prototype given a defined enum value
func NewEmployee(employeeType int, fulltime bool) Employee {
	switch employeeType {
	case Developer:
		if fulltime {
			return &FullTimeEmployee{
				employee{"", "developer", 48e3},
			}
		}
		return &PartTimeEmployee{
			employee{"", "developer", 48e3}, 24,
		}
	case Manager:
		if fulltime {
			return &FullTimeEmployee{
				employee{"", "manager", 80e3},
			}
		}
		return &PartTimeEmployee{
			employee{"", "developer", 40e3}, 24,
		}
	}
	return nil
}

// FactoryGenerator - returns a factory for different employee positions
func FactoryGenerator(position string, salary float64, fulltime bool) func(string) Employee {
	if fulltime {
		return func(name string) Employee {
			return &FullTimeEmployee{
				employee{
					Name:     name,
					Position: position,
					Salary:   salary,
				},
			}
		}
	}
	return func(name string) Employee {
		return &PartTimeEmployee{
			employee{
				Name:     name,
				Position: position,
				Salary:   salary,
			},
			20,
		}
	}
}
