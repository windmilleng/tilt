// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/proto/daemon.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	internal/proto/daemon.proto

It has these top-level messages:
	Service
	Cmd
	Mount
	Repo
	GitRepo
	CreateServiceReply
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type Service struct {
	K8SYaml        string   `protobuf:"bytes,1,opt,name=k8s_yaml,json=k8sYaml" json:"k8s_yaml,omitempty"`
	DockerfileText string   `protobuf:"bytes,2,opt,name=dockerfile_text,json=dockerfileText" json:"dockerfile_text,omitempty"`
	Mounts         []*Mount `protobuf:"bytes,3,rep,name=mounts" json:"mounts,omitempty"`
	Steps          []*Cmd   `protobuf:"bytes,4,rep,name=steps" json:"steps,omitempty"`
	Entrypoint     *Cmd     `protobuf:"bytes,5,opt,name=entrypoint" json:"entrypoint,omitempty"`
	DockerfileTag  string   `protobuf:"bytes,6,opt,name=dockerfile_tag,json=dockerfileTag" json:"dockerfile_tag,omitempty"`
}

func (m *Service) Reset()                    { *m = Service{} }
func (m *Service) String() string            { return proto1.CompactTextString(m) }
func (*Service) ProtoMessage()               {}
func (*Service) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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
	Argv []string `protobuf:"bytes,1,rep,name=argv" json:"argv,omitempty"`
}

func (m *Cmd) Reset()                    { *m = Cmd{} }
func (m *Cmd) String() string            { return proto1.CompactTextString(m) }
func (*Cmd) ProtoMessage()               {}
func (*Cmd) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Cmd) GetArgv() []string {
	if m != nil {
		return m.Argv
	}
	return nil
}

type Mount struct {
	Repo          *Repo  `protobuf:"bytes,1,opt,name=repo" json:"repo,omitempty"`
	ContainerPath string `protobuf:"bytes,2,opt,name=container_path,json=containerPath" json:"container_path,omitempty"`
}

func (m *Mount) Reset()                    { *m = Mount{} }
func (m *Mount) String() string            { return proto1.CompactTextString(m) }
func (*Mount) ProtoMessage()               {}
func (*Mount) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

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
	RepoType isRepo_RepoType `protobuf_oneof:"repo_type"`
}

