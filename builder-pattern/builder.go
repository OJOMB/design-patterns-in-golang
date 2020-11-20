package main

import (
	"fmt"

	"github.com/OJOMB/design-patterns-in-golang/builder-pattern/email"
	"github.com/OJOMB/design-patterns-in-golang/builder-pattern/html"
	"github.com/OJOMB/design-patterns-in-golang/builder-pattern/person"
)

// The point of using a Builder is so that we don't require users to type out mad stuff like this:
var elem html.Element = html.Element{
	Name: "html",
	Children: []*html.Element{
		{
			Name: "head",
			Children: []*html.Element{
				{
					Name:  "meta",
					Attrs: map[string]string{"charset": "utf-8"},
				},
				{
					Name:  "meta",
					Attrs: map[string]string{"name": "viewport", "content": "width=device-width, initial-scale=1"},
				},
				{
					Name:       "title",
					MiddleText: "This is the Title",
				},
			},
		},
		{
			Name: "body",
			Children: []*html.Element{
				{
					Name:  "div",
					Attrs: map[string]string{"class": "para1"},
					Children: []*html.Element{
						{
							Name:       "p",
							MiddleText: "This is the first para",
						},
					},
				},
				{
					Name:  "div",
					Attrs: map[string]string{"class": "para2"},
					Children: []*html.Element{
						{
							Name:       "p",
							MiddleText: "This is the second para",
						},
					},
				},
			},
		},
	},
}

func main() {
	fmt.Println("####################literal####################")

	fmt.Println(elem.String())
	fmt.Print("\n\n")

	fmt.Println("####################Constructed####################")

	headBuilder := html.NewBuilder("head", html.Element{})
	headBuilder.
		AddChild("meta", "", map[string]string{"name": "viewport", "content": "width=device-width, initial-scale=1"}).
		AddChild("meta", "", map[string]string{"charset": "utf-8"}).
		AddChild("title", "This is a Title", nil)

	para1 := html.NewBuilder("p", html.Element{MiddleText: "This is the first para"})

	bodyBuilder := html.NewBuilder("body", html.Element{}).
		AddChild("div", "", map[string]string{"class": "para"})

	fmt.Println("####################Person####################")
	personBuilder := person.NewBuilder()

	personBuilder.
		Works(). // returns a *person.JobBuilder
		CurrentlyEmployed(true).
		WorksAs("programmer").
		WorksFor("home24").
		Earns(5e4).
		Lives(). // returns a *person.AddressBuilder
		LivesAt("16a").
		LivesOn("Althea Street").
		InCity("London").
		InCountry("UK").
		WithPostcode("sw62ry")

	p := personBuilder.Build()

	fmt.Println("Here's a person: ", p)

	build := func(eb *email.Builder) {
		eb.To("oscar@real.com")
		eb.From("oscar@fake.com")
		eb.Subject("subject")
		eb.Body("body")
	}

	email.SendEmail(build)
}
