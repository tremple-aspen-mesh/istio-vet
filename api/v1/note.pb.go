// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/v1/note.proto

/*
Package istio_vet_v1 is a generated protocol buffer package.

It is generated from these files:
	api/v1/note.proto

It has these top-level messages:
	Info
	Note
*/
package istio_vet_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// NoteLevel indicates the severity level of the note
type NoteLevel int32

const (
	NoteLevel_UNUSED  NoteLevel = 0
	NoteLevel_INFO    NoteLevel = 1
	NoteLevel_WARNING NoteLevel = 2
	NoteLevel_ERROR   NoteLevel = 3
)

var NoteLevel_name = map[int32]string{
	0: "UNUSED",
	1: "INFO",
	2: "WARNING",
	3: "ERROR",
}
var NoteLevel_value = map[string]int32{
	"UNUSED":  0,
	"INFO":    1,
	"WARNING": 2,
	"ERROR":   3,
}

func (x NoteLevel) String() string {
	return proto.EnumName(NoteLevel_name, int32(x))
}
func (NoteLevel) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Vetters use Info to provide their information
type Info struct {
	// Name of the vetter
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// Semver string of the vetter
	Version string `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
}

func (m *Info) Reset()                    { *m = Info{} }
func (m *Info) String() string            { return proto.CompactTextString(m) }
func (*Info) ProtoMessage()               {}
func (*Info) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Info) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Info) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

// Vetters generate Notes after inspecting the mesh configuration
type Note struct {
	// MD5 checksum of the generated note
	//
	// Used as UUID for notes
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// Type of the note
	//
	// Example "istio-component-mismatch", "missing-service-port-prefix"
	Type string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	// Short description of the note
	//
	// Summary may contain python template strings "${var}" which will be
	// substituted from values in Attr map described below. Summary
	// should only refer to template variables present in Attr map.
	Summary string `protobuf:"bytes,3,opt,name=summary" json:"summary,omitempty"`
	// Long description of the note
	//
	// Similar to Summary, Msg can contain python template strings "${var}" which
	// will be substituted from values in Attr map described below. Msg
	// should only refer to template variables present in Attr map.
	Msg string `protobuf:"bytes,4,opt,name=msg" json:"msg,omitempty"`
	// Severity of the note
	Level NoteLevel `protobuf:"varint,5,opt,name=level,enum=istio.vet.v1.NoteLevel" json:"level,omitempty"`
	// Map of template variables which can be used by Summary and Msg
	Attr map[string]string `protobuf:"bytes,6,rep,name=attr" json:"attr,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Note) Reset()                    { *m = Note{} }
func (m *Note) String() string            { return proto.CompactTextString(m) }
func (*Note) ProtoMessage()               {}
func (*Note) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Note) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Note) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Note) GetSummary() string {
	if m != nil {
		return m.Summary
	}
	return ""
}

func (m *Note) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *Note) GetLevel() NoteLevel {
	if m != nil {
		return m.Level
	}
	return NoteLevel_UNUSED
}

func (m *Note) GetAttr() map[string]string {
	if m != nil {
		return m.Attr
	}
	return nil
}

func init() {
	proto.RegisterType((*Info)(nil), "istio.vet.v1.Info")
	proto.RegisterType((*Note)(nil), "istio.vet.v1.Note")
	proto.RegisterEnum("istio.vet.v1.NoteLevel", NoteLevel_name, NoteLevel_value)
}

func init() { proto.RegisterFile("api/v1/note.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x5f, 0x4b, 0xc3, 0x30,
	0x14, 0xc5, 0x4d, 0x9b, 0x6e, 0xf6, 0x4e, 0x46, 0xbd, 0x08, 0x06, 0xf1, 0x61, 0xec, 0x69, 0x08,
	0x66, 0x7f, 0x7c, 0xf0, 0xcf, 0xdb, 0xc0, 0x29, 0x03, 0xe9, 0x20, 0x32, 0x7c, 0xae, 0x2c, 0x4a,
	0x70, 0x6b, 0x46, 0x9a, 0x05, 0xfa, 0xad, 0xfd, 0x08, 0x92, 0xb6, 0x1b, 0xc2, 0xde, 0xce, 0x39,
	0xf9, 0xe5, 0x9e, 0x9b, 0xc0, 0x79, 0xb6, 0x55, 0x43, 0x37, 0x1e, 0xe6, 0xda, 0x4a, 0xbe, 0x35,
	0xda, 0x6a, 0x3c, 0x53, 0x85, 0x55, 0x9a, 0x3b, 0x69, 0xb9, 0x1b, 0xf7, 0x47, 0x40, 0xe7, 0xf9,
	0x97, 0xc6, 0x2e, 0x04, 0x6a, 0xc5, 0x48, 0x8f, 0x0c, 0x62, 0x11, 0xa8, 0x15, 0x32, 0x68, 0x3b,
	0x69, 0x0a, 0xa5, 0x73, 0x16, 0x54, 0xe1, 0xde, 0xf6, 0x7f, 0x09, 0xd0, 0x54, 0x5b, 0x79, 0x74,
	0x05, 0x81, 0xda, 0x72, 0x2b, 0x1b, 0xbe, 0xd2, 0x7e, 0x4c, 0xb1, 0xdb, 0x6c, 0x32, 0x53, 0xb2,
	0xb0, 0x1e, 0xd3, 0x58, 0x4c, 0x20, 0xdc, 0x14, 0xdf, 0x8c, 0x56, 0xa9, 0x97, 0x78, 0x0b, 0xd1,
	0x5a, 0x3a, 0xb9, 0x66, 0x51, 0x8f, 0x0c, 0xba, 0x93, 0x4b, 0xfe, 0x7f, 0x51, 0xee, 0x2b, 0xdf,
	0xfc, 0xb1, 0xa8, 0x29, 0x1c, 0x01, 0xcd, 0xac, 0x35, 0xac, 0xd5, 0x0b, 0x07, 0x9d, 0xc9, 0xf5,
	0x31, 0xcd, 0xa7, 0xd6, 0x9a, 0x59, 0x6e, 0x4d, 0x29, 0x2a, 0xf2, 0xea, 0x1e, 0xe2, 0x43, 0xe4,
	0xfb, 0x7f, 0x64, 0xd9, 0xac, 0xef, 0x25, 0x5e, 0x40, 0xe4, 0xb2, 0xf5, 0x6e, 0xff, 0x80, 0xda,
	0x3c, 0x05, 0x0f, 0xe4, 0xe6, 0x11, 0xe2, 0x43, 0x3d, 0x02, 0xb4, 0x96, 0xe9, 0xf2, 0x7d, 0xf6,
	0x9c, 0x9c, 0xe0, 0x29, 0xd0, 0x79, 0xfa, 0xb2, 0x48, 0x08, 0x76, 0xa0, 0xfd, 0x31, 0x15, 0xe9,
	0x3c, 0x7d, 0x4d, 0x02, 0x8c, 0x21, 0x9a, 0x09, 0xb1, 0x10, 0x49, 0xf8, 0xd9, 0xaa, 0x3e, 0xfd,
	0xee, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xa6, 0x4c, 0x42, 0x45, 0x89, 0x01, 0x00, 0x00,
}
