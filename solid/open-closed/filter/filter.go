package filter

import (
	"github.com/OJOMB/design-patterns-in-golang/solid/open-closed/colour"
	"github.com/OJOMB/design-patterns-in-golang/solid/open-closed/product"
	"github.com/OJOMB/design-patterns-in-golang/solid/open-closed/size"
	"github.com/OJOMB/design-patterns-in-golang/solid/open-closed/specification"
)

/*
Filter type violates the OCP
everytime we need a new filter we have to add a new method
*/

// Filter for filtering Products
type Filter struct {
	Inventory []product.Product
}

// ByColour filters Product slice By Colour
func (f *Filter) ByColour(c colour.Colour) (result []product.Product) {
	for _, p := range f.Inventory {
		if p.Colour == c {
			result = append(result, p)
		}
	}
	return
}

// BySize filters Product slice By Colour
func (f *Filter) BySize(s size.Size) (result []product.Product) {
	for _, p := range f.Inventory {
		if p.Size == s {
			result = append(result, p)
		}
	}
	return
}

// BySizeAndColour filters Product slice By Colour and Size
func (f *Filter) BySizeAndColour(s size.Size, c colour.Colour) (result []product.Product) {
	for _, p := range f.Inventory {
		if p.Size == s && p.Colour == c {
			result = append(result, p)
		}
	}
	return
}

/*
Spec Filter has just one method: "Filter"
we abstracted the filtering logic to specifications that define our filter criteria
if the spec is satisfied by the item then it goes into the filter results
SpecFilter is defined and tested and does not require constant modification as the requirements change
*/

// SpecFilter for filtering by a specification type
type SpecFilter struct {
	Inventory []product.Product
}

// Filter filters product list with the given spec
func (sf *SpecFilter) Filter(spec specification.Specification) (result []product.Product) {
	for _, v := range sf.Inventory {
		if spec.IsSatisfied(&v) {
			result = append(result, v)
		}
	}
	return
}
