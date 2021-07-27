// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.6.1
// source: v1/pb/imagehub.proto

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

type CheckStatus int32

const (
	CheckStatus_UpToDate    CheckStatus = 0
	CheckStatus_UpdateFound CheckStatus = 1
)

// Enum value maps for CheckStatus.
var (
	CheckStatus_name = map[int32]string{
		0: "UpToDate",
		1: "UpdateFound",
	}
	CheckStatus_value = map[string]int32{
		"UpToDate":    0,
		"UpdateFound": 1,
	}
)

func (x CheckStatus) Enum() *CheckStatus {
	p := new(CheckStatus)
	*p = x
	return p
}

func (x CheckStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CheckStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_pb_imagehub_proto_enumTypes[0].Descriptor()
}

func (CheckStatus) Type() protoreflect.EnumType {
	return &file_v1_pb_imagehub_proto_enumTypes[0]
}

func (x CheckStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CheckStatus.Descriptor instead.
func (CheckStatus) EnumDescriptor() ([]byte, []int) {
	return file_v1_pb_imagehub_proto_rawDescGZIP(), []int{0}
}

type CloneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReposPath string `protobuf:"bytes,1,opt,name=repos_path,json=reposPath,proto3" json:"repos_path,omitempty"`
}

func (x *CloneRequest) Reset() {
	*x = CloneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_pb_imagehub_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloneRequest) ProtoMessage() {}

func (x *CloneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_pb_imagehub_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloneRequest.ProtoReflect.Descriptor instead.
func (*CloneRequest) Descriptor() ([]byte, []int) {
	return file_v1_pb_imagehub_proto_rawDescGZIP(), []int{0}
}

func (x *CloneRequest) GetReposPath() string {
	if x != nil {
		return x.ReposPath
	}
	return ""
}

type MetaData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash       uint32 `protobuf:"varint,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Owner      string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	FolderName string `protobuf:"bytes,3,opt,name=folder_name,json=folderName,proto3" json:"folder_name,omitempty"`
}

func (x *MetaData) Reset() {
	*x = MetaData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_pb_imagehub_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetaData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetaData) ProtoMessage() {}

func (x *MetaData) ProtoReflect() protoreflect.Message {
	mi := &file_v1_pb_imagehub_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetaData.ProtoReflect.Descriptor instead.
func (*MetaData) Descriptor() ([]byte, []int) {
	return file_v1_pb_imagehub_proto_rawDescGZIP(), []int{1}
}

func (x *MetaData) GetHash() uint32 {
	if x != nil {
		return x.Hash
	}
	return 0
}

func (x *MetaData) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *MetaData) GetFolderName() string {
	if x != nil {
		return x.FolderName
	}
	return ""
}

type CloneResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//	*CloneResponse_Metadata
	//	*CloneResponse_ChunkData
	Data isCloneResponse_Data `protobuf_oneof:"data"`
}

func (x *CloneResponse) Reset() {
	*x = CloneResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_pb_imagehub_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloneResponse) ProtoMessage() {}

func (x *CloneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_pb_imagehub_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloneResponse.ProtoReflect.Descriptor instead.
func (*CloneResponse) Descriptor() ([]byte, []int) {
	return file_v1_pb_imagehub_proto_rawDescGZIP(), []int{2}
}

func (m *CloneResponse) GetData() isCloneResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *CloneResponse) GetMetadata() *MetaData {
	if x, ok := x.GetData().(*CloneResponse_Metadata); ok {
		return x.Metadata
	}
	return nil
}

func (x *CloneResponse) GetChunkData() []byte {
	if x, ok := x.GetData().(*CloneResponse_ChunkData); ok {
		return x.ChunkData
	}
	return nil
}

type isCloneResponse_Data interface {
	isCloneResponse_Data()
}

type CloneResponse_Metadata struct {
	Metadata *MetaData `protobuf:"bytes,1,opt,name=metadata,proto3,oneof"`
}

type CloneResponse_ChunkData struct {
	ChunkData []byte `protobuf:"bytes,2,opt,name=chunk_data,json=chunkData,proto3,oneof"`
}

func (*CloneResponse_Metadata) isCloneResponse_Data() {}

func (*CloneResponse_ChunkData) isCloneResponse_Data() {}

type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_pb_imagehub_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_pb_imagehub_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_v1_pb_imagehub_proto_rawDescGZIP(), []int{3}
}

func (x *RegisterResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RegisterResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RegisterResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username  string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Email     string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password  string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Password2 string `protobuf:"bytes,4,opt,name=password2,proto3" json:"password2,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_pb_imagehub_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_pb_imagehub_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_v1_pb_imagehub_proto_rawDescGZIP(), []int{4}
}

