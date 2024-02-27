// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.3
// source: crashlytics.proto

package proto

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

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exception  string `protobuf:"bytes,1,opt,name=exception,proto3" json:"exception,omitempty"`
	StackTrace string `protobuf:"bytes,2,opt,name=stack_trace,json=stackTrace,proto3" json:"stack_trace,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crashlytics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_crashlytics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_crashlytics_proto_rawDescGZIP(), []int{0}
}

func (x *Error) GetException() string {
	if x != nil {
		return x.Exception
	}
	return ""
}

func (x *Error) GetStackTrace() string {
	if x != nil {
		return x.StackTrace
	}
	return ""
}

var File_crashlytics_proto protoreflect.FileDescriptor

var file_crashlytics_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x72, 0x61, 0x73, 0x68, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x22, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x62, 0x72,
	0x61, 0x69, 0x6e, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x2e, 0x63, 0x72, 0x61, 0x73, 0x68, 0x6c, 0x79,
	0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1c, 0x0a,
	0x09, 0x65, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x65, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x63, 0x65, 0x32, 0x66, 0x0a, 0x12,
	0x43, 0x72, 0x61, 0x73, 0x68, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x50, 0x0a, 0x0b, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x29, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x62, 0x72, 0x61,
	0x69, 0x6e, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x2e, 0x63, 0x72, 0x61, 0x73, 0x68, 0x6c, 0x79, 0x74,
	0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_crashlytics_proto_rawDescOnce sync.Once
	file_crashlytics_proto_rawDescData = file_crashlytics_proto_rawDesc
)

func file_crashlytics_proto_rawDescGZIP() []byte {
	file_crashlytics_proto_rawDescOnce.Do(func() {
		file_crashlytics_proto_rawDescData = protoimpl.X.CompressGZIP(file_crashlytics_proto_rawDescData)
	})
	return file_crashlytics_proto_rawDescData
}

var file_crashlytics_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_crashlytics_proto_goTypes = []interface{}{
	(*Error)(nil),         // 0: endpoint.brainboost.crashlytics.v1.Error
	(*emptypb.Empty)(nil), // 1: google.protobuf.Empty
}
var file_crashlytics_proto_depIdxs = []int32{
	0, // 0: endpoint.brainboost.crashlytics.v1.CrashlyticsService.RecordError:input_type -> endpoint.brainboost.crashlytics.v1.Error
	1, // 1: endpoint.brainboost.crashlytics.v1.CrashlyticsService.RecordError:output_type -> google.protobuf.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_crashlytics_proto_init() }
func file_crashlytics_proto_init() {
	if File_crashlytics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_crashlytics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
			RawDescriptor: file_crashlytics_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_crashlytics_proto_goTypes,
		DependencyIndexes: file_crashlytics_proto_depIdxs,
		MessageInfos:      file_crashlytics_proto_msgTypes,
	}.Build()
	File_crashlytics_proto = out.File
	file_crashlytics_proto_rawDesc = nil
	file_crashlytics_proto_goTypes = nil
	file_crashlytics_proto_depIdxs = nil
}