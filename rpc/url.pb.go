// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.11.2
// source: rpc/url.proto

package rpc

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type UrlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url      string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Interval int32  `protobuf:"varint,2,opt,name=interval,proto3" json:"interval,omitempty"`
}

func (x *UrlRequest) Reset() {
	*x = UrlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_url_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UrlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UrlRequest) ProtoMessage() {}

func (x *UrlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_url_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UrlRequest.ProtoReflect.Descriptor instead.
func (*UrlRequest) Descriptor() ([]byte, []int) {
	return file_rpc_url_proto_rawDescGZIP(), []int{0}
}

func (x *UrlRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *UrlRequest) GetInterval() int32 {
	if x != nil {
		return x.Interval
	}
	return 0
}

type UrlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *UrlResponse) Reset() {
	*x = UrlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_url_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UrlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UrlResponse) ProtoMessage() {}

func (x *UrlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_url_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UrlResponse.ProtoReflect.Descriptor instead.
func (*UrlResponse) Descriptor() ([]byte, []int) {
	return file_rpc_url_proto_rawDescGZIP(), []int{1}
}

func (x *UrlResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type UrlListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Urls []string `protobuf:"bytes,1,rep,name=urls,proto3" json:"urls,omitempty"`
}

func (x *UrlListResponse) Reset() {
	*x = UrlListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_url_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UrlListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UrlListResponse) ProtoMessage() {}

func (x *UrlListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_url_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UrlListResponse.ProtoReflect.Descriptor instead.
func (*UrlListResponse) Descriptor() ([]byte, []int) {
	return file_rpc_url_proto_rawDescGZIP(), []int{2}
}

func (x *UrlListResponse) GetUrls() []string {
	if x != nil {
		return x.Urls
	}
	return nil
}

type UrlIntervalResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UrlInterval map[string]int32 `protobuf:"bytes,1,rep,name=UrlInterval,proto3" json:"UrlInterval,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *UrlIntervalResponse) Reset() {
	*x = UrlIntervalResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_url_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UrlIntervalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UrlIntervalResponse) ProtoMessage() {}

func (x *UrlIntervalResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_url_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UrlIntervalResponse.ProtoReflect.Descriptor instead.
func (*UrlIntervalResponse) Descriptor() ([]byte, []int) {
	return file_rpc_url_proto_rawDescGZIP(), []int{3}
}

func (x *UrlIntervalResponse) GetUrlInterval() map[string]int32 {
	if x != nil {
		return x.UrlInterval
	}
	return nil
}

type ProxyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Proxy string `protobuf:"bytes,1,opt,name=proxy,proto3" json:"proxy,omitempty"`
}

func (x *ProxyRequest) Reset() {
	*x = ProxyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_url_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyRequest) ProtoMessage() {}

func (x *ProxyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_url_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyRequest.ProtoReflect.Descriptor instead.
func (*ProxyRequest) Descriptor() ([]byte, []int) {
	return file_rpc_url_proto_rawDescGZIP(), []int{4}
}

func (x *ProxyRequest) GetProxy() string {
	if x != nil {
		return x.Proxy
	}
	return ""
}

type ProxyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *ProxyResponse) Reset() {
	*x = ProxyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_url_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyResponse) ProtoMessage() {}

func (x *ProxyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_url_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyResponse.ProtoReflect.Descriptor instead.
func (*ProxyResponse) Descriptor() ([]byte, []int) {
	return file_rpc_url_proto_rawDescGZIP(), []int{5}
}

func (x *ProxyResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type ProxyListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProxyList []string `protobuf:"bytes,1,rep,name=proxyList,proto3" json:"proxyList,omitempty"`
}

func (x *ProxyListResponse) Reset() {
	*x = ProxyListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_url_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyListResponse) ProtoMessage() {}

func (x *ProxyListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_url_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyListResponse.ProtoReflect.Descriptor instead.
func (*ProxyListResponse) Descriptor() ([]byte, []int) {
	return file_rpc_url_proto_rawDescGZIP(), []int{6}
}

func (x *ProxyListResponse) GetProxyList() []string {
	if x != nil {
		return x.ProxyList
	}
	return nil
}

var File_rpc_url_proto protoreflect.FileDescriptor

var file_rpc_url_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x70, 0x63, 0x2f, 0x75, 0x72, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x72, 0x70, 0x63, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x3a, 0x0a, 0x0a, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x22, 0x25, 0x0a,
	0x0b, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x22, 0x25, 0x0a, 0x0f, 0x55, 0x72, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x22, 0xa2, 0x01, 0x0a, 0x13,
	0x55, 0x72, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0b, 0x55, 0x72, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x55,
	0x72, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x55, 0x72, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x0b, 0x55, 0x72, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x1a, 0x3e, 0x0a, 0x10, 0x55, 0x72, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x24, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x22, 0x27, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x31, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x4c, 0x69, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x4c, 0x69,
	0x73, 0x74, 0x32, 0x96, 0x03, 0x0a, 0x0a, 0x55, 0x72, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x2b, 0x0a, 0x06, 0x53, 0x65, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x0f, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36,
	0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x14, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x72, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x41, 0x6e, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x55, 0x72, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x72, 0x6c,
	0x12, 0x0f, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x10, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12,
	0x11, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x12, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x50, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x78, 0x79, 0x12, 0x11, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x72,
	0x6f, 0x78, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2f,
	0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_url_proto_rawDescOnce sync.Once
	file_rpc_url_proto_rawDescData = file_rpc_url_proto_rawDesc
)

func file_rpc_url_proto_rawDescGZIP() []byte {
	file_rpc_url_proto_rawDescOnce.Do(func() {
		file_rpc_url_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_url_proto_rawDescData)
	})
	return file_rpc_url_proto_rawDescData
}

var file_rpc_url_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_rpc_url_proto_goTypes = []interface{}{
	(*UrlRequest)(nil),          // 0: rpc.UrlRequest
	(*UrlResponse)(nil),         // 1: rpc.UrlResponse
	(*UrlListResponse)(nil),     // 2: rpc.UrlListResponse
	(*UrlIntervalResponse)(nil), // 3: rpc.UrlIntervalResponse
	(*ProxyRequest)(nil),        // 4: rpc.ProxyRequest
	(*ProxyResponse)(nil),       // 5: rpc.ProxyResponse
	(*ProxyListResponse)(nil),   // 6: rpc.ProxyListResponse
	nil,                         // 7: rpc.UrlIntervalResponse.UrlIntervalEntry
	(*empty.Empty)(nil),         // 8: google.protobuf.Empty
}
var file_rpc_url_proto_depIdxs = []int32{
	7, // 0: rpc.UrlIntervalResponse.UrlInterval:type_name -> rpc.UrlIntervalResponse.UrlIntervalEntry
	0, // 1: rpc.UrlService.SetUrl:input_type -> rpc.UrlRequest
	8, // 2: rpc.UrlService.GetAll:input_type -> google.protobuf.Empty
	8, // 3: rpc.UrlService.GetAllDomainAndInterval:input_type -> google.protobuf.Empty
	0, // 4: rpc.UrlService.DeleteUrl:input_type -> rpc.UrlRequest
	4, // 5: rpc.UrlService.SetProxy:input_type -> rpc.ProxyRequest
	8, // 6: rpc.UrlService.GetAllProxy:input_type -> google.protobuf.Empty
	4, // 7: rpc.UrlService.DeleteProxy:input_type -> rpc.ProxyRequest
	1, // 8: rpc.UrlService.SetUrl:output_type -> rpc.UrlResponse
	2, // 9: rpc.UrlService.GetAll:output_type -> rpc.UrlListResponse
	3, // 10: rpc.UrlService.GetAllDomainAndInterval:output_type -> rpc.UrlIntervalResponse
	1, // 11: rpc.UrlService.DeleteUrl:output_type -> rpc.UrlResponse
	5, // 12: rpc.UrlService.SetProxy:output_type -> rpc.ProxyResponse
	6, // 13: rpc.UrlService.GetAllProxy:output_type -> rpc.ProxyListResponse
	5, // 14: rpc.UrlService.DeleteProxy:output_type -> rpc.ProxyResponse
	8, // [8:15] is the sub-list for method output_type
	1, // [1:8] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_url_proto_init() }
func file_rpc_url_proto_init() {
	if File_rpc_url_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_url_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UrlRequest); i {
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
		file_rpc_url_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UrlResponse); i {
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
		file_rpc_url_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UrlListResponse); i {
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
		file_rpc_url_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UrlIntervalResponse); i {
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
		file_rpc_url_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyRequest); i {
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
		file_rpc_url_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyResponse); i {
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
		file_rpc_url_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyListResponse); i {
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
			RawDescriptor: file_rpc_url_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_url_proto_goTypes,
		DependencyIndexes: file_rpc_url_proto_depIdxs,
		MessageInfos:      file_rpc_url_proto_msgTypes,
	}.Build()
	File_rpc_url_proto = out.File
	file_rpc_url_proto_rawDesc = nil
	file_rpc_url_proto_goTypes = nil
	file_rpc_url_proto_depIdxs = nil
}
