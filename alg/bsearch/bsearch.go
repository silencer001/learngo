package bsearch

/*无重复数据的二分查找，返回index*/

func Bsearch(a []int, value int) int {
	low := 0
	high := len(a) - 1
	for low <= high { //low等于high时，mid=low=high
		mid := (low + high) / 2 //有可能溢出
		if a[mid] == value {
			return mid
		} else if a[mid] > value {
			high = mid - 1 //必须减一，否则会死循环
		} else {
			low = mid + 1 //必须加1，否则会死循环
		}
	}
	return -1 //notfound
}

/*递归方式的二分查找*/
func BsearchRecurs(a []int, value int) int {
	if len(a) == 0 || (len(a) == 1 && a[0] != value) {
		return -1
	}
	mid := len(a) / 2
	if a[mid] > value {
		return BsearchRecurs(a[:mid], value)
	}
	if a[mid] < value {
		ret := BsearchRecurs(a[mid+1:], value)
		if ret == -1 {
			return -1
		} else {
			return mid + 1 + ret
		}
	}
	if a[mid] == value {
		return mid
	}
	return -1
}
func BsearchFirst(a []int, value int) int {
	low := 0
	high := len(a) - 1
	for low <= high {
		mid := (low + high) / 2
		if a[mid] == value {
			if mid > 0 && a[mid-1] == a[mid] {
				high = mid - 1
			} else {
				return mid
			}
		} else if a[mid] > value {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
