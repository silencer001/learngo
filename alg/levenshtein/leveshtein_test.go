package levenshtein

import (
	"fmt"
	"testing"
)

func TestLevenshtein1(t *testing.T) {
	fmt.Println(levenshtein1("abc", "Abc"))
	fmt.Println(levenshtein1("abcd", "edafbghc"))
}
