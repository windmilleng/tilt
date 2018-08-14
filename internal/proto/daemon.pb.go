// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/proto/daemon.proto

package proto // import "github.com/windmilleng/tilt/internal/proto"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Service struct {
	K8SYaml              string   `protobuf:"bytes,1,opt,name=k8s_yaml,json=k8sYaml,proto3" json:"k8s_yaml,omitempty"`
	DockerfileText       string   `protobuf:"bytes,2,opt,name=dockerfile_text,json=dockerfileText,proto3" json:"dockerfile_text,omitempty"`
	Mounts               []*Mount `protobuf:"bytes,3,rep,name=mounts,proto3" json:"mounts,omitempty"`
	Steps                []*Cmd   `protobuf:"bytes,4,rep,name=steps,proto3" json:"steps,omitempty"`
	Entrypoint           *Cmd     `protobuf:"bytes,5,opt,name=entrypoint,proto3" json:"entrypoint,omitempty"`
	DockerfileTag        string   `protobuf:"bytes,6,opt,name=dockerfile_tag,json=dockerfileTag,proto3" json:"dockerfile_tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Service) Reset()         { *m = Service{} }
func (m *Service) String() string { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()    {}
func (*Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_daemon_6382cf112eee2b2c, []int{0}
}
func (m *Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service.Unmarshal(m, b)
}
func (m *Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service.Marshal(b, m, deterministic)
}
func (dst *Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service.Merge(dst, src)
}
func (m *Service) XXX_Size() int {
	return xxx_messageInfo_Service.Size(m)
}
func (m *Service) XXX_DiscardUnknown() {
	xxx_messageInfo_Service.DiscardUnknown(m)
}

var xxx_messageInfo_Service proto.InternalMessageInfo

func (m *Service) GetK8SYaml() string {
	if m != nil {
		return m.K8SYaml
	}
	return ""
}

func (m *Service) GetDockerfileText() string {
	if m != nil {
		return m.DockerfileText
	}
	return ""
}

func (m *Service) GetMounts() []*Mount {
	if m != nil {
		return m.Mounts
	}
	return nil
}

func (m *Service) GetSteps() []*Cmd {
	if m != nil {
		return m.Steps
	}
	return nil
}

func (m *Service) GetEntrypoint() *Cmd {
	if m != nil {
		return m.Entrypoint
	}
	return nil
}

func (m *Service) GetDockerfileTag() string {
	if m != nil {
		return m.DockerfileTag
	}
	return ""
}

type Cmd struct {
	Argv                 []string `protobuf:"bytes,1,rep,name=argv,proto3" json:"argv,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cmd) Reset()         { *m = Cmd{} }
func (m *Cmd) String() string { return proto.CompactTextString(m) }
func (*Cmd) ProtoMessage()    {}
func (*Cmd) Descriptor() ([]byte, []int) {
	return fileDescriptor_daemon_6382cf112eee2b2c, []int{1}
}
func (m *Cmd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cmd.Unmarshal(m, b)
}
func (m *Cmd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cmd.Marshal(b, m, deterministic)
}
func (dst *Cmd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cmd.Merge(dst, src)
}
func (m *Cmd) XXX_Size() int {
	return xxx_messageInfo_Cmd.Size(m)
}
func (m *Cmd) XXX_DiscardUnknown() {
	xxx_messageInfo_Cmd.DiscardUnknown(m)
}

var xxx_messageInfo_Cmd proto.InternalMessageInfo

func (m *Cmd) GetArgv() []string {
	if m != nil {
		return m.Argv
	}
	return nil
}

