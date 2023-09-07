package repl

import (
	//	"bufio"
	//	"fmt"
	"io"
	"mwnci/evaluator"
	"mwnci/lexer"
	"mwnci/object"
	"mwnci/parser"
	"github.com/chzyer/readline"
)

const PROMPT = "mwnci> "

func Start(in io.Reader, out io.Writer) {
	//	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		
		//		fmt.Printf(PROMPT)
		//		scanned := scanner.Scan()
		//		if !scanned {
		//			return
		//		}

		//		line := scanner.Text()
		line := ""
		rl, _ := readline.New(PROMPT)
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
