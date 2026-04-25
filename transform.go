package main

import (
	"strconv"
	"unicode"
)

func is_nonterminal(t token) bool {
	return !t.Is_empty && len(t.Name) > 0 && unicode.IsUpper(rune(t.Name[0]))
}

func is_terminal(t token) bool {
	return !t.Is_empty && !is_nonterminal(t)
}

func to_cnf(g grammar) grammar {
	newgram := grammar{
		Tokens:       make(map[int]token),
		Rules:        make(map[int]rule),
		Start_symbol: g.Start_symbol,
	}

	for _, t := range g.Tokens {
		newgram.Tokens[len(newgram.Tokens)] = t
	}

	term_to_nt := make(map[string]token)
	get_or_create_term_nt := func(t token) token {
		if nt, ok := term_to_nt[t.Name]; ok {
			return nt
		}
		new_nt := token{
			Name:     t.Name + "^",
			Is_empty: false,
		}
		term_to_nt[t.Name] = new_nt
		newgram.Tokens[len(newgram.Tokens)] = new_nt
		return new_nt
	}

	add_rule := func(left token, right []token) {
		newgram.Rules[len(newgram.Rules)] = rule{Left: left, Right: right}
	}

	new_nt_counter := 0

	// основной проход по правилам
	for _, r := range g.Rules {
		big_A := r.Left
		right := r.Right

		// оставить S -> empty (если было)
		if len(right) == 1 && right[0].Is_empty {
			if big_A == g.Start_symbol {
				add_rule(big_A, right)
			}
			continue
		}

		// 1) A -> a
		if len(right) == 1 && is_terminal(right[0]) {
			add_rule(big_A, right)
			continue
		}

		xs := make([]token, len(right))
		copy(xs, right)
		if len(xs) >= 2 {
			for i := 0; i < len(xs); i++ {
				if is_terminal(xs[i]) {
					xs[i] = get_or_create_term_nt(xs[i])
				}
			}
		}

		if len(xs) == 2 {
			add_rule(big_A, xs)
			continue
		}

		cur_left := big_A
		for i := 0; i < len(xs)-2; i++ {
			new_nt_counter++
			next_nt := token{
				Name:     "N" + strconv.Itoa(new_nt_counter),
				Is_empty: false,
			}
			newgram.Tokens[len(newgram.Tokens)] = next_nt
			add_rule(cur_left, []token{xs[i], next_nt})
			cur_left = next_nt
		}
		add_rule(cur_left, []token{xs[len(xs)-2], xs[len(xs)-1]})
	}

	for term_name, nt := range term_to_nt {
		add_rule(nt, []token{{Name: term_name, Is_empty: false}})
	}

	return newgram
}
