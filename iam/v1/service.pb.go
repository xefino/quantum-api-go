// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: protos/frontend/iam/v1/service.proto

package v1

import (
	data "github.com/xefino/quantum-api-go/data"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Request to register a new client associated with the current user
type RegisterClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"` // The domain associated with the client; this value may be used to
	// deny access to certain endpoints
	ClientType data.ClientType `protobuf:"varint,2,opt,name=client_type,json=clientType,proto3,enum=protos.frontend.data.ClientType" json:"client_type,omitempty"` // The type of client (web, CLI, other)
}

func (x *RegisterClientRequest) Reset() {
	*x = RegisterClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_frontend_iam_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterClientRequest) ProtoMessage() {}

func (x *RegisterClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_frontend_iam_v1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterClientRequest.ProtoReflect.Descriptor instead.
func (*RegisterClientRequest) Descriptor() ([]byte, []int) {
	return file_protos_frontend_iam_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterClientRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *RegisterClientRequest) GetClientType() data.ClientType {
	if x != nil {
		return x.ClientType
	}
	return data.ClientType(0)
}

// Response retrieved after the client registration is complete
type RegisterClientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                           // The ID of the client, as recorded in the system. This value should be saved for later
	Secret string          `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`                                   // The secret associated with this client. This value will only be made available once
	Domain string          `protobuf:"bytes,3,opt,name=domain,proto3" json:"domain,omitempty"`                                   // The domain for which the client is registered. This value may be used to deny access to certain resources
	Type   data.ClientType `protobuf:"varint,4,opt,name=type,proto3,enum=protos.frontend.data.ClientType" json:"type,omitempty"` // The type of the client. This value may be used to deny access to certain resources
}

func (x *RegisterClientResponse) Reset() {
	*x = RegisterClientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_frontend_iam_v1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterClientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterClientResponse) ProtoMessage() {}

func (x *RegisterClientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_frontend_iam_v1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterClientResponse.ProtoReflect.Descriptor instead.
func (*RegisterClientResponse) Descriptor() ([]byte, []int) {
	return file_protos_frontend_iam_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterClientResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RegisterClientResponse) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *RegisterClientResponse) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *RegisterClientResponse) GetType() data.ClientType {
	if x != nil {
		return x.Type
	}
	return data.ClientType(0)
}

// Request for getting all clients for the current user
type GetClientsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetClientsRequest) Reset() {
	*x = GetClientsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_frontend_iam_v1_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClientsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClientsRequest) ProtoMessage() {}

func (x *GetClientsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_frontend_iam_v1_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClientsRequest.ProtoReflect.Descriptor instead.
func (*GetClientsRequest) Descriptor() ([]byte, []int) {
	return file_protos_frontend_iam_v1_service_proto_rawDescGZIP(), []int{2}
}

// Request for getting a particular client associated with the current user
type GetClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"` // The ID of the client being retrieved
}

func (x *GetClientRequest) Reset() {
	*x = GetClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_frontend_iam_v1_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClientRequest) ProtoMessage() {}

