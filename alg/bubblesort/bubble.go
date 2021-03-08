package bubblesort

func Bubble(a []int) []int {
	flag := false
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	return a
}

func InsertSort(a []int) []int {
	for i := 1; i < len(a); i++ {
		value := a[i]
		for j := i - 1; j >= 0; j-- {
			if value < a[j] {
				a[j+1] = a[j]
			} else {
				a[j+1] = value
				break
			}

		}
	}
	return a
}
