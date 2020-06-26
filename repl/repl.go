package repl

import (
	"io"
	"os"

	"github.com/c-bata/go-prompt"
	"github.com/shric/monkey/evaluator"
	"github.com/shric/monkey/lexer"
	"github.com/shric/monkey/object"
	"github.com/shric/monkey/parser"
)

const PROMPT = ">> "

func executor(env *object.Environment) func(string) {
	return func(in string) {
		l := lexer.New(in)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(os.Stdout, p.Errors())
			return
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(os.Stdout, evaluated.Inspect())
			io.WriteString(os.Stdout, "\n")
		}
	}
}

func Start(env *object.Environment) {
	prom := prompt.New(executor(env), completer, prompt.OptionPrefix(PROMPT))
	prom.Run()
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
