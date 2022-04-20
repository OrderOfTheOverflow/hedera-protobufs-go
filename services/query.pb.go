// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.17.3
// source: query.proto

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

//*
// A single query, which is sent from the client to a node. This includes all possible queries. Each
// Query should not have more than 50 levels.
type Query struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Query:
	//	*Query_GetByKey
	//	*Query_GetBySolidityID
	//	*Query_ContractCallLocal
	//	*Query_ContractGetInfo
	//	*Query_ContractGetBytecode
	//	*Query_ContractGetRecords
	//	*Query_CryptogetAccountBalance
	//	*Query_CryptoGetAccountRecords
	//	*Query_CryptoGetInfo
	//	*Query_CryptoGetLiveHash
	//	*Query_CryptoGetProxyStakers
	//	*Query_FileGetContents
	//	*Query_FileGetInfo
	//	*Query_TransactionGetReceipt
	//	*Query_TransactionGetRecord
	//	*Query_TransactionGetFastRecord
	//	*Query_ConsensusGetTopicInfo
	//	*Query_NetworkGetVersionInfo
	//	*Query_TokenGetInfo
	//	*Query_ScheduleGetInfo
	//	*Query_TokenGetAccountNftInfos
	//	*Query_TokenGetNftInfo
	//	*Query_TokenGetNftInfos
	//	*Query_NetworkGetExecutionTime
	//	*Query_AccountDetails
	Query isQuery_Query `protobuf_oneof:"query"`
}

func (x *Query) Reset() {
	*x = Query{}
	if protoimpl.UnsafeEnabled {
		mi := &file_query_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Query) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Query) ProtoMessage() {}

