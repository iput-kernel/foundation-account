// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: account/user/v1/statement.proto

package accountv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StatementRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *StatementRequest) Reset() {
	*x = StatementRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_user_v1_statement_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatementRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatementRequest) ProtoMessage() {}

func (x *StatementRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_user_v1_statement_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatementRequest.ProtoReflect.Descriptor instead.
func (*StatementRequest) Descriptor() ([]byte, []int) {
	return file_account_user_v1_statement_proto_rawDescGZIP(), []int{0}
}

func (x *StatementRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *StatementRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type StatementResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Statements []*Statement `protobuf:"bytes,1,rep,name=statements,proto3" json:"statements,omitempty"`
}

func (x *StatementResponse) Reset() {
	*x = StatementResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_user_v1_statement_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatementResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatementResponse) ProtoMessage() {}

func (x *StatementResponse) ProtoReflect() protoreflect.Message {
	mi := &file_account_user_v1_statement_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatementResponse.ProtoReflect.Descriptor instead.
func (*StatementResponse) Descriptor() ([]byte, []int) {
	return file_account_user_v1_statement_proto_rawDescGZIP(), []int{1}
}

func (x *StatementResponse) GetStatements() []*Statement {
	if x != nil {
		return x.Statements
	}
	return nil
}

type Statement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Amount    string                 `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Type      string                 `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Statement) Reset() {
	*x = Statement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_user_v1_statement_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Statement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Statement) ProtoMessage() {}

func (x *Statement) ProtoReflect() protoreflect.Message {
	mi := &file_account_user_v1_statement_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Statement.ProtoReflect.Descriptor instead.
func (*Statement) Descriptor() ([]byte, []int) {
	return file_account_user_v1_statement_proto_rawDescGZIP(), []int{2}
}

func (x *Statement) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Statement) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *Statement) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Statement) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_account_user_v1_statement_proto protoreflect.FileDescriptor

var file_account_user_v1_statement_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x10, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x4f, 0x0a, 0x11, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x8b, 0x01, 0x0a, 0x09, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x39, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0xd4, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42,
	0x0e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x70,
	0x75, 0x74, 0x2d, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x2f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x55, 0x58, 0xaa, 0x02, 0x0f, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0f, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5c, 0x55, 0x73, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1b, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5c, 0x55, 0x73, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x3a, 0x55, 0x73, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_account_user_v1_statement_proto_rawDescOnce sync.Once
	file_account_user_v1_statement_proto_rawDescData = file_account_user_v1_statement_proto_rawDesc
)

func file_account_user_v1_statement_proto_rawDescGZIP() []byte {
	file_account_user_v1_statement_proto_rawDescOnce.Do(func() {
		file_account_user_v1_statement_proto_rawDescData = protoimpl.X.CompressGZIP(file_account_user_v1_statement_proto_rawDescData)
	})
	return file_account_user_v1_statement_proto_rawDescData
}

var file_account_user_v1_statement_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_account_user_v1_statement_proto_goTypes = []any{
	(*StatementRequest)(nil),      // 0: account.user.v1.StatementRequest
	(*StatementResponse)(nil),     // 1: account.user.v1.StatementResponse
	(*Statement)(nil),             // 2: account.user.v1.Statement
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_account_user_v1_statement_proto_depIdxs = []int32{
	2, // 0: account.user.v1.StatementResponse.statements:type_name -> account.user.v1.Statement
	3, // 1: account.user.v1.Statement.created_at:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_account_user_v1_statement_proto_init() }
func file_account_user_v1_statement_proto_init() {
	if File_account_user_v1_statement_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_account_user_v1_statement_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*StatementRequest); i {
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
		file_account_user_v1_statement_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*StatementResponse); i {
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
		file_account_user_v1_statement_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Statement); i {
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
			RawDescriptor: file_account_user_v1_statement_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_account_user_v1_statement_proto_goTypes,
		DependencyIndexes: file_account_user_v1_statement_proto_depIdxs,
		MessageInfos:      file_account_user_v1_statement_proto_msgTypes,
	}.Build()
	File_account_user_v1_statement_proto = out.File
	file_account_user_v1_statement_proto_rawDesc = nil
	file_account_user_v1_statement_proto_goTypes = nil
	file_account_user_v1_statement_proto_depIdxs = nil
}
