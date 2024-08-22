// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.2
// source: bank.proto

package bank_v1

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

type CreateAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jwt         string `protobuf:"bytes,1,opt,name=jwt,proto3" json:"jwt,omitempty"`
	FullName    string `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Citizenship string `protobuf:"bytes,3,opt,name=citizenship,proto3" json:"citizenship,omitempty"`
	Balance     int64  `protobuf:"varint,4,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *CreateAccountRequest) Reset() {
	*x = CreateAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountRequest) ProtoMessage() {}

func (x *CreateAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bank_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountRequest.ProtoReflect.Descriptor instead.
func (*CreateAccountRequest) Descriptor() ([]byte, []int) {
	return file_bank_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAccountRequest) GetJwt() string {
	if x != nil {
		return x.Jwt
	}
	return ""
}

func (x *CreateAccountRequest) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *CreateAccountRequest) GetCitizenship() string {
	if x != nil {
		return x.Citizenship
	}
	return ""
}

func (x *CreateAccountRequest) GetBalance() int64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

type CreateAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId int64 `protobuf:"varint,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"` // ID of created account
	Balance   int64 `protobuf:"varint,2,opt,name=balance,proto3" json:"balance,omitempty"`                      // Created account balance
}

func (x *CreateAccountResponse) Reset() {
	*x = CreateAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountResponse) ProtoMessage() {}

func (x *CreateAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bank_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountResponse.ProtoReflect.Descriptor instead.
func (*CreateAccountResponse) Descriptor() ([]byte, []int) {
	return file_bank_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAccountResponse) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *CreateAccountResponse) GetBalance() int64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

type AccountTopUpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jwt         string `protobuf:"bytes,1,opt,name=jwt,proto3" json:"jwt,omitempty"`
	AccountId   int64  `protobuf:"varint,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	TopUpAmount int64  `protobuf:"varint,3,opt,name=top_up_amount,json=topUpAmount,proto3" json:"top_up_amount,omitempty"`
}

func (x *AccountTopUpRequest) Reset() {
	*x = AccountTopUpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountTopUpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountTopUpRequest) ProtoMessage() {}

func (x *AccountTopUpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bank_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountTopUpRequest.ProtoReflect.Descriptor instead.
func (*AccountTopUpRequest) Descriptor() ([]byte, []int) {
	return file_bank_proto_rawDescGZIP(), []int{2}
}

func (x *AccountTopUpRequest) GetJwt() string {
	if x != nil {
		return x.Jwt
	}
	return ""
}

func (x *AccountTopUpRequest) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *AccountTopUpRequest) GetTopUpAmount() int64 {
	if x != nil {
		return x.TopUpAmount
	}
	return 0
}

type AccountTopUpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Balance int64 `protobuf:"varint,1,opt,name=balance,proto3" json:"balance,omitempty"` // Account balance after top up
}

func (x *AccountTopUpResponse) Reset() {
	*x = AccountTopUpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountTopUpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountTopUpResponse) ProtoMessage() {}

func (x *AccountTopUpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bank_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountTopUpResponse.ProtoReflect.Descriptor instead.
func (*AccountTopUpResponse) Descriptor() ([]byte, []int) {
	return file_bank_proto_rawDescGZIP(), []int{3}
}

func (x *AccountTopUpResponse) GetBalance() int64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

type AccountWithdrawRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jwt            string `protobuf:"bytes,1,opt,name=jwt,proto3" json:"jwt,omitempty"`
	AccountId      int64  `protobuf:"varint,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	WithdrawAmount int64  `protobuf:"varint,3,opt,name=withdraw_amount,json=withdrawAmount,proto3" json:"withdraw_amount,omitempty"`
}

func (x *AccountWithdrawRequest) Reset() {
	*x = AccountWithdrawRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountWithdrawRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountWithdrawRequest) ProtoMessage() {}

func (x *AccountWithdrawRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bank_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountWithdrawRequest.ProtoReflect.Descriptor instead.
func (*AccountWithdrawRequest) Descriptor() ([]byte, []int) {
	return file_bank_proto_rawDescGZIP(), []int{4}
}

func (x *AccountWithdrawRequest) GetJwt() string {
	if x != nil {
		return x.Jwt
	}
	return ""
}

func (x *AccountWithdrawRequest) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *AccountWithdrawRequest) GetWithdrawAmount() int64 {
	if x != nil {
		return x.WithdrawAmount
	}
	return 0
}

type AccountWithdrawResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Balance int64 `protobuf:"varint,1,opt,name=balance,proto3" json:"balance,omitempty"` // Account balance after withdraw
}

func (x *AccountWithdrawResponse) Reset() {
	*x = AccountWithdrawResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountWithdrawResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountWithdrawResponse) ProtoMessage() {}

