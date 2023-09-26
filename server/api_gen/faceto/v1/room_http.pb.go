// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.2
// - protoc             v4.23.3
// source: faceto/v1/room.proto

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

const OperationRoomAuth = "/faceto.v1.Room/Auth"
const OperationRoomAuthCreate = "/faceto.v1.Room/AuthCreate"
const OperationRoomAuthExchange = "/faceto.v1.Room/AuthExchange"
const OperationRoomLink = "/faceto.v1.Room/Link"
const OperationRoomToken = "/faceto.v1.Room/Token"

type RoomHTTPServer interface {
	// Auth OAuth Token API
	Auth(context.Context, *AuthRequest) (*AuthReply, error)
	// AuthCreate OAuth Create API
	AuthCreate(context.Context, *AuthCreateRequest) (*AuthCreateReply, error)
	// AuthExchange OAuth Token Exchange API
	AuthExchange(context.Context, *AuthExchangeRequest) (*AuthReply, error)
	// Link Token Get Room Link
	Link(context.Context, *RoomLinkRequest) (*RoomLinkReply, error)
	// Token Token Create API Auth Token
	Token(context.Context, *TokenRequest) (*TokenReply, error)
}

func RegisterRoomHTTPServer(s *http.Server, srv RoomHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/token", _Room_Token0_HTTP_Handler(srv))
	r.POST("/v1/room/link", _Room_Link0_HTTP_Handler(srv))
	r.POST("/v1/auth/create", _Room_AuthCreate0_HTTP_Handler(srv))
	r.POST("/v1/auth", _Room_Auth0_HTTP_Handler(srv))
	r.POST("/v1/auth/exchange", _Room_AuthExchange0_HTTP_Handler(srv))
}

func _Room_Token0_HTTP_Handler(srv RoomHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TokenRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRoomToken)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Token(ctx, req.(*TokenRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*TokenReply)
		return ctx.Result(200, reply)
	}
}

func _Room_Link0_HTTP_Handler(srv RoomHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RoomLinkRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRoomLink)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Link(ctx, req.(*RoomLinkRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RoomLinkReply)
		return ctx.Result(200, reply)
	}
}

func _Room_AuthCreate0_HTTP_Handler(srv RoomHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AuthCreateRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRoomAuthCreate)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AuthCreate(ctx, req.(*AuthCreateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AuthCreateReply)
		return ctx.Result(200, reply)
	}
}

func _Room_Auth0_HTTP_Handler(srv RoomHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AuthRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRoomAuth)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Auth(ctx, req.(*AuthRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AuthReply)
		return ctx.Result(200, reply)
	}
}

func _Room_AuthExchange0_HTTP_Handler(srv RoomHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AuthExchangeRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRoomAuthExchange)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AuthExchange(ctx, req.(*AuthExchangeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AuthReply)
		return ctx.Result(200, reply)
	}
}

type RoomHTTPClient interface {
	Auth(ctx context.Context, req *AuthRequest, opts ...http.CallOption) (rsp *AuthReply, err error)
	AuthCreate(ctx context.Context, req *AuthCreateRequest, opts ...http.CallOption) (rsp *AuthCreateReply, err error)
	AuthExchange(ctx context.Context, req *AuthExchangeRequest, opts ...http.CallOption) (rsp *AuthReply, err error)
	Link(ctx context.Context, req *RoomLinkRequest, opts ...http.CallOption) (rsp *RoomLinkReply, err error)
	Token(ctx context.Context, req *TokenRequest, opts ...http.CallOption) (rsp *TokenReply, err error)
}

type RoomHTTPClientImpl struct {
	cc *http.Client
}

func NewRoomHTTPClient(client *http.Client) RoomHTTPClient {
	return &RoomHTTPClientImpl{client}
}

func (c *RoomHTTPClientImpl) Auth(ctx context.Context, in *AuthRequest, opts ...http.CallOption) (*AuthReply, error) {
	var out AuthReply
	pattern := "/v1/auth"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRoomAuth))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RoomHTTPClientImpl) AuthCreate(ctx context.Context, in *AuthCreateRequest, opts ...http.CallOption) (*AuthCreateReply, error) {
	var out AuthCreateReply
	pattern := "/v1/auth/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRoomAuthCreate))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RoomHTTPClientImpl) AuthExchange(ctx context.Context, in *AuthExchangeRequest, opts ...http.CallOption) (*AuthReply, error) {
	var out AuthReply
	pattern := "/v1/auth/exchange"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRoomAuthExchange))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RoomHTTPClientImpl) Link(ctx context.Context, in *RoomLinkRequest, opts ...http.CallOption) (*RoomLinkReply, error) {
	var out RoomLinkReply
	pattern := "/v1/room/link"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRoomLink))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RoomHTTPClientImpl) Token(ctx context.Context, in *TokenRequest, opts ...http.CallOption) (*TokenReply, error) {
	var out TokenReply
	pattern := "/v1/token"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRoomToken))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
