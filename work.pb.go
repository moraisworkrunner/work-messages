// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: work.proto

package work_messages

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Context struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Context) Reset() {
	*x = Context{}
	if protoimpl.UnsafeEnabled {
		mi := &file_work_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Context) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Context) ProtoMessage() {}

func (x *Context) ProtoReflect() protoreflect.Message {
	mi := &file_work_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Context.ProtoReflect.Descriptor instead.
func (*Context) Descriptor() ([]byte, []int) {
	return file_work_proto_rawDescGZIP(), []int{0}
}

func (x *Context) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SvcWorkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceFile   string        `protobuf:"bytes,1,opt,name=sourceFile,proto3" json:"sourceFile,omitempty"`
	WebhookUrl   string        `protobuf:"bytes,2,opt,name=webhookUrl,proto3" json:"webhookUrl,omitempty"`
	FileMetadata *FileMetadata `protobuf:"bytes,3,opt,name=fileMetadata,proto3" json:"fileMetadata,omitempty"`
	Context      *Context      `protobuf:"bytes,4,opt,name=context,proto3" json:"context,omitempty"`
}

func (x *SvcWorkRequest) Reset() {
	*x = SvcWorkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_work_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SvcWorkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SvcWorkRequest) ProtoMessage() {}

func (x *SvcWorkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_work_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SvcWorkRequest.ProtoReflect.Descriptor instead.
func (*SvcWorkRequest) Descriptor() ([]byte, []int) {
	return file_work_proto_rawDescGZIP(), []int{1}
}

func (x *SvcWorkRequest) GetSourceFile() string {
	if x != nil {
		return x.SourceFile
	}
	return ""
}

func (x *SvcWorkRequest) GetWebhookUrl() string {
	if x != nil {
		return x.WebhookUrl
	}
	return ""
}

func (x *SvcWorkRequest) GetFileMetadata() *FileMetadata {
	if x != nil {
		return x.FileMetadata
	}
	return nil
}

func (x *SvcWorkRequest) GetContext() *Context {
	if x != nil {
		return x.Context
	}
	return nil
}

type SvcWorkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Context *Context `protobuf:"bytes,2,opt,name=context,proto3" json:"context,omitempty"`
	Error   *Error   `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *SvcWorkResponse) Reset() {
	*x = SvcWorkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_work_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SvcWorkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SvcWorkResponse) ProtoMessage() {}

func (x *SvcWorkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_work_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SvcWorkResponse.ProtoReflect.Descriptor instead.
func (*SvcWorkResponse) Descriptor() ([]byte, []int) {
	return file_work_proto_rawDescGZIP(), []int{2}
}

func (x *SvcWorkResponse) GetContext() *Context {
	if x != nil {
		return x.Context
	}
	return nil
}

func (x *SvcWorkResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type FileMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mimetype string `protobuf:"bytes,1,opt,name=mimetype,proto3" json:"mimetype,omitempty"`
	Md5      string `protobuf:"bytes,2,opt,name=md5,proto3" json:"md5,omitempty"`
	Size     int64  `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *FileMetadata) Reset() {
	*x = FileMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_work_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileMetadata) ProtoMessage() {}

func (x *FileMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_work_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileMetadata.ProtoReflect.Descriptor instead.
func (*FileMetadata) Descriptor() ([]byte, []int) {
	return file_work_proto_rawDescGZIP(), []int{3}
}

func (x *FileMetadata) GetMimetype() string {
	if x != nil {
		return x.Mimetype
	}
	return ""
}

func (x *FileMetadata) GetMd5() string {
	if x != nil {
		return x.Md5
	}
	return ""
}

func (x *FileMetadata) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_work_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_work_proto_msgTypes[4]
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
	return file_work_proto_rawDescGZIP(), []int{4}
}

func (x *Error) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_work_proto protoreflect.FileDescriptor

var file_work_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6d, 0x6f,
	0x72, 0x61, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0x19, 0x0a, 0x07, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xc3, 0x01, 0x0a, 0x0e, 0x53, 0x76, 0x63, 0x57, 0x6f,
	0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x77, 0x65, 0x62,
	0x68, 0x6f, 0x6f, 0x6b, 0x55, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77,
	0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x55, 0x72, 0x6c, 0x12, 0x3f, 0x0a, 0x0c, 0x66, 0x69, 0x6c,
	0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x6d, 0x6f, 0x72, 0x61, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x0c, 0x66, 0x69,
	0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x30, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x6f,
	0x72, 0x61, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x22, 0x6f, 0x0a, 0x0f,
	0x53, 0x76, 0x63, 0x57, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x30, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x6d, 0x6f, 0x72, 0x61, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x12, 0x2a, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x6d, 0x6f, 0x72, 0x61, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x50, 0x0a,
	0x0c, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a,
	0x08, 0x6d, 0x69, 0x6d, 0x65, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6d, 0x69, 0x6d, 0x65, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x64, 0x35,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x64, 0x35, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22,
	0x35, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x72, 0x61, 0x69, 0x73, 0x77, 0x6f, 0x72, 0x6b, 0x72,
	0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x2d, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x3b, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_work_proto_rawDescOnce sync.Once
	file_work_proto_rawDescData = file_work_proto_rawDesc
)

func file_work_proto_rawDescGZIP() []byte {
	file_work_proto_rawDescOnce.Do(func() {
		file_work_proto_rawDescData = protoimpl.X.CompressGZIP(file_work_proto_rawDescData)
	})
	return file_work_proto_rawDescData
}

var file_work_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_work_proto_goTypes = []interface{}{
	(*Context)(nil),         // 0: morais_protos.Context
	(*SvcWorkRequest)(nil),  // 1: morais_protos.SvcWorkRequest
	(*SvcWorkResponse)(nil), // 2: morais_protos.SvcWorkResponse
	(*FileMetadata)(nil),    // 3: morais_protos.FileMetadata
	(*Error)(nil),           // 4: morais_protos.Error
}
var file_work_proto_depIdxs = []int32{
	3, // 0: morais_protos.SvcWorkRequest.fileMetadata:type_name -> morais_protos.FileMetadata
	0, // 1: morais_protos.SvcWorkRequest.context:type_name -> morais_protos.Context
	0, // 2: morais_protos.SvcWorkResponse.context:type_name -> morais_protos.Context
	4, // 3: morais_protos.SvcWorkResponse.error:type_name -> morais_protos.Error
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_work_proto_init() }
func file_work_proto_init() {
	if File_work_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_work_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Context); i {
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
		file_work_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SvcWorkRequest); i {
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
		file_work_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SvcWorkResponse); i {
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
		file_work_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileMetadata); i {
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
		file_work_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_work_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_work_proto_goTypes,
		DependencyIndexes: file_work_proto_depIdxs,
		MessageInfos:      file_work_proto_msgTypes,
	}.Build()
	File_work_proto = out.File
	file_work_proto_rawDesc = nil
	file_work_proto_goTypes = nil
	file_work_proto_depIdxs = nil
}