func (x *RegisterRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RegisterRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterRequest) GetPassword2() string {
	if x != nil {
		return x.Password2
	}
	return ""
}

type UserCredentials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username  string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password  string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Hash      uint32 `protobuf:"varint,3,opt,name=hash,proto3" json:"hash,omitempty"`
	ReposPath string `protobuf:"bytes,4,opt,name=repos_path,json=reposPath,proto3" json:"repos_path,omitempty"`
}

func (x *UserCredentials) Reset() {
	*x = UserCredentials{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_pb_imagehub_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCredentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCredentials) ProtoMessage() {}

func (x *UserCredentials) ProtoReflect() protoreflect.Message {
	mi := &file_v1_pb_imagehub_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCredentials.ProtoReflect.Descriptor instead.
func (*UserCredentials) Descriptor() ([]byte, []int) {
	return file_v1_pb_imagehub_proto_rawDescGZIP(), []int{5}
}

func (x *UserCredentials) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserCredentials) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserCredentials) GetHash() uint32 {
	if x != nil {
		return x.Hash
	}
	return 0
}

func (x *UserCredentials) GetReposPath() string {
	if x != nil {
		return x.ReposPath
	}
	return ""
}

type PushRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//	*PushRequest_Info
	//	*PushRequest_ChunkData
	Data isPushRequest_Data `protobuf_oneof:"data"`
}

func (x *PushRequest) Reset() {
	*x = PushRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_pb_imagehub_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushRequest) ProtoMessage() {}

func (x *PushRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_pb_imagehub_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushRequest.ProtoReflect.Descriptor instead.
func (*PushRequest) Descriptor() ([]byte, []int) {
	return file_v1_pb_imagehub_proto_rawDescGZIP(), []int{6}
}

func (m *PushRequest) GetData() isPushRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *PushRequest) GetInfo() *UserCredentials {
	if x, ok := x.GetData().(*PushRequest_Info); ok {
		return x.Info
	}
	return nil
}

func (x *PushRequest) GetChunkData() []byte {
	if x, ok := x.GetData().(*PushRequest_ChunkData); ok {
		return x.ChunkData
	}
	return nil
}

type isPushRequest_Data interface {
	isPushRequest_Data()
}

type PushRequest_Info struct {
	Info *UserCredentials `protobuf:"bytes,1,opt,name=info,proto3,oneof"`
}

type PushRequest_ChunkData struct {
	ChunkData []byte `protobuf:"bytes,2,opt,name=chunk_data,json=chunkData,proto3,oneof"`
}

func (*PushRequest_Info) isPushRequest_Data() {}

func (*PushRequest_ChunkData) isPushRequest_Data() {}

type PushResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *PushResponse) Reset() {
	*x = PushResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_pb_imagehub_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushResponse) ProtoMessage() {}

