package akomantoso

import "fmt"

// Specification
// https://sayit.mysociety.org/about/developers
//akomaNtoso and debate
//The akomaNtoso and debate element wrap the entire document. The debate element has a name attribute to express the correct name for the document's type, for example, "hansard", "transcript", "play", or simply "debate".
//
//preface / coverPage
//The preface or coverPage element can contain block element children. Within that, you may use various inline elements to signify things such as the title, type, number, purpose, or jurisdiction of the document – SayIt currently only spots the docDate or docTitle elements.

type Preface struct{}

type AkomaNtoso struct {
	Preface Preface
	Debate  Debate
}

func ValidateAkomantosoDoc() {
	fmt.Println("INside validate ..")
}
