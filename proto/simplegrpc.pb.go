// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: proto/simplegrpc.proto

package proto

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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceName      string  `protobuf:"bytes,1,opt,name=SourceName,proto3" json:"SourceName,omitempty"`
	DestinationName string  `protobuf:"bytes,2,opt,name=DestinationName,proto3" json:"DestinationName,omitempty"`
	Body            string  `protobuf:"bytes,3,opt,name=Body,proto3" json:"Body,omitempty"`
	Count           *uint32 `protobuf:"varint,4,opt,name=Count,proto3,oneof" json:"Count,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_simplegrpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_proto_simplegrpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_proto_simplegrpc_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetSourceName() string {
	if x != nil {
		return x.SourceName
	}
	return ""
}

func (x *Message) GetDestinationName() string {
	if x != nil {
		return x.DestinationName
	}
	return ""
}

func (x *Message) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *Message) GetCount() uint32 {
	if x != nil && x.Count != nil {
		return *x.Count
	}
	return 0
}

type MultiEchoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message        *Message `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	Repeats        *uint32  `protobuf:"varint,2,opt,name=Repeats,proto3,oneof" json:"Repeats,omitempty"`
	DelayInSeconds *uint32  `protobuf:"varint,3,opt,name=DelayInSeconds,proto3,oneof" json:"DelayInSeconds,omitempty"`
}

func (x *MultiEchoRequest) Reset() {
	*x = MultiEchoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_simplegrpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiEchoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiEchoRequest) ProtoMessage() {}

func (x *MultiEchoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_simplegrpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiEchoRequest.ProtoReflect.Descriptor instead.
func (*MultiEchoRequest) Descriptor() ([]byte, []int) {
	return file_proto_simplegrpc_proto_rawDescGZIP(), []int{1}
}

func (x *MultiEchoRequest) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *MultiEchoRequest) GetRepeats() uint32 {
	if x != nil && x.Repeats != nil {
		return *x.Repeats
	}
	return 0
}

func (x *MultiEchoRequest) GetDelayInSeconds() uint32 {
	if x != nil && x.DelayInSeconds != nil {
		return *x.DelayInSeconds
	}
	return 0
}

type Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"zigzag32,1,opt,name=Number,proto3" json:"Number,omitempty"`
}

func (x *Value) Reset() {
	*x = Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_simplegrpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value) ProtoMessage() {}

func (x *Value) ProtoReflect() protoreflect.Message {
	mi := &file_proto_simplegrpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Value.ProtoReflect.Descriptor instead.
func (*Value) Descriptor() ([]byte, []int) {
	return file_proto_simplegrpc_proto_rawDescGZIP(), []int{2}
}

func (x *Value) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

var File_proto_simplegrpc_proto protoreflect.FileDescriptor

var file_proto_simplegrpc_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x8c, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x44,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x19, 0x0a, 0x05, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x05, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xa7,
	0x01, 0x0a, 0x10, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a,
	0x07, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00,
	0x52, 0x07, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x73, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x0e,
	0x44, 0x65, 0x6c, 0x61, 0x79, 0x49, 0x6e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x48, 0x01, 0x52, 0x0e, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x49, 0x6e, 0x53,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x52, 0x65,
	0x70, 0x65, 0x61, 0x74, 0x73, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x49,
	0x6e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x22, 0x1f, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x11, 0x52, 0x06, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x32, 0x69, 0x0a, 0x04, 0x45, 0x63, 0x68,
	0x6f, 0x12, 0x29, 0x0a, 0x07, 0x4f, 0x6e, 0x65, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x0e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x36, 0x0a, 0x09,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x30, 0x01, 0x32, 0x63, 0x0a, 0x0a, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x6f, 0x72, 0x12, 0x23, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x28, 0x01, 0x12, 0x30, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x65, 0x69, 0x66, 0x75, 0x6b, 0x61, 0x6d,
	0x69, 0x2f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_simplegrpc_proto_rawDescOnce sync.Once
	file_proto_simplegrpc_proto_rawDescData = file_proto_simplegrpc_proto_rawDesc
)

func file_proto_simplegrpc_proto_rawDescGZIP() []byte {
	file_proto_simplegrpc_proto_rawDescOnce.Do(func() {
		file_proto_simplegrpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_simplegrpc_proto_rawDescData)
	})
	return file_proto_simplegrpc_proto_rawDescData
}

var file_proto_simplegrpc_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_simplegrpc_proto_goTypes = []interface{}{
	(*Message)(nil),          // 0: proto.Message
	(*MultiEchoRequest)(nil), // 1: proto.MultiEchoRequest
	(*Value)(nil),            // 2: proto.Value
}
var file_proto_simplegrpc_proto_depIdxs = []int32{
	0, // 0: proto.MultiEchoRequest.Message:type_name -> proto.Message
	0, // 1: proto.Echo.OneEcho:input_type -> proto.Message
	1, // 2: proto.Echo.MultiEcho:input_type -> proto.MultiEchoRequest
	2, // 3: proto.Calculator.Add:input_type -> proto.Value
	2, // 4: proto.Calculator.AddInteractive:input_type -> proto.Value
	0, // 5: proto.Echo.OneEcho:output_type -> proto.Message
	0, // 6: proto.Echo.MultiEcho:output_type -> proto.Message
	2, // 7: proto.Calculator.Add:output_type -> proto.Value
	2, // 8: proto.Calculator.AddInteractive:output_type -> proto.Value
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_simplegrpc_proto_init() }
func file_proto_simplegrpc_proto_init() {
	if File_proto_simplegrpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_simplegrpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_proto_simplegrpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiEchoRequest); i {
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
		file_proto_simplegrpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Value); i {
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
	file_proto_simplegrpc_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_proto_simplegrpc_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_simplegrpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_proto_simplegrpc_proto_goTypes,
		DependencyIndexes: file_proto_simplegrpc_proto_depIdxs,
		MessageInfos:      file_proto_simplegrpc_proto_msgTypes,
	}.Build()
	File_proto_simplegrpc_proto = out.File
	file_proto_simplegrpc_proto_rawDesc = nil
	file_proto_simplegrpc_proto_goTypes = nil
	file_proto_simplegrpc_proto_depIdxs = nil
}
