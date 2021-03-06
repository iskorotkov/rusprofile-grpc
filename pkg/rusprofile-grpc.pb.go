// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.2
// source: rusprofile-grpc.proto

package pkg

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

type INN struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	INN string `protobuf:"bytes,1,opt,name=INN,proto3" json:"INN,omitempty"`
}

func (x *INN) Reset() {
	*x = INN{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rusprofile_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *INN) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*INN) ProtoMessage() {}

func (x *INN) ProtoReflect() protoreflect.Message {
	mi := &file_rusprofile_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use INN.ProtoReflect.Descriptor instead.
func (*INN) Descriptor() ([]byte, []int) {
	return file_rusprofile_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *INN) GetINN() string {
	if x != nil {
		return x.INN
	}
	return ""
}

type Company struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	INN  string `protobuf:"bytes,1,opt,name=INN,proto3" json:"INN,omitempty"`
	KPP  string `protobuf:"bytes,2,opt,name=KPP,proto3" json:"KPP,omitempty"`
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	CEO  string `protobuf:"bytes,4,opt,name=CEO,proto3" json:"CEO,omitempty"`
}

func (x *Company) Reset() {
	*x = Company{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rusprofile_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Company) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Company) ProtoMessage() {}

func (x *Company) ProtoReflect() protoreflect.Message {
	mi := &file_rusprofile_grpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Company.ProtoReflect.Descriptor instead.
func (*Company) Descriptor() ([]byte, []int) {
	return file_rusprofile_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *Company) GetINN() string {
	if x != nil {
		return x.INN
	}
	return ""
}

func (x *Company) GetKPP() string {
	if x != nil {
		return x.KPP
	}
	return ""
}

func (x *Company) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Company) GetCEO() string {
	if x != nil {
		return x.CEO
	}
	return ""
}

var File_rusprofile_grpc_proto protoreflect.FileDescriptor

var file_rusprofile_grpc_proto_rawDesc = []byte{
	0x0a, 0x15, 0x72, 0x75, 0x73, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2d, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x72, 0x75, 0x73, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x17, 0x0a, 0x03, 0x49, 0x4e, 0x4e, 0x12, 0x10, 0x0a, 0x03, 0x49,
	0x4e, 0x4e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x49, 0x4e, 0x4e, 0x22, 0x53, 0x0a,
	0x07, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x49, 0x4e, 0x4e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x49, 0x4e, 0x4e, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x50,
	0x50, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x50, 0x50, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x43, 0x45, 0x4f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x43,
	0x45, 0x4f, 0x32, 0x5f, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x46, 0x69, 0x6e,
	0x64, 0x65, 0x72, 0x12, 0x4e, 0x0a, 0x05, 0x42, 0x79, 0x49, 0x4e, 0x4e, 0x12, 0x12, 0x2e, 0x72,
	0x75, 0x73, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x4e, 0x4e,
	0x1a, 0x16, 0x2e, 0x72, 0x75, 0x73, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13,
	0x12, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2f, 0x7b, 0x49,
	0x4e, 0x4e, 0x7d, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x69, 0x73, 0x6b, 0x6f, 0x72, 0x6f, 0x74, 0x6b, 0x6f, 0x76, 0x2f, 0x72, 0x75, 0x73,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x6b, 0x67,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rusprofile_grpc_proto_rawDescOnce sync.Once
	file_rusprofile_grpc_proto_rawDescData = file_rusprofile_grpc_proto_rawDesc
)

func file_rusprofile_grpc_proto_rawDescGZIP() []byte {
	file_rusprofile_grpc_proto_rawDescOnce.Do(func() {
		file_rusprofile_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_rusprofile_grpc_proto_rawDescData)
	})
	return file_rusprofile_grpc_proto_rawDescData
}

var file_rusprofile_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rusprofile_grpc_proto_goTypes = []interface{}{
	(*INN)(nil),     // 0: rusprofile.v1.INN
	(*Company)(nil), // 1: rusprofile.v1.Company
}
var file_rusprofile_grpc_proto_depIdxs = []int32{
	0, // 0: rusprofile.v1.CompanyFinder.ByINN:input_type -> rusprofile.v1.INN
	1, // 1: rusprofile.v1.CompanyFinder.ByINN:output_type -> rusprofile.v1.Company
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rusprofile_grpc_proto_init() }
func file_rusprofile_grpc_proto_init() {
	if File_rusprofile_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rusprofile_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*INN); i {
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
		file_rusprofile_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Company); i {
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
			RawDescriptor: file_rusprofile_grpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rusprofile_grpc_proto_goTypes,
		DependencyIndexes: file_rusprofile_grpc_proto_depIdxs,
		MessageInfos:      file_rusprofile_grpc_proto_msgTypes,
	}.Build()
	File_rusprofile_grpc_proto = out.File
	file_rusprofile_grpc_proto_rawDesc = nil
	file_rusprofile_grpc_proto_goTypes = nil
	file_rusprofile_grpc_proto_depIdxs = nil
}
