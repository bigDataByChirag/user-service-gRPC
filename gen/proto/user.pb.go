// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: proto/user.proto

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

type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Fname   string `protobuf:"bytes,2,opt,name=fname,proto3" json:"fname,omitempty"`
	City    string `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	Phone   uint64 `protobuf:"varint,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Height  string `protobuf:"bytes,5,opt,name=height,proto3" json:"height,omitempty"`
	Married bool   `protobuf:"varint,6,opt,name=married,proto3" json:"married,omitempty"`
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{0}
}

func (x *Person) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Person) GetFname() string {
	if x != nil {
		return x.Fname
	}
	return ""
}

func (x *Person) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Person) GetPhone() uint64 {
	if x != nil {
		return x.Phone
	}
	return 0
}

func (x *Person) GetHeight() string {
	if x != nil {
		return x.Height
	}
	return ""
}

func (x *Person) GetMarried() bool {
	if x != nil {
		return x.Married
	}
	return false
}

type SignupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fname   string `protobuf:"bytes,2,opt,name=fname,proto3" json:"fname,omitempty"`
	City    string `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	Phone   uint64 `protobuf:"varint,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Height  string `protobuf:"bytes,5,opt,name=height,proto3" json:"height,omitempty"`
	Married bool   `protobuf:"varint,6,opt,name=married,proto3" json:"married,omitempty"`
}

func (x *SignupRequest) Reset() {
	*x = SignupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignupRequest) ProtoMessage() {}

func (x *SignupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignupRequest.ProtoReflect.Descriptor instead.
func (*SignupRequest) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{1}
}

func (x *SignupRequest) GetFname() string {
	if x != nil {
		return x.Fname
	}
	return ""
}

func (x *SignupRequest) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *SignupRequest) GetPhone() uint64 {
	if x != nil {
		return x.Phone
	}
	return 0
}

func (x *SignupRequest) GetHeight() string {
	if x != nil {
		return x.Height
	}
	return ""
}

func (x *SignupRequest) GetMarried() bool {
	if x != nil {
		return x.Married
	}
	return false
}

type SignupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Person *Person `protobuf:"bytes,1,opt,name=person,proto3" json:"person,omitempty"`
}

func (x *SignupResponse) Reset() {
	*x = SignupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignupResponse) ProtoMessage() {}

func (x *SignupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignupResponse.ProtoReflect.Descriptor instead.
func (*SignupResponse) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{2}
}

func (x *SignupResponse) GetPerson() *Person {
	if x != nil {
		return x.Person
	}
	return nil
}

type SearchUserByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SearchUserByIdRequest) Reset() {
	*x = SearchUserByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchUserByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchUserByIdRequest) ProtoMessage() {}

func (x *SearchUserByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchUserByIdRequest.ProtoReflect.Descriptor instead.
func (*SearchUserByIdRequest) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{3}
}

func (x *SearchUserByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SearchUserByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Person *Person `protobuf:"bytes,1,opt,name=person,proto3" json:"person,omitempty"`
}

func (x *SearchUserByIdResponse) Reset() {
	*x = SearchUserByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchUserByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchUserByIdResponse) ProtoMessage() {}

func (x *SearchUserByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchUserByIdResponse.ProtoReflect.Descriptor instead.
func (*SearchUserByIdResponse) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{4}
}

func (x *SearchUserByIdResponse) GetPerson() *Person {
	if x != nil {
		return x.Person
	}
	return nil
}

type SearchUsersByIdsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *SearchUsersByIdsRequest) Reset() {
	*x = SearchUsersByIdsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchUsersByIdsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchUsersByIdsRequest) ProtoMessage() {}

func (x *SearchUsersByIdsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchUsersByIdsRequest.ProtoReflect.Descriptor instead.
func (*SearchUsersByIdsRequest) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{5}
}

func (x *SearchUsersByIdsRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type SearchUsersByIdsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	People []*Person `protobuf:"bytes,1,rep,name=people,proto3" json:"people,omitempty"`
}

func (x *SearchUsersByIdsResponse) Reset() {
	*x = SearchUsersByIdsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchUsersByIdsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchUsersByIdsResponse) ProtoMessage() {}

