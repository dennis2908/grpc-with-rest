package client

import (
	"context"
	"fmt"
	"log"

	"github.com/flakrimjusufi/grpc-with-rest/models"
	customerpb "github.com/flakrimjusufi/grpc-with-rest/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	colorBlue   = "\033[34m"
	colorYellow = "\033[33m"
)

// CustomerServer - the grpc server of customers
type CustomerServer struct {
	customerpb.UnimplementedCustomerServiceServer
}

// SayHello - the service that prints a given name in the output
func (us *CustomerServer) SayHello(ctx context.Context, in *customerpb.Customer) (*customerpb.Message, error) {
	return &customerpb.Message{Message: "Hello " + in.Name}, nil
}

// CreateCustomer - the service that creates a customer in the Customers table and returns customerpb.Customer
func (us *CustomerServer) CreateCustomer(ctx context.Context, in *customerpb.Customer) (*customerpb.Customer, error) {
	customer := models.Customer{Name: in.Name, Email: in.Email, PhoneNumber: in.PhoneNumber}

	database.NewRecord(customer)
	database.Create(&customer)

	return &customerpb.Customer{Id: uint32(customer.ID), Name: customer.Name, Email: customer.Email, PhoneNumber: customer.PhoneNumber}, nil
}

// UpdateCustomerByName - the service that updates a customer by its name in the Customers table and returns customerpb.Customer
func (us *CustomerServer) UpdateCustomerByName(ctx context.Context, in *customerpb.Customer) (*customerpb.Customer, error) {

	name := in.GetName()
	if name == "" {
		return &customerpb.Customer{}, status.Error(codes.InvalidArgument, "Customer's name not specified")
	}
	email := in.GetEmail()
	phoneNumber := in.GetPhoneNumber()

	var customer models.Customer
	database.Where("name =?", name).Find(&customer)

	customer.Email = email
	customer.PhoneNumber = phoneNumber

	database.Save(&customer)

	return &customerpb.Customer{Id: uint32(customer.ID), Name: customer.Name, Email: customer.Email, PhoneNumber: customer.PhoneNumber}, nil
}

// UpdateCustomerByID - the service that updates a customer by its ID in the Customers table and returns customerpb.Customer
func (us *CustomerServer) UpdateCustomerByID(ctx context.Context, in *customerpb.Customer) (*customerpb.Customer, error) {

	id := in.GetId()
	name := in.GetName()
	email := in.GetEmail()
	phoneNumber := in.GetPhoneNumber()

	var customer models.Customer
	database.Where("id =?", id).Find(&customer)

	customer.Name = name
	customer.Email = email
	customer.PhoneNumber = phoneNumber

	database.Save(&customer)

	return &customerpb.Customer{Id: uint32(customer.ID), Name: customer.Name, Email: customer.Email, PhoneNumber: customer.PhoneNumber}, nil
}

// DeleteCustomer - the service that deletes a customer by its name in the Customers table and returns customerpb.Customer
func (us *CustomerServer) DeleteCustomer(ctx context.Context, in *customerpb.Customer) (*customerpb.Message, error) {
	name := in.GetName()
	if name == "" {
		return &customerpb.Message{}, status.Error(codes.InvalidArgument, "Customer's name not specified")
	}
	var customer models.Customer
	rowsAffected := database.Where("name =?", name).Find(&customer).RowsAffected

	if rowsAffected == 0 {
		return &customerpb.Message{}, status.Error(codes.NotFound, "Cannot find a Customer with this name!")
	}
	database.Delete(&customer)

	return &customerpb.Message{Message: customer.Name + " Deleted successfully!"}, nil
}

// ListCustomers - the service that lists the customers in the Customers table and returns customerpb.Customer
func (us *CustomerServer) ListCustomers(ctx context.Context, in *customerpb.Customer) (*customerpb.ListCustomer, error) {

	list := make([]*customerpb.Customer, 0)
	database.Where("deleted_at is null").Order("created_at desc").Limit(100).Find(&list)
	return &customerpb.ListCustomer{
		Customers: list,
	}, nil
}

// GetCustomerByName - the service that gets the customer by name in the Customers table and returns customerpb.Customer
func (us *CustomerServer) GetCustomerByName(ctx context.Context, in *customerpb.Customer) (*customerpb.Customer, error) {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	fmt.Println(colorYellow, "__________________________________________________________________________________"+
		"_______________________________________________________________________________________________________")
	log.Println(colorPurple, "[customerService] - [rpc GetCustomerByName] -> ",
		colorBlue, "Received person's name: ", in.GetName())
	fmt.Println(colorYellow, "__________________________________________________________________________________"+
		"_______________________________________________________________________________________________________")
	name := in.GetName()
	if name == "" {
		return &customerpb.Customer{}, status.Error(codes.InvalidArgument, "Customer's name not specified")
	}
	var customer models.Customer
	database.Where(&models.Customer{Name: name}).Find(&customer)

	return &customerpb.Customer{Id: uint32(customer.ID), Name: customer.Name, Email: customer.Email, PhoneNumber: customer.PhoneNumber}, nil
}

// GetCustomerByID - the service that gets the customer by ID in the Customers table and returns customerpb.Customer
func (us *CustomerServer) GetCustomerByID(ctx context.Context, in *customerpb.Customer) (*customerpb.Customer, error) {
	id := in.GetId()
	var customer models.Customer
	rowsAffected := database.Where("id = ?", id).Find(&customer).RowsAffected

	if rowsAffected == 0 {
		return &customerpb.Customer{}, status.Error(codes.NotFound, "Cannot find a Customer with this id!")
	}

	return &customerpb.Customer{Id: uint32(customer.ID), Name: customer.Name, Email: customer.Email, PhoneNumber: customer.PhoneNumber}, nil
}
