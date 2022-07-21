// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.15.8
// source: proto/user.proto

package rpc

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SetDriverRatingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rating int32 `protobuf:"varint,1,opt,name=rating,proto3" json:"rating,omitempty"`
}

func (x *SetDriverRatingRequest) Reset() {
	*x = SetDriverRatingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetDriverRatingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetDriverRatingRequest) ProtoMessage() {}

func (x *SetDriverRatingRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use SetDriverRatingRequest.ProtoReflect.Descriptor instead.
func (*SetDriverRatingRequest) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{0}
}

func (x *SetDriverRatingRequest) GetRating() int32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

type SetUserRatingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rating int32 `protobuf:"varint,1,opt,name=rating,proto3" json:"rating,omitempty"`
}

func (x *SetUserRatingRequest) Reset() {
	*x = SetUserRatingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetUserRatingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetUserRatingRequest) ProtoMessage() {}

func (x *SetUserRatingRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use SetUserRatingRequest.ProtoReflect.Descriptor instead.
func (*SetUserRatingRequest) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{1}
}

func (x *SetUserRatingRequest) GetRating() int32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

type GetTaxiForUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserUuid    string  `protobuf:"bytes,1,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	Origin      string  `protobuf:"bytes,2,opt,name=origin,proto3" json:"origin,omitempty"`
	Destination string  `protobuf:"bytes,3,opt,name=destination,proto3" json:"destination,omitempty"`
	TaxiType    string  `protobuf:"bytes,4,opt,name=taxi_type,json=taxiType,proto3" json:"taxi_type,omitempty"`
	UserRating  float32 `protobuf:"fixed32,5,opt,name=user_rating,json=userRating,proto3" json:"user_rating,omitempty"`
}

func (x *GetTaxiForUserRequest) Reset() {
	*x = GetTaxiForUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTaxiForUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaxiForUserRequest) ProtoMessage() {}

func (x *GetTaxiForUserRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetTaxiForUserRequest.ProtoReflect.Descriptor instead.
func (*GetTaxiForUserRequest) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{2}
}

func (x *GetTaxiForUserRequest) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

func (x *GetTaxiForUserRequest) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *GetTaxiForUserRequest) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

func (x *GetTaxiForUserRequest) GetTaxiType() string {
	if x != nil {
		return x.TaxiType
	}
	return ""
}

func (x *GetTaxiForUserRequest) GetUserRating() float32 {
	if x != nil {
		return x.UserRating
	}
	return 0
}

type GetTaxiForUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DriverUuid   string  `protobuf:"bytes,1,opt,name=driver_uuid,json=driverUuid,proto3" json:"driver_uuid,omitempty"`
	DriverRating float32 `protobuf:"fixed32,2,opt,name=driver_rating,json=driverRating,proto3" json:"driver_rating,omitempty"`
}

func (x *GetTaxiForUserResponse) Reset() {
	*x = GetTaxiForUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTaxiForUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaxiForUserResponse) ProtoMessage() {}

func (x *GetTaxiForUserResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetTaxiForUserResponse.ProtoReflect.Descriptor instead.
func (*GetTaxiForUserResponse) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{3}
}

func (x *GetTaxiForUserResponse) GetDriverUuid() string {
	if x != nil {
		return x.DriverUuid
	}
	return ""
}

func (x *GetTaxiForUserResponse) GetDriverRating() float32 {
	if x != nil {
		return x.DriverRating
	}
	return 0
}

type GetOrderForDriverRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DriverUuid   string  `protobuf:"bytes,1,opt,name=driver_uuid,json=driverUuid,proto3" json:"driver_uuid,omitempty"`
	TaxiType     string  `protobuf:"bytes,2,opt,name=taxi_type,json=taxiType,proto3" json:"taxi_type,omitempty"`
	DriverRating float32 `protobuf:"fixed32,3,opt,name=driver_rating,json=driverRating,proto3" json:"driver_rating,omitempty"`
}

func (x *GetOrderForDriverRequest) Reset() {
	*x = GetOrderForDriverRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderForDriverRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderForDriverRequest) ProtoMessage() {}

func (x *GetOrderForDriverRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetOrderForDriverRequest.ProtoReflect.Descriptor instead.
func (*GetOrderForDriverRequest) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{4}
}

func (x *GetOrderForDriverRequest) GetDriverUuid() string {
	if x != nil {
		return x.DriverUuid
	}
	return ""
}

func (x *GetOrderForDriverRequest) GetTaxiType() string {
	if x != nil {
		return x.TaxiType
	}
	return ""
}

func (x *GetOrderForDriverRequest) GetDriverRating() float32 {
	if x != nil {
		return x.DriverRating
	}
	return 0
}

type GetOrderForDriverResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserUuid    string  `protobuf:"bytes,1,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	Origin      string  `protobuf:"bytes,2,opt,name=origin,proto3" json:"origin,omitempty"`
	Destination string  `protobuf:"bytes,3,opt,name=destination,proto3" json:"destination,omitempty"`
	UserRating  float32 `protobuf:"fixed32,4,opt,name=user_rating,json=userRating,proto3" json:"user_rating,omitempty"`
}

