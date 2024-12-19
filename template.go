package main

import (
	"flag"
	"fmt"
	"os"
)

// **  GLOBALS  **//
var verbose bool

const (
	exampleConst = ""
)

// **  HELPERS  **//
func verbosePrint(str string) {
	if verbose {
		fmt.Println("V >> _v " + str + " v_")
	}
}

func normalPrint(str string) {
	fmt.Println("I >> " + str)
}

// **  BEGINCD  **//
func main() {
	// FLAGTIME
	var exampleStrFlag string
	var exampleIntFlag int
	var verboseFlag bool

	flag.StringVar(&exampleStrFlag, "example", "", "full word flag, --example, second arg is default")
	flag.StringVar(&exampleStrFlag, "e", "", "single letter flag, -e, second arg is default")
	flag.IntVar(&exampleIntFlag, "integer", 1, "--integer help, second arg is default")
	flag.IntVar(&exampleIntFlag, "i", 1, "-i help, second arg is default")

	flag.BoolVar(&verboseFlag, "verbose", false, "Enable more verbose output.")
	flag.BoolVar(&verboseFlag, "v", false, "Enable more verbose output.")

	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])

		//flag.PrintDefaults()
		flagSet := flag.CommandLine
		order := [][]string{
			[]string{"", "Required:"},
			[]string{"example", "e"},
			[]string{"", "Optionals:"},
			[]string{"integer", "i"},
			[]string{"verbose", "v"}}

		for _, set := range order {
			if set[0] == "" {
				fmt.Println("\n  " + set[1])
				continue
			}
			flag := flagSet.Lookup(set[0])
			fmt.Println("    --" + set[0] + " (-" + set[1] + ")  ->  " + flag.Usage)
		}
	}

	flag.Parse()

	if verboseFlag {
		verbose = true
	} else {
		verbose = false
	}

	// CODETIME
	normalPrint("functional template!")
	verbosePrint("verbose set!")
}
