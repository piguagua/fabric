// Copyright the Hyperledger Fabric contributors. All rights reserved.
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: orderer/etcdraft/configuration.proto

package etcdraft

import (
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

// ConfigMetadata is serialized and set as the value of ConsensusType.Metadata in
// a channel configuration when the ConsensusType.Type is set "etcdraft".
type ConfigMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Consenters []*Consenter `protobuf:"bytes,1,rep,name=consenters,proto3" json:"consenters,omitempty"`
	Options    *Options     `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *ConfigMetadata) Reset() {
	*x = ConfigMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderer_etcdraft_configuration_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigMetadata) ProtoMessage() {}

func (x *ConfigMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_orderer_etcdraft_configuration_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigMetadata.ProtoReflect.Descriptor instead.
func (*ConfigMetadata) Descriptor() ([]byte, []int) {
	return file_orderer_etcdraft_configuration_proto_rawDescGZIP(), []int{0}
}

func (x *ConfigMetadata) GetConsenters() []*Consenter {
	if x != nil {
		return x.Consenters
	}
	return nil
}

func (x *ConfigMetadata) GetOptions() *Options {
	if x != nil {
		return x.Options
	}
	return nil
}

// Consenter represents a consenting node (i.e. replica).
type Consenter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host          string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port          uint32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	ClientTlsCert []byte `protobuf:"bytes,3,opt,name=client_tls_cert,json=clientTlsCert,proto3" json:"client_tls_cert,omitempty"`
	ServerTlsCert []byte `protobuf:"bytes,4,opt,name=server_tls_cert,json=serverTlsCert,proto3" json:"server_tls_cert,omitempty"`
}

func (x *Consenter) Reset() {
	*x = Consenter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderer_etcdraft_configuration_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Consenter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Consenter) ProtoMessage() {}

func (x *Consenter) ProtoReflect() protoreflect.Message {
	mi := &file_orderer_etcdraft_configuration_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Consenter.ProtoReflect.Descriptor instead.
func (*Consenter) Descriptor() ([]byte, []int) {
	return file_orderer_etcdraft_configuration_proto_rawDescGZIP(), []int{1}
}

func (x *Consenter) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Consenter) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Consenter) GetClientTlsCert() []byte {
	if x != nil {
		return x.ClientTlsCert
	}
	return nil
}

func (x *Consenter) GetServerTlsCert() []byte {
	if x != nil {
		return x.ServerTlsCert
	}
	return nil
}

// Options to be specified for all the etcd/raft nodes. These can be modified on a
// per-channel basis.
type Options struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TickInterval      string `protobuf:"bytes,1,opt,name=tick_interval,json=tickInterval,proto3" json:"tick_interval,omitempty"` // time duration format, e.g. 500ms
	ElectionTick      uint32 `protobuf:"varint,2,opt,name=election_tick,json=electionTick,proto3" json:"election_tick,omitempty"`
	HeartbeatTick     uint32 `protobuf:"varint,3,opt,name=heartbeat_tick,json=heartbeatTick,proto3" json:"heartbeat_tick,omitempty"`
	MaxInflightBlocks uint32 `protobuf:"varint,4,opt,name=max_inflight_blocks,json=maxInflightBlocks,proto3" json:"max_inflight_blocks,omitempty"`
	// Take snapshot when cumulative data exceeds certain size in bytes.
	SnapshotIntervalSize uint32 `protobuf:"varint,5,opt,name=snapshot_interval_size,json=snapshotIntervalSize,proto3" json:"snapshot_interval_size,omitempty"`
}

func (x *Options) Reset() {
	*x = Options{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderer_etcdraft_configuration_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Options) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Options) ProtoMessage() {}

func (x *Options) ProtoReflect() protoreflect.Message {
	mi := &file_orderer_etcdraft_configuration_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Options.ProtoReflect.Descriptor instead.
func (*Options) Descriptor() ([]byte, []int) {
	return file_orderer_etcdraft_configuration_proto_rawDescGZIP(), []int{2}
}

func (x *Options) GetTickInterval() string {
	if x != nil {
		return x.TickInterval
	}
	return ""
}

func (x *Options) GetElectionTick() uint32 {
	if x != nil {
		return x.ElectionTick
	}
	return 0
}

func (x *Options) GetHeartbeatTick() uint32 {
	if x != nil {
		return x.HeartbeatTick
	}
	return 0
}

func (x *Options) GetMaxInflightBlocks() uint32 {
	if x != nil {
		return x.MaxInflightBlocks
	}
	return 0
}

func (x *Options) GetSnapshotIntervalSize() uint32 {
	if x != nil {
		return x.SnapshotIntervalSize
	}
	return 0
}

var File_orderer_etcdraft_configuration_proto protoreflect.FileDescriptor

var file_orderer_etcdraft_configuration_proto_rawDesc = []byte{
	0x0a, 0x24, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x2f, 0x65, 0x74, 0x63, 0x64, 0x72, 0x61,
	0x66, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x65, 0x74, 0x63, 0x64, 0x72, 0x61, 0x66, 0x74,
	0x22, 0x72, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x33, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65, 0x74, 0x63, 0x64, 0x72, 0x61, 0x66,
	0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x0a, 0x63, 0x6f, 0x6e,
	0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x12, 0x2b, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x65, 0x74, 0x63, 0x64, 0x72,
	0x61, 0x66, 0x74, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x22, 0x83, 0x01, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x6c, 0x73, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x6c, 0x73, 0x43, 0x65,
	0x72, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x6c, 0x73,
	0x5f, 0x63, 0x65, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x54, 0x6c, 0x73, 0x43, 0x65, 0x72, 0x74, 0x22, 0xe0, 0x01, 0x0a, 0x07, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x69, 0x63, 0x6b, 0x5f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74,
	0x69, 0x63, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0c, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x63, 0x6b,
	0x12, 0x25, 0x0a, 0x0e, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x5f, 0x74, 0x69,
	0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62,
	0x65, 0x61, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x12, 0x2e, 0x0a, 0x13, 0x6d, 0x61, 0x78, 0x5f, 0x69,
	0x6e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x6d, 0x61, 0x78, 0x49, 0x6e, 0x66, 0x6c, 0x69, 0x67, 0x68,
	0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x34, 0x0a, 0x16, 0x73, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x14, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f,
	0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x42, 0xc4, 0x01,
	0x0a, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x2e, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x2e, 0x65, 0x74, 0x63, 0x64, 0x72, 0x61, 0x66, 0x74,
	0x42, 0x12, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x66,
	0x61, 0x62, 0x72, 0x69, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2d, 0x67, 0x6f, 0x2d,
	0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x2f, 0x65, 0x74,
	0x63, 0x64, 0x72, 0x61, 0x66, 0x74, 0xa2, 0x02, 0x03, 0x45, 0x58, 0x58, 0xaa, 0x02, 0x08, 0x45,
	0x74, 0x63, 0x64, 0x72, 0x61, 0x66, 0x74, 0xca, 0x02, 0x08, 0x45, 0x74, 0x63, 0x64, 0x72, 0x61,
	0x66, 0x74, 0xe2, 0x02, 0x14, 0x45, 0x74, 0x63, 0x64, 0x72, 0x61, 0x66, 0x74, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x45, 0x74, 0x63, 0x64,
	0x72, 0x61, 0x66, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orderer_etcdraft_configuration_proto_rawDescOnce sync.Once
	file_orderer_etcdraft_configuration_proto_rawDescData = file_orderer_etcdraft_configuration_proto_rawDesc
)

func file_orderer_etcdraft_configuration_proto_rawDescGZIP() []byte {
	file_orderer_etcdraft_configuration_proto_rawDescOnce.Do(func() {
		file_orderer_etcdraft_configuration_proto_rawDescData = protoimpl.X.CompressGZIP(file_orderer_etcdraft_configuration_proto_rawDescData)
	})
	return file_orderer_etcdraft_configuration_proto_rawDescData
}

var file_orderer_etcdraft_configuration_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_orderer_etcdraft_configuration_proto_goTypes = []any{
	(*ConfigMetadata)(nil), // 0: etcdraft.ConfigMetadata
	(*Consenter)(nil),      // 1: etcdraft.Consenter
	(*Options)(nil),        // 2: etcdraft.Options
}
var file_orderer_etcdraft_configuration_proto_depIdxs = []int32{
	1, // 0: etcdraft.ConfigMetadata.consenters:type_name -> etcdraft.Consenter
	2, // 1: etcdraft.ConfigMetadata.options:type_name -> etcdraft.Options
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_orderer_etcdraft_configuration_proto_init() }
func file_orderer_etcdraft_configuration_proto_init() {
	if File_orderer_etcdraft_configuration_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orderer_etcdraft_configuration_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ConfigMetadata); i {
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
		file_orderer_etcdraft_configuration_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Consenter); i {
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
		file_orderer_etcdraft_configuration_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Options); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_orderer_etcdraft_configuration_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_orderer_etcdraft_configuration_proto_goTypes,
		DependencyIndexes: file_orderer_etcdraft_configuration_proto_depIdxs,
		MessageInfos:      file_orderer_etcdraft_configuration_proto_msgTypes,
	}.Build()
	File_orderer_etcdraft_configuration_proto = out.File
	file_orderer_etcdraft_configuration_proto_rawDesc = nil
	file_orderer_etcdraft_configuration_proto_goTypes = nil
	file_orderer_etcdraft_configuration_proto_depIdxs = nil
}