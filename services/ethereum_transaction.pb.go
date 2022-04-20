// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.17.3
// source: ethereum_transaction.proto

package services

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

type EthereumTransactionBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// The raw Ethereum transaction (RLP encoded type 0, 1, and 2). Complete
	// unless the callData field is set.
	EthereumData []byte `protobuf:"bytes,1,opt,name=ethereum_data,json=ethereumData,proto3" json:"ethereum_data,omitempty"`
	//*
	// For large transactions (for example contract create) this is the callData
	// of the ethereumData. The data in the ethereumData will be re-written with
	// the callData element as a zero length string with the original contents in
	// the referenced file at time of execution. The ethereumData will need to be
	// "rehydrated" with the callData for signature validation to pass.
	CallData *FileID `protobuf:"bytes,2,opt,name=call_data,json=callData,proto3" json:"call_data,omitempty"`
	//*
	// The maximum amount, in tinybars, that the payer of the hedera transaction
	// is willing to pay to complete the transaction.
	//
	// Ordinarily the account with the ECDSA alias corresponding to the public
	// key that is extracted from the ethereum_data signature is responsible for
	// fees that result from the execution of the transaction. If that amount of
	// authorized fees is not sufficient then the payer of the transaction can be
	// charged, up to but not exceeding this amount. If the ethereum_data
	// transaction authorized an amount that was insufficient then the payer will
	// only be charged the amount needed to make up the difference. If the gas
	// price in the transaction was set to zero then the payer will be assessed
	// the entire fee.
	MaxGasAllowance int64 `protobuf:"varint,3,opt,name=max_gas_allowance,json=maxGasAllowance,proto3" json:"max_gas_allowance,omitempty"`
}

func (x *EthereumTransactionBody) Reset() {
	*x = EthereumTransactionBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ethereum_transaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EthereumTransactionBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EthereumTransactionBody) ProtoMessage() {}

func (x *EthereumTransactionBody) ProtoReflect() protoreflect.Message {
	mi := &file_ethereum_transaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EthereumTransactionBody.ProtoReflect.Descriptor instead.
func (*EthereumTransactionBody) Descriptor() ([]byte, []int) {
	return file_ethereum_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *EthereumTransactionBody) GetEthereumData() []byte {
	if x != nil {
		return x.EthereumData
	}
	return nil
}

func (x *EthereumTransactionBody) GetCallData() *FileID {
	if x != nil {
		return x.CallData
	}
	return nil
}

func (x *EthereumTransactionBody) GetMaxGasAllowance() int64 {
	if x != nil {
		return x.MaxGasAllowance
	}
	return 0
}

var File_ethereum_transaction_proto protoreflect.FileDescriptor

var file_ethereum_transaction_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96, 0x01, 0x0a, 0x17, 0x45, 0x74, 0x68, 0x65, 0x72,
	0x65, 0x75, 0x6d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6f,
	0x64, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x65, 0x74, 0x68, 0x65, 0x72,
	0x65, 0x75, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2a, 0x0a, 0x09, 0x63, 0x61, 0x6c, 0x6c, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x44, 0x52, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x2a, 0x0a, 0x11, 0x6d, 0x61, 0x78, 0x5f, 0x67, 0x61, 0x73, 0x5f, 0x61,
	0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f,
	0x6d, 0x61, 0x78, 0x47, 0x61, 0x73, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x42,
	0x26, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x68, 0x61, 0x73,
	0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x6a, 0x61, 0x76, 0x61, 0x50, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ethereum_transaction_proto_rawDescOnce sync.Once
	file_ethereum_transaction_proto_rawDescData = file_ethereum_transaction_proto_rawDesc
)

func file_ethereum_transaction_proto_rawDescGZIP() []byte {
	file_ethereum_transaction_proto_rawDescOnce.Do(func() {
		file_ethereum_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_ethereum_transaction_proto_rawDescData)
	})
	return file_ethereum_transaction_proto_rawDescData
}

var file_ethereum_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ethereum_transaction_proto_goTypes = []interface{}{
	(*EthereumTransactionBody)(nil), // 0: proto.EthereumTransactionBody
	(*FileID)(nil),                  // 1: proto.FileID
}
var file_ethereum_transaction_proto_depIdxs = []int32{
	1, // 0: proto.EthereumTransactionBody.call_data:type_name -> proto.FileID
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ethereum_transaction_proto_init() }
func file_ethereum_transaction_proto_init() {
	if File_ethereum_transaction_proto != nil {
		return
	}
	file_basic_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ethereum_transaction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EthereumTransactionBody); i {
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
			RawDescriptor: file_ethereum_transaction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ethereum_transaction_proto_goTypes,
		DependencyIndexes: file_ethereum_transaction_proto_depIdxs,
		MessageInfos:      file_ethereum_transaction_proto_msgTypes,
	}.Build()
	File_ethereum_transaction_proto = out.File
	file_ethereum_transaction_proto_rawDesc = nil
	file_ethereum_transaction_proto_goTypes = nil
	file_ethereum_transaction_proto_depIdxs = nil
}
