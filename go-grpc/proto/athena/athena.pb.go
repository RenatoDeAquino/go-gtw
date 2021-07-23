// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: athena/athena.proto

package athena

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type BaseStatus_EnumStatus int32

const (
	BaseStatus_SUCESSO BaseStatus_EnumStatus = 0
	BaseStatus_NEGOCIO BaseStatus_EnumStatus = 1
	BaseStatus_SISTEMA BaseStatus_EnumStatus = 2
)

// Enum value maps for BaseStatus_EnumStatus.
var (
	BaseStatus_EnumStatus_name = map[int32]string{
		0: "SUCESSO",
		1: "NEGOCIO",
		2: "SISTEMA",
	}
	BaseStatus_EnumStatus_value = map[string]int32{
		"SUCESSO": 0,
		"NEGOCIO": 1,
		"SISTEMA": 2,
	}
)

func (x BaseStatus_EnumStatus) Enum() *BaseStatus_EnumStatus {
	p := new(BaseStatus_EnumStatus)
	*p = x
	return p
}

func (x BaseStatus_EnumStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BaseStatus_EnumStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_athena_athena_proto_enumTypes[0].Descriptor()
}

func (BaseStatus_EnumStatus) Type() protoreflect.EnumType {
	return &file_athena_athena_proto_enumTypes[0]
}

func (x BaseStatus_EnumStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BaseStatus_EnumStatus.Descriptor instead.
func (BaseStatus_EnumStatus) EnumDescriptor() ([]byte, []int) {
	return file_athena_athena_proto_rawDescGZIP(), []int{1, 0}
}

type AccountProduct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Banco   int64  `protobuf:"varint,1,opt,name=Banco,proto3" json:"Banco,omitempty"`
	Produto int64  `protobuf:"varint,2,opt,name=Produto,proto3" json:"Produto,omitempty"`
	Agencia int64  `protobuf:"varint,3,opt,name=Agencia,proto3" json:"Agencia,omitempty"`
	Conta   int64  `protobuf:"varint,4,opt,name=Conta,proto3" json:"Conta,omitempty"`
	CPF     string `protobuf:"bytes,5,opt,name=CPF,proto3" json:"CPF,omitempty"`
}

func (x *AccountProduct) Reset() {
	*x = AccountProduct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_athena_athena_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountProduct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountProduct) ProtoMessage() {}

func (x *AccountProduct) ProtoReflect() protoreflect.Message {
	mi := &file_athena_athena_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountProduct.ProtoReflect.Descriptor instead.
func (*AccountProduct) Descriptor() ([]byte, []int) {
	return file_athena_athena_proto_rawDescGZIP(), []int{0}
}

func (x *AccountProduct) GetBanco() int64 {
	if x != nil {
		return x.Banco
	}
	return 0
}

func (x *AccountProduct) GetProduto() int64 {
	if x != nil {
		return x.Produto
	}
	return 0
}

func (x *AccountProduct) GetAgencia() int64 {
	if x != nil {
		return x.Agencia
	}
	return 0
}

func (x *AccountProduct) GetConta() int64 {
	if x != nil {
		return x.Conta
	}
	return 0
}

func (x *AccountProduct) GetCPF() string {
	if x != nil {
		return x.CPF
	}
	return ""
}

type BaseStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status       BaseStatus_EnumStatus `protobuf:"varint,1,opt,name=Status,proto3,enum=athena.BaseStatus_EnumStatus" json:"Status,omitempty"`
	Mensagem     string                `protobuf:"bytes,2,opt,name=Mensagem,proto3" json:"Mensagem,omitempty"`
	DataContabil string                `protobuf:"bytes,3,opt,name=DataContabil,proto3" json:"DataContabil,omitempty"`
}

func (x *BaseStatus) Reset() {
	*x = BaseStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_athena_athena_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseStatus) ProtoMessage() {}

