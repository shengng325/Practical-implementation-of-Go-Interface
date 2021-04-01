package main

import (
	"bufio"
	"io"
	"os"

	bt "github.com/shengng325/Practical-implementation-of-Go-Interface/booleantree"
)

func main() {
	question := "What are the things you can’t eat in the afternoon?"

	// Answer: (breakfast || dinner || supper) && !(lunch || branch)
	keyword := bt.AND(
		bt.OR(bt.Word("breakfast"), bt.Word("dinner"), bt.Word("supper")),
		bt.NOT(
			bt.OR(bt.Word("lunch"), bt.Word("brunch")),
		),
	)

	run(os.Stdin, os.Stdout, question, keyword)
}

func run(in io.Reader, out io.Writer, question string, keyword bt.Node) {
	scanner := bufio.NewScanner(in)
	io.WriteString(out, question+"\n")

	for {
		io.WriteString(out, ">> ")
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()

		input := bt.NewSentence(line)
		result := keyword.Eval(input)
		if result {
			io.WriteString(out, "Correct answer\n")
		} else {
			io.WriteString(out, "Wrong answer\n")
		}
	}
}
