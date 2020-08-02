package pipe_filter

import "testing"

func TestStraightPipeline(t *testing.T) {
	spliter := NewSplitFilter(",")
	converter := NewToIntFilter()
	sumer := NewSumFilter()
	sper := NewStraightPipeline("ju", spliter, converter, sumer)
	ret, err := sper.Process("1,2,3,4,5")
	if err != nil {
		t.Fatal(err)
	}
	if ret != 15 {
		t.Fatalf("the expected is 10, but the actual is %d", ret)
	}
}
