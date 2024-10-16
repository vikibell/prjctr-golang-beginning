// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.1
// source: api.proto

package gocourse17

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

type Rating int32

const (
	Rating_UNDEFINED Rating = 0
	Rating_POOR      Rating = 1
	Rating_FAIR      Rating = 2
	Rating_GOOD      Rating = 3
	Rating_GREAT     Rating = 4
	Rating_EXCELLENT Rating = 5
)

// Enum value maps for Rating.
var (
	Rating_name = map[int32]string{
		0: "UNDEFINED",
		1: "POOR",
		2: "FAIR",
		3: "GOOD",
		4: "GREAT",
		5: "EXCELLENT",
	}
	Rating_value = map[string]int32{
		"UNDEFINED": 0,
		"POOR":      1,
		"FAIR":      2,
		"GOOD":      3,
		"GREAT":     4,
		"EXCELLENT": 5,
	}
)

func (x Rating) Enum() *Rating {
	p := new(Rating)
	*p = x
	return p
}

func (x Rating) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Rating) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_enumTypes[0].Descriptor()
}

func (Rating) Type() protoreflect.EnumType {
	return &file_api_proto_enumTypes[0]
}

func (x Rating) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Rating.Descriptor instead.
func (Rating) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

type GetHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DriverId int32 `protobuf:"varint,1,opt,name=driver_id,json=driverId,proto3" json:"driver_id,omitempty"`
}

func (x *GetHistoryRequest) Reset() {
	*x = GetHistoryRequest{}
	mi := &file_api_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHistoryRequest) ProtoMessage() {}

func (x *GetHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHistoryRequest.ProtoReflect.Descriptor instead.
func (*GetHistoryRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *GetHistoryRequest) GetDriverId() int32 {
	if x != nil {
		return x.DriverId
	}
	return 0
}

type GetHistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reviews []*ReviewData `protobuf:"bytes,1,rep,name=reviews,proto3" json:"reviews,omitempty"`
}

func (x *GetHistoryResponse) Reset() {
	*x = GetHistoryResponse{}
	mi := &file_api_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHistoryResponse) ProtoMessage() {}

func (x *GetHistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHistoryResponse.ProtoReflect.Descriptor instead.
func (*GetHistoryResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *GetHistoryResponse) GetReviews() []*ReviewData {
	if x != nil {
		return x.Reviews
	}
	return nil
}

type SendReviewRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DriverId int32       `protobuf:"varint,1,opt,name=driver_id,json=driverId,proto3" json:"driver_id,omitempty"`
	Review   *ReviewData `protobuf:"bytes,2,opt,name=review,proto3" json:"review,omitempty"`
}

func (x *SendReviewRequest) Reset() {
	*x = SendReviewRequest{}
	mi := &file_api_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendReviewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendReviewRequest) ProtoMessage() {}

func (x *SendReviewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendReviewRequest.ProtoReflect.Descriptor instead.
func (*SendReviewRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *SendReviewRequest) GetDriverId() int32 {
	if x != nil {
		return x.DriverId
	}
	return 0
}

func (x *SendReviewRequest) GetReview() *ReviewData {
	if x != nil {
		return x.Review
	}
	return nil
}

type SendReviewResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SendReviewResponse) Reset() {
	*x = SendReviewResponse{}
	mi := &file_api_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendReviewResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendReviewResponse) ProtoMessage() {}

