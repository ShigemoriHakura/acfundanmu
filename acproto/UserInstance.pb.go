// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: UserInstance.proto

package acproto

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

type UserInstance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User       *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	InstanceId int64 `protobuf:"varint,2,opt,name=instanceId,proto3" json:"instanceId,omitempty"`
}

func (x *UserInstance) Reset() {
	*x = UserInstance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UserInstance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInstance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInstance) ProtoMessage() {}

func (x *UserInstance) ProtoReflect() protoreflect.Message {
	mi := &file_UserInstance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInstance.ProtoReflect.Descriptor instead.
func (*UserInstance) Descriptor() ([]byte, []int) {
	return file_UserInstance_proto_rawDescGZIP(), []int{0}
}

func (x *UserInstance) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UserInstance) GetInstanceId() int64 {
	if x != nil {
		return x.InstanceId
	}
	return 0
}

var File_UserInstance_proto protoreflect.FileDescriptor

var file_UserInstance_proto_rawDesc = []byte{
	0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75,
	0x1a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x0c,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x41, 0x63, 0x46,
	0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x49, 0x64, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x3b, 0x61, 0x63, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_UserInstance_proto_rawDescOnce sync.Once
	file_UserInstance_proto_rawDescData = file_UserInstance_proto_rawDesc
)

func file_UserInstance_proto_rawDescGZIP() []byte {
	file_UserInstance_proto_rawDescOnce.Do(func() {
		file_UserInstance_proto_rawDescData = protoimpl.X.CompressGZIP(file_UserInstance_proto_rawDescData)
	})
	return file_UserInstance_proto_rawDescData
}

var file_UserInstance_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_UserInstance_proto_goTypes = []interface{}{
	(*UserInstance)(nil), // 0: AcFunDanmu.UserInstance
	(*User)(nil),         // 1: AcFunDanmu.User
}
var file_UserInstance_proto_depIdxs = []int32{
	1, // 0: AcFunDanmu.UserInstance.user:type_name -> AcFunDanmu.User
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_UserInstance_proto_init() }
func file_UserInstance_proto_init() {
	if File_UserInstance_proto != nil {
		return
	}
	file_User_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_UserInstance_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInstance); i {
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
			RawDescriptor: file_UserInstance_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_UserInstance_proto_goTypes,
		DependencyIndexes: file_UserInstance_proto_depIdxs,
		MessageInfos:      file_UserInstance_proto_msgTypes,
	}.Build()
	File_UserInstance_proto = out.File
	file_UserInstance_proto_rawDesc = nil
	file_UserInstance_proto_goTypes = nil
	file_UserInstance_proto_depIdxs = nil
}
