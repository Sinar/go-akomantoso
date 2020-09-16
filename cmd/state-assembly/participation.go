package main

import "fmt"

type ParticipantCmd struct {
	DebateRawFile string `help:"Where is raw?"`
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
