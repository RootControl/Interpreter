package relp

import (
	"bufio"
	"fmt"
	"io"
	lx "github.com/RootControl/Interpreter/internal/lexer"
	tk "github.com/RootControl/Interpreter/internal/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		lexer := lx.NewLexer(line)

		for token, err := lexer.NextToken(); token.Type != tk.EoF; token, err = lexer.NextToken() {
			if err != nil {
				fmt.Fprintln(out, "%v", err)
			}
			
			fmt.Fprintf(out, "%+v\n", token)
		}
	}
}
