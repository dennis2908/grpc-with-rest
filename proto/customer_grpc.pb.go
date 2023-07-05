// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CustomerServiceClient is the client API for CustomerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomerServiceClient interface {
	CreateCustomer(ctx context.Context, in *Customer, opts ...grpc.CallOption) (*Customer, error)
	UpdateCustomerByName(ctx context.Context, in *Customer, opts ...grpc.CallOption) (*Customer, error)
	UpdateCustomerByID(ctx context.Context, in *Customer, opts ...grpc.CallOption) (*Customer, error)
	DeleteCustomer(ctx context.Context, in *Customer, opts ...grpc.CallOption) (*Message, error)
	ListCustomers(ctx context.Context, in *Customer, opts ...grpc.CallOption) (*ListCustomer, error)
	GetCustomerByName(ctx context.Context, in *Customer, opts ...grpc.CallOption) (*Customer, error)
	GetCustomerByID(ctx context.Context, in *Customer, opts ...grpc.CallOption) (*Customer, error)
	SayHello(ctx context.Context, in *Customer, opts ...grpc.CallOption) (*Message, error)
}

type customerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerServiceClient(cc grpc.ClientConnInterface) CustomerServiceClient {
	return &customerServiceClient{cc}
}

func (c *customerServiceClient) CreateCustomer(ctx context.Context, in *Customer, opts ...grpc.CallOption) (*Customer, error) {
	out := new(Customer)
	err := c.cc.Invoke(ctx, "/customer.CustomerService/CreateCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) UpdateCustomerByName(ctx context.Context, in *Customer, opts ...grpc.CallOption) (*Customer, error) {
	out := new(Customer)
	err := c.cc.Invoke(ctx, "/customer.CustomerService/UpdateCustomerByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) UpdateCustomerByID(ctx context.Context, in *Customer, opts ...grpc.CallOption) (*Customer, error) {
	out := new(Customer)
	err := c.cc.Invoke(ctx, "/customer.CustomerService/UpdateCustomerByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) DeleteCustomer(ctx context.Context, in *Customer, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/customer.CustomerService/DeleteCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) ListCustomers(ctx context.Context, in *Customer, opts ...grpc.CallOption) (*ListCustomer, error) {
	out := new(ListCustomer)
	err := c.cc.Invoke(ctx, "/customer.CustomerService/ListCustomers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) GetCustomerByName(ctx context.Context, in *Customer, opts ...grpc.CallOption) (*Customer, error) {
	out := new(Customer)
	err := c.cc.Invoke(ctx, "/customer.CustomerService/GetCustomerByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) GetCustomerByID(ctx context.Context, in *Customer, opts ...grpc.CallOption) (*Customer, error) {
	out := new(Customer)
	err := c.cc.Invoke(ctx, "/customer.CustomerService/GetCustomerByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) SayHello(ctx context.Context, in *Customer, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/customer.CustomerService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerServiceServer is the server API for CustomerService service.
// All implementations must embed UnimplementedCustomerServiceServer
// for forward compatibility
type CustomerServiceServer interface {
	CreateCustomer(context.Context, *Customer) (*Customer, error)
	UpdateCustomerByName(context.Context, *Customer) (*Customer, error)
	UpdateCustomerByID(context.Context, *Customer) (*Customer, error)
	DeleteCustomer(context.Context, *Customer) (*Message, error)
	ListCustomers(context.Context, *Customer) (*ListCustomer, error)
	GetCustomerByName(context.Context, *Customer) (*Customer, error)
	GetCustomerByID(context.Context, *Customer) (*Customer, error)
	SayHello(context.Context, *Customer) (*Message, error)
	mustEmbedUnimplementedCustomerServiceServer()
}

// UnimplementedCustomerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCustomerServiceServer struct {
}

func (UnimplementedCustomerServiceServer) CreateCustomer(context.Context, *Customer) (*Customer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCustomer not implemented")
}
func (UnimplementedCustomerServiceServer) UpdateCustomerByName(context.Context, *Customer) (*Customer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCustomerByName not implemented")
}
func (UnimplementedCustomerServiceServer) UpdateCustomerByID(context.Context, *Customer) (*Customer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCustomerByID not implemented")
}
func (UnimplementedCustomerServiceServer) DeleteCustomer(context.Context, *Customer) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCustomer not implemented")
}
func (UnimplementedCustomerServiceServer) ListCustomers(context.Context, *Customer) (*ListCustomer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCustomers not implemented")
}
func (UnimplementedCustomerServiceServer) GetCustomerByName(context.Context, *Customer) (*Customer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomerByName not implemented")
}
func (UnimplementedCustomerServiceServer) GetCustomerByID(context.Context, *Customer) (*Customer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomerByID not implemented")
}
func (UnimplementedCustomerServiceServer) SayHello(context.Context, *Customer) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedCustomerServiceServer) mustEmbedUnimplementedCustomerServiceServer() {}

// UnsafeCustomerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomerServiceServer will
// result in compilation errors.
type UnsafeCustomerServiceServer interface {
	mustEmbedUnimplementedCustomerServiceServer()
}

func RegisterCustomerServiceServer(s grpc.ServiceRegistrar, srv CustomerServiceServer) {
	s.RegisterService(&CustomerService_ServiceDesc, srv)
}

func _CustomerService_CreateCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Customer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).CreateCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer.CustomerService/CreateCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).CreateCustomer(ctx, req.(*Customer))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_UpdateCustomerByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Customer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).UpdateCustomerByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer.CustomerService/UpdateCustomerByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).UpdateCustomerByName(ctx, req.(*Customer))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_UpdateCustomerByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Customer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).UpdateCustomerByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer.CustomerService/UpdateCustomerByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).UpdateCustomerByID(ctx, req.(*Customer))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_DeleteCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Customer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).DeleteCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer.CustomerService/DeleteCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).DeleteCustomer(ctx, req.(*Customer))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_ListCustomers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Customer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).ListCustomers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer.CustomerService/ListCustomers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).ListCustomers(ctx, req.(*Customer))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_GetCustomerByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Customer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).GetCustomerByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer.CustomerService/GetCustomerByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).GetCustomerByName(ctx, req.(*Customer))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_GetCustomerByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Customer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).GetCustomerByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer.CustomerService/GetCustomerByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).GetCustomerByID(ctx, req.(*Customer))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Customer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer.CustomerService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).SayHello(ctx, req.(*Customer))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomerService_ServiceDesc is the grpc.ServiceDesc for CustomerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "customer.CustomerService",
	HandlerType: (*CustomerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCustomer",
			Handler:    _CustomerService_CreateCustomer_Handler,
		},
		{
			MethodName: "UpdateCustomerByName",
			Handler:    _CustomerService_UpdateCustomerByName_Handler,
		},
		{
			MethodName: "UpdateCustomerByID",
			Handler:    _CustomerService_UpdateCustomerByID_Handler,
		},
		{
			MethodName: "DeleteCustomer",
			Handler:    _CustomerService_DeleteCustomer_Handler,
		},
		{
			MethodName: "ListCustomers",
			Handler:    _CustomerService_ListCustomers_Handler,
		},
		{
			MethodName: "GetCustomerByName",
			Handler:    _CustomerService_GetCustomerByName_Handler,
		},
		{
			MethodName: "GetCustomerByID",
			Handler:    _CustomerService_GetCustomerByID_Handler,
		},
		{
			MethodName: "SayHello",
			Handler:    _CustomerService_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/customer.proto",
}