func (x *Query) ProtoReflect() protoreflect.Message {
	mi := &file_query_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Query.ProtoReflect.Descriptor instead.
func (*Query) Descriptor() ([]byte, []int) {
	return file_query_proto_rawDescGZIP(), []int{0}
}

func (m *Query) GetQuery() isQuery_Query {
	if m != nil {
		return m.Query
	}
	return nil
}

func (x *Query) GetGetByKey() *GetByKeyQuery {
	if x, ok := x.GetQuery().(*Query_GetByKey); ok {
		return x.GetByKey
	}
	return nil
}

func (x *Query) GetGetBySolidityID() *GetBySolidityIDQuery {
	if x, ok := x.GetQuery().(*Query_GetBySolidityID); ok {
		return x.GetBySolidityID
	}
	return nil
}

func (x *Query) GetContractCallLocal() *ContractCallLocalQuery {
	if x, ok := x.GetQuery().(*Query_ContractCallLocal); ok {
		return x.ContractCallLocal
	}
	return nil
}

func (x *Query) GetContractGetInfo() *ContractGetInfoQuery {
	if x, ok := x.GetQuery().(*Query_ContractGetInfo); ok {
		return x.ContractGetInfo
	}
	return nil
}

func (x *Query) GetContractGetBytecode() *ContractGetBytecodeQuery {
	if x, ok := x.GetQuery().(*Query_ContractGetBytecode); ok {
		return x.ContractGetBytecode
	}
	return nil
}

func (x *Query) GetContractGetRecords() *ContractGetRecordsQuery {
	if x, ok := x.GetQuery().(*Query_ContractGetRecords); ok {
		return x.ContractGetRecords
	}
	return nil
}

func (x *Query) GetCryptogetAccountBalance() *CryptoGetAccountBalanceQuery {
	if x, ok := x.GetQuery().(*Query_CryptogetAccountBalance); ok {
		return x.CryptogetAccountBalance
	}
	return nil
}

func (x *Query) GetCryptoGetAccountRecords() *CryptoGetAccountRecordsQuery {
	if x, ok := x.GetQuery().(*Query_CryptoGetAccountRecords); ok {
		return x.CryptoGetAccountRecords
	}
	return nil
}

func (x *Query) GetCryptoGetInfo() *CryptoGetInfoQuery {
	if x, ok := x.GetQuery().(*Query_CryptoGetInfo); ok {
		return x.CryptoGetInfo
	}
	return nil
}

func (x *Query) GetCryptoGetLiveHash() *CryptoGetLiveHashQuery {
	if x, ok := x.GetQuery().(*Query_CryptoGetLiveHash); ok {
		return x.CryptoGetLiveHash
	}
	return nil
}

func (x *Query) GetCryptoGetProxyStakers() *CryptoGetStakersQuery {
	if x, ok := x.GetQuery().(*Query_CryptoGetProxyStakers); ok {
		return x.CryptoGetProxyStakers
	}
	return nil
}

func (x *Query) GetFileGetContents() *FileGetContentsQuery {
	if x, ok := x.GetQuery().(*Query_FileGetContents); ok {
		return x.FileGetContents
	}
	return nil
}

func (x *Query) GetFileGetInfo() *FileGetInfoQuery {
	if x, ok := x.GetQuery().(*Query_FileGetInfo); ok {
		return x.FileGetInfo
	}
	return nil
}

func (x *Query) GetTransactionGetReceipt() *TransactionGetReceiptQuery {
	if x, ok := x.GetQuery().(*Query_TransactionGetReceipt); ok {
		return x.TransactionGetReceipt
	}
	return nil
}

func (x *Query) GetTransactionGetRecord() *TransactionGetRecordQuery {
	if x, ok := x.GetQuery().(*Query_TransactionGetRecord); ok {
		return x.TransactionGetRecord
	}
	return nil
}

func (x *Query) GetTransactionGetFastRecord() *TransactionGetFastRecordQuery {
	if x, ok := x.GetQuery().(*Query_TransactionGetFastRecord); ok {
		return x.TransactionGetFastRecord
	}
	return nil
}

func (x *Query) GetConsensusGetTopicInfo() *ConsensusGetTopicInfoQuery {
	if x, ok := x.GetQuery().(*Query_ConsensusGetTopicInfo); ok {
		return x.ConsensusGetTopicInfo
	}
	return nil
}

func (x *Query) GetNetworkGetVersionInfo() *NetworkGetVersionInfoQuery {
	if x, ok := x.GetQuery().(*Query_NetworkGetVersionInfo); ok {
		return x.NetworkGetVersionInfo
	}
	return nil
}

func (x *Query) GetTokenGetInfo() *TokenGetInfoQuery {
	if x, ok := x.GetQuery().(*Query_TokenGetInfo); ok {
		return x.TokenGetInfo
	}
	return nil
}

func (x *Query) GetScheduleGetInfo() *ScheduleGetInfoQuery {
	if x, ok := x.GetQuery().(*Query_ScheduleGetInfo); ok {
		return x.ScheduleGetInfo
	}
	return nil
}

func (x *Query) GetTokenGetAccountNftInfos() *TokenGetAccountNftInfosQuery {
	if x, ok := x.GetQuery().(*Query_TokenGetAccountNftInfos); ok {
		return x.TokenGetAccountNftInfos
	}
	return nil
}

func (x *Query) GetTokenGetNftInfo() *TokenGetNftInfoQuery {
	if x, ok := x.GetQuery().(*Query_TokenGetNftInfo); ok {
		return x.TokenGetNftInfo
	}
	return nil
}

func (x *Query) GetTokenGetNftInfos() *TokenGetNftInfosQuery {
	if x, ok := x.GetQuery().(*Query_TokenGetNftInfos); ok {
		return x.TokenGetNftInfos
	}
	return nil
}

func (x *Query) GetNetworkGetExecutionTime() *NetworkGetExecutionTimeQuery {
	if x, ok := x.GetQuery().(*Query_NetworkGetExecutionTime); ok {
		return x.NetworkGetExecutionTime
	}
	return nil
}

func (x *Query) GetAccountDetails() *GetAccountDetailsQuery {
	if x, ok := x.GetQuery().(*Query_AccountDetails); ok {
		return x.AccountDetails
	}
	return nil
}

type isQuery_Query interface {
	isQuery_Query()
}

type Query_GetByKey struct {
	//*
	// Get all entities associated with a given key
	GetByKey *GetByKeyQuery `protobuf:"bytes,1,opt,name=getByKey,proto3,oneof"`
}

type Query_GetBySolidityID struct {
	//*
	// Get the IDs in the format used in transactions, given the format used in Solidity
	GetBySolidityID *GetBySolidityIDQuery `protobuf:"bytes,2,opt,name=getBySolidityID,proto3,oneof"`
}

type Query_ContractCallLocal struct {
	//*
	// Call a function of a smart contract instance
	ContractCallLocal *ContractCallLocalQuery `protobuf:"bytes,3,opt,name=contractCallLocal,proto3,oneof"`
}

type Query_ContractGetInfo struct {
	//*
	// Get information about a smart contract instance
	ContractGetInfo *ContractGetInfoQuery `protobuf:"bytes,4,opt,name=contractGetInfo,proto3,oneof"`
}

type Query_ContractGetBytecode struct {
	//*
	// Get runtime code used by a smart contract instance
	ContractGetBytecode *ContractGetBytecodeQuery `protobuf:"bytes,5,opt,name=contractGetBytecode,proto3,oneof"`
}

type Query_ContractGetRecords struct {
	//*
	// Get Records of the contract instance
	ContractGetRecords *ContractGetRecordsQuery `protobuf:"bytes,6,opt,name=ContractGetRecords,proto3,oneof"`
}

type Query_CryptogetAccountBalance struct {
	//*
	// Get the current balance in a cryptocurrency account
	CryptogetAccountBalance *CryptoGetAccountBalanceQuery `protobuf:"bytes,7,opt,name=cryptogetAccountBalance,proto3,oneof"`
}

type Query_CryptoGetAccountRecords struct {
	//*
	// Get all the records that currently exist for transactions involving an account
	CryptoGetAccountRecords *CryptoGetAccountRecordsQuery `protobuf:"bytes,8,opt,name=cryptoGetAccountRecords,proto3,oneof"`
}

type Query_CryptoGetInfo struct {
	//*
	// Get all information about an account
	CryptoGetInfo *CryptoGetInfoQuery `protobuf:"bytes,9,opt,name=cryptoGetInfo,proto3,oneof"`
}

type Query_CryptoGetLiveHash struct {
	//*
	// Get a single livehash from a single account, if present
	CryptoGetLiveHash *CryptoGetLiveHashQuery `protobuf:"bytes,10,opt,name=cryptoGetLiveHash,proto3,oneof"`
}

type Query_CryptoGetProxyStakers struct {
	//*
	// Get all the accounts that proxy stake to a given account, and how much they proxy stake
	// (not yet implemented in the current API)
	CryptoGetProxyStakers *CryptoGetStakersQuery `protobuf:"bytes,11,opt,name=cryptoGetProxyStakers,proto3,oneof"`
}

type Query_FileGetContents struct {
	//*
	// Get the contents of a file (the bytes stored in it)
	FileGetContents *FileGetContentsQuery `protobuf:"bytes,12,opt,name=fileGetContents,proto3,oneof"`
}

type Query_FileGetInfo struct {
	//*
	// Get information about a file, such as its expiration date
	FileGetInfo *FileGetInfoQuery `protobuf:"bytes,13,opt,name=fileGetInfo,proto3,oneof"`
}

type Query_TransactionGetReceipt struct {
	//*
	// Get a receipt for a transaction (lasts 180 seconds)
	TransactionGetReceipt *TransactionGetReceiptQuery `protobuf:"bytes,14,opt,name=transactionGetReceipt,proto3,oneof"`
}

type Query_TransactionGetRecord struct {
	//*
	// Get a record for a transaction
	TransactionGetRecord *TransactionGetRecordQuery `protobuf:"bytes,15,opt,name=transactionGetRecord,proto3,oneof"`
}

type Query_TransactionGetFastRecord struct {
	//*
	// Get a record for a transaction (lasts 180 seconds)
	TransactionGetFastRecord *TransactionGetFastRecordQuery `protobuf:"bytes,16,opt,name=transactionGetFastRecord,proto3,oneof"`
}

type Query_ConsensusGetTopicInfo struct {
	//*
	// Get the parameters of and state of a consensus topic.
	ConsensusGetTopicInfo *ConsensusGetTopicInfoQuery `protobuf:"bytes,50,opt,name=consensusGetTopicInfo,proto3,oneof"`
}

type Query_NetworkGetVersionInfo struct {
	//*
	// Get the versions of the HAPI protobuf and Hedera Services software deployed on the
	// responding node.
	NetworkGetVersionInfo *NetworkGetVersionInfoQuery `protobuf:"bytes,51,opt,name=networkGetVersionInfo,proto3,oneof"`
}

type Query_TokenGetInfo struct {
	//*
	// Get all information about a token
	TokenGetInfo *TokenGetInfoQuery `protobuf:"bytes,52,opt,name=tokenGetInfo,proto3,oneof"`
}

type Query_ScheduleGetInfo struct {
	//*
	// Get all information about a scheduled entity
	ScheduleGetInfo *ScheduleGetInfoQuery `protobuf:"bytes,53,opt,name=scheduleGetInfo,proto3,oneof"`
}

type Query_TokenGetAccountNftInfos struct {
	//*
	// Get a list of NFTs associated with the account
	TokenGetAccountNftInfos *TokenGetAccountNftInfosQuery `protobuf:"bytes,54,opt,name=tokenGetAccountNftInfos,proto3,oneof"`
}

type Query_TokenGetNftInfo struct {
	//*
	// Get all information about a NFT
	TokenGetNftInfo *TokenGetNftInfoQuery `protobuf:"bytes,55,opt,name=tokenGetNftInfo,proto3,oneof"`
}

type Query_TokenGetNftInfos struct {
	//*
	// Get a list of NFTs for the token
	TokenGetNftInfos *TokenGetNftInfosQuery `protobuf:"bytes,56,opt,name=tokenGetNftInfos,proto3,oneof"`
}

type Query_NetworkGetExecutionTime struct {
	//*
	// Gets <tt>handleTransaction</tt> times for one or more "sufficiently recent" TransactionIDs
	NetworkGetExecutionTime *NetworkGetExecutionTimeQuery `protobuf:"bytes,57,opt,name=networkGetExecutionTime,proto3,oneof"`
}

type Query_AccountDetails struct {
	//*
	// Gets all information about an account including allowances granted by the account
	AccountDetails *GetAccountDetailsQuery `protobuf:"bytes,58,opt,name=accountDetails,proto3,oneof"`
}

func (*Query_GetByKey) isQuery_Query() {}

func (*Query_GetBySolidityID) isQuery_Query() {}

func (*Query_ContractCallLocal) isQuery_Query() {}

func (*Query_ContractGetInfo) isQuery_Query() {}

func (*Query_ContractGetBytecode) isQuery_Query() {}

func (*Query_ContractGetRecords) isQuery_Query() {}

func (*Query_CryptogetAccountBalance) isQuery_Query() {}

func (*Query_CryptoGetAccountRecords) isQuery_Query() {}

func (*Query_CryptoGetInfo) isQuery_Query() {}

func (*Query_CryptoGetLiveHash) isQuery_Query() {}

func (*Query_CryptoGetProxyStakers) isQuery_Query() {}

func (*Query_FileGetContents) isQuery_Query() {}

func (*Query_FileGetInfo) isQuery_Query() {}

func (*Query_TransactionGetReceipt) isQuery_Query() {}

func (*Query_TransactionGetRecord) isQuery_Query() {}

func (*Query_TransactionGetFastRecord) isQuery_Query() {}

func (*Query_ConsensusGetTopicInfo) isQuery_Query() {}

func (*Query_NetworkGetVersionInfo) isQuery_Query() {}

func (*Query_TokenGetInfo) isQuery_Query() {}

func (*Query_ScheduleGetInfo) isQuery_Query() {}

func (*Query_TokenGetAccountNftInfos) isQuery_Query() {}

func (*Query_TokenGetNftInfo) isQuery_Query() {}

func (*Query_TokenGetNftInfos) isQuery_Query() {}

func (*Query_NetworkGetExecutionTime) isQuery_Query() {}

func (*Query_AccountDetails) isQuery_Query() {}

var File_query_proto protoreflect.FileDescriptor

var file_query_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x67, 0x65, 0x74, 0x5f, 0x62, 0x79, 0x5f, 0x6b, 0x65, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x67, 0x65, 0x74, 0x5f, 0x62, 0x79, 0x5f, 0x73,
	0x6f, 0x6c, 0x69, 0x64, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x5f,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x67,
	0x65, 0x74, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x67, 0x65, 0x74, 0x5f,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x15, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x5f, 0x67, 0x65, 0x74, 0x5f, 0x6c, 0x69, 0x76, 0x65, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x74,
	0x5f, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x67, 0x65,
	0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x70, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x66, 0x61, 0x73, 0x74, 0x5f,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x63, 0x6f,
	0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x67,
	0x65, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x6e, 0x66, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x18, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x6e, 0x66, 0x74, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x6e, 0x66, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xf2, 0x0f, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x32, 0x0a, 0x08, 0x67, 0x65,
	0x74, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x48, 0x00, 0x52, 0x08, 0x67, 0x65, 0x74, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x47,
	0x0a, 0x0f, 0x67, 0x65, 0x74, 0x42, 0x79, 0x53, 0x6f, 0x6c, 0x69, 0x64, 0x69, 0x74, 0x79, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x47, 0x65, 0x74, 0x42, 0x79, 0x53, 0x6f, 0x6c, 0x69, 0x64, 0x69, 0x74, 0x79, 0x49, 0x44, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x48, 0x00, 0x52, 0x0f, 0x67, 0x65, 0x74, 0x42, 0x79, 0x53, 0x6f, 0x6c,
	0x69, 0x64, 0x69, 0x74, 0x79, 0x49, 0x44, 0x12, 0x4d, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x48, 0x00, 0x52, 0x11, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x43, 0x61, 0x6c,
	0x6c, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x12, 0x47, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x00, 0x52, 0x0f,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x53, 0x0a, 0x13, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x47, 0x65, 0x74, 0x42, 0x79,
	0x74, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x47, 0x65, 0x74,
	0x42, 0x79, 0x74, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x00, 0x52,
	0x13, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x47, 0x65, 0x74, 0x42, 0x79, 0x74, 0x65,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x50, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x48, 0x00, 0x52, 0x12, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x5f, 0x0a, 0x17, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x67, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x00, 0x52, 0x17,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x67, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x5f, 0x0a, 0x17, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x00, 0x52,
	0x17, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x41, 0x0a, 0x0d, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x47, 0x65,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x00, 0x52, 0x0d, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x4d, 0x0a, 0x11, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x48, 0x61, 0x73, 0x68,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x48, 0x61, 0x73, 0x68,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x00, 0x52, 0x11, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x47,
	0x65, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12, 0x54, 0x0a, 0x15, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x53, 0x74, 0x61, 0x6b,
	0x65, 0x72, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x6b, 0x65,
	0x72, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x00, 0x52, 0x15, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x73,
	0x12, 0x47, 0x0a, 0x0f, 0x66, 0x69, 0x6c, 0x65, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x00, 0x52, 0x0f, 0x66, 0x69, 0x6c, 0x65, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x3b, 0x0a, 0x0b, 0x66, 0x69, 0x6c,
	0x65, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x47, 0x65, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x00, 0x52, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x47,
	0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x59, 0x0a, 0x15, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65,
	0x69, 0x70, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x00, 0x52, 0x15, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70,
	0x74, 0x12, 0x56, 0x0a, 0x14, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x48, 0x00, 0x52, 0x14, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x62, 0x0a, 0x18, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x46, 0x61, 0x73, 0x74, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x47,
	0x65, 0x74, 0x46, 0x61, 0x73, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x48, 0x00, 0x52, 0x18, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x47, 0x65, 0x74, 0x46, 0x61, 0x73, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x59, 0x0a,
	0x15, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70,
	0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x32, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x47, 0x65,
	0x74, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x48,
	0x00, 0x52, 0x15, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x47, 0x65, 0x74, 0x54,
	0x6f, 0x70, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x59, 0x0a, 0x15, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x18, 0x33, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x00, 0x52, 0x15, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x3e, 0x0a, 0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x47, 0x65, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x34, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x48, 0x00, 0x52, 0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x47, 0x65, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x47, 0x0a, 0x0f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x47,
	0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x35, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x47, 0x65, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x00, 0x52, 0x0f, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x5f, 0x0a, 0x17,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e,
	0x66, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x36, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x47, 0x65, 0x74, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x66, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x48, 0x00, 0x52, 0x17, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x47, 0x65, 0x74, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x66, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x47, 0x0a,
	0x0f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x47, 0x65, 0x74, 0x4e, 0x66, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x18, 0x37, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x47, 0x65, 0x74, 0x4e, 0x66, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x48, 0x00, 0x52, 0x0f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x47, 0x65, 0x74, 0x4e,
	0x66, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x4a, 0x0a, 0x10, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x47,
	0x65, 0x74, 0x4e, 0x66, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x38, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x47, 0x65,
	0x74, 0x4e, 0x66, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x00,
	0x52, 0x10, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x47, 0x65, 0x74, 0x4e, 0x66, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x73, 0x12, 0x5f, 0x0a, 0x17, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x47, 0x65, 0x74,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x39, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x47, 0x65, 0x74, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x69, 0x6d, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x00, 0x52, 0x17, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x47, 0x65, 0x74, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x47, 0x0a, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x3a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x00, 0x52, 0x0e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x07, 0x0a, 0x05,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x42, 0x26, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64,
	0x65, 0x72, 0x61, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6a, 0x61, 0x76, 0x61, 0x50, 0x01, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_query_proto_rawDescOnce sync.Once
	file_query_proto_rawDescData = file_query_proto_rawDesc
)

