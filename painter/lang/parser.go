package lang

import (
	"bufio"
	"github.com/vladimirvikulin/Lab3-Software-Architecture/painter"
	"io"
	"strings"
)

// Parser уміє прочитати дані з вхідного io.Reader та повернути список операцій представлені вхідним скриптом.
type Parser struct {
}

func (p *Parser) Parse(in io.Reader, state *UIState) ([]painter.Operation, error) {
	scanner := bufio.NewScanner(in)
	var commands []string
	for scanner.Scan() {
		commandLine := scanner.Text()
		splitCommands := strings.Split(commandLine, ",")
		if len(splitCommands) == 0 {
			continue
		}
		commands = append(commands, splitCommands...)
	}

	if err := p.ProcessCommands(commands, state); err != nil {
		return nil, err
	}

	return state.GetOperations(), nil
}

func (p *Parser) ProcessCommands(commands []string, state *UIState) error {
	for _, command := range commands {
		part := strings.Fields(command)

		switch part[0] {
		case "white":
			state.WhiteBackground()

		case "green":
			state.GreenBackground()

		case "update":
			state.SetUpdateOperation()

		case "bgrect":
			err := state.SetBackgroundRect(part[1:])
			if err != nil {
				return err
			}

		case "figure":
			err := state.AddFigure(part[1:])
			if err != nil {
				return err
			}

		case "move":
			err := state.MoveFigures(part[1:])
			if err != nil {
				return err
			}

		case "reset":
			state.Reset()
		}
	}

	return nil
}
