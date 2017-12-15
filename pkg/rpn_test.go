package rpn

import (
	"strings"
	"testing"
)

func TestInfixToPostfix(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{input: []string{"2"}, expected: []string{"2"}},
		{input: []string{"2", "+", "2"}, expected: []string{"2", "2", "+"}},
		{input: []string{"2", "+", "(", "4", "-", "3", ")"}, expected: []string{"2", "4", "3", "-", "+"}},
		{input: []string{"2", "+", "3", "*", "4"}, expected: []string{"3", "4", "*", "2", "+"}},
	}

	for _, test := range tests {
		rpn, err := InfixToPostfix(test.input)

		if err != nil {
			t.Fatalf("error was non-nil: %s", err.Error())
		}

		actual := strings.Join(rpn, ",")
		expected := strings.Join(test.expected, ",")
		if actual != expected {
			t.Errorf("expected=%s. got=%s.",
				expected, actual)
		}
	}
}
