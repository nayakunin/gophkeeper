// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: proto/api.proto

package generated

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

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	AuthenticateUser(ctx context.Context, in *AuthenticateUserRequest, opts ...grpc.CallOption) (*AuthenticateUserResponse, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) AuthenticateUser(ctx context.Context, in *AuthenticateUserRequest, opts ...grpc.CallOption) (*AuthenticateUserResponse, error) {
	out := new(AuthenticateUserResponse)
	err := c.cc.Invoke(ctx, "/gophkeeper.AuthService/AuthenticateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	AuthenticateUser(context.Context, *AuthenticateUserRequest) (*AuthenticateUserResponse, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) AuthenticateUser(context.Context, *AuthenticateUserRequest) (*AuthenticateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthenticateUser not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_AuthenticateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AuthenticateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gophkeeper.AuthService/AuthenticateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AuthenticateUser(ctx, req.(*AuthenticateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gophkeeper.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuthenticateUser",
			Handler:    _AuthService_AuthenticateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/api.proto",
}

// RegistrationServiceClient is the client API for RegistrationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RegistrationServiceClient interface {
	RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*Empty, error)
}

type registrationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRegistrationServiceClient(cc grpc.ClientConnInterface) RegistrationServiceClient {
	return &registrationServiceClient{cc}
}

func (c *registrationServiceClient) RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/gophkeeper.RegistrationService/RegisterUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegistrationServiceServer is the server API for RegistrationService service.
// All implementations must embed UnimplementedRegistrationServiceServer
// for forward compatibility
type RegistrationServiceServer interface {
	RegisterUser(context.Context, *RegisterUserRequest) (*Empty, error)
	mustEmbedUnimplementedRegistrationServiceServer()
}

// UnimplementedRegistrationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRegistrationServiceServer struct {
}

