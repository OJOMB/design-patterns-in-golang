package specification

import (
	"github.com/OJOMB/design-patterns-in-golang/solid/open-closed/colour"
	"github.com/OJOMB/design-patterns-in-golang/solid/open-closed/size"
	"google.golang.org/genproto/googleapis/type/date"
)

// Product interface
type Product interface {
	GetID() string
	GetName() string
	GetUnitPrice() float64
	GetLaunchDate() *date.Date
	GetSize() size.Size
	GetBrand() string
	GetColour() colour.Colour
}

// Specification pattern interface
type Specification interface {
	IsSatisfied(Product) bool
}

// Colour colour spec
type Colour struct {
	Colour colour.Colour
}

// IsSatisfied is satisfied if the product colour matches the specified colour
func (cs *Colour) IsSatisfied(p Product) bool {
	if p.GetColour() == cs.Colour {
		return true
	}
	return false
}

// Size colour spec
type Size struct {
	Size size.Size
}

// IsSatisfied is satisfied if the product size matches the specified size
func (ss *Size) IsSatisfied(p Product) bool {
	if p.GetSize() == ss.Size {
		return true
	}
	return false
}

// And combinatorial specification for logical AND
type And struct {
	First, Second Specification
}

// IsSatisfied returns true if the product size matches the specified size
func (as *And) IsSatisfied(p Product) bool {
	if as.First.IsSatisfied(p) && as.Second.IsSatisfied(p) {
		return true
	}
	return false
}
