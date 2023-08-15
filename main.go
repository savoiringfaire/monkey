package main

import (
	"fmt"
	"io"
	"os"
	"os/user"

	"monkey/evaluator"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
	"monkey/repl"
)

func main() {

	if len(os.Args) == 2 {
		data, err := os.ReadFile(os.Args[1])
		if err != nil {
			fmt.Printf("Error opening `%s`: %s", os.Args[1], err)
		}

		l := lexer.New(string(data))
		p := parser.New(l)

		env := object.NewEnvironment()

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(os.Stderr, p.Errors())
			return
		}

		evaluated := evaluator.Eval(program, env)
		io.WriteString(os.Stdout, evaluated.Inspect())
		io.WriteString(os.Stdout, "\n")
	} else {
		user, err := user.Current()
		if err != nil {
			panic(err)
		}

		fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
		fmt.Printf("Feel free to type in commands\n")
		repl.Start(os.Stdin, os.Stdout)
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