// CreditCardCustomerServiceClient is the client API for CreditCardCustomerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CreditCardCustomerServiceClient interface {
	CreditCardCustomers(ctx context.Context, in *CreditCardCustomer, opts ...grpc.CallOption) (*ListCreditCardCustomers, error)
	GetCreditCardCustomerByCustomerName(ctx context.Context, in *CreditCardCustomer, opts ...grpc.CallOption) (*CreditCardCustomer, error)
	CreateCreditCardCustomerApplication(ctx context.Context, in *CreditCardCustomerApplication, opts ...grpc.CallOption) (*CreditCardCustomerApplication, error)
	GetCreditCardCustomerApplicationByName(ctx context.Context, in *CreditCardCustomerApplication, opts ...grpc.CallOption) (*CreditCardCustomerApplication, error)
}

type creditCardCustomerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCreditCardCustomerServiceClient(cc grpc.ClientConnInterface) CreditCardCustomerServiceClient {
	return &creditCardCustomerServiceClient{cc}
}

func (c *creditCardCustomerServiceClient) CreditCardCustomers(ctx context.Context, in *CreditCardCustomer, opts ...grpc.CallOption) (*ListCreditCardCustomers, error) {
	out := new(ListCreditCardCustomers)
	err := c.cc.Invoke(ctx, "/customer.CreditCardCustomerService/CreditCardCustomers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creditCardCustomerServiceClient) GetCreditCardCustomerByCustomerName(ctx context.Context, in *CreditCardCustomer, opts ...grpc.CallOption) (*CreditCardCustomer, error) {
	out := new(CreditCardCustomer)
	err := c.cc.Invoke(ctx, "/customer.CreditCardCustomerService/GetCreditCardCustomerByCustomerName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creditCardCustomerServiceClient) CreateCreditCardCustomerApplication(ctx context.Context, in *CreditCardCustomerApplication, opts ...grpc.CallOption) (*CreditCardCustomerApplication, error) {
	out := new(CreditCardCustomerApplication)
	err := c.cc.Invoke(ctx, "/customer.CreditCardCustomerService/CreateCreditCardCustomerApplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creditCardCustomerServiceClient) GetCreditCardCustomerApplicationByName(ctx context.Context, in *CreditCardCustomerApplication, opts ...grpc.CallOption) (*CreditCardCustomerApplication, error) {
	out := new(CreditCardCustomerApplication)
	err := c.cc.Invoke(ctx, "/customer.CreditCardCustomerService/GetCreditCardCustomerApplicationByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CreditCardCustomerServiceServer is the server API for CreditCardCustomerService service.
// All implementations must embed UnimplementedCreditCardCustomerServiceServer
// for forward compatibility
type CreditCardCustomerServiceServer interface {
	CreditCardCustomers(context.Context, *CreditCardCustomer) (*ListCreditCardCustomers, error)
	GetCreditCardCustomerByCustomerName(context.Context, *CreditCardCustomer) (*CreditCardCustomer, error)
	CreateCreditCardCustomerApplication(context.Context, *CreditCardCustomerApplication) (*CreditCardCustomerApplication, error)
	GetCreditCardCustomerApplicationByName(context.Context, *CreditCardCustomerApplication) (*CreditCardCustomerApplication, error)
	mustEmbedUnimplementedCreditCardCustomerServiceServer()
}

// UnimplementedCreditCardCustomerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCreditCardCustomerServiceServer struct {
}

func (UnimplementedCreditCardCustomerServiceServer) CreditCardCustomers(context.Context, *CreditCardCustomer) (*ListCreditCardCustomers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreditCardCustomers not implemented")
}
func (UnimplementedCreditCardCustomerServiceServer) GetCreditCardCustomerByCustomerName(context.Context, *CreditCardCustomer) (*CreditCardCustomer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCreditCardCustomerByCustomerName not implemented")
}
func (UnimplementedCreditCardCustomerServiceServer) CreateCreditCardCustomerApplication(context.Context, *CreditCardCustomerApplication) (*CreditCardCustomerApplication, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCreditCardCustomerApplication not implemented")
}
func (UnimplementedCreditCardCustomerServiceServer) GetCreditCardCustomerApplicationByName(context.Context, *CreditCardCustomerApplication) (*CreditCardCustomerApplication, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCreditCardCustomerApplicationByName not implemented")
}
func (UnimplementedCreditCardCustomerServiceServer) mustEmbedUnimplementedCreditCardCustomerServiceServer() {}

// UnsafeCreditCardCustomerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CreditCardCustomerServiceServer will
// result in compilation errors.
type UnsafeCreditCardCustomerServiceServer interface {
	mustEmbedUnimplementedCreditCardCustomerServiceServer()
}

func RegisterCreditCardCustomerServiceServer(s grpc.ServiceRegistrar, srv CreditCardCustomerServiceServer) {
	s.RegisterService(&CreditCardCustomerService_ServiceDesc, srv)
}

func _CreditCardCustomerService_CreditCardCustomers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreditCardCustomer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreditCardCustomerServiceServer).CreditCardCustomers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer.CreditCardCustomerService/CreditCardCustomers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreditCardCustomerServiceServer).CreditCardCustomers(ctx, req.(*CreditCardCustomer))
	}
	return interceptor(ctx, in, info, handler)
}