func (x *GetClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_frontend_iam_v1_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClientRequest.ProtoReflect.Descriptor instead.
func (*GetClientRequest) Descriptor() ([]byte, []int) {
	return file_protos_frontend_iam_v1_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetClientRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

// Request for deleting an existing client associated with the current user
type DeleteClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"` // The ID of the client being deleted
}

func (x *DeleteClientRequest) Reset() {
	*x = DeleteClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_frontend_iam_v1_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteClientRequest) ProtoMessage() {}

func (x *DeleteClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_frontend_iam_v1_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteClientRequest.ProtoReflect.Descriptor instead.
func (*DeleteClientRequest) Descriptor() ([]byte, []int) {
	return file_protos_frontend_iam_v1_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteClientRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

// Request to create a new user
type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`       // The email address associated with the user
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"` // The cleartext password the user will logon with
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_frontend_iam_v1_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_frontend_iam_v1_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_protos_frontend_iam_v1_service_proto_rawDescGZIP(), []int{5}
}

func (x *CreateUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

// Request for getting the current user
type GetUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_frontend_iam_v1_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_frontend_iam_v1_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_protos_frontend_iam_v1_service_proto_rawDescGZIP(), []int{6}
}

// Request to update a user's email address
type UpdateEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewEmail string `protobuf:"bytes,1,opt,name=new_email,json=newEmail,proto3" json:"new_email,omitempty"` // The user's new email address
}

func (x *UpdateEmailRequest) Reset() {
	*x = UpdateEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_frontend_iam_v1_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEmailRequest) ProtoMessage() {}

func (x *UpdateEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_frontend_iam_v1_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEmailRequest.ProtoReflect.Descriptor instead.
func (*UpdateEmailRequest) Descriptor() ([]byte, []int) {
	return file_protos_frontend_iam_v1_service_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateEmailRequest) GetNewEmail() string {
	if x != nil {
		return x.NewEmail
	}
	return ""
}

// Request to update a user's password
type UpdatePasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewPassword string `protobuf:"bytes,1,opt,name=new_password,json=newPassword,proto3" json:"new_password,omitempty"` // The user's new cleartext password
}

func (x *UpdatePasswordRequest) Reset() {
	*x = UpdatePasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_frontend_iam_v1_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePasswordRequest) ProtoMessage() {}

func (x *UpdatePasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_frontend_iam_v1_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePasswordRequest.ProtoReflect.Descriptor instead.
func (*UpdatePasswordRequest) Descriptor() ([]byte, []int) {
	return file_protos_frontend_iam_v1_service_proto_rawDescGZIP(), []int{8}
}

func (x *UpdatePasswordRequest) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

var File_protos_frontend_iam_v1_service_proto protoreflect.FileDescriptor

var file_protos_frontend_iam_v1_service_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x64, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x66,
	0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x2f, 0x69, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x72, 0x0a, 0x15,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x41, 0x0a,
	0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x66, 0x72, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x22, 0x8e, 0x01, 0x0a, 0x16, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x34, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2f, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x32, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x45, 0x0a, 0x11, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x31, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x65,
	0x77, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e,
	0x65, 0x77, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x3a, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x32, 0xad, 0x07, 0x0a, 0x0a, 0x49, 0x61, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x8d, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x2d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x66,
	0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x66, 0x72,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x11, 0x2f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x3a,
	0x01, 0x2a, 0x12, 0x69, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x64, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0a, 0x12, 0x08, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x30, 0x01, 0x12, 0x71, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x28, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x69, 0x61, 0x6d,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x66, 0x72,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x7d,
	0x12, 0x77, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x12, 0x2b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x64, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x1c, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x16, 0x2a, 0x14, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x65, 0x0a, 0x0a, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x66, 0x72, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x10,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x22, 0x05, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x3a, 0x01, 0x2a,
	0x12, 0x5c, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x26, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x69, 0x61,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x66, 0x72, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22,
	0x0d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x07, 0x12, 0x05, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x12, 0x74,
	0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x2a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e,
	0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x1a, 0x12, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2d, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x3a, 0x01, 0x2a, 0x12, 0x7d, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x2d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x66,
	0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x1a, 0x15, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2d, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x3a, 0x01, 0x2a, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x78, 0x65, 0x66, 0x69, 0x6e, 0x6f, 0x2f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x75, 0x6d,
	0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_frontend_iam_v1_service_proto_rawDescOnce sync.Once
	file_protos_frontend_iam_v1_service_proto_rawDescData = file_protos_frontend_iam_v1_service_proto_rawDesc
)

func file_protos_frontend_iam_v1_service_proto_rawDescGZIP() []byte {
	file_protos_frontend_iam_v1_service_proto_rawDescOnce.Do(func() {
		file_protos_frontend_iam_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_frontend_iam_v1_service_proto_rawDescData)
	})
	return file_protos_frontend_iam_v1_service_proto_rawDescData
}

