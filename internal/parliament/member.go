package parliament

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Sinar/go-akomantoso/internal/akomantoso"
)

func looksLikeRep(line string) (bool, string) {
	// Assumes: Has been trimmed both left + right ..
	// Take choice to remove the char '.'; rep name .
	// For cases of Dr.; is still OK Dr
	line = strings.ReplaceAll(line, ".", "")
	// As per: https://stackoverflow.com/a/42251527
	normalizedLine := strings.Join(strings.Fields(strings.ToLower(line)), " ")
	// Detect Member of Parliament (MPs)
	// Naive heuristics (NOT needed)
	// Tuan/Puan [? :?
	// Dato [? :?
	// Datuk [? :?
	// Tun [? :?
	// Tan Sri [? :?
	// Menteri [? :?
	// == BEGIN  PROCESSING ===
	// Exception; if start only with [; it NOT
	// Rule #0: If starts  with [; fail immediately!
	r0, r0err := regexp.MatchString(`^\[`, normalizedLine)
	if r0err != nil {
		panic(r0err)
	}
	if r0 {
		return false, ""
	}
	//  Rule #1: If got ]: at ending sure win!
	matched, err := regexp.MatchString(`(\]\:|\]|\:)$`, normalizedLine)
	if err != nil {
		panic(err)
	}
	if matched {
		// Clean up unwanted ':' character
		line = strings.ReplaceAll(line, ":", "")
		// Look for the last number from the left
		li := strings.LastIndexAny(line, "1234567890")
		if li != -1 {
			// Found a char index
			// DEBUG
			//fmt.Println("====> SPLIT: ", line, " pos: ", li)
			// Look for the index after the number found
			line = strings.Trim(line[li+1:], " ")
			//fmt.Println("-- AFTER ***", line)
		}
		// DEBUG
		//fmt.Println("************** MATCHED: **********", line)
		return matched, line
	}
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
		isRep, normalizedRep := looksLikeRep(strings.Trim(line, " "))
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
