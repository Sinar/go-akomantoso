package akomantoso

//Each of these elements contains zero or one num, heading, and subheading elements, followed by more speech section elements, or speech-like elements.
//The num, heading and subheading elements can contain inline text, and heading must have an id attribute (though examples on the official Akoma Ntoso website do not).

// DebateSection can include DebateSections!! or Speech for leaf nodes?
type DebateSection struct {
	ID             string
	Name           string
	Num            Num
	Heading        Heading
	SubHeading     SubHeading
	Speech         Speech
	DebateSections []DebateSection
}

// EXAMPLE:
// <questions id="…">
//  <debateSection id="…">
//    <heading id="…">…</heading>
//    <question by="#…">…</question>
//    <answer by="#…" as="#…">…</answer>
//    …
//  </debateSection>
//  …
//</questions>
//<ministerialStatements>
//  <heading id="…">…</heading>
//  <debateSection id="…">
//    <heading id="…">…</heading>
//    <speech by="#…" as="#…">…</speech>
//  </debateSection>
//</ministerialStatements>
