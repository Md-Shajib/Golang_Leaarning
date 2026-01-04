package paymentinterface

import (
	"fmt"
	"go-basic/payment_interface/internal/gateway"
	"go-basic/payment_interface/internal/logger"
	"go-basic/payment_interface/internal/payment"
	"go-basic/payment_interface/internal/repository"
	"go-basic/payment_interface/internal/service"
)

func PaymantCmd(){
	log := logger.NewConsoleLogger()
	repo := repository.NewInMemoryPaymentRepository()

	deps := struct {
		payment.Logger
		payment.PaymentRepository
	}{
		log,
		repo,
	}

	gateway := &gateway.StripeGateway{}

	paymentService := service.NewPaymentService(gateway, deps)

	paymentService.ProcessPayment(1000, "USD")

	fmt.Println("=====================End of Example===================")
}