func (x *BaseStatus) ProtoReflect() protoreflect.Message {
	mi := &file_athena_athena_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseStatus.ProtoReflect.Descriptor instead.
func (*BaseStatus) Descriptor() ([]byte, []int) {
	return file_athena_athena_proto_rawDescGZIP(), []int{1}
}

func (x *BaseStatus) GetStatus() BaseStatus_EnumStatus {
	if x != nil {
		return x.Status
	}
	return BaseStatus_SUCESSO
}

func (x *BaseStatus) GetMensagem() string {
	if x != nil {
		return x.Mensagem
	}
	return ""
}

func (x *BaseStatus) GetDataContabil() string {
	if x != nil {
		return x.DataContabil
	}
	return ""
}

var File_athena_athena_proto protoreflect.FileDescriptor

var file_athena_athena_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x74, 0x68, 0x65, 0x6e, 0x61, 0x2f, 0x61, 0x74, 0x68, 0x65, 0x6e, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x61, 0x74, 0x68, 0x65, 0x6e, 0x61, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x01, 0x0a, 0x0e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x42, 0x61, 0x6e, 0x63, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x42,
	0x61, 0x6e, 0x63, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x74, 0x6f, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x74, 0x6f, 0x12, 0x18,
	0x0a, 0x07, 0x41, 0x67, 0x65, 0x6e, 0x63, 0x69, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x41, 0x67, 0x65, 0x6e, 0x63, 0x69, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x12, 0x10,
	0x0a, 0x03, 0x43, 0x50, 0x46, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x43, 0x50, 0x46,
	0x22, 0xb8, 0x01, 0x0a, 0x0a, 0x42, 0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x35, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1d, 0x2e, 0x61, 0x74, 0x68, 0x65, 0x6e, 0x61, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x4d, 0x65, 0x6e, 0x73, 0x61, 0x67,
	0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4d, 0x65, 0x6e, 0x73, 0x61, 0x67,
	0x65, 0x6d, 0x12, 0x22, 0x0a, 0x0c, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x62,
	0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x62, 0x69, 0x6c, 0x22, 0x33, 0x0a, 0x0a, 0x45, 0x6e, 0x75, 0x6d, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x45, 0x53, 0x53, 0x4f, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x07, 0x4e, 0x45, 0x47, 0x4f, 0x43, 0x49, 0x4f, 0x10, 0x01, 0x12, 0x0b,
	0x0a, 0x07, 0x53, 0x49, 0x53, 0x54, 0x45, 0x4d, 0x41, 0x10, 0x02, 0x32, 0xc8, 0x01, 0x0a, 0x0f,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12,
	0x5e, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x61, 0x74, 0x68, 0x65, 0x6e, 0x61, 0x2e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x1a, 0x12, 0x2e, 0x61,
	0x74, 0x68, 0x65, 0x6e, 0x61, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x74,
	0x68, 0x65, 0x6e, 0x61, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x01, 0x2a, 0x12,
	0x55, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x61,
	0x74, 0x68, 0x65, 0x6e, 0x61, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x1a, 0x12, 0x2e, 0x61, 0x74, 0x68, 0x65, 0x6e, 0x61, 0x2e, 0x42, 0x61,
	0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16,
	0x12, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x74, 0x68, 0x65, 0x6e, 0x61, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x3a, 0x01, 0x2a, 0x42, 0x22, 0x5a, 0x20, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x74, 0x68, 0x65, 0x6e, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_athena_athena_proto_rawDescOnce sync.Once
	file_athena_athena_proto_rawDescData = file_athena_athena_proto_rawDesc
)

func file_athena_athena_proto_rawDescGZIP() []byte {
	file_athena_athena_proto_rawDescOnce.Do(func() {
		file_athena_athena_proto_rawDescData = protoimpl.X.CompressGZIP(file_athena_athena_proto_rawDescData)
	})
	return file_athena_athena_proto_rawDescData
}

var file_athena_athena_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_athena_athena_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_athena_athena_proto_goTypes = []interface{}{
	(BaseStatus_EnumStatus)(0), // 0: athena.BaseStatus.EnumStatus
	(*AccountProduct)(nil),     // 1: athena.AccountProduct
	(*BaseStatus)(nil),         // 2: athena.BaseStatus
}
var file_athena_athena_proto_depIdxs = []int32{
	0, // 0: athena.BaseStatus.Status:type_name -> athena.BaseStatus.EnumStatus
	1, // 1: athena.AccountServices.GetProductAccount:input_type -> athena.AccountProduct
	1, // 2: athena.AccountServices.GetClient:input_type -> athena.AccountProduct
	2, // 3: athena.AccountServices.GetProductAccount:output_type -> athena.BaseStatus
	2, // 4: athena.AccountServices.GetClient:output_type -> athena.BaseStatus
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_athena_athena_proto_init() }
func file_athena_athena_proto_init() {
	if File_athena_athena_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_athena_athena_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountProduct); i {
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
		file_athena_athena_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseStatus); i {
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
			RawDescriptor: file_athena_athena_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_athena_athena_proto_goTypes,
		DependencyIndexes: file_athena_athena_proto_depIdxs,
		EnumInfos:         file_athena_athena_proto_enumTypes,
		MessageInfos:      file_athena_athena_proto_msgTypes,
	}.Build()
	File_athena_athena_proto = out.File
	file_athena_athena_proto_rawDesc = nil
	file_athena_athena_proto_goTypes = nil
	file_athena_athena_proto_depIdxs = nil
}
