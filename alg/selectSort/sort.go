package selectSort

/*选择排序，每次在未排序的列表中选择最小的，放在已排序列表的末尾*/
func Sort(a []int) []int {
	for i := 0; i < len(a); i++ {
		min := i
		for j := i; j < len(a); j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		a[i], a[min] = a[min], a[i]
	}
	return a
}
