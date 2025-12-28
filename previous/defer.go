package previous

import "fmt"

func defer_test() {
	func() {
		i := 0
		defer fmt.Println(i)
		i++
	}()
}
