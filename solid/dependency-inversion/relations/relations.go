package relations

import "github.com/OJOMB/design-patterns-in-golang/solid/dependency-inversion/persons"

// RelationshipStorer - package interface for relationship storage type
type RelationshipStorer interface {
	GetAllRelationships() []RelationshipModeller
}

// RelationshipModeller - package interface for relationship data type
type RelationshipModeller interface {
	GetRelationshipType() RelationshipType
	GetSubject() *persons.Person
	GetObject() *persons.Person
}

// RelationshipType - enum denoting the category of relationship
type RelationshipType int

const (
	// Parent - object is parent of subject
	Parent RelationshipType = iota
	// Child - object is child of subject
	Child
	// Sibling - object is sibling of subject
	Sibling
	// Partner - object is partner of subject
	Partner
	// Friend - object is friend of subject
	Friend
	// Enemy - object is enemy of subject
	Enemy
	// Colleague - object is colleague of subject
	Colleague
	// Neighbour - object is neighbour of subject
	Neighbour
)

// Relationship - defines a relation between two people
type Relationship struct {
	Subject          *persons.Person
	Object           *persons.Person
	RelationshipType RelationshipType
}

// NewRelationship - returns a pointer to a new relationship
func NewRelationship(subject, object *persons.Person, relationshipType RelationshipType) *Relationship {
	return &Relationship{
		Subject:          subject,
		Object:           object,
		RelationshipType: relationshipType,
	}
}

// GetRelationshipType - returns the relationship type between subject and object
func (r *Relationship) GetRelationshipType() RelationshipType {
	return r.RelationshipType
}

// GetSubject - returns the subject of the relationship
func (r *Relationship) GetSubject() *persons.Person {
	return r.Subject
}

// GetObject - returns the object of the relationship
func (r *Relationship) GetObject() *persons.Person {
	return r.Object
}

// Relationships - stores instances of Relationship
type Relationships struct {
	store []RelationshipModeller
}

// AddParentChild - adds a parent child rel
func (r *Relationships) AddParentChild(p, c *persons.Person) {
	r.store = append(r.store, NewRelationship(p, c, Parent))
}

// GetAllRelationships - returns slice of all stored relationships
func (r *Relationships) GetAllRelationships() []RelationshipModeller {
	return r.store
}
