package html

import (
	"fmt"
	"strings"
)

const indentSize int = 4

// Element - represents an html element
type Element struct {
	Name       string
	Attrs      map[string]string
	MiddleText string
	Children   []*Element
}

func (e *Element) String() string {
	return e.string(0)
}

func (e *Element) string(indent int) string {
	sb := strings.Builder{}
	in := strings.Repeat(" ", indent*indentSize)
	sb.WriteString(fmt.Sprintf("%s<%s", in, e.Name))

	for k, v := range e.Attrs {
		sb.WriteString(fmt.Sprintf(` %s="%s"`, k, v))
	}

	sb.WriteString(">\n")

	if len(e.MiddleText) > 0 {
		sb.WriteString(in + strings.Repeat(" ", indentSize))
		sb.WriteString(e.MiddleText + "\n")
	}

	if len(e.Children) > 0 {
		for _, child := range e.Children {
			sb.WriteString(child.string(indent + 1))
		}
	}

	sb.WriteString(fmt.Sprintf("%s</%s>\n", in, e.Name))

	return sb.String()
}

// Builder - Builds an HTML element so you don't have to
type Builder struct {
	Name string
	Root Element
}

// NewBuilder - constructor for Builder
func NewBuilder(name string, root Element) *Builder {
	return &Builder{
		Name: name,
		Root: root,
	}
}

func (b *Builder) String() string {
	return b.Root.String()
}

// AddChild - adds a child to the html element and returns *Builder to allow fluent method chaining
func (b *Builder) AddChild(name, middleText string, attr map[string]string) *Builder {
	e := Element{Name: name, MiddleText: middleText, Attrs: attr}
	b.Root.Children = append(b.Root.Children, &e)
	return b
}
