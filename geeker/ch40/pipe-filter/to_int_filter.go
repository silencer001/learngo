package pipe_filter

import (
	"errors"
	"strconv"
)

var ToIntFilterBadFormatFilter = errors.New("Input should be a string slice")

type ToIntFilter struct {
}

func NewToIntFilter() *ToIntFilter {
	return &ToIntFilter{}
}

func (tif *ToIntFilter) Process(data Request) (Response, error) {
	s, ok := data.([]string)
	if !ok {
		return nil, ToIntFilterBadFormatFilter
	}
	var si []int
	for _, v := range s {
		i, _ := strconv.Atoi(v)
		si = append(si, i)
	}
	return si, nil
}
