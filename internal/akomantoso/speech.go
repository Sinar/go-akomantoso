package akomantoso

// Inside Speech - Questions + Answers are annotations?
//There are seven elements for holding speech-like entries:
//
//speech
//question
//answer
//scene
//narrative
//summary
//other
//speech, question and answer require a by attribute – a URI to an entry in references (probably a TLCPerson). You may also optionally include as (a URI to a reference of the role this speech is made in), to (a URI to a reference of who this speech is addressed to), and startTime and endTime (in ISO format YYYY-MM-DDThh:mm:ss).
//
//Each of these three elements contains optional num, heading, and subheading elements (as with speech section elements), an optional from element and then one or more block elements.
//
//The from element should contain the text used in the transcript for this speaker (their identifier is handled by the attributes on the speech element itself).
//
//There are three elements for descriptive entries, that can contain inline elements and text:
//
//scene (e.g. “applause”)
//narrative (e.g. “Mr X takes the Chair”)
//summary (e.g. “Question agreed to”)
//Lastly, the other element is the container for parts of a debate that are not speeches nor scene comments (e.g. lists of papers). It requires an id attribute, and contains block elements.
type Speech struct {
	ID string
}

// TODO: Map as Popolo?
type Entity string
type Role string

type SpeechType string

type SpeechLikeElement struct {
	SpeechType SpeechType // Valid are speech, question, answer
	By         Entity     // Required
	As         Role
	To         Entity
}

// Example:
//<narrative>…</narrative>
//<speech by="#caliban">
//<from>CALIBAN</from>
//<p>……</p>
//</speech>
//<narrative>Enter TRINCULO</narrative>
//<speech by="#caliban">
//<from>CALIBAN</from>
//<p>Lo, now lo!…</p>
//</speech>
//<speech by="#trinculo">
//<from>TRINCULO</from>
//<p>……</p>
//</speech>

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

type ContentBlock struct{}

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
