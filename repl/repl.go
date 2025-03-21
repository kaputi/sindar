package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/kaputi/sindar/lexer"
	"github.com/kaputi/sindar/parser"
	"github.com/kaputi/sindar/token"
)

const PROMPT = ">> "

// func Start(in io.Reader, out io.Writer) {
// }

func StartLexer(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		if line == "exit" {
			return
		}

		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}

}

func StartParser(in io.Reader, out io.Writer) {

	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		if line == "exit" {
			return
		}

		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		_, err := io.WriteString(out, program.String())
		if err != nil {
			panic(err)
		}
		_, err = io.WriteString(out, "\n")
		if err != nil {
			panic(err)
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	_, err := io.WriteString(out, " parser errors:\n")
	if err != nil {
		panic(err)
	}
	for _, msg := range errors {
		_, err := io.WriteString(out, "\t"+msg+"\n")
		if err != nil {
			panic(err)
		}
	}
}
