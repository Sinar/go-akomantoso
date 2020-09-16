package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"

	state_assembly "github.com/Sinar/go-akomantoso/internal/state-assembly"
)

type ParticipantCmd struct {
	ID            int    `flag:"-"`
	DebateRawFile string `help:"Where is raw?" flag:"file"`
	DebateType    string `help:"Debate Type? dun,par?"`
}

func NewParticipantCmd(conf Config) *ParticipantCmd { return &ParticipantCmd{DebateType: "dun"} }

func (m *ParticipantCmd) Run() error {
	if m.DebateRawFile == "" {
		return fmt.Errorf("Select filename plz!!")
	}
	fmt.Println("EXECUTE simethin ..")
	return nil
}

// ListActiveParticipants scans through the raw file and
// extract out the participants into YAML or Pretty-print
func ListActiveParticipants(conf Config) error {
	saDebateAnalyzer := state_assembly.NewDebateAnalyzer(conf.rawFolder + conf.fileName)
	err, reps := saDebateAnalyzer.Process()
	if err != nil {
		return err
	}
	// TODO: Depending on the config; dump into stdout or YAML
	spew.Dump(reps)
	return nil
}
