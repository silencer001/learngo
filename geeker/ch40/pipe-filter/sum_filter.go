package pipe_filter

import "errors"

var SumFilterBadFomatError = errors.New("Input should be int slice")

type SumFilter struct {
}

func NewSumFilter() *SumFilter {
	return &SumFilter{}
}

func (sf *SumFilter) Process(data Request) (Response, error) {
	s, ok := data.([]int)
	if !ok {
		return nil, SumFilterBadFomatError
	}
	var sum int
	for _, v := range s {
		sum += v
	}
	return sum, nil
}
