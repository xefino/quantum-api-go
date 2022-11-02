// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: protos/frontend/ohlc/v1/service.proto

package v1

import (
	gopb "github.com/xefino/protobuf-gen-go/gopb"
	data "github.com/xefino/quantum-api-go/data"
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

// Describes the frequency of the bar
type Frequency int32

const (
	// Describes an invalid frequency. This is used for deserializing the frequency from JSON
	Frequency_InvalidFrequency Frequency = 0
	// Bars should be aggregated in terms of seconds
	Frequency_Second Frequency = 1
	// Bars should be aggregated in terms of minutes
	Frequency_Minute Frequency = 2
	// Bars should be aggregated in terms of hours
	Frequency_Hour Frequency = 3
	// Bars should be aggregated in terms of days
	Frequency_Day Frequency = 4
	// Bars should be aggregated in terms of weeks
	Frequency_Week Frequency = 5
	// Bars should be aggregated in terms of months
	Frequency_Month Frequency = 6
	// Bars should be aggregated in terms of quarters
	Frequency_Quarter Frequency = 7
	// Bars should be aggregated in terms of years
	Frequency_Year Frequency = 8
)

// Enum value maps for Frequency.
var (
	Frequency_name = map[int32]string{
		0: "InvalidFrequency",
		1: "Second",
		2: "Minute",
		3: "Hour",
		4: "Day",
		5: "Week",
		6: "Month",
		7: "Quarter",
		8: "Year",
	}
	Frequency_value = map[string]int32{
		"InvalidFrequency": 0,
		"Second":           1,
		"Minute":           2,
		"Hour":             3,
		"Day":              4,
		"Week":             5,
		"Month":            6,
		"Quarter":          7,
		"Year":             8,
	}
)

func (x Frequency) Enum() *Frequency {
	p := new(Frequency)
	*p = x
	return p
}

func (x Frequency) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Frequency) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_frontend_ohlc_v1_service_proto_enumTypes[0].Descriptor()
}

func (Frequency) Type() protoreflect.EnumType {
	return &file_protos_frontend_ohlc_v1_service_proto_enumTypes[0]
}

func (x Frequency) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Frequency.Descriptor instead.
func (Frequency) EnumDescriptor() ([]byte, []int) {
	return file_protos_frontend_ohlc_v1_service_proto_rawDescGZIP(), []int{0}
}