var file_protos_frontend_iam_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_protos_frontend_iam_v1_service_proto_goTypes = []interface{}{
	(*RegisterClientRequest)(nil),  // 0: protos.frontend.iam.v1.RegisterClientRequest
	(*RegisterClientResponse)(nil), // 1: protos.frontend.iam.v1.RegisterClientResponse
	(*GetClientsRequest)(nil),      // 2: protos.frontend.iam.v1.GetClientsRequest
	(*GetClientRequest)(nil),       // 3: protos.frontend.iam.v1.GetClientRequest
	(*DeleteClientRequest)(nil),    // 4: protos.frontend.iam.v1.DeleteClientRequest
	(*CreateUserRequest)(nil),      // 5: protos.frontend.iam.v1.CreateUserRequest
	(*GetUserRequest)(nil),         // 6: protos.frontend.iam.v1.GetUserRequest
	(*UpdateEmailRequest)(nil),     // 7: protos.frontend.iam.v1.UpdateEmailRequest
	(*UpdatePasswordRequest)(nil),  // 8: protos.frontend.iam.v1.UpdatePasswordRequest
	(data.ClientType)(0),           // 9: protos.frontend.data.ClientType
	(*data.Client)(nil),            // 10: protos.frontend.data.Client
	(*data.User)(nil),              // 11: protos.frontend.data.User
}
var file_protos_frontend_iam_v1_service_proto_depIdxs = []int32{
	9,  // 0: protos.frontend.iam.v1.RegisterClientRequest.client_type:type_name -> protos.frontend.data.ClientType
	9,  // 1: protos.frontend.iam.v1.RegisterClientResponse.type:type_name -> protos.frontend.data.ClientType
	0,  // 2: protos.frontend.iam.v1.IamService.RegisterClient:input_type -> protos.frontend.iam.v1.RegisterClientRequest
	2,  // 3: protos.frontend.iam.v1.IamService.GetClients:input_type -> protos.frontend.iam.v1.GetClientsRequest
	3,  // 4: protos.frontend.iam.v1.IamService.GetClient:input_type -> protos.frontend.iam.v1.GetClientRequest
	4,  // 5: protos.frontend.iam.v1.IamService.DeleteClient:input_type -> protos.frontend.iam.v1.DeleteClientRequest
	5,  // 6: protos.frontend.iam.v1.IamService.CreateUser:input_type -> protos.frontend.iam.v1.CreateUserRequest
	6,  // 7: protos.frontend.iam.v1.IamService.GetUser:input_type -> protos.frontend.iam.v1.GetUserRequest
	7,  // 8: protos.frontend.iam.v1.IamService.UpdateEmail:input_type -> protos.frontend.iam.v1.UpdateEmailRequest
	8,  // 9: protos.frontend.iam.v1.IamService.UpdatePassword:input_type -> protos.frontend.iam.v1.UpdatePasswordRequest
	1,  // 10: protos.frontend.iam.v1.IamService.RegisterClient:output_type -> protos.frontend.iam.v1.RegisterClientResponse
	10, // 11: protos.frontend.iam.v1.IamService.GetClients:output_type -> protos.frontend.data.Client
	10, // 12: protos.frontend.iam.v1.IamService.GetClient:output_type -> protos.frontend.data.Client
	10, // 13: protos.frontend.iam.v1.IamService.DeleteClient:output_type -> protos.frontend.data.Client
	11, // 14: protos.frontend.iam.v1.IamService.CreateUser:output_type -> protos.frontend.data.User
	11, // 15: protos.frontend.iam.v1.IamService.GetUser:output_type -> protos.frontend.data.User
	11, // 16: protos.frontend.iam.v1.IamService.UpdateEmail:output_type -> protos.frontend.data.User
	11, // 17: protos.frontend.iam.v1.IamService.UpdatePassword:output_type -> protos.frontend.data.User
	10, // [10:18] is the sub-list for method output_type
	2,  // [2:10] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_protos_frontend_iam_v1_service_proto_init() }
func file_protos_frontend_iam_v1_service_proto_init() {
	if File_protos_frontend_iam_v1_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_frontend_iam_v1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterClientRequest); i {
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
		file_protos_frontend_iam_v1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterClientResponse); i {
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
		file_protos_frontend_iam_v1_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClientsRequest); i {
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
		file_protos_frontend_iam_v1_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClientRequest); i {
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
		file_protos_frontend_iam_v1_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteClientRequest); i {
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
		file_protos_frontend_iam_v1_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
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
		file_protos_frontend_iam_v1_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserRequest); i {
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
		file_protos_frontend_iam_v1_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEmailRequest); i {
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
		file_protos_frontend_iam_v1_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePasswordRequest); i {
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
			RawDescriptor: file_protos_frontend_iam_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_frontend_iam_v1_service_proto_goTypes,
		DependencyIndexes: file_protos_frontend_iam_v1_service_proto_depIdxs,
		MessageInfos:      file_protos_frontend_iam_v1_service_proto_msgTypes,
	}.Build()
	File_protos_frontend_iam_v1_service_proto = out.File
	file_protos_frontend_iam_v1_service_proto_rawDesc = nil
	file_protos_frontend_iam_v1_service_proto_goTypes = nil
	file_protos_frontend_iam_v1_service_proto_depIdxs = nil
}
