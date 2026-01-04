package previous

import (
	"fmt"
	"time"
)

func printSomething(num int) {
	fmt.Println("Hello Shajib ", num)
}

func mainPrint() {
	var a = 10
	const b = 20

	go printSomething(1)
	go printSomething(2)
	go printSomething(3)
	go printSomething(4)
	go printSomething(5)

	fmt.Println(a," ", b)

	time.Sleep(time.Second*1)
}


/*
            simulation
--------------------------------
1. Compile Phase
2. Runtime Phase


****************Compile Phase***************
go build main.go => compile it => ./main (Main name akta file create hobe)

b = 20
printSomething
mainPrint

*/