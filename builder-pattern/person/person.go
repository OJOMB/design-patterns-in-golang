package person

// Person - models a person
type Person struct {
	// Address
	HouseNumber, Street, City, Country, Postcode string
	// Job
	Employed          bool
	Position, Company string
	Salary            float64
}

// Builder - assists with the construction of a Person type
type Builder struct {
	person *Person
}

// NewBuilder - creates a new Person Builder
func NewBuilder() *Builder {
	return &Builder{person: &Person{}}
}

// Build - builds a Person from the builder
func (b *Builder) Build() *Person {
	return b.person
}

// Lives - returns a new AddressBuilder for piecing together the Job elements
func (b *Builder) Lives() *AddressBuilder {
	return &AddressBuilder{*b}
}

// Works - returns a new JobBuilder for piecing together the Employment elements
func (b *Builder) Works() *JobBuilder {
	return &JobBuilder{*b}
}

// AddressBuilder - assists with the construction of the Address elements of a Person type
type AddressBuilder struct {
	Builder
}

// LivesAt - builds the HouseNumber piece
func (ab *AddressBuilder) LivesAt(houseNumber string) *AddressBuilder {
	ab.person.HouseNumber = houseNumber
	return ab
}

// LivesOn - builds the Street piece
func (ab *AddressBuilder) LivesOn(street string) *AddressBuilder {
	ab.person.Street = street
	return ab
}

// InCity - builds the City piece
func (ab *AddressBuilder) InCity(city string) *AddressBuilder {
	ab.person.City = city
	return ab
}

// InCountry - builds the InCountry piece
func (ab *AddressBuilder) InCountry(country string) *AddressBuilder {
	ab.person.Country = country
	return ab
}

// WithPostcode - builds the postcode piece
func (ab *AddressBuilder) WithPostcode(postcode string) *AddressBuilder {
	ab.person.Postcode = postcode
	return ab
}

// JobBuilder - assists with the construction of the Job elements of a Person type
type JobBuilder struct {
	Builder
}

// CurrentlyEmployed - builds the Employed piece
func (jb *JobBuilder) CurrentlyEmployed(employed bool) *JobBuilder {
	jb.person.Employed = employed
	return jb
}

// WorksAs - builds the position piece
func (jb *JobBuilder) WorksAs(position string) *JobBuilder {
	jb.person.Position = position
	return jb
}

// WorksFor - builds the occupation piece
func (jb *JobBuilder) WorksFor(company string) *JobBuilder {
	jb.person.Company = company
	return jb
}

// Earns - builds the occupation piece
func (jb *JobBuilder) Earns(salary float64) *JobBuilder {
	jb.person.Salary = salary
	return jb
}
