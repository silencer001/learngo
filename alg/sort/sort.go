package mysort

import "fmt"

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

/*计数排序：1、遍历数据，统计出每个值的个数，得到count数组
2、然后对count数组进行累加求和，得出小于等于每个值的累计元素个数
3、从后向前遍历数据，根据count数组的值放入相应临时数组中的index中，同时将count数组的值减1
限制条件：a中元素均为正整数，范围有限，元素最大值小于a的个数
*/
func CountSort(a []int) []int {
	/*创建count数组*/
	max := 0
	for _, v := range a {
		if v > max {
			max = v
		}
	}
	fmt.Println("max", max)
	count := make([]int, max+1)
	//遍历a，对每个值进行统计计数
	for _, v := range a {
		count[v]++
	}
	fmt.Println("count:", count)
	//遍历count，获得每个值的排名
	for i := 1; i < len(count); i++ {
		count[i] = count[i] + count[i-1]
	}
	fmt.Println("after count:", count)
	res := make([]int, len(a))
	for i := len(a) - 1; i >= 0; i-- {
		resIndex := count[a[i]] - 1
		res[resIndex] = a[i]
		count[a[i]]--
	}
	return res
}

/*找到数组中从小到大第k个的值，时间O(N)，空间O(1)*/
/*使用分治策略*/
func FindK(a []int, k int) int {
	pivot := partition2(a)
	if pivot == k-1 {
		return a[pivot]
	}
	if pivot > k-1 {
		return FindK(a[:pivot], k)
	} else {
		return FindK(a[pivot+1:], k-pivot-1)
	}
}

/*radix sort*/

func radix(a []int) []int {

}
