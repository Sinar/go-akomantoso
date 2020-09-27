package main

import (
	"fmt"
	"os"

	"github.com/jaffee/commandeer/cobrafy"
)

type Config struct {
	rawFolder  string
	dataFolder string
}

func main() {
	fmt.Println("Parliament CLI ..")
	defaultConfig := Config{
		rawFolder:  "raw/StateAssembly/Hansard/",
		dataFolder: "data/StateAssembly/Hansard/",
	}
	// Commands will have subcommand: PlanIt, SayIt
	if len(os.Args) < 2 {
		fmt.Println("expected 'planit' or 'sayit' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {

	case "planit":
		err := cobrafy.Execute(NewPlanItCmd(defaultConfig))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "sayit":
		// Create output in SayIt format ..
		//err := cobrafy.Execute(NewSayItCmd(defaultConfig))
		//if err != nil {
		//	fmt.Println(err)
		//	os.Exit(1)
		//}
	default:
		fmt.Println("expected 'planit' or 'sayit' subcommands")
		os.Exit(1)
	}

}
