// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: backend.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PostUsedAcc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nip      uint64 `protobuf:"varint,1,opt,name=nip,proto3" json:"nip,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *PostUsedAcc) Reset() {
	*x = PostUsedAcc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostUsedAcc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostUsedAcc) ProtoMessage() {}

func (x *PostUsedAcc) ProtoReflect() protoreflect.Message {
	mi := &file_backend_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostUsedAcc.ProtoReflect.Descriptor instead.
func (*PostUsedAcc) Descriptor() ([]byte, []int) {
	return file_backend_proto_rawDescGZIP(), []int{0}
}

func (x *PostUsedAcc) GetNip() uint64 {
	if x != nil {
		return x.Nip
	}
	return 0
}

func (x *PostUsedAcc) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type GetNipResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nip uint64 `protobuf:"varint,1,opt,name=nip,proto3" json:"nip,omitempty"`
}

func (x *GetNipResponse) Reset() {
	*x = GetNipResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNipResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNipResponse) ProtoMessage() {}

func (x *GetNipResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backend_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNipResponse.ProtoReflect.Descriptor instead.
func (*GetNipResponse) Descriptor() ([]byte, []int) {
	return file_backend_proto_rawDescGZIP(), []int{1}
}

func (x *GetNipResponse) GetNip() uint64 {
	if x != nil {
		return x.Nip
	}
	return 0
}

var File_backend_proto protoreflect.FileDescriptor

var file_backend_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x3b, 0x0a, 0x0b, 0x50, 0x6f, 0x73, 0x74, 0x55, 0x73, 0x65, 0x64, 0x41, 0x63, 0x63, 0x12,
	0x10, 0x0a, 0x03, 0x6e, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6e, 0x69,
	0x70, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x22, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x4e, 0x69, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6e, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6e, 0x69,
	0x70, 0x32, 0xdf, 0x02, 0x0a, 0x0a, 0x4e, 0x49, 0x50, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x36, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x49, 0x74, 0x4e, 0x69, 0x70, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x69, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4e,
	0x75, 0x72, 0x73, 0x65, 0x4e, 0x69, 0x70, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x12, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x69, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0a, 0x50, 0x6f, 0x73, 0x74, 0x55, 0x73, 0x65, 0x64, 0x49,
	0x54, 0x12, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x55, 0x73, 0x65, 0x64, 0x41,
	0x63, 0x63, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x38, 0x0a, 0x0d, 0x50, 0x6f,
	0x73, 0x74, 0x55, 0x73, 0x65, 0x64, 0x4e, 0x75, 0x72, 0x73, 0x65, 0x12, 0x0f, 0x2e, 0x70, 0x62,
	0x2e, 0x50, 0x6f, 0x73, 0x74, 0x55, 0x73, 0x65, 0x64, 0x41, 0x63, 0x63, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x34, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x64, 0x49,
	0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x50,
	0x6f, 0x73, 0x74, 0x55, 0x73, 0x65, 0x64, 0x41, 0x63, 0x63, 0x12, 0x37, 0x0a, 0x0c, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x64, 0x4e, 0x75, 0x72, 0x73, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x55, 0x73, 0x65, 0x64,
	0x41, 0x63, 0x63, 0x42, 0x0b, 0x5a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_backend_proto_rawDescOnce sync.Once
	file_backend_proto_rawDescData = file_backend_proto_rawDesc
)

func file_backend_proto_rawDescGZIP() []byte {
	file_backend_proto_rawDescOnce.Do(func() {
		file_backend_proto_rawDescData = protoimpl.X.CompressGZIP(file_backend_proto_rawDescData)
	})
	return file_backend_proto_rawDescData
}

var file_backend_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_backend_proto_goTypes = []interface{}{
	(*PostUsedAcc)(nil),    // 0: pb.PostUsedAcc
	(*GetNipResponse)(nil), // 1: pb.GetNipResponse
	(*emptypb.Empty)(nil),  // 2: google.protobuf.Empty
}
var file_backend_proto_depIdxs = []int32{
	2, // 0: pb.NIPService.GetItNip:input_type -> google.protobuf.Empty
	2, // 1: pb.NIPService.GetNurseNip:input_type -> google.protobuf.Empty
	0, // 2: pb.NIPService.PostUsedIT:input_type -> pb.PostUsedAcc
	0, // 3: pb.NIPService.PostUsedNurse:input_type -> pb.PostUsedAcc
	2, // 4: pb.NIPService.GetUsedIt:input_type -> google.protobuf.Empty
	2, // 5: pb.NIPService.GetUsedNurse:input_type -> google.protobuf.Empty
	1, // 6: pb.NIPService.GetItNip:output_type -> pb.GetNipResponse
	1, // 7: pb.NIPService.GetNurseNip:output_type -> pb.GetNipResponse
	2, // 8: pb.NIPService.PostUsedIT:output_type -> google.protobuf.Empty
	2, // 9: pb.NIPService.PostUsedNurse:output_type -> google.protobuf.Empty
	0, // 10: pb.NIPService.GetUsedIt:output_type -> pb.PostUsedAcc
	0, // 11: pb.NIPService.GetUsedNurse:output_type -> pb.PostUsedAcc
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_backend_proto_init() }
func file_backend_proto_init() {
	if File_backend_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_backend_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostUsedAcc); i {
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
		file_backend_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNipResponse); i {
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
			RawDescriptor: file_backend_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_backend_proto_goTypes,
		DependencyIndexes: file_backend_proto_depIdxs,
		MessageInfos:      file_backend_proto_msgTypes,
	}.Build()
	File_backend_proto = out.File
	file_backend_proto_rawDesc = nil
	file_backend_proto_goTypes = nil
	file_backend_proto_depIdxs = nil
}
