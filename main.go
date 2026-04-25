package main

import (
	"fmt"
)

func main() {
	test_grammar := parse_file()
	print_gram(test_grammar)
	fmt.Println()
	fmt.Println()
	cnf_grammar := to_cnf(test_grammar)
	new_grammar := left_rec_elimination(test_grammar)
	fmt.Printf("А вот новая грамматика: \n")
	print_gram(new_grammar)
	fmt.Println("\n")
	fmt.Println("Приведенная к форме Хомского: ")
	print_gram(cnf_grammar)
}
