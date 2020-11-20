package product

import (
	"github.com/OJOMB/design-patterns-in-golang/solid/open-closed/colour"
	"github.com/OJOMB/design-patterns-in-golang/solid/open-closed/size"
	"google.golang.org/genproto/googleapis/type/date"
)

// Product inventory item
type Product struct {
	ID         string        `json:"pid"`
	Name       string        `json:"name"`
	UnitPrice  float64       `json:"unit_price"`
	LaunchDate *date.Date    `json:"launch_date"`
	Size       size.Size     `json:"size"`
	Brand      string        `json:"brand"`
	Colour     colour.Colour `json:"colour"`
}

func (p *Product) GetID() string {
	return p.ID
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetUnitPrice() float64 {
	return p.UnitPrice
}

func (p *Product) GetLaunchDate() *date.Date {
	return p.LaunchDate
}

func (p *Product) GetSize() size.Size {
	return p.Size
}

func (p *Product) GetBrand() string {
	return p.Brand
}

func (p *Product) GetColour() colour.Colour {
	return p.Colour
}
