// Copyright Istio Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        (unknown)
// source: dns/proto/nds.proto

package istio_networking_nds_v1

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

// Table of hostnames and their IPs to br used for DNS resolution at the agent
// Sent by istiod to istio agents via xds
type NameTable struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Map of hostname to resolution attributes.
	Table         map[string]*NameTable_NameInfo `protobuf:"bytes,1,rep,name=table,proto3" json:"table,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NameTable) Reset() {
	*x = NameTable{}
	mi := &file_dns_proto_nds_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NameTable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NameTable) ProtoMessage() {}

func (x *NameTable) ProtoReflect() protoreflect.Message {
	mi := &file_dns_proto_nds_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NameTable.ProtoReflect.Descriptor instead.
func (*NameTable) Descriptor() ([]byte, []int) {
	return file_dns_proto_nds_proto_rawDescGZIP(), []int{0}
}

func (x *NameTable) GetTable() map[string]*NameTable_NameInfo {
	if x != nil {
		return x.Table
	}
	return nil
}

type NameTable_NameInfo struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// List of IPs for the host.
	Ips []string `protobuf:"bytes,1,rep,name=ips,proto3" json:"ips,omitempty"`
	// The name of the service registry containing the service (e.g. 'Kubernetes').
	Registry string `protobuf:"bytes,2,opt,name=registry,proto3" json:"registry,omitempty"`
	// The k8s service name. Only applies when registry=`Kubernetes`
	Shortname string `protobuf:"bytes,3,opt,name=shortname,proto3" json:"shortname,omitempty"`
	// The k8s namespace for the service. Only applies when registry=`Kubernetes`
	Namespace string `protobuf:"bytes,4,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Deprecated. Was added for experimentation only.
	//
	// Deprecated: Marked as deprecated in dns/proto/nds.proto.
	AltHosts      []string `protobuf:"bytes,5,rep,name=alt_hosts,json=altHosts,proto3" json:"alt_hosts,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NameTable_NameInfo) Reset() {
	*x = NameTable_NameInfo{}
	mi := &file_dns_proto_nds_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NameTable_NameInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NameTable_NameInfo) ProtoMessage() {}

func (x *NameTable_NameInfo) ProtoReflect() protoreflect.Message {
	mi := &file_dns_proto_nds_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NameTable_NameInfo.ProtoReflect.Descriptor instead.
func (*NameTable_NameInfo) Descriptor() ([]byte, []int) {
	return file_dns_proto_nds_proto_rawDescGZIP(), []int{0, 0}
}

func (x *NameTable_NameInfo) GetIps() []string {
	if x != nil {
		return x.Ips
	}
	return nil
}

func (x *NameTable_NameInfo) GetRegistry() string {
	if x != nil {
		return x.Registry
	}
	return ""
}

func (x *NameTable_NameInfo) GetShortname() string {
	if x != nil {
		return x.Shortname
	}
	return ""
}

func (x *NameTable_NameInfo) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

// Deprecated: Marked as deprecated in dns/proto/nds.proto.
func (x *NameTable_NameInfo) GetAltHosts() []string {
	if x != nil {
		return x.AltHosts
	}
	return nil
}

var File_dns_proto_nds_proto protoreflect.FileDescriptor

var file_dns_proto_nds_proto_rawDesc = []byte{
	0x0a, 0x13, 0x64, 0x6e, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x64, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x22, 0xcf,
	0x02, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x43, 0x0a, 0x05,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x69, 0x73,
	0x74, 0x69, 0x6f, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x6e,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x1a, 0x95, 0x01, 0x0a, 0x08, 0x4e, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10,
	0x0a, 0x03, 0x69, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x70, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x09, 0x61, 0x6c, 0x74, 0x5f,
	0x68, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52,
	0x08, 0x61, 0x6c, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x1a, 0x65, 0x0a, 0x0a, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x41, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f,
	0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x6e, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x4e, 0x61, 0x6d,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x42, 0x36, 0x5a, 0x34, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x69, 0x73, 0x74,
	0x69, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x64, 0x6e, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e,
	0x67, 0x5f, 0x6e, 0x64, 0x73, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dns_proto_nds_proto_rawDescOnce sync.Once
	file_dns_proto_nds_proto_rawDescData = file_dns_proto_nds_proto_rawDesc
)

func file_dns_proto_nds_proto_rawDescGZIP() []byte {
	file_dns_proto_nds_proto_rawDescOnce.Do(func() {
		file_dns_proto_nds_proto_rawDescData = protoimpl.X.CompressGZIP(file_dns_proto_nds_proto_rawDescData)
	})
	return file_dns_proto_nds_proto_rawDescData
}

var file_dns_proto_nds_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_dns_proto_nds_proto_goTypes = []any{
	(*NameTable)(nil),          // 0: istio.networking.nds.v1.NameTable
	(*NameTable_NameInfo)(nil), // 1: istio.networking.nds.v1.NameTable.NameInfo
	nil,                        // 2: istio.networking.nds.v1.NameTable.TableEntry
}
var file_dns_proto_nds_proto_depIdxs = []int32{
	2, // 0: istio.networking.nds.v1.NameTable.table:type_name -> istio.networking.nds.v1.NameTable.TableEntry
	1, // 1: istio.networking.nds.v1.NameTable.TableEntry.value:type_name -> istio.networking.nds.v1.NameTable.NameInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_dns_proto_nds_proto_init() }
func file_dns_proto_nds_proto_init() {
	if File_dns_proto_nds_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_dns_proto_nds_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_dns_proto_nds_proto_goTypes,
		DependencyIndexes: file_dns_proto_nds_proto_depIdxs,
		MessageInfos:      file_dns_proto_nds_proto_msgTypes,
	}.Build()
	File_dns_proto_nds_proto = out.File
	file_dns_proto_nds_proto_rawDesc = nil
	file_dns_proto_nds_proto_goTypes = nil
	file_dns_proto_nds_proto_depIdxs = nil
}
