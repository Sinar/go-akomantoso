package main

import (
	"flag"
	"fmt"
	"os"

	nsos "github.com/leowmjw/nonstdlib/os"

	nsflag "github.com/leowmjw/nonstdlib/flag"
)

type Config struct {
	rawFolder  string
	dataFolder string
	fileName   string
}

func main() {
	fmt.Println("State Assembly CLI ..")
	// Get Command line flags ..
	// based on type ..
	defaultConfig := Config{
		rawFolder:  "raw/StateAssembly/Hansard/",
		dataFolder: "data/StateAssembly/Hansard/",
		fileName:   "HANSARD-15-JULAI-2020.pdf",
	}
	// Example of how to read the directoriees .. give them as choice later on?
	dr := nsos.NewDirReader(defaultConfig.rawFolder)
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
	// DEBUG
	//spew.Dump(conf)
	if Run(defaultConfig) != nil {
		os.Exit(1)
	}
}

func Run(conf Config) error {
	return nil
}
