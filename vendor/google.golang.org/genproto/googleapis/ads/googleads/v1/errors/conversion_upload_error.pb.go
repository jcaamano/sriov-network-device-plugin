// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/errors/conversion_upload_error.proto

package errors

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Enum describing possible conversion upload errors.
type ConversionUploadErrorEnum_ConversionUploadError int32

const (
	// Enum unspecified.
	ConversionUploadErrorEnum_UNSPECIFIED ConversionUploadErrorEnum_ConversionUploadError = 0
	// The received error code is not known in this version.
	ConversionUploadErrorEnum_UNKNOWN ConversionUploadErrorEnum_ConversionUploadError = 1
	// The request contained more than 2000 conversions.
	ConversionUploadErrorEnum_TOO_MANY_CONVERSIONS_IN_REQUEST ConversionUploadErrorEnum_ConversionUploadError = 2
	// The specified gclid could not be decoded.
	ConversionUploadErrorEnum_UNPARSEABLE_GCLID ConversionUploadErrorEnum_ConversionUploadError = 3
	// The specified conversion_date_time is before the event time
	// associated with the given gclid.
	ConversionUploadErrorEnum_CONVERSION_PRECEDES_GCLID ConversionUploadErrorEnum_ConversionUploadError = 4
	// The click associated with the given gclid is either too old to be
	// imported or occurred outside of the click through lookback window for the
	// specified conversion action.
	ConversionUploadErrorEnum_EXPIRED_GCLID ConversionUploadErrorEnum_ConversionUploadError = 5
	// The click associated with the given gclid occurred too recently. Please
	// try uploading again after 24 hours have passed since the click occurred.
	ConversionUploadErrorEnum_TOO_RECENT_GCLID ConversionUploadErrorEnum_ConversionUploadError = 6
	// The click associated with the given gclid could not be found in the
	// system. This can happen if Google Click IDs are collected for non Google
	// Ads clicks.
	ConversionUploadErrorEnum_GCLID_NOT_FOUND ConversionUploadErrorEnum_ConversionUploadError = 7
	// The click associated with the given gclid is owned by a customer
	// account that the uploading customer does not manage.
	ConversionUploadErrorEnum_UNAUTHORIZED_CUSTOMER ConversionUploadErrorEnum_ConversionUploadError = 8
	// No upload eligible conversion action that matches the provided
	// information can be found for the customer.
	ConversionUploadErrorEnum_INVALID_CONVERSION_ACTION ConversionUploadErrorEnum_ConversionUploadError = 9
	// The specified conversion action was created too recently.
	// Please try the upload again after 4-6 hours have passed since the
	// conversion action was created.
	ConversionUploadErrorEnum_TOO_RECENT_CONVERSION_ACTION ConversionUploadErrorEnum_ConversionUploadError = 10
	// The click associated with the given gclid does not contain conversion
	// tracking information.
	ConversionUploadErrorEnum_CONVERSION_TRACKING_NOT_ENABLED_AT_IMPRESSION_TIME ConversionUploadErrorEnum_ConversionUploadError = 11
	// The specified conversion action does not use an external attribution
	// model, but external_attribution_data was set.
	ConversionUploadErrorEnum_EXTERNAL_ATTRIBUTION_DATA_SET_FOR_NON_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION ConversionUploadErrorEnum_ConversionUploadError = 12
	// The specified conversion action uses an external attribution model, but
	// external_attribution_data or one of its contained fields was not set.
	// Both external_attribution_credit and external_attribution_model must be
	// set for externally attributed conversion actions.
	ConversionUploadErrorEnum_EXTERNAL_ATTRIBUTION_DATA_NOT_SET_FOR_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION ConversionUploadErrorEnum_ConversionUploadError = 13
	// Order IDs are not supported for conversion actions which use an external
	// attribution model.
	ConversionUploadErrorEnum_ORDER_ID_NOT_PERMITTED_FOR_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION ConversionUploadErrorEnum_ConversionUploadError = 14
	// A conversion with the same order id and conversion action combination
	// already exists in our system.
	ConversionUploadErrorEnum_ORDER_ID_ALREADY_IN_USE ConversionUploadErrorEnum_ConversionUploadError = 15
	// The request contained two or more conversions with the same order id and
	// conversion action combination.
	ConversionUploadErrorEnum_DUPLICATE_ORDER_ID ConversionUploadErrorEnum_ConversionUploadError = 16
	// The call occurred too recently. Please try uploading again after 24 hours
	// have passed since the call occurred.
	ConversionUploadErrorEnum_TOO_RECENT_CALL ConversionUploadErrorEnum_ConversionUploadError = 17
	// The click that initiated the call is too old for this conversion to be
	// imported.
	ConversionUploadErrorEnum_EXPIRED_CALL ConversionUploadErrorEnum_ConversionUploadError = 18
	// The call or the click leading to the call was not found.
	ConversionUploadErrorEnum_CALL_NOT_FOUND ConversionUploadErrorEnum_ConversionUploadError = 19
	// The specified conversion_date_time is before the call_start_date_time.
	ConversionUploadErrorEnum_CONVERSION_PRECEDES_CALL ConversionUploadErrorEnum_ConversionUploadError = 20
	// The click associated with the call does not contain conversion tracking
	// information.
	ConversionUploadErrorEnum_CONVERSION_TRACKING_NOT_ENABLED_AT_CALL_TIME ConversionUploadErrorEnum_ConversionUploadError = 21
	// The caller’s phone number cannot be parsed. It should be formatted either
	// as E.164 "+16502531234", International "+64 3-331 6005" or US national
	// number "6502531234".
	ConversionUploadErrorEnum_UNPARSEABLE_CALLERS_PHONE_NUMBER ConversionUploadErrorEnum_ConversionUploadError = 22
)

