// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/endpoint/v3/endpoint.proto

package envoy_config_endpoint_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/type/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ClusterLoadAssignment struct {
	ClusterName          string                        `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	Endpoints            []*LocalityLbEndpoints        `protobuf:"bytes,2,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	NamedEndpoints       map[string]*Endpoint          `protobuf:"bytes,5,rep,name=named_endpoints,json=namedEndpoints,proto3" json:"named_endpoints,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Policy               *ClusterLoadAssignment_Policy `protobuf:"bytes,4,opt,name=policy,proto3" json:"policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *ClusterLoadAssignment) Reset()         { *m = ClusterLoadAssignment{} }
func (m *ClusterLoadAssignment) String() string { return proto.CompactTextString(m) }
func (*ClusterLoadAssignment) ProtoMessage()    {}
func (*ClusterLoadAssignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_3459cc50e28e3760, []int{0}
}

func (m *ClusterLoadAssignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterLoadAssignment.Unmarshal(m, b)
}
func (m *ClusterLoadAssignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterLoadAssignment.Marshal(b, m, deterministic)
}
func (m *ClusterLoadAssignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterLoadAssignment.Merge(m, src)
}
func (m *ClusterLoadAssignment) XXX_Size() int {
	return xxx_messageInfo_ClusterLoadAssignment.Size(m)
}
func (m *ClusterLoadAssignment) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterLoadAssignment.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterLoadAssignment proto.InternalMessageInfo

func (m *ClusterLoadAssignment) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *ClusterLoadAssignment) GetEndpoints() []*LocalityLbEndpoints {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func (m *ClusterLoadAssignment) GetNamedEndpoints() map[string]*Endpoint {
	if m != nil {
		return m.NamedEndpoints
	}
	return nil
}

func (m *ClusterLoadAssignment) GetPolicy() *ClusterLoadAssignment_Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

type ClusterLoadAssignment_Policy struct {
	DropOverloads                                []*ClusterLoadAssignment_Policy_DropOverload `protobuf:"bytes,2,rep,name=drop_overloads,json=dropOverloads,proto3" json:"drop_overloads,omitempty"`
	OverprovisioningFactor                       *wrappers.UInt32Value                        `protobuf:"bytes,3,opt,name=overprovisioning_factor,json=overprovisioningFactor,proto3" json:"overprovisioning_factor,omitempty"`
	EndpointStaleAfter                           *duration.Duration                           `protobuf:"bytes,4,opt,name=endpoint_stale_after,json=endpointStaleAfter,proto3" json:"endpoint_stale_after,omitempty"`
	HiddenEnvoyDeprecatedDisableOverprovisioning bool                                         `protobuf:"varint,5,opt,name=hidden_envoy_deprecated_disable_overprovisioning,json=hiddenEnvoyDeprecatedDisableOverprovisioning,proto3" json:"hidden_envoy_deprecated_disable_overprovisioning,omitempty"` // Deprecated: Do not use.
	XXX_NoUnkeyedLiteral                         struct{}                                     `json:"-"`
	XXX_unrecognized                             []byte                                       `json:"-"`
	XXX_sizecache                                int32                                        `json:"-"`
}

func (m *ClusterLoadAssignment_Policy) Reset()         { *m = ClusterLoadAssignment_Policy{} }
func (m *ClusterLoadAssignment_Policy) String() string { return proto.CompactTextString(m) }
func (*ClusterLoadAssignment_Policy) ProtoMessage()    {}
func (*ClusterLoadAssignment_Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_3459cc50e28e3760, []int{0, 0}
}

func (m *ClusterLoadAssignment_Policy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterLoadAssignment_Policy.Unmarshal(m, b)
}
func (m *ClusterLoadAssignment_Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterLoadAssignment_Policy.Marshal(b, m, deterministic)
}
func (m *ClusterLoadAssignment_Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterLoadAssignment_Policy.Merge(m, src)
}
func (m *ClusterLoadAssignment_Policy) XXX_Size() int {
	return xxx_messageInfo_ClusterLoadAssignment_Policy.Size(m)
}
func (m *ClusterLoadAssignment_Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterLoadAssignment_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterLoadAssignment_Policy proto.InternalMessageInfo

func (m *ClusterLoadAssignment_Policy) GetDropOverloads() []*ClusterLoadAssignment_Policy_DropOverload {
	if m != nil {
		return m.DropOverloads
	}
	return nil
}

func (m *ClusterLoadAssignment_Policy) GetOverprovisioningFactor() *wrappers.UInt32Value {
	if m != nil {
		return m.OverprovisioningFactor
	}
	return nil
}

func (m *ClusterLoadAssignment_Policy) GetEndpointStaleAfter() *duration.Duration {
	if m != nil {
		return m.EndpointStaleAfter
	}
	return nil
}

// Deprecated: Do not use.
func (m *ClusterLoadAssignment_Policy) GetHiddenEnvoyDeprecatedDisableOverprovisioning() bool {
	if m != nil {
		return m.HiddenEnvoyDeprecatedDisableOverprovisioning
	}
	return false
}

type ClusterLoadAssignment_Policy_DropOverload struct {
	Category             string                `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	DropPercentage       *v3.FractionalPercent `protobuf:"bytes,2,opt,name=drop_percentage,json=dropPercentage,proto3" json:"drop_percentage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ClusterLoadAssignment_Policy_DropOverload) Reset() {
	*m = ClusterLoadAssignment_Policy_DropOverload{}
}
func (m *ClusterLoadAssignment_Policy_DropOverload) String() string { return proto.CompactTextString(m) }
func (*ClusterLoadAssignment_Policy_DropOverload) ProtoMessage()    {}
func (*ClusterLoadAssignment_Policy_DropOverload) Descriptor() ([]byte, []int) {
	return fileDescriptor_3459cc50e28e3760, []int{0, 0, 0}
}

func (m *ClusterLoadAssignment_Policy_DropOverload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload.Unmarshal(m, b)
}
func (m *ClusterLoadAssignment_Policy_DropOverload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload.Marshal(b, m, deterministic)
}
func (m *ClusterLoadAssignment_Policy_DropOverload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload.Merge(m, src)
}
func (m *ClusterLoadAssignment_Policy_DropOverload) XXX_Size() int {
	return xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload.Size(m)
}
func (m *ClusterLoadAssignment_Policy_DropOverload) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload proto.InternalMessageInfo

func (m *ClusterLoadAssignment_Policy_DropOverload) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *ClusterLoadAssignment_Policy_DropOverload) GetDropPercentage() *v3.FractionalPercent {
	if m != nil {
		return m.DropPercentage
	}
	return nil
}

func init() {
	proto.RegisterType((*ClusterLoadAssignment)(nil), "envoy.config.endpoint.v3.ClusterLoadAssignment")
	proto.RegisterMapType((map[string]*Endpoint)(nil), "envoy.config.endpoint.v3.ClusterLoadAssignment.NamedEndpointsEntry")
	proto.RegisterType((*ClusterLoadAssignment_Policy)(nil), "envoy.config.endpoint.v3.ClusterLoadAssignment.Policy")
	proto.RegisterType((*ClusterLoadAssignment_Policy_DropOverload)(nil), "envoy.config.endpoint.v3.ClusterLoadAssignment.Policy.DropOverload")
}

func init() {
	proto.RegisterFile("envoy/config/endpoint/v3/endpoint.proto", fileDescriptor_3459cc50e28e3760)
}

var fileDescriptor_3459cc50e28e3760 = []byte{
	// 699 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x4d, 0x4f, 0x13, 0x4f,
	0x1c, 0xc7, 0xd9, 0x85, 0xf6, 0x5f, 0x86, 0xc7, 0xcc, 0x5f, 0x65, 0xad, 0x48, 0x2a, 0x1a, 0x2d,
	0xa8, 0xbb, 0xa4, 0x4d, 0x08, 0x69, 0xe2, 0x81, 0x52, 0x48, 0x50, 0x02, 0xcd, 0x1a, 0x3c, 0xba,
	0x99, 0xee, 0x4c, 0xeb, 0xe0, 0x76, 0x66, 0x32, 0x3b, 0x5d, 0xed, 0xcd, 0xa3, 0x37, 0xef, 0xbe,
	0x04, 0x5f, 0x82, 0x89, 0x47, 0x13, 0xaf, 0xbe, 0x15, 0x8f, 0x9c, 0xcc, 0xec, 0x13, 0x58, 0x68,
	0x40, 0x6f, 0xd3, 0xfd, 0xfd, 0x3e, 0xdf, 0xef, 0xfc, 0x1e, 0xa6, 0xe0, 0x11, 0x61, 0x11, 0x1f,
	0x3a, 0x3e, 0x67, 0x5d, 0xda, 0x73, 0x08, 0xc3, 0x82, 0x53, 0xa6, 0x9c, 0xa8, 0x9e, 0x9f, 0x6d,
	0x21, 0xb9, 0xe2, 0xd0, 0x8a, 0x13, 0xed, 0x24, 0xd1, 0xce, 0x83, 0x51, 0xbd, 0x5c, 0xbb, 0x52,
	0xc2, 0xf3, 0x79, 0x5f, 0x70, 0x46, 0x98, 0x0a, 0x13, 0xb5, 0xf2, 0x9d, 0x84, 0x51, 0x43, 0x41,
	0x74, 0xa2, 0x20, 0xd2, 0x27, 0x99, 0x55, 0x79, 0xb9, 0xc7, 0x79, 0x2f, 0x20, 0x0e, 0x12, 0xd4,
	0x41, 0x8c, 0x71, 0x85, 0x14, 0xe5, 0x2c, 0x43, 0x57, 0xd2, 0x68, 0xfc, 0xab, 0x33, 0xe8, 0x3a,
	0x78, 0x20, 0xe3, 0x84, 0x71, 0xf1, 0x77, 0x12, 0x09, 0x41, 0x64, 0xc6, 0xdf, 0x1d, 0x60, 0x81,
	0xce, 0xeb, 0x3a, 0xa1, 0x42, 0x6a, 0x90, 0x85, 0xef, 0x5d, 0x08, 0x47, 0x44, 0x86, 0x94, 0x33,
	0xca, 0x7a, 0x69, 0xca, 0x52, 0x84, 0x02, 0x8a, 0x91, 0x22, 0x4e, 0x76, 0x48, 0x02, 0xab, 0xbf,
	0x4a, 0xe0, 0xe6, 0x4e, 0x30, 0x08, 0x15, 0x91, 0x07, 0x1c, 0xe1, 0xed, 0x30, 0xa4, 0x3d, 0xd6,
	0x27, 0x4c, 0xc1, 0x75, 0x30, 0xeb, 0x27, 0x01, 0x8f, 0xa1, 0x3e, 0xb1, 0x8c, 0x8a, 0x51, 0x9d,
	0x6e, 0xfe, 0x77, 0xda, 0x9c, 0x92, 0x66, 0xc5, 0x70, 0x67, 0xd2, 0xe0, 0x21, 0xea, 0x13, 0xf8,
	0x02, 0x4c, 0x67, 0x8d, 0x0b, 0x2d, 0xb3, 0x32, 0x59, 0x9d, 0xa9, 0x3d, 0xb5, 0xc7, 0x75, 0xdf,
	0x3e, 0xe0, 0x3e, 0x0a, 0xa8, 0x1a, 0x1e, 0x74, 0x76, 0x33, 0xc8, 0x3d, 0xe3, 0x61, 0x00, 0x16,
	0xb4, 0x21, 0xf6, 0xce, 0x24, 0x0b, 0xb1, 0xe4, 0xce, 0x78, 0xc9, 0x4b, 0x4b, 0xb0, 0xf5, 0xdd,
	0x70, 0xee, 0xb1, 0xcb, 0x94, 0x1c, 0xba, 0xf3, 0xec, 0x8f, 0x8f, 0xf0, 0x10, 0x14, 0x05, 0x0f,
	0xa8, 0x3f, 0xb4, 0xa6, 0x2a, 0x46, 0x75, 0xa6, 0xb6, 0xf9, 0xb7, 0x26, 0xed, 0x98, 0x76, 0x53,
	0x95, 0xf2, 0xa7, 0x02, 0x28, 0x26, 0x9f, 0xe0, 0x09, 0x98, 0xc7, 0x92, 0x0b, 0x8f, 0x47, 0x44,
	0x06, 0x1c, 0xe1, 0xac, 0x35, 0x3b, 0xff, 0x66, 0x61, 0xb7, 0x24, 0x17, 0x47, 0xa9, 0x96, 0x3b,
	0x87, 0xcf, 0xfd, 0x0a, 0xe1, 0x6b, 0xb0, 0xa4, 0x6d, 0x84, 0xe4, 0x11, 0x4d, 0x47, 0xef, 0x75,
	0x91, 0xaf, 0xb8, 0xb4, 0x26, 0xe3, 0xba, 0x96, 0xed, 0x64, 0xc9, 0xec, 0x6c, 0xc9, 0xec, 0xe3,
	0x7d, 0xa6, 0xea, 0xb5, 0x57, 0x28, 0x18, 0x90, 0x78, 0xac, 0xeb, 0x66, 0x65, 0xc2, 0xbd, 0x35,
	0xaa, 0xb2, 0x17, 0x8b, 0xc0, 0x63, 0x70, 0x23, 0x7f, 0x1a, 0xa1, 0x42, 0x01, 0xf1, 0x50, 0x57,
	0x11, 0x99, 0x36, 0xed, 0xf6, 0x05, 0xf1, 0x56, 0xba, 0xe1, 0xcd, 0xd2, 0x69, 0xb3, 0xf0, 0xc5,
	0x30, 0xd7, 0x27, 0x5c, 0x98, 0x09, 0xbc, 0xd4, 0xfc, 0xb6, 0xc6, 0xe1, 0x09, 0xd8, 0x78, 0x43,
	0x31, 0x26, 0xcc, 0x8b, 0x5b, 0xe2, 0x61, 0x22, 0x24, 0xf1, 0x91, 0x22, 0xd8, 0xc3, 0x34, 0x44,
	0x9d, 0x80, 0x78, 0xa3, 0x17, 0xb2, 0x0a, 0x15, 0xa3, 0x5a, 0x6a, 0x9a, 0x96, 0xe1, 0x3e, 0x49,
	0xd8, 0x5d, 0x8d, 0xb6, 0x72, 0xb2, 0x95, 0x80, 0x47, 0x23, 0x5c, 0xf9, 0x9b, 0x01, 0x66, 0xcf,
	0xb7, 0x10, 0xde, 0x07, 0x25, 0x0d, 0xf4, 0xb8, 0x1c, 0x8e, 0x6e, 0x77, 0x1e, 0x80, 0xfb, 0x60,
	0x21, 0x1e, 0x62, 0xfa, 0xde, 0x51, 0x8f, 0x58, 0x66, 0x5c, 0x73, 0x25, 0x9d, 0xa2, 0xfe, 0x43,
	0xd0, 0xa3, 0xdb, 0x93, 0xc8, 0xd7, 0x15, 0xa3, 0xa0, 0x9d, 0xa4, 0xba, 0xf1, 0xf4, 0xdb, 0x39,
	0xd7, 0x78, 0xf6, 0xf9, 0xfb, 0xc7, 0x95, 0x2d, 0xb0, 0x99, 0x70, 0x48, 0x50, 0x3b, 0xaa, 0x5d,
	0x7f, 0xe2, 0x8d, 0x0d, 0x8d, 0x3f, 0x06, 0x6b, 0xd7, 0xc6, 0x9f, 0x4f, 0x95, 0x8c, 0x45, 0xb3,
	0x4c, 0xc0, 0xff, 0x97, 0x3c, 0x04, 0xb8, 0x08, 0x26, 0xdf, 0x92, 0xb4, 0x70, 0x57, 0x1f, 0xe1,
	0x16, 0x28, 0x44, 0x7a, 0x1b, 0xd2, 0x02, 0x57, 0xc7, 0xaf, 0x69, 0x26, 0xe5, 0x26, 0x40, 0xc3,
	0xdc, 0x32, 0x1a, 0x6b, 0xfa, 0x7a, 0x0f, 0xc0, 0xea, 0xd5, 0xd7, 0x6b, 0x6e, 0x7f, 0xfd, 0xf0,
	0xe3, 0x67, 0xd1, 0x5c, 0x34, 0xc1, 0x43, 0xca, 0x13, 0x17, 0x21, 0xf9, 0xfb, 0xe1, 0x58, 0xc3,
	0xe6, 0x5c, 0xe6, 0xd8, 0xd6, 0x0b, 0xd6, 0x36, 0x3a, 0xc5, 0x78, 0xd3, 0xea, 0xbf, 0x03, 0x00,
	0x00, 0xff, 0xff, 0xfc, 0xd6, 0x14, 0x94, 0x0d, 0x06, 0x00, 0x00,
}