type Mount struct {
	Repo                 *Repo    `protobuf:"bytes,1,opt,name=repo,proto3" json:"repo,omitempty"`
	ContainerPath        string   `protobuf:"bytes,2,opt,name=container_path,json=containerPath,proto3" json:"container_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Mount) Reset()         { *m = Mount{} }
func (m *Mount) String() string { return proto.CompactTextString(m) }
func (*Mount) ProtoMessage()    {}
func (*Mount) Descriptor() ([]byte, []int) {
	return fileDescriptor_daemon_6382cf112eee2b2c, []int{2}
}
func (m *Mount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Mount.Unmarshal(m, b)
}
func (m *Mount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Mount.Marshal(b, m, deterministic)
}
func (dst *Mount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mount.Merge(dst, src)
}
func (m *Mount) XXX_Size() int {
	return xxx_messageInfo_Mount.Size(m)
}
func (m *Mount) XXX_DiscardUnknown() {
	xxx_messageInfo_Mount.DiscardUnknown(m)
}

var xxx_messageInfo_Mount proto.InternalMessageInfo

func (m *Mount) GetRepo() *Repo {
	if m != nil {
		return m.Repo
	}
	return nil
}

func (m *Mount) GetContainerPath() string {
	if m != nil {
		return m.ContainerPath
	}
	return ""
}

type Repo struct {
	// Types that are valid to be assigned to RepoType:
	//	*Repo_GitRepo
	RepoType             isRepo_RepoType `protobuf_oneof:"repo_type"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Repo) Reset()         { *m = Repo{} }
func (m *Repo) String() string { return proto.CompactTextString(m) }
func (*Repo) ProtoMessage()    {}
func (*Repo) Descriptor() ([]byte, []int) {
	return fileDescriptor_daemon_6382cf112eee2b2c, []int{3}
}
func (m *Repo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Repo.Unmarshal(m, b)
}
func (m *Repo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Repo.Marshal(b, m, deterministic)
}
func (dst *Repo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Repo.Merge(dst, src)
}
func (m *Repo) XXX_Size() int {
	return xxx_messageInfo_Repo.Size(m)
}
func (m *Repo) XXX_DiscardUnknown() {
	xxx_messageInfo_Repo.DiscardUnknown(m)
}

var xxx_messageInfo_Repo proto.InternalMessageInfo

type isRepo_RepoType interface {
	isRepo_RepoType()
}

type Repo_GitRepo struct {
	GitRepo *GitRepo `protobuf:"bytes,1,opt,name=git_repo,json=gitRepo,proto3,oneof"`
}

func (*Repo_GitRepo) isRepo_RepoType() {}

func (m *Repo) GetRepoType() isRepo_RepoType {
	if m != nil {
		return m.RepoType
	}
	return nil
}

func (m *Repo) GetGitRepo() *GitRepo {
	if x, ok := m.GetRepoType().(*Repo_GitRepo); ok {
		return x.GitRepo
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Repo) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Repo_OneofMarshaler, _Repo_OneofUnmarshaler, _Repo_OneofSizer, []interface{}{
		(*Repo_GitRepo)(nil),
	}
}

func _Repo_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Repo)
	// repo_type
	switch x := m.RepoType.(type) {
	case *Repo_GitRepo:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.GitRepo); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Repo.RepoType has unexpected type %T", x)
	}
	return nil
}

