package state_assembly

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	akomantoso "github.com/Sinar/go-akomantoso/internal/akomantoso"
)

func removeNonASCII(line string) string {
	// https://programming-idioms.org/idiom/147/remove-all-non-ascii-characters/1848/go
	return strings.Map(func(r rune) rune {
		if r > unicode.MaxASCII {
			return -1
		}
		return r
	}, line)
}

func generateRepresentativeID(line string) string {
	// TODO: To be added into test cases?
	// Replace the [] so have standardized ID
	line = strings.ReplaceAll(line, "[", "")
	line = strings.ReplaceAll(line, "]", "")
	// Replace the () so have standardized ID
	line = strings.ReplaceAll(line, "(", "")
	line = strings.ReplaceAll(line, ")", "")
	// Remove common chars like @ or / or ; ..
	line = strings.ReplaceAll(line, "@", "")
	line = strings.ReplaceAll(line, "/", "")
	line = strings.ReplaceAll(line, ";", "")
	return strings.ToLower(strings.ReplaceAll(line, " ", "-"))
}

func fastLikeRep(line string) (bool, string) {
	// Quick fast rule
	if strings.Count(line, " ") < 6 {
		// DEBUG
		//fmt.Println("FAST: ", line)
		// Step 0: Remove ALL non-ASCII
		line = removeNonASCII(line)
		// DEBUG
		//fmt.Println("RAW: ", line)
		// Take choice to remove the char '.'; rep name .
		// For cases of Dr.; is still OK Dr
		line = strings.ReplaceAll(line, ".", "")
		// As per: https://stackoverflow.com/a/42251527
		normalizedLine := strings.Join(strings.Fields(strings.ToLower(line)), " ")
		// Short-circuit; bad HACK!
		matchSC, scerr := regexp.MatchString(`kampung|persekutuan|unfortunately|kepada`, normalizedLine)
		if scerr != nil {
			panic(scerr)
		}
		if matchSC {
			return false, ""
		}
		// If YB/YAB must match the start
		// If speaker must end with speaker: tuan speaker, puan timbalan speaker; nothign else!
		if strings.Contains(normalizedLine, "speaker") {
			matchSpeaker, err := regexp.MatchString(`^tuan speaker$|^tuan timbalan speaker$|^puan speaker$|^puan timbalan speaker$`, normalizedLine)
			if err != nil {
				panic(err)
			}
			if matchSpeaker {
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
				// Step -1: Trim both left + right ..
				line = strings.Trim(line, " ")
				// Remove all extra space which will fail the has mapping
				// https://programming-idioms.org/idiom/219/replace-multiple-spaces-with-single-space
				whitespaces := regexp.MustCompile(`\s+`)
				line = whitespaces.ReplaceAllString(line, " ")
				// If still got content after all processing; say it is OK
				if line != "" {
					// DEBUG
					//fmt.Println("************** RETURNED: **********", line, "*****")
					return true, line
				}
			}
			// DEFAULT action; no MATCH!
			return false, ""
		}
		matchYB, err := regexp.MatchString(`^yb|^yab`, normalizedLine)
		if err != nil {
			panic(err)
		}
		matchOthers, err := regexp.MatchString(`tuan|puan|dato|datuk|tun|tan sri`, normalizedLine)
		if err != nil {
			panic(err)
		}
		if matchYB || matchOthers {
			// DEBUG
			//fmt.Println("MATCHED FAST: ", line)
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
			// Step -1: Trim both left + right ..
			line = strings.Trim(line, " ")
			// Remove all extra space which will fail the has mapping
			// https://programming-idioms.org/idiom/219/replace-multiple-spaces-with-single-space
			whitespaces := regexp.MustCompile(`\s+`)
			line = whitespaces.ReplaceAllString(line, " ")
			// If still got content after all processing; say it is OK
			if line != "" {
				// DEBUG
				//fmt.Println("************** RETURNED: **********", line, "*****")
				return true, line
			}
		}
	}
	// DEFAULT action; no MATCH!
	return false, ""
}

func looksLikeRep(line string) (bool, string) {
	// Step 0: Remove ALL non-ASCII
	line = removeNonASCII(line)
	// DEBUG
	//fmt.Println("RAW: ", line)
	// Take choice to remove the char '.'; rep name .
	// For cases of Dr.; is still OK Dr
	line = strings.ReplaceAll(line, ".", "")
	// As per: https://stackoverflow.com/a/42251527
	normalizedLine := strings.Join(strings.Fields(strings.ToLower(line)), " ")
	// Detect Member of Parliament (MPs)
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
		// DEBUG
		//fmt.Println("MATCHED: ", line)
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
		// Step -1: Trim both left + right ..
		line = strings.Trim(line, " ")
		// Remove all extra space which will fail the has mapping
		// https://programming-idioms.org/idiom/219/replace-multiple-spaces-with-single-space
		whitespaces := regexp.MustCompile(`\s+`)
		line = whitespaces.ReplaceAllString(line, " ")
		// If still got content after all processing; say it is OK
		if line != "" {
			// DEBUG
			//fmt.Println("************** RETURNED: **********", line, "*****")
			return matched, line
		}
	}
	// DEFAULT action; no MATCH!
	return false, ""
}

func cleanExtractedDebaters(normalizedReps []string) []akomantoso.Representative {
	// Dedupe MPs
	// Attach roles, with name etc ..
	// Search short cut

	return []akomantoso.Representative{}
}

func isRepTitle(line string) bool {
	// Naive heuristics
	line = strings.ToLower(line)
	// Short-circuit; bad HACK!
	matchSC, scerr := regexp.MatchString(`jawapan|mempengerusikan|unfortunately|kepada`, line)
	if scerr != nil {
		panic(scerr)
	}
	if matchSC {
		return false
	}
	// Representative names definitely will NOT be more than 10 words!
	if strings.Count(line, " ") > 10 {
		return false
	}
	// MAIN Rule Match
	// YB,YAB
	// Tuan
	// Puan
	// Dato
	// Datuk
	// Tun
	// Tan Sri
	// Menteri
	matchYB, err := regexp.MatchString(`yb|yab|tuan|puan|dato|datuk|tun|tan sri`, line)
	if err != nil {
		panic(err)
	}

	return matchYB
}

func extractDebaters(allLines []string) []string {
	var allReps []string
	allMapReps := make(map[string]bool, 100)
	//  DEBUG
	fmt.Println("========= Cover Pages ====================")
	fmt.Println("NO LINES: ", len(allLines))
	//Debug allLines
	for _, line := range allLines {
		// DEBUG
		//fmt.Println("\"", line, "\",")
		// If look like Reps; flag it ..
		isFastRep, normalizedRep := fastLikeRep(strings.Trim(line, " "))
		if isFastRep {
			if allMapReps[normalizedRep] {
				continue
			}
			// Unique new; set the map to seen
			allMapReps[normalizedRep] = true
			// New one, attach it for use; unsorted?
			allReps = append(allReps, normalizedRep)
			continue
		}
		// Need to merge and make it unique? OrderedMap? No order guarantee ..
		isRep, normalizedRep := looksLikeRep(strings.Trim(line, " "))
		if isRep {
			//  DEBUG
			//fmt.Println(fmt.Sprintf("IN: \"%s\"", line))
			//fmt.Println(fmt.Sprintf("OUT: \"%s\"", normalizedRep))
			// Skip if it is NOT Rep pattern
			if !isRepTitle(normalizedRep) {
				continue
			}

			if allMapReps[normalizedRep] {
				continue
			}
			// Unique new; set the map to seen
			allMapReps[normalizedRep] = true
			// New one, attach it for use; unsorted?
			allReps = append(allReps, normalizedRep)
		}
	}
	fmt.Println("========= END ====================")
	// For creating test cases
	//for _, uniqueRep := range allReps {
	//	fmt.Println(fmt.Sprintf("\"%s\":\"%s\",", uniqueRep, generateRepresentativeID(uniqueRep)))
	//}
	// DEBUG
	//spew.Dump(allReps)
	// TODO: Persist it into rep file for next phase of processing
	return allReps
}
