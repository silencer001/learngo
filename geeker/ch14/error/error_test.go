package err_test

import (
	"testing"

	"errors"
)

var TOO_SMALL_ERROR = errors.New("n should be bigger than 2")
var TOO_BIG_ERROR = errors.New("n should be smaller than 50")

func GetFibonacci(n int) ([]int64, error) {
	fibList := []int64{1, 1}
	if n < 2 {
		return nil, TOO_SMALL_ERROR
	}
	if n > 50 {
		return nil, TOO_BIG_ERROR
	}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	s, err := GetFibonacci(100)
	if err != nil {
		if err == TOO_SMALL_ERROR {
			t.Log("small error:", err)
		}
		if err == TOO_BIG_ERROR {
			t.Log("big error: ", err)
		}
		return
	}
	t.Log(s)
}
