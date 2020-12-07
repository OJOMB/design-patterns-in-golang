package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type monster struct {
	Name             string
	Eyes, Legs, Arms int
	Friends          []*monster
}

func (m *monster) DeepCopySerialization() (*monster, error) {
	b := bytes.Buffer{}

	// encode struct to bytes in buffer
	e := gob.NewEncoder(&b)
	if err := e.Encode(m); err != nil {
		return nil, err
	}

	// decode bytes in buffer to struct
	new := &monster{}
	d := gob.NewDecoder(&b)
	if err := d.Decode(new); err != nil {
		return nil, err
	}
	return new, nil
}

func main() {
	m1 := &monster{
		"jack", 1, 2, 3, []*monster{},
	}
	m2 := &monster{
		"clark", 1, 2, 3, []*monster{m1},
	}

	cp, err := m2.DeepCopySerialization()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("original: %v\n", m2)
	fmt.Printf("copy    : %v\n", cp)
	fmt.Println(`
As you can see we've copied the original, but instead of just copyting the pointer value, we've copied the values in memory
and we have a new pointer to the copied object.`,
	)
}