func (x *AccountWithdrawResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bank_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountWithdrawResponse.ProtoReflect.Descriptor instead.
func (*AccountWithdrawResponse) Descriptor() ([]byte, []int) {
	return file_bank_proto_rawDescGZIP(), []int{5}
}

func (x *AccountWithdrawResponse) GetBalance() int64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

type AccountTransferRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jwt                  string `protobuf:"bytes,1,opt,name=jwt,proto3" json:"jwt,omitempty"`
	WriteOffAccountId    int64  `protobuf:"varint,2,opt,name=write_off_account_id,json=writeOffAccountId,proto3" json:"write_off_account_id,omitempty"`
	BeneficiaryAccountId int64  `protobuf:"varint,3,opt,name=beneficiary_account_id,json=beneficiaryAccountId,proto3" json:"beneficiary_account_id,omitempty"`
	TransferAmount       int64  `protobuf:"varint,4,opt,name=transfer_amount,json=transferAmount,proto3" json:"transfer_amount,omitempty"`
}

func (x *AccountTransferRequest) Reset() {
	*x = AccountTransferRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountTransferRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountTransferRequest) ProtoMessage() {}

func (x *AccountTransferRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bank_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountTransferRequest.ProtoReflect.Descriptor instead.
func (*AccountTransferRequest) Descriptor() ([]byte, []int) {
	return file_bank_proto_rawDescGZIP(), []int{6}
}

func (x *AccountTransferRequest) GetJwt() string {
	if x != nil {
		return x.Jwt
	}
	return ""
}

func (x *AccountTransferRequest) GetWriteOffAccountId() int64 {
	if x != nil {
		return x.WriteOffAccountId
	}
	return 0
}

func (x *AccountTransferRequest) GetBeneficiaryAccountId() int64 {
	if x != nil {
		return x.BeneficiaryAccountId
	}
	return 0
}

func (x *AccountTransferRequest) GetTransferAmount() int64 {
	if x != nil {
		return x.TransferAmount
	}
	return 0
}

type AccountTransferResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WriteOffAccountBalance    int64 `protobuf:"varint,1,opt,name=write_off_account_balance,json=writeOffAccountBalance,proto3" json:"write_off_account_balance,omitempty"`
	BeneficiaryAccountBalance int64 `protobuf:"varint,2,opt,name=beneficiary_account_balance,json=beneficiaryAccountBalance,proto3" json:"beneficiary_account_balance,omitempty"`
}

func (x *AccountTransferResponse) Reset() {
	*x = AccountTransferResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountTransferResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountTransferResponse) ProtoMessage() {}

func (x *AccountTransferResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bank_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountTransferResponse.ProtoReflect.Descriptor instead.
func (*AccountTransferResponse) Descriptor() ([]byte, []int) {
	return file_bank_proto_rawDescGZIP(), []int{7}
}

func (x *AccountTransferResponse) GetWriteOffAccountBalance() int64 {
	if x != nil {
		return x.WriteOffAccountBalance
	}
	return 0
}

func (x *AccountTransferResponse) GetBeneficiaryAccountBalance() int64 {
	if x != nil {
		return x.BeneficiaryAccountBalance
	}
	return 0
}

type AccountLockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId int64 `protobuf:"varint,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"` // Account ID to lock
}

func (x *AccountLockRequest) Reset() {
	*x = AccountLockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountLockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountLockRequest) ProtoMessage() {}

func (x *AccountLockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bank_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountLockRequest.ProtoReflect.Descriptor instead.
func (*AccountLockRequest) Descriptor() ([]byte, []int) {
	return file_bank_proto_rawDescGZIP(), []int{8}
}

func (x *AccountLockRequest) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

type AccountLockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AccountLockResponse) Reset() {
	*x = AccountLockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountLockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountLockResponse) ProtoMessage() {}

