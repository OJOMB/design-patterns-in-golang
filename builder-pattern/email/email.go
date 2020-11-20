package email

import (
	"errors"
	"fmt"
	"strings"
)

type email struct {
	from, to, subject, body string
}

func (e *email) send() error {
	if e.to == "" || e.from == "" {
		return errors.New("email fields incomplete")
	} else if !strings.Contains(e.to, "@") || !strings.Contains(e.from, "@") {
		return errors.New("invalid email address")
	}
	fmt.Printf("Sent email from: %s, to: %s\n", e.from, e.to)
	return nil
}

// Builder - builds a email with validation at each stage of the process
type Builder struct {
	email email
}

// NewBuilder - returns a new email builder
func NewBuilder() *Builder {
	return &Builder{email: email{}}
}

// From - handles the From piece
func (b *Builder) From(from string) *Builder {
	// validation
	b.email.from = from
	return b
}

// To - handles the To piece
func (b *Builder) To(to string) *Builder {
	// validation
	b.email.to = to
	return b
}

// Subject - handles the Subject piece
func (b *Builder) Subject(subject string) *Builder {
	// validation
	b.email.subject = subject
	return b
}

// Build - returns the constructed email
func (b *Builder) build() email {
	return b.email
}

// Body - handles the Body piece
func (b *Builder) Body(body string) *Builder {
	// validation
	b.email.body = body
	return b
}

type build func(*Builder)

// SendEmail - takes a build function and sends the resultant email
func SendEmail(action build) error {
	b := Builder{}
	action(&b)
	e := b.build()
	return e.send()
}