var ConversionUploadErrorEnum_ConversionUploadError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "TOO_MANY_CONVERSIONS_IN_REQUEST",
	3:  "UNPARSEABLE_GCLID",
	4:  "CONVERSION_PRECEDES_GCLID",
	5:  "EXPIRED_GCLID",
	6:  "TOO_RECENT_GCLID",
	7:  "GCLID_NOT_FOUND",
	8:  "UNAUTHORIZED_CUSTOMER",
	9:  "INVALID_CONVERSION_ACTION",
	10: "TOO_RECENT_CONVERSION_ACTION",
	11: "CONVERSION_TRACKING_NOT_ENABLED_AT_IMPRESSION_TIME",
	12: "EXTERNAL_ATTRIBUTION_DATA_SET_FOR_NON_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION",
	13: "EXTERNAL_ATTRIBUTION_DATA_NOT_SET_FOR_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION",
	14: "ORDER_ID_NOT_PERMITTED_FOR_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION",
	15: "ORDER_ID_ALREADY_IN_USE",
	16: "DUPLICATE_ORDER_ID",
	17: "TOO_RECENT_CALL",
	18: "EXPIRED_CALL",
	19: "CALL_NOT_FOUND",
	20: "CONVERSION_PRECEDES_CALL",
	21: "CONVERSION_TRACKING_NOT_ENABLED_AT_CALL_TIME",
	22: "UNPARSEABLE_CALLERS_PHONE_NUMBER",
}

