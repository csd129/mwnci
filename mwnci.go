// Monkey is a scripting language implemented in golang, based upon
// the book "Write an Interpreter in Go", written by Thorsten Ball.
//
// This implementation adds a number of tweaks, improvements, and new
// features.  For example we support file-based I/O, regular expressions,
// the ternary operator, and more.
//
// For full details please consult the project homepage http://mwnci.ploogie.co.uk
package main

import (
	"flag"
	"fmt"
	"os"

	"mwnci/evaluator"
	"mwnci/lexer"
	"mwnci/object"
	"mwnci/parser"
	"mwnci/repl"
)

// The current version
var version = "0.1.9.r05-20240621"

var stdlib string

// Implemention of "version()" function.
func versionFun(args ...object.Object) object.Object {
	return &object.String{Value: version}
}

// Execute the supplied string as a program.
func Execute(input string) int {

	env := object.NewEnvironment()
	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		for _, msg := range p.Errors() {
			fmt.Printf("\t%s\n", msg)
		}
		os.Exit(1)
	}

	// Register a function called version()
	// that the script can call.
	evaluator.RegisterBuiltin("version",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (versionFun(args...))
		})

	//
	//  Parse and evaluate our standard-library.
	//
	stdlib := "INCLUDE={} include(\"main\")"
	initL := lexer.New(stdlib)
	initP := parser.New(initL)
	initProg := initP.ParseProgram()
	evaluator.Eval(initProg, env)

	//
	//  Now evaluate the code the user wanted to load.
	//
	//  Note that here our environment will still contain
	// the code we just loaded from our data-resource
	//
	//  (i.e. Our monkey-based standard library.)
	//
	evaluator.Eval(program, env)
	return 0
}

func main() {

	//
	// Setup some flags.
	//
	eval := flag.String("eval", "", "Code to execute.")
	vers := flag.Bool("version", false, "Show our version and exit.")

	//
	// Parse the flags
	//
	flag.Parse()

	//
	// Showing the version?
	//
	if *vers {
		fmt.Printf("mwnci %s\n", version)
		os.Exit(1)
	}

	//
	// Executing code?
	//
	if *eval != "" {
		Execute(*eval)
		os.Exit(1)
	}

	//
	// Otherwise we're either reading from STDIN, or the
	// named file containing source-code.
	//
	var input []byte
	var err error

	if len(flag.Args()) > 0 {
		input, err = os.ReadFile(os.Args[1])
	} else {
		fmt.Printf("Mwnci %s\n", version)
		repl.Start(os.Stdin, os.Stdout)
	}

	if err != nil {
		fmt.Printf("Error reading: %s\n", err.Error())
	}

	Execute(string(input))
}
