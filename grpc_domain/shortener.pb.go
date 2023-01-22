// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: shortener/shortener.proto

package grpc_domain

import (
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

type CreateLinkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Link string `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *CreateLinkRequest) Reset() {
	*x = CreateLinkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortener_shortener_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLinkRequest) ProtoMessage() {}

func (x *CreateLinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_shortener_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLinkRequest.ProtoReflect.Descriptor instead.
func (*CreateLinkRequest) Descriptor() ([]byte, []int) {
	return file_shortener_shortener_proto_rawDescGZIP(), []int{0}
}

func (x *CreateLinkRequest) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

type CreateLinkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortLink string `protobuf:"bytes,1,opt,name=shortLink,proto3" json:"shortLink,omitempty"`
}

func (x *CreateLinkResponse) Reset() {
	*x = CreateLinkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortener_shortener_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLinkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLinkResponse) ProtoMessage() {}

func (x *CreateLinkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_shortener_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLinkResponse.ProtoReflect.Descriptor instead.
func (*CreateLinkResponse) Descriptor() ([]byte, []int) {
	return file_shortener_shortener_proto_rawDescGZIP(), []int{1}
}

func (x *CreateLinkResponse) GetShortLink() string {
	if x != nil {
		return x.ShortLink
	}
	return ""
}

type GetLinkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Link string `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *GetLinkRequest) Reset() {
	*x = GetLinkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortener_shortener_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLinkRequest) ProtoMessage() {}

func (x *GetLinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_shortener_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLinkRequest.ProtoReflect.Descriptor instead.
func (*GetLinkRequest) Descriptor() ([]byte, []int) {
	return file_shortener_shortener_proto_rawDescGZIP(), []int{2}
}

func (x *GetLinkRequest) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

type GetLinkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrigLink string `protobuf:"bytes,1,opt,name=origLink,proto3" json:"origLink,omitempty"`
}

func (x *GetLinkResponse) Reset() {
	*x = GetLinkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortener_shortener_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLinkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLinkResponse) ProtoMessage() {}

func (x *GetLinkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_shortener_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLinkResponse.ProtoReflect.Descriptor instead.
func (*GetLinkResponse) Descriptor() ([]byte, []int) {
	return file_shortener_shortener_proto_rawDescGZIP(), []int{3}
}

func (x *GetLinkResponse) GetOrigLink() string {
	if x != nil {
		return x.OrigLink
	}
	return ""
}

var File_shortener_shortener_proto protoreflect.FileDescriptor

var file_shortener_shortener_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2f, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x27, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69,
	0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0x32, 0x0a,
	0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e,
	0x6b, 0x22, 0x24, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0x2d, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72,
	0x69, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x72,
	0x69, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x32, 0xb9, 0x01, 0x0a, 0x09, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x65, 0x6e, 0x65, 0x72, 0x12, 0x59, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1c,
	0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c,
	0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0c, 0x22, 0x07, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12,
	0x51, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x19, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e,
	0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x7b, 0x6c, 0x69, 0x6e,
	0x6b, 0x7d, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shortener_shortener_proto_rawDescOnce sync.Once
	file_shortener_shortener_proto_rawDescData = file_shortener_shortener_proto_rawDesc
)

func file_shortener_shortener_proto_rawDescGZIP() []byte {
	file_shortener_shortener_proto_rawDescOnce.Do(func() {
		file_shortener_shortener_proto_rawDescData = protoimpl.X.CompressGZIP(file_shortener_shortener_proto_rawDescData)
	})
	return file_shortener_shortener_proto_rawDescData
}

var file_shortener_shortener_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_shortener_shortener_proto_goTypes = []interface{}{
	(*CreateLinkRequest)(nil),  // 0: shortener.CreateLinkRequest
	(*CreateLinkResponse)(nil), // 1: shortener.CreateLinkResponse
	(*GetLinkRequest)(nil),     // 2: shortener.GetLinkRequest
	(*GetLinkResponse)(nil),    // 3: shortener.GetLinkResponse
}
var file_shortener_shortener_proto_depIdxs = []int32{
	0, // 0: shortener.Shortener.Create:input_type -> shortener.CreateLinkRequest
	2, // 1: shortener.Shortener.Get:input_type -> shortener.GetLinkRequest
	1, // 2: shortener.Shortener.Create:output_type -> shortener.CreateLinkResponse
	3, // 3: shortener.Shortener.Get:output_type -> shortener.GetLinkResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_shortener_shortener_proto_init() }
func file_shortener_shortener_proto_init() {
	if File_shortener_shortener_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shortener_shortener_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLinkRequest); i {
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
		file_shortener_shortener_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLinkResponse); i {
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
		file_shortener_shortener_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLinkRequest); i {
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
		file_shortener_shortener_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLinkResponse); i {
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
			RawDescriptor: file_shortener_shortener_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shortener_shortener_proto_goTypes,
		DependencyIndexes: file_shortener_shortener_proto_depIdxs,
		MessageInfos:      file_shortener_shortener_proto_msgTypes,
	}.Build()
	File_shortener_shortener_proto = out.File
	file_shortener_shortener_proto_rawDesc = nil
	file_shortener_shortener_proto_goTypes = nil
	file_shortener_shortener_proto_depIdxs = nil
}
