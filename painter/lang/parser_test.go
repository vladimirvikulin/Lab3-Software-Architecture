package lang

import (
	"strings"
	"testing"
)

func TestParse_InvalidNumberOfArguments(t *testing.T) {
	state := NewUIState()
	parser := Parser{}
	input := "bgrect 0.1 0.1 0.9, update\n"

	_, err := parser.Parse(strings.NewReader(input), state)
	if err == nil {
		t.Fatalf("Parse() did not return an error")
	}
}

func TestParse_InvalidArguments(t *testing.T) {
	state := NewUIState()
	parser := Parser{}
	input := "move 0.1 abc, update\n"

	_, err := parser.Parse(strings.NewReader(input), state)
	if err == nil {
		t.Fatalf("Parse() did not return an error")
	}
}

func TestParse_LargeInput(t *testing.T) {
	state := NewUIState()
	parser := Parser{}
	input := strings.NewReader(strings.Repeat("figure 0.1 0.2\n", 1000))
	_, err := parser.Parse(input, state)

	if err != nil {
		t.Fatalf("Unexpected error. Expected nil, Actual: %v", err)
	}
}

func TestParse_OneCommand(t *testing.T) {
	state := NewUIState()
	parser := Parser{}
	input := "figure 0.5 0.5, update\n"
	expectedLen := 2

	result, err := parser.Parse(strings.NewReader(input), state)
	if err != nil {
		t.Fatalf("Parse() error: %s", err)
	}

	if len(result) != expectedLen {
		t.Fatalf("Incorrect parse result")
	}
}

func TestParse_MultipleCommands(t *testing.T) {
	state := NewUIState()
	parser := Parser{}
	input := "figure 0.5 0.5, move 0.1 0.1, bgrect 0.1 0.1 0.9 0.9, update\n"
	expectedLen := 4

	result, err := parser.Parse(strings.NewReader(input), state)
	if err != nil {
		t.Fatalf("Parse() error: %s", err)
	}

	if len(result) != expectedLen {
		t.Fatalf("Incorrect parse result")
	}
}

func TestParse_Background(t *testing.T) {
	state := NewUIState()
	parser := Parser{}
	input := "white, update, green, update\n"
	expectedLen := 2

	result, err := parser.Parse(strings.NewReader(input), state)
	if err != nil {
		t.Fatalf("Parse() error: %s", err)
	}

	if len(result) != expectedLen {
		t.Fatalf("Incorrect parse result")
	}
}