func _Repo_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Repo)
	switch tag {
	case 1: // repo_type.git_repo
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GitRepo)
		err := b.DecodeMessage(msg)
		m.RepoType = &Repo_GitRepo{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Repo_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Repo)
	// repo_type
	switch x := m.RepoType.(type) {
	case *Repo_GitRepo:
		s := proto.Size(x.GitRepo)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type GitRepo struct {
	LocalPath            string   `protobuf:"bytes,1,opt,name=local_path,json=localPath,proto3" json:"local_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GitRepo) Reset()         { *m = GitRepo{} }
func (m *GitRepo) String() string { return proto.CompactTextString(m) }
func (*GitRepo) ProtoMessage()    {}
func (*GitRepo) Descriptor() ([]byte, []int) {
	return fileDescriptor_daemon_6382cf112eee2b2c, []int{4}
}
func (m *GitRepo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GitRepo.Unmarshal(m, b)
}
func (m *GitRepo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GitRepo.Marshal(b, m, deterministic)
}
func (dst *GitRepo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GitRepo.Merge(dst, src)
}
func (m *GitRepo) XXX_Size() int {
	return xxx_messageInfo_GitRepo.Size(m)
}
func (m *GitRepo) XXX_DiscardUnknown() {
	xxx_messageInfo_GitRepo.DiscardUnknown(m)
}

var xxx_messageInfo_GitRepo proto.InternalMessageInfo

func (m *GitRepo) GetLocalPath() string {
	if m != nil {
		return m.LocalPath
	}
	return ""
}

type Output struct {
	Stdout               []byte   `protobuf:"bytes,1,opt,name=stdout,proto3" json:"stdout,omitempty"`
	Stderr               []byte   `protobuf:"bytes,2,opt,name=stderr,proto3" json:"stderr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Output) Reset()         { *m = Output{} }
func (m *Output) String() string { return proto.CompactTextString(m) }
func (*Output) ProtoMessage()    {}
func (*Output) Descriptor() ([]byte, []int) {
	return fileDescriptor_daemon_6382cf112eee2b2c, []int{5}
}
func (m *Output) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Output.Unmarshal(m, b)
}
func (m *Output) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Output.Marshal(b, m, deterministic)
}
func (dst *Output) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Output.Merge(dst, src)
}
func (m *Output) XXX_Size() int {
	return xxx_messageInfo_Output.Size(m)
}
func (m *Output) XXX_DiscardUnknown() {
	xxx_messageInfo_Output.DiscardUnknown(m)
}

var xxx_messageInfo_Output proto.InternalMessageInfo

func (m *Output) GetStdout() []byte {
	if m != nil {
		return m.Stdout
	}
	return nil
}

func (m *Output) GetStderr() []byte {
	if m != nil {
		return m.Stderr
	}
	return nil
}

