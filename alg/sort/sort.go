package mysort

/*选择排序，每次在未排序的列表中选择最小的，放在已排序列表的末尾*/
func SelectSort(a []int) []int {
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

/*归并排序使用的是分治思想，适合递归编程*/
func MergeSort(a []int) []int {
	len := len(a)
	if len <= 1 {
		return a
	}
	a1 := MergeSort(a[:len/2])
	a2 := MergeSort(a[len/2:])
	return merge(a1, a2)
}

/*合并两个有序的切片*/
func merge(a1, a2 []int) []int {
	var i, j int = 0, 0
	var tmp int = 0
	res := make([]int, len(a1)+len(a2))
	for i < len(a1) && j < len(a2) {
		if a1[i] < a2[j] {
			res[tmp] = a1[i]
			tmp++
			i++
		} else {
			res[tmp] = a2[j]
			tmp++
			j++
		}
	}
	for i < len(a1) {
		res[tmp] = a1[i]
		tmp++
		i++
	}
	for j < len(a2) {
		res[tmp] = a2[j]
		tmp++
		j++
	}
	return res
}

/*快速排序同样是采用分治思想，选择任意一个数据作为pivot，然后将小于pivot的放在左边，大于pivot的放在右边*/
func QuikSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	pivotIndex := partition2(a)
	QuikSort(a[:pivotIndex])
	QuikSort(a[pivotIndex+1:])
	return a
}

/*使用冒泡排序的思想，将pivot放在数组中正确位置，此时是稳定的排序算法*/
/*冒泡排序本身是稳定的，但初始pivot和i的交换可能不稳定，例如
2，3，1，7，5a，10，5b，5c
交换第一次后，此时5c在5a前边，就不稳定了
2，3，1，5c，5a，10，5b，7 (pivot=3)
但如果值相同时仍然互换，就可以变稳定
2，3，1，5a，5c，10，5b，7 (pivot=4)
2，3，1，5a，10，5c，5b,7 (pivot=5)
2，3，1，5a，10，5b，5c，7 (pivot=6)此时又重新稳定了

*/
func partition(a []int) int {
	pivot := len(a) - 1
	var i int = 0
	for ; i < len(a)-1; i++ {
		if a[i] > a[pivot] { //可以使用>=，将不稳定转换为稳定
			a[i], a[pivot] = a[pivot], a[i] //当pivot的值有多个时，pivot和i的互换导致不稳定
			pivot = i
		}
	}
	return pivot
}

/*使用选择排序的思想，取得pivot的位置，此时一定不稳定*/
func partition2(a []int) int {
	pivot := a[len(a)-1]
	i := 0
	for j := 0; j < len(a)-1; j++ {
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[len(a)-1] = a[len(a)-1], a[i]
	return i
}