func (x *SearchUsersByIdsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchUsersByIdsResponse.ProtoReflect.Descriptor instead.
func (*SearchUsersByIdsResponse) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{6}
}

func (x *SearchUsersByIdsResponse) GetPeople() []*Person {
	if x != nil {
		return x.People
	}
	return nil
}

type GetAllUserIdsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllUserIdsRequest) Reset() {
	*x = GetAllUserIdsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllUserIdsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllUserIdsRequest) ProtoMessage() {}

func (x *GetAllUserIdsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllUserIdsRequest.ProtoReflect.Descriptor instead.
func (*GetAllUserIdsRequest) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{7}
}

type GetAllUserIdsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *GetAllUserIdsResponse) Reset() {
	*x = GetAllUserIdsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllUserIdsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllUserIdsResponse) ProtoMessage() {}

func (x *GetAllUserIdsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllUserIdsResponse.ProtoReflect.Descriptor instead.
func (*GetAllUserIdsResponse) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{8}
}

func (x *GetAllUserIdsResponse) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type SearchUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fname   string `protobuf:"bytes,1,opt,name=fname,proto3" json:"fname,omitempty"`
	City    string `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	Phone   uint64 `protobuf:"varint,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Height  string `protobuf:"bytes,4,opt,name=height,proto3" json:"height,omitempty"`
	Married bool   `protobuf:"varint,5,opt,name=married,proto3" json:"married,omitempty"`
}

func (x *SearchUsersRequest) Reset() {
	*x = SearchUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchUsersRequest) ProtoMessage() {}

func (x *SearchUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchUsersRequest.ProtoReflect.Descriptor instead.
func (*SearchUsersRequest) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{9}
}

func (x *SearchUsersRequest) GetFname() string {
	if x != nil {
		return x.Fname
	}
	return ""
}

func (x *SearchUsersRequest) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *SearchUsersRequest) GetPhone() uint64 {
	if x != nil {
		return x.Phone
	}
	return 0
}

func (x *SearchUsersRequest) GetHeight() string {
	if x != nil {
		return x.Height
	}
	return ""
}

func (x *SearchUsersRequest) GetMarried() bool {
	if x != nil {
		return x.Married
	}
	return false
}

type SearchUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	People []*Person `protobuf:"bytes,1,rep,name=people,proto3" json:"people,omitempty"`
}

func (x *SearchUsersResponse) Reset() {
	*x = SearchUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchUsersResponse) ProtoMessage() {}

func (x *SearchUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchUsersResponse.ProtoReflect.Descriptor instead.
func (*SearchUsersResponse) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{10}
}

func (x *SearchUsersResponse) GetPeople() []*Person {
	if x != nil {
		return x.People
	}
	return nil
}

var File_proto_user_proto protoreflect.FileDescriptor

var file_proto_user_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x01, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x66, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x72,
	0x72, 0x69, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x6d, 0x61, 0x72, 0x72,
	0x69, 0x65, 0x64, 0x22, 0x81, 0x01, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x61, 0x72, 0x72, 0x69, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x6d, 0x61, 0x72, 0x72, 0x69, 0x65, 0x64, 0x22, 0x37, 0x0a, 0x0e, 0x53, 0x69, 0x67, 0x6e, 0x75,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x22, 0x27, 0x0a, 0x15, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3f, 0x0a, 0x16, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x52, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x22, 0x2b, 0x0a, 0x17, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x73, 0x42, 0x79, 0x49, 0x64, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x41, 0x0a, 0x18, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x55, 0x73, 0x65, 0x72, 0x73, 0x42, 0x79, 0x49, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x52, 0x06, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x22, 0x16, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x29, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x86, 0x01,
	0x0a, 0x12, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69,
	0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x61, 0x72, 0x72, 0x69, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x6d,
	0x61, 0x72, 0x72, 0x69, 0x65, 0x64, 0x22, 0x3c, 0x0a, 0x13, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a,
	0x06, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x70, 0x65,
	0x6f, 0x70, 0x6c, 0x65, 0x32, 0xf5, 0x02, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x12, 0x14,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x69, 0x67,
	0x6e, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x10, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x73, 0x42, 0x79, 0x49, 0x64, 0x73, 0x12, 0x1e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x42, 0x79, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x42, 0x79, 0x49, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x45, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x12, 0x5a, 0x10,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_user_proto_rawDescOnce sync.Once
	file_proto_user_proto_rawDescData = file_proto_user_proto_rawDesc
)