type CreateServiceReply struct {
	Output               *Output  `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateServiceReply) Reset()         { *m = CreateServiceReply{} }
func (m *CreateServiceReply) String() string { return proto.CompactTextString(m) }
func (*CreateServiceReply) ProtoMessage()    {}
func (*CreateServiceReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_daemon_6382cf112eee2b2c, []int{6}
}
func (m *CreateServiceReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateServiceReply.Unmarshal(m, b)
}
func (m *CreateServiceReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateServiceReply.Marshal(b, m, deterministic)
}
func (dst *CreateServiceReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateServiceReply.Merge(dst, src)
}
func (m *CreateServiceReply) XXX_Size() int {
	return xxx_messageInfo_CreateServiceReply.Size(m)
}
func (m *CreateServiceReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateServiceReply.DiscardUnknown(m)
}

var xxx_messageInfo_CreateServiceReply proto.InternalMessageInfo

func (m *CreateServiceReply) GetOutput() *Output {
	if m != nil {
		return m.Output
	}
	return nil
}

func init() {
	proto.RegisterType((*Service)(nil), "daemon.Service")
	proto.RegisterType((*Cmd)(nil), "daemon.Cmd")
	proto.RegisterType((*Mount)(nil), "daemon.Mount")
	proto.RegisterType((*Repo)(nil), "daemon.Repo")
	proto.RegisterType((*GitRepo)(nil), "daemon.GitRepo")
	proto.RegisterType((*Output)(nil), "daemon.Output")
	proto.RegisterType((*CreateServiceReply)(nil), "daemon.CreateServiceReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DaemonClient is the client API for Daemon service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DaemonClient interface {
	CreateService(ctx context.Context, in *Service, opts ...grpc.CallOption) (Daemon_CreateServiceClient, error)
}

type daemonClient struct {
	cc *grpc.ClientConn
}

func NewDaemonClient(cc *grpc.ClientConn) DaemonClient {
	return &daemonClient{cc}
}

func (c *daemonClient) CreateService(ctx context.Context, in *Service, opts ...grpc.CallOption) (Daemon_CreateServiceClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Daemon_serviceDesc.Streams[0], "/daemon.Daemon/CreateService", opts...)
	if err != nil {
		return nil, err
	}
	x := &daemonCreateServiceClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Daemon_CreateServiceClient interface {
	Recv() (*CreateServiceReply, error)
	grpc.ClientStream
}

type daemonCreateServiceClient struct {
	grpc.ClientStream
}

func (x *daemonCreateServiceClient) Recv() (*CreateServiceReply, error) {
	m := new(CreateServiceReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DaemonServer is the server API for Daemon service.
type DaemonServer interface {
	CreateService(*Service, Daemon_CreateServiceServer) error
}

func RegisterDaemonServer(s *grpc.Server, srv DaemonServer) {
	s.RegisterService(&_Daemon_serviceDesc, srv)
}

func _Daemon_CreateService_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Service)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DaemonServer).CreateService(m, &daemonCreateServiceServer{stream})
}

type Daemon_CreateServiceServer interface {
	Send(*CreateServiceReply) error
	grpc.ServerStream
}

type daemonCreateServiceServer struct {
	grpc.ServerStream
}

func (x *daemonCreateServiceServer) Send(m *CreateServiceReply) error {
	return x.ServerStream.SendMsg(m)
}

var _Daemon_serviceDesc = grpc.ServiceDesc{
	ServiceName: "daemon.Daemon",
	HandlerType: (*DaemonServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateService",
			Handler:       _Daemon_CreateService_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "internal/proto/daemon.proto",
}

func init() { proto.RegisterFile("internal/proto/daemon.proto", fileDescriptor_daemon_6382cf112eee2b2c) }

var fileDescriptor_daemon_6382cf112eee2b2c = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x52, 0x41, 0x6f, 0xd3, 0x30,
	0x18, 0x5d, 0x68, 0x9b, 0xae, 0x5f, 0xd7, 0x4e, 0xf2, 0x01, 0x65, 0x43, 0x48, 0x25, 0xd2, 0xa0,
	0x12, 0xa8, 0x9d, 0xca, 0xa5, 0x12, 0x1c, 0x60, 0x45, 0x02, 0x21, 0x21, 0x26, 0xc3, 0x05, 0x2e,
	0x91, 0x97, 0x7c, 0xa4, 0x56, 0x1d, 0x3b, 0x72, 0xbe, 0x8c, 0xf5, 0x37, 0xf3, 0x27, 0x50, 0x1c,
	0xb7, 0x74, 0xec, 0x92, 0xf8, 0x7b, 0xef, 0xe5, 0xbd, 0x17, 0xdb, 0xf0, 0x44, 0x6a, 0x42, 0xab,
	0x85, 0x9a, 0x97, 0xd6, 0x90, 0x99, 0x67, 0x02, 0x0b, 0xa3, 0x67, 0x6e, 0x60, 0x61, 0x3b, 0xc5,
	0x7f, 0x02, 0xe8, 0x7f, 0x43, 0x7b, 0x2b, 0x53, 0x64, 0x67, 0x70, 0xbc, 0x59, 0x56, 0xc9, 0x56,
	0x14, 0x2a, 0x0a, 0x26, 0xc1, 0x74, 0xc0, 0xfb, 0x9b, 0x65, 0xf5, 0x43, 0x14, 0x8a, 0xbd, 0x80,
	0xd3, 0xcc, 0xa4, 0x1b, 0xb4, 0xbf, 0xa4, 0xc2, 0x84, 0xf0, 0x8e, 0xa2, 0x47, 0x4e, 0x31, 0xfe,
	0x07, 0x7f, 0xc7, 0x3b, 0x62, 0x17, 0x10, 0x16, 0xa6, 0xd6, 0x54, 0x45, 0x9d, 0x49, 0x67, 0x3a,
	0x5c, 0x8c, 0x66, 0x3e, 0xf6, 0x4b, 0x83, 0x72, 0x4f, 0xb2, 0x67, 0xd0, 0xab, 0x08, 0xcb, 0x2a,
	0xea, 0x3a, 0xd5, 0x70, 0xa7, 0x5a, 0x15, 0x19, 0x6f, 0x19, 0xf6, 0x12, 0x00, 0x35, 0xd9, 0x6d,
	0x69, 0xa4, 0xa6, 0xa8, 0x37, 0x09, 0xfe, 0xd7, 0x1d, 0xd0, 0xec, 0x02, 0xc6, 0x87, 0xfd, 0x44,
	0x1e, 0x85, 0xae, 0xde, 0xe8, 0xa0, 0x9e, 0xc8, 0xe3, 0x33, 0xe8, 0xac, 0x8a, 0x8c, 0x31, 0xe8,
	0x0a, 0x9b, 0xdf, 0x46, 0xc1, 0xa4, 0x33, 0x1d, 0x70, 0xb7, 0x8e, 0xaf, 0xa1, 0xe7, 0x2a, 0xb2,
	0x09, 0x74, 0x2d, 0x96, 0xc6, 0xed, 0xc0, 0x70, 0x71, 0xb2, 0x4b, 0xe4, 0x58, 0x1a, 0xee, 0x98,
	0x26, 0x2c, 0x35, 0x9a, 0x84, 0xd4, 0x68, 0x93, 0x52, 0xd0, 0xda, 0xef, 0xc5, 0x68, 0x8f, 0x5e,
	0x0b, 0x5a, 0xc7, 0xef, 0xa1, 0xdb, 0x7c, 0xc4, 0x5e, 0xc1, 0x71, 0x2e, 0x29, 0x39, 0x30, 0x3d,
	0xdd, 0x99, 0x7e, 0x94, 0xd4, 0x48, 0x3e, 0x1d, 0xf1, 0x7e, 0xde, 0x2e, 0xaf, 0x86, 0x30, 0x68,
	0x94, 0x09, 0x6d, 0x4b, 0x8c, 0xa7, 0xd0, 0xf7, 0x12, 0xf6, 0x14, 0x40, 0x99, 0x54, 0xa8, 0x36,
	0xb0, 0x3d, 0x9e, 0x81, 0x43, 0x5c, 0xd8, 0x12, 0xc2, 0xaf, 0x35, 0x95, 0x35, 0xb1, 0xc7, 0x10,
	0x56, 0x94, 0x99, 0x9a, 0x9c, 0xe8, 0x84, 0xfb, 0xc9, 0xe3, 0x68, 0xad, 0x6b, 0xdb, 0xe2, 0x68,
	0x6d, 0xfc, 0x16, 0xd8, 0xca, 0xa2, 0x20, 0xf4, 0xd7, 0x80, 0x63, 0xa9, 0xb6, 0xec, 0x39, 0x84,
	0xc6, 0xf9, 0xf9, 0xca, 0xe3, 0x5d, 0xe5, 0x36, 0x85, 0x7b, 0x76, 0xf1, 0x19, 0xc2, 0x0f, 0x8e,
	0x60, 0xef, 0x60, 0x74, 0xcf, 0x87, 0xed, 0xff, 0xd2, 0x03, 0xe7, 0xe7, 0xfb, 0xd3, 0x7b, 0x90,
	0x17, 0x1f, 0x5d, 0x06, 0x57, 0x8b, 0x9f, 0x97, 0xb9, 0xa4, 0x75, 0x7d, 0x33, 0x4b, 0x4d, 0x31,
	0xff, 0x2d, 0x75, 0x56, 0x48, 0xa5, 0x50, 0xe7, 0x73, 0x92, 0x8a, 0xe6, 0xf7, 0xaf, 0xf3, 0x1b,
	0xf7, 0xbc, 0x09, 0xdd, 0xeb, 0xf5, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa1, 0x7b, 0x9c, 0x05,
	0xed, 0x02, 0x00, 0x00,
}
