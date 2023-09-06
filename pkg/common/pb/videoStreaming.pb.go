// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.12
// source: pkg/common/proto/videoStreaming.proto

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

type UploadVideoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Data     []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UploadVideoRequest) Reset() {
	*x = UploadVideoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_common_proto_videoStreaming_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadVideoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadVideoRequest) ProtoMessage() {}

func (x *UploadVideoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_common_proto_videoStreaming_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadVideoRequest.ProtoReflect.Descriptor instead.
func (*UploadVideoRequest) Descriptor() ([]byte, []int) {
	return file_pkg_common_proto_videoStreaming_proto_rawDescGZIP(), []int{0}
}

func (x *UploadVideoRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *UploadVideoRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type UploadVideoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	VideoId string `protobuf:"bytes,3,opt,name=videoId,proto3" json:"videoId,omitempty"`
}

func (x *UploadVideoResponse) Reset() {
	*x = UploadVideoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_common_proto_videoStreaming_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadVideoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadVideoResponse) ProtoMessage() {}

func (x *UploadVideoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_common_proto_videoStreaming_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadVideoResponse.ProtoReflect.Descriptor instead.
func (*UploadVideoResponse) Descriptor() ([]byte, []int) {
	return file_pkg_common_proto_videoStreaming_proto_rawDescGZIP(), []int{1}
}

func (x *UploadVideoResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *UploadVideoResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UploadVideoResponse) GetVideoId() string {
	if x != nil {
		return x.VideoId
	}
	return ""
}

type StreamVideoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoId string `protobuf:"bytes,1,opt,name=videoId,proto3" json:"videoId,omitempty"`
}

func (x *StreamVideoRequest) Reset() {
	*x = StreamVideoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_common_proto_videoStreaming_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamVideoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamVideoRequest) ProtoMessage() {}

func (x *StreamVideoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_common_proto_videoStreaming_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamVideoRequest.ProtoReflect.Descriptor instead.
func (*StreamVideoRequest) Descriptor() ([]byte, []int) {
	return file_pkg_common_proto_videoStreaming_proto_rawDescGZIP(), []int{2}
}

func (x *StreamVideoRequest) GetVideoId() string {
	if x != nil {
		return x.VideoId
	}
	return ""
}

type StreamVideoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoChunk []byte `protobuf:"bytes,1,opt,name=video_chunk,json=videoChunk,proto3" json:"video_chunk,omitempty"`
}

func (x *StreamVideoResponse) Reset() {
	*x = StreamVideoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_common_proto_videoStreaming_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamVideoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamVideoResponse) ProtoMessage() {}

func (x *StreamVideoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_common_proto_videoStreaming_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamVideoResponse.ProtoReflect.Descriptor instead.
func (*StreamVideoResponse) Descriptor() ([]byte, []int) {
	return file_pkg_common_proto_videoStreaming_proto_rawDescGZIP(), []int{3}
}

func (x *StreamVideoResponse) GetVideoChunk() []byte {
	if x != nil {
		return x.VideoChunk
	}
	return nil
}

type FindAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FindAllRequest) Reset() {
	*x = FindAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_common_proto_videoStreaming_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllRequest) ProtoMessage() {}

func (x *FindAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_common_proto_videoStreaming_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllRequest.ProtoReflect.Descriptor instead.
func (*FindAllRequest) Descriptor() ([]byte, []int) {
	return file_pkg_common_proto_videoStreaming_proto_rawDescGZIP(), []int{4}
}

type VideoID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoId string `protobuf:"bytes,1,opt,name=VideoId,proto3" json:"VideoId,omitempty"`
}

func (x *VideoID) Reset() {
	*x = VideoID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_common_proto_videoStreaming_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoID) ProtoMessage() {}

func (x *VideoID) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_common_proto_videoStreaming_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoID.ProtoReflect.Descriptor instead.
func (*VideoID) Descriptor() ([]byte, []int) {
	return file_pkg_common_proto_videoStreaming_proto_rawDescGZIP(), []int{5}
}

func (x *VideoID) GetVideoId() string {
	if x != nil {
		return x.VideoId
	}
	return ""
}

type FindAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32      `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Videos []*VideoID `protobuf:"bytes,2,rep,name=videos,proto3" json:"videos,omitempty"`
}

func (x *FindAllResponse) Reset() {
	*x = FindAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_common_proto_videoStreaming_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllResponse) ProtoMessage() {}

func (x *FindAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_common_proto_videoStreaming_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllResponse.ProtoReflect.Descriptor instead.
func (*FindAllResponse) Descriptor() ([]byte, []int) {
	return file_pkg_common_proto_videoStreaming_proto_rawDescGZIP(), []int{6}
}

func (x *FindAllResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *FindAllResponse) GetVideos() []*VideoID {
	if x != nil {
		return x.Videos
	}
	return nil
}

var File_pkg_common_proto_videoStreaming_proto protoreflect.FileDescriptor

var file_pkg_common_proto_videoStreaming_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x44, 0x0a, 0x12, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x61, 0x0a, 0x13, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x49, 0x64, 0x22, 0x2e, 0x0a, 0x12, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x49, 0x64, 0x22, 0x36, 0x0a, 0x13, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x5f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0a, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x22, 0x10, 0x0a, 0x0e,
	0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x23,
	0x0a, 0x07, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x49, 0x64, 0x22, 0x4e, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23,
	0x0a, 0x06, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x70, 0x62, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x44, 0x52, 0x06, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x73, 0x32, 0xd1, 0x01, 0x0a, 0x0c, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x0b, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62,
	0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x42, 0x0a, 0x0b, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x39, 0x0a, 0x0c,
	0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x12, 0x2e, 0x70,
	0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pkg_common_proto_videoStreaming_proto_rawDescOnce sync.Once
	file_pkg_common_proto_videoStreaming_proto_rawDescData = file_pkg_common_proto_videoStreaming_proto_rawDesc
)

func file_pkg_common_proto_videoStreaming_proto_rawDescGZIP() []byte {
	file_pkg_common_proto_videoStreaming_proto_rawDescOnce.Do(func() {
		file_pkg_common_proto_videoStreaming_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_common_proto_videoStreaming_proto_rawDescData)
	})
	return file_pkg_common_proto_videoStreaming_proto_rawDescData
}

var file_pkg_common_proto_videoStreaming_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pkg_common_proto_videoStreaming_proto_goTypes = []interface{}{
	(*UploadVideoRequest)(nil),  // 0: pb.UploadVideoRequest
	(*UploadVideoResponse)(nil), // 1: pb.UploadVideoResponse
	(*StreamVideoRequest)(nil),  // 2: pb.StreamVideoRequest
	(*StreamVideoResponse)(nil), // 3: pb.StreamVideoResponse
	(*FindAllRequest)(nil),      // 4: pb.FindAllRequest
	(*VideoID)(nil),             // 5: pb.VideoID
	(*FindAllResponse)(nil),     // 6: pb.FindAllResponse
}
var file_pkg_common_proto_videoStreaming_proto_depIdxs = []int32{
	5, // 0: pb.FindAllResponse.videos:type_name -> pb.VideoID
	0, // 1: pb.VideoService.UploadVideo:input_type -> pb.UploadVideoRequest
	2, // 2: pb.VideoService.StreamVideo:input_type -> pb.StreamVideoRequest
	4, // 3: pb.VideoService.FindAllVideo:input_type -> pb.FindAllRequest
	1, // 4: pb.VideoService.UploadVideo:output_type -> pb.UploadVideoResponse
	3, // 5: pb.VideoService.StreamVideo:output_type -> pb.StreamVideoResponse
	6, // 6: pb.VideoService.FindAllVideo:output_type -> pb.FindAllResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_common_proto_videoStreaming_proto_init() }
func file_pkg_common_proto_videoStreaming_proto_init() {
	if File_pkg_common_proto_videoStreaming_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_common_proto_videoStreaming_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadVideoRequest); i {
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
		file_pkg_common_proto_videoStreaming_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadVideoResponse); i {
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
		file_pkg_common_proto_videoStreaming_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamVideoRequest); i {
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
		file_pkg_common_proto_videoStreaming_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamVideoResponse); i {
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
		file_pkg_common_proto_videoStreaming_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllRequest); i {
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
		file_pkg_common_proto_videoStreaming_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoID); i {
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
		file_pkg_common_proto_videoStreaming_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllResponse); i {
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
			RawDescriptor: file_pkg_common_proto_videoStreaming_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_common_proto_videoStreaming_proto_goTypes,
		DependencyIndexes: file_pkg_common_proto_videoStreaming_proto_depIdxs,
		MessageInfos:      file_pkg_common_proto_videoStreaming_proto_msgTypes,
	}.Build()
	File_pkg_common_proto_videoStreaming_proto = out.File
	file_pkg_common_proto_videoStreaming_proto_rawDesc = nil
	file_pkg_common_proto_videoStreaming_proto_goTypes = nil
	file_pkg_common_proto_videoStreaming_proto_depIdxs = nil
}