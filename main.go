package main

import (
	"fmt"

	"github.com/Sinar/go-akomantoso/internal/akomantoso"
)

//go:generate /Users/leow/go/bin/xsdgen schema.xsd

func main() {
	fmt.Println("Welcome to Sinar Project go-akomantoso (go-akn) helper libs!!")
	akomantoso.ValidateAkomantosoDoc()
}
