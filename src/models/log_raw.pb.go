// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: log_raw.proto

package models

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

type LogRaw struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type             string `protobuf:"bytes,1,opt,name=type,proto3" json:"type"`
	LogIndex         uint64 `protobuf:"varint,2,opt,name=log_index,json=logIndex,proto3" json:"log_index"`
	MaxLogIndex      uint64 `protobuf:"varint,3,opt,name=max_log_index,json=maxLogIndex,proto3" json:"max_log_index"`
	TransactionHash  string `protobuf:"bytes,4,opt,name=transaction_hash,json=transactionHash,proto3" json:"transaction_hash"`
	TransactionIndex uint32 `protobuf:"varint,5,opt,name=transaction_index,json=transactionIndex,proto3" json:"transaction_index"`
	Address          string `protobuf:"bytes,6,opt,name=address,proto3" json:"address"`
	Data             string `protobuf:"bytes,7,opt,name=data,proto3" json:"data"`
	Indexed          string `protobuf:"bytes,8,opt,name=indexed,proto3" json:"indexed"`
	BlockNumber      uint64 `protobuf:"varint,9,opt,name=block_number,json=blockNumber,proto3" json:"block_number"`
	BlockTimestamp   uint64 `protobuf:"varint,10,opt,name=block_timestamp,json=blockTimestamp,proto3" json:"block_timestamp"`
	BlockHash        string `protobuf:"bytes,11,opt,name=block_hash,json=blockHash,proto3" json:"block_hash"`
	ItemId           string `protobuf:"bytes,12,opt,name=item_id,json=itemId,proto3" json:"item_id"`
	ItemTimestamp    string `protobuf:"bytes,13,opt,name=item_timestamp,json=itemTimestamp,proto3" json:"item_timestamp"`
}

func (x *LogRaw) Reset() {
	*x = LogRaw{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_raw_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogRaw) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogRaw) ProtoMessage() {}

func (x *LogRaw) ProtoReflect() protoreflect.Message {
	mi := &file_log_raw_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogRaw.ProtoReflect.Descriptor instead.
func (*LogRaw) Descriptor() ([]byte, []int) {
	return file_log_raw_proto_rawDescGZIP(), []int{0}
}

func (x *LogRaw) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *LogRaw) GetLogIndex() uint64 {
	if x != nil {
		return x.LogIndex
	}
	return 0
}

func (x *LogRaw) GetMaxLogIndex() uint64 {
	if x != nil {
		return x.MaxLogIndex
	}
	return 0
}

func (x *LogRaw) GetTransactionHash() string {
	if x != nil {
		return x.TransactionHash
	}
	return ""
}

func (x *LogRaw) GetTransactionIndex() uint32 {
	if x != nil {
		return x.TransactionIndex
	}
	return 0
}

func (x *LogRaw) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *LogRaw) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *LogRaw) GetIndexed() string {
	if x != nil {
		return x.Indexed
	}
	return ""
}

func (x *LogRaw) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *LogRaw) GetBlockTimestamp() uint64 {
	if x != nil {
		return x.BlockTimestamp
	}
	return 0
}

func (x *LogRaw) GetBlockHash() string {
	if x != nil {
		return x.BlockHash
	}
	return ""
}

func (x *LogRaw) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

func (x *LogRaw) GetItemTimestamp() string {
	if x != nil {
		return x.ItemTimestamp
	}
	return ""
}

var File_log_raw_proto protoreflect.FileDescriptor

var file_log_raw_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6c, 0x6f, 0x67, 0x5f, 0x72, 0x61, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x22, 0xa8, 0x03, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x52,
	0x61, 0x77, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x22, 0x0a, 0x0d, 0x6d, 0x61, 0x78, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x6d, 0x61, 0x78, 0x4c,
	0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61,
	0x73, 0x68, 0x12, 0x2b, 0x0a, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a,
	0x07, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x61, 0x73,
	0x68, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61,
	0x73, 0x68, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x69,
	0x74, 0x65, 0x6d, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x74, 0x65, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_log_raw_proto_rawDescOnce sync.Once
	file_log_raw_proto_rawDescData = file_log_raw_proto_rawDesc
)

func file_log_raw_proto_rawDescGZIP() []byte {
	file_log_raw_proto_rawDescOnce.Do(func() {
		file_log_raw_proto_rawDescData = protoimpl.X.CompressGZIP(file_log_raw_proto_rawDescData)
	})
	return file_log_raw_proto_rawDescData
}

var file_log_raw_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_log_raw_proto_goTypes = []interface{}{
	(*LogRaw)(nil), // 0: models.LogRaw
}
var file_log_raw_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_log_raw_proto_init() }
func file_log_raw_proto_init() {
	if File_log_raw_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_log_raw_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogRaw); i {
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
			RawDescriptor: file_log_raw_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_log_raw_proto_goTypes,
		DependencyIndexes: file_log_raw_proto_depIdxs,
		MessageInfos:      file_log_raw_proto_msgTypes,
	}.Build()
	File_log_raw_proto = out.File
	file_log_raw_proto_rawDesc = nil
	file_log_raw_proto_goTypes = nil
	file_log_raw_proto_depIdxs = nil
}
