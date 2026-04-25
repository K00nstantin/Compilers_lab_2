package main

import (
	"fmt"
)

func main() {
	test_grammar := parse_file()
	print_gram(test_grammar)
	fmt.Println()
	fmt.Println()
	new_grammar := left_rec_removal(test_grammar)
	fmt.Printf("А вот новая грамматика: \n")
	print_gram(new_grammar)

}
