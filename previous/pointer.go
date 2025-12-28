package previous

// import "fmt"

func Pointers() {
	i, j := 30, 40

	p := &i
	q := &j

	*p = 35
	*q = 45

	// fmt.Println("i = ",i , " j = ",j)
	// defer_test()
	// embedding_main()
	mainPrint()
}
