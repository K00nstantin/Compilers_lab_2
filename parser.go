package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse(str string) []token {
	tokens := []token{}
	ts := strings.Fields(str)
	for _, strs := range ts {
		if strs == "empty" {
			tokens = append(tokens, token{Name: "empty", Is_empty: true})
		} else {
			tokens = append(tokens, token{Name: strs, Is_empty: false})
		}
	}
	return tokens
}

func parse_file() grammar {
	file, err := os.Open("grammar.txt")
	if err != nil {
		fmt.Printf("error while opening file")
	}
	defer file.Close()

	parsed_grammar := grammar{
		Tokens:       make(map[int]token),
		Rules:        make(map[int]rule),
		Start_symbol: token{},
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	fmt.Printf("Колличество нетерминалов: %s \n", line)
	scanner.Scan()
	line = scanner.Text()
	parsed_tokens_nt := parse(line)
	for _, tk := range parsed_tokens_nt {
		parsed_grammar.Tokens[len(parsed_grammar.Tokens)] = tk
	}
	scanner.Scan()
	line = scanner.Text()
	fmt.Printf("Колличество терминалов: %s \n", line)
	scanner.Scan()
	line = scanner.Text()
	parsed_tokens_t := parse(line)
	for _, tk := range parsed_tokens_t {
		parsed_grammar.Tokens[len(parsed_grammar.Tokens)] = tk
	}
	scanner.Scan()
	line = scanner.Text()
	rule_num, _ := strconv.Atoi(line)
	for i := 0; i < rule_num; i++ {
		scanner.Scan()
		line = scanner.Text()
		ts := parse(line)
		parsed_grammar.Rules[len(parsed_grammar.Rules)] = rule{Left: ts[0], Right: ts[2:]}
	}
	scanner.Scan()
	line = scanner.Text()
	ts := parse(line)
	parsed_grammar.Start_symbol = ts[0]
	return parsed_grammar
}
