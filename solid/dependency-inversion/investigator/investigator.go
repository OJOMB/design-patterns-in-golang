package investigator

import (
	"github.com/OJOMB/design-patterns-in-golang/solid/dependency-inversion/persons"
	"github.com/OJOMB/design-patterns-in-golang/solid/dependency-inversion/relations"
)

// Investigator - type for handling analysis on a set of relationships
type Investigator struct {
	Relationships relations.RelationshipStorer
}

// NoOfChildren - returns the number of children for a given Person
func (i *Investigator) NoOfChildren(parent *persons.Person) (n int) {
	for _, rel := range i.Relationships.GetAllRelationships() {
		if rel.GetSubject() == parent || rel.GetRelationshipType() == relations.Parent {
			n++
		}
	}
	return
}
