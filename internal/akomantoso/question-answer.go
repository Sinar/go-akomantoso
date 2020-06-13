package akomantoso

type Title string

type QAContent struct {
	ID       string
	Content  []string // Raw question content extracted out
	Title    Title
	QContent []string
	QBy      Representative
	AContent []string
	ABy      Representative
}

type Header string

type QAReader struct {
	ID      string
	Header  Header
	Content []QAContent
}

func NewQAContent() {
	// Loop over Hansard Raw Content (per bunch  of pages)

	// Match existing Representative or append one to the bottom

}
