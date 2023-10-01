// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: pkg/proto/repayment.proto

package protogen

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

type Repayment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	LoanId    string                 `protobuf:"bytes,2,opt,name=loan_id,json=loanId,proto3" json:"loan_id,omitempty"`
	Amount    float32                `protobuf:"fixed32,3,opt,name=amount,proto3" json:"amount,omitempty"`
	State     string                 `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty"`
	DueDate   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=due_date,json=dueDate,proto3" json:"due_date,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,87,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Repayment) Reset() {
	*x = Repayment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_repayment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Repayment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Repayment) ProtoMessage() {}

func (x *Repayment) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_repayment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Repayment.ProtoReflect.Descriptor instead.
func (*Repayment) Descriptor() ([]byte, []int) {
	return file_pkg_proto_repayment_proto_rawDescGZIP(), []int{0}
}

func (x *Repayment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Repayment) GetLoanId() string {
	if x != nil {
		return x.LoanId
	}
	return ""
}

func (x *Repayment) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Repayment) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Repayment) GetDueDate() *timestamppb.Timestamp {
	if x != nil {
		return x.DueDate
	}
	return nil
}

func (x *Repayment) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Repayment) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type RepaymentCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LoanId  string                 `protobuf:"bytes,1,opt,name=loan_id,json=loanId,proto3" json:"loan_id,omitempty"`
	Amount  float32                `protobuf:"fixed32,2,opt,name=amount,proto3" json:"amount,omitempty"`
	State   string                 `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	DueDate *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=due_date,json=dueDate,proto3" json:"due_date,omitempty"`
}

func (x *RepaymentCreateRequest) Reset() {
	*x = RepaymentCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_repayment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepaymentCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepaymentCreateRequest) ProtoMessage() {}

func (x *RepaymentCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_repayment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepaymentCreateRequest.ProtoReflect.Descriptor instead.
func (*RepaymentCreateRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_repayment_proto_rawDescGZIP(), []int{1}
}

func (x *RepaymentCreateRequest) GetLoanId() string {
	if x != nil {
		return x.LoanId
	}
	return ""
}

func (x *RepaymentCreateRequest) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *RepaymentCreateRequest) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *RepaymentCreateRequest) GetDueDate() *timestamppb.Timestamp {
	if x != nil {
		return x.DueDate
	}
	return nil
}

type RepaymentCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *Repayment `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Error string     `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Code  int32      `protobuf:"varint,3,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *RepaymentCreateResponse) Reset() {
	*x = RepaymentCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_repayment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepaymentCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepaymentCreateResponse) ProtoMessage() {}

func (x *RepaymentCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_repayment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepaymentCreateResponse.ProtoReflect.Descriptor instead.
func (*RepaymentCreateResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_repayment_proto_rawDescGZIP(), []int{2}
}

func (x *RepaymentCreateResponse) GetData() *Repayment {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *RepaymentCreateResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *RepaymentCreateResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

var File_pkg_proto_repayment_proto protoreflect.FileDescriptor

var file_pkg_proto_repayment_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6c, 0x6f, 0x61,
	0x6e, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x8f, 0x02, 0x0a, 0x09, 0x52, 0x65, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x6c, 0x6f, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6c, 0x6f, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x75, 0x65, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x64, 0x75, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x39,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x57, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x96, 0x01, 0x0a, 0x16, 0x52, 0x65, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x6c, 0x6f, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6c, 0x6f, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x75, 0x65, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x64, 0x75, 0x65, 0x44, 0x61, 0x74, 0x65, 0x22, 0x68, 0x0a,
	0x17, 0x52, 0x65, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6c, 0x6f, 0x61, 0x6e, 0x2e, 0x52, 0x65,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6e, 0x6e, 0x67, 0x75, 0x79, 0x65, 0x6e, 0x35, 0x38,
	0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x6c, 0x61, 0x79, 0x6f,
	0x75, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_repayment_proto_rawDescOnce sync.Once
	file_pkg_proto_repayment_proto_rawDescData = file_pkg_proto_repayment_proto_rawDesc
)

func file_pkg_proto_repayment_proto_rawDescGZIP() []byte {
	file_pkg_proto_repayment_proto_rawDescOnce.Do(func() {
		file_pkg_proto_repayment_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_repayment_proto_rawDescData)
	})
	return file_pkg_proto_repayment_proto_rawDescData
}

var file_pkg_proto_repayment_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_proto_repayment_proto_goTypes = []interface{}{
	(*Repayment)(nil),               // 0: loan.Repayment
	(*RepaymentCreateRequest)(nil),  // 1: loan.RepaymentCreateRequest
	(*RepaymentCreateResponse)(nil), // 2: loan.RepaymentCreateResponse
	(*timestamppb.Timestamp)(nil),   // 3: google.protobuf.Timestamp
}
var file_pkg_proto_repayment_proto_depIdxs = []int32{
	3, // 0: loan.Repayment.due_date:type_name -> google.protobuf.Timestamp
	3, // 1: loan.Repayment.created_at:type_name -> google.protobuf.Timestamp
	3, // 2: loan.Repayment.updated_at:type_name -> google.protobuf.Timestamp
	3, // 3: loan.RepaymentCreateRequest.due_date:type_name -> google.protobuf.Timestamp
	0, // 4: loan.RepaymentCreateResponse.data:type_name -> loan.Repayment
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_pkg_proto_repayment_proto_init() }
func file_pkg_proto_repayment_proto_init() {
	if File_pkg_proto_repayment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_repayment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Repayment); i {
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
		file_pkg_proto_repayment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepaymentCreateRequest); i {
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
		file_pkg_proto_repayment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepaymentCreateResponse); i {
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
			RawDescriptor: file_pkg_proto_repayment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_proto_repayment_proto_goTypes,
		DependencyIndexes: file_pkg_proto_repayment_proto_depIdxs,
		MessageInfos:      file_pkg_proto_repayment_proto_msgTypes,
	}.Build()
	File_pkg_proto_repayment_proto = out.File
	file_pkg_proto_repayment_proto_rawDesc = nil
	file_pkg_proto_repayment_proto_goTypes = nil
	file_pkg_proto_repayment_proto_depIdxs = nil
}
