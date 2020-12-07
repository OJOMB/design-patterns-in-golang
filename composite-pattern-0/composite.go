package main

import (
	"fmt"
	"strings"
)

// Builder - piecemeal GraphicObject constructor
type Builder struct {
	Name, Colour string
	Children     []*Builder
}

// NewGraphicObjectBuilder - GraphicObject factory function
func NewGraphicObjectBuilder(name, colour string) *Builder {
	return &Builder{
		Name:   name,
		Colour: colour,
	}
}

// Build - recursively builds the GraphicObject and all its children.
func (b *Builder) Build() *GraphicObject {
	var children []*GraphicObject
	for _, child := range b.Children {
		children = append(children, child.Build())
	}
	return &GraphicObject{
		Name:     b.Name,
		Colour:   b.Colour,
		Children: children,
	}
}

// AddChild - adds a child to the parent
func (b *Builder) AddChild(name, colour string) {
	b.Children = append(b.Children, &Builder{Name: name, Colour: colour})
}

// AddChildToChild - adds a child to a nested builder
func (b *Builder) AddChildToChild(parentName, childName, childColour string, depth int) {
	if depth > 0 {
		for _, child := range b.Children {
			child.AddChildToChild(parentName, childName, childColour, depth-1)
		}
		return
	}
	if b.Name == parentName {
		b.Children = append(b.Children, NewGraphicObjectBuilder(childName, childColour))
	} else {
		fmt.Println("wrong parent")
	}
}

// GraphicObject - object
type GraphicObject struct {
	Name, Colour string
	Children     []*GraphicObject
}

// print - print graphic object to stdout
func (gObj *GraphicObject) print(sb *strings.Builder, depth int) {
	sb.WriteString(strings.Repeat("*", depth))
	sb.WriteRune(' ')
	sb.WriteString(gObj.Name)
	if len(gObj.Colour) > 0 {
		sb.WriteString(" - ")
		sb.WriteString(gObj.Colour)
	}
	sb.WriteRune('\n')

	for _, child := range gObj.Children {
		child.print(sb, depth+1)
	}

	if depth == 0 {
		fmt.Println(sb.String())
	}
}

func main() {
	manualObj := &GraphicObject{
		Name:   "oscar",
		Colour: "blue",
		Children: []*GraphicObject{
			{
				Name:   "otto",
				Colour: "light blue",
				Children: []*GraphicObject{
					{
						Name:   "rosie",
						Colour: "orange",
						Children: []*GraphicObject{
							{
								Name:   "daisy",
								Colour: "yellow",
							},
						},
					},
					{
						Name:     "Cosmin",
						Colour:   "light green",
						Children: nil,
					},
				},
			},
		},
	}

	manualObj.print(&strings.Builder{}, 0)

	gObjBuilder := NewGraphicObjectBuilder("oscar", "blue")
	gObjBuilder.AddChild("otto", "light blue")
	gObjBuilder.AddChildToChild("otto", "rosie", "orange", 1)
	gObjBuilder.AddChildToChild("rosie", "daisy", "yellow", 2)
	gObjBuilder.AddChildToChild("otto", "cosmin", "light green", 1)
	// next statement won't work
	gObjBuilder.AddChildToChild("kent", "dwight", "braun", 4)

	builtObj := gObjBuilder.Build()
	builtObj.print(&strings.Builder{}, 0)
}
