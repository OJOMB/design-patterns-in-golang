package main

import "fmt"

// Totaller - types implement Totaller if we can get the total price of the node and all nested nodes.
type Totaller interface {
	GetTotalPrice() float64
}

// Product - Item for Sale
type Product struct {
	ID, Name string
	Price    float64
}

// GetTotalPrice - returns the total price of the product
func (p *Product) GetTotalPrice() float64 {
	return p.Price
}

// Box - Product container
type Box struct {
	ID string
	// Box is a Tree structure, the nodes are any object that implements Totaller
	Contents []Totaller
}

// GetTotalPrice - returns the recursively accumulated total price of all the containing Products
func (b *Box) GetTotalPrice() (total float64) {
	for _, child := range b.Contents {
		total += child.GetTotalPrice()
	}
	return
}

// ProductBuilder - for use in piecemeal box construction
type ProductBuilder struct {
	ID, Name string
	Price    float64
}

// NewProductBuilder - returns a pointer to a new ProductBuilder
func NewProductBuilder(id, name string, price float64) *ProductBuilder {
	return &ProductBuilder{
		ID: id, Name: name, Price: price,
	}
}

// Build - constructs a new product from a product builder
func (pb *ProductBuilder) Build() *Product {
	return &Product{
		ID: pb.ID, Name: pb.Name, Price: pb.Price,
	}
}

// BoxBuilder - piecemeal Box constructor
type BoxBuilder struct {
	Type     string
	ID       string
	Contents []*BoxBuilder
}

// Build - recursively builds a Box and returns it's pointer
func (bb *BoxBuilder) Build() *Box {
	var children []Totaller
	for _, child := range bb.Contents {
		children = append(children, child.Build())
	}
	return &Box{
		ID:       bb.ID,
		Contents: children,
	}
}

// NewBoxBuilder - returns a pointer to a new BoxBuilder
func NewBoxBuilder(ID string) *BoxBuilder {
	return &BoxBuilder{
		ID: ID,
	}
}

// AddChildBox - addes a child to Box contents
func (bb *BoxBuilder) AddChildBox(childID string) *BoxBuilder {
	bb.Contents = append(bb.Contents, NewBoxBuilder(childID))
	return bb
}

// AddChildProduct - addes a child to Box contents
func (bb *BoxBuilder) AddChildProduct(childID string) *BoxBuilder {
	bb.Contents = append(bb.Contents, NewBoxBuilder(childID))
	return bb
}

// AddChildToChild - adds a child to the
func (bb *BoxBuilder) AddChildBoxToChild(parentID, childID string) *BoxBuilder {
	for _, child := range bb.Contents {
		if child.Typechild.ID == parentID {
			bb.Contents = append(bb.Contents, NewBoxBuilder(childID))
			return bb
		}
		child.AddChildToChild(parentID, childID)
	}
	return bb
}

// RemoveChild - removes a child from Box contents with the given ID
func (bb *BoxBuilder) RemoveChild(childID string) *BoxBuilder {
	for i, child := range bb.Contents {
		if child.ID == childID {
			last := len(bb.Contents) - 1
			bb.Contents[i], bb.Contents[last] = bb.Contents[last], bb.Contents[i]
			bb.Contents = bb.Contents[:last]
			return bb
		}
	}
	return bb
}

func test(b *Box) {
	fmt.Println("OK")
}

func main() {
	order := &Box{
		ID: "00",
		Contents: []Totaller{
			&Box{
				ID: "01",
				Contents: []Totaller{
					&Product{
						ID: "00", Name: "Handy", Price: 100.,
					},
					&Product{
						ID: "00", Name: "Beamer", Price: 500.,
					},
				},
			},
			&Product{
				ID: "00", Name: "Kase", Price: 4.,
			},
		},
	}

	fmt.Printf("Order Total Price: %.2f\n", order.GetTotalPrice())

	bb := NewBoxBuilder("00")
	bb.AddChild()

	test(order.Contents[0].(*Box))
}
