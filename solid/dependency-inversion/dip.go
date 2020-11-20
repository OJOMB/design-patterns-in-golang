/*
Dependency Inversion Principle
High-level modules shoud not depend on low-level modules
both should instead depend on abstractions/contracts/interfaces

In this example we have a low-level module, relations.
relations exposes a type for modelling human relationships - Relationship
relations also exposes a relationship storage type - Relationships

We also have a high-level module called investigator
investigator takes a relationship storage abstraction which is the RelationshipStorer interface
this means that we can easily swap out the storage type - say we want to swtch from an in memory DS to a database
being able to swap makes our code more reusable
it also makes the code easier to test since we can create and use mock implementations




*/
package main

import (
	"fmt"

	"github.com/OJOMB/design-patterns-in-golang/solid/dependency-inversion/investigator"
	"github.com/OJOMB/design-patterns-in-golang/solid/dependency-inversion/persons"
	"github.com/OJOMB/design-patterns-in-golang/solid/dependency-inversion/relations"
	"google.golang.org/genproto/googleapis/type/date"
)

func main() {
	Ligeia := persons.Person{
		Name: "Ligeia Marsh",
		DOB: date.Date{
			Year:  1968,
			Month: 3,
			Day:   30,
		},
		Gender:      1,
		Nationality: "British",
	}
	Oscar := persons.Person{
		Name: "Oscar Oram",
		DOB: date.Date{
			Year:  1991,
			Month: 4,
			Day:   23,
		},
		Gender:      0,
		Nationality: "British",
	}
	Ruby := persons.Person{
		Name: "Ruby White",
		DOB: date.Date{
			Year:  2012,
			Month: 9,
			Day:   4,
		},
		Gender:      1,
		Nationality: "British",
	}
	Mimi := persons.Person{
		Name: "Mimi White",
		DOB: date.Date{
			Year:  2008,
			Month: 12,
			Day:   8,
		},
		Gender:      1,
		Nationality: "British",
	}

	relationships := &relations.Relationships{}

	relationships.AddParentChild(&Ligeia, &Oscar)
	relationships.AddParentChild(&Ligeia, &Ruby)
	relationships.AddParentChild(&Ligeia, &Mimi)

	inv := investigator.Investigator{Relationships: relationships}

	fmt.Printf("%s has %d children\n", Ligeia.Name, inv.NoOfChildren(&Ligeia))

}
