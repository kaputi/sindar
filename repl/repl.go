package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/kaputi/sindar/lexer"
	"github.com/kaputi/sindar/token"
)

const PROMPT = ">> "

func StartLexer(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		// fmt.Fprintf(out, PROMPT)
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

// func Start(in io.Reader, out io.Writer) {
// }
