package pws

import "fmt"

func SliceDeclarationAndOperation() {
	numbers := []int{10, 20, 30}

	// pointer to slice
	ptr := &numbers
	fmt.Println("Before update", numbers)

	//update the slices
	updateSlice(ptr)
	fmt.Println("After update", numbers)

	// remove first element from the slices
	deleteFirst(ptr)
	fmt.Println("After deleting first:", numbers)

	fmt.Println("========================End of SLICE=======================")
}

func updateSlice(s *[]int){
	(*s)[0] = 100 // modify existing element
	*s = append(*s, 40, 50) // change slice length
}

func deleteFirst(s *[]int){
	*s = (*s)[1:]
}
