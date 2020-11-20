/*
Journal is a collection of written texts modelled as Entries
Journal's responsibility is to manage a collection of Entries and nothing else
the commented out methods on Journal violate the Single-Responsibility Principle


*/
package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"google.golang.org/genproto/googleapis/type/date"
)

// CollectionOfTextsGetter defines the method set for a type that has a collection of texts that are gettable
type CollectionOfTextsGetter interface {
	GetTexts() []string
}

// CollectionOfTextsPersister abstracts away the logic for persisting our objects from the objects themselves.
// this is so that the objects comply with the Single-Responsibility Principle
// i.e. they are responsible for their primary function and not their primary function + persisitence
type CollectionOfTextsPersister struct {
	lineSeperator string
}

// SaveToFile saves the entries of a given object to a file
func (p *CollectionOfTextsPersister) SaveToFile(filename string, object CollectionOfTextsGetter) error {
	return ioutil.WriteFile(
		filename,
		[]byte(strings.Join(object.GetTexts(), p.lineSeperator)),
		0644,
	)
}

// Entry models a written and dated text
type Entry struct {
	Date    *date.Date
	Text    string
	Private bool
}

// NewEntry returns a pointer to a new Entry
func NewEntry(text string, d *date.Date) *Entry {
	return &Entry{Date: d, Text: text}
}

// Journal models a collection of entries
type Journal struct {
	Owner   string            `json:"owner"`
	Entries map[string]*Entry `json:"entries"`
}

// NewJournal returns a pointer to a new Journal
func NewJournal(owner string, entries ...*Entry) *Journal {
	m := make(map[string]*Entry)
	for _, e := range entries {
		m[e.Date.String()] = e
	}
	return &Journal{
		Owner:   owner,
		Entries: m,
	}
}

func (j *Journal) String() string {
	return strings.Join(j.GetTexts(), "\n")
}

// AddEntry adds a new entry to the journal
func (j *Journal) AddEntry(entry *Entry) {
	j.Entries[entry.Date.String()] = entry
}

// DeleteEntry removes the entry from the journal at the given date
func (j *Journal) DeleteEntry(d *date.Date) error {
	if _, ok := j.Entries[d.String()]; !ok {
		return fmt.Errorf("No entry for at given date: %s", d.String())
	}
	delete(j.Entries, d.String())
	return nil
}

//////////////////////
// BAD METHOD ALERT //
//       |		 	//
//       ▼          //
//////////////////////

// SaveToFile extends the Journal types responsibilities beyond just managing entries
// Now Journal is responsible for persistence of a collection of entries on top of management of that collection
func (j *Journal) SaveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(j.String()), 0644)
}

// GetTexts collates all the entries in the journal to a slice of texts
func (j *Journal) GetTexts() (texts []string) {
	for _, entry := range j.Entries {
		texts = append(texts, entry.Text)
	}
	return texts
}
