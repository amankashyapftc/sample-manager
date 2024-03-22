// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.14.0
// source: proto/sample-manager.proto

package sample_manager

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

type GetSampleItemIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClmSegments []string `protobuf:"bytes,1,rep,name=clm_segments,json=clmSegments,proto3" json:"clm_segments,omitempty"`
	ItemId      string   `protobuf:"bytes,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
}

func (x *GetSampleItemIDRequest) Reset() {
	*x = GetSampleItemIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sample_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSampleItemIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSampleItemIDRequest) ProtoMessage() {}

func (x *GetSampleItemIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sample_manager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSampleItemIDRequest.ProtoReflect.Descriptor instead.
func (*GetSampleItemIDRequest) Descriptor() ([]byte, []int) {
	return file_proto_sample_manager_proto_rawDescGZIP(), []int{0}
}

func (x *GetSampleItemIDRequest) GetClmSegments() []string {
	if x != nil {
		return x.ClmSegments
	}
	return nil
}

func (x *GetSampleItemIDRequest) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

type GetSampleItemIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SampleItemId string `protobuf:"bytes,1,opt,name=sample_item_id,json=sampleItemId,proto3" json:"sample_item_id,omitempty"`
}

func (x *GetSampleItemIDResponse) Reset() {
	*x = GetSampleItemIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sample_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSampleItemIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSampleItemIDResponse) ProtoMessage() {}

func (x *GetSampleItemIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sample_manager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSampleItemIDResponse.ProtoReflect.Descriptor instead.
func (*GetSampleItemIDResponse) Descriptor() ([]byte, []int) {
	return file_proto_sample_manager_proto_rawDescGZIP(), []int{1}
}

func (x *GetSampleItemIDResponse) GetSampleItemId() string {
	if x != nil {
		return x.SampleItemId
	}
	return ""
}

type CreateSampleItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SampleItemId string   `protobuf:"bytes,1,opt,name=sample_item_id,json=sampleItemId,proto3" json:"sample_item_id,omitempty"`
	ClmSegments  []string `protobuf:"bytes,2,rep,name=clm_segments,json=clmSegments,proto3" json:"clm_segments,omitempty"`
	ItemId       string   `protobuf:"bytes,3,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
}

func (x *CreateSampleItemRequest) Reset() {
	*x = CreateSampleItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sample_manager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSampleItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSampleItemRequest) ProtoMessage() {}

func (x *CreateSampleItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sample_manager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSampleItemRequest.ProtoReflect.Descriptor instead.
func (*CreateSampleItemRequest) Descriptor() ([]byte, []int) {
	return file_proto_sample_manager_proto_rawDescGZIP(), []int{2}
}

func (x *CreateSampleItemRequest) GetSampleItemId() string {
	if x != nil {
		return x.SampleItemId
	}
	return ""
}

func (x *CreateSampleItemRequest) GetClmSegments() []string {
	if x != nil {
		return x.ClmSegments
	}
	return nil
}

func (x *CreateSampleItemRequest) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

type CreateSampleItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateSampleItemResponse) Reset() {
	*x = CreateSampleItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sample_manager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSampleItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSampleItemResponse) ProtoMessage() {}

func (x *CreateSampleItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sample_manager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSampleItemResponse.ProtoReflect.Descriptor instead.
func (*CreateSampleItemResponse) Descriptor() ([]byte, []int) {
	return file_proto_sample_manager_proto_rawDescGZIP(), []int{3}
}

func (x *CreateSampleItemResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_sample_manager_proto protoreflect.FileDescriptor

var file_proto_sample_manager_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2d, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x16,
	0x47, 0x65, 0x74, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x44, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6c, 0x6d, 0x5f, 0x73, 0x65,
	0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c,
	0x6d, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65,
	0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d,
	0x49, 0x64, 0x22, 0x3f, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a,
	0x0e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x49, 0x74, 0x65,
	0x6d, 0x49, 0x64, 0x22, 0x7b, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24,
	0x0a, 0x0e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6c, 0x6d, 0x5f, 0x73, 0x65, 0x67, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x6d, 0x53,
	0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64,
	0x22, 0x34, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x9e, 0x01, 0x0a, 0x0d, 0x53, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x44, 0x12, 0x17, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x44, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47,
	0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x18, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x12, 0x5a, 0x10, 0x2e, 0x2f, 0x73, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_sample_manager_proto_rawDescOnce sync.Once
	file_proto_sample_manager_proto_rawDescData = file_proto_sample_manager_proto_rawDesc
)

func file_proto_sample_manager_proto_rawDescGZIP() []byte {
	file_proto_sample_manager_proto_rawDescOnce.Do(func() {
		file_proto_sample_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_sample_manager_proto_rawDescData)
	})
	return file_proto_sample_manager_proto_rawDescData
}

var file_proto_sample_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_sample_manager_proto_goTypes = []interface{}{
	(*GetSampleItemIDRequest)(nil),   // 0: GetSampleItemIDRequest
	(*GetSampleItemIDResponse)(nil),  // 1: GetSampleItemIDResponse
	(*CreateSampleItemRequest)(nil),  // 2: CreateSampleItemRequest
	(*CreateSampleItemResponse)(nil), // 3: CreateSampleItemResponse
}
var file_proto_sample_manager_proto_depIdxs = []int32{
	0, // 0: SampleService.GetSampleItemID:input_type -> GetSampleItemIDRequest
	2, // 1: SampleService.CreateSampleItem:input_type -> CreateSampleItemRequest
	1, // 2: SampleService.GetSampleItemID:output_type -> GetSampleItemIDResponse
	3, // 3: SampleService.CreateSampleItem:output_type -> CreateSampleItemResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_sample_manager_proto_init() }
func file_proto_sample_manager_proto_init() {
	if File_proto_sample_manager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_sample_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSampleItemIDRequest); i {
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
		file_proto_sample_manager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSampleItemIDResponse); i {
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
		file_proto_sample_manager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSampleItemRequest); i {
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
		file_proto_sample_manager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSampleItemResponse); i {
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
			RawDescriptor: file_proto_sample_manager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_sample_manager_proto_goTypes,
		DependencyIndexes: file_proto_sample_manager_proto_depIdxs,
		MessageInfos:      file_proto_sample_manager_proto_msgTypes,
	}.Build()
	File_proto_sample_manager_proto = out.File
	file_proto_sample_manager_proto_rawDesc = nil
	file_proto_sample_manager_proto_goTypes = nil
	file_proto_sample_manager_proto_depIdxs = nil
}
