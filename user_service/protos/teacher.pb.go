// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: protos/teacher.proto

package protos

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

type Teacher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	TeacherName string   `protobuf:"bytes,2,opt,name=TeacherName,proto3" json:"TeacherName,omitempty"`
	ClassID     []string `protobuf:"bytes,3,rep,name=ClassID,proto3" json:"ClassID,omitempty"`
	SubjectID   []string `protobuf:"bytes,4,rep,name=SubjectID,proto3" json:"SubjectID,omitempty"`
}

func (x *Teacher) Reset() {
	*x = Teacher{}
	mi := &file_protos_teacher_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Teacher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Teacher) ProtoMessage() {}

func (x *Teacher) ProtoReflect() protoreflect.Message {
	mi := &file_protos_teacher_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Teacher.ProtoReflect.Descriptor instead.
func (*Teacher) Descriptor() ([]byte, []int) {
	return file_protos_teacher_proto_rawDescGZIP(), []int{0}
}

func (x *Teacher) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Teacher) GetTeacherName() string {
	if x != nil {
		return x.TeacherName
	}
	return ""
}

func (x *Teacher) GetClassID() []string {
	if x != nil {
		return x.ClassID
	}
	return nil
}

func (x *Teacher) GetSubjectID() []string {
	if x != nil {
		return x.SubjectID
	}
	return nil
}

type CreateTeacherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Email       string `protobuf:"bytes,2,opt,name=Email,proto3" json:"Email,omitempty"`
	PhoneNumber string `protobuf:"bytes,3,opt,name=PhoneNumber,proto3" json:"PhoneNumber,omitempty"`
	Username    string `protobuf:"bytes,4,opt,name=Username,proto3" json:"Username,omitempty"`
	Password    string `protobuf:"bytes,5,opt,name=Password,proto3" json:"Password,omitempty"`
	TeacherName string `protobuf:"bytes,6,opt,name=TeacherName,proto3" json:"TeacherName,omitempty"`
}

func (x *CreateTeacherRequest) Reset() {
	*x = CreateTeacherRequest{}
	mi := &file_protos_teacher_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTeacherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTeacherRequest) ProtoMessage() {}

func (x *CreateTeacherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_teacher_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTeacherRequest.ProtoReflect.Descriptor instead.
func (*CreateTeacherRequest) Descriptor() ([]byte, []int) {
	return file_protos_teacher_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTeacherRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *CreateTeacherRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateTeacherRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *CreateTeacherRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateTeacherRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateTeacherRequest) GetTeacherName() string {
	if x != nil {
		return x.TeacherName
	}
	return ""
}

type CreateTeacherResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateTeacherResponse) Reset() {
	*x = CreateTeacherResponse{}
	mi := &file_protos_teacher_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTeacherResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTeacherResponse) ProtoMessage() {}

