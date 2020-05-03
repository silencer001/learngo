package main

import "fmt"

func main() {
	// id1 := 1
	// id2 := 2
	// id3 := 3
	//数组，同一个类型的集合

	var id [50]int
	id[0] = 1
	id[1] = 2
	id[3] = 3
	for i := 0; i < len(id); i++ {
		id[i] = i + 1
		fmt.Printf("id[%d] = %d\n", i, id[i])
	}
}
