/*
Open-Closed Principle (OCP)
Open for extension / Closed for modification

demonstrated via the Specification Pattern
*/
package main

import (
	"fmt"

	"github.com/OJOMB/design-patterns-in-golang/solid/open-closed/colour"
	"github.com/OJOMB/design-patterns-in-golang/solid/open-closed/filter"
	"github.com/OJOMB/design-patterns-in-golang/solid/open-closed/product"
	"github.com/OJOMB/design-patterns-in-golang/solid/open-closed/size"
	"github.com/OJOMB/design-patterns-in-golang/solid/open-closed/specification"
	"google.golang.org/genproto/googleapis/type/date"
)

var inventory = []product.Product{
	{
		ID:         "0",
		Name:       "Air Max One",
		UnitPrice:  130.,
		LaunchDate: &date.Date{Day: 1, Month: 12, Year: 1996},
		Brand:      "Nike",
		Colour:     colour.Black,
		Size:       size.Small,
	},
	{
		ID:         "1",
		Name:       "Air Max Two",
		UnitPrice:  130.,
		LaunchDate: &date.Date{Day: 1, Month: 12, Year: 1998},
		Brand:      "Nike",
		Colour:     colour.Black,
		Size:       size.Small,
	},
	{
		ID:         "1",
		Name:       "Green Flash",
		UnitPrice:  40.,
		LaunchDate: &date.Date{Day: 1, Month: 12, Year: 1982},
		Brand:      "Dunlop",
		Colour:     colour.White,
		Size:       size.Large,
	},
	{
		ID:         "1",
		Name:       "silver shadows",
		UnitPrice:  40.,
		LaunchDate: &date.Date{Day: 1, Month: 12, Year: 1982},
		Brand:      "hi-Tec",
		Colour:     colour.White,
		Size:       size.Small,
	},
}

func main() {
	badFilterType := &filter.Filter{Inventory: inventory}
	specFilterType := &filter.SpecFilter{Inventory: inventory}

	fmt.Println("___________________________________________")
	fmt.Println("|++++++++++++++=========++++++++++++++++++|")
	fmt.Println("|++++++++++++++BadFilter++++++++++++++++++|")
	fmt.Println("|++++++++++++++=========++++++++++++++++++|")
	fmt.Println("___________________________________________")
	whites := badFilterType.ByColour(colour.White)
	fmt.Println("White products:")
	for _, t := range whites {
		fmt.Printf("--- %v\n", t)
	}

	smalls := badFilterType.BySize(size.Small)
	fmt.Println("Small products:")
	for _, t := range smalls {
		fmt.Printf("--- %v\n", t)
	}

	smallsWhites := badFilterType.BySizeAndColour(size.Small, colour.White)
	fmt.Println("Small white products:")
	for _, t := range smallsWhites {
		fmt.Printf("--- %v\n", t)
	}

	fmt.Println("____________________________________________")
	fmt.Println("|++++++++++++++==========++++++++++++++++++|")
	fmt.Println("|++++++++++++++SpecFilter++++++++++++++++++|")
	fmt.Println("|++++++++++++++==========++++++++++++++++++|")
	fmt.Println("____________________________________________")

	// create a spec with colour criteria set to white
	whiteSpec := &specification.Colour{Colour: colour.White}
	whites = specFilterType.Filter(whiteSpec)
	fmt.Println("White products:")
	for _, t := range whites {
		fmt.Printf("--- %v\n", t)
	}

	// create a spec with size criteria set to small
	smallSpec := &specification.Size{Size: size.Small}
	smalls = specFilterType.Filter(smallSpec)
	fmt.Println("Small products:")
	for _, t := range whites {
		fmt.Printf("--- %v\n", t)
	}

	smallWhiteSpec := &specification.And{
		First:  smallSpec,
		Second: whiteSpec,
	}
	smallsWhites = specFilterType.Filter(smallWhiteSpec)
	fmt.Println("Small white products:")
	for _, t := range smallsWhites {
		fmt.Printf("--- %v\n", t)
	}
}
