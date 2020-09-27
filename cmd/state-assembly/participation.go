package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/davecgh/go-spew/spew"

	state_assembly "github.com/Sinar/go-akomantoso/internal/state-assembly"
)

type ParticipantCmd struct {
	ID            int    `flag:"-"`
	Conf          Config `flag:"-"`
	DebateRawFile string `help:"Where is raw?" flag:"source"`
	DebateType    string `help:"Debate Type? dun,par?"`
}

func NewParticipantCmd(conf Config) *ParticipantCmd {
	return &ParticipantCmd{Conf: conf, DebateType: "dun"}
}

func createDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}

func (m *ParticipantCmd) Run() error {
	if m.DebateRawFile == "" {
		return fmt.Errorf("Select filename plz!!")
	}
	fmt.Println("EXECUTE simethin ..")
	pdfPath := m.Conf.rawFolder + m.DebateRawFile
	// Create data folder as per needed; based on the extracted label
	dataLabel := extractLabelFromFileName(pdfPath)
	fmt.Println("Creating folder ..", m.Conf.dataFolder, dataLabel)
	createDirIfNotExist(m.Conf.dataFolder + dataLabel)
	currentDPS := GenerateDebateProcessorState(pdfPath)
	// Default is to output; but have optional Output of Plan ...
	da := state_assembly.NewDebateAnalyzer(pdfPath)
	err, mapAllReps := da.Process()
	//spew.Dump(allReps)
	if err != nil {
		panic(err)
	}
	currentDPS.RepresentativesMap = mapAllReps
	spew.Dump(currentDPS)
	// Output into YAML ..
	fmt.Println("Writing DPS into ..", dataLabel)
	return nil
}

func extractLabelFromFileName(pdfPath string) string {
	_, fileName := filepath.Split(pdfPath)
	return strings.Split(fileName, filepath.Ext(fileName))[0]
}

// GenerateDebateProcessorState creates the Replist, the header, start of session
func GenerateDebateProcessorState(pdfPath string) state_assembly.DebateProcessorState {
	currentDPS := state_assembly.DebateProcessorState{}
	// Fill in the metadata info
	currentDPS.SectionMarkers.DatePageMarker, currentDPS.SectionMarkers.SessionStartMarkerLine = state_assembly.ExtractSessionInfo(pdfPath)
	currentDPS.SectionMarkers.SessionStartMarkerLine = currentDPS.SectionMarkers.SessionStartMarkerLine/30 + 1
	// Process the rep

	// DEBUG
	//spew.Dump(currentDPS)

	return currentDPS
}

// ListActiveParticipants scans through the raw file and
// extract out the participants into YAML or Pretty-print
//func ListActiveParticipants(conf Config) error {
//	saDebateAnalyzer := state_assembly.NewDebateAnalyzer(conf.rawFolder + conf.fileName)
//	err, reps := saDebateAnalyzer.Process()
//	if err != nil {
//		return err
//	}
//	// TODO: Depending on the config; dump into stdout or YAML
//	spew.Dump(reps)
//
//	return nil
//}
