package state_assembly

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/Sinar/go-akomantoso/internal/akomantoso"
	"gopkg.in/yaml.v2"
)

type StateAssemblyQA struct {
	StateAssemblySession StateAssemblySession
	Type                 string
	SplitPlan            SplitPlan
	QAHansard            akomantoso.QAHansard
}

func NewStateAssemblyQA(splitPlanPath string) StateAssemblyQA {
	return StateAssemblyQA{
		StateAssemblySession: StateAssemblySession{},
		Type:                 "",
		SplitPlan: SplitPlan{
			dataDir:         "",
			PlanDir:         "",
			HansardDocument: HansardDocument{},
		},
		QAHansard: akomantoso.QAHansard{},
	}
}

func (saqa StateAssemblyQA) ExtractQAHansard() error {
	return nil
}

//  Helpers ..

type HansardType int

const (
	HANSARD_INVALID HansardType = iota
	HANSARD_SPOKEN
	HANSARD_WRITTEN
	HANSARD_DEBATE
)

type HansardQuestion struct {
	QuestionNum  string
	PageNumStart int
	PageNumEnd   int
}

type HansardDocument struct {
	StateAssemblySession string
	HansardType          HansardType
	HansardQuestions     []HansardQuestion
}

type SplitPlan struct {
	dataDir         string
	PlanDir         string
	HansardDocument HansardDocument
}

func NewEmptySplitHansardDocumentPlan(absoluteDataDir, absolutePlanFile, sessionName string) *SplitPlan {
	// Assume: sourcePDFFilename stripped off; validation here??
	// Assume: dataDir and PlanDir must become absolute before passing it back? Validate?
	if !(filepath.IsAbs(absoluteDataDir) && filepath.IsAbs(absolutePlanFile)) {
		panic(fmt.Errorf("DATA: %s + PLAN: %s MUST BE ABSOLUTE!", absoluteDataDir, absolutePlanFile))
	}
	// If absolute dataDir; just take it  as is, no use for workingDir
	//absoluteDataDir := GetAbsoluteDataDir(workingDir, dataDir)
	// Extract out filename as folder for split.yml plan
	// https://stackoverflow.com/questions/13027912/trim-strings-suffix-or-extension
	//basePDFPath := filepath.Base(sourcePDFPath)
	//planFile := absoluteDataDir + fmt.Sprintf("/%s", strings.TrimSuffix(basePDFPath, filepath.Ext(basePDFPath)))
	//// DEBUG
	//fmt.Println("PLAN_PATH: ", planFile)
	// Do abs conversion here?? for PlanDir only? Is it needed; is relative good enough? Maybe ..
	//<TODO>??
	// Assemble the pieces here ..
	splitPlan := SplitPlan{
		dataDir: absoluteDataDir,
		PlanDir: absolutePlanFile,
		HansardDocument: HansardDocument{
			StateAssemblySession: sessionName,
			HansardQuestions:     []HansardQuestion{},
		},
	}
	return &splitPlan
}
func (s *SplitPlan) LoadPlan() error {
	// Load into the struct HansardDoc from the persistent storage ..
	hansardDoc := HansardDocument{}
	b, rerr := ioutil.ReadFile(s.PlanDir)
	if rerr != nil {
		return rerr
	}
	umerr := yaml.Unmarshal(b, &hansardDoc)
	if umerr != nil {
		return umerr
	}
	// attach plan
	s.HansardDocument = hansardDoc
	// All OK!
	return nil
}
