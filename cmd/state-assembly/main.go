package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jaffee/commandeer/cobrafy"

	nsos "github.com/leowmjw/nonstdlib/os"

	nsflag "github.com/leowmjw/nonstdlib/flag"
)

type Config struct {
	institution string // StateAssembly, Parliament
	rawFolder   string
	dataFolder  string
	fileName    string
}

func main() {
	fmt.Println("State Assembly CLI ..")
	Participant()
}

func Participant() {
	// Get Command line flags ..
	// based on type ..
	defaultConfig := Config{
		rawFolder:  "raw/StateAssembly/Hansard/",
		dataFolder: "data/StateAssembly/Hansard/",
	}
	// DEBUG
	//spew.Dump(conf)
	//if Run(defaultConfig) != nil {
	//	os.Exit(1)
	//}
	if len(os.Args) < 2 {
		fmt.Println("expected 'participants' or 'bar' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {

	case "planit":
		//err := commandeer.Run(NewParticipantCmd(defaultConfig))
		err := cobrafy.Execute(NewParticipantCmd(defaultConfig))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "sayit":
		// Create output in SayIt format ..
		err := cobrafy.Execute(NewSayItCmd(defaultConfig))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	default:
		fmt.Println("expected 'participants' or 'sayit' subcommands")
		os.Exit(1)
	}
}

// Run is an example without commandeer library
func Run(conf Config) error {
	// Example of how to read the directoriees .. give them as choice later on?
	dr := nsos.NewDirReader(conf.rawFolder)
	files, err := dr.Read()
	if err != nil {
		panic(err)
	}
	//spew.Dump(files)
	// Try out flagSet
	var strings nsflag.Strings // typically a field in a struct
	strings = nsflag.NewStringsWithConstraint(files...)
	typeParticipants := nsflag.NewChoice("dun", "par")

	participantsCmd := flag.NewFlagSet("participants", flag.ExitOnError)
	participantsCmd.Var(
		&strings,
		"x",
		"Files available: "+strings.ValidValuesDescription(),
	)
	participantsCmd.Var(
		&typeParticipants,
		"y",
		"BIB")

	if len(os.Args) < 2 {
		fmt.Println("expected 'participants' or 'bar' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {

	case "participants":
		participantsCmd.Parse(os.Args[2:])
	case "bar":
	default:
		fmt.Println("expected 'participants' or 'bar' subcommands")
		os.Exit(1)
	}

	for _, selected := range strings.Strings {
		fmt.Println(selected)
	}
	fmt.Println("Choice is: ", typeParticipants.String())
	return nil
}
