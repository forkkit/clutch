// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: chaos/serverexperimentation/v1/serverexperimentation.proto

package serverexperimentationv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/lyft/clutch/backend/api/api/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type TestSpecification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Config:
	//	*TestSpecification_Abort
	//	*TestSpecification_Latency
	Config isTestSpecification_Config `protobuf_oneof:"config"`
}

func (x *TestSpecification) Reset() {
	*x = TestSpecification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chaos_serverexperimentation_v1_serverexperimentation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestSpecification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestSpecification) ProtoMessage() {}

func (x *TestSpecification) ProtoReflect() protoreflect.Message {
	mi := &file_chaos_serverexperimentation_v1_serverexperimentation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestSpecification.ProtoReflect.Descriptor instead.
func (*TestSpecification) Descriptor() ([]byte, []int) {
	return file_chaos_serverexperimentation_v1_serverexperimentation_proto_rawDescGZIP(), []int{0}
}

func (m *TestSpecification) GetConfig() isTestSpecification_Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (x *TestSpecification) GetAbort() *AbortFault {
	if x, ok := x.GetConfig().(*TestSpecification_Abort); ok {
		return x.Abort
	}
	return nil
}

func (x *TestSpecification) GetLatency() *LatencyFault {
	if x, ok := x.GetConfig().(*TestSpecification_Latency); ok {
		return x.Latency
	}
	return nil
}

type isTestSpecification_Config interface {
	isTestSpecification_Config()
}

type TestSpecification_Abort struct {
	Abort *AbortFault `protobuf:"bytes,1,opt,name=abort,proto3,oneof"`
}

type TestSpecification_Latency struct {
	Latency *LatencyFault `protobuf:"bytes,2,opt,name=latency,proto3,oneof"`
}

func (*TestSpecification_Abort) isTestSpecification_Config() {}

func (*TestSpecification_Latency) isTestSpecification_Config() {}

// Targets requests from downstream_cluster -> upstream_cluster
type ClusterPairTarget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the downstream cluster
	DownstreamCluster string `protobuf:"bytes,1,opt,name=downstream_cluster,json=downstreamCluster,proto3" json:"downstream_cluster,omitempty"`
	// The name of the upstream cluster
	UpstreamCluster string `protobuf:"bytes,2,opt,name=upstream_cluster,json=upstreamCluster,proto3" json:"upstream_cluster,omitempty"`
}

func (x *ClusterPairTarget) Reset() {
	*x = ClusterPairTarget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chaos_serverexperimentation_v1_serverexperimentation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterPairTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterPairTarget) ProtoMessage() {}

func (x *ClusterPairTarget) ProtoReflect() protoreflect.Message {
	mi := &file_chaos_serverexperimentation_v1_serverexperimentation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterPairTarget.ProtoReflect.Descriptor instead.
func (*ClusterPairTarget) Descriptor() ([]byte, []int) {
	return file_chaos_serverexperimentation_v1_serverexperimentation_proto_rawDescGZIP(), []int{1}
}

func (x *ClusterPairTarget) GetDownstreamCluster() string {
	if x != nil {
		return x.DownstreamCluster
	}
	return ""
}

func (x *ClusterPairTarget) GetUpstreamCluster() string {
	if x != nil {
		return x.UpstreamCluster
	}
	return ""
}

type AbortFault struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Target:
	//	*AbortFault_ClusterPair
	Target isAbortFault_Target `protobuf_oneof:"target"`
	// The percentage of requests that will be faulted.
	Percent float32 `protobuf:"fixed32,2,opt,name=percent,proto3" json:"percent,omitempty"`
	// The abort HTTP status that will be returned.
	HttpStatus int32 `protobuf:"varint,3,opt,name=http_status,json=httpStatus,proto3" json:"http_status,omitempty"`
}

func (x *AbortFault) Reset() {
	*x = AbortFault{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chaos_serverexperimentation_v1_serverexperimentation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbortFault) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbortFault) ProtoMessage() {}

func (x *AbortFault) ProtoReflect() protoreflect.Message {
	mi := &file_chaos_serverexperimentation_v1_serverexperimentation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AbortFault.ProtoReflect.Descriptor instead.
func (*AbortFault) Descriptor() ([]byte, []int) {
	return file_chaos_serverexperimentation_v1_serverexperimentation_proto_rawDescGZIP(), []int{2}
}

func (m *AbortFault) GetTarget() isAbortFault_Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (x *AbortFault) GetClusterPair() *ClusterPairTarget {
	if x, ok := x.GetTarget().(*AbortFault_ClusterPair); ok {
		return x.ClusterPair
	}
	return nil
}

func (x *AbortFault) GetPercent() float32 {
	if x != nil {
		return x.Percent
	}
	return 0
}

func (x *AbortFault) GetHttpStatus() int32 {
	if x != nil {
		return x.HttpStatus
	}
	return 0
}

type isAbortFault_Target interface {
	isAbortFault_Target()
}

type AbortFault_ClusterPair struct {
	ClusterPair *ClusterPairTarget `protobuf:"bytes,1,opt,name=cluster_pair,json=clusterPair,proto3,oneof"`
}

func (*AbortFault_ClusterPair) isAbortFault_Target() {}

type LatencyFault struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Target:
	//	*LatencyFault_ClusterPair
	Target isLatencyFault_Target `protobuf_oneof:"target"`
	// The percentage of requests that will be slowed down.
	Percent float32 `protobuf:"fixed32,2,opt,name=percent,proto3" json:"percent,omitempty"`
	// The latency duration in milliseconds.
	DurationMs int32 `protobuf:"varint,3,opt,name=duration_ms,json=durationMs,proto3" json:"duration_ms,omitempty"`
}

