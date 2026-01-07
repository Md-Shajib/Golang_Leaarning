package errorhandling

import (
	"errors"
	"fmt"
)

type error interface {
	Error() string
}

func division(a, b float64) (float64, error){
	if b == 0{
		return 0, errors.New("Division by 0 is not allowed")
	}
	return a/b, nil
}

func ErrorMain() {
	fmt.Println("---------Error Handling-------------")

	result, err := division(10,6)
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println(result)
	}
}

// Need to learn more about error handling

func Need() string {
	return "Hi"
}