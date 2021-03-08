package findv

/*使用递归方式，最主要的是采用快速排序的思想，将数组分为三部分:小于pivot、pivot、大于pivot*/
/*如何将数组分为三部分：使用选择排序的思想，遍历整个数组，小于pivot的放数组前边，大于pivot的放数组后边*/
func FindV(a []int, v int) int {
	if len(a) == 1 {
		if a[0] == v {
			return 1
		} else {
			return -1 //not found
		}
	}
	pivot := a[len(a)-1]
	j := 0
	for i := 0; i < len(a)-1; i++ {
		if a[i] < pivot {
			a[j], a[i] = a[i], a[j]
			j++ //前j个为小于pivot的值
		}
	}
	a[j], a[len(a)-1] = a[len(a)-1], a[j]
	if pivot == v {
		return j + 1
	} else if pivot > v {
		return FindV(a[:j], v)
	} else {
		index := FindV(a[j+1:], v)
		if index == -1 {
			return -1
		} else {
			return j + 1 + index
		}
	}

}