func (x *AccountLockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bank_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountLockResponse.ProtoReflect.Descriptor instead.
func (*AccountLockResponse) Descriptor() ([]byte, []int) {
	return file_bank_proto_rawDescGZIP(), []int{9}
}

var File_bank_proto protoreflect.FileDescriptor

var file_bank_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x62, 0x61,
	0x6e, 0x6b, 0x22, 0x81, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6a,
	0x77, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6a, 0x77, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x69,
	0x74, 0x69, 0x7a, 0x65, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x63, 0x69, 0x74, 0x69, 0x7a, 0x65, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x12, 0x18, 0x0a, 0x07,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x62,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x50, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x6a, 0x0a, 0x13, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x54, 0x6f, 0x70, 0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x6a, 0x77, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6a, 0x77,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x22, 0x0a, 0x0d, 0x74, 0x6f, 0x70, 0x5f, 0x75, 0x70, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x6f, 0x70, 0x55, 0x70, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x30, 0x0a, 0x14, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54,
	0x6f, 0x70, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x62,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x72, 0x0a, 0x16, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x6a, 0x77, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6a,
	0x77, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x27, 0x0a, 0x0f, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x5f, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x77, 0x69, 0x74, 0x68,
	0x64, 0x72, 0x61, 0x77, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x33, 0x0a, 0x17, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22,
	0xba, 0x01, 0x0a, 0x16, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6a, 0x77,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6a, 0x77, 0x74, 0x12, 0x2f, 0x0a, 0x14,
	0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x6f, 0x66, 0x66, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x77, 0x72, 0x69, 0x74,
	0x65, 0x4f, 0x66, 0x66, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x34, 0x0a,
	0x16, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x63, 0x69, 0x61, 0x72, 0x79, 0x5f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x62,
	0x65, 0x6e, 0x65, 0x66, 0x69, 0x63, 0x69, 0x61, 0x72, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x94, 0x01, 0x0a,
	0x17, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x19, 0x77, 0x72, 0x69, 0x74,
	0x65, 0x5f, 0x6f, 0x66, 0x66, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x62, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x16, 0x77, 0x72, 0x69,
	0x74, 0x65, 0x4f, 0x66, 0x66, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x1b, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x63, 0x69, 0x61,
	0x72, 0x79, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x19, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69,
	0x63, 0x69, 0x61, 0x72, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x22, 0x33, 0x0a, 0x12, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x6f,
	0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0xfb, 0x02, 0x0a, 0x04, 0x42, 0x61, 0x6e, 0x6b, 0x12, 0x48, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x2e, 0x62, 0x61, 0x6e, 0x6b,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x45, 0x0a, 0x0c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x6f, 0x70,
	0x55, 0x70, 0x12, 0x19, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x54, 0x6f, 0x70, 0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x6f, 0x70, 0x55,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0f, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x12, 0x1c, 0x2e, 0x62,
	0x61, 0x6e, 0x6b, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x57, 0x69, 0x74, 0x68, 0x64,
	0x72, 0x61, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x62, 0x61, 0x6e,
	0x6b, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61,
	0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0f, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x62,
	0x61, 0x6e, 0x6b, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x62, 0x61, 0x6e,
	0x6b, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0b, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x6f, 0x63, 0x6b, 0x12, 0x18, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1f, 0x5a,
	0x1d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x3b, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bank_proto_rawDescOnce sync.Once
	file_bank_proto_rawDescData = file_bank_proto_rawDesc
)

func file_bank_proto_rawDescGZIP() []byte {
	file_bank_proto_rawDescOnce.Do(func() {
		file_bank_proto_rawDescData = protoimpl.X.CompressGZIP(file_bank_proto_rawDescData)
	})
	return file_bank_proto_rawDescData
}

var file_bank_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_bank_proto_goTypes = []interface{}{
	(*CreateAccountRequest)(nil),    // 0: bank.CreateAccountRequest
	(*CreateAccountResponse)(nil),   // 1: bank.CreateAccountResponse
	(*AccountTopUpRequest)(nil),     // 2: bank.AccountTopUpRequest
	(*AccountTopUpResponse)(nil),    // 3: bank.AccountTopUpResponse
	(*AccountWithdrawRequest)(nil),  // 4: bank.AccountWithdrawRequest
	(*AccountWithdrawResponse)(nil), // 5: bank.AccountWithdrawResponse
	(*AccountTransferRequest)(nil),  // 6: bank.AccountTransferRequest
	(*AccountTransferResponse)(nil), // 7: bank.AccountTransferResponse
	(*AccountLockRequest)(nil),      // 8: bank.AccountLockRequest
	(*AccountLockResponse)(nil),     // 9: bank.AccountLockResponse
}
var file_bank_proto_depIdxs = []int32{
	0, // 0: bank.Bank.CreateAccount:input_type -> bank.CreateAccountRequest
	2, // 1: bank.Bank.AccountTopUp:input_type -> bank.AccountTopUpRequest
	4, // 2: bank.Bank.AccountWithdraw:input_type -> bank.AccountWithdrawRequest
	6, // 3: bank.Bank.AccountTransfer:input_type -> bank.AccountTransferRequest
	8, // 4: bank.Bank.AccountLock:input_type -> bank.AccountLockRequest
	1, // 5: bank.Bank.CreateAccount:output_type -> bank.CreateAccountResponse
	3, // 6: bank.Bank.AccountTopUp:output_type -> bank.AccountTopUpResponse
	5, // 7: bank.Bank.AccountWithdraw:output_type -> bank.AccountWithdrawResponse
	7, // 8: bank.Bank.AccountTransfer:output_type -> bank.AccountTransferResponse
	9, // 9: bank.Bank.AccountLock:output_type -> bank.AccountLockResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bank_proto_init() }
func file_bank_proto_init() {
	if File_bank_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bank_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccountRequest); i {
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
		file_bank_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccountResponse); i {
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
		file_bank_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountTopUpRequest); i {
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
		file_bank_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountTopUpResponse); i {
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
		file_bank_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountWithdrawRequest); i {
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
		file_bank_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountWithdrawResponse); i {
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
		file_bank_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountTransferRequest); i {
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
		file_bank_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountTransferResponse); i {
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
		file_bank_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountLockRequest); i {
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
		file_bank_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountLockResponse); i {
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
			RawDescriptor: file_bank_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bank_proto_goTypes,
		DependencyIndexes: file_bank_proto_depIdxs,
		MessageInfos:      file_bank_proto_msgTypes,
	}.Build()
	File_bank_proto = out.File
	file_bank_proto_rawDesc = nil
	file_bank_proto_goTypes = nil
	file_bank_proto_depIdxs = nil
}