func (x *GetOrderForDriverResponse) Reset() {
	*x = GetOrderForDriverResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderForDriverResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderForDriverResponse) ProtoMessage() {}

func (x *GetOrderForDriverResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetOrderForDriverResponse.ProtoReflect.Descriptor instead.
func (*GetOrderForDriverResponse) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{5}
}

func (x *GetOrderForDriverResponse) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

func (x *GetOrderForDriverResponse) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *GetOrderForDriverResponse) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

func (x *GetOrderForDriverResponse) GetUserRating() float32 {
	if x != nil {
		return x.UserRating
	}
	return 0
}

type GetOrderHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *GetOrderHistoryRequest) Reset() {
	*x = GetOrderHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderHistoryRequest) ProtoMessage() {}

func (x *GetOrderHistoryRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetOrderHistoryRequest.ProtoReflect.Descriptor instead.
func (*GetOrderHistoryRequest) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{6}
}

func (x *GetOrderHistoryRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type GetOrderHistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders []*Order `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
}

func (x *GetOrderHistoryResponse) Reset() {
	*x = GetOrderHistoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderHistoryResponse) ProtoMessage() {}

func (x *GetOrderHistoryResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetOrderHistoryResponse.ProtoReflect.Descriptor instead.
func (*GetOrderHistoryResponse) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{7}
}

func (x *GetOrderHistoryResponse) GetOrders() []*Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserUuid    string `protobuf:"bytes,1,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	DriverUuid  string `protobuf:"bytes,2,opt,name=driver_uuid,json=driverUuid,proto3" json:"driver_uuid,omitempty"`
	Origin      string `protobuf:"bytes,3,opt,name=origin,proto3" json:"origin,omitempty"`
	Destination string `protobuf:"bytes,4,opt,name=destination,proto3" json:"destination,omitempty"`
	TaxiType    string `protobuf:"bytes,5,opt,name=taxi_type,json=taxiType,proto3" json:"taxi_type,omitempty"`
	Date        string `protobuf:"bytes,6,opt,name=date,proto3" json:"date,omitempty"`
	Duration    string `protobuf:"bytes,7,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{8}
}

func (x *Order) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

func (x *Order) GetDriverUuid() string {
	if x != nil {
		return x.DriverUuid
	}
	return ""
}

func (x *Order) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *Order) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

func (x *Order) GetTaxiType() string {
	if x != nil {
		return x.TaxiType
	}
	return ""
}

func (x *Order) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *Order) GetDuration() string {
	if x != nil {
		return x.Duration
	}
	return ""
}

type EmptyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyResponse) Reset() {
	*x = EmptyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResponse) ProtoMessage() {}

func (x *EmptyResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use EmptyResponse.ProtoReflect.Descriptor instead.
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{9}
}

var File_proto_user_proto protoreflect.FileDescriptor

var file_proto_user_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x30, 0x0a, 0x16, 0x53, 0x65, 0x74,
	0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x2e, 0x0a, 0x14, 0x53,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0xac, 0x01, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x54, 0x61, 0x78, 0x69, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x55, 0x75,
	0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09,
	0x74, 0x61, 0x78, 0x69, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x74, 0x61, 0x78, 0x69, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a,
	0x75, 0x73, 0x65, 0x72, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x5e, 0x0a, 0x16, 0x47, 0x65,
	0x74, 0x54, 0x61, 0x78, 0x69, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x75,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x55, 0x75, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f,
	0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x64, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x7d, 0x0a, 0x18, 0x47, 0x65,
	0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61, 0x78, 0x69, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x78, 0x69,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x72,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x64, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x93, 0x01, 0x0a, 0x19, 0x47, 0x65,
	0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x55, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f,
	0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22,
	0x2c, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x3f, 0x0a,
	0x17, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x22, 0xcc,
	0x01, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x55, 0x75, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f,
	0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61, 0x78, 0x69, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x78, 0x69, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x0f, 0x0a,
	0x0d, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x53,
	0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a,
	0x0d, 0x53, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x32, 0x9d, 0x03, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x53, 0x65, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44,
	0x0a, 0x0d, 0x53, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12,
	0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x78, 0x69, 0x46,
	0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47,
	0x65, 0x74, 0x54, 0x61, 0x78, 0x69, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x61, 0x78, 0x69, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x46, 0x6f, 0x72, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x44, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x44,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x52, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x62, 0x61, 0x64, 0x66, 0x61, 0x6e, 0x2f, 0x69, 0x6e, 0x6e, 0x6f, 0x2d, 0x74, 0x61,
	0x78, 0x69, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_proto_user_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_user_proto_goTypes = []interface{}{
	(*SetDriverRatingRequest)(nil),    // 0: proto.SetDriverRatingRequest
	(*SetUserRatingRequest)(nil),      // 1: proto.SetUserRatingRequest
	(*GetTaxiForUserRequest)(nil),     // 2: proto.GetTaxiForUserRequest
	(*GetTaxiForUserResponse)(nil),    // 3: proto.GetTaxiForUserResponse
	(*GetOrderForDriverRequest)(nil),  // 4: proto.GetOrderForDriverRequest
	(*GetOrderForDriverResponse)(nil), // 5: proto.GetOrderForDriverResponse
	(*GetOrderHistoryRequest)(nil),    // 6: proto.GetOrderHistoryRequest
	(*GetOrderHistoryResponse)(nil),   // 7: proto.GetOrderHistoryResponse
	(*Order)(nil),                     // 8: proto.Order
	(*EmptyResponse)(nil),             // 9: proto.EmptyResponse
}
var file_proto_user_proto_depIdxs = []int32{
	8, // 0: proto.GetOrderHistoryResponse.orders:type_name -> proto.Order
	1, // 1: proto.UserService.SetUserRating:input_type -> proto.SetUserRatingRequest
	0, // 2: proto.OrderService.SetDriverRating:input_type -> proto.SetDriverRatingRequest
	1, // 3: proto.OrderService.SetUserRating:input_type -> proto.SetUserRatingRequest
	2, // 4: proto.OrderService.GetTaxiForUser:input_type -> proto.GetTaxiForUserRequest
	4, // 5: proto.OrderService.GetOrderForDriver:input_type -> proto.GetOrderForDriverRequest
	6, // 6: proto.OrderService.GetOrderHistory:input_type -> proto.GetOrderHistoryRequest
	9, // 7: proto.UserService.SetUserRating:output_type -> proto.EmptyResponse
	9, // 8: proto.OrderService.SetDriverRating:output_type -> proto.EmptyResponse
	9, // 9: proto.OrderService.SetUserRating:output_type -> proto.EmptyResponse
	3, // 10: proto.OrderService.GetTaxiForUser:output_type -> proto.GetTaxiForUserResponse
	5, // 11: proto.OrderService.GetOrderForDriver:output_type -> proto.GetOrderForDriverResponse
	7, // 12: proto.OrderService.GetOrderHistory:output_type -> proto.GetOrderHistoryResponse
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_user_proto_init() }
func file_proto_user_proto_init() {
	if File_proto_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetDriverRatingRequest); i {
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
		file_proto_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetUserRatingRequest); i {
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
		file_proto_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTaxiForUserRequest); i {
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
		file_proto_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTaxiForUserResponse); i {
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
		file_proto_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderForDriverRequest); i {
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
		file_proto_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderForDriverResponse); i {
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
		file_proto_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderHistoryRequest); i {
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
		file_proto_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderHistoryResponse); i {
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
		file_proto_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_proto_user_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyResponse); i {
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
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   2,
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