var ConversionUploadErrorEnum_ConversionUploadError_value = map[string]int32{
	"UNSPECIFIED":                     0,
	"UNKNOWN":                         1,
	"TOO_MANY_CONVERSIONS_IN_REQUEST": 2,
	"UNPARSEABLE_GCLID":               3,
	"CONVERSION_PRECEDES_GCLID":       4,
	"EXPIRED_GCLID":                   5,
	"TOO_RECENT_GCLID":                6,
	"GCLID_NOT_FOUND":                 7,
	"UNAUTHORIZED_CUSTOMER":           8,
	"INVALID_CONVERSION_ACTION":       9,
	"TOO_RECENT_CONVERSION_ACTION":    10,
	"CONVERSION_TRACKING_NOT_ENABLED_AT_IMPRESSION_TIME":                            11,
	"EXTERNAL_ATTRIBUTION_DATA_SET_FOR_NON_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION": 12,
	"EXTERNAL_ATTRIBUTION_DATA_NOT_SET_FOR_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION": 13,
	"ORDER_ID_NOT_PERMITTED_FOR_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION":            14,
	"ORDER_ID_ALREADY_IN_USE":                      15,
	"DUPLICATE_ORDER_ID":                           16,
	"TOO_RECENT_CALL":                              17,
	"EXPIRED_CALL":                                 18,
	"CALL_NOT_FOUND":                               19,
	"CONVERSION_PRECEDES_CALL":                     20,
	"CONVERSION_TRACKING_NOT_ENABLED_AT_CALL_TIME": 21,
	"UNPARSEABLE_CALLERS_PHONE_NUMBER":             22,
}

func (x ConversionUploadErrorEnum_ConversionUploadError) String() string {
	return proto.EnumName(ConversionUploadErrorEnum_ConversionUploadError_name, int32(x))
}

func (ConversionUploadErrorEnum_ConversionUploadError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ba5bceb565e055e5, []int{0, 0}
}

// Container for enum describing possible conversion upload errors.
type ConversionUploadErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConversionUploadErrorEnum) Reset()         { *m = ConversionUploadErrorEnum{} }
func (m *ConversionUploadErrorEnum) String() string { return proto.CompactTextString(m) }
func (*ConversionUploadErrorEnum) ProtoMessage()    {}
func (*ConversionUploadErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba5bceb565e055e5, []int{0}
}

func (m *ConversionUploadErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConversionUploadErrorEnum.Unmarshal(m, b)
}
func (m *ConversionUploadErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConversionUploadErrorEnum.Marshal(b, m, deterministic)
}
func (m *ConversionUploadErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConversionUploadErrorEnum.Merge(m, src)
}
func (m *ConversionUploadErrorEnum) XXX_Size() int {
	return xxx_messageInfo_ConversionUploadErrorEnum.Size(m)
}
func (m *ConversionUploadErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ConversionUploadErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ConversionUploadErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.errors.ConversionUploadErrorEnum_ConversionUploadError", ConversionUploadErrorEnum_ConversionUploadError_name, ConversionUploadErrorEnum_ConversionUploadError_value)
	proto.RegisterType((*ConversionUploadErrorEnum)(nil), "google.ads.googleads.v1.errors.ConversionUploadErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/errors/conversion_upload_error.proto", fileDescriptor_ba5bceb565e055e5)
}

