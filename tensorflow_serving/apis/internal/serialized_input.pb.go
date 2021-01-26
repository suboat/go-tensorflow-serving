// Copyright 2017 Google Inc. All Rights Reserved.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//==============================================================================

// Serialized counterparts of the messages in input.proto. These protos enable
// us to keep the original tensorflow.serving.Input's structure but with the
// tensorflow.Examples in their serialized form. When combined with lazy
// parsing, this improves performance by allowing us to skip a redundant
// deserialization/serialization loop.
//
// WARNING: These are internal implementation details and not part of the public
// API.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: tensorflow_serving/apis/internal/serialized_input.proto

package internal

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SerializedExampleList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Examples [][]byte `protobuf:"bytes,1,rep,name=examples,proto3" json:"examples,omitempty"`
}

func (x *SerializedExampleList) Reset() {
	*x = SerializedExampleList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_apis_internal_serialized_input_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SerializedExampleList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SerializedExampleList) ProtoMessage() {}

func (x *SerializedExampleList) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_apis_internal_serialized_input_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SerializedExampleList.ProtoReflect.Descriptor instead.
func (*SerializedExampleList) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_apis_internal_serialized_input_proto_rawDescGZIP(), []int{0}
}

func (x *SerializedExampleList) GetExamples() [][]byte {
	if x != nil {
		return x.Examples
	}
	return nil
}

type SerializedExampleListWithContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Examples [][]byte `protobuf:"bytes,1,rep,name=examples,proto3" json:"examples,omitempty"`
	Context  []byte   `protobuf:"bytes,2,opt,name=context,proto3" json:"context,omitempty"`
}

func (x *SerializedExampleListWithContext) Reset() {
	*x = SerializedExampleListWithContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_apis_internal_serialized_input_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SerializedExampleListWithContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SerializedExampleListWithContext) ProtoMessage() {}

func (x *SerializedExampleListWithContext) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_apis_internal_serialized_input_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SerializedExampleListWithContext.ProtoReflect.Descriptor instead.
func (*SerializedExampleListWithContext) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_apis_internal_serialized_input_proto_rawDescGZIP(), []int{1}
}

func (x *SerializedExampleListWithContext) GetExamples() [][]byte {
	if x != nil {
		return x.Examples
	}
	return nil
}

func (x *SerializedExampleListWithContext) GetContext() []byte {
	if x != nil {
		return x.Context
	}
	return nil
}

type SerializedInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Kind:
	//	*SerializedInput_ExampleList
	//	*SerializedInput_ExampleListWithContext
	Kind isSerializedInput_Kind `protobuf_oneof:"kind"`
}

func (x *SerializedInput) Reset() {
	*x = SerializedInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_apis_internal_serialized_input_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SerializedInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SerializedInput) ProtoMessage() {}

func (x *SerializedInput) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_apis_internal_serialized_input_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SerializedInput.ProtoReflect.Descriptor instead.
func (*SerializedInput) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_apis_internal_serialized_input_proto_rawDescGZIP(), []int{2}
}

func (m *SerializedInput) GetKind() isSerializedInput_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (x *SerializedInput) GetExampleList() *SerializedExampleList {
	if x, ok := x.GetKind().(*SerializedInput_ExampleList); ok {
		return x.ExampleList
	}
	return nil
}

func (x *SerializedInput) GetExampleListWithContext() *SerializedExampleListWithContext {
	if x, ok := x.GetKind().(*SerializedInput_ExampleListWithContext); ok {
		return x.ExampleListWithContext
	}
	return nil
}

type isSerializedInput_Kind interface {
	isSerializedInput_Kind()
}

type SerializedInput_ExampleList struct {
	ExampleList *SerializedExampleList `protobuf:"bytes,1,opt,name=example_list,json=exampleList,proto3,oneof"`
}

type SerializedInput_ExampleListWithContext struct {
	ExampleListWithContext *SerializedExampleListWithContext `protobuf:"bytes,2,opt,name=example_list_with_context,json=exampleListWithContext,proto3,oneof"`
}

