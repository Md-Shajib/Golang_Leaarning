package interfaces

import "fmt"

type i interface{}

func EmptyInterface() {
	fmt.Println("===================Empty Interface===================")
	emptyIntetrfaceExample(23)
	emptyIntetrfaceExample(true)
}

func emptyIntetrfaceExample(i interface{}) {

	fmt.Println(i)
}