func (m *Repo) Reset()                    { *m = Repo{} }
func (m *Repo) String() string            { return proto1.CompactTextString(m) }
func (*Repo) ProtoMessage()               {}
func (*Repo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type isRepo_RepoType interface{ isRepo_RepoType() }

type Repo_GitRepo struct {
	GitRepo *GitRepo `protobuf:"bytes,1,opt,name=git_repo,json=gitRepo,oneof"`
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
func (*Repo) XXX_OneofFuncs() (func(msg proto1.Message, b *proto1.Buffer) error, func(msg proto1.Message, tag, wire int, b *proto1.Buffer) (bool, error), func(msg proto1.Message) (n int), []interface{}) {
	return _Repo_OneofMarshaler, _Repo_OneofUnmarshaler, _Repo_OneofSizer, []interface{}{
		(*Repo_GitRepo)(nil),
	}
}

func _Repo_OneofMarshaler(msg proto1.Message, b *proto1.Buffer) error {
	m := msg.(*Repo)
	// repo_type
	switch x := m.RepoType.(type) {
	case *Repo_GitRepo:
		b.EncodeVarint(1<<3 | proto1.WireBytes)
		if err := b.EncodeMessage(x.GitRepo); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Repo.RepoType has unexpected type %T", x)
	}
	return nil
}

func _Repo_OneofUnmarshaler(msg proto1.Message, tag, wire int, b *proto1.Buffer) (bool, error) {
	m := msg.(*Repo)
	switch tag {
	case 1: // repo_type.git_repo
		if wire != proto1.WireBytes {
			return true, proto1.ErrInternalBadWireType
		}
		msg := new(GitRepo)
		err := b.DecodeMessage(msg)
		m.RepoType = &Repo_GitRepo{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Repo_OneofSizer(msg proto1.Message) (n int) {
	m := msg.(*Repo)
	// repo_type
	switch x := m.RepoType.(type) {
	case *Repo_GitRepo:
		s := proto1.Size(x.GitRepo)
		n += proto1.SizeVarint(1<<3 | proto1.WireBytes)
		n += proto1.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type GitRepo struct {
	LocalPath string `protobuf:"bytes,1,opt,name=local_path,json=localPath" json:"local_path,omitempty"`
}

func (m *GitRepo) Reset()                    { *m = GitRepo{} }
func (m *GitRepo) String() string            { return proto1.CompactTextString(m) }
func (*GitRepo) ProtoMessage()               {}
func (*GitRepo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GitRepo) GetLocalPath() string {
	if m != nil {
		return m.LocalPath
	}
	return ""
}

type CreateServiceReply struct {
}

func (m *CreateServiceReply) Reset()                    { *m = CreateServiceReply{} }
func (m *CreateServiceReply) String() string            { return proto1.CompactTextString(m) }
func (*CreateServiceReply) ProtoMessage()               {}
func (*CreateServiceReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func init() {
	proto1.RegisterType((*Service)(nil), "daemon.Service")
	proto1.RegisterType((*Cmd)(nil), "daemon.Cmd")
	proto1.RegisterType((*Mount)(nil), "daemon.Mount")
	proto1.RegisterType((*Repo)(nil), "daemon.Repo")
	proto1.RegisterType((*GitRepo)(nil), "daemon.GitRepo")
	proto1.RegisterType((*CreateServiceReply)(nil), "daemon.CreateServiceReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Daemon service

type DaemonClient interface {
	CreateService(ctx context.Context, in *Service, opts ...grpc.CallOption) (*CreateServiceReply, error)
}

type daemonClient struct {
	cc *grpc.ClientConn
}

func NewDaemonClient(cc *grpc.ClientConn) DaemonClient {
	return &daemonClient{cc}
}

func (c *daemonClient) CreateService(ctx context.Context, in *Service, opts ...grpc.CallOption) (*CreateServiceReply, error) {
	out := new(CreateServiceReply)
	err := grpc.Invoke(ctx, "/daemon.Daemon/CreateService", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Daemon service

type DaemonServer interface {
	CreateService(context.Context, *Service) (*CreateServiceReply, error)
}

func RegisterDaemonServer(s *grpc.Server, srv DaemonServer) {
	s.RegisterService(&_Daemon_serviceDesc, srv)
}

func _Daemon_CreateService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Service)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).CreateService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daemon.Daemon/CreateService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).CreateService(ctx, req.(*Service))
	}
	return interceptor(ctx, in, info, handler)
}

var _Daemon_serviceDesc = grpc.ServiceDesc{
	ServiceName: "daemon.Daemon",
	HandlerType: (*DaemonServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateService",
			Handler:    _Daemon_CreateService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/proto/daemon.proto",
}

func init() { proto1.RegisterFile("internal/proto/daemon.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 406 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x52, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0x8d, 0x6a, 0x5b, 0x8e, 0xc7, 0x75, 0x02, 0x4b, 0x0f, 0x4a, 0x4a, 0x41, 0x15, 0x84, 0x1a,
	0x5a, 0x2c, 0xea, 0x5c, 0x02, 0x85, 0x42, 0xe3, 0x42, 0x73, 0x29, 0x04, 0xb5, 0x97, 0xf6, 0x22,
	0x36, 0xd2, 0x54, 0x5e, 0xbc, 0x5f, 0xac, 0x26, 0x69, 0xfc, 0x9b, 0xfb, 0x27, 0x8a, 0x56, 0xb2,
	0xab, 0x3a, 0xa7, 0x9d, 0x79, 0xef, 0x69, 0xde, 0x63, 0x46, 0xf0, 0x52, 0x68, 0x42, 0xa7, 0xb9,
	0x4c, 0xad, 0x33, 0x64, 0xd2, 0x92, 0xa3, 0x32, 0x7a, 0xe1, 0x1b, 0x16, 0xb6, 0x5d, 0xf2, 0x27,
	0x80, 0xf1, 0x37, 0x74, 0x0f, 0xa2, 0x40, 0x76, 0x06, 0xc7, 0x9b, 0xab, 0x3a, 0xdf, 0x72, 0x25,
	0xa3, 0x20, 0x0e, 0xe6, 0x93, 0x6c, 0xbc, 0xb9, 0xaa, 0x7f, 0x70, 0x25, 0xd9, 0x1b, 0x38, 0x2d,
	0x4d, 0xb1, 0x41, 0xf7, 0x4b, 0x48, 0xcc, 0x09, 0x1f, 0x29, 0x7a, 0xe6, 0x15, 0x27, 0xff, 0xe0,
	0xef, 0xf8, 0x48, 0xec, 0x02, 0x42, 0x65, 0xee, 0x35, 0xd5, 0xd1, 0x20, 0x1e, 0xcc, 0xa7, 0xcb,
	0xd9, 0xa2, 0xb3, 0xfd, 0xda, 0xa0, 0x59, 0x47, 0xb2, 0xd7, 0x30, 0xaa, 0x09, 0x6d, 0x1d, 0x0d,
	0xbd, 0x6a, 0xba, 0x53, 0xad, 0x54, 0x99, 0xb5, 0x0c, 0x7b, 0x0b, 0x80, 0x9a, 0xdc, 0xd6, 0x1a,
	0xa1, 0x29, 0x1a, 0xc5, 0xc1, 0xa1, 0xae, 0x47, 0xb3, 0x0b, 0x38, 0xe9, 0xe7, 0xe3, 0x55, 0x14,
	0xfa, 0x78, 0xb3, 0x5e, 0x3c, 0x5e, 0x25, 0x67, 0x30, 0x58, 0xa9, 0x92, 0x31, 0x18, 0x72, 0x57,
	0x3d, 0x44, 0x41, 0x3c, 0x98, 0x4f, 0x32, 0x5f, 0x27, 0xb7, 0x30, 0xf2, 0x11, 0x59, 0x0c, 0x43,
	0x87, 0xd6, 0xf8, 0x0d, 0x4c, 0x97, 0xcf, 0x77, 0x8e, 0x19, 0x5a, 0x93, 0x79, 0xa6, 0x31, 0x2b,
	0x8c, 0x26, 0x2e, 0x34, 0xba, 0xdc, 0x72, 0x5a, 0x77, 0xbb, 0x98, 0xed, 0xd1, 0x5b, 0x4e, 0xeb,
	0xe4, 0x13, 0x0c, 0x9b, 0x8f, 0xd8, 0x3b, 0x38, 0xae, 0x04, 0xe5, 0xbd, 0xa1, 0xa7, 0xbb, 0xa1,
	0x5f, 0x04, 0x35, 0x92, 0x9b, 0xa3, 0x6c, 0x5c, 0xb5, 0xe5, 0xf5, 0x14, 0x26, 0x8d, 0x32, 0xa7,
	0xad, 0xc5, 0x64, 0x0e, 0xe3, 0x4e, 0xc2, 0x5e, 0x01, 0x48, 0x53, 0x70, 0xd9, 0x1a, 0xb6, 0xe7,
	0x99, 0x78, 0xc4, 0x9b, 0xbd, 0x00, 0xb6, 0x72, 0xc8, 0x09, 0xbb, 0x63, 0x66, 0x68, 0xe5, 0x76,
	0x79, 0x03, 0xe1, 0x67, 0xef, 0xc4, 0x3e, 0xc2, 0xec, 0x3f, 0x9e, 0xed, 0x33, 0x74, 0xc0, 0xf9,
	0xf9, 0x7e, 0xb7, 0x4f, 0xe6, 0x24, 0x47, 0xd7, 0x97, 0x3f, 0xdf, 0x57, 0x82, 0xd6, 0xf7, 0x77,
	0x8b, 0xc2, 0xa8, 0xf4, 0xb7, 0xd0, 0xa5, 0x12, 0x52, 0xa2, 0xae, 0x52, 0x12, 0x92, 0xd2, 0x83,
	0x5f, 0xed, 0x83, 0x7f, 0xee, 0x42, 0xff, 0x5c, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x07, 0x19,
	0x19, 0x30, 0x8a, 0x02, 0x00, 0x00,
}