func file_query_proto_rawDescGZIP() []byte {
	file_query_proto_rawDescOnce.Do(func() {
		file_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_query_proto_rawDescData)
	})
	return file_query_proto_rawDescData
}

var file_query_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_query_proto_goTypes = []interface{}{
	(*Query)(nil),                         // 0: proto.Query
	(*GetByKeyQuery)(nil),                 // 1: proto.GetByKeyQuery
	(*GetBySolidityIDQuery)(nil),          // 2: proto.GetBySolidityIDQuery
	(*ContractCallLocalQuery)(nil),        // 3: proto.ContractCallLocalQuery
	(*ContractGetInfoQuery)(nil),          // 4: proto.ContractGetInfoQuery
	(*ContractGetBytecodeQuery)(nil),      // 5: proto.ContractGetBytecodeQuery
	(*ContractGetRecordsQuery)(nil),       // 6: proto.ContractGetRecordsQuery
	(*CryptoGetAccountBalanceQuery)(nil),  // 7: proto.CryptoGetAccountBalanceQuery
	(*CryptoGetAccountRecordsQuery)(nil),  // 8: proto.CryptoGetAccountRecordsQuery
	(*CryptoGetInfoQuery)(nil),            // 9: proto.CryptoGetInfoQuery
	(*CryptoGetLiveHashQuery)(nil),        // 10: proto.CryptoGetLiveHashQuery
	(*CryptoGetStakersQuery)(nil),         // 11: proto.CryptoGetStakersQuery
	(*FileGetContentsQuery)(nil),          // 12: proto.FileGetContentsQuery
	(*FileGetInfoQuery)(nil),              // 13: proto.FileGetInfoQuery
	(*TransactionGetReceiptQuery)(nil),    // 14: proto.TransactionGetReceiptQuery
	(*TransactionGetRecordQuery)(nil),     // 15: proto.TransactionGetRecordQuery
	(*TransactionGetFastRecordQuery)(nil), // 16: proto.TransactionGetFastRecordQuery
	(*ConsensusGetTopicInfoQuery)(nil),    // 17: proto.ConsensusGetTopicInfoQuery
	(*NetworkGetVersionInfoQuery)(nil),    // 18: proto.NetworkGetVersionInfoQuery
	(*TokenGetInfoQuery)(nil),             // 19: proto.TokenGetInfoQuery
	(*ScheduleGetInfoQuery)(nil),          // 20: proto.ScheduleGetInfoQuery
	(*TokenGetAccountNftInfosQuery)(nil),  // 21: proto.TokenGetAccountNftInfosQuery
	(*TokenGetNftInfoQuery)(nil),          // 22: proto.TokenGetNftInfoQuery
	(*TokenGetNftInfosQuery)(nil),         // 23: proto.TokenGetNftInfosQuery
	(*NetworkGetExecutionTimeQuery)(nil),  // 24: proto.NetworkGetExecutionTimeQuery
	(*GetAccountDetailsQuery)(nil),        // 25: proto.GetAccountDetailsQuery
}
var file_query_proto_depIdxs = []int32{
	1,  // 0: proto.Query.getByKey:type_name -> proto.GetByKeyQuery
	2,  // 1: proto.Query.getBySolidityID:type_name -> proto.GetBySolidityIDQuery
	3,  // 2: proto.Query.contractCallLocal:type_name -> proto.ContractCallLocalQuery
	4,  // 3: proto.Query.contractGetInfo:type_name -> proto.ContractGetInfoQuery
	5,  // 4: proto.Query.contractGetBytecode:type_name -> proto.ContractGetBytecodeQuery
	6,  // 5: proto.Query.ContractGetRecords:type_name -> proto.ContractGetRecordsQuery
	7,  // 6: proto.Query.cryptogetAccountBalance:type_name -> proto.CryptoGetAccountBalanceQuery
	8,  // 7: proto.Query.cryptoGetAccountRecords:type_name -> proto.CryptoGetAccountRecordsQuery
	9,  // 8: proto.Query.cryptoGetInfo:type_name -> proto.CryptoGetInfoQuery
	10, // 9: proto.Query.cryptoGetLiveHash:type_name -> proto.CryptoGetLiveHashQuery
	11, // 10: proto.Query.cryptoGetProxyStakers:type_name -> proto.CryptoGetStakersQuery
	12, // 11: proto.Query.fileGetContents:type_name -> proto.FileGetContentsQuery
	13, // 12: proto.Query.fileGetInfo:type_name -> proto.FileGetInfoQuery
	14, // 13: proto.Query.transactionGetReceipt:type_name -> proto.TransactionGetReceiptQuery
	15, // 14: proto.Query.transactionGetRecord:type_name -> proto.TransactionGetRecordQuery
	16, // 15: proto.Query.transactionGetFastRecord:type_name -> proto.TransactionGetFastRecordQuery
	17, // 16: proto.Query.consensusGetTopicInfo:type_name -> proto.ConsensusGetTopicInfoQuery
	18, // 17: proto.Query.networkGetVersionInfo:type_name -> proto.NetworkGetVersionInfoQuery
	19, // 18: proto.Query.tokenGetInfo:type_name -> proto.TokenGetInfoQuery
	20, // 19: proto.Query.scheduleGetInfo:type_name -> proto.ScheduleGetInfoQuery
	21, // 20: proto.Query.tokenGetAccountNftInfos:type_name -> proto.TokenGetAccountNftInfosQuery
	22, // 21: proto.Query.tokenGetNftInfo:type_name -> proto.TokenGetNftInfoQuery
	23, // 22: proto.Query.tokenGetNftInfos:type_name -> proto.TokenGetNftInfosQuery
	24, // 23: proto.Query.networkGetExecutionTime:type_name -> proto.NetworkGetExecutionTimeQuery
	25, // 24: proto.Query.accountDetails:type_name -> proto.GetAccountDetailsQuery
	25, // [25:25] is the sub-list for method output_type
	25, // [25:25] is the sub-list for method input_type
	25, // [25:25] is the sub-list for extension type_name
	25, // [25:25] is the sub-list for extension extendee
	0,  // [0:25] is the sub-list for field type_name
}

