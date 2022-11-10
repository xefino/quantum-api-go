// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.3
// source: protos/frontend/data/ohlc.proto

package data

import (
	gopb "github.com/xefino/protobuf-gen-go/gopb"
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
	return file_protos_frontend_data_ohlc_proto_enumTypes[0].Descriptor()
}

func (Frequency) Type() protoreflect.EnumType {
	return &file_protos_frontend_data_ohlc_proto_enumTypes[0]
}

func (x Frequency) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Frequency.Descriptor instead.
func (Frequency) EnumDescriptor() ([]byte, []int) {
	return file_protos_frontend_data_ohlc_proto_rawDescGZIP(), []int{0}
}

// Describes a single aggregated bar of data
type Bar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Volume    int64               `protobuf:"varint,1,opt,name=volume,proto3" json:"volume,omitempty"`
	Weighted  float64             `protobuf:"fixed64,2,opt,name=weighted,proto3" json:"weighted,omitempty"`
	Open      float64             `protobuf:"fixed64,3,opt,name=open,proto3" json:"open,omitempty"`
	Close     float64             `protobuf:"fixed64,4,opt,name=close,proto3" json:"close,omitempty"`
	High      float64             `protobuf:"fixed64,5,opt,name=high,proto3" json:"high,omitempty"`
	Low       float64             `protobuf:"fixed64,6,opt,name=low,proto3" json:"low,omitempty"`
	Number    int32               `protobuf:"varint,7,opt,name=number,proto3" json:"number,omitempty"`
	Otc       bool                `protobuf:"varint,8,opt,name=otc,proto3" json:"otc,omitempty"`
	Timestamp *gopb.UnixTimestamp `protobuf:"bytes,9,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Bar) Reset() {
	*x = Bar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_frontend_data_ohlc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bar) ProtoMessage() {}

func (x *Bar) ProtoReflect() protoreflect.Message {
	mi := &file_protos_frontend_data_ohlc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bar.ProtoReflect.Descriptor instead.
func (*Bar) Descriptor() ([]byte, []int) {
	return file_protos_frontend_data_ohlc_proto_rawDescGZIP(), []int{0}
}

func (x *Bar) GetVolume() int64 {
	if x != nil {
		return x.Volume
	}
	return 0
}

func (x *Bar) GetWeighted() float64 {
	if x != nil {
		return x.Weighted
	}
	return 0
}

func (x *Bar) GetOpen() float64 {
	if x != nil {
		return x.Open
	}
	return 0
}

func (x *Bar) GetClose() float64 {
	if x != nil {
		return x.Close
	}
	return 0
}

func (x *Bar) GetHigh() float64 {
	if x != nil {
		return x.High
	}
	return 0
}

func (x *Bar) GetLow() float64 {
	if x != nil {
		return x.Low
	}
	return 0
}

func (x *Bar) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *Bar) GetOtc() bool {
	if x != nil {
		return x.Otc
	}
	return false
}

func (x *Bar) GetTimestamp() *gopb.UnixTimestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

var File_protos_frontend_data_ohlc_proto protoreflect.FileDescriptor

var file_protos_frontend_data_ohlc_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x64, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x6f, 0x68, 0x6c, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xef, 0x01, 0x0a, 0x03, 0x42, 0x61, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x08, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x65, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x6f, 0x70, 0x65,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x67, 0x68, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x68, 0x69, 0x67, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x6c,
	0x6f, 0x77, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x6f, 0x77, 0x12, 0x16, 0x0a,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x74, 0x63, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x03, 0x6f, 0x74, 0x63, 0x12, 0x3a, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x55, 0x6e, 0x69, 0x78, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2a, 0x78, 0x0a, 0x09, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79,
	0x12, 0x14, 0x0a, 0x10, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x46, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x79, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x10, 0x02, 0x12, 0x08,
	0x0a, 0x04, 0x48, 0x6f, 0x75, 0x72, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x44, 0x61, 0x79, 0x10,
	0x04, 0x12, 0x08, 0x0a, 0x04, 0x57, 0x65, 0x65, 0x6b, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x4d,
	0x6f, 0x6e, 0x74, 0x68, 0x10, 0x06, 0x12, 0x0b, 0x0a, 0x07, 0x51, 0x75, 0x61, 0x72, 0x74, 0x65,
	0x72, 0x10, 0x07, 0x12, 0x08, 0x0a, 0x04, 0x59, 0x65, 0x61, 0x72, 0x10, 0x08, 0x42, 0x27, 0x5a,
	0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x65, 0x66, 0x69,
	0x6e, 0x6f, 0x2f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x75, 0x6d, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67,
	0x6f, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_frontend_data_ohlc_proto_rawDescOnce sync.Once
	file_protos_frontend_data_ohlc_proto_rawDescData = file_protos_frontend_data_ohlc_proto_rawDesc
)

func file_protos_frontend_data_ohlc_proto_rawDescGZIP() []byte {
	file_protos_frontend_data_ohlc_proto_rawDescOnce.Do(func() {
		file_protos_frontend_data_ohlc_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_frontend_data_ohlc_proto_rawDescData)
	})
	return file_protos_frontend_data_ohlc_proto_rawDescData
}

var file_protos_frontend_data_ohlc_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_frontend_data_ohlc_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protos_frontend_data_ohlc_proto_goTypes = []interface{}{
	(Frequency)(0),             // 0: protos.frontend.data.Frequency
	(*Bar)(nil),                // 1: protos.frontend.data.Bar
	(*gopb.UnixTimestamp)(nil), // 2: protos.common.UnixTimestamp
}
var file_protos_frontend_data_ohlc_proto_depIdxs = []int32{
	2, // 0: protos.frontend.data.Bar.timestamp:type_name -> protos.common.UnixTimestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_frontend_data_ohlc_proto_init() }
func file_protos_frontend_data_ohlc_proto_init() {
	if File_protos_frontend_data_ohlc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_frontend_data_ohlc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bar); i {
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
			RawDescriptor: file_protos_frontend_data_ohlc_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_frontend_data_ohlc_proto_goTypes,
		DependencyIndexes: file_protos_frontend_data_ohlc_proto_depIdxs,
		EnumInfos:         file_protos_frontend_data_ohlc_proto_enumTypes,
		MessageInfos:      file_protos_frontend_data_ohlc_proto_msgTypes,
	}.Build()
	File_protos_frontend_data_ohlc_proto = out.File
	file_protos_frontend_data_ohlc_proto_rawDesc = nil
	file_protos_frontend_data_ohlc_proto_goTypes = nil
	file_protos_frontend_data_ohlc_proto_depIdxs = nil
}
