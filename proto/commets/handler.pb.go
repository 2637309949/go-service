// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.14.0
// source: proto/commets/handler.proto

package commets

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_proto_commets_handler_proto protoreflect.FileDescriptor

var file_proto_commets_handler_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x74, 0x73, 0x2f,
	0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x74, 0x73, 0x1a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x74, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x74, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0x82, 0x03, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x74, 0x73, 0x12,
	0x47, 0x0a, 0x0a, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x74, 0x73, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x74, 0x73, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x74, 0x73,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x74, 0x73, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x47, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x74, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x74, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x09, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x74,
	0x73, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x74, 0x73, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x56, 0x0a, 0x0f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x12, 0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x74, 0x73, 0x2e, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x74, 0x73, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_proto_commets_handler_proto_goTypes = []interface{}{
	(*InsertInfoRequest)(nil),       // 0: commets.InsertInfoRequest
	(*DeleteInfoRequest)(nil),       // 1: commets.DeleteInfoRequest
	(*UpdateInfoRequest)(nil),       // 2: commets.UpdateInfoRequest
	(*QueryInfoRequest)(nil),        // 3: commets.QueryInfoRequest
	(*QueryInfoDetailRequest)(nil),  // 4: commets.QueryInfoDetailRequest
	(*InsertInfoResponse)(nil),      // 5: commets.InsertInfoResponse
	(*DeleteInfoResponse)(nil),      // 6: commets.DeleteInfoResponse
	(*UpdateInfoResponse)(nil),      // 7: commets.UpdateInfoResponse
	(*QueryInfoResponse)(nil),       // 8: commets.QueryInfoResponse
	(*QueryInfoDetailResponse)(nil), // 9: commets.QueryInfoDetailResponse
}
var file_proto_commets_handler_proto_depIdxs = []int32{
	0, // 0: commets.Commets.InsertInfo:input_type -> commets.InsertInfoRequest
	1, // 1: commets.Commets.DeleteInfo:input_type -> commets.DeleteInfoRequest
	2, // 2: commets.Commets.UpdateInfo:input_type -> commets.UpdateInfoRequest
	3, // 3: commets.Commets.QueryInfo:input_type -> commets.QueryInfoRequest
	4, // 4: commets.Commets.QueryInfoDetail:input_type -> commets.QueryInfoDetailRequest
	5, // 5: commets.Commets.InsertInfo:output_type -> commets.InsertInfoResponse
	6, // 6: commets.Commets.DeleteInfo:output_type -> commets.DeleteInfoResponse
	7, // 7: commets.Commets.UpdateInfo:output_type -> commets.UpdateInfoResponse
	8, // 8: commets.Commets.QueryInfo:output_type -> commets.QueryInfoResponse
	9, // 9: commets.Commets.QueryInfoDetail:output_type -> commets.QueryInfoDetailResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_commets_handler_proto_init() }
func file_proto_commets_handler_proto_init() {
	if File_proto_commets_handler_proto != nil {
		return
	}
	file_proto_commets_commets_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_commets_handler_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_commets_handler_proto_goTypes,
		DependencyIndexes: file_proto_commets_handler_proto_depIdxs,
	}.Build()
	File_proto_commets_handler_proto = out.File
	file_proto_commets_handler_proto_rawDesc = nil
	file_proto_commets_handler_proto_goTypes = nil
	file_proto_commets_handler_proto_depIdxs = nil
}