func (x *CreateTeacherResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_teacher_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTeacherResponse.ProtoReflect.Descriptor instead.
func (*CreateTeacherResponse) Descriptor() ([]byte, []int) {
	return file_protos_teacher_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTeacherResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetAllTeacherResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message  string     `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Teachers []*Teacher `protobuf:"bytes,2,rep,name=teachers,proto3" json:"teachers,omitempty"`
}

func (x *GetAllTeacherResponse) Reset() {
	*x = GetAllTeacherResponse{}
	mi := &file_protos_teacher_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllTeacherResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTeacherResponse) ProtoMessage() {}

func (x *GetAllTeacherResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_teacher_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTeacherResponse.ProtoReflect.Descriptor instead.
func (*GetAllTeacherResponse) Descriptor() ([]byte, []int) {
	return file_protos_teacher_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllTeacherResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllTeacherResponse) GetTeachers() []*Teacher {
	if x != nil {
		return x.Teachers
	}
	return nil
}

type UpdateTeacherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UpdateTeacherRequest) Reset() {
	*x = UpdateTeacherRequest{}
	mi := &file_protos_teacher_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTeacherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTeacherRequest) ProtoMessage() {}

func (x *UpdateTeacherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_teacher_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTeacherRequest.ProtoReflect.Descriptor instead.
func (*UpdateTeacherRequest) Descriptor() ([]byte, []int) {
	return file_protos_teacher_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateTeacherRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateTeacherRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UpdateTeacherResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UpdateTeacherResponse) Reset() {
	*x = UpdateTeacherResponse{}
	mi := &file_protos_teacher_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTeacherResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTeacherResponse) ProtoMessage() {}

func (x *UpdateTeacherResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_teacher_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTeacherResponse.ProtoReflect.Descriptor instead.
func (*UpdateTeacherResponse) Descriptor() ([]byte, []int) {
	return file_protos_teacher_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateTeacherResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Delete
type DeleteTeacherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteTeacherRequest) Reset() {
	*x = DeleteTeacherRequest{}
	mi := &file_protos_teacher_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTeacherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTeacherRequest) ProtoMessage() {}

func (x *DeleteTeacherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_teacher_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTeacherRequest.ProtoReflect.Descriptor instead.
func (*DeleteTeacherRequest) Descriptor() ([]byte, []int) {
	return file_protos_teacher_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteTeacherRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteTeacherResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteTeacherResponse) Reset() {
	*x = DeleteTeacherResponse{}
	mi := &file_protos_teacher_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTeacherResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTeacherResponse) ProtoMessage() {}

func (x *DeleteTeacherResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_teacher_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTeacherResponse.ProtoReflect.Descriptor instead.
func (*DeleteTeacherResponse) Descriptor() ([]byte, []int) {
	return file_protos_teacher_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteTeacherResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_protos_teacher_proto protoreflect.FileDescriptor

var file_protos_teacher_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x73, 0x0a, 0x07, 0x54, 0x65, 0x61,
	0x63, 0x68, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x54, 0x65, 0x61, 0x63, 0x68,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49,
	0x44, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x44,
	0x12, 0x1c, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x09, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x22, 0xb8,
	0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x20, 0x0a,
	0x0b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x65, 0x61, 0x63, 0x68,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x54, 0x65,
	0x61, 0x63, 0x68, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x31, 0x0a, 0x15, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x5c, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x29, 0x0a, 0x08, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72,
	0x52, 0x08, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x73, 0x22, 0x3a, 0x0a, 0x14, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x31, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x26, 0x0a, 0x14, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x31, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x61, 0x63, 0x68,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x32, 0xb7, 0x02, 0x0a, 0x0e, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4e, 0x65, 0x77, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x12, 0x1a, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x65,
	0x61, 0x63, 0x68, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1b, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x65, 0x61, 0x63, 0x68,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0d, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x12, 0x1a, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65,
	0x61, 0x63, 0x68, 0x65, 0x72, 0x12, 0x1a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54,
	0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x15,
	0x5a, 0x13, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_teacher_proto_rawDescOnce sync.Once
	file_protos_teacher_proto_rawDescData = file_protos_teacher_proto_rawDesc
)

func file_protos_teacher_proto_rawDescGZIP() []byte {
	file_protos_teacher_proto_rawDescOnce.Do(func() {
		file_protos_teacher_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_teacher_proto_rawDescData)
	})
	return file_protos_teacher_proto_rawDescData
}

var file_protos_teacher_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_protos_teacher_proto_goTypes = []any{
	(*Teacher)(nil),               // 0: auth.Teacher
	(*CreateTeacherRequest)(nil),  // 1: auth.CreateTeacherRequest
	(*CreateTeacherResponse)(nil), // 2: auth.CreateTeacherResponse
	(*GetAllTeacherResponse)(nil), // 3: auth.GetAllTeacherResponse
	(*UpdateTeacherRequest)(nil),  // 4: auth.UpdateTeacherRequest
	(*UpdateTeacherResponse)(nil), // 5: auth.UpdateTeacherResponse
	(*DeleteTeacherRequest)(nil),  // 6: auth.DeleteTeacherRequest
	(*DeleteTeacherResponse)(nil), // 7: auth.DeleteTeacherResponse
	(*emptypb.Empty)(nil),         // 8: google.protobuf.Empty
}
var file_protos_teacher_proto_depIdxs = []int32{
	0, // 0: auth.GetAllTeacherResponse.teachers:type_name -> auth.Teacher
	1, // 1: auth.TeacherService.CreateNewTeacher:input_type -> auth.CreateTeacherRequest
	8, // 2: auth.TeacherService.GetAllTeacher:input_type -> google.protobuf.Empty
	4, // 3: auth.TeacherService.UpdateTeacher:input_type -> auth.UpdateTeacherRequest
	6, // 4: auth.TeacherService.DeleteTeacher:input_type -> auth.DeleteTeacherRequest
	2, // 5: auth.TeacherService.CreateNewTeacher:output_type -> auth.CreateTeacherResponse
	3, // 6: auth.TeacherService.GetAllTeacher:output_type -> auth.GetAllTeacherResponse
	5, // 7: auth.TeacherService.UpdateTeacher:output_type -> auth.UpdateTeacherResponse
	7, // 8: auth.TeacherService.DeleteTeacher:output_type -> auth.DeleteTeacherResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_teacher_proto_init() }
func file_protos_teacher_proto_init() {
	if File_protos_teacher_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_teacher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_teacher_proto_goTypes,
		DependencyIndexes: file_protos_teacher_proto_depIdxs,
		MessageInfos:      file_protos_teacher_proto_msgTypes,
	}.Build()
	File_protos_teacher_proto = out.File
	file_protos_teacher_proto_rawDesc = nil
	file_protos_teacher_proto_goTypes = nil
	file_protos_teacher_proto_depIdxs = nil
}
