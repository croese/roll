package rpn

import (
	"fmt"
	"unicode"

	"github.com/golang-collections/collections/queue"
	"github.com/golang-collections/collections/stack"
)

type Direction int

const (
	Left Direction = iota
	Right
)

const (
	leftParen  = "("
	rightParen = ")"
)

type OperatorDetails struct {
	Precedence    int
	Associativity Direction
}

type OperatorTable map[string]OperatorDetails

func InfixToPostfix(tokens []string) ([]string, error) {
	operators := stack.New()
	output := queue.New()

	for _, t := range tokens {
		if isOperand(t) {
			output.Enqueue(t)
		} else if t == leftParen {
			operators.Push(t)
		} else if t == rightParen {
			for operators.Len() > 0 && operators.Peek() != leftParen {
				output.Enqueue(operators.Pop())
			}

			if operators.Peek() != leftParen {
				return nil, fmt.Errorf("mismatched parentheses")
			}

			operators.Pop() // pop left paren
		} else {
			operators.Push(t)
		}
	}

	for operators.Len() > 0 {
		output.Enqueue(operators.Pop())
	}

	return toStringSlice(output)
}

func toStringSlice(q *queue.Queue) ([]string, error) {
	result := make([]string, 0, q.Len())
	for q.Len() > 0 {
		item := q.Dequeue()
		if s, ok := item.(string); ok {
			result = append(result, s)
		} else {
			return nil, fmt.Errorf("'%v' not a string", item)
		}
	}

	return result, nil
}

func isOperand(token string) bool {
	for _, r := range token {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}
