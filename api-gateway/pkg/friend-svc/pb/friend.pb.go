// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: pkg/friend-svc/pb/friend.proto

package pb

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

type FriendRequestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID   string `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	FriendID string `protobuf:"bytes,2,opt,name=FriendID,proto3" json:"FriendID,omitempty"`
}

func (x *FriendRequestRequest) Reset() {
	*x = FriendRequestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_friend_svc_pb_friend_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendRequestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendRequestRequest) ProtoMessage() {}

func (x *FriendRequestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_friend_svc_pb_friend_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendRequestRequest.ProtoReflect.Descriptor instead.
func (*FriendRequestRequest) Descriptor() ([]byte, []int) {
	return file_pkg_friend_svc_pb_friend_proto_rawDescGZIP(), []int{0}
}

func (x *FriendRequestRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *FriendRequestRequest) GetFriendID() string {
	if x != nil {
		return x.FriendID
	}
	return ""
}

type FriendRequestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FriendsID string `protobuf:"bytes,1,opt,name=FriendsID,proto3" json:"FriendsID,omitempty"`
	UserID    string `protobuf:"bytes,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Status    string `protobuf:"bytes,4,opt,name=Status,proto3" json:"Status,omitempty"`
	UpdateAt  string `protobuf:"bytes,5,opt,name=UpdateAt,proto3" json:"UpdateAt,omitempty"`
}

func (x *FriendRequestResponse) Reset() {
	*x = FriendRequestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_friend_svc_pb_friend_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendRequestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendRequestResponse) ProtoMessage() {}

func (x *FriendRequestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_friend_svc_pb_friend_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendRequestResponse.ProtoReflect.Descriptor instead.
func (*FriendRequestResponse) Descriptor() ([]byte, []int) {
	return file_pkg_friend_svc_pb_friend_proto_rawDescGZIP(), []int{1}
}

func (x *FriendRequestResponse) GetFriendsID() string {
	if x != nil {
		return x.FriendsID
	}
	return ""
}

func (x *FriendRequestResponse) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *FriendRequestResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *FriendRequestResponse) GetUpdateAt() string {
	if x != nil {
		return x.UpdateAt
	}
	return ""
}

type FriendListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID string `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	OffSet string `protobuf:"bytes,2,opt,name=OffSet,proto3" json:"OffSet,omitempty"`
	Limit  string `protobuf:"bytes,3,opt,name=Limit,proto3" json:"Limit,omitempty"`
	Status string `protobuf:"bytes,4,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *FriendListRequest) Reset() {
	*x = FriendListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_friend_svc_pb_friend_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendListRequest) ProtoMessage() {}

