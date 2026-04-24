package main

import (
	"fmt"
)

func main() {
	E := token{
		Name:     "E",
		Is_empty: false,
	}
	T := token{
		Name:     "T",
		Is_empty: false,
	}
	plus := token{
		Name:     "+",
		Is_empty: false,
	}

	first_rule := rule{
		Left:  E,
		Right: []token{E, plus, T},
	}
	second_rule := rule{
		Left:  E,
		Right: []token{T},
	}

	ts := map[int]token{
		0: E,
		1: T,
		2: plus,
	}

	rs := map[int]rule{
		0: first_rule,
		1: second_rule,
	}

	test_grammar := grammar{
		Tokens: ts,
		Rules:  rs,
	}

	print_gram(test_grammar)
	fmt.Println()
	new_grammar := simple_rec_removal(test_grammar)
	fmt.Printf("А вот новая грамматика: \n")
	print_gram(new_grammar)

}