func (x *LatencyFault) Reset() {
	*x = LatencyFault{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chaos_serverexperimentation_v1_serverexperimentation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LatencyFault) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LatencyFault) ProtoMessage() {}

func (x *LatencyFault) ProtoReflect() protoreflect.Message {
	mi := &file_chaos_serverexperimentation_v1_serverexperimentation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LatencyFault.ProtoReflect.Descriptor instead.
func (*LatencyFault) Descriptor() ([]byte, []int) {
	return file_chaos_serverexperimentation_v1_serverexperimentation_proto_rawDescGZIP(), []int{3}
}

func (m *LatencyFault) GetTarget() isLatencyFault_Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (x *LatencyFault) GetClusterPair() *ClusterPairTarget {
	if x, ok := x.GetTarget().(*LatencyFault_ClusterPair); ok {
		return x.ClusterPair
	}
	return nil
}

func (x *LatencyFault) GetPercent() float32 {
	if x != nil {
		return x.Percent
	}
	return 0
}

func (x *LatencyFault) GetDurationMs() int32 {
	if x != nil {
		return x.DurationMs
	}
	return 0
}

type isLatencyFault_Target interface {
	isLatencyFault_Target()
}

type LatencyFault_ClusterPair struct {
	ClusterPair *ClusterPairTarget `protobuf:"bytes,1,opt,name=cluster_pair,json=clusterPair,proto3,oneof"`
}

func (*LatencyFault_ClusterPair) isLatencyFault_Target() {}

var File_chaos_serverexperimentation_v1_serverexperimentation_proto protoreflect.FileDescriptor

var file_chaos_serverexperimentation_v1_serverexperimentation_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x63, 0x68, 0x61, 0x6f, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x65, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x25, 0x63, 0x6c,
	0x75, 0x74, 0x63, 0x68, 0x2e, 0x63, 0x68, 0x61, 0x6f, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x01, 0x0a, 0x11, 0x54, 0x65, 0x73, 0x74, 0x53, 0x70, 0x65,
	0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x49, 0x0a, 0x05, 0x61, 0x62,
	0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x63, 0x6c, 0x75, 0x74,
	0x63, 0x68, 0x2e, 0x63, 0x68, 0x61, 0x6f, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x65,
	0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x62, 0x6f, 0x72, 0x74, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x05,
	0x61, 0x62, 0x6f, 0x72, 0x74, 0x12, 0x4f, 0x0a, 0x07, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e,
	0x63, 0x68, 0x61, 0x6f, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x65, 0x78, 0x70, 0x65,
	0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x07, 0x6c,
	0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x42, 0x08, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x22, 0x7f, 0x0a, 0x11, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x61, 0x69, 0x72, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x36, 0x0a, 0x12, 0x64, 0x6f, 0x77, 0x6e, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x11, 0x64, 0x6f, 0x77, 0x6e,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x32, 0x0a,
	0x10, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01,
	0x52, 0x0f, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x22, 0xcd, 0x01, 0x0a, 0x0a, 0x41, 0x62, 0x6f, 0x72, 0x74, 0x46, 0x61, 0x75, 0x6c, 0x74,
	0x12, 0x5d, 0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x69, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e,
	0x63, 0x68, 0x61, 0x6f, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x65, 0x78, 0x70, 0x65,
	0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x61, 0x69, 0x72, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x48, 0x00, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x61, 0x69, 0x72, 0x12,
	0x29, 0x0a, 0x07, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02,
	0x42, 0x0f, 0xfa, 0x42, 0x0c, 0x0a, 0x0a, 0x1d, 0x00, 0x00, 0xc8, 0x42, 0x25, 0x00, 0x00, 0x00,
	0x00, 0x52, 0x07, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x0b, 0x68, 0x74,
	0x74, 0x70, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42,
	0x0a, 0xfa, 0x42, 0x07, 0x1a, 0x05, 0x10, 0xd8, 0x04, 0x20, 0x63, 0x52, 0x0a, 0x68, 0x74, 0x74,
	0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x08, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x22, 0xcc, 0x01, 0x0a, 0x0c, 0x4c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x46, 0x61, 0x75,
	0x6c, 0x74, 0x12, 0x5d, 0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x70, 0x61,
	0x69, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63,
	0x68, 0x2e, 0x63, 0x68, 0x61, 0x6f, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x65, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x61, 0x69, 0x72, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x48, 0x00, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x61, 0x69,
	0x72, 0x12, 0x29, 0x0a, 0x07, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x02, 0x42, 0x0f, 0xfa, 0x42, 0x0c, 0x0a, 0x0a, 0x1d, 0x00, 0x00, 0xc8, 0x42, 0x25, 0x00,
	0x00, 0x00, 0x00, 0x52, 0x07, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x0b,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x20, 0x00, 0x52, 0x0a, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x73, 0x42, 0x08, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x42, 0x19, 0x5a, 0x17, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69,
	0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_chaos_serverexperimentation_v1_serverexperimentation_proto_rawDescOnce sync.Once
	file_chaos_serverexperimentation_v1_serverexperimentation_proto_rawDescData = file_chaos_serverexperimentation_v1_serverexperimentation_proto_rawDesc
)

