package pwm

import "fmt"

func MapDeclaration(){
	// Basic map declaration with a literal
	var map1 map[string]int = map[string]int{
		"apple": 5,
		"banana": 10,
		"mango": 50,
	}

	// Map declaration with make function
	map2 := make(map[string]int)
	map2["total"] = 65
	map2["west"] = 2

	fmt.Println("Previous operation map 1 = ", map1)
	fmt.Println("Previous operation map 2 = ", map2)

	// Add and changes in map are same, if exist the key then change the value, else not exist then create a new key.
	map1["apple"] = 10
	map1["orange"] = 20
	fmt.Println("After changing map 1 = ", map1)

	// Delete from map
	delete(map1, "apple")
	fmt.Println("After deleting element map 1 = ", map1)

	// Checking the key is exist or not
	const key = "orange"
	value, ok := map1[key]
	if ok == false {
		fmt.Println("The ", key, " is not present in the map")
	}else{
		fmt.Println(" Existing status:", ok, "\n The key is", key, "\n And the value is:", value)
	}

	// View total:
	fmt.Println("\nTotal element of map1 = ",len(map1))
	for key, val := range map1 {
		fmt.Println(key, "=>", val)
	}
}