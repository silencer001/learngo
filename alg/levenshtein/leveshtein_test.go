package levenshtein

import (
	"fmt"
	"testing"
)

func TestLevenshtein1(t *testing.T) {
	fmt.Println(levenshtein1("abc", "Abc"))
	fmt.Println(levenshtein1("abcd", "edafbghc"))
	fmt.Println(levenshtein2("abc", "Abc"))
	fmt.Println(levenshtein2("abcd", "edafbghc"))
	fmt.Println(levenshtein1("ACBD", "ADCB"))
	fmt.Println(levenshtein2("ACBD", "ADCB"))
	fmt.Println(lcs("abc", "abc"))
	fmt.Println(lcs("abcd", "edafbghc"))
}
