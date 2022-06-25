// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.11.2
// source: internal/rpc/url.proto

package rpc

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
		mi := &file_internal_rpc_url_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UrlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UrlRequest) ProtoMessage() {}

func (x *UrlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_rpc_url_proto_msgTypes[0]
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
	return file_internal_rpc_url_proto_rawDescGZIP(), []int{0}
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

	Result string `protobuf:"bytes,3,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *UrlResponse) Reset() {
	*x = UrlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_rpc_url_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UrlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UrlResponse) ProtoMessage() {}

func (x *UrlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_rpc_url_proto_msgTypes[1]
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
	return file_internal_rpc_url_proto_rawDescGZIP(), []int{1}
}

func (x *UrlResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type ProxyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Proxy string `protobuf:"bytes,4,opt,name=proxy,proto3" json:"proxy,omitempty"`
}

func (x *ProxyRequest) Reset() {
	*x = ProxyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_rpc_url_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyRequest) ProtoMessage() {}

func (x *ProxyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_rpc_url_proto_msgTypes[2]
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
	return file_internal_rpc_url_proto_rawDescGZIP(), []int{2}
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

	Result string `protobuf:"bytes,5,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *ProxyResponse) Reset() {
	*x = ProxyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_rpc_url_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyResponse) ProtoMessage() {}

func (x *ProxyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_rpc_url_proto_msgTypes[3]
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
	return file_internal_rpc_url_proto_rawDescGZIP(), []int{3}
}

func (x *ProxyResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_internal_rpc_url_proto protoreflect.FileDescriptor

var file_internal_rpc_url_proto_rawDesc = []byte{
	0x0a, 0x16, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x75,
	0x72, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x70, 0x63, 0x22, 0x3a, 0x0a,
	0x0a, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1a, 0x0a,
	0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x22, 0x25, 0x0a, 0x0b, 0x55, 0x72, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x22, 0x24, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x22, 0x27, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32,
	0x6c, 0x0a, 0x0a, 0x55, 0x72, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2b, 0x0a,
	0x06, 0x53, 0x65, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x0f, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x72,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x55,
	0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x53, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x11, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x72, 0x6f,
	0x78, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0e, 0x5a,
	0x0c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_rpc_url_proto_rawDescOnce sync.Once
	file_internal_rpc_url_proto_rawDescData = file_internal_rpc_url_proto_rawDesc
)

func file_internal_rpc_url_proto_rawDescGZIP() []byte {
	file_internal_rpc_url_proto_rawDescOnce.Do(func() {
		file_internal_rpc_url_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_rpc_url_proto_rawDescData)
	})
	return file_internal_rpc_url_proto_rawDescData
}

var file_internal_rpc_url_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_internal_rpc_url_proto_goTypes = []interface{}{
	(*UrlRequest)(nil),    // 0: rpc.UrlRequest
	(*UrlResponse)(nil),   // 1: rpc.UrlResponse
	(*ProxyRequest)(nil),  // 2: rpc.ProxyRequest
	(*ProxyResponse)(nil), // 3: rpc.ProxyResponse
}
var file_internal_rpc_url_proto_depIdxs = []int32{
	0, // 0: rpc.UrlService.SetUrl:input_type -> rpc.UrlRequest
	2, // 1: rpc.UrlService.SetProxy:input_type -> rpc.ProxyRequest
	1, // 2: rpc.UrlService.SetUrl:output_type -> rpc.UrlResponse
	3, // 3: rpc.UrlService.SetProxy:output_type -> rpc.ProxyResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_rpc_url_proto_init() }
func file_internal_rpc_url_proto_init() {
	if File_internal_rpc_url_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_rpc_url_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_internal_rpc_url_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_internal_rpc_url_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_internal_rpc_url_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_rpc_url_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_rpc_url_proto_goTypes,
		DependencyIndexes: file_internal_rpc_url_proto_depIdxs,
		MessageInfos:      file_internal_rpc_url_proto_msgTypes,
	}.Build()
	File_internal_rpc_url_proto = out.File
	file_internal_rpc_url_proto_rawDesc = nil
	file_internal_rpc_url_proto_goTypes = nil
	file_internal_rpc_url_proto_depIdxs = nil
}
