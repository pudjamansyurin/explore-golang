package poker

import (
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	playerStore PlayerStore
	in          *bufio.Scanner
}

func NewCLI(store PlayerStore, in io.Reader) *CLI {
	return &CLI{
		playerStore: store,
		in:          bufio.NewScanner(in),
	}
}

func (cli *CLI) PlayPoker() {
	winner := extractWinner(cli.readLine())
	cli.playerStore.RecordWin(winner)
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}

func extractWinner(input string) string {
	return strings.Replace(input, " wins", "", 1)
}
