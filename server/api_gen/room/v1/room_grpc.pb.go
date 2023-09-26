// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: room/v1/room.proto

package v1

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
	Room_Health_FullMethodName               = "/interview.v1.Room/Health"
	Room_CheckRoom_FullMethodName            = "/interview.v1.Room/CheckRoom"
	Room_JoinRoom_FullMethodName             = "/interview.v1.Room/JoinRoom"
	Room_RoomTranscriptOnline_FullMethodName = "/interview.v1.Room/RoomTranscriptOnline"
	Room_RoomTranscript_FullMethodName       = "/interview.v1.Room/RoomTranscript"
	Room_SetRoomVoice_FullMethodName         = "/interview.v1.Room/SetRoomVoice"
)

// RoomClient is the client API for Room service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoomClient interface {
	// Health check
	Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HelloReply, error)
	// CheckRoom check room valid
	CheckRoom(ctx context.Context, in *CheckRoomRequest, opts ...grpc.CallOption) (*CheckRoomReply, error)
	// JoinRoom local join room
	JoinRoom(ctx context.Context, in *JoinRoomRequest, opts ...grpc.CallOption) (*JoinRoomReply, error)
	// RoomTranscript. Get conversation transcripts in real time
	RoomTranscriptOnline(ctx context.Context, in *RoomTranscriptRequest, opts ...grpc.CallOption) (*RoomTranscriptReply, error)
	// RoomTranscript. Get room transcript in real time or cache, compatible with above interface
	RoomTranscript(ctx context.Context, in *RoomTranscriptRequest, opts ...grpc.CallOption) (*RoomTranscriptReply, error)
	// Set Room Voice
	SetRoomVoice(ctx context.Context, in *SetRoomVoiceRequest, opts ...grpc.CallOption) (*NilReply, error)
}

type roomClient struct {
	cc grpc.ClientConnInterface
}

func NewRoomClient(cc grpc.ClientConnInterface) RoomClient {
	return &roomClient{cc}
}

func (c *roomClient) Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, Room_Health_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomClient) CheckRoom(ctx context.Context, in *CheckRoomRequest, opts ...grpc.CallOption) (*CheckRoomReply, error) {
	out := new(CheckRoomReply)
	err := c.cc.Invoke(ctx, Room_CheckRoom_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomClient) JoinRoom(ctx context.Context, in *JoinRoomRequest, opts ...grpc.CallOption) (*JoinRoomReply, error) {
	out := new(JoinRoomReply)
	err := c.cc.Invoke(ctx, Room_JoinRoom_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomClient) RoomTranscriptOnline(ctx context.Context, in *RoomTranscriptRequest, opts ...grpc.CallOption) (*RoomTranscriptReply, error) {
	out := new(RoomTranscriptReply)
	err := c.cc.Invoke(ctx, Room_RoomTranscriptOnline_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomClient) RoomTranscript(ctx context.Context, in *RoomTranscriptRequest, opts ...grpc.CallOption) (*RoomTranscriptReply, error) {
	out := new(RoomTranscriptReply)
	err := c.cc.Invoke(ctx, Room_RoomTranscript_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomClient) SetRoomVoice(ctx context.Context, in *SetRoomVoiceRequest, opts ...grpc.CallOption) (*NilReply, error) {
	out := new(NilReply)
	err := c.cc.Invoke(ctx, Room_SetRoomVoice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoomServer is the server API for Room service.
// All implementations must embed UnimplementedRoomServer
// for forward compatibility
type RoomServer interface {
	// Health check
	Health(context.Context, *HealthRequest) (*HelloReply, error)
	// CheckRoom check room valid
	CheckRoom(context.Context, *CheckRoomRequest) (*CheckRoomReply, error)
	// JoinRoom local join room
	JoinRoom(context.Context, *JoinRoomRequest) (*JoinRoomReply, error)
	// RoomTranscript. Get conversation transcripts in real time
	RoomTranscriptOnline(context.Context, *RoomTranscriptRequest) (*RoomTranscriptReply, error)
	// RoomTranscript. Get room transcript in real time or cache, compatible with above interface
	RoomTranscript(context.Context, *RoomTranscriptRequest) (*RoomTranscriptReply, error)
	// Set Room Voice
	SetRoomVoice(context.Context, *SetRoomVoiceRequest) (*NilReply, error)
	mustEmbedUnimplementedRoomServer()
}

// UnimplementedRoomServer must be embedded to have forward compatible implementations.
type UnimplementedRoomServer struct {
}

func (UnimplementedRoomServer) Health(context.Context, *HealthRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (UnimplementedRoomServer) CheckRoom(context.Context, *CheckRoomRequest) (*CheckRoomReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckRoom not implemented")
}
func (UnimplementedRoomServer) JoinRoom(context.Context, *JoinRoomRequest) (*JoinRoomReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinRoom not implemented")
}
func (UnimplementedRoomServer) RoomTranscriptOnline(context.Context, *RoomTranscriptRequest) (*RoomTranscriptReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoomTranscriptOnline not implemented")
}
func (UnimplementedRoomServer) RoomTranscript(context.Context, *RoomTranscriptRequest) (*RoomTranscriptReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoomTranscript not implemented")
}
func (UnimplementedRoomServer) SetRoomVoice(context.Context, *SetRoomVoiceRequest) (*NilReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetRoomVoice not implemented")
}
func (UnimplementedRoomServer) mustEmbedUnimplementedRoomServer() {}

// UnsafeRoomServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoomServer will
// result in compilation errors.
type UnsafeRoomServer interface {
	mustEmbedUnimplementedRoomServer()
}

func RegisterRoomServer(s grpc.ServiceRegistrar, srv RoomServer) {
	s.RegisterService(&Room_ServiceDesc, srv)
}

func _Room_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Room_Health_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServer).Health(ctx, req.(*HealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Room_CheckRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServer).CheckRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Room_CheckRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServer).CheckRoom(ctx, req.(*CheckRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Room_JoinRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServer).JoinRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Room_JoinRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServer).JoinRoom(ctx, req.(*JoinRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Room_RoomTranscriptOnline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoomTranscriptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServer).RoomTranscriptOnline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Room_RoomTranscriptOnline_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServer).RoomTranscriptOnline(ctx, req.(*RoomTranscriptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Room_RoomTranscript_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoomTranscriptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServer).RoomTranscript(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Room_RoomTranscript_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServer).RoomTranscript(ctx, req.(*RoomTranscriptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Room_SetRoomVoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRoomVoiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServer).SetRoomVoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Room_SetRoomVoice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServer).SetRoomVoice(ctx, req.(*SetRoomVoiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Room_ServiceDesc is the grpc.ServiceDesc for Room service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Room_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "interview.v1.Room",
	HandlerType: (*RoomServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Health",
			Handler:    _Room_Health_Handler,
		},
		{
			MethodName: "CheckRoom",
			Handler:    _Room_CheckRoom_Handler,
		},
		{
			MethodName: "JoinRoom",
			Handler:    _Room_JoinRoom_Handler,
		},
		{
			MethodName: "RoomTranscriptOnline",
			Handler:    _Room_RoomTranscriptOnline_Handler,
		},
		{
			MethodName: "RoomTranscript",
			Handler:    _Room_RoomTranscript_Handler,
		},
		{
			MethodName: "SetRoomVoice",
			Handler:    _Room_SetRoomVoice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "room/v1/room.proto",
}