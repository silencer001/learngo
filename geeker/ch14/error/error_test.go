package err_test

import (
	"testing"

	"errors"
)

func GetFibonacci(n int) ([]int, error) {
	fibList := []int{1, 1}
	if n < 2 {
		return nil, errors.New("n should be in[2:]")
	}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	s, err := GetFibonacci(-10)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(s)
}
