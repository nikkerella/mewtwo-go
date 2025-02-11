// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v3.21.1
// source: stock.proto

package stock

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

type CheckStockRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductId     string                 `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckStockRequest) Reset() {
	*x = CheckStockRequest{}
	mi := &file_stock_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckStockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckStockRequest) ProtoMessage() {}

func (x *CheckStockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stock_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckStockRequest.ProtoReflect.Descriptor instead.
func (*CheckStockRequest) Descriptor() ([]byte, []int) {
	return file_stock_proto_rawDescGZIP(), []int{0}
}

func (x *CheckStockRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

type CheckStockResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Available     bool                   `protobuf:"varint,1,opt,name=available,proto3" json:"available,omitempty"`
	CurrentStock  int32                  `protobuf:"varint,2,opt,name=current_stock,json=currentStock,proto3" json:"current_stock,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckStockResponse) Reset() {
	*x = CheckStockResponse{}
	mi := &file_stock_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckStockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckStockResponse) ProtoMessage() {}

func (x *CheckStockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stock_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckStockResponse.ProtoReflect.Descriptor instead.
func (*CheckStockResponse) Descriptor() ([]byte, []int) {
	return file_stock_proto_rawDescGZIP(), []int{1}
}

func (x *CheckStockResponse) GetAvailable() bool {
	if x != nil {
		return x.Available
	}
	return false
}

func (x *CheckStockResponse) GetCurrentStock() int32 {
	if x != nil {
		return x.CurrentStock
	}
	return 0
}

type DeductStockRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductId     string                 `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeductStockRequest) Reset() {
	*x = DeductStockRequest{}
	mi := &file_stock_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeductStockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeductStockRequest) ProtoMessage() {}

func (x *DeductStockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stock_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeductStockRequest.ProtoReflect.Descriptor instead.
func (*DeductStockRequest) Descriptor() ([]byte, []int) {
	return file_stock_proto_rawDescGZIP(), []int{2}
}

func (x *DeductStockRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

type DeductStockResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	UpdatedStock  int32                  `protobuf:"varint,2,opt,name=updated_stock,json=updatedStock,proto3" json:"updated_stock,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeductStockResponse) Reset() {
	*x = DeductStockResponse{}
	mi := &file_stock_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeductStockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeductStockResponse) ProtoMessage() {}

func (x *DeductStockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stock_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeductStockResponse.ProtoReflect.Descriptor instead.
func (*DeductStockResponse) Descriptor() ([]byte, []int) {
	return file_stock_proto_rawDescGZIP(), []int{3}
}

func (x *DeductStockResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *DeductStockResponse) GetUpdatedStock() int32 {
	if x != nil {
		return x.UpdatedStock
	}
	return 0
}

var File_stock_proto protoreflect.FileDescriptor

var file_stock_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x22, 0x32, 0x0a, 0x11, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x22, 0x57, 0x0a, 0x12, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x22, 0x33, 0x0a, 0x12, 0x44, 0x65, 0x64, 0x75, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x22, 0x54, 0x0a, 0x13, 0x44, 0x65, 0x64, 0x75, 0x63, 0x74,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x32, 0x97, 0x01, 0x0a,
	0x0c, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a,
	0x0a, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x18, 0x2e, 0x73, 0x74,
	0x6f, 0x63, 0x6b, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x44, 0x0a, 0x0b, 0x44, 0x65, 0x64, 0x75, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12,
	0x19, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x44, 0x65, 0x64, 0x75, 0x63, 0x74, 0x53, 0x74,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x74, 0x6f,
	0x63, 0x6b, 0x2e, 0x44, 0x65, 0x64, 0x75, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x14, 0x5a, 0x12, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x3b, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stock_proto_rawDescOnce sync.Once
	file_stock_proto_rawDescData = file_stock_proto_rawDesc
)

func file_stock_proto_rawDescGZIP() []byte {
	file_stock_proto_rawDescOnce.Do(func() {
		file_stock_proto_rawDescData = protoimpl.X.CompressGZIP(file_stock_proto_rawDescData)
	})
	return file_stock_proto_rawDescData
}

var file_stock_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_stock_proto_goTypes = []any{
	(*CheckStockRequest)(nil),   // 0: stock.CheckStockRequest
	(*CheckStockResponse)(nil),  // 1: stock.CheckStockResponse
	(*DeductStockRequest)(nil),  // 2: stock.DeductStockRequest
	(*DeductStockResponse)(nil), // 3: stock.DeductStockResponse
}
var file_stock_proto_depIdxs = []int32{
	0, // 0: stock.StockService.CheckStock:input_type -> stock.CheckStockRequest
	2, // 1: stock.StockService.DeductStock:input_type -> stock.DeductStockRequest
	1, // 2: stock.StockService.CheckStock:output_type -> stock.CheckStockResponse
	3, // 3: stock.StockService.DeductStock:output_type -> stock.DeductStockResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_stock_proto_init() }
func file_stock_proto_init() {
	if File_stock_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_stock_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stock_proto_goTypes,
		DependencyIndexes: file_stock_proto_depIdxs,
		MessageInfos:      file_stock_proto_msgTypes,
	}.Build()
	File_stock_proto = out.File
	file_stock_proto_rawDesc = nil
	file_stock_proto_goTypes = nil
	file_stock_proto_depIdxs = nil
}
