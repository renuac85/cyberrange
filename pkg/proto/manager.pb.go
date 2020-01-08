// Code generated by protoc-gen-go. DO NOT EDIT.
// source: manager.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Machine struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Ip                   string   `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	Status               string   `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Machine) Reset()         { *m = Machine{} }
func (m *Machine) String() string { return proto.CompactTextString(m) }
func (*Machine) ProtoMessage()    {}
func (*Machine) Descriptor() ([]byte, []int) {
	return fileDescriptor_cde9ec64f0d2c859, []int{0}
}

func (m *Machine) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Machine.Unmarshal(m, b)
}
func (m *Machine) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Machine.Marshal(b, m, deterministic)
}
func (m *Machine) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Machine.Merge(m, src)
}
func (m *Machine) XXX_Size() int {
	return xxx_messageInfo_Machine.Size(m)
}
func (m *Machine) XXX_DiscardUnknown() {
	xxx_messageInfo_Machine.DiscardUnknown(m)
}

var xxx_messageInfo_Machine proto.InternalMessageInfo

func (m *Machine) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Machine) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *Machine) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type MachineList struct {
	Machines             []*Machine `protobuf:"bytes,1,rep,name=machines,proto3" json:"machines,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *MachineList) Reset()         { *m = MachineList{} }
func (m *MachineList) String() string { return proto.CompactTextString(m) }
func (*MachineList) ProtoMessage()    {}
func (*MachineList) Descriptor() ([]byte, []int) {
	return fileDescriptor_cde9ec64f0d2c859, []int{1}
}

func (m *MachineList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MachineList.Unmarshal(m, b)
}
func (m *MachineList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MachineList.Marshal(b, m, deterministic)
}
func (m *MachineList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MachineList.Merge(m, src)
}
func (m *MachineList) XXX_Size() int {
	return xxx_messageInfo_MachineList.Size(m)
}
func (m *MachineList) XXX_DiscardUnknown() {
	xxx_messageInfo_MachineList.DiscardUnknown(m)
}

var xxx_messageInfo_MachineList proto.InternalMessageInfo

func (m *MachineList) GetMachines() []*Machine {
	if m != nil {
		return m.Machines
	}
	return nil
}

type Response struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_cde9ec64f0d2c859, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_cde9ec64f0d2c859, []int{3}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Machine)(nil), "proto.Machine")
	proto.RegisterType((*MachineList)(nil), "proto.MachineList")
	proto.RegisterType((*Response)(nil), "proto.Response")
	proto.RegisterType((*Empty)(nil), "proto.Empty")
}

func init() { proto.RegisterFile("manager.proto", fileDescriptor_cde9ec64f0d2c859) }

var fileDescriptor_cde9ec64f0d2c859 = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0x95, 0x34, 0xff, 0xb8, 0x40, 0x91, 0x3c, 0xa0, 0xa8, 0x53, 0x95, 0x85, 0x50, 0xa4,
	0x0c, 0x65, 0x62, 0xae, 0xba, 0xc1, 0xe2, 0x3c, 0x81, 0xa9, 0x4e, 0x24, 0x82, 0xd8, 0x96, 0x7d,
	0x45, 0xe2, 0x81, 0x79, 0x0f, 0x94, 0xab, 0x55, 0x09, 0x96, 0x66, 0x4a, 0xbe, 0xef, 0x3e, 0xff,
	0xee, 0xce, 0x86, 0x9b, 0x51, 0x69, 0xf5, 0x8e, 0xae, 0xb5, 0xce, 0x90, 0x11, 0x29, 0x7f, 0xea,
	0x3d, 0xe4, 0xaf, 0xea, 0xd0, 0x0f, 0x1a, 0x85, 0x80, 0x44, 0xab, 0x11, 0xab, 0x68, 0x1d, 0x35,
	0x57, 0x92, 0xff, 0xc5, 0x12, 0xe2, 0xc1, 0x56, 0x31, 0x3b, 0xf1, 0x60, 0xc5, 0x1d, 0x64, 0x9e,
	0x14, 0x1d, 0x7d, 0xb5, 0x60, 0x2f, 0xa8, 0xfa, 0x19, 0xca, 0x80, 0x79, 0x19, 0x3c, 0x89, 0x0d,
	0x14, 0xe3, 0x49, 0xfa, 0x2a, 0x5a, 0x2f, 0x9a, 0x72, 0xbb, 0x3c, 0xb5, 0x6d, 0x43, 0x4a, 0x9e,
	0xeb, 0x75, 0x0d, 0x85, 0x44, 0x6f, 0x8d, 0xf6, 0x38, 0xe1, 0x1d, 0xfa, 0xe3, 0x27, 0xf1, 0x10,
	0x85, 0x0c, 0xaa, 0xce, 0x21, 0xdd, 0x8f, 0x96, 0xbe, 0xb7, 0x3f, 0xf1, 0x34, 0x2f, 0xef, 0x21,
	0x5a, 0x28, 0x77, 0x3d, 0x1e, 0x3e, 0x76, 0x0e, 0x15, 0xa1, 0xf8, 0xd7, 0x61, 0x75, 0x1b, 0xf4,
	0x19, 0xfe, 0x00, 0xd9, 0xdc, 0x68, 0x03, 0x69, 0x47, 0xca, 0xd1, 0xe5, 0xe4, 0x3d, 0x24, 0x1d,
	0x19, 0x7b, 0x39, 0xb8, 0x81, 0x5c, 0xa2, 0x9f, 0x07, 0x7d, 0x84, 0xa2, 0xd3, 0xca, 0xfa, 0xde,
	0xd0, 0xac, 0xb5, 0x24, 0x7e, 0xe1, 0x1c, 0x6e, 0x03, 0x09, 0x3f, 0xcf, 0x75, 0x28, 0xf0, 0x9d,
	0xae, 0xc4, 0xdf, 0x63, 0x53, 0xe2, 0x2d, 0x63, 0xeb, 0xe9, 0x37, 0x00, 0x00, 0xff, 0xff, 0xb6,
	0x9c, 0x0f, 0xef, 0x35, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ManagerClient is the client API for Manager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ManagerClient interface {
	CheckCreate(ctx context.Context, in *Machine, opts ...grpc.CallOption) (*Response, error)
	Create(ctx context.Context, in *Machine, opts ...grpc.CallOption) (*Response, error)
	Start(ctx context.Context, in *Machine, opts ...grpc.CallOption) (*Response, error)
	Stop(ctx context.Context, in *Machine, opts ...grpc.CallOption) (*Response, error)
	Restart(ctx context.Context, in *Machine, opts ...grpc.CallOption) (*Response, error)
	Snapshot(ctx context.Context, in *Machine, opts ...grpc.CallOption) (*Response, error)
	Revert(ctx context.Context, in *Machine, opts ...grpc.CallOption) (*Response, error)
	List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MachineList, error)
}

type managerClient struct {
	cc *grpc.ClientConn
}

func NewManagerClient(cc *grpc.ClientConn) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) CheckCreate(ctx context.Context, in *Machine, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.Manager/CheckCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) Create(ctx context.Context, in *Machine, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.Manager/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) Start(ctx context.Context, in *Machine, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.Manager/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) Stop(ctx context.Context, in *Machine, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.Manager/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) Restart(ctx context.Context, in *Machine, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.Manager/Restart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) Snapshot(ctx context.Context, in *Machine, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.Manager/Snapshot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) Revert(ctx context.Context, in *Machine, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.Manager/Revert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MachineList, error) {
	out := new(MachineList)
	err := c.cc.Invoke(ctx, "/proto.Manager/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagerServer is the server API for Manager service.
type ManagerServer interface {
	CheckCreate(context.Context, *Machine) (*Response, error)
	Create(context.Context, *Machine) (*Response, error)
	Start(context.Context, *Machine) (*Response, error)
	Stop(context.Context, *Machine) (*Response, error)
	Restart(context.Context, *Machine) (*Response, error)
	Snapshot(context.Context, *Machine) (*Response, error)
	Revert(context.Context, *Machine) (*Response, error)
	List(context.Context, *Empty) (*MachineList, error)
}

// UnimplementedManagerServer can be embedded to have forward compatible implementations.
type UnimplementedManagerServer struct {
}

func (*UnimplementedManagerServer) CheckCreate(ctx context.Context, req *Machine) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckCreate not implemented")
}
func (*UnimplementedManagerServer) Create(ctx context.Context, req *Machine) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedManagerServer) Start(ctx context.Context, req *Machine) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (*UnimplementedManagerServer) Stop(ctx context.Context, req *Machine) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (*UnimplementedManagerServer) Restart(ctx context.Context, req *Machine) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restart not implemented")
}
func (*UnimplementedManagerServer) Snapshot(ctx context.Context, req *Machine) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Snapshot not implemented")
}
func (*UnimplementedManagerServer) Revert(ctx context.Context, req *Machine) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Revert not implemented")
}
func (*UnimplementedManagerServer) List(ctx context.Context, req *Empty) (*MachineList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterManagerServer(s *grpc.Server, srv ManagerServer) {
	s.RegisterService(&_Manager_serviceDesc, srv)
}

func _Manager_CheckCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Machine)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CheckCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Manager/CheckCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CheckCreate(ctx, req.(*Machine))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Machine)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Manager/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).Create(ctx, req.(*Machine))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Machine)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Manager/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).Start(ctx, req.(*Machine))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Machine)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Manager/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).Stop(ctx, req.(*Machine))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_Restart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Machine)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).Restart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Manager/Restart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).Restart(ctx, req.(*Machine))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_Snapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Machine)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).Snapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Manager/Snapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).Snapshot(ctx, req.(*Machine))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_Revert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Machine)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).Revert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Manager/Revert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).Revert(ctx, req.(*Machine))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Manager/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).List(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Manager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Manager",
	HandlerType: (*ManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckCreate",
			Handler:    _Manager_CheckCreate_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Manager_Create_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _Manager_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Manager_Stop_Handler,
		},
		{
			MethodName: "Restart",
			Handler:    _Manager_Restart_Handler,
		},
		{
			MethodName: "Snapshot",
			Handler:    _Manager_Snapshot_Handler,
		},
		{
			MethodName: "Revert",
			Handler:    _Manager_Revert_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Manager_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "manager.proto",
}
