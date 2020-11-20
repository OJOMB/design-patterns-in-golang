package main

import "fmt"

// Printer - interface with all the behaviours of a modern full-featured printer
type Printer interface {
	Print()
	Fax()
	Scan()
	GetID() string
}

/*
So we have a Printer interface that sets out the behaviours we'd expect from a modern office printer.
Note that the features go beyond a Printer's primary function. We got faxing and scanning specified too.


Below we have a concrete printer type thats all shiny and modern.
It has all the features and fully implements Printer.
So far so good.
*/

// BrandNewShinyPrinter - this is a hella modern printer with all the things you need and more
type BrandNewShinyPrinter struct {
	ID string
}

// GetID - returns the printers unique ID
func (new *BrandNewShinyPrinter) GetID() string {
	return new.ID
}

// Print - shiny new printer has printing
func (new *BrandNewShinyPrinter) Print() {
	fmt.Println("I'm printing")
}

// Fax - shiny new printer has faxing (lol)
func (new *BrandNewShinyPrinter) Fax() {
	fmt.Println("I'm printing")
}

// Scan - shiny new printer has scanning
func (new *BrandNewShinyPrinter) Scan() {
	fmt.Println("I'm printing")
}

/////////////////////////////////////////////////////////////////
// Suite of functions for working with Printers on the networ //
///////////////////////////////////////////////////////////////

// RegisterNetworkPrinter - takes a Printer and registers it to the network
func RegisterNetworkPrinter(p Printer) {
	fmt.Printf("Successfully registered Printer with ID: %s\n", p.GetID())
}

/*
The company downstairs from us goes bankrupt and firesales all it's yellow-beige printers from the 90s.
The IT guy can't help himself and buys the whole lot.
But these new printers are old school, they have fax but no scanning
*/

// OldSchoolPrinter - printers of this type are basic, they do the job but nothing more
type OldSchoolPrinter struct {
	ID string
}

// GetID - returns the printers unique ID
func (old *OldSchoolPrinter) GetID() string {
	return old.ID
}

// Print - shiny old printer has printing
func (old *OldSchoolPrinter) Print() {
	fmt.Println("I'm printing")
}

// Fax - shiny old printer has faxing (lol)
func (old *OldSchoolPrinter) Fax() {
	fmt.Println("I'm printing")
}

/*
Now we have a problem because all the old printers we bought don't implement our Printer interface.
All our functions for working with the printers on the network won't accept the new school printer type.

The issue we're having is a result of our bloated interface
Printer should only specify methods directly related to printing. Scan and Fax are extraneous.
Most of our printer network functions would work perfectly fine without requiring a Scan and Fax method.
we should have a Faxer/Scanner/PrintScanner/PrintFaxer/ScanFaxer/PrintScanFaxer interface for functions that actually do different combos.

Only include the methods that are specifically required and nothing more. that way our functions are as reusable as possible.
*/

// Scanner - lean method set for objects that scan
type Scanner interface {
	Scan()
	GetID() string
}

// ScanFaxer - lean method set for objects that scan and fax
type ScanFaxer interface {
	Scan()
	Fax()
	GetID() string
}

// Faxer - lean method set for objects that fax
type Faxer interface {
	Fax()
	GetID() string
}

// PrintScanner - lean method set for objects that print and scan
type PrintScanner interface {
	Print()
	Scan()
	GetID() string
}

// PrintFaxer - lean method set for objects that print and fax
type PrintFaxer interface {
	Print()
	Fax()
	GetID() string
}

// PrintScanFaxer - lean method set for objects that print, scan and fax
// this is the interface that our original Printer interface should have been
type PrintScanFaxer interface {
	Print()
	Fax()
	Scan()
	GetID() string
}
