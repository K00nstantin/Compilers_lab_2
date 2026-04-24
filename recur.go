package main

import (
	"fmt"
)

type grammar struct {
	Tokens       map[int]token
	Rules        map[int]rule
	Start_symbol token
}

type token struct {
	Name     string
	Is_empty bool
}

type rule struct {
	Left  token
	Right []token
}

func create_new_token(t token, g *grammar) token {
	if t.Is_empty == true {
		fmt.Printf("Попытка сделать токен из пустого слова")
	}
	new_token := token{
		Name:     t.Name + "^",
		Is_empty: t.Is_empty,
	}
	g.Tokens[len(g.Tokens)] = new_token
	return new_token
}

func simple_rec_removal(g grammar) grammar {
	newgrammar := g
	for _, f_rule := range newgrammar.Rules {
		if f_rule.Left == f_rule.Right[0] {
			big_A := f_rule.Left
			big_A_new := create_new_token(big_A, &newgrammar)
			need_to_add_empty := false
			for num, s_rule := range newgrammar.Rules {
				if s_rule.Left == big_A {
					if s_rule.Right[0] == big_A {
						alpha := s_rule.Right[1:]
						alpha = append(alpha, big_A_new)
						newgrammar.Rules[num] = rule{Left: big_A_new, Right: alpha}
						need_to_add_empty = true
					} else {
						beta := s_rule.Right
						beta = append(beta, big_A_new)
						newgrammar.Rules[num] = rule{Left: big_A, Right: beta}
					}

				}
			}
			if need_to_add_empty {
				newgrammar.Rules[len(newgrammar.Rules)] = rule{Left: big_A_new, Right: []token{{Name: "empty", Is_empty: true}}}
			}
		}
	}
	return newgrammar
}

func print_gram(g grammar) {
	fmt.Println("Токены :")
	for _, t := range g.Tokens {
		fmt.Println(t.Name)
	}
	fmt.Println()
	fmt.Println("Правила : ")
	for _, r := range g.Rules {
		fmt.Printf("%s -> ", r.Left.Name)
		for _, t := range r.Right {
			fmt.Printf("%s ", t.Name)
		}
		fmt.Println()
	}
}
