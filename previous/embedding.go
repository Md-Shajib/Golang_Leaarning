package previous

import "fmt"

type Person struct {
	name string
	phn  int
	addr string
}

type Employee struct {
	Person
	id     int
	salary float32
	dept   string
}

func (E *Employee) setInfo() {
	fmt.Scan(&E.name, &E.phn, &E.addr, &E.id, &E.salary, &E.dept)
}

func (E Employee) getInfo() {
	fmt.Println("Name: ", E.name)
	fmt.Println("ID: ", E.id)
	fmt.Println("Dept ", E.dept)
	fmt.Println("Salary: ", E.salary)
	fmt.Println("Phone: ", E.phn)
	fmt.Println("Address: ", E.addr)
}

func embedding_main() {
	emp := Employee{}
	emp.setInfo()
	emp.getInfo()
}