func file_proto_user_proto_rawDescGZIP() []byte {
	file_proto_user_proto_rawDescOnce.Do(func() {
		file_proto_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_user_proto_rawDescData)
	})
	return file_proto_user_proto_rawDescData
}

var file_proto_user_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_user_proto_goTypes = []any{
	(*Person)(nil),                   // 0: proto.Person
	(*SignupRequest)(nil),            // 1: proto.SignupRequest
	(*SignupResponse)(nil),           // 2: proto.SignupResponse
	(*SearchUserByIdRequest)(nil),    // 3: proto.SearchUserByIdRequest
	(*SearchUserByIdResponse)(nil),   // 4: proto.SearchUserByIdResponse
	(*SearchUsersByIdsRequest)(nil),  // 5: proto.SearchUsersByIdsRequest
	(*SearchUsersByIdsResponse)(nil), // 6: proto.SearchUsersByIdsResponse
	(*GetAllUserIdsRequest)(nil),     // 7: proto.GetAllUserIdsRequest
	(*GetAllUserIdsResponse)(nil),    // 8: proto.GetAllUserIdsResponse
	(*SearchUsersRequest)(nil),       // 9: proto.SearchUsersRequest
	(*SearchUsersResponse)(nil),      // 10: proto.SearchUsersResponse
	(*emptypb.Empty)(nil),            // 11: google.protobuf.Empty
}
var file_proto_user_proto_depIdxs = []int32{
	0,  // 0: proto.SignupResponse.person:type_name -> proto.Person
	0,  // 1: proto.SearchUserByIdResponse.person:type_name -> proto.Person
	0,  // 2: proto.SearchUsersByIdsResponse.people:type_name -> proto.Person
	0,  // 3: proto.SearchUsersResponse.people:type_name -> proto.Person
	1,  // 4: proto.UserService.Signup:input_type -> proto.SignupRequest
	3,  // 5: proto.UserService.SearchUserById:input_type -> proto.SearchUserByIdRequest
	5,  // 6: proto.UserService.SearchUsersByIds:input_type -> proto.SearchUsersByIdsRequest
	11, // 7: proto.UserService.GetAllUserIds:input_type -> google.protobuf.Empty
	9,  // 8: proto.UserService.SearchUsers:input_type -> proto.SearchUsersRequest
	2,  // 9: proto.UserService.Signup:output_type -> proto.SignupResponse
	4,  // 10: proto.UserService.SearchUserById:output_type -> proto.SearchUserByIdResponse
	6,  // 11: proto.UserService.SearchUsersByIds:output_type -> proto.SearchUsersByIdsResponse
	8,  // 12: proto.UserService.GetAllUserIds:output_type -> proto.GetAllUserIdsResponse
	10, // 13: proto.UserService.SearchUsers:output_type -> proto.SearchUsersResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_proto_user_proto_init() }
func file_proto_user_proto_init() {
	if File_proto_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_user_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Person); i {
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
		file_proto_user_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SignupRequest); i {
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
		file_proto_user_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*SignupResponse); i {
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
		file_proto_user_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*SearchUserByIdRequest); i {
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
		file_proto_user_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*SearchUserByIdResponse); i {
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
		file_proto_user_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*SearchUsersByIdsRequest); i {
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
		file_proto_user_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*SearchUsersByIdsResponse); i {
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
		file_proto_user_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllUserIdsRequest); i {
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
		file_proto_user_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllUserIdsResponse); i {
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
		file_proto_user_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*SearchUsersRequest); i {
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
		file_proto_user_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*SearchUsersResponse); i {
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
			RawDescriptor: file_proto_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_user_proto_goTypes,
		DependencyIndexes: file_proto_user_proto_depIdxs,
		MessageInfos:      file_proto_user_proto_msgTypes,
	}.Build()
	File_proto_user_proto = out.File
	file_proto_user_proto_rawDesc = nil
	file_proto_user_proto_goTypes = nil
	file_proto_user_proto_depIdxs = nil
}