var fileDescriptor_ba5bceb565e055e5 = []byte{
	// 638 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x6e, 0xd3, 0x4e,
	0x10, 0xff, 0x37, 0xfd, 0x93, 0xc2, 0xf6, 0x6b, 0xbb, 0x6d, 0x0a, 0x2d, 0xa5, 0x54, 0x81, 0x23,
	0x72, 0x08, 0x48, 0x1c, 0x0c, 0x97, 0x8d, 0x3d, 0x4d, 0x57, 0x75, 0xd6, 0xee, 0x7a, 0x37, 0xb4,
	0x55, 0xa4, 0x55, 0x68, 0xa2, 0x28, 0x52, 0xeb, 0x8d, 0xe2, 0xb6, 0x0f, 0xc4, 0x91, 0x07, 0xe0,
	0x21, 0x38, 0xf3, 0x14, 0xdc, 0xb9, 0xa3, 0xb5, 0xe3, 0x10, 0x89, 0x50, 0x95, 0x53, 0x46, 0x33,
	0xbf, 0x8f, 0x89, 0x67, 0x66, 0xd1, 0x87, 0x81, 0x31, 0x83, 0xcb, 0x7e, 0xad, 0xdb, 0x4b, 0x6b,
	0x79, 0x68, 0xa3, 0xdb, 0x7a, 0xad, 0x3f, 0x1e, 0x9b, 0x71, 0x5a, 0xbb, 0x30, 0xc9, 0x6d, 0x7f,
	0x9c, 0x0e, 0x4d, 0xa2, 0x6f, 0x46, 0x97, 0xa6, 0xdb, 0xd3, 0x59, 0xc1, 0x19, 0x8d, 0xcd, 0xb5,
	0x21, 0xfb, 0x39, 0xc5, 0xe9, 0xf6, 0x52, 0x67, 0xca, 0x76, 0x6e, 0xeb, 0x4e, 0xce, 0xde, 0xdd,
	0x2b, 0xd4, 0x47, 0xc3, 0x5a, 0x37, 0x49, 0xcc, 0x75, 0xf7, 0x7a, 0x68, 0x92, 0x34, 0x67, 0x57,
	0xbf, 0x97, 0xd1, 0x8e, 0x37, 0xd5, 0x57, 0x99, 0x3c, 0x58, 0x22, 0x24, 0x37, 0x57, 0xd5, 0xaf,
	0x65, 0x54, 0x99, 0x5b, 0x25, 0xeb, 0x68, 0x59, 0xf1, 0x38, 0x02, 0x8f, 0x1d, 0x32, 0xf0, 0xf1,
	0x7f, 0x64, 0x19, 0x2d, 0x29, 0x7e, 0xcc, 0xc3, 0x8f, 0x1c, 0x2f, 0x90, 0x17, 0xe8, 0xb9, 0x0c,
	0x43, 0xdd, 0xa2, 0xfc, 0x4c, 0x7b, 0x21, 0x6f, 0x83, 0x88, 0x59, 0xc8, 0x63, 0xcd, 0xb8, 0x16,
	0x70, 0xa2, 0x20, 0x96, 0xb8, 0x44, 0x2a, 0x68, 0x43, 0xf1, 0x88, 0x8a, 0x18, 0x68, 0x23, 0x00,
	0xdd, 0xf4, 0x02, 0xe6, 0xe3, 0x45, 0xf2, 0x0c, 0xed, 0xfc, 0xa6, 0xe8, 0x48, 0x80, 0x07, 0x3e,
	0xc4, 0x93, 0xf2, 0xff, 0x64, 0x03, 0xad, 0xc2, 0x69, 0xc4, 0x04, 0xf8, 0x93, 0xd4, 0x03, 0xb2,
	0x85, 0xb0, 0x75, 0xb3, 0x48, 0x2e, 0x27, 0xd9, 0x32, 0xd9, 0x44, 0xeb, 0x59, 0xa8, 0x79, 0x28,
	0xf5, 0x61, 0xa8, 0xb8, 0x8f, 0x97, 0xc8, 0x0e, 0xaa, 0x28, 0x4e, 0x95, 0x3c, 0x0a, 0x05, 0x3b,
	0x07, 0x5f, 0x7b, 0x2a, 0x96, 0x61, 0x0b, 0x04, 0x7e, 0x68, 0x7d, 0x19, 0x6f, 0x53, 0xcb, 0x98,
	0xf1, 0xa7, 0x9e, 0x64, 0x21, 0xc7, 0x8f, 0xc8, 0x01, 0xda, 0x9b, 0x31, 0xf9, 0x13, 0x81, 0xc8,
	0x3b, 0xf4, 0x66, 0x26, 0x2d, 0x05, 0xf5, 0x8e, 0x19, 0x6f, 0x66, 0xf6, 0xc0, 0xed, 0x5f, 0xf4,
	0x35, 0x95, 0x9a, 0xb5, 0x22, 0x01, 0x71, 0x0e, 0x61, 0x2d, 0xc0, 0xcb, 0xe4, 0x04, 0xb5, 0xe0,
	0x54, 0x82, 0xe0, 0x34, 0xd0, 0x54, 0x4a, 0xc1, 0x1a, 0xca, 0x2a, 0x6a, 0x9f, 0x4a, 0xaa, 0x63,
	0xb0, 0xcd, 0x0b, 0xcd, 0x43, 0xae, 0x0b, 0x54, 0x70, 0x36, 0xc5, 0xc1, 0xbc, 0x66, 0x57, 0xee,
	0x96, 0xb4, 0x0d, 0x15, 0xb2, 0xf7, 0x95, 0x5c, 0x25, 0x87, 0xa8, 0x11, 0x0a, 0x1f, 0x84, 0x9e,
	0x7c, 0xd1, 0x08, 0x44, 0x8b, 0x49, 0x8b, 0xfe, 0x17, 0x9d, 0x35, 0xf2, 0x14, 0x3d, 0x9e, 0xea,
	0xd0, 0x40, 0x00, 0xf5, 0xcf, 0xec, 0x5a, 0xa8, 0x18, 0xf0, 0x3a, 0xd9, 0x46, 0xc4, 0x57, 0x51,
	0xc0, 0x3c, 0x2a, 0x41, 0x17, 0x30, 0x8c, 0xed, 0x2c, 0x67, 0x3f, 0x3e, 0x0d, 0x02, 0xbc, 0x41,
	0x30, 0x5a, 0x29, 0x36, 0x21, 0xcb, 0x10, 0x42, 0xd0, 0x9a, 0x8d, 0x66, 0x26, 0xbe, 0x49, 0xf6,
	0xd0, 0x93, 0x79, 0xeb, 0x94, 0x31, 0xb6, 0xc8, 0x6b, 0xf4, 0xea, 0x1e, 0x33, 0xcb, 0x44, 0xb3,
	0x69, 0x55, 0xc8, 0x4b, 0x74, 0x30, 0xbb, 0xb5, 0xb6, 0x04, 0x22, 0xd6, 0xd1, 0x51, 0xc8, 0x41,
	0x73, 0xd5, 0x6a, 0x80, 0xc0, 0xdb, 0x8d, 0x9f, 0x0b, 0xa8, 0x7a, 0x61, 0xae, 0x9c, 0xbb, 0x6f,
	0xb3, 0xb1, 0x3b, 0xf7, 0xb8, 0x22, 0x7b, 0x99, 0xd1, 0xc2, 0xb9, 0x3f, 0x61, 0x0f, 0xcc, 0x65,
	0x37, 0x19, 0x38, 0x66, 0x3c, 0xa8, 0x0d, 0xfa, 0x49, 0x76, 0xb7, 0xc5, 0x3b, 0x31, 0x1a, 0xa6,
	0x7f, 0x7b, 0x36, 0xde, 0xe7, 0x3f, 0x9f, 0x4b, 0x8b, 0x4d, 0x4a, 0xbf, 0x94, 0xf6, 0x9b, 0xb9,
	0x18, 0xed, 0xa5, 0x4e, 0x1e, 0xda, 0xa8, 0x5d, 0x77, 0x32, 0xcb, 0xf4, 0x5b, 0x01, 0xe8, 0xd0,
	0x5e, 0xda, 0x99, 0x02, 0x3a, 0xed, 0x7a, 0x27, 0x07, 0xfc, 0x28, 0x55, 0xf3, 0xac, 0xeb, 0xd2,
	0x5e, 0xea, 0xba, 0x53, 0x88, 0xeb, 0xb6, 0xeb, 0xae, 0x9b, 0x83, 0x3e, 0x95, 0xb3, 0xee, 0xde,
	0xfe, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x09, 0x85, 0xd8, 0x22, 0xd3, 0x04, 0x00, 0x00,
}
