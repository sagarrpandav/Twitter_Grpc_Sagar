// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: authService.proto

package auth

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

// The request message containing the user's information.
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	FirstName string `protobuf:"bytes,2,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName  string `protobuf:"bytes,3,opt,name=LastName,proto3" json:"LastName,omitempty"`
	Email     string `protobuf:"bytes,4,opt,name=Email,proto3" json:"Email,omitempty"`
	Password  string `protobuf:"bytes,5,opt,name=Password,proto3" json:"Password,omitempty"`
	UserHash  string `protobuf:"bytes,6,opt,name=UserHash,proto3" json:"UserHash,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_authService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_authService_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetUserHash() string {
	if x != nil {
		return x.UserHash
	}
	return ""
}

type LoginUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *LoginUser) Reset() {
	*x = LoginUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginUser) ProtoMessage() {}

func (x *LoginUser) ProtoReflect() protoreflect.Message {
	mi := &file_authService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginUser.ProtoReflect.Descriptor instead.
func (*LoginUser) Descriptor() ([]byte, []int) {
	return file_authService_proto_rawDescGZIP(), []int{1}
}

func (x *LoginUser) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginUser) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserHash string `protobuf:"bytes,1,opt,name=UserHash,proto3" json:"UserHash,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_authService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_authService_proto_rawDescGZIP(), []int{2}
}

func (x *Token) GetUserHash() string {
	if x != nil {
		return x.UserHash
	}
	return ""
}

// The response message containing the greetings
type ResponseMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ResponseMessage) Reset() {
	*x = ResponseMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseMessage) ProtoMessage() {}

func (x *ResponseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_authService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseMessage.ProtoReflect.Descriptor instead.
func (*ResponseMessage) Descriptor() ([]byte, []int) {
	return file_authService_proto_rawDescGZIP(), []int{3}
}

func (x *ResponseMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type EmptyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyRequest) Reset() {
	*x = EmptyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRequest) ProtoMessage() {}

func (x *EmptyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRequest.ProtoReflect.Descriptor instead.
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return file_authService_proto_rawDescGZIP(), []int{4}
}

var File_authService_proto protoreflect.FileDescriptor

var file_authService_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x61,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x48, 0x61, 0x73, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x48, 0x61, 0x73, 0x68, 0x22, 0x3d, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x22, 0x23, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x48, 0x61, 0x73, 0x68, 0x22, 0x2b, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x0e, 0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0x9a, 0x01, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x18,
	0x0a, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x12, 0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a,
	0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x12, 0x1d, 0x0a, 0x06, 0x73, 0x69, 0x67, 0x6e,
	0x49, 0x6e, 0x12, 0x0a, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x05,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x12, 0x25, 0x0a, 0x07, 0x73, 0x69, 0x67, 0x6e, 0x4f,
	0x75, 0x74, 0x12, 0x06, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x10, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x32,
	0x0a, 0x14, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f,
	0x67, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x12, 0x06, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x10,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_authService_proto_rawDescOnce sync.Once
	file_authService_proto_rawDescData = file_authService_proto_rawDesc
)

func file_authService_proto_rawDescGZIP() []byte {
	file_authService_proto_rawDescOnce.Do(func() {
		file_authService_proto_rawDescData = protoimpl.X.CompressGZIP(file_authService_proto_rawDescData)
	})
	return file_authService_proto_rawDescData
}

var file_authService_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_authService_proto_goTypes = []interface{}{
	(*User)(nil),            // 0: User
	(*LoginUser)(nil),       // 1: LoginUser
	(*Token)(nil),           // 2: token
	(*ResponseMessage)(nil), // 3: ResponseMessage
	(*EmptyRequest)(nil),    // 4: EmptyRequest
}
var file_authService_proto_depIdxs = []int32{
	0, // 0: Auth.signUp:input_type -> User
	1, // 1: Auth.signIn:input_type -> LoginUser
	2, // 2: Auth.signOut:input_type -> token
	2, // 3: Auth.validateUserLoggedIn:input_type -> token
	0, // 4: Auth.signUp:output_type -> User
	0, // 5: Auth.signIn:output_type -> User
	3, // 6: Auth.signOut:output_type -> ResponseMessage
	3, // 7: Auth.validateUserLoggedIn:output_type -> ResponseMessage
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_authService_proto_init() }
func file_authService_proto_init() {
	if File_authService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_authService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_authService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginUser); i {
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
		file_authService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
		file_authService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseMessage); i {
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
		file_authService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyRequest); i {
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
			RawDescriptor: file_authService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_authService_proto_goTypes,
		DependencyIndexes: file_authService_proto_depIdxs,
		MessageInfos:      file_authService_proto_msgTypes,
	}.Build()
	File_authService_proto = out.File
	file_authService_proto_rawDesc = nil
	file_authService_proto_goTypes = nil
	file_authService_proto_depIdxs = nil
}