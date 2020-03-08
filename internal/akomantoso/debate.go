package akomantoso

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

// Debate cam be made out of many categories; here Generics may be of use ..
type Debate struct {
	QuestionAnswer QuestionAnswer
}

// QuestionAnswer will have a fixed structure; who asked, which Ministry, who answered; answer ..
type QuestionAnswer struct {
	Heading        string
	SubHeading     string
	DebateSections []DebateSection
}