func (x *FriendListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_friend_svc_pb_friend_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendListRequest.ProtoReflect.Descriptor instead.
func (*FriendListRequest) Descriptor() ([]byte, []int) {
	return file_pkg_friend_svc_pb_friend_proto_rawDescGZIP(), []int{2}
}

func (x *FriendListRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *FriendListRequest) GetOffSet() string {
	if x != nil {
		return x.OffSet
	}
	return ""
}

func (x *FriendListRequest) GetLimit() string {
	if x != nil {
		return x.Limit
	}
	return ""
}

func (x *FriendListRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type FriendListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FriendID        string `protobuf:"bytes,1,opt,name=FriendID,proto3" json:"FriendID,omitempty"`
	UpdateAt        string `protobuf:"bytes,2,opt,name=UpdateAt,proto3" json:"UpdateAt,omitempty"`
	Messsage        string `protobuf:"bytes,3,opt,name=Messsage,proto3" json:"Messsage,omitempty"`
	UniqueFriendsID string `protobuf:"bytes,4,opt,name=UniqueFriendsID,proto3" json:"UniqueFriendsID,omitempty"`
}

func (x *FriendListResponse) Reset() {
	*x = FriendListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_friend_svc_pb_friend_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendListResponse) ProtoMessage() {}

func (x *FriendListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_friend_svc_pb_friend_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendListResponse.ProtoReflect.Descriptor instead.
func (*FriendListResponse) Descriptor() ([]byte, []int) {
	return file_pkg_friend_svc_pb_friend_proto_rawDescGZIP(), []int{3}
}

func (x *FriendListResponse) GetFriendID() string {
	if x != nil {
		return x.FriendID
	}
	return ""
}

func (x *FriendListResponse) GetUpdateAt() string {
	if x != nil {
		return x.UpdateAt
	}
	return ""
}

func (x *FriendListResponse) GetMesssage() string {
	if x != nil {
		return x.Messsage
	}
	return ""
}

func (x *FriendListResponse) GetUniqueFriendsID() string {
	if x != nil {
		return x.UniqueFriendsID
	}
	return ""
}

var File_pkg_friend_svc_pb_friend_proto protoreflect.FileDescriptor

var file_pkg_friend_svc_pb_friend_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x2d, 0x73, 0x76, 0x63,
	0x2f, 0x70, 0x62, 0x2f, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x4a, 0x0a, 0x14, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x1a, 0x0a, 0x08, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x44, 0x22, 0x81, 0x01, 0x0a,
	0x15, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x73, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x46, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x73, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74,
	0x22, 0x71, 0x0a, 0x11, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a,
	0x06, 0x4f, 0x66, 0x66, 0x53, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4f,
	0x66, 0x66, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x92, 0x01, 0x0a, 0x12, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x41, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x4d, 0x65, 0x73, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4d, 0x65, 0x73, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x28,
	0x0a, 0x0f, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x49,
	0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x46,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x49, 0x44, 0x32, 0x88, 0x01, 0x0a, 0x0d, 0x46, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x0d, 0x46, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x2e, 0x46, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x0a,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x2e, 0x46, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13,
	0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x16, 0x5a, 0x14, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x2d, 0x73, 0x76, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_pkg_friend_svc_pb_friend_proto_rawDescOnce sync.Once
	file_pkg_friend_svc_pb_friend_proto_rawDescData = file_pkg_friend_svc_pb_friend_proto_rawDesc
)

func file_pkg_friend_svc_pb_friend_proto_rawDescGZIP() []byte {
	file_pkg_friend_svc_pb_friend_proto_rawDescOnce.Do(func() {
		file_pkg_friend_svc_pb_friend_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_friend_svc_pb_friend_proto_rawDescData)
	})
	return file_pkg_friend_svc_pb_friend_proto_rawDescData
}

var file_pkg_friend_svc_pb_friend_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_friend_svc_pb_friend_proto_goTypes = []interface{}{
	(*FriendRequestRequest)(nil),  // 0: FriendRequestRequest
	(*FriendRequestResponse)(nil), // 1: FriendRequestResponse
	(*FriendListRequest)(nil),     // 2: FriendListRequest
	(*FriendListResponse)(nil),    // 3: FriendListResponse
}
var file_pkg_friend_svc_pb_friend_proto_depIdxs = []int32{
	0, // 0: FriendService.FriendRequest:input_type -> FriendRequestRequest
	2, // 1: FriendService.FriendList:input_type -> FriendListRequest
	1, // 2: FriendService.FriendRequest:output_type -> FriendRequestResponse
	3, // 3: FriendService.FriendList:output_type -> FriendListResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_friend_svc_pb_friend_proto_init() }
func file_pkg_friend_svc_pb_friend_proto_init() {
	if File_pkg_friend_svc_pb_friend_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_friend_svc_pb_friend_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendRequestRequest); i {
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
		file_pkg_friend_svc_pb_friend_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendRequestResponse); i {
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
		file_pkg_friend_svc_pb_friend_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendListRequest); i {
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
		file_pkg_friend_svc_pb_friend_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendListResponse); i {
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
			RawDescriptor: file_pkg_friend_svc_pb_friend_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_friend_svc_pb_friend_proto_goTypes,
		DependencyIndexes: file_pkg_friend_svc_pb_friend_proto_depIdxs,
		MessageInfos:      file_pkg_friend_svc_pb_friend_proto_msgTypes,
	}.Build()
	File_pkg_friend_svc_pb_friend_proto = out.File
	file_pkg_friend_svc_pb_friend_proto_rawDesc = nil
	file_pkg_friend_svc_pb_friend_proto_goTypes = nil
	file_pkg_friend_svc_pb_friend_proto_depIdxs = nil
}
