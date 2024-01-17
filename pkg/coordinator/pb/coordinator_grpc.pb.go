// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: coordinator.proto

package __

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

const (
	Coordinator_CoordinatorSignupRequest_FullMethodName        = "/pb.Coordinator/CoordinatorSignupRequest"
	Coordinator_CoordinatorSignupVerifyRequest_FullMethodName  = "/pb.Coordinator/CoordinatorSignupVerifyRequest"
	Coordinator_CoordinatorLoginRequest_FullMethodName         = "/pb.Coordinator/CoordinatorLoginRequest"
	Coordinator_CoordinatorAddPackage_FullMethodName           = "/pb.Coordinator/CoordinatorAddPackage"
	Coordinator_CoordinatorAddDestination_FullMethodName       = "/pb.Coordinator/CoordinatorAddDestination"
	Coordinator_CoordinatorAddActivity_FullMethodName          = "/pb.Coordinator/CoordinatorAddActivity"
	Coordinator_CoordinatorViewPackage_FullMethodName          = "/pb.Coordinator/CoordinatorViewPackage"
	Coordinator_CoordinatorViewDestination_FullMethodName      = "/pb.Coordinator/CoordinatorViewDestination"
	Coordinator_CoordinatorViewActivity_FullMethodName         = "/pb.Coordinator/CoordinatorViewActivity"
	Coordinator_CoordinatorForgetPassword_FullMethodName       = "/pb.Coordinator/CoordinatorForgetPassword"
	Coordinator_CoordinatorForgetPasswordVerify_FullMethodName = "/pb.Coordinator/CoordinatorForgetPasswordVerify"
	Coordinator_CoordinatorNewPassword_FullMethodName          = "/pb.Coordinator/CoordinatorNewPassword"
	Coordinator_AvailablePackages_FullMethodName               = "/pb.Coordinator/AvailablePackages"
	Coordinator_AddCatagory_FullMethodName                     = "/pb.Coordinator/AddCatagory"
	Coordinator_AdminAvailablePackages_FullMethodName          = "/pb.Coordinator/AdminAvailablePackages"
	Coordinator_AdminPacakgeStatus_FullMethodName              = "/pb.Coordinator/AdminPacakgeStatus"
)

// CoordinatorClient is the client API for Coordinator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoordinatorClient interface {
	CoordinatorSignupRequest(ctx context.Context, in *Signup, opts ...grpc.CallOption) (*Responce, error)
	CoordinatorSignupVerifyRequest(ctx context.Context, in *Verify, opts ...grpc.CallOption) (*Responce, error)
	CoordinatorLoginRequest(ctx context.Context, in *Login, opts ...grpc.CallOption) (*LoginResponce, error)
	CoordinatorAddPackage(ctx context.Context, in *Package, opts ...grpc.CallOption) (*Responce, error)
	CoordinatorAddDestination(ctx context.Context, in *Destination, opts ...grpc.CallOption) (*Responce, error)
	CoordinatorAddActivity(ctx context.Context, in *Activity, opts ...grpc.CallOption) (*Responce, error)
	CoordinatorViewPackage(ctx context.Context, in *View, opts ...grpc.CallOption) (*Package, error)
	CoordinatorViewDestination(ctx context.Context, in *View, opts ...grpc.CallOption) (*Destination, error)
	CoordinatorViewActivity(ctx context.Context, in *View, opts ...grpc.CallOption) (*Activity, error)
	CoordinatorForgetPassword(ctx context.Context, in *ForgetPassword, opts ...grpc.CallOption) (*Responce, error)
	CoordinatorForgetPasswordVerify(ctx context.Context, in *ForgetPasswordVerify, opts ...grpc.CallOption) (*Responce, error)
	CoordinatorNewPassword(ctx context.Context, in *Newpassword, opts ...grpc.CallOption) (*Responce, error)
	AvailablePackages(ctx context.Context, in *Packages, opts ...grpc.CallOption) (*PackagesResponce, error)
	AddCatagory(ctx context.Context, in *Category, opts ...grpc.CallOption) (*Responce, error)
	AdminAvailablePackages(ctx context.Context, in *Packages, opts ...grpc.CallOption) (*PackagesResponce, error)
	AdminPacakgeStatus(ctx context.Context, in *View, opts ...grpc.CallOption) (*Responce, error)
}

type coordinatorClient struct {
	cc grpc.ClientConnInterface
}

func NewCoordinatorClient(cc grpc.ClientConnInterface) CoordinatorClient {
	return &coordinatorClient{cc}
}

