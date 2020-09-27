package main

import "os"

type PlanItCmd struct {
	ID            int    `flag:"-"`
	Conf          Config `flag:"-"`
	DebateRawFile string `help:"Where is raw?" flag:"source"`
}

func NewPlanItCmd(conf Config) *PlanItCmd {
	return &PlanItCmd{Conf: conf}
}

func (m *PlanItCmd) Run() error {
	return nil
}

// Helper funcs
func createDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}
