package interfaces

import "fmt"

type AllUser interface {
	ViewDetails() 
}

type User struct {
	name string
	age float32
	address string
}

type PremiumUser struct{
	User
	subscriptionType string
	amount float32
	from string
	since string
}

func (u User) ViewDetails(){
	fmt.Println("-----------User Details-------------")
	fmt.Println("Name: ", u.name)
	fmt.Println("Age: ", u.age)
	fmt.Println("Address: ", u.address)
}

func (p PremiumUser) ViewDetails() {
	fmt.Println("-----------Premium User Details--------------")
	p.User.ViewDetails()
	fmt.Println("Subscription Type: ",p.subscriptionType)
	fmt.Println("Amount: ", p.amount)
	fmt.Println("From: ", p.from, "to ", p.since)
}

func View(a AllUser){
	a.ViewDetails()
}

func InterfaceExample(){
	user1 := User{name: "Shajib", age: 26, address: "Dhaka"}
	View(user1)

	user2 := PremiumUser{
		User: user1,
		subscriptionType: "Monthly", 
		amount: 10.99, 
		from: "January", 
		since: "April",
	}
	View(user2)
}
