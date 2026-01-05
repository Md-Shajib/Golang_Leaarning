package cmd

import (
	"go-basic/generics"
	interfaces "go-basic/interface"
	paymentinterface "go-basic/payment_interface"
	"go-basic/pointers_with_map_and_slice/pwm"
	"go-basic/pointers_with_map_and_slice/pws"
)

func Serve(){
	pwm.MapDeclarationAndOperation()
	pws.SliceDeclarationAndOperation()
	paymentinterface.PaymantCmd()
	interfaces.InterfaceExample()
	generics.GenericsFunction()
}