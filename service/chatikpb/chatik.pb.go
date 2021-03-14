// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.5
// source: service/chatik.proto

package chatikpb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	typepb "github.com/jamolh/chatik-sharedpb/type/typepb"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_chatik_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_chatik_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_service_chatik_proto_rawDescGZIP(), []int{0}
}

func (x *PingRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_chatik_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_chatik_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_service_chatik_proto_rawDescGZIP(), []int{1}
}

func (x *PingResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GenerateJWTRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *GenerateJWTRequest) Reset() {
	*x = GenerateJWTRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_chatik_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateJWTRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateJWTRequest) ProtoMessage() {}

func (x *GenerateJWTRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_chatik_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateJWTRequest.ProtoReflect.Descriptor instead.
func (*GenerateJWTRequest) Descriptor() ([]byte, []int) {
	return file_service_chatik_proto_rawDescGZIP(), []int{2}
}

func (x *GenerateJWTRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GenerateJWTRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type GenerateJWTResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	//   type.Status status = 2;
	Token string       `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	User  *typepb.User `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GenerateJWTResponse) Reset() {
	*x = GenerateJWTResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_chatik_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateJWTResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateJWTResponse) ProtoMessage() {}

func (x *GenerateJWTResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_chatik_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateJWTResponse.ProtoReflect.Descriptor instead.
func (*GenerateJWTResponse) Descriptor() ([]byte, []int) {
	return file_service_chatik_proto_rawDescGZIP(), []int{3}
}

func (x *GenerateJWTResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *GenerateJWTResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GenerateJWTResponse) GetUser() *typepb.User {
	if x != nil {
		return x.User
	}
	return nil
}

type ValidateJWTRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Token     string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ValidateJWTRequest) Reset() {
	*x = ValidateJWTRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_chatik_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateJWTRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateJWTRequest) ProtoMessage() {}

func (x *ValidateJWTRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_chatik_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateJWTRequest.ProtoReflect.Descriptor instead.
func (*ValidateJWTRequest) Descriptor() ([]byte, []int) {
	return file_service_chatik_proto_rawDescGZIP(), []int{4}
}

func (x *ValidateJWTRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *ValidateJWTRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ValidateJWTResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"` // //   type.Status status = 2;
}

func (x *ValidateJWTResponse) Reset() {
	*x = ValidateJWTResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_chatik_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateJWTResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateJWTResponse) ProtoMessage() {}

func (x *ValidateJWTResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_chatik_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateJWTResponse.ProtoReflect.Descriptor instead.
func (*ValidateJWTResponse) Descriptor() ([]byte, []int) {
	return file_service_chatik_proto_rawDescGZIP(), []int{5}
}

func (x *ValidateJWTResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

type JWTAuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Token     string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Api       string `protobuf:"bytes,3,opt,name=api,proto3" json:"api,omitempty"`
	Action    string `protobuf:"bytes,4,opt,name=action,proto3" json:"action,omitempty"`
}

func (x *JWTAuthRequest) Reset() {
	*x = JWTAuthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_chatik_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JWTAuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JWTAuthRequest) ProtoMessage() {}

func (x *JWTAuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_chatik_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JWTAuthRequest.ProtoReflect.Descriptor instead.
func (*JWTAuthRequest) Descriptor() ([]byte, []int) {
	return file_service_chatik_proto_rawDescGZIP(), []int{6}
}

func (x *JWTAuthRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *JWTAuthRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *JWTAuthRequest) GetApi() string {
	if x != nil {
		return x.Api
	}
	return ""
}

func (x *JWTAuthRequest) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

type JWTAuthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"` // //   type.Status status = 2;
}

func (x *JWTAuthResponse) Reset() {
	*x = JWTAuthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_chatik_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JWTAuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JWTAuthResponse) ProtoMessage() {}

func (x *JWTAuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_chatik_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JWTAuthResponse.ProtoReflect.Descriptor instead.
func (*JWTAuthResponse) Descriptor() ([]byte, []int) {
	return file_service_chatik_proto_rawDescGZIP(), []int{7}
}

func (x *JWTAuthResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

var File_service_chatik_proto protoreflect.FileDescriptor

var file_service_chatik_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x69, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x0f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x27, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x28, 0x0a, 0x0c, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x4c, 0x0a, 0x12, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4a, 0x57, 0x54, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x71, 0x0a, 0x13, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x4a, 0x57, 0x54, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x25, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x49, 0x0a, 0x12, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x57, 0x54, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x34, 0x0a, 0x13, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x4a, 0x57, 0x54, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x22, 0x6f, 0x0a, 0x0e,
	0x4a, 0x57, 0x54, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x61, 0x70, 0x69, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x30, 0x0a,
	0x0f, 0x4a, 0x57, 0x54, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x32,
	0x52, 0x0a, 0x0d, 0x43, 0x68, 0x61, 0x74, 0x69, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x41, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x1b, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6a, 0x61, 0x6d, 0x6f, 0x6c, 0x68, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x69, 0x6b, 0x2d,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x70, 0x62, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x63, 0x68, 0x61, 0x74, 0x69, 0x6b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_service_chatik_proto_rawDescOnce sync.Once
	file_service_chatik_proto_rawDescData = file_service_chatik_proto_rawDesc
)

func file_service_chatik_proto_rawDescGZIP() []byte {
	file_service_chatik_proto_rawDescOnce.Do(func() {
		file_service_chatik_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_chatik_proto_rawDescData)
	})
	return file_service_chatik_proto_rawDescData
}

var file_service_chatik_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_service_chatik_proto_goTypes = []interface{}{
	(*PingRequest)(nil),         // 0: shared.service.PingRequest
	(*PingResponse)(nil),        // 1: shared.service.PingResponse
	(*GenerateJWTRequest)(nil),  // 2: shared.service.GenerateJWTRequest
	(*GenerateJWTResponse)(nil), // 3: shared.service.GenerateJWTResponse
	(*ValidateJWTRequest)(nil),  // 4: shared.service.ValidateJWTRequest
	(*ValidateJWTResponse)(nil), // 5: shared.service.ValidateJWTResponse
	(*JWTAuthRequest)(nil),      // 6: shared.service.JWTAuthRequest
	(*JWTAuthResponse)(nil),     // 7: shared.service.JWTAuthResponse
	(*typepb.User)(nil),         // 8: shared.type.User
}
var file_service_chatik_proto_depIdxs = []int32{
	8, // 0: shared.service.GenerateJWTResponse.user:type_name -> shared.type.User
	0, // 1: shared.service.ChatikService.Ping:input_type -> shared.service.PingRequest
	1, // 2: shared.service.ChatikService.Ping:output_type -> shared.service.PingResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_service_chatik_proto_init() }
func file_service_chatik_proto_init() {
	if File_service_chatik_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_chatik_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
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
		file_service_chatik_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResponse); i {
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
		file_service_chatik_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateJWTRequest); i {
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
		file_service_chatik_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateJWTResponse); i {
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
		file_service_chatik_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateJWTRequest); i {
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
		file_service_chatik_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateJWTResponse); i {
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
		file_service_chatik_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JWTAuthRequest); i {
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
		file_service_chatik_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JWTAuthResponse); i {
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
			RawDescriptor: file_service_chatik_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_chatik_proto_goTypes,
		DependencyIndexes: file_service_chatik_proto_depIdxs,
		MessageInfos:      file_service_chatik_proto_msgTypes,
	}.Build()
	File_service_chatik_proto = out.File
	file_service_chatik_proto_rawDesc = nil
	file_service_chatik_proto_goTypes = nil
	file_service_chatik_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChatikServiceClient is the client API for ChatikService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatikServiceClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
}

type chatikServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatikServiceClient(cc grpc.ClientConnInterface) ChatikServiceClient {
	return &chatikServiceClient{cc}
}

func (c *chatikServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/shared.service.ChatikService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatikServiceServer is the server API for ChatikService service.
type ChatikServiceServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
}

// UnimplementedChatikServiceServer can be embedded to have forward compatible implementations.
type UnimplementedChatikServiceServer struct {
}

func (*UnimplementedChatikServiceServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}

func RegisterChatikServiceServer(s *grpc.Server, srv ChatikServiceServer) {
	s.RegisterService(&_ChatikService_serviceDesc, srv)
}

func _ChatikService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatikServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shared.service.ChatikService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatikServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChatikService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "shared.service.ChatikService",
	HandlerType: (*ChatikServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _ChatikService_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/chatik.proto",
}
