// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v3.21.6
// source: userapi/v1/user.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationUserChangePassword = "/userapi.v1.User/ChangePassword"
const OperationUserDeleteUser = "/userapi.v1.User/DeleteUser"
const OperationUserLogin = "/userapi.v1.User/Login"
const OperationUserRegister = "/userapi.v1.User/Register"
const OperationUserSendVerificationCode = "/userapi.v1.User/SendVerificationCode"

type UserHTTPServer interface {
	ChangePassword(context.Context, *ChangePasswordRequest) (*ChangePasswordReply, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserReply, error)
	// Login 用户登录
	Login(context.Context, *LoginRequest) (*UserReply, error)
	// Register 用户注册
	Register(context.Context, *RegisterRequest) (*UserReply, error)
	// SendVerificationCode 发送邮箱验证码
	SendVerificationCode(context.Context, *SendVerificationCodeRequest) (*SendVerificationCodeReply, error)
}

func RegisterUserHTTPServer(s *http.Server, srv UserHTTPServer) {
	r := s.Route("/")
	r.POST("/register", _User_Register0_HTTP_Handler(srv))
	r.POST("/login", _User_Login0_HTTP_Handler(srv))
	r.POST("/send_verification_code", _User_SendVerificationCode0_HTTP_Handler(srv))
	r.POST("/change_password", _User_ChangePassword0_HTTP_Handler(srv))
	r.POST("/delete_user", _User_DeleteUser0_HTTP_Handler(srv))
}

func _User_Register0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserRegister)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Register(ctx, req.(*RegisterRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserReply)
		return ctx.Result(200, reply)
	}
}

func _User_Login0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserReply)
		return ctx.Result(200, reply)
	}
}

func _User_SendVerificationCode0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SendVerificationCodeRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserSendVerificationCode)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SendVerificationCode(ctx, req.(*SendVerificationCodeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SendVerificationCodeReply)
		return ctx.Result(200, reply)
	}
}

func _User_ChangePassword0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ChangePasswordRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserChangePassword)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ChangePassword(ctx, req.(*ChangePasswordRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ChangePasswordReply)
		return ctx.Result(200, reply)
	}
}

func _User_DeleteUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserDeleteUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteUser(ctx, req.(*DeleteUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteUserReply)
		return ctx.Result(200, reply)
	}
}

type UserHTTPClient interface {
	ChangePassword(ctx context.Context, req *ChangePasswordRequest, opts ...http.CallOption) (rsp *ChangePasswordReply, err error)
	DeleteUser(ctx context.Context, req *DeleteUserRequest, opts ...http.CallOption) (rsp *DeleteUserReply, err error)
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *UserReply, err error)
	Register(ctx context.Context, req *RegisterRequest, opts ...http.CallOption) (rsp *UserReply, err error)
	SendVerificationCode(ctx context.Context, req *SendVerificationCodeRequest, opts ...http.CallOption) (rsp *SendVerificationCodeReply, err error)
}

type UserHTTPClientImpl struct {
	cc *http.Client
}

func NewUserHTTPClient(client *http.Client) UserHTTPClient {
	return &UserHTTPClientImpl{client}
}

func (c *UserHTTPClientImpl) ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...http.CallOption) (*ChangePasswordReply, error) {
	var out ChangePasswordReply
	pattern := "/change_password"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserChangePassword))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserHTTPClientImpl) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...http.CallOption) (*DeleteUserReply, error) {
	var out DeleteUserReply
	pattern := "/delete_user"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserDeleteUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserHTTPClientImpl) Login(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*UserReply, error) {
	var out UserReply
	pattern := "/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserHTTPClientImpl) Register(ctx context.Context, in *RegisterRequest, opts ...http.CallOption) (*UserReply, error) {
	var out UserReply
	pattern := "/register"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserRegister))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserHTTPClientImpl) SendVerificationCode(ctx context.Context, in *SendVerificationCodeRequest, opts ...http.CallOption) (*SendVerificationCodeReply, error) {
	var out SendVerificationCodeReply
	pattern := "/send_verification_code"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserSendVerificationCode))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
