package repl

import (
	"io"
	"os"
	"mwnci/evaluator"
	"mwnci/lexer"
	"mwnci/object"
	"mwnci/parser"
	"fmt"
	"github.com/chzyer/readline"
)

const PROMPT = "mwnci> "

func Start(in io.Reader, out io.Writer) {
	env := object.NewEnvironment()
	homedir, _ := os.UserHomeDir()
	historyfile := fmt.Sprintf("%s/.mwnci_history", homedir)
	
	line := "INCLUDE={} include(\"main\")"
	l := lexer.New(line)
	p := parser.New(l)
	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printParserErrors(out, p.Errors())
	}
	evaluator.Eval(program, env)
	for {
		line := ""
		rl, _ := readline.NewEx(&readline.Config{
			Prompt:            PROMPT,
			HistoryFile:       historyfile,
		})
		defer rl.Close()
		line, _ = rl.Readline()

		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