func init() { file_query_proto_init() }
func file_query_proto_init() {
	if File_query_proto != nil {
		return
	}
	file_get_by_key_proto_init()
	file_get_by_solidity_id_proto_init()
	file_contract_call_local_proto_init()
	file_contract_get_info_proto_init()
	file_contract_get_bytecode_proto_init()
	file_contract_get_records_proto_init()
	file_crypto_get_account_balance_proto_init()
	file_crypto_get_account_records_proto_init()
	file_crypto_get_info_proto_init()
	file_crypto_get_live_hash_proto_init()
	file_crypto_get_stakers_proto_init()
	file_file_get_contents_proto_init()
	file_file_get_info_proto_init()
	file_transaction_get_receipt_proto_init()
	file_transaction_get_record_proto_init()
	file_transaction_get_fast_record_proto_init()
	file_consensus_get_topic_info_proto_init()
	file_network_get_version_info_proto_init()
	file_network_get_execution_time_proto_init()
	file_token_get_info_proto_init()
	file_schedule_get_info_proto_init()
	file_token_get_account_nft_infos_proto_init()
	file_token_get_nft_info_proto_init()
	file_token_get_nft_infos_proto_init()
	file_get_account_details_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_query_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Query); i {
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
	file_query_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Query_GetByKey)(nil),
		(*Query_GetBySolidityID)(nil),
		(*Query_ContractCallLocal)(nil),
		(*Query_ContractGetInfo)(nil),
		(*Query_ContractGetBytecode)(nil),
		(*Query_ContractGetRecords)(nil),
		(*Query_CryptogetAccountBalance)(nil),
		(*Query_CryptoGetAccountRecords)(nil),
		(*Query_CryptoGetInfo)(nil),
		(*Query_CryptoGetLiveHash)(nil),
		(*Query_CryptoGetProxyStakers)(nil),
		(*Query_FileGetContents)(nil),
		(*Query_FileGetInfo)(nil),
		(*Query_TransactionGetReceipt)(nil),
		(*Query_TransactionGetRecord)(nil),
		(*Query_TransactionGetFastRecord)(nil),
		(*Query_ConsensusGetTopicInfo)(nil),
		(*Query_NetworkGetVersionInfo)(nil),
		(*Query_TokenGetInfo)(nil),
		(*Query_ScheduleGetInfo)(nil),
		(*Query_TokenGetAccountNftInfos)(nil),
		(*Query_TokenGetNftInfo)(nil),
		(*Query_TokenGetNftInfos)(nil),
		(*Query_NetworkGetExecutionTime)(nil),
		(*Query_AccountDetails)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_query_proto_goTypes,
		DependencyIndexes: file_query_proto_depIdxs,
		MessageInfos:      file_query_proto_msgTypes,
	}.Build()
	File_query_proto = out.File
	file_query_proto_rawDesc = nil
	file_query_proto_goTypes = nil
	file_query_proto_depIdxs = nil
}
