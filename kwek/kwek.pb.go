// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: kwek.proto

package kwek

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateKwek struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text     string                 `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	UserId   string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PostedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=posted_at,json=postedAt,proto3" json:"posted_at,omitempty"`
	KwekGuid string                 `protobuf:"bytes,4,opt,name=kwek_guid,json=kwekGuid,proto3" json:"kwek_guid,omitempty"`
}

func (x *CreateKwek) Reset() {
	*x = CreateKwek{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kwek_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateKwek) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateKwek) ProtoMessage() {}

func (x *CreateKwek) ProtoReflect() protoreflect.Message {
	mi := &file_kwek_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateKwek.ProtoReflect.Descriptor instead.
func (*CreateKwek) Descriptor() ([]byte, []int) {
	return file_kwek_proto_rawDescGZIP(), []int{0}
}

func (x *CreateKwek) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *CreateKwek) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateKwek) GetPostedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.PostedAt
	}
	return nil
}

func (x *CreateKwek) GetKwekGuid() string {
	if x != nil {
		return x.KwekGuid
	}
	return ""
}

type UpdateKwek struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text      string                 `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	KwekGuid  string                 `protobuf:"bytes,4,opt,name=kwek_guid,json=kwekGuid,proto3" json:"kwek_guid,omitempty"`
}

func (x *UpdateKwek) Reset() {
	*x = UpdateKwek{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kwek_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateKwek) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateKwek) ProtoMessage() {}

func (x *UpdateKwek) ProtoReflect() protoreflect.Message {
	mi := &file_kwek_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateKwek.ProtoReflect.Descriptor instead.
func (*UpdateKwek) Descriptor() ([]byte, []int) {
	return file_kwek_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateKwek) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *UpdateKwek) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *UpdateKwek) GetKwekGuid() string {
	if x != nil {
		return x.KwekGuid
	}
	return ""
}

type DeleteKwek struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KwekGuid string `protobuf:"bytes,2,opt,name=kwek_guid,json=kwekGuid,proto3" json:"kwek_guid,omitempty"`
}

func (x *DeleteKwek) Reset() {
	*x = DeleteKwek{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kwek_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteKwek) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteKwek) ProtoMessage() {}

func (x *DeleteKwek) ProtoReflect() protoreflect.Message {
	mi := &file_kwek_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteKwek.ProtoReflect.Descriptor instead.
func (*DeleteKwek) Descriptor() ([]byte, []int) {
	return file_kwek_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteKwek) GetKwekGuid() string {
	if x != nil {
		return x.KwekGuid
	}
	return ""
}

var File_kwek_proto protoreflect.FileDescriptor

var file_kwek_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6b, 0x77, 0x65, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6b, 0x77,
	0x65, 0x6b, 0x6b, 0x65, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x01, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4b, 0x77, 0x65, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x37, 0x0a, 0x09, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6b,
	0x77, 0x65, 0x6b, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6b, 0x77, 0x65, 0x6b, 0x47, 0x75, 0x69, 0x64, 0x22, 0x7e, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4b, 0x77, 0x65, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6b, 0x77, 0x65, 0x6b, 0x5f, 0x67, 0x75,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x77, 0x65, 0x6b, 0x47, 0x75,
	0x69, 0x64, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x22, 0x2f, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4b, 0x77, 0x65, 0x6b, 0x12, 0x1b, 0x0a, 0x09, 0x6b, 0x77, 0x65, 0x6b, 0x5f, 0x67,
	0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x77, 0x65, 0x6b, 0x47,
	0x75, 0x69, 0x64, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x42, 0x58, 0x5a, 0x32, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6f, 0x6c, 0x70, 0x6c,
	0x65, 0x78, 0x2d, 0x73, 0x36, 0x2f, 0x6b, 0x77, 0x65, 0x6b, 0x6b, 0x65, 0x72, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x6b, 0x77, 0x65, 0x6b, 0xaa,
	0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6f, 0x6c, 0x70, 0x6c, 0x65, 0x78, 0x2e, 0x4b, 0x77, 0x65,
	0x6b, 0x6b, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x2e, 0x4b,
	0x77, 0x65, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kwek_proto_rawDescOnce sync.Once
	file_kwek_proto_rawDescData = file_kwek_proto_rawDesc
)

func file_kwek_proto_rawDescGZIP() []byte {
	file_kwek_proto_rawDescOnce.Do(func() {
		file_kwek_proto_rawDescData = protoimpl.X.CompressGZIP(file_kwek_proto_rawDescData)
	})
	return file_kwek_proto_rawDescData
}

var file_kwek_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_kwek_proto_goTypes = []interface{}{
	(*CreateKwek)(nil),            // 0: kwekker.CreateKwek
	(*UpdateKwek)(nil),            // 1: kwekker.UpdateKwek
	(*DeleteKwek)(nil),            // 2: kwekker.DeleteKwek
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_kwek_proto_depIdxs = []int32{
	3, // 0: kwekker.CreateKwek.posted_at:type_name -> google.protobuf.Timestamp
	3, // 1: kwekker.UpdateKwek.updated_at:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_kwek_proto_init() }
func file_kwek_proto_init() {
	if File_kwek_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kwek_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateKwek); i {
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
		file_kwek_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateKwek); i {
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
		file_kwek_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteKwek); i {
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
			RawDescriptor: file_kwek_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kwek_proto_goTypes,
		DependencyIndexes: file_kwek_proto_depIdxs,
		MessageInfos:      file_kwek_proto_msgTypes,
	}.Build()
	File_kwek_proto = out.File
	file_kwek_proto_rawDesc = nil
	file_kwek_proto_goTypes = nil
	file_kwek_proto_depIdxs = nil
}
