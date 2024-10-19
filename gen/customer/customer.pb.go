// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: proto/customer/customer.proto

package __

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

type RequestTaxi struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId         uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PickupLocation string `protobuf:"bytes,2,opt,name=pickup_location,json=pickupLocation,proto3" json:"pickup_location,omitempty"`
	DropLocation   string `protobuf:"bytes,3,opt,name=drop_location,json=dropLocation,proto3" json:"drop_location,omitempty"`
}

func (x *RequestTaxi) Reset() {
	*x = RequestTaxi{}
	mi := &file_proto_customer_customer_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequestTaxi) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestTaxi) ProtoMessage() {}

func (x *RequestTaxi) ProtoReflect() protoreflect.Message {
	mi := &file_proto_customer_customer_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestTaxi.ProtoReflect.Descriptor instead.
func (*RequestTaxi) Descriptor() ([]byte, []int) {
	return file_proto_customer_customer_proto_rawDescGZIP(), []int{0}
}

func (x *RequestTaxi) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *RequestTaxi) GetPickupLocation() string {
	if x != nil {
		return x.PickupLocation
	}
	return ""
}

func (x *RequestTaxi) GetDropLocation() string {
	if x != nil {
		return x.DropLocation
	}
	return ""
}

type ResponseTaxi struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookingId uint64 `protobuf:"varint,1,opt,name=booking_id,json=bookingId,proto3" json:"booking_id,omitempty"`
	Status    string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ResponseTaxi) Reset() {
	*x = ResponseTaxi{}
	mi := &file_proto_customer_customer_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResponseTaxi) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseTaxi) ProtoMessage() {}

func (x *ResponseTaxi) ProtoReflect() protoreflect.Message {
	mi := &file_proto_customer_customer_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseTaxi.ProtoReflect.Descriptor instead.
func (*ResponseTaxi) Descriptor() ([]byte, []int) {
	return file_proto_customer_customer_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseTaxi) GetBookingId() uint64 {
	if x != nil {
		return x.BookingId
	}
	return 0
}

func (x *ResponseTaxi) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_proto_customer_customer_proto protoreflect.FileDescriptor

var file_proto_customer_customer_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x22, 0x74, 0x0a, 0x0b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x54, 0x61, 0x78, 0x69, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x69, 0x63, 0x6b,
	0x75, 0x70, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x72,
	0x6f, 0x70, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x64, 0x72, 0x6f, 0x70, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x45, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x61, 0x78, 0x69, 0x12,
	0x1d, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x4f, 0x0a, 0x0f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x52, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x54, 0x61, 0x78, 0x69, 0x12, 0x15, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x61, 0x78, 0x69, 0x1a,
	0x16, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x54, 0x61, 0x78, 0x69, 0x42, 0x03, 0x5a, 0x01, 0x2f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_customer_customer_proto_rawDescOnce sync.Once
	file_proto_customer_customer_proto_rawDescData = file_proto_customer_customer_proto_rawDesc
)

func file_proto_customer_customer_proto_rawDescGZIP() []byte {
	file_proto_customer_customer_proto_rawDescOnce.Do(func() {
		file_proto_customer_customer_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_customer_customer_proto_rawDescData)
	})
	return file_proto_customer_customer_proto_rawDescData
}

var file_proto_customer_customer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_customer_customer_proto_goTypes = []any{
	(*RequestTaxi)(nil),  // 0: customer.RequestTaxi
	(*ResponseTaxi)(nil), // 1: customer.ResponseTaxi
}
var file_proto_customer_customer_proto_depIdxs = []int32{
	0, // 0: customer.CustomerService.ReserveTaxi:input_type -> customer.RequestTaxi
	1, // 1: customer.CustomerService.ReserveTaxi:output_type -> customer.ResponseTaxi
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_customer_customer_proto_init() }
func file_proto_customer_customer_proto_init() {
	if File_proto_customer_customer_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_customer_customer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_customer_customer_proto_goTypes,
		DependencyIndexes: file_proto_customer_customer_proto_depIdxs,
		MessageInfos:      file_proto_customer_customer_proto_msgTypes,
	}.Build()
	File_proto_customer_customer_proto = out.File
	file_proto_customer_customer_proto_rawDesc = nil
	file_proto_customer_customer_proto_goTypes = nil
	file_proto_customer_customer_proto_depIdxs = nil
}
