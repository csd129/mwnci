package repl

import (
	"fmt"
	"io"
	"mwnci/evaluator"
	"mwnci/lexer"
	"mwnci/object"
	"mwnci/parser"
	"os"
	"strings"

	"github.com/chzyer/readline"
)

var (
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
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
	evaluator.Eval(program, env, false)
	line = ""
	rl, _ := readline.NewEx(&readline.Config{
		Prompt:      PROMPT,
		HistoryFile: historyfile,
	})
	defer rl.Close()
	var cmds []string
	var err error
	for {
		line, err = rl.Readline()
		if err != nil {
			break
		}
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		cmds = append(cmds, line)
		if !strings.HasSuffix(line, ":") {
			continue
		}
		cmd := strings.Join(cmds, " ")
		cmd = strings.TrimSpace(cmd)
		cmd = strings.TrimSuffix(cmd, ":")
		if len(cmd) > 0 {
			rl.SaveHistory(cmd)
			l := lexer.New(cmd)
			p := parser.New(l)

			program := p.ParseProgram()
			if len(p.Errors()) != 0 {
				printParserErrors(out, p.Errors())
				cmds = cmds[:0]
				continue
			}

			evaluated := evaluator.Eval(program, env, false)
			if evaluated != NULL {
				io.WriteString(out, evaluated.Inspect())
				io.WriteString(out, "\n")
			}
		}
		cmds = cmds[:0]
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
