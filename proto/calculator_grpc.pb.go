// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: proto/calculator.proto

package calculator

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AdditionService_Add_FullMethodName = "/calculator.AdditionService/Add"
)

// AdditionServiceClient is the client API for AdditionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Addition Service
type AdditionServiceClient interface {
	Add(ctx context.Context, in *CalculationRequest, opts ...grpc.CallOption) (*CalculationResponse, error)
}

type additionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdditionServiceClient(cc grpc.ClientConnInterface) AdditionServiceClient {
	return &additionServiceClient{cc}
}

func (c *additionServiceClient) Add(ctx context.Context, in *CalculationRequest, opts ...grpc.CallOption) (*CalculationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CalculationResponse)
	err := c.cc.Invoke(ctx, AdditionService_Add_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdditionServiceServer is the server API for AdditionService service.
// All implementations must embed UnimplementedAdditionServiceServer
// for forward compatibility.
//
// Addition Service
type AdditionServiceServer interface {
	Add(context.Context, *CalculationRequest) (*CalculationResponse, error)
	mustEmbedUnimplementedAdditionServiceServer()
}

// UnimplementedAdditionServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAdditionServiceServer struct{}

func (UnimplementedAdditionServiceServer) Add(context.Context, *CalculationRequest) (*CalculationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedAdditionServiceServer) mustEmbedUnimplementedAdditionServiceServer() {}
func (UnimplementedAdditionServiceServer) testEmbeddedByValue()                         {}

// UnsafeAdditionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdditionServiceServer will
// result in compilation errors.
type UnsafeAdditionServiceServer interface {
	mustEmbedUnimplementedAdditionServiceServer()
}

func RegisterAdditionServiceServer(s grpc.ServiceRegistrar, srv AdditionServiceServer) {
	// If the following call pancis, it indicates UnimplementedAdditionServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AdditionService_ServiceDesc, srv)
}

func _AdditionService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdditionServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdditionService_Add_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdditionServiceServer).Add(ctx, req.(*CalculationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdditionService_ServiceDesc is the grpc.ServiceDesc for AdditionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdditionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.AdditionService",
	HandlerType: (*AdditionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _AdditionService_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/calculator.proto",
}

const (
	MultiplicationService_Multiply_FullMethodName = "/calculator.MultiplicationService/Multiply"
)

// MultiplicationServiceClient is the client API for MultiplicationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Multiplication Service
type MultiplicationServiceClient interface {
	Multiply(ctx context.Context, in *CalculationRequest, opts ...grpc.CallOption) (*CalculationResponse, error)
}

type multiplicationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMultiplicationServiceClient(cc grpc.ClientConnInterface) MultiplicationServiceClient {
	return &multiplicationServiceClient{cc}
}

func (c *multiplicationServiceClient) Multiply(ctx context.Context, in *CalculationRequest, opts ...grpc.CallOption) (*CalculationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CalculationResponse)
	err := c.cc.Invoke(ctx, MultiplicationService_Multiply_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MultiplicationServiceServer is the server API for MultiplicationService service.
// All implementations must embed UnimplementedMultiplicationServiceServer
// for forward compatibility.
//
// Multiplication Service
type MultiplicationServiceServer interface {
	Multiply(context.Context, *CalculationRequest) (*CalculationResponse, error)
	mustEmbedUnimplementedMultiplicationServiceServer()
}

// UnimplementedMultiplicationServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMultiplicationServiceServer struct{}

func (UnimplementedMultiplicationServiceServer) Multiply(context.Context, *CalculationRequest) (*CalculationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Multiply not implemented")
}
func (UnimplementedMultiplicationServiceServer) mustEmbedUnimplementedMultiplicationServiceServer() {}
func (UnimplementedMultiplicationServiceServer) testEmbeddedByValue()                               {}

// UnsafeMultiplicationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MultiplicationServiceServer will
// result in compilation errors.
type UnsafeMultiplicationServiceServer interface {
	mustEmbedUnimplementedMultiplicationServiceServer()
}

func RegisterMultiplicationServiceServer(s grpc.ServiceRegistrar, srv MultiplicationServiceServer) {
	// If the following call pancis, it indicates UnimplementedMultiplicationServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MultiplicationService_ServiceDesc, srv)
}

func _MultiplicationService_Multiply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MultiplicationServiceServer).Multiply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MultiplicationService_Multiply_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MultiplicationServiceServer).Multiply(ctx, req.(*CalculationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MultiplicationService_ServiceDesc is the grpc.ServiceDesc for MultiplicationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MultiplicationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.MultiplicationService",
	HandlerType: (*MultiplicationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Multiply",
			Handler:    _MultiplicationService_Multiply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/calculator.proto",
}