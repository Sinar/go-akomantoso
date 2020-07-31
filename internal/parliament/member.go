package parliament

import (
	"fmt"

	"github.com/Sinar/go-akomantoso/internal/akomantoso"
)

// Detect Member of Parliament (MPs)
// Naive heuristics
// Tuan/Puan [? :?
// Dato [? :?
// Datuk [? :?
// Tun [? :?
// Tan Sri [? :?
// Menteri [? :?

func looksLikeRep(line string) (bool, string) {
	return false, ""
}

func extractDebaters(allLines []string) []akomantoso.Representative {
	var allReps []akomantoso.Representative
	allMapReps := make(map[string]akomantoso.Representative, 100)
	//  DEBUG
	fmt.Println("========= Cover Pages ====================")
	fmt.Println("NO LINES: ", len(allLines))
	//Debug allLines
	for _, line := range allLines {
		fmt.Println("\"", line, "\",")
		// If look like Reps; flag it ..
		// Need to merge and make it unique? OrderedMap? No order guarantee ..
		isRep, normalizedRep := looksLikeRep(line)
		if isRep {
			fmt.Println("FOUND: ", normalizedRep)
		}
	}
	fmt.Println("========= END ====================")
	//  Gather all the unique folsk together ..
	for _, uniqueRep := range allMapReps {
		allReps = append(allReps, uniqueRep)
	}
	return allReps
}