func (*SerializedInput_ExampleList) isSerializedInput_Kind() {}

func (*SerializedInput_ExampleListWithContext) isSerializedInput_Kind() {}

var File_tensorflow_serving_apis_internal_serialized_input_proto protoreflect.FileDescriptor

var file_tensorflow_serving_apis_internal_serialized_input_proto_rawDesc = []byte{
	0x0a, 0x37, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x22, 0x33, 0x0a, 0x15, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x69, 0x7a, 0x65, 0x64, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0c, 0x52, 0x08, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x22, 0x58, 0x0a, 0x20, 0x53,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x57, 0x69, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0c, 0x52, 0x08, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x22, 0xee, 0x01, 0x0a, 0x0f, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x69, 0x7a, 0x65, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x57, 0x0a, 0x0c, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x32, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x53, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x7a, 0x0a, 0x19, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2e, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x45, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x69, 0x74, 0x68, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x48, 0x00, 0x52, 0x16, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x57, 0x69, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x42, 0x06,
	0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x42, 0x4d, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x75, 0x62, 0x6f, 0x61, 0x74, 0x2f, 0x67, 0x6f, 0x2d, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e,
	0x67, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_serving_apis_internal_serialized_input_proto_rawDescOnce sync.Once
	file_tensorflow_serving_apis_internal_serialized_input_proto_rawDescData = file_tensorflow_serving_apis_internal_serialized_input_proto_rawDesc
)

func file_tensorflow_serving_apis_internal_serialized_input_proto_rawDescGZIP() []byte {
	file_tensorflow_serving_apis_internal_serialized_input_proto_rawDescOnce.Do(func() {
		file_tensorflow_serving_apis_internal_serialized_input_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_serving_apis_internal_serialized_input_proto_rawDescData)
	})
	return file_tensorflow_serving_apis_internal_serialized_input_proto_rawDescData
}

var file_tensorflow_serving_apis_internal_serialized_input_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tensorflow_serving_apis_internal_serialized_input_proto_goTypes = []interface{}{
	(*SerializedExampleList)(nil),            // 0: tensorflow.serving.internal.SerializedExampleList
	(*SerializedExampleListWithContext)(nil), // 1: tensorflow.serving.internal.SerializedExampleListWithContext
	(*SerializedInput)(nil),                  // 2: tensorflow.serving.internal.SerializedInput
}
var file_tensorflow_serving_apis_internal_serialized_input_proto_depIdxs = []int32{
	0, // 0: tensorflow.serving.internal.SerializedInput.example_list:type_name -> tensorflow.serving.internal.SerializedExampleList
	1, // 1: tensorflow.serving.internal.SerializedInput.example_list_with_context:type_name -> tensorflow.serving.internal.SerializedExampleListWithContext
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_tensorflow_serving_apis_internal_serialized_input_proto_init() }
func file_tensorflow_serving_apis_internal_serialized_input_proto_init() {
	if File_tensorflow_serving_apis_internal_serialized_input_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_serving_apis_internal_serialized_input_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SerializedExampleList); i {
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
		file_tensorflow_serving_apis_internal_serialized_input_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SerializedExampleListWithContext); i {
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
		file_tensorflow_serving_apis_internal_serialized_input_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SerializedInput); i {
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
	file_tensorflow_serving_apis_internal_serialized_input_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*SerializedInput_ExampleList)(nil),
		(*SerializedInput_ExampleListWithContext)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tensorflow_serving_apis_internal_serialized_input_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_serving_apis_internal_serialized_input_proto_goTypes,
		DependencyIndexes: file_tensorflow_serving_apis_internal_serialized_input_proto_depIdxs,
		MessageInfos:      file_tensorflow_serving_apis_internal_serialized_input_proto_msgTypes,
	}.Build()
	File_tensorflow_serving_apis_internal_serialized_input_proto = out.File
	file_tensorflow_serving_apis_internal_serialized_input_proto_rawDesc = nil
	file_tensorflow_serving_apis_internal_serialized_input_proto_goTypes = nil
	file_tensorflow_serving_apis_internal_serialized_input_proto_depIdxs = nil
}
