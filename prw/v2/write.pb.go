// Copyright 2023 Prometheus Team
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: prw/v2/write.proto

package v2

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

type WriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The symbols table for all strings used in WriteRequest message.
	Symbols    []string      `protobuf:"bytes,1,rep,name=symbols,proto3" json:"symbols,omitempty"`
	Timeseries []*TimeSeries `protobuf:"bytes,2,rep,name=timeseries,proto3" json:"timeseries,omitempty"`
}

func (x *WriteRequest) Reset() {
	*x = WriteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prw_v2_write_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteRequest) ProtoMessage() {}

func (x *WriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_prw_v2_write_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteRequest.ProtoReflect.Descriptor instead.
func (*WriteRequest) Descriptor() ([]byte, []int) {
	return file_prw_v2_write_proto_rawDescGZIP(), []int{0}
}

func (x *WriteRequest) GetSymbols() []string {
	if x != nil {
		return x.Symbols
	}
	return nil
}

func (x *WriteRequest) GetTimeseries() []*TimeSeries {
	if x != nil {
		return x.Timeseries
	}
	return nil
}

// TimeSeries represents samples and labels for a single time series.
type TimeSeries struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Sorted list of label name-value pair references. This list's len is always multiple of 2,
	// packing tuples of (label name ref, label value ref) refs to WriteRequests.symbols.
	LabelSymbols []uint32 `protobuf:"varint,1,rep,packed,name=label_symbols,json=labelSymbols,proto3" json:"label_symbols,omitempty"`
	// Sorted by time, oldest sample first.
	Samples []*Sample `protobuf:"bytes,2,rep,name=samples,proto3" json:"samples,omitempty"`
}

func (x *TimeSeries) Reset() {
	*x = TimeSeries{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prw_v2_write_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeSeries) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeSeries) ProtoMessage() {}

func (x *TimeSeries) ProtoReflect() protoreflect.Message {
	mi := &file_prw_v2_write_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeSeries.ProtoReflect.Descriptor instead.
func (*TimeSeries) Descriptor() ([]byte, []int) {
	return file_prw_v2_write_proto_rawDescGZIP(), []int{1}
}

func (x *TimeSeries) GetLabelSymbols() []uint32 {
	if x != nil {
		return x.LabelSymbols
	}
	return nil
}

func (x *TimeSeries) GetSamples() []*Sample {
	if x != nil {
		return x.Samples
	}
	return nil
}

type Sample struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value float64 `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	// timestamp is in ms format, see model/timestamp/timestamp.go for
	// conversion from time.Time to Prometheus timestamp.
	Timestamp int64 `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Sample) Reset() {
	*x = Sample{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prw_v2_write_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sample) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sample) ProtoMessage() {}

func (x *Sample) ProtoReflect() protoreflect.Message {
	mi := &file_prw_v2_write_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sample.ProtoReflect.Descriptor instead.
func (*Sample) Descriptor() ([]byte, []int) {
	return file_prw_v2_write_proto_rawDescGZIP(), []int{2}
}

func (x *Sample) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Sample) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

var File_prw_v2_write_proto protoreflect.FileDescriptor

var file_prw_v2_write_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x77, 0x2f, 0x76, 0x32, 0x2f, 0x77, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x77, 0x2e, 0x76, 0x32, 0x22, 0x5c, 0x0a, 0x0c,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x12, 0x32, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x65,
	0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x77,
	0x2e, 0x76, 0x32, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x52, 0x0a,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x22, 0x73, 0x0a, 0x0a, 0x54, 0x69,
	0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x5f, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52,
	0x0c, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x12, 0x28, 0x0a,
	0x07, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x70, 0x72, 0x77, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x07,
	0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x4a, 0x04, 0x08,
	0x04, 0x10, 0x05, 0x4a, 0x04, 0x08, 0x05, 0x10, 0x06, 0x4a, 0x04, 0x08, 0x06, 0x10, 0x07, 0x22,
	0x3c, 0x0a, 0x06, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_prw_v2_write_proto_rawDescOnce sync.Once
	file_prw_v2_write_proto_rawDescData = file_prw_v2_write_proto_rawDesc
)

func file_prw_v2_write_proto_rawDescGZIP() []byte {
	file_prw_v2_write_proto_rawDescOnce.Do(func() {
		file_prw_v2_write_proto_rawDescData = protoimpl.X.CompressGZIP(file_prw_v2_write_proto_rawDescData)
	})
	return file_prw_v2_write_proto_rawDescData
}

var file_prw_v2_write_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_prw_v2_write_proto_goTypes = []interface{}{
	(*WriteRequest)(nil), // 0: prw.v2.WriteRequest
	(*TimeSeries)(nil),   // 1: prw.v2.TimeSeries
	(*Sample)(nil),       // 2: prw.v2.Sample
}
var file_prw_v2_write_proto_depIdxs = []int32{
	1, // 0: prw.v2.WriteRequest.timeseries:type_name -> prw.v2.TimeSeries
	2, // 1: prw.v2.TimeSeries.samples:type_name -> prw.v2.Sample
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_prw_v2_write_proto_init() }
func file_prw_v2_write_proto_init() {
	if File_prw_v2_write_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_prw_v2_write_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteRequest); i {
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
		file_prw_v2_write_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeSeries); i {
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
		file_prw_v2_write_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sample); i {
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
			RawDescriptor: file_prw_v2_write_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_prw_v2_write_proto_goTypes,
		DependencyIndexes: file_prw_v2_write_proto_depIdxs,
		MessageInfos:      file_prw_v2_write_proto_msgTypes,
	}.Build()
	File_prw_v2_write_proto = out.File
	file_prw_v2_write_proto_rawDesc = nil
	file_prw_v2_write_proto_goTypes = nil
	file_prw_v2_write_proto_depIdxs = nil
}