func file_chaos_serverexperimentation_v1_serverexperimentation_proto_rawDescGZIP() []byte {
	file_chaos_serverexperimentation_v1_serverexperimentation_proto_rawDescOnce.Do(func() {
		file_chaos_serverexperimentation_v1_serverexperimentation_proto_rawDescData = protoimpl.X.CompressGZIP(file_chaos_serverexperimentation_v1_serverexperimentation_proto_rawDescData)
	})
	return file_chaos_serverexperimentation_v1_serverexperimentation_proto_rawDescData
}

var file_chaos_serverexperimentation_v1_serverexperimentation_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_chaos_serverexperimentation_v1_serverexperimentation_proto_goTypes = []interface{}{
	(*TestSpecification)(nil), // 0: clutch.chaos.serverexperimentation.v1.TestSpecification
	(*ClusterPairTarget)(nil), // 1: clutch.chaos.serverexperimentation.v1.ClusterPairTarget
	(*AbortFault)(nil),        // 2: clutch.chaos.serverexperimentation.v1.AbortFault
	(*LatencyFault)(nil),      // 3: clutch.chaos.serverexperimentation.v1.LatencyFault
}
var file_chaos_serverexperimentation_v1_serverexperimentation_proto_depIdxs = []int32{
	2, // 0: clutch.chaos.serverexperimentation.v1.TestSpecification.abort:type_name -> clutch.chaos.serverexperimentation.v1.AbortFault
	3, // 1: clutch.chaos.serverexperimentation.v1.TestSpecification.latency:type_name -> clutch.chaos.serverexperimentation.v1.LatencyFault
	1, // 2: clutch.chaos.serverexperimentation.v1.AbortFault.cluster_pair:type_name -> clutch.chaos.serverexperimentation.v1.ClusterPairTarget
	1, // 3: clutch.chaos.serverexperimentation.v1.LatencyFault.cluster_pair:type_name -> clutch.chaos.serverexperimentation.v1.ClusterPairTarget
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_chaos_serverexperimentation_v1_serverexperimentation_proto_init() }
func file_chaos_serverexperimentation_v1_serverexperimentation_proto_init() {
	if File_chaos_serverexperimentation_v1_serverexperimentation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chaos_serverexperimentation_v1_serverexperimentation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestSpecification); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chaos_serverexperimentation_v1_serverexperimentation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterPairTarget); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chaos_serverexperimentation_v1_serverexperimentation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AbortFault); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chaos_serverexperimentation_v1_serverexperimentation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LatencyFault); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_chaos_serverexperimentation_v1_serverexperimentation_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*TestSpecification_Abort)(nil),
		(*TestSpecification_Latency)(nil),
	}
	file_chaos_serverexperimentation_v1_serverexperimentation_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*AbortFault_ClusterPair)(nil),
	}
	file_chaos_serverexperimentation_v1_serverexperimentation_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*LatencyFault_ClusterPair)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chaos_serverexperimentation_v1_serverexperimentation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chaos_serverexperimentation_v1_serverexperimentation_proto_goTypes,
		DependencyIndexes: file_chaos_serverexperimentation_v1_serverexperimentation_proto_depIdxs,
		MessageInfos:      file_chaos_serverexperimentation_v1_serverexperimentation_proto_msgTypes,
	}.Build()
	File_chaos_serverexperimentation_v1_serverexperimentation_proto = out.File
	file_chaos_serverexperimentation_v1_serverexperimentation_proto_rawDesc = nil
	file_chaos_serverexperimentation_v1_serverexperimentation_proto_goTypes = nil
	file_chaos_serverexperimentation_v1_serverexperimentation_proto_depIdxs = nil
}