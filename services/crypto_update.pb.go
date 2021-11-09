// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: crypto_update.proto

package services

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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
// Change properties for the given account. Any null field is ignored (left unchanged). This
// transaction must be signed by the existing key for this account. If the transaction is changing
// the key field, then the transaction must be signed by both the old key (from before the change)
// and the new key. The old key must sign for security. The new key must sign as a safeguard to
// avoid accidentally changing to an invalid key, and then having no way to recover.
type CryptoUpdateTransactionBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// The account ID which is being updated in this transaction
	AccountIDToUpdate *AccountID `protobuf:"bytes,2,opt,name=accountIDToUpdate,proto3" json:"accountIDToUpdate,omitempty"`
	//*
	// The new key
	Key *Key `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	//*
	// ID of the account to which this account is proxy staked. If proxyAccountID is null, or is an
	// invalid account, or is an account that isn't a node, then this account is automatically proxy
	// staked to a node chosen by the network, but without earning payments. If the proxyAccountID
	// account refuses to accept proxy staking , or if it is not currently running a node, then it
	// will behave as if proxyAccountID was null.
	ProxyAccountID *AccountID `protobuf:"bytes,4,opt,name=proxyAccountID,proto3" json:"proxyAccountID,omitempty"`
	//*
	// [Deprecated]. Payments earned from proxy staking are shared between the node and this
	// account, with proxyFraction / 10000 going to this account
	//
	// Deprecated: Do not use.
	ProxyFraction int32 `protobuf:"varint,5,opt,name=proxyFraction,proto3" json:"proxyFraction,omitempty"`
	// Types that are assignable to SendRecordThresholdField:
	//	*CryptoUpdateTransactionBody_SendRecordThreshold
	//	*CryptoUpdateTransactionBody_SendRecordThresholdWrapper
	SendRecordThresholdField isCryptoUpdateTransactionBody_SendRecordThresholdField `protobuf_oneof:"sendRecordThresholdField"`
	// Types that are assignable to ReceiveRecordThresholdField:
	//	*CryptoUpdateTransactionBody_ReceiveRecordThreshold
	//	*CryptoUpdateTransactionBody_ReceiveRecordThresholdWrapper
	ReceiveRecordThresholdField isCryptoUpdateTransactionBody_ReceiveRecordThresholdField `protobuf_oneof:"receiveRecordThresholdField"`
	//*
	// The duration in which it will automatically extend the expiration period. If it doesn't have
	// enough balance, it extends as long as possible. If it is empty when it expires, then it is
	// deleted.
	AutoRenewPeriod *Duration `protobuf:"bytes,8,opt,name=autoRenewPeriod,proto3" json:"autoRenewPeriod,omitempty"`
	//*
	// The new expiration time to extend to (ignored if equal to or before the current one)
	ExpirationTime *Timestamp `protobuf:"bytes,9,opt,name=expirationTime,proto3" json:"expirationTime,omitempty"`
	// Types that are assignable to ReceiverSigRequiredField:
	//	*CryptoUpdateTransactionBody_ReceiverSigRequired
	//	*CryptoUpdateTransactionBody_ReceiverSigRequiredWrapper
	ReceiverSigRequiredField isCryptoUpdateTransactionBody_ReceiverSigRequiredField `protobuf_oneof:"receiverSigRequiredField"`
	//*
	// If set, the new memo to be associated with the account (UTF-8 encoding max 100 bytes)
	Memo *wrapperspb.StringValue `protobuf:"bytes,14,opt,name=memo,proto3" json:"memo,omitempty"`
	//*
	// The maximum number of tokens that an Account can be implicitly associated with. Up to a 1000
	// including implicit and explicit associations.
	MaxAutomaticTokenAssociations *wrapperspb.Int32Value `protobuf:"bytes,15,opt,name=max_automatic_token_associations,json=maxAutomaticTokenAssociations,proto3" json:"max_automatic_token_associations,omitempty"`
}

func (x *CryptoUpdateTransactionBody) Reset() {
	*x = CryptoUpdateTransactionBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crypto_update_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CryptoUpdateTransactionBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CryptoUpdateTransactionBody) ProtoMessage() {}

func (x *CryptoUpdateTransactionBody) ProtoReflect() protoreflect.Message {
	mi := &file_crypto_update_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CryptoUpdateTransactionBody.ProtoReflect.Descriptor instead.
func (*CryptoUpdateTransactionBody) Descriptor() ([]byte, []int) {
	return file_crypto_update_proto_rawDescGZIP(), []int{0}
}

func (x *CryptoUpdateTransactionBody) GetAccountIDToUpdate() *AccountID {
	if x != nil {
		return x.AccountIDToUpdate
	}
	return nil
}

func (x *CryptoUpdateTransactionBody) GetKey() *Key {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *CryptoUpdateTransactionBody) GetProxyAccountID() *AccountID {
	if x != nil {
		return x.ProxyAccountID
	}
	return nil
}

// Deprecated: Do not use.
func (x *CryptoUpdateTransactionBody) GetProxyFraction() int32 {
	if x != nil {
		return x.ProxyFraction
	}
	return 0
}

func (m *CryptoUpdateTransactionBody) GetSendRecordThresholdField() isCryptoUpdateTransactionBody_SendRecordThresholdField {
	if m != nil {
		return m.SendRecordThresholdField
	}
	return nil
}

// Deprecated: Do not use.
func (x *CryptoUpdateTransactionBody) GetSendRecordThreshold() uint64 {
	if x, ok := x.GetSendRecordThresholdField().(*CryptoUpdateTransactionBody_SendRecordThreshold); ok {
		return x.SendRecordThreshold
	}
	return 0
}

// Deprecated: Do not use.
func (x *CryptoUpdateTransactionBody) GetSendRecordThresholdWrapper() *wrapperspb.UInt64Value {
	if x, ok := x.GetSendRecordThresholdField().(*CryptoUpdateTransactionBody_SendRecordThresholdWrapper); ok {
		return x.SendRecordThresholdWrapper
	}
	return nil
}

func (m *CryptoUpdateTransactionBody) GetReceiveRecordThresholdField() isCryptoUpdateTransactionBody_ReceiveRecordThresholdField {
	if m != nil {
		return m.ReceiveRecordThresholdField
	}
	return nil
}

// Deprecated: Do not use.
func (x *CryptoUpdateTransactionBody) GetReceiveRecordThreshold() uint64 {
	if x, ok := x.GetReceiveRecordThresholdField().(*CryptoUpdateTransactionBody_ReceiveRecordThreshold); ok {
		return x.ReceiveRecordThreshold
	}
	return 0
}

// Deprecated: Do not use.
func (x *CryptoUpdateTransactionBody) GetReceiveRecordThresholdWrapper() *wrapperspb.UInt64Value {
	if x, ok := x.GetReceiveRecordThresholdField().(*CryptoUpdateTransactionBody_ReceiveRecordThresholdWrapper); ok {
		return x.ReceiveRecordThresholdWrapper
	}
	return nil
}

func (x *CryptoUpdateTransactionBody) GetAutoRenewPeriod() *Duration {
	if x != nil {
		return x.AutoRenewPeriod
	}
	return nil
}

func (x *CryptoUpdateTransactionBody) GetExpirationTime() *Timestamp {
	if x != nil {
		return x.ExpirationTime
	}
	return nil
}

func (m *CryptoUpdateTransactionBody) GetReceiverSigRequiredField() isCryptoUpdateTransactionBody_ReceiverSigRequiredField {
	if m != nil {
		return m.ReceiverSigRequiredField
	}
	return nil
}

// Deprecated: Do not use.
func (x *CryptoUpdateTransactionBody) GetReceiverSigRequired() bool {
	if x, ok := x.GetReceiverSigRequiredField().(*CryptoUpdateTransactionBody_ReceiverSigRequired); ok {
		return x.ReceiverSigRequired
	}
	return false
}

func (x *CryptoUpdateTransactionBody) GetReceiverSigRequiredWrapper() *wrapperspb.BoolValue {
	if x, ok := x.GetReceiverSigRequiredField().(*CryptoUpdateTransactionBody_ReceiverSigRequiredWrapper); ok {
		return x.ReceiverSigRequiredWrapper
	}
	return nil
}

func (x *CryptoUpdateTransactionBody) GetMemo() *wrapperspb.StringValue {
	if x != nil {
		return x.Memo
	}
	return nil
}

func (x *CryptoUpdateTransactionBody) GetMaxAutomaticTokenAssociations() *wrapperspb.Int32Value {
	if x != nil {
		return x.MaxAutomaticTokenAssociations
	}
	return nil
}

type isCryptoUpdateTransactionBody_SendRecordThresholdField interface {
	isCryptoUpdateTransactionBody_SendRecordThresholdField()
}

type CryptoUpdateTransactionBody_SendRecordThreshold struct {
	//*
	// [Deprecated]. The new threshold amount (in tinybars) for which an account record is
	// created for any send/withdraw transaction
	//
	// Deprecated: Do not use.
	SendRecordThreshold uint64 `protobuf:"varint,6,opt,name=sendRecordThreshold,proto3,oneof"`
}

type CryptoUpdateTransactionBody_SendRecordThresholdWrapper struct {
	//*
	// [Deprecated]. The new threshold amount (in tinybars) for which an account record is
	// created for any send/withdraw transaction
	//
	// Deprecated: Do not use.
	SendRecordThresholdWrapper *wrapperspb.UInt64Value `protobuf:"bytes,11,opt,name=sendRecordThresholdWrapper,proto3,oneof"`
}

func (*CryptoUpdateTransactionBody_SendRecordThreshold) isCryptoUpdateTransactionBody_SendRecordThresholdField() {
}

func (*CryptoUpdateTransactionBody_SendRecordThresholdWrapper) isCryptoUpdateTransactionBody_SendRecordThresholdField() {
}

type isCryptoUpdateTransactionBody_ReceiveRecordThresholdField interface {
	isCryptoUpdateTransactionBody_ReceiveRecordThresholdField()
}

type CryptoUpdateTransactionBody_ReceiveRecordThreshold struct {
	//*
	// [Deprecated]. The new threshold amount (in tinybars) for which an account record is
	// created for any receive/deposit transaction.
	//
	// Deprecated: Do not use.
	ReceiveRecordThreshold uint64 `protobuf:"varint,7,opt,name=receiveRecordThreshold,proto3,oneof"`
}

type CryptoUpdateTransactionBody_ReceiveRecordThresholdWrapper struct {
	//*
	// [Deprecated]. The new threshold amount (in tinybars) for which an account record is
	// created for any receive/deposit transaction.
	//
	// Deprecated: Do not use.
	ReceiveRecordThresholdWrapper *wrapperspb.UInt64Value `protobuf:"bytes,12,opt,name=receiveRecordThresholdWrapper,proto3,oneof"`
}

func (*CryptoUpdateTransactionBody_ReceiveRecordThreshold) isCryptoUpdateTransactionBody_ReceiveRecordThresholdField() {
}

func (*CryptoUpdateTransactionBody_ReceiveRecordThresholdWrapper) isCryptoUpdateTransactionBody_ReceiveRecordThresholdField() {
}

type isCryptoUpdateTransactionBody_ReceiverSigRequiredField interface {
	isCryptoUpdateTransactionBody_ReceiverSigRequiredField()
}

type CryptoUpdateTransactionBody_ReceiverSigRequired struct {
	//*
	// [Deprecated] Do NOT use this field to set a false value because the server cannot
	// distinguish from the default value. Use receiverSigRequiredWrapper field for this
	// purpose.
	//
	// Deprecated: Do not use.
	ReceiverSigRequired bool `protobuf:"varint,10,opt,name=receiverSigRequired,proto3,oneof"`
}

type CryptoUpdateTransactionBody_ReceiverSigRequiredWrapper struct {
	//*
	// If true, this account's key must sign any transaction depositing into this account (in
	// addition to all withdrawals)
	ReceiverSigRequiredWrapper *wrapperspb.BoolValue `protobuf:"bytes,13,opt,name=receiverSigRequiredWrapper,proto3,oneof"`
}

func (*CryptoUpdateTransactionBody_ReceiverSigRequired) isCryptoUpdateTransactionBody_ReceiverSigRequiredField() {
}

func (*CryptoUpdateTransactionBody_ReceiverSigRequiredWrapper) isCryptoUpdateTransactionBody_ReceiverSigRequiredField() {
}

var File_crypto_update_proto protoreflect.FileDescriptor

var file_crypto_update_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x62, 0x61,
	0x73, 0x69, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0e, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x9d, 0x08, 0x0a, 0x1b, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6f, 0x64, 0x79,
	0x12, 0x3e, 0x0a, 0x11, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x54, 0x6f, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x11, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x54, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x1c, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x38,
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x28, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x46, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x42,
	0x02, 0x18, 0x01, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x46, 0x72, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x13, 0x73, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x42,
	0x02, 0x18, 0x01, 0x48, 0x00, 0x52, 0x13, 0x73, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x62, 0x0a, 0x1a, 0x73, 0x65,
	0x6e, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c,
	0x64, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x02, 0x18, 0x01,
	0x48, 0x00, 0x52, 0x1a, 0x73, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x68,
	0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x12, 0x3c,
	0x0a, 0x16, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54,
	0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02,
	0x18, 0x01, 0x48, 0x01, 0x52, 0x16, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x68, 0x0a, 0x1d,
	0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x68, 0x72,
	0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x42, 0x02, 0x18, 0x01, 0x48, 0x01, 0x52, 0x1d, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x57,
	0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x0f, 0x61, 0x75, 0x74, 0x6f, 0x52, 0x65,
	0x6e, 0x65, 0x77, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0f, 0x61, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x50, 0x65, 0x72, 0x69, 0x6f,
	0x64, 0x12, 0x38, 0x0a, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x13, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x53, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x42, 0x02, 0x18, 0x01, 0x48, 0x02, 0x52, 0x13,
	0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x53, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x64, 0x12, 0x5c, 0x0a, 0x1a, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x53,
	0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x48, 0x02, 0x52, 0x1a, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x53,
	0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x12, 0x30, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x6d,
	0x65, 0x6d, 0x6f, 0x12, 0x64, 0x0a, 0x20, 0x6d, 0x61, 0x78, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x61, 0x74, 0x69, 0x63, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x61, 0x73, 0x73, 0x6f, 0x63,
	0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x1d, 0x6d, 0x61, 0x78, 0x41,
	0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x73, 0x73,
	0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x1a, 0x0a, 0x18, 0x73, 0x65, 0x6e,
	0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x1d, 0x0a, 0x1b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x42, 0x1a, 0x0a, 0x18, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72,
	0x53, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x42, 0x26, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x68, 0x61,
	0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x6a, 0x61, 0x76, 0x61, 0x50, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_crypto_update_proto_rawDescOnce sync.Once
	file_crypto_update_proto_rawDescData = file_crypto_update_proto_rawDesc
)

func file_crypto_update_proto_rawDescGZIP() []byte {
	file_crypto_update_proto_rawDescOnce.Do(func() {
		file_crypto_update_proto_rawDescData = protoimpl.X.CompressGZIP(file_crypto_update_proto_rawDescData)
	})
	return file_crypto_update_proto_rawDescData
}

var file_crypto_update_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_crypto_update_proto_goTypes = []interface{}{
	(*CryptoUpdateTransactionBody)(nil), // 0: proto.CryptoUpdateTransactionBody
	(*AccountID)(nil),                   // 1: proto.AccountID
	(*Key)(nil),                         // 2: proto.Key
	(*wrapperspb.UInt64Value)(nil),      // 3: google.protobuf.UInt64Value
	(*Duration)(nil),                    // 4: proto.Duration
	(*Timestamp)(nil),                   // 5: proto.Timestamp
	(*wrapperspb.BoolValue)(nil),        // 6: google.protobuf.BoolValue
	(*wrapperspb.StringValue)(nil),      // 7: google.protobuf.StringValue
	(*wrapperspb.Int32Value)(nil),       // 8: google.protobuf.Int32Value
}
var file_crypto_update_proto_depIdxs = []int32{
	1,  // 0: proto.CryptoUpdateTransactionBody.accountIDToUpdate:type_name -> proto.AccountID
	2,  // 1: proto.CryptoUpdateTransactionBody.key:type_name -> proto.Key
	1,  // 2: proto.CryptoUpdateTransactionBody.proxyAccountID:type_name -> proto.AccountID
	3,  // 3: proto.CryptoUpdateTransactionBody.sendRecordThresholdWrapper:type_name -> google.protobuf.UInt64Value
	3,  // 4: proto.CryptoUpdateTransactionBody.receiveRecordThresholdWrapper:type_name -> google.protobuf.UInt64Value
	4,  // 5: proto.CryptoUpdateTransactionBody.autoRenewPeriod:type_name -> proto.Duration
	5,  // 6: proto.CryptoUpdateTransactionBody.expirationTime:type_name -> proto.Timestamp
	6,  // 7: proto.CryptoUpdateTransactionBody.receiverSigRequiredWrapper:type_name -> google.protobuf.BoolValue
	7,  // 8: proto.CryptoUpdateTransactionBody.memo:type_name -> google.protobuf.StringValue
	8,  // 9: proto.CryptoUpdateTransactionBody.max_automatic_token_associations:type_name -> google.protobuf.Int32Value
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_crypto_update_proto_init() }
func file_crypto_update_proto_init() {
	if File_crypto_update_proto != nil {
		return
	}
	file_basic_types_proto_init()
	file_duration_proto_init()
	file_timestamp_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_crypto_update_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CryptoUpdateTransactionBody); i {
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
	file_crypto_update_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*CryptoUpdateTransactionBody_SendRecordThreshold)(nil),
		(*CryptoUpdateTransactionBody_SendRecordThresholdWrapper)(nil),
		(*CryptoUpdateTransactionBody_ReceiveRecordThreshold)(nil),
		(*CryptoUpdateTransactionBody_ReceiveRecordThresholdWrapper)(nil),
		(*CryptoUpdateTransactionBody_ReceiverSigRequired)(nil),
		(*CryptoUpdateTransactionBody_ReceiverSigRequiredWrapper)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_crypto_update_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_crypto_update_proto_goTypes,
		DependencyIndexes: file_crypto_update_proto_depIdxs,
		MessageInfos:      file_crypto_update_proto_msgTypes,
	}.Build()
	File_crypto_update_proto = out.File
	file_crypto_update_proto_rawDesc = nil
	file_crypto_update_proto_goTypes = nil
	file_crypto_update_proto_depIdxs = nil
}