func (UnimplementedRegistrationServiceServer) RegisterUser(context.Context, *RegisterUserRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (UnimplementedRegistrationServiceServer) mustEmbedUnimplementedRegistrationServiceServer() {}

// UnsafeRegistrationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RegistrationServiceServer will
// result in compilation errors.
type UnsafeRegistrationServiceServer interface {
	mustEmbedUnimplementedRegistrationServiceServer()
}

func RegisterRegistrationServiceServer(s grpc.ServiceRegistrar, srv RegistrationServiceServer) {
	s.RegisterService(&RegistrationService_ServiceDesc, srv)
}

func _RegistrationService_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationServiceServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gophkeeper.RegistrationService/RegisterUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationServiceServer).RegisterUser(ctx, req.(*RegisterUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegistrationService_ServiceDesc is the grpc.ServiceDesc for RegistrationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RegistrationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gophkeeper.RegistrationService",
	HandlerType: (*RegistrationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterUser",
			Handler:    _RegistrationService_RegisterUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/api.proto",
}

// DataServiceClient is the client API for DataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataServiceClient interface {
	AddLoginPasswordPair(ctx context.Context, in *AddLoginPasswordPairRequest, opts ...grpc.CallOption) (*Empty, error)
	AddTextData(ctx context.Context, in *AddTextDataRequest, opts ...grpc.CallOption) (*Empty, error)
	AddBinaryData(ctx context.Context, in *AddBinaryDataRequest, opts ...grpc.CallOption) (*Empty, error)
	AddBankCardDetail(ctx context.Context, in *AddBankCardDetailRequest, opts ...grpc.CallOption) (*Empty, error)
	GetLoginPasswordPairs(ctx context.Context, in *GetLoginPasswordPairsRequest, opts ...grpc.CallOption) (*GetLoginPasswordPairsResponse, error)
	GetTextData(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetTextDataResponse, error)
	GetBinaryData(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetBinaryDataResponse, error)
	GetBankCardDetails(ctx context.Context, in *GetBankCardDetailsRequest, opts ...grpc.CallOption) (*GetBankCardDetailsResponse, error)
}

type dataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataServiceClient(cc grpc.ClientConnInterface) DataServiceClient {
	return &dataServiceClient{cc}
}

func (c *dataServiceClient) AddLoginPasswordPair(ctx context.Context, in *AddLoginPasswordPairRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/gophkeeper.DataService/AddLoginPasswordPair", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataServiceClient) AddTextData(ctx context.Context, in *AddTextDataRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/gophkeeper.DataService/AddTextData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataServiceClient) AddBinaryData(ctx context.Context, in *AddBinaryDataRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/gophkeeper.DataService/AddBinaryData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataServiceClient) AddBankCardDetail(ctx context.Context, in *AddBankCardDetailRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/gophkeeper.DataService/AddBankCardDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataServiceClient) GetLoginPasswordPairs(ctx context.Context, in *GetLoginPasswordPairsRequest, opts ...grpc.CallOption) (*GetLoginPasswordPairsResponse, error) {
	out := new(GetLoginPasswordPairsResponse)
	err := c.cc.Invoke(ctx, "/gophkeeper.DataService/GetLoginPasswordPairs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataServiceClient) GetTextData(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetTextDataResponse, error) {
	out := new(GetTextDataResponse)
	err := c.cc.Invoke(ctx, "/gophkeeper.DataService/GetTextData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataServiceClient) GetBinaryData(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetBinaryDataResponse, error) {
	out := new(GetBinaryDataResponse)
	err := c.cc.Invoke(ctx, "/gophkeeper.DataService/GetBinaryData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataServiceClient) GetBankCardDetails(ctx context.Context, in *GetBankCardDetailsRequest, opts ...grpc.CallOption) (*GetBankCardDetailsResponse, error) {
	out := new(GetBankCardDetailsResponse)
	err := c.cc.Invoke(ctx, "/gophkeeper.DataService/GetBankCardDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataServiceServer is the server API for DataService service.
// All implementations must embed UnimplementedDataServiceServer
// for forward compatibility
type DataServiceServer interface {
	AddLoginPasswordPair(context.Context, *AddLoginPasswordPairRequest) (*Empty, error)
	AddTextData(context.Context, *AddTextDataRequest) (*Empty, error)
	AddBinaryData(context.Context, *AddBinaryDataRequest) (*Empty, error)
	AddBankCardDetail(context.Context, *AddBankCardDetailRequest) (*Empty, error)
	GetLoginPasswordPairs(context.Context, *GetLoginPasswordPairsRequest) (*GetLoginPasswordPairsResponse, error)
	GetTextData(context.Context, *Empty) (*GetTextDataResponse, error)
	GetBinaryData(context.Context, *Empty) (*GetBinaryDataResponse, error)
	GetBankCardDetails(context.Context, *GetBankCardDetailsRequest) (*GetBankCardDetailsResponse, error)
	mustEmbedUnimplementedDataServiceServer()
}

// UnimplementedDataServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDataServiceServer struct {
}

func (UnimplementedDataServiceServer) AddLoginPasswordPair(context.Context, *AddLoginPasswordPairRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddLoginPasswordPair not implemented")
}
func (UnimplementedDataServiceServer) AddTextData(context.Context, *AddTextDataRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTextData not implemented")
}
func (UnimplementedDataServiceServer) AddBinaryData(context.Context, *AddBinaryDataRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBinaryData not implemented")
}
func (UnimplementedDataServiceServer) AddBankCardDetail(context.Context, *AddBankCardDetailRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBankCardDetail not implemented")
}
func (UnimplementedDataServiceServer) GetLoginPasswordPairs(context.Context, *GetLoginPasswordPairsRequest) (*GetLoginPasswordPairsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLoginPasswordPairs not implemented")
}
func (UnimplementedDataServiceServer) GetTextData(context.Context, *Empty) (*GetTextDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTextData not implemented")
}
func (UnimplementedDataServiceServer) GetBinaryData(context.Context, *Empty) (*GetBinaryDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBinaryData not implemented")
}
func (UnimplementedDataServiceServer) GetBankCardDetails(context.Context, *GetBankCardDetailsRequest) (*GetBankCardDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBankCardDetails not implemented")
}
func (UnimplementedDataServiceServer) mustEmbedUnimplementedDataServiceServer() {}

// UnsafeDataServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataServiceServer will
// result in compilation errors.
type UnsafeDataServiceServer interface {
	mustEmbedUnimplementedDataServiceServer()
}

func RegisterDataServiceServer(s grpc.ServiceRegistrar, srv DataServiceServer) {
	s.RegisterService(&DataService_ServiceDesc, srv)
}

func _DataService_AddLoginPasswordPair_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddLoginPasswordPairRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceServer).AddLoginPasswordPair(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gophkeeper.DataService/AddLoginPasswordPair",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceServer).AddLoginPasswordPair(ctx, req.(*AddLoginPasswordPairRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataService_AddTextData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTextDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceServer).AddTextData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gophkeeper.DataService/AddTextData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceServer).AddTextData(ctx, req.(*AddTextDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataService_AddBinaryData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBinaryDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceServer).AddBinaryData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gophkeeper.DataService/AddBinaryData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceServer).AddBinaryData(ctx, req.(*AddBinaryDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataService_AddBankCardDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBankCardDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceServer).AddBankCardDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gophkeeper.DataService/AddBankCardDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceServer).AddBankCardDetail(ctx, req.(*AddBankCardDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataService_GetLoginPasswordPairs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLoginPasswordPairsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceServer).GetLoginPasswordPairs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gophkeeper.DataService/GetLoginPasswordPairs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceServer).GetLoginPasswordPairs(ctx, req.(*GetLoginPasswordPairsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataService_GetTextData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceServer).GetTextData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gophkeeper.DataService/GetTextData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceServer).GetTextData(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataService_GetBinaryData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceServer).GetBinaryData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gophkeeper.DataService/GetBinaryData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceServer).GetBinaryData(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataService_GetBankCardDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBankCardDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceServer).GetBankCardDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gophkeeper.DataService/GetBankCardDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceServer).GetBankCardDetails(ctx, req.(*GetBankCardDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DataService_ServiceDesc is the grpc.ServiceDesc for DataService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gophkeeper.DataService",
	HandlerType: (*DataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddLoginPasswordPair",
			Handler:    _DataService_AddLoginPasswordPair_Handler,
		},
		{
			MethodName: "AddTextData",
			Handler:    _DataService_AddTextData_Handler,
		},
		{
			MethodName: "AddBinaryData",
			Handler:    _DataService_AddBinaryData_Handler,
		},
		{
			MethodName: "AddBankCardDetail",
			Handler:    _DataService_AddBankCardDetail_Handler,
		},
		{
			MethodName: "GetLoginPasswordPairs",
			Handler:    _DataService_GetLoginPasswordPairs_Handler,
		},
		{
			MethodName: "GetTextData",
			Handler:    _DataService_GetTextData_Handler,
		},
		{
			MethodName: "GetBinaryData",
			Handler:    _DataService_GetBinaryData_Handler,
		},
		{
			MethodName: "GetBankCardDetails",
			Handler:    _DataService_GetBankCardDetails_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/api.proto",
}
