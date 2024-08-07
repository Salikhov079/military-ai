// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.12.4
// source: groups.proto

package soldiers

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

type GroupReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DepartmentId string `protobuf:"bytes,2,opt,name=department_id,json=departmentId,proto3" json:"department_id,omitempty"`
}

func (x *GroupReq) Reset() {
	*x = GroupReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_groups_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupReq) ProtoMessage() {}

func (x *GroupReq) ProtoReflect() protoreflect.Message {
	mi := &file_groups_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupReq.ProtoReflect.Descriptor instead.
func (*GroupReq) Descriptor() ([]byte, []int) {
	return file_groups_proto_rawDescGZIP(), []int{0}
}

func (x *GroupReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GroupReq) GetDepartmentId() string {
	if x != nil {
		return x.DepartmentId
	}
	return ""
}

type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Department *Department `protobuf:"bytes,3,opt,name=department,proto3" json:"department,omitempty"`
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_groups_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_groups_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_groups_proto_rawDescGZIP(), []int{1}
}

func (x *Group) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Group) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Group) GetDepartment() *Department {
	if x != nil {
		return x.Department
	}
	return nil
}

type AllGroups struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Groups []*Group `protobuf:"bytes,1,rep,name=groups,proto3" json:"groups,omitempty"`
}

func (x *AllGroups) Reset() {
	*x = AllGroups{}
	if protoimpl.UnsafeEnabled {
		mi := &file_groups_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllGroups) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllGroups) ProtoMessage() {}

func (x *AllGroups) ProtoReflect() protoreflect.Message {
	mi := &file_groups_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllGroups.ProtoReflect.Descriptor instead.
func (*AllGroups) Descriptor() ([]byte, []int) {
	return file_groups_proto_rawDescGZIP(), []int{2}
}

func (x *AllGroups) GetGroups() []*Group {
	if x != nil {
		return x.Groups
	}
	return nil
}

var File_groups_proto protoreflect.FileDescriptor

var file_groups_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x1a, 0x11, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x08, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64,
	0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x22, 0x64, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a,
	0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x61,
	0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x32, 0x0a, 0x09, 0x41, 0x6c, 0x6c, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x73, 0x12, 0x25, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x32, 0xec, 0x01, 0x0a, 0x0c, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x65, 0x72, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x29, 0x0a, 0x06, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x0d, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x2e,
	0x56, 0x6f, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x10,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x42, 0x79, 0x49, 0x64,
	0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x56, 0x6f,
	0x69, 0x64, 0x12, 0x26, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x0d, 0x2e, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x2d, 0x0a, 0x06, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x12, 0x10, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e,
	0x41, 0x6c, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x42, 0x14, 0x5a, 0x12, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x6f, 0x6c, 0x64, 0x69, 0x65, 0x72, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_groups_proto_rawDescOnce sync.Once
	file_groups_proto_rawDescData = file_groups_proto_rawDesc
)

func file_groups_proto_rawDescGZIP() []byte {
	file_groups_proto_rawDescOnce.Do(func() {
		file_groups_proto_rawDescData = protoimpl.X.CompressGZIP(file_groups_proto_rawDescData)
	})
	return file_groups_proto_rawDescData
}

var file_groups_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_groups_proto_goTypes = []interface{}{
	(*GroupReq)(nil),   // 0: groups.GroupReq
	(*Group)(nil),      // 1: groups.Group
	(*AllGroups)(nil),  // 2: groups.AllGroups
	(*Department)(nil), // 3: departments.Department
	(*ById)(nil),       // 4: commanders.ById
	(*Void)(nil),       // 5: commanders.Void
}
var file_groups_proto_depIdxs = []int32{
	3, // 0: groups.Group.department:type_name -> departments.Department
	1, // 1: groups.AllGroups.groups:type_name -> groups.Group
	0, // 2: groups.GroupService.Create:input_type -> groups.GroupReq
	1, // 3: groups.GroupService.Update:input_type -> groups.Group
	4, // 4: groups.GroupService.Delete:input_type -> commanders.ById
	4, // 5: groups.GroupService.Get:input_type -> commanders.ById
	0, // 6: groups.GroupService.GetAll:input_type -> groups.GroupReq
	5, // 7: groups.GroupService.Create:output_type -> commanders.Void
	5, // 8: groups.GroupService.Update:output_type -> commanders.Void
	5, // 9: groups.GroupService.Delete:output_type -> commanders.Void
	1, // 10: groups.GroupService.Get:output_type -> groups.Group
	2, // 11: groups.GroupService.GetAll:output_type -> groups.AllGroups
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_groups_proto_init() }
func file_groups_proto_init() {
	if File_groups_proto != nil {
		return
	}
	file_departments_proto_init()
	file_commanders_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_groups_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupReq); i {
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
		file_groups_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Group); i {
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
		file_groups_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllGroups); i {
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
			RawDescriptor: file_groups_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_groups_proto_goTypes,
		DependencyIndexes: file_groups_proto_depIdxs,
		MessageInfos:      file_groups_proto_msgTypes,
	}.Build()
	File_groups_proto = out.File
	file_groups_proto_rawDesc = nil
	file_groups_proto_goTypes = nil
	file_groups_proto_depIdxs = nil
}
