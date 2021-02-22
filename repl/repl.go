package repl

import (
	"../evaluator"
	"../formatter"
	"../lexer"
	"../object"
	"../parser"
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

const PROMPT = ">> "

var fileLocation = "src/"

var input_files = os.Args[1:]

func Start(in io.Reader, out io.Writer) {

	if len(input_files) < 1 {
		fmt.Println("Not detected any input files.")
	} else {
		fmt.Println("File name is: ", input_files[0], "\n")
		fileLocation = fileLocation + input_files[0]
	}

	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	if fileLocation != "src/" {
		file, err := os.Open(fileLocation)
		//fmt.Println("MapReduce - Reading from Data Source")
		if err != nil { //If there is an error then log the error
			log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file) //For every line within the txt
		unformattedText, formattedText := "", ""
		for scanner.Scan() {
			//if strings.HasSuffix(scanner.Text(), "{") {
			//} else if strings.HasSuffix(scanner.Text(), "}") {
			//} else if strings.HasSuffix(scanner.Text(), ";") {
			//	else {
			//
			//				}
			unformattedText = unformattedText + scanner.Text() + "\n"
		}
		formattedText = formatter.FormatInput(unformattedText, input_files[0])
		//formattedText = strings.Replace(unformattedText, "\n", "", -1)
		//Format the text and save it

		l := lexer.New(formattedText)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}

	} else {

		for {
			fmt.Printf(PROMPT)
			scanned := scanner.Scan()
			if !scanned {
				return
			}

			line := scanner.Text()
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
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "An error has occured!\n")
	io.WriteString(out, " - Parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
