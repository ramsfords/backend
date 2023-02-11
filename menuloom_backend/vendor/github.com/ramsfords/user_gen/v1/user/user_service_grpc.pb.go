// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: user_service.proto

package user

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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	PingUserService(ctx context.Context, in *UserServicePingData, opts ...grpc.CallOption) (*UserServicePingData, error)
	SignUp(ctx context.Context, in *SignUpData, opts ...grpc.CallOption) (*Ok, error)
	LogIn(ctx context.Context, in *LoginUserData, opts ...grpc.CallOption) (*OkWithData, error)
	LogOut(ctx context.Context, in *LogOutData, opts ...grpc.CallOption) (*Ok, error)
	ForgotPassword(ctx context.Context, in *ForgotPasswordData, opts ...grpc.CallOption) (*Ok, error)
	ResetPassword(ctx context.Context, in *ResetPasswordData, opts ...grpc.CallOption) (*OkWithData, error)
	UpdateStaffRole(ctx context.Context, in *UpdateUserRoleData, opts ...grpc.CallOption) (*OkWithData, error)
	AddUser(ctx context.Context, in *AddStaffData, opts ...grpc.CallOption) (*OkWithData, error)
	UpdateUser(ctx context.Context, in *UserData, opts ...grpc.CallOption) (*OkWithData, error)
	GetMe(ctx context.Context, in *MeData, opts ...grpc.CallOption) (*UserData, error)
	ConfirmEmail(ctx context.Context, in *ConfirmEmailData, opts ...grpc.CallOption) (*OkWithData, error)
	Home(ctx context.Context, in *UserHomeData, opts ...grpc.CallOption) (*OkWithData, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) PingUserService(ctx context.Context, in *UserServicePingData, opts ...grpc.CallOption) (*UserServicePingData, error) {
	out := new(UserServicePingData)
	err := c.cc.Invoke(ctx, "/user.UserService/PingUserService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SignUp(ctx context.Context, in *SignUpData, opts ...grpc.CallOption) (*Ok, error) {
	out := new(Ok)
	err := c.cc.Invoke(ctx, "/user.UserService/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) LogIn(ctx context.Context, in *LoginUserData, opts ...grpc.CallOption) (*OkWithData, error) {
	out := new(OkWithData)
	err := c.cc.Invoke(ctx, "/user.UserService/LogIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) LogOut(ctx context.Context, in *LogOutData, opts ...grpc.CallOption) (*Ok, error) {
	out := new(Ok)
	err := c.cc.Invoke(ctx, "/user.UserService/LogOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ForgotPassword(ctx context.Context, in *ForgotPasswordData, opts ...grpc.CallOption) (*Ok, error) {
	out := new(Ok)
	err := c.cc.Invoke(ctx, "/user.UserService/ForgotPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ResetPassword(ctx context.Context, in *ResetPasswordData, opts ...grpc.CallOption) (*OkWithData, error) {
	out := new(OkWithData)
	err := c.cc.Invoke(ctx, "/user.UserService/ResetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateStaffRole(ctx context.Context, in *UpdateUserRoleData, opts ...grpc.CallOption) (*OkWithData, error) {
	out := new(OkWithData)
	err := c.cc.Invoke(ctx, "/user.UserService/UpdateStaffRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AddUser(ctx context.Context, in *AddStaffData, opts ...grpc.CallOption) (*OkWithData, error) {
	out := new(OkWithData)
	err := c.cc.Invoke(ctx, "/user.UserService/AddUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUser(ctx context.Context, in *UserData, opts ...grpc.CallOption) (*OkWithData, error) {
	out := new(OkWithData)
	err := c.cc.Invoke(ctx, "/user.UserService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetMe(ctx context.Context, in *MeData, opts ...grpc.CallOption) (*UserData, error) {
	out := new(UserData)
	err := c.cc.Invoke(ctx, "/user.UserService/GetMe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ConfirmEmail(ctx context.Context, in *ConfirmEmailData, opts ...grpc.CallOption) (*OkWithData, error) {
	out := new(OkWithData)
	err := c.cc.Invoke(ctx, "/user.UserService/ConfirmEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Home(ctx context.Context, in *UserHomeData, opts ...grpc.CallOption) (*OkWithData, error) {
	out := new(OkWithData)
	err := c.cc.Invoke(ctx, "/user.UserService/Home", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations should embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	PingUserService(context.Context, *UserServicePingData) (*UserServicePingData, error)
	SignUp(context.Context, *SignUpData) (*Ok, error)
	LogIn(context.Context, *LoginUserData) (*OkWithData, error)
	LogOut(context.Context, *LogOutData) (*Ok, error)
	ForgotPassword(context.Context, *ForgotPasswordData) (*Ok, error)
	ResetPassword(context.Context, *ResetPasswordData) (*OkWithData, error)
	UpdateStaffRole(context.Context, *UpdateUserRoleData) (*OkWithData, error)
	AddUser(context.Context, *AddStaffData) (*OkWithData, error)
	UpdateUser(context.Context, *UserData) (*OkWithData, error)
	GetMe(context.Context, *MeData) (*UserData, error)
	ConfirmEmail(context.Context, *ConfirmEmailData) (*OkWithData, error)
	Home(context.Context, *UserHomeData) (*OkWithData, error)
}

// UnimplementedUserServiceServer should be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) PingUserService(context.Context, *UserServicePingData) (*UserServicePingData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PingUserService not implemented")
}
func (UnimplementedUserServiceServer) SignUp(context.Context, *SignUpData) (*Ok, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (UnimplementedUserServiceServer) LogIn(context.Context, *LoginUserData) (*OkWithData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogIn not implemented")
}
func (UnimplementedUserServiceServer) LogOut(context.Context, *LogOutData) (*Ok, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogOut not implemented")
}
func (UnimplementedUserServiceServer) ForgotPassword(context.Context, *ForgotPasswordData) (*Ok, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForgotPassword not implemented")
}
func (UnimplementedUserServiceServer) ResetPassword(context.Context, *ResetPasswordData) (*OkWithData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}
func (UnimplementedUserServiceServer) UpdateStaffRole(context.Context, *UpdateUserRoleData) (*OkWithData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStaffRole not implemented")
}
func (UnimplementedUserServiceServer) AddUser(context.Context, *AddStaffData) (*OkWithData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}
func (UnimplementedUserServiceServer) UpdateUser(context.Context, *UserData) (*OkWithData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUserServiceServer) GetMe(context.Context, *MeData) (*UserData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMe not implemented")
}
func (UnimplementedUserServiceServer) ConfirmEmail(context.Context, *ConfirmEmailData) (*OkWithData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmEmail not implemented")
}
func (UnimplementedUserServiceServer) Home(context.Context, *UserHomeData) (*OkWithData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Home not implemented")
}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_PingUserService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserServicePingData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).PingUserService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/PingUserService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).PingUserService(ctx, req.(*UserServicePingData))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SignUp(ctx, req.(*SignUpData))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_LogIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginUserData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).LogIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/LogIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).LogIn(ctx, req.(*LoginUserData))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_LogOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogOutData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).LogOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/LogOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).LogOut(ctx, req.(*LogOutData))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ForgotPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForgotPasswordData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ForgotPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/ForgotPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ForgotPassword(ctx, req.(*ForgotPasswordData))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetPasswordData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/ResetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ResetPassword(ctx, req.(*ResetPasswordData))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateStaffRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRoleData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateStaffRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UpdateStaffRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateStaffRole(ctx, req.(*UpdateUserRoleData))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddStaffData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddUser(ctx, req.(*AddStaffData))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUser(ctx, req.(*UserData))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetMe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MeData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetMe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetMe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetMe(ctx, req.(*MeData))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ConfirmEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmEmailData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ConfirmEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/ConfirmEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ConfirmEmail(ctx, req.(*ConfirmEmailData))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Home_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserHomeData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Home(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/Home",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Home(ctx, req.(*UserHomeData))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PingUserService",
			Handler:    _UserService_PingUserService_Handler,
		},
		{
			MethodName: "SignUp",
			Handler:    _UserService_SignUp_Handler,
		},
		{
			MethodName: "LogIn",
			Handler:    _UserService_LogIn_Handler,
		},
		{
			MethodName: "LogOut",
			Handler:    _UserService_LogOut_Handler,
		},
		{
			MethodName: "ForgotPassword",
			Handler:    _UserService_ForgotPassword_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _UserService_ResetPassword_Handler,
		},
		{
			MethodName: "UpdateStaffRole",
			Handler:    _UserService_UpdateStaffRole_Handler,
		},
		{
			MethodName: "AddUser",
			Handler:    _UserService_AddUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserService_UpdateUser_Handler,
		},
		{
			MethodName: "GetMe",
			Handler:    _UserService_GetMe_Handler,
		},
		{
			MethodName: "ConfirmEmail",
			Handler:    _UserService_ConfirmEmail_Handler,
		},
		{
			MethodName: "Home",
			Handler:    _UserService_Home_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_service.proto",
}
