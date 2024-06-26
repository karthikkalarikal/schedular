// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.14.0
// source: employee-management/pkg/pb/employee.proto

package pb

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

type Employee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Department string `protobuf:"bytes,2,opt,name=department,proto3" json:"department,omitempty"`
	Role       string `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *Employee) Reset() {
	*x = Employee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_employee_management_pkg_pb_employee_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Employee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Employee) ProtoMessage() {}

func (x *Employee) ProtoReflect() protoreflect.Message {
	mi := &file_employee_management_pkg_pb_employee_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Employee.ProtoReflect.Descriptor instead.
func (*Employee) Descriptor() ([]byte, []int) {
	return file_employee_management_pkg_pb_employee_proto_rawDescGZIP(), []int{0}
}

func (x *Employee) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Employee) GetDepartment() string {
	if x != nil {
		return x.Department
	}
	return ""
}

func (x *Employee) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type EmployeeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Department string `protobuf:"bytes,3,opt,name=department,proto3" json:"department,omitempty"`
	Role       string `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty"`
	CreatedAt  string `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *EmployeeReply) Reset() {
	*x = EmployeeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_employee_management_pkg_pb_employee_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmployeeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmployeeReply) ProtoMessage() {}

func (x *EmployeeReply) ProtoReflect() protoreflect.Message {
	mi := &file_employee_management_pkg_pb_employee_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmployeeReply.ProtoReflect.Descriptor instead.
func (*EmployeeReply) Descriptor() ([]byte, []int) {
	return file_employee_management_pkg_pb_employee_proto_rawDescGZIP(), []int{1}
}

func (x *EmployeeReply) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *EmployeeReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EmployeeReply) GetDepartment() string {
	if x != nil {
		return x.Department
	}
	return ""
}

func (x *EmployeeReply) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *EmployeeReply) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

var File_employee_management_pkg_pb_employee_proto protoreflect.FileDescriptor

var file_employee_management_pkg_pb_employee_proto_rawDesc = []byte{
	0x0a, 0x29, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x65, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22,
	0x52, 0x0a, 0x08, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72,
	0x6f, 0x6c, 0x65, 0x22, 0x86, 0x01, 0x0a, 0x0d, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x70,
	0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64,
	0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0x3e, 0x0a, 0x0f,
	0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x2b, 0x0a, 0x08, 0x74, 0x65, 0x73, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x12, 0x0c, 0x2e, 0x70, 0x62,
	0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x1a, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x45,
	0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x1c, 0x5a, 0x1a,
	0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_employee_management_pkg_pb_employee_proto_rawDescOnce sync.Once
	file_employee_management_pkg_pb_employee_proto_rawDescData = file_employee_management_pkg_pb_employee_proto_rawDesc
)

func file_employee_management_pkg_pb_employee_proto_rawDescGZIP() []byte {
	file_employee_management_pkg_pb_employee_proto_rawDescOnce.Do(func() {
		file_employee_management_pkg_pb_employee_proto_rawDescData = protoimpl.X.CompressGZIP(file_employee_management_pkg_pb_employee_proto_rawDescData)
	})
	return file_employee_management_pkg_pb_employee_proto_rawDescData
}

var file_employee_management_pkg_pb_employee_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_employee_management_pkg_pb_employee_proto_goTypes = []interface{}{
	(*Employee)(nil),      // 0: pb.Employee
	(*EmployeeReply)(nil), // 1: pb.EmployeeReply
}
var file_employee_management_pkg_pb_employee_proto_depIdxs = []int32{
	0, // 0: pb.EmployeeService.testFunc:input_type -> pb.Employee
	1, // 1: pb.EmployeeService.testFunc:output_type -> pb.EmployeeReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_employee_management_pkg_pb_employee_proto_init() }
func file_employee_management_pkg_pb_employee_proto_init() {
	if File_employee_management_pkg_pb_employee_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_employee_management_pkg_pb_employee_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Employee); i {
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
		file_employee_management_pkg_pb_employee_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmployeeReply); i {
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
			RawDescriptor: file_employee_management_pkg_pb_employee_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_employee_management_pkg_pb_employee_proto_goTypes,
		DependencyIndexes: file_employee_management_pkg_pb_employee_proto_depIdxs,
		MessageInfos:      file_employee_management_pkg_pb_employee_proto_msgTypes,
	}.Build()
	File_employee_management_pkg_pb_employee_proto = out.File
	file_employee_management_pkg_pb_employee_proto_rawDesc = nil
	file_employee_management_pkg_pb_employee_proto_goTypes = nil
	file_employee_management_pkg_pb_employee_proto_depIdxs = nil
}