func (x *PushResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_pb_imagehub_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushResponse.ProtoReflect.Descriptor instead.
func (*PushResponse) Descriptor() ([]byte, []int) {
	return file_v1_pb_imagehub_proto_rawDescGZIP(), []int{7}
}

func (x *PushResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type CheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *MetaData `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *CheckRequest) Reset() {
	*x = CheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_pb_imagehub_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckRequest) ProtoMessage() {}

func (x *CheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_pb_imagehub_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckRequest.ProtoReflect.Descriptor instead.
func (*CheckRequest) Descriptor() ([]byte, []int) {
	return file_v1_pb_imagehub_proto_rawDescGZIP(), []int{8}
}

func (x *CheckRequest) GetMetadata() *MetaData {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type CheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status CheckStatus `protobuf:"varint,1,opt,name=status,proto3,enum=imagehub.CheckStatus" json:"status,omitempty"`
}

func (x *CheckResponse) Reset() {
	*x = CheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_pb_imagehub_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckResponse) ProtoMessage() {}

func (x *CheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_pb_imagehub_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckResponse.ProtoReflect.Descriptor instead.
func (*CheckResponse) Descriptor() ([]byte, []int) {
	return file_v1_pb_imagehub_proto_rawDescGZIP(), []int{9}
}

func (x *CheckResponse) GetStatus() CheckStatus {
	if x != nil {
		return x.Status
	}
	return CheckStatus_UpToDate
}

var File_v1_pb_imagehub_proto protoreflect.FileDescriptor

var file_v1_pb_imagehub_proto_rawDesc = []byte{
	0x0a, 0x14, 0x76, 0x31, 0x2f, 0x70, 0x62, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x68, 0x75, 0x62,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x68, 0x75, 0x62,
	0x22, 0x2d, 0x0a, 0x0c, 0x43, 0x6c, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x50, 0x61, 0x74, 0x68, 0x22,
	0x55, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x68,
	0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12,
	0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x6a, 0x0a, 0x0d, 0x43, 0x6c, 0x6f, 0x6e, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x68, 0x75, 0x62, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0a, 0x63, 0x68, 0x75,
	0x6e, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52,
	0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x54, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x7d, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x32, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x32, 0x22, 0x7c, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x43,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x5f,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x70, 0x6f,
	0x73, 0x50, 0x61, 0x74, 0x68, 0x22, 0x67, 0x0a, 0x0b, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x68, 0x75, 0x62, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x48, 0x00, 0x52,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0a, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x09, 0x63, 0x68, 0x75,
	0x6e, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x26,
	0x0a, 0x0c, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x3e, 0x0a, 0x0c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x68, 0x75, 0x62, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3e, 0x0a, 0x0d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x68,
	0x75, 0x62, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x2c, 0x0a, 0x0b, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x70, 0x54, 0x6f, 0x44, 0x61, 0x74,
	0x65, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x75,
	0x6e, 0x64, 0x10, 0x01, 0x32, 0xfe, 0x01, 0x0a, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x70, 0x6f, 0x73, 0x12, 0x3a, 0x0a, 0x05, 0x43, 0x6c, 0x6f, 0x6e, 0x65, 0x12, 0x16, 0x2e, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x68, 0x75, 0x62, 0x2e, 0x43, 0x6c, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x68, 0x75, 0x62, 0x2e,
	0x43, 0x6c, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12,
	0x41, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x68, 0x75, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x68, 0x75,
	0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x50, 0x75, 0x73, 0x68, 0x12, 0x15, 0x2e, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x68, 0x75, 0x62, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x68, 0x75, 0x62, 0x2e, 0x50, 0x75, 0x73,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x38, 0x0a, 0x05, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x12, 0x16, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x68, 0x75, 0x62, 0x2e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x68, 0x75, 0x62, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_pb_imagehub_proto_rawDescOnce sync.Once
	file_v1_pb_imagehub_proto_rawDescData = file_v1_pb_imagehub_proto_rawDesc
)

func file_v1_pb_imagehub_proto_rawDescGZIP() []byte {
	file_v1_pb_imagehub_proto_rawDescOnce.Do(func() {
		file_v1_pb_imagehub_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_pb_imagehub_proto_rawDescData)
	})
	return file_v1_pb_imagehub_proto_rawDescData
}

var file_v1_pb_imagehub_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_v1_pb_imagehub_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_v1_pb_imagehub_proto_goTypes = []interface{}{
	(CheckStatus)(0),         // 0: imagehub.CheckStatus
	(*CloneRequest)(nil),     // 1: imagehub.CloneRequest
	(*MetaData)(nil),         // 2: imagehub.MetaData
	(*CloneResponse)(nil),    // 3: imagehub.CloneResponse
	(*RegisterResponse)(nil), // 4: imagehub.RegisterResponse
	(*RegisterRequest)(nil),  // 5: imagehub.RegisterRequest
	(*UserCredentials)(nil),  // 6: imagehub.UserCredentials
	(*PushRequest)(nil),      // 7: imagehub.PushRequest
	(*PushResponse)(nil),     // 8: imagehub.PushResponse
	(*CheckRequest)(nil),     // 9: imagehub.CheckRequest
	(*CheckResponse)(nil),    // 10: imagehub.CheckResponse
}
var file_v1_pb_imagehub_proto_depIdxs = []int32{
	2,  // 0: imagehub.CloneResponse.metadata:type_name -> imagehub.MetaData
	6,  // 1: imagehub.PushRequest.info:type_name -> imagehub.UserCredentials
	2,  // 2: imagehub.CheckRequest.metadata:type_name -> imagehub.MetaData
	0,  // 3: imagehub.CheckResponse.status:type_name -> imagehub.CheckStatus
	1,  // 4: imagehub.imageRepos.Clone:input_type -> imagehub.CloneRequest
	5,  // 5: imagehub.imageRepos.Register:input_type -> imagehub.RegisterRequest
	7,  // 6: imagehub.imageRepos.Push:input_type -> imagehub.PushRequest
	9,  // 7: imagehub.imageRepos.Check:input_type -> imagehub.CheckRequest
	3,  // 8: imagehub.imageRepos.Clone:output_type -> imagehub.CloneResponse
	4,  // 9: imagehub.imageRepos.Register:output_type -> imagehub.RegisterResponse
	8,  // 10: imagehub.imageRepos.Push:output_type -> imagehub.PushResponse
	10, // 11: imagehub.imageRepos.Check:output_type -> imagehub.CheckResponse
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_v1_pb_imagehub_proto_init() }
func file_v1_pb_imagehub_proto_init() {
	if File_v1_pb_imagehub_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_pb_imagehub_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloneRequest); i {
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
		file_v1_pb_imagehub_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetaData); i {
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
		file_v1_pb_imagehub_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloneResponse); i {
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
		file_v1_pb_imagehub_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResponse); i {
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
		file_v1_pb_imagehub_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_v1_pb_imagehub_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCredentials); i {
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
		file_v1_pb_imagehub_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushRequest); i {
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
		file_v1_pb_imagehub_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushResponse); i {
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
		file_v1_pb_imagehub_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckRequest); i {
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
		file_v1_pb_imagehub_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckResponse); i {
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
	file_v1_pb_imagehub_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*CloneResponse_Metadata)(nil),
		(*CloneResponse_ChunkData)(nil),
	}
	file_v1_pb_imagehub_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*PushRequest_Info)(nil),
		(*PushRequest_ChunkData)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_pb_imagehub_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_pb_imagehub_proto_goTypes,
		DependencyIndexes: file_v1_pb_imagehub_proto_depIdxs,
		EnumInfos:         file_v1_pb_imagehub_proto_enumTypes,
		MessageInfos:      file_v1_pb_imagehub_proto_msgTypes,
	}.Build()
	File_v1_pb_imagehub_proto = out.File
	file_v1_pb_imagehub_proto_rawDesc = nil
	file_v1_pb_imagehub_proto_goTypes = nil
	file_v1_pb_imagehub_proto_depIdxs = nil
}