func (c *coordinatorClient) CoordinatorSignupRequest(ctx context.Context, in *Signup, opts ...grpc.CallOption) (*Responce, error) {
	out := new(Responce)
	err := c.cc.Invoke(ctx, Coordinator_CoordinatorSignupRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coordinatorClient) CoordinatorSignupVerifyRequest(ctx context.Context, in *Verify, opts ...grpc.CallOption) (*Responce, error) {
	out := new(Responce)
	err := c.cc.Invoke(ctx, Coordinator_CoordinatorSignupVerifyRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coordinatorClient) CoordinatorLoginRequest(ctx context.Context, in *Login, opts ...grpc.CallOption) (*LoginResponce, error) {
	out := new(LoginResponce)
	err := c.cc.Invoke(ctx, Coordinator_CoordinatorLoginRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coordinatorClient) CoordinatorAddPackage(ctx context.Context, in *Package, opts ...grpc.CallOption) (*Responce, error) {
	out := new(Responce)
	err := c.cc.Invoke(ctx, Coordinator_CoordinatorAddPackage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coordinatorClient) CoordinatorAddDestination(ctx context.Context, in *Destination, opts ...grpc.CallOption) (*Responce, error) {
	out := new(Responce)
	err := c.cc.Invoke(ctx, Coordinator_CoordinatorAddDestination_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coordinatorClient) CoordinatorAddActivity(ctx context.Context, in *Activity, opts ...grpc.CallOption) (*Responce, error) {
	out := new(Responce)
	err := c.cc.Invoke(ctx, Coordinator_CoordinatorAddActivity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coordinatorClient) CoordinatorViewPackage(ctx context.Context, in *View, opts ...grpc.CallOption) (*Package, error) {
	out := new(Package)
	err := c.cc.Invoke(ctx, Coordinator_CoordinatorViewPackage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coordinatorClient) CoordinatorViewDestination(ctx context.Context, in *View, opts ...grpc.CallOption) (*Destination, error) {
	out := new(Destination)
	err := c.cc.Invoke(ctx, Coordinator_CoordinatorViewDestination_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coordinatorClient) CoordinatorViewActivity(ctx context.Context, in *View, opts ...grpc.CallOption) (*Activity, error) {
	out := new(Activity)
	err := c.cc.Invoke(ctx, Coordinator_CoordinatorViewActivity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coordinatorClient) CoordinatorForgetPassword(ctx context.Context, in *ForgetPassword, opts ...grpc.CallOption) (*Responce, error) {
	out := new(Responce)
	err := c.cc.Invoke(ctx, Coordinator_CoordinatorForgetPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coordinatorClient) CoordinatorForgetPasswordVerify(ctx context.Context, in *ForgetPasswordVerify, opts ...grpc.CallOption) (*Responce, error) {
	out := new(Responce)
	err := c.cc.Invoke(ctx, Coordinator_CoordinatorForgetPasswordVerify_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coordinatorClient) CoordinatorNewPassword(ctx context.Context, in *Newpassword, opts ...grpc.CallOption) (*Responce, error) {
	out := new(Responce)
	err := c.cc.Invoke(ctx, Coordinator_CoordinatorNewPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coordinatorClient) AvailablePackages(ctx context.Context, in *Packages, opts ...grpc.CallOption) (*PackagesResponce, error) {
	out := new(PackagesResponce)
	err := c.cc.Invoke(ctx, Coordinator_AvailablePackages_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coordinatorClient) AddCatagory(ctx context.Context, in *Category, opts ...grpc.CallOption) (*Responce, error) {
	out := new(Responce)
	err := c.cc.Invoke(ctx, Coordinator_AddCatagory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coordinatorClient) AdminAvailablePackages(ctx context.Context, in *Packages, opts ...grpc.CallOption) (*PackagesResponce, error) {
	out := new(PackagesResponce)
	err := c.cc.Invoke(ctx, Coordinator_AdminAvailablePackages_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coordinatorClient) AdminPacakgeStatus(ctx context.Context, in *View, opts ...grpc.CallOption) (*Responce, error) {
	out := new(Responce)
	err := c.cc.Invoke(ctx, Coordinator_AdminPacakgeStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoordinatorServer is the server API for Coordinator service.
// All implementations must embed UnimplementedCoordinatorServer
// for forward compatibility
type CoordinatorServer interface {
	CoordinatorSignupRequest(context.Context, *Signup) (*Responce, error)
	CoordinatorSignupVerifyRequest(context.Context, *Verify) (*Responce, error)
	CoordinatorLoginRequest(context.Context, *Login) (*LoginResponce, error)
	CoordinatorAddPackage(context.Context, *Package) (*Responce, error)
	CoordinatorAddDestination(context.Context, *Destination) (*Responce, error)
	CoordinatorAddActivity(context.Context, *Activity) (*Responce, error)
	CoordinatorViewPackage(context.Context, *View) (*Package, error)
	CoordinatorViewDestination(context.Context, *View) (*Destination, error)
	CoordinatorViewActivity(context.Context, *View) (*Activity, error)
	CoordinatorForgetPassword(context.Context, *ForgetPassword) (*Responce, error)
	CoordinatorForgetPasswordVerify(context.Context, *ForgetPasswordVerify) (*Responce, error)
	CoordinatorNewPassword(context.Context, *Newpassword) (*Responce, error)
	AvailablePackages(context.Context, *Packages) (*PackagesResponce, error)
	AddCatagory(context.Context, *Category) (*Responce, error)
	AdminAvailablePackages(context.Context, *Packages) (*PackagesResponce, error)
	AdminPacakgeStatus(context.Context, *View) (*Responce, error)
	mustEmbedUnimplementedCoordinatorServer()
}

// UnimplementedCoordinatorServer must be embedded to have forward compatible implementations.
type UnimplementedCoordinatorServer struct {
}

func (UnimplementedCoordinatorServer) CoordinatorSignupRequest(context.Context, *Signup) (*Responce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CoordinatorSignupRequest not implemented")
}
func (UnimplementedCoordinatorServer) CoordinatorSignupVerifyRequest(context.Context, *Verify) (*Responce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CoordinatorSignupVerifyRequest not implemented")
}
func (UnimplementedCoordinatorServer) CoordinatorLoginRequest(context.Context, *Login) (*LoginResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CoordinatorLoginRequest not implemented")
}
func (UnimplementedCoordinatorServer) CoordinatorAddPackage(context.Context, *Package) (*Responce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CoordinatorAddPackage not implemented")
}
func (UnimplementedCoordinatorServer) CoordinatorAddDestination(context.Context, *Destination) (*Responce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CoordinatorAddDestination not implemented")
}
func (UnimplementedCoordinatorServer) CoordinatorAddActivity(context.Context, *Activity) (*Responce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CoordinatorAddActivity not implemented")
}
func (UnimplementedCoordinatorServer) CoordinatorViewPackage(context.Context, *View) (*Package, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CoordinatorViewPackage not implemented")
}
func (UnimplementedCoordinatorServer) CoordinatorViewDestination(context.Context, *View) (*Destination, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CoordinatorViewDestination not implemented")
}
func (UnimplementedCoordinatorServer) CoordinatorViewActivity(context.Context, *View) (*Activity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CoordinatorViewActivity not implemented")
}
func (UnimplementedCoordinatorServer) CoordinatorForgetPassword(context.Context, *ForgetPassword) (*Responce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CoordinatorForgetPassword not implemented")
}
func (UnimplementedCoordinatorServer) CoordinatorForgetPasswordVerify(context.Context, *ForgetPasswordVerify) (*Responce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CoordinatorForgetPasswordVerify not implemented")
}
func (UnimplementedCoordinatorServer) CoordinatorNewPassword(context.Context, *Newpassword) (*Responce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CoordinatorNewPassword not implemented")
}
func (UnimplementedCoordinatorServer) AvailablePackages(context.Context, *Packages) (*PackagesResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AvailablePackages not implemented")
}
func (UnimplementedCoordinatorServer) AddCatagory(context.Context, *Category) (*Responce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCatagory not implemented")
}
func (UnimplementedCoordinatorServer) AdminAvailablePackages(context.Context, *Packages) (*PackagesResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminAvailablePackages not implemented")
}
func (UnimplementedCoordinatorServer) AdminPacakgeStatus(context.Context, *View) (*Responce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminPacakgeStatus not implemented")
}
func (UnimplementedCoordinatorServer) mustEmbedUnimplementedCoordinatorServer() {}

// UnsafeCoordinatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CoordinatorServer will
// result in compilation errors.
type UnsafeCoordinatorServer interface {
	mustEmbedUnimplementedCoordinatorServer()
}

func RegisterCoordinatorServer(s grpc.ServiceRegistrar, srv CoordinatorServer) {
	s.RegisterService(&Coordinator_ServiceDesc, srv)
}

func _Coordinator_CoordinatorSignupRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Signup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).CoordinatorSignupRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Coordinator_CoordinatorSignupRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).CoordinatorSignupRequest(ctx, req.(*Signup))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coordinator_CoordinatorSignupVerifyRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Verify)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).CoordinatorSignupVerifyRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Coordinator_CoordinatorSignupVerifyRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).CoordinatorSignupVerifyRequest(ctx, req.(*Verify))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coordinator_CoordinatorLoginRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Login)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).CoordinatorLoginRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Coordinator_CoordinatorLoginRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).CoordinatorLoginRequest(ctx, req.(*Login))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coordinator_CoordinatorAddPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Package)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).CoordinatorAddPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Coordinator_CoordinatorAddPackage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).CoordinatorAddPackage(ctx, req.(*Package))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coordinator_CoordinatorAddDestination_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Destination)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).CoordinatorAddDestination(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Coordinator_CoordinatorAddDestination_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).CoordinatorAddDestination(ctx, req.(*Destination))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coordinator_CoordinatorAddActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Activity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).CoordinatorAddActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Coordinator_CoordinatorAddActivity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).CoordinatorAddActivity(ctx, req.(*Activity))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coordinator_CoordinatorViewPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(View)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).CoordinatorViewPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Coordinator_CoordinatorViewPackage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).CoordinatorViewPackage(ctx, req.(*View))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coordinator_CoordinatorViewDestination_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(View)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).CoordinatorViewDestination(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Coordinator_CoordinatorViewDestination_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).CoordinatorViewDestination(ctx, req.(*View))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coordinator_CoordinatorViewActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(View)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).CoordinatorViewActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Coordinator_CoordinatorViewActivity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).CoordinatorViewActivity(ctx, req.(*View))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coordinator_CoordinatorForgetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForgetPassword)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).CoordinatorForgetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Coordinator_CoordinatorForgetPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).CoordinatorForgetPassword(ctx, req.(*ForgetPassword))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coordinator_CoordinatorForgetPasswordVerify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForgetPasswordVerify)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).CoordinatorForgetPasswordVerify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Coordinator_CoordinatorForgetPasswordVerify_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).CoordinatorForgetPasswordVerify(ctx, req.(*ForgetPasswordVerify))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coordinator_CoordinatorNewPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Newpassword)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).CoordinatorNewPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Coordinator_CoordinatorNewPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).CoordinatorNewPassword(ctx, req.(*Newpassword))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coordinator_AvailablePackages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Packages)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).AvailablePackages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Coordinator_AvailablePackages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).AvailablePackages(ctx, req.(*Packages))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coordinator_AddCatagory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Category)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).AddCatagory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Coordinator_AddCatagory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).AddCatagory(ctx, req.(*Category))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coordinator_AdminAvailablePackages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Packages)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).AdminAvailablePackages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Coordinator_AdminAvailablePackages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).AdminAvailablePackages(ctx, req.(*Packages))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coordinator_AdminPacakgeStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(View)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).AdminPacakgeStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Coordinator_AdminPacakgeStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).AdminPacakgeStatus(ctx, req.(*View))
	}
	return interceptor(ctx, in, info, handler)
}

