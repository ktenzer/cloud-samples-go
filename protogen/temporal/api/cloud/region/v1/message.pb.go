// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: temporal/api/cloud/region/v1/message.proto

package region

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

type Region struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of the temporal cloud region.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the cloud provider that's hosting the region.
	// Currently only "aws" is supported.
	CloudProvider string `protobuf:"bytes,2,opt,name=cloud_provider,json=cloudProvider,proto3" json:"cloud_provider,omitempty"`
	// The region identifier as defined by the cloud provider.
	CloudProviderRegion string `protobuf:"bytes,3,opt,name=cloud_provider_region,json=cloudProviderRegion,proto3" json:"cloud_provider_region,omitempty"`
	// The human readable location of the region.
	Location string `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *Region) Reset() {
	*x = Region{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_api_cloud_region_v1_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Region) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Region) ProtoMessage() {}

func (x *Region) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_cloud_region_v1_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Region.ProtoReflect.Descriptor instead.
func (*Region) Descriptor() ([]byte, []int) {
	return file_temporal_api_cloud_region_v1_message_proto_rawDescGZIP(), []int{0}
}

func (x *Region) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Region) GetCloudProvider() string {
	if x != nil {
		return x.CloudProvider
	}
	return ""
}

func (x *Region) GetCloudProviderRegion() string {
	if x != nil {
		return x.CloudProviderRegion
	}
	return ""
}

func (x *Region) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

var File_temporal_api_cloud_region_v1_message_proto protoreflect.FileDescriptor

var file_temporal_api_cloud_region_v1_message_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x74, 0x65,
	0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x22, 0x8f, 0x01, 0x0a, 0x06, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x15,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x2b, 0x5a, 0x29,
	0x67, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x69, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2f,
	0x76, 0x31, 0x3b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_temporal_api_cloud_region_v1_message_proto_rawDescOnce sync.Once
	file_temporal_api_cloud_region_v1_message_proto_rawDescData = file_temporal_api_cloud_region_v1_message_proto_rawDesc
)

func file_temporal_api_cloud_region_v1_message_proto_rawDescGZIP() []byte {
	file_temporal_api_cloud_region_v1_message_proto_rawDescOnce.Do(func() {
		file_temporal_api_cloud_region_v1_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_temporal_api_cloud_region_v1_message_proto_rawDescData)
	})
	return file_temporal_api_cloud_region_v1_message_proto_rawDescData
}

var file_temporal_api_cloud_region_v1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_temporal_api_cloud_region_v1_message_proto_goTypes = []interface{}{
	(*Region)(nil), // 0: temporal.api.cloud.region.v1.Region
}
var file_temporal_api_cloud_region_v1_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_temporal_api_cloud_region_v1_message_proto_init() }
func file_temporal_api_cloud_region_v1_message_proto_init() {
	if File_temporal_api_cloud_region_v1_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_temporal_api_cloud_region_v1_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Region); i {
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
			RawDescriptor: file_temporal_api_cloud_region_v1_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_api_cloud_region_v1_message_proto_goTypes,
		DependencyIndexes: file_temporal_api_cloud_region_v1_message_proto_depIdxs,
		MessageInfos:      file_temporal_api_cloud_region_v1_message_proto_msgTypes,
	}.Build()
	File_temporal_api_cloud_region_v1_message_proto = out.File
	file_temporal_api_cloud_region_v1_message_proto_rawDesc = nil
	file_temporal_api_cloud_region_v1_message_proto_goTypes = nil
	file_temporal_api_cloud_region_v1_message_proto_depIdxs = nil
}
