package matrix

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var m [][]int = [][]int{
	{0, 3, 5, 9},
	{2, 1, 3, 4},
	{5, 2, 6, 7},
	{6, 8, 4, 3},
}

func TestFindCostMin1(t *testing.T) {
	fmt.Println(findCostMin1(m, 3, 3))
}

func TestFindCostMin2(t *testing.T) {
	fmt.Println(findCostMin2(m, 3, 3))
}

func TestFindCostMin3(t *testing.T) {
	fmt.Println(findCostMin3(m, 3, 3))
}
func TestFindCostMin(t *testing.T) {
	m := createMatrix(10)
	fmt.Printf("%d\n%d\n%d\n%d\n",
		findCostMin1(m, 9, 9),
		findCostMin2(m, 9, 9),
		findCostMin3(m, 9, 9),
		findCostMin4(m, 9, 9))
}

func createMatrix(n int) [][]int {
	rand.Seed(time.Now().UnixNano())
	m := make([][]int, n)
	for i, _ := range m {
		m[i] = make([]int, n)
		for j, _ := range m[i] {
			m[i][j] = rand.Int() % 100
		}
	}
	return m
}

func BenchmarkFindCostMin1(b *testing.B) {
	m := createMatrix(17)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		findCostMin1(m, 16, 16)
	}
}

func BenchmarkFindCostMin2(b *testing.B) {
	m := createMatrix(17)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		findCostMin2(m, 16, 16)
	}
}

func BenchmarkFindCostMin3(b *testing.B) {
	m := createMatrix(17)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		findCostMin3(m, 16, 16)
	}
}

func BenchmarkFindCostMin4(b *testing.B) {
	m := createMatrix(17)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		findCostMin4(m, 16, 16)
	}
}
