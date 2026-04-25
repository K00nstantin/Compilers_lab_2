package main

import (
	"fmt"
	"unicode"
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
						if beta[0].Is_empty {
							beta = beta[1:]
						}
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

func left_rec_elimination(g grammar) grammar {
	newgram := g
	nonterminals := []token{}
	terminals := []token{}
	allrules := []rule{}
	for _, t := range g.Tokens {
		if unicode.IsUpper(rune(t.Name[0])) {
			nonterminals = append(nonterminals, t)
		} else {
			terminals = append(terminals, t)
		}
	}
	for _, r := range g.Rules {
		allrules = append(allrules, r)
	}
	for i := 0; i < len(nonterminals); i++ {
		for j := 0; j < i; j++ {
			for num, r := range allrules {
				if r.Left == nonterminals[i] && r.Right[0] == nonterminals[j] {
					combined_left_l := []token{}
					combined_left_l = r.Right[1:]
					temprules := []rule{}
					for _, s_r := range allrules {
						if s_r.Left == nonterminals[j] {
							combined_left := s_r.Right
							combined_left = append(combined_left, combined_left_l...)
							temprules = append(temprules, rule{Left: nonterminals[i], Right: combined_left})
						}
					}
					allrules[num] = temprules[0]
					allrules = append(allrules, temprules[1:]...)
				}
			}
		}
		rule_map := make(map[int]rule)
		for i, r := range allrules {
			rule_map[i] = r
		}
		newgram.Rules = rule_map
		newgram = simple_rec_removal(newgram)
	}
	return newgram
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
