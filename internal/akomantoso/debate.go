package akomantoso

// =================================
// Block Elements
//Block elements handled by SayIt are the HTML elements:
//
//p
//ul
//ol
//table
//All these besides p require an id attribute. ul and ol contain lis as in HTML (which can optionally have a value attribue), and lis can contain p, ul, ol, or inline text.
//Other Akoma Ntoso block elements are ignored (though not their contents).
// Example:
//<speech by="#…" as="#…">
//<from>Mr Block</from>
//<p>Here is a list:</p>
//<ul id="">
//<li>First item</li>
//<li>Second item</li>
//</ul>
//<p>And here is a table:</p>
//<table id="">
//<tr> <td>A</td> <td>B</td> </tr>
//<tr> <td>A</td> <td>D</td> </tr>
//</table>
//</speech>

type ContentBlock struct{}

type Num struct {
	ContentBlock ContentBlock
}
type Heading struct {
	ID           string
	ContentBlock ContentBlock
}
type SubHeading struct {
	ContentBlock ContentBlock
}

// -- Speech sections --
// The following elements (which all require an id attribute) can be used to create a hierarchy of speech-like elements. The generic element is debateSection, which requires a name attribute to describe what type of section it is. Most of the specific elements are only useful in a Parliamentary-style debate context; do use them if applicable, but generally you may find debateSection is what you use. SayIt doesn't handle different types of section differently at present.
//
//debateSection (additionally requires a name attribute to describe the type of section)
//administrationOfOath, rollCall, prayers
//oralStatements, writtenStatements, personalStatements, ministerialStatements
//resolutions, nationalInterest
//declarationOfVote
//communication
//petitions, papers, noticesOfMotion
//questions
//address
//proceduralMotions
//pointOfOrder
//adjournment

// One of the Debate Section type is Q&A; translate to element <questions></questions>
// QuestionAnswer will have a fixed structure; who asked, which Ministry, who answered; answer ..
type QuestionAnswer struct {
	Heading        Heading
	SubHeading     SubHeading
	DebateSections []DebateSection
}