// Coordinator_ServiceDesc is the grpc.ServiceDesc for Coordinator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Coordinator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Coordinator",
	HandlerType: (*CoordinatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CoordinatorSignupRequest",
			Handler:    _Coordinator_CoordinatorSignupRequest_Handler,
		},
		{
			MethodName: "CoordinatorSignupVerifyRequest",
			Handler:    _Coordinator_CoordinatorSignupVerifyRequest_Handler,
		},
		{
			MethodName: "CoordinatorLoginRequest",
			Handler:    _Coordinator_CoordinatorLoginRequest_Handler,
		},
		{
			MethodName: "CoordinatorAddPackage",
			Handler:    _Coordinator_CoordinatorAddPackage_Handler,
		},
		{
			MethodName: "CoordinatorAddDestination",
			Handler:    _Coordinator_CoordinatorAddDestination_Handler,
		},
		{
			MethodName: "CoordinatorAddActivity",
			Handler:    _Coordinator_CoordinatorAddActivity_Handler,
		},
		{
			MethodName: "CoordinatorViewPackage",
			Handler:    _Coordinator_CoordinatorViewPackage_Handler,
		},
		{
			MethodName: "CoordinatorViewDestination",
			Handler:    _Coordinator_CoordinatorViewDestination_Handler,
		},
		{
			MethodName: "CoordinatorViewActivity",
			Handler:    _Coordinator_CoordinatorViewActivity_Handler,
		},
		{
			MethodName: "CoordinatorForgetPassword",
			Handler:    _Coordinator_CoordinatorForgetPassword_Handler,
		},
		{
			MethodName: "CoordinatorForgetPasswordVerify",
			Handler:    _Coordinator_CoordinatorForgetPasswordVerify_Handler,
		},
		{
			MethodName: "CoordinatorNewPassword",
			Handler:    _Coordinator_CoordinatorNewPassword_Handler,
		},
		{
			MethodName: "AvailablePackages",
			Handler:    _Coordinator_AvailablePackages_Handler,
		},
		{
			MethodName: "AddCatagory",
			Handler:    _Coordinator_AddCatagory_Handler,
		},
		{
			MethodName: "AdminAvailablePackages",
			Handler:    _Coordinator_AdminAvailablePackages_Handler,
		},
		{
			MethodName: "AdminPacakgeStatus",
			Handler:    _Coordinator_AdminPacakgeStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "coordinator.proto",
}