func (x *SendReviewResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendReviewResponse.ProtoReflect.Descriptor instead.
func (*SendReviewResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *SendReviewResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ReviewData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CargoState       Rating `protobuf:"varint,1,opt,name=cargo_state,json=cargoState,proto3,enum=review.Rating" json:"cargo_state,omitempty"`
	ServiceQuality   Rating `protobuf:"varint,2,opt,name=service_quality,json=serviceQuality,proto3,enum=review.Rating" json:"service_quality,omitempty"`
	FulfillmentSpeed Rating `protobuf:"varint,3,opt,name=fulfillment_speed,json=fulfillmentSpeed,proto3,enum=review.Rating" json:"fulfillment_speed,omitempty"`
}

func (x *ReviewData) Reset() {
	*x = ReviewData{}
	mi := &file_api_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReviewData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReviewData) ProtoMessage() {}

func (x *ReviewData) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReviewData.ProtoReflect.Descriptor instead.
func (*ReviewData) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

func (x *ReviewData) GetCargoState() Rating {
	if x != nil {
		return x.CargoState
	}
	return Rating_UNDEFINED
}

func (x *ReviewData) GetServiceQuality() Rating {
	if x != nil {
		return x.ServiceQuality
	}
	return Rating_UNDEFINED
}

func (x *ReviewData) GetFulfillmentSpeed() Rating {
	if x != nil {
		return x.FulfillmentSpeed
	}
	return Rating_UNDEFINED
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x22, 0x30, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x49, 0x64, 0x22, 0x42, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x07, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x22, 0x5c, 0x0a, 0x11, 0x53, 0x65, 0x6e,
	0x64, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x06, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x06, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x22, 0x2e, 0x0a, 0x12, 0x53, 0x65, 0x6e, 0x64, 0x52,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xb3, 0x01, 0x0a, 0x0a, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2f, 0x0a, 0x0b, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x2e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x0a, 0x63, 0x61, 0x72,
	0x67, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x37, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0e, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x52, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79,
	0x12, 0x3b, 0x0a, 0x11, 0x66, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x2e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x10, 0x66, 0x75, 0x6c,
	0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x70, 0x65, 0x65, 0x64, 0x2a, 0x4f, 0x0a,
	0x06, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e, 0x44, 0x45, 0x46,
	0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x4f, 0x4f, 0x52, 0x10, 0x01,
	0x12, 0x08, 0x0a, 0x04, 0x46, 0x41, 0x49, 0x52, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x47, 0x4f,
	0x4f, 0x44, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x52, 0x45, 0x41, 0x54, 0x10, 0x04, 0x12,
	0x0d, 0x0a, 0x09, 0x45, 0x58, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x4e, 0x54, 0x10, 0x05, 0x32, 0x96,
	0x01, 0x0a, 0x06, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x45, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x19, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x2e, 0x47, 0x65, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x47, 0x65, 0x74, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x45, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x19,
	0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x69, 0x6b, 0x69, 0x62, 0x65, 0x6c, 0x6c, 0x2f, 0x70,
	0x72, 0x6a, 0x63, 0x74, 0x72, 0x2d, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2d, 0x62, 0x65, 0x67,
	0x69, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x31,
	0x37, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_proto_goTypes = []any{
	(Rating)(0),                // 0: review.Rating
	(*GetHistoryRequest)(nil),  // 1: review.GetHistoryRequest
	(*GetHistoryResponse)(nil), // 2: review.GetHistoryResponse
	(*SendReviewRequest)(nil),  // 3: review.SendReviewRequest
	(*SendReviewResponse)(nil), // 4: review.SendReviewResponse
	(*ReviewData)(nil),         // 5: review.ReviewData
}
var file_api_proto_depIdxs = []int32{
	5, // 0: review.GetHistoryResponse.reviews:type_name -> review.ReviewData
	5, // 1: review.SendReviewRequest.review:type_name -> review.ReviewData
	0, // 2: review.ReviewData.cargo_state:type_name -> review.Rating
	0, // 3: review.ReviewData.service_quality:type_name -> review.Rating
	0, // 4: review.ReviewData.fulfillment_speed:type_name -> review.Rating
	1, // 5: review.Review.GetHistory:input_type -> review.GetHistoryRequest
	3, // 6: review.Review.SendReview:input_type -> review.SendReviewRequest
	2, // 7: review.Review.GetHistory:output_type -> review.GetHistoryResponse
	4, // 8: review.Review.SendReview:output_type -> review.SendReviewResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		EnumInfos:         file_api_proto_enumTypes,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
