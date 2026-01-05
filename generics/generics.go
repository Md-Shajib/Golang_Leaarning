package generics

import "fmt"

func isNameExist(names []string, query string) bool {
	for _, name := range names {
		if name == query {
			return true
		}
	}
	return false
}

func isNumberExist(nums []int, query int) bool {
	for _, num := range nums {
		if num == query {
			return true
		}
	}
	return false
}

func isFloatExist(fnums []float32, query float32) bool {
	for _, num := range fnums {
		if num == query {
			return true
		}
	}
	return false
}

func genericIsAvailable[T string | int | float32](valueList []T, query T) bool {
	for _, val := range valueList {
		if val == query {
			return true
		}
	}
	return false
}

func GenericsFunction() {
	fmt.Println("---------------Generics Start-------------------")

	names := []string{"shajib", "firoj", "rony", "emon", "bappy"}
	query := "shajib"
	fmt.Println(isNameExist(names, query))

	nums := []int{23, 12, 32, 43, 54}
	searchNum := 50
	fmt.Println(isNumberExist(nums, searchNum))

	fnums := []float32{12.21, 23.12, 54.23, 65.23}
	searchFnum := 12.21
	fmt.Println(isFloatExist(fnums, float32(searchFnum)))

	fmt.Println("----Generic Call----")
	fmt.Println(genericIsAvailable[string](names, query))
	fmt.Println(genericIsAvailable(nums, searchNum))
	fmt.Println(genericIsAvailable(fnums, float32(searchFnum)))
	
}