func _CreditCardCustomerService_GetCreditCardCustomerByCustomerName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreditCardCustomer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreditCardCustomerServiceServer).GetCreditCardCustomerByCustomerName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer.CreditCardCustomerService/GetCreditCardCustomerByCustomerName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreditCardCustomerServiceServer).GetCreditCardCustomerByCustomerName(ctx, req.(*CreditCardCustomer))
	}
	return interceptor(ctx, in, info, handler)
}

func _CreditCardCustomerService_CreateCreditCardCustomerApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreditCardCustomerApplication)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreditCardCustomerServiceServer).CreateCreditCardCustomerApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer.CreditCardCustomerService/CreateCreditCardCustomerApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreditCardCustomerServiceServer).CreateCreditCardCustomerApplication(ctx, req.(*CreditCardCustomerApplication))
	}
	return interceptor(ctx, in, info, handler)
}

func _CreditCardCustomerService_GetCreditCardCustomerApplicationByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreditCardCustomerApplication)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreditCardCustomerServiceServer).GetCreditCardCustomerApplicationByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer.CreditCardCustomerService/GetCreditCardCustomerApplicationByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreditCardCustomerServiceServer).GetCreditCardCustomerApplicationByName(ctx, req.(*CreditCardCustomerApplication))
	}
	return interceptor(ctx, in, info, handler)
}

// CreditCardCustomerService_ServiceDesc is the grpc.ServiceDesc for CreditCardCustomerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CreditCardCustomerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "customer.CreditCardCustomerService",
	HandlerType: (*CreditCardCustomerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreditCardCustomers",
			Handler:    _CreditCardCustomerService_CreditCardCustomers_Handler,
		},
		{
			MethodName: "GetCreditCardCustomerByCustomerName",
			Handler:    _CreditCardCustomerService_GetCreditCardCustomerByCustomerName_Handler,
		},
		{
			MethodName: "CreateCreditCardCustomerApplication",
			Handler:    _CreditCardCustomerService_CreateCreditCardCustomerApplication_Handler,
		},
		{
			MethodName: "GetCreditCardCustomerApplicationByName",
			Handler:    _CreditCardCustomerService_GetCreditCardCustomerApplicationByName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/customer.proto",
}