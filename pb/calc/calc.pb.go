// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: calc.proto

package calc

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

type SumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A int32 `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B int32 `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *SumRequest) Reset() {
	*x = SumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumRequest) ProtoMessage() {}

func (x *SumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumRequest.ProtoReflect.Descriptor instead.
func (*SumRequest) Descriptor() ([]byte, []int) {
	return file_calc_proto_rawDescGZIP(), []int{0}
}

func (x *SumRequest) GetA() int32 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *SumRequest) GetB() int32 {
	if x != nil {
		return x.B
	}
	return 0
}

type SumReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SumReply) Reset() {
	*x = SumReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumReply) ProtoMessage() {}

func (x *SumReply) ProtoReflect() protoreflect.Message {
	mi := &file_calc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumReply.ProtoReflect.Descriptor instead.
func (*SumReply) Descriptor() ([]byte, []int) {
	return file_calc_proto_rawDescGZIP(), []int{1}
}

func (x *SumReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_calc_proto protoreflect.FileDescriptor

var file_calc_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a, 0x0a,
	0x53, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x01, 0x62, 0x22, 0x24, 0x0a, 0x08, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x27, 0x0a, 0x04,
	0x43, 0x61, 0x6c, 0x63, 0x12, 0x1f, 0x0a, 0x03, 0x53, 0x75, 0x6d, 0x12, 0x0b, 0x2e, 0x53, 0x75,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x53, 0x75, 0x6d, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x5e, 0x0a, 0x17, 0x69, 0x6f, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65,
	0x42, 0x0b, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x2d, 0x6d,
	0x69, 0x79, 0x61, 0x7a, 0x61, 0x6b, 0x69, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d,
	0x6e, 0x65, 0x6f, 0x34, 0x6a, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x70, 0x62,
	0x2f, 0x63, 0x61, 0x6c, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_calc_proto_rawDescOnce sync.Once
	file_calc_proto_rawDescData = file_calc_proto_rawDesc
)

func file_calc_proto_rawDescGZIP() []byte {
	file_calc_proto_rawDescOnce.Do(func() {
		file_calc_proto_rawDescData = protoimpl.X.CompressGZIP(file_calc_proto_rawDescData)
	})
	return file_calc_proto_rawDescData
}

var file_calc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_calc_proto_goTypes = []interface{}{
	(*SumRequest)(nil), // 0: SumRequest
	(*SumReply)(nil),   // 1: SumReply
}
var file_calc_proto_depIdxs = []int32{
	0, // 0: Calc.Sum:input_type -> SumRequest
	1, // 1: Calc.Sum:output_type -> SumReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_calc_proto_init() }
func file_calc_proto_init() {
	if File_calc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_calc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumRequest); i {
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
		file_calc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumReply); i {
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
			RawDescriptor: file_calc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calc_proto_goTypes,
		DependencyIndexes: file_calc_proto_depIdxs,
		MessageInfos:      file_calc_proto_msgTypes,
	}.Build()
	File_calc_proto = out.File
	file_calc_proto_rawDesc = nil
	file_calc_proto_goTypes = nil
	file_calc_proto_depIdxs = nil
}
