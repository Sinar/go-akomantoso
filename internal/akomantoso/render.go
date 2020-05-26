package akomantoso

// Rendering we'll use gomplate  ..

// renderOralQA renders the QA struct into Akomantoso template
func renderOralQA() {

	//<questions id="…">
	//<debateSection id="…">
	//<heading id="…">…</heading>
	//<question by="#…">…</question>
	//<answer by="#…" as="#…">…</answer>
	//…
	//</debateSection>
	//…
	//</questions>
}

func renderWrittenQA() {
	// Prepare session metadata .. members
	// prepare Heading; <type> + <folder name>
	// Loop through all questions
	// Format question, newline with </br> or <p>
	// Format answer
	// Stitch them all together .. write to file ..
}

//<question by="#azminali">
//	</question>
//	<answer by="#stevensim" as="">
//		<p><text here ...></p>
//		<table></table>
//
//		</answer>