// Describes a request to get aggregated trades data.
type GetAggregatesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The starting time from which to pull aggregated data as a timestamp
	From *gopb.UnixTimestamp `protobuf:"bytes,1,opt,name=from,proto3" json:"from"` 
	// The ending time from which to pull aggregated data as a timestamp
	To *gopb.UnixTimestamp `protobuf:"bytes,2,opt,name=to,proto3" json:"to"` 
	// The symbol of the asset for which data was being requested
	Symbol string `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
	// The multiplier to associate with the frequency to describe the size of the bar being requested
	Multiplier uint32 `protobuf:"varint,4,opt,name=multiplier,proto3" json:"multiplier,omitempty"`
	// The frequency with which bars should be requested. This value is used together with the multiplier to describe the size of the bar
	Frequency Frequency `protobuf:"varint,5,opt,name=frequency,proto3,enum=protos.frontend.ohlc.v1.Frequency" json:"frequency"` 
	// Whether or not partial bars should be ignored. If this value is false then the last bar will be retrieved, even if it
	// extends past the end of the request period. Otherwise, the last bar will not be included in the results
	IgnorePartial bool `protobuf:"varint,6,opt,name=ignore_partial,json=ignorePartial,proto3" json:"ignore_partial,omitempty"`
	// Whether or not the aggregated data should be adjusted for splits. This is the functionality by default.  If it is disabled,
	// then the pricing and volume information will not be adjusted to account for splits.
	UnadjustSplits bool `protobuf:"varint,7,opt,name=unadjust_splits,json=unadjustSplits,proto3" json:"unadjust_splits,omitempty"`
}

func (x *GetAggregatesRequest) Reset() {
	*x = GetAggregatesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_frontend_ohlc_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAggregatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAggregatesRequest) ProtoMessage() {}

func (x *GetAggregatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_frontend_ohlc_v1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAggregatesRequest.ProtoReflect.Descriptor instead.
func (*GetAggregatesRequest) Descriptor() ([]byte, []int) {
	return file_protos_frontend_ohlc_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetAggregatesRequest) GetFrom() *gopb.UnixTimestamp {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *GetAggregatesRequest) GetTo() *gopb.UnixTimestamp {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *GetAggregatesRequest) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *GetAggregatesRequest) GetMultiplier() uint32 {
	if x != nil {
		return x.Multiplier
	}
	return 0
}

func (x *GetAggregatesRequest) GetFrequency() Frequency {
	if x != nil {
		return x.Frequency
	}
	return Frequency_InvalidFrequency
}

func (x *GetAggregatesRequest) GetIgnorePartial() bool {
	if x != nil {
		return x.IgnorePartial
	}
	return false
}

func (x *GetAggregatesRequest) GetUnadjustSplits() bool {
	if x != nil {
		return x.UnadjustSplits
	}
	return false
}

var File_protos_frontend_ohlc_v1_service_proto protoreflect.FileDescriptor

var file_protos_frontend_ohlc_v1_service_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x64, 0x2f, 0x6f, 0x68, 0x6c, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x6f, 0x68, 0x6c, 0x63, 0x2e, 0x76, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x2f, 0x6f, 0x68, 0x6c, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x02, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x30, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x55, 0x6e, 0x69, 0x78, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04,
	0x66, 0x72, 0x6f, 0x6d, 0x12, 0x2c, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x55, 0x6e, 0x69, 0x78, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02,
	0x74, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x75,
	0x6c, 0x74, 0x69, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x12, 0x40, 0x0a, 0x09, 0x66, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e,
	0x6f, 0x68, 0x6c, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63,
	0x79, 0x52, 0x09, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x25, 0x0a, 0x0e,
	0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x50, 0x61, 0x72, 0x74,
	0x69, 0x61, 0x6c, 0x12, 0x27, 0x0a, 0x0f, 0x75, 0x6e, 0x61, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x5f,
	0x73, 0x70, 0x6c, 0x69, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x75, 0x6e,
	0x61, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x73, 0x2a, 0x78, 0x0a, 0x09,
	0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x14, 0x0a, 0x10, 0x49, 0x6e, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4d,
	0x69, 0x6e, 0x75, 0x74, 0x65, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x6f, 0x75, 0x72, 0x10,
	0x03, 0x12, 0x07, 0x0a, 0x03, 0x44, 0x61, 0x79, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04, 0x57, 0x65,
	0x65, 0x6b, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x10, 0x06, 0x12,
	0x0b, 0x0a, 0x07, 0x51, 0x75, 0x61, 0x72, 0x74, 0x65, 0x72, 0x10, 0x07, 0x12, 0x08, 0x0a, 0x04,
	0x59, 0x65, 0x61, 0x72, 0x10, 0x08, 0x32, 0xac, 0x01, 0x0a, 0x0b, 0x4f, 0x68, 0x6c, 0x63, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x9c, 0x01, 0x0a, 0x0a, 0x41, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x73, 0x12, 0x2d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x66,
	0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x6f, 0x68, 0x6c, 0x63, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x66, 0x72,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x42, 0x61, 0x72, 0x22,
	0x42, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3c, 0x22, 0x37, 0x2f, 0x6f, 0x68, 0x6c, 0x63, 0x2f, 0x61,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2f, 0x7b, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c,
	0x7d, 0x2f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x2f, 0x7b, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c,
	0x69, 0x65, 0x72, 0x7d, 0x2f, 0x7b, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x7d,
	0x3a, 0x01, 0x2a, 0x30, 0x01, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x65, 0x66, 0x69, 0x6e, 0x6f, 0x2f, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x75, 0x6d, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x6f, 0x68, 0x6c, 0x63, 0x2f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_frontend_ohlc_v1_service_proto_rawDescOnce sync.Once
	file_protos_frontend_ohlc_v1_service_proto_rawDescData = file_protos_frontend_ohlc_v1_service_proto_rawDesc
)

func file_protos_frontend_ohlc_v1_service_proto_rawDescGZIP() []byte {
	file_protos_frontend_ohlc_v1_service_proto_rawDescOnce.Do(func() {
		file_protos_frontend_ohlc_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_frontend_ohlc_v1_service_proto_rawDescData)
	})
	return file_protos_frontend_ohlc_v1_service_proto_rawDescData
}

var file_protos_frontend_ohlc_v1_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_frontend_ohlc_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protos_frontend_ohlc_v1_service_proto_goTypes = []interface{}{
	(Frequency)(0),               // 0: protos.frontend.ohlc.v1.Frequency
	(*GetAggregatesRequest)(nil), // 1: protos.frontend.ohlc.v1.GetAggregatesRequest
	(*gopb.UnixTimestamp)(nil),   // 2: protos.common.UnixTimestamp
	(*data.Bar)(nil),             // 3: protos.frontend.data.Bar
}
var file_protos_frontend_ohlc_v1_service_proto_depIdxs = []int32{
	2, // 0: protos.frontend.ohlc.v1.GetAggregatesRequest.from:type_name -> protos.common.UnixTimestamp
	2, // 1: protos.frontend.ohlc.v1.GetAggregatesRequest.to:type_name -> protos.common.UnixTimestamp
	0, // 2: protos.frontend.ohlc.v1.GetAggregatesRequest.frequency:type_name -> protos.frontend.ohlc.v1.Frequency
	1, // 3: protos.frontend.ohlc.v1.OhlcService.Aggregates:input_type -> protos.frontend.ohlc.v1.GetAggregatesRequest
	3, // 4: protos.frontend.ohlc.v1.OhlcService.Aggregates:output_type -> protos.frontend.data.Bar
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_protos_frontend_ohlc_v1_service_proto_init() }
func file_protos_frontend_ohlc_v1_service_proto_init() {
	if File_protos_frontend_ohlc_v1_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_frontend_ohlc_v1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAggregatesRequest); i {
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
			RawDescriptor: file_protos_frontend_ohlc_v1_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_frontend_ohlc_v1_service_proto_goTypes,
		DependencyIndexes: file_protos_frontend_ohlc_v1_service_proto_depIdxs,
		EnumInfos:         file_protos_frontend_ohlc_v1_service_proto_enumTypes,
		MessageInfos:      file_protos_frontend_ohlc_v1_service_proto_msgTypes,
	}.Build()
	File_protos_frontend_ohlc_v1_service_proto = out.File
	file_protos_frontend_ohlc_v1_service_proto_rawDesc = nil
	file_protos_frontend_ohlc_v1_service_proto_goTypes = nil
	file_protos_frontend_ohlc_v1_service_proto_depIdxs = nil
}
