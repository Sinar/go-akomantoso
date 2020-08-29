package akomantoso

type RepresentativeID string

type Representative struct {
	ID          string
	PopItID     string
	DisplayName string
	Represents  string
	Roles       []string
}
