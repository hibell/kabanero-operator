// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/hotel_placeholder_field.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v0/enums"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Possible values for Hotel placeholder fields.
type HotelPlaceholderFieldEnum_HotelPlaceholderField int32

const (
	// Not specified.
	HotelPlaceholderFieldEnum_UNSPECIFIED HotelPlaceholderFieldEnum_HotelPlaceholderField = 0
	// Used for return value only. Represents value unknown in this version.
	HotelPlaceholderFieldEnum_UNKNOWN HotelPlaceholderFieldEnum_HotelPlaceholderField = 1
	// Data Type: STRING. Required. Unique ID.
	HotelPlaceholderFieldEnum_PROPERTY_ID HotelPlaceholderFieldEnum_HotelPlaceholderField = 2
	// Data Type: STRING. Required. Main headline with property name to be shown
	// in dynamic ad.
	HotelPlaceholderFieldEnum_PROPERTY_NAME HotelPlaceholderFieldEnum_HotelPlaceholderField = 3
	// Data Type: STRING. Name of destination to be shown in dynamic ad.
	HotelPlaceholderFieldEnum_DESTINATION_NAME HotelPlaceholderFieldEnum_HotelPlaceholderField = 4
	// Data Type: STRING. Description of destination to be shown in dynamic ad.
	HotelPlaceholderFieldEnum_DESCRIPTION HotelPlaceholderFieldEnum_HotelPlaceholderField = 5
	// Data Type: STRING. Complete property address, including postal code.
	HotelPlaceholderFieldEnum_ADDRESS HotelPlaceholderFieldEnum_HotelPlaceholderField = 6
	// Data Type: STRING. Price to be shown in the ad.
	// Example: "100.00 USD"
	HotelPlaceholderFieldEnum_PRICE HotelPlaceholderFieldEnum_HotelPlaceholderField = 7
	// Data Type: STRING. Formatted price to be shown in the ad.
	// Example: "Starting at $100.00 USD", "$80 - $100"
	HotelPlaceholderFieldEnum_FORMATTED_PRICE HotelPlaceholderFieldEnum_HotelPlaceholderField = 8
	// Data Type: STRING. Sale price to be shown in the ad.
	// Example: "80.00 USD"
	HotelPlaceholderFieldEnum_SALE_PRICE HotelPlaceholderFieldEnum_HotelPlaceholderField = 9
	// Data Type: STRING. Formatted sale price to be shown in the ad.
	// Example: "On sale for $80.00", "$60 - $80"
	HotelPlaceholderFieldEnum_FORMATTED_SALE_PRICE HotelPlaceholderFieldEnum_HotelPlaceholderField = 10
	// Data Type: URL. Image to be displayed in the ad.
	HotelPlaceholderFieldEnum_IMAGE_URL HotelPlaceholderFieldEnum_HotelPlaceholderField = 11
	// Data Type: STRING. Category of property used to group like items together
	// for recommendation engine.
	HotelPlaceholderFieldEnum_CATEGORY HotelPlaceholderFieldEnum_HotelPlaceholderField = 12
	// Data Type: INT64. Star rating (1 to 5) used to group like items
	// together for recommendation engine.
	HotelPlaceholderFieldEnum_STAR_RATING HotelPlaceholderFieldEnum_HotelPlaceholderField = 13
	// Data Type: STRING_LIST. Keywords used for product retrieval.
	HotelPlaceholderFieldEnum_CONTEXTUAL_KEYWORDS HotelPlaceholderFieldEnum_HotelPlaceholderField = 14
	// Data Type: URL_LIST. Required. Final URLs for the ad when using Upgraded
	// URLs. User will be redirected to these URLs when they click on an ad, or
	// when they click on a specific flight for ads that show multiple
	// flights.
	HotelPlaceholderFieldEnum_FINAL_URLS HotelPlaceholderFieldEnum_HotelPlaceholderField = 15
	// Data Type: URL_LIST. Final mobile URLs for the ad when using Upgraded
	// URLs.
	HotelPlaceholderFieldEnum_FINAL_MOBILE_URLS HotelPlaceholderFieldEnum_HotelPlaceholderField = 16
	// Data Type: URL. Tracking template for the ad when using Upgraded URLs.
	HotelPlaceholderFieldEnum_TRACKING_URL HotelPlaceholderFieldEnum_HotelPlaceholderField = 17
	// Data Type: STRING. Android app link. Must be formatted as:
	// android-app://{package_id}/{scheme}/{host_path}.
	// The components are defined as follows:
	// package_id: app ID as specified in Google Play.
	// scheme: the scheme to pass to the application. Can be HTTP, or a custom
	//   scheme.
	// host_path: identifies the specific content within your application.
	HotelPlaceholderFieldEnum_ANDROID_APP_LINK HotelPlaceholderFieldEnum_HotelPlaceholderField = 18
	// Data Type: STRING_LIST. List of recommended property IDs to show together
	// with this item.
	HotelPlaceholderFieldEnum_SIMILAR_PROPERTY_IDS HotelPlaceholderFieldEnum_HotelPlaceholderField = 19
	// Data Type: STRING. iOS app link.
	HotelPlaceholderFieldEnum_IOS_APP_LINK HotelPlaceholderFieldEnum_HotelPlaceholderField = 20
	// Data Type: INT64. iOS app store ID.
	HotelPlaceholderFieldEnum_IOS_APP_STORE_ID HotelPlaceholderFieldEnum_HotelPlaceholderField = 21
)

var HotelPlaceholderFieldEnum_HotelPlaceholderField_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "PROPERTY_ID",
	3:  "PROPERTY_NAME",
	4:  "DESTINATION_NAME",
	5:  "DESCRIPTION",
	6:  "ADDRESS",
	7:  "PRICE",
	8:  "FORMATTED_PRICE",
	9:  "SALE_PRICE",
	10: "FORMATTED_SALE_PRICE",
	11: "IMAGE_URL",
	12: "CATEGORY",
	13: "STAR_RATING",
	14: "CONTEXTUAL_KEYWORDS",
	15: "FINAL_URLS",
	16: "FINAL_MOBILE_URLS",
	17: "TRACKING_URL",
	18: "ANDROID_APP_LINK",
	19: "SIMILAR_PROPERTY_IDS",
	20: "IOS_APP_LINK",
	21: "IOS_APP_STORE_ID",
}
var HotelPlaceholderFieldEnum_HotelPlaceholderField_value = map[string]int32{
	"UNSPECIFIED":          0,
	"UNKNOWN":              1,
	"PROPERTY_ID":          2,
	"PROPERTY_NAME":        3,
	"DESTINATION_NAME":     4,
	"DESCRIPTION":          5,
	"ADDRESS":              6,
	"PRICE":                7,
	"FORMATTED_PRICE":      8,
	"SALE_PRICE":           9,
	"FORMATTED_SALE_PRICE": 10,
	"IMAGE_URL":            11,
	"CATEGORY":             12,
	"STAR_RATING":          13,
	"CONTEXTUAL_KEYWORDS":  14,
	"FINAL_URLS":           15,
	"FINAL_MOBILE_URLS":    16,
	"TRACKING_URL":         17,
	"ANDROID_APP_LINK":     18,
	"SIMILAR_PROPERTY_IDS": 19,
	"IOS_APP_LINK":         20,
	"IOS_APP_STORE_ID":     21,
}

func (x HotelPlaceholderFieldEnum_HotelPlaceholderField) String() string {
	return proto.EnumName(HotelPlaceholderFieldEnum_HotelPlaceholderField_name, int32(x))
}
func (HotelPlaceholderFieldEnum_HotelPlaceholderField) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_hotel_placeholder_field_15aba0b561d9fb8c, []int{0, 0}
}

// Values for Hotel placeholder fields.
// For more information about dynamic remarketing feeds, see
// https://support.google.com/google-ads/answer/6053288.
type HotelPlaceholderFieldEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HotelPlaceholderFieldEnum) Reset()         { *m = HotelPlaceholderFieldEnum{} }
func (m *HotelPlaceholderFieldEnum) String() string { return proto.CompactTextString(m) }
func (*HotelPlaceholderFieldEnum) ProtoMessage()    {}
func (*HotelPlaceholderFieldEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_hotel_placeholder_field_15aba0b561d9fb8c, []int{0}
}
func (m *HotelPlaceholderFieldEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HotelPlaceholderFieldEnum.Unmarshal(m, b)
}
func (m *HotelPlaceholderFieldEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HotelPlaceholderFieldEnum.Marshal(b, m, deterministic)
}
func (dst *HotelPlaceholderFieldEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HotelPlaceholderFieldEnum.Merge(dst, src)
}
func (m *HotelPlaceholderFieldEnum) XXX_Size() int {
	return xxx_messageInfo_HotelPlaceholderFieldEnum.Size(m)
}
func (m *HotelPlaceholderFieldEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_HotelPlaceholderFieldEnum.DiscardUnknown(m)
}

var xxx_messageInfo_HotelPlaceholderFieldEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*HotelPlaceholderFieldEnum)(nil), "google.ads.googleads.v0.enums.HotelPlaceholderFieldEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.HotelPlaceholderFieldEnum_HotelPlaceholderField", HotelPlaceholderFieldEnum_HotelPlaceholderField_name, HotelPlaceholderFieldEnum_HotelPlaceholderField_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/hotel_placeholder_field.proto", fileDescriptor_hotel_placeholder_field_15aba0b561d9fb8c)
}

var fileDescriptor_hotel_placeholder_field_15aba0b561d9fb8c = []byte{
	// 514 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xdd, 0x6e, 0xda, 0x30,
	0x18, 0x1d, 0xb0, 0xfe, 0x60, 0xa0, 0x18, 0x03, 0xda, 0xa6, 0xa9, 0x17, 0xed, 0x03, 0x04, 0xa4,
	0xdd, 0xa5, 0x57, 0x26, 0x31, 0x99, 0x45, 0x62, 0x47, 0xb6, 0xa1, 0x63, 0x42, 0xb2, 0x58, 0x93,
	0xa5, 0x95, 0x02, 0x41, 0xa4, 0xed, 0x03, 0xed, 0x72, 0x37, 0x7b, 0x8f, 0x49, 0x7b, 0x91, 0x5d,
	0xed, 0x11, 0x26, 0x27, 0x05, 0xaa, 0x69, 0xdb, 0x4d, 0xf4, 0xf9, 0x9c, 0xef, 0x1c, 0x7f, 0xf1,
	0x77, 0xc0, 0x55, 0x92, 0x65, 0x49, 0x1a, 0x0f, 0x96, 0x51, 0x3e, 0x28, 0x4b, 0x53, 0x3d, 0x0e,
	0x07, 0xf1, 0xfa, 0x61, 0x95, 0x0f, 0x6e, 0xb3, 0xfb, 0x38, 0xd5, 0x9b, 0x74, 0x79, 0x13, 0xdf,
	0x66, 0x69, 0x14, 0x6f, 0xf5, 0xe7, 0xbb, 0x38, 0x8d, 0xac, 0xcd, 0x36, 0xbb, 0xcf, 0xd0, 0x79,
	0xa9, 0xb0, 0x96, 0x51, 0x6e, 0xed, 0xc5, 0xd6, 0xe3, 0xd0, 0x2a, 0xc4, 0x97, 0x3f, 0x6a, 0xe0,
	0xcd, 0x7b, 0x63, 0x10, 0x1e, 0xf4, 0x63, 0x23, 0x27, 0xeb, 0x87, 0xd5, 0xe5, 0xb7, 0x1a, 0xe8,
	0xff, 0x95, 0x45, 0x6d, 0xd0, 0x98, 0x32, 0x19, 0x12, 0x87, 0x8e, 0x29, 0x71, 0xe1, 0x0b, 0xd4,
	0x00, 0x27, 0x53, 0x36, 0x61, 0xfc, 0x9a, 0xc1, 0x8a, 0x61, 0x43, 0xc1, 0x43, 0x22, 0xd4, 0x5c,
	0x53, 0x17, 0x56, 0x51, 0x07, 0xb4, 0xf6, 0x00, 0xc3, 0x01, 0x81, 0x35, 0xd4, 0x03, 0xd0, 0x25,
	0x52, 0x51, 0x86, 0x15, 0xe5, 0xac, 0x44, 0x5f, 0x1a, 0xa5, 0x4b, 0xa4, 0x23, 0x68, 0x68, 0x50,
	0x78, 0x64, 0x7c, 0xb1, 0xeb, 0x0a, 0x22, 0x25, 0x3c, 0x46, 0x75, 0x70, 0x14, 0x0a, 0xea, 0x10,
	0x78, 0x82, 0xba, 0xa0, 0x3d, 0xe6, 0x22, 0xc0, 0x4a, 0x11, 0x57, 0x97, 0xe0, 0x29, 0x3a, 0x03,
	0x40, 0x62, 0x9f, 0x3c, 0x9d, 0xeb, 0xe8, 0x35, 0xe8, 0x1d, 0x9a, 0x9e, 0x31, 0x00, 0xb5, 0x40,
	0x9d, 0x06, 0xd8, 0x23, 0x7a, 0x2a, 0x7c, 0xd8, 0x40, 0x4d, 0x70, 0xea, 0x60, 0x45, 0x3c, 0x2e,
	0xe6, 0xb0, 0x69, 0x86, 0x90, 0x0a, 0x0b, 0x2d, 0xb0, 0xa2, 0xcc, 0x83, 0x2d, 0xf4, 0x0a, 0x74,
	0x1d, 0xce, 0x14, 0xf9, 0xa0, 0xa6, 0xd8, 0xd7, 0x13, 0x32, 0xbf, 0xe6, 0xc2, 0x95, 0xf0, 0xcc,
	0x5c, 0x38, 0xa6, 0x0c, 0xfb, 0xc6, 0x46, 0xc2, 0x36, 0xea, 0x83, 0x4e, 0x79, 0x0e, 0xf8, 0x88,
	0xfa, 0xa4, 0x84, 0x21, 0x82, 0xa0, 0xa9, 0x04, 0x76, 0x26, 0x94, 0x79, 0xc5, 0x85, 0x1d, 0xf3,
	0xf7, 0x98, 0xb9, 0x82, 0x53, 0x57, 0xe3, 0x30, 0xd4, 0x3e, 0x65, 0x13, 0x88, 0xcc, 0xbc, 0x92,
	0x06, 0xd4, 0xc7, 0x42, 0x3f, 0x7b, 0x3f, 0x09, 0xbb, 0xc6, 0x81, 0x72, 0x79, 0xe8, 0xed, 0x19,
	0x87, 0x1d, 0x22, 0x15, 0x17, 0xc4, 0x3c, 0x74, 0x7f, 0xf4, 0xab, 0x02, 0x2e, 0x6e, 0xb2, 0x95,
	0xf5, 0xdf, 0xad, 0x8f, 0xde, 0x16, 0x4b, 0xcd, 0xff, 0xdc, 0x6a, 0x68, 0x12, 0x13, 0x56, 0x3e,
	0x8e, 0x9e, 0xd4, 0x49, 0x96, 0x2e, 0xd7, 0x89, 0x95, 0x6d, 0x93, 0x41, 0x12, 0xaf, 0x8b, 0x3c,
	0xed, 0x02, 0xb8, 0xb9, 0xcb, 0xff, 0x91, 0xc7, 0xab, 0xe2, 0xfb, 0xa5, 0x5a, 0xf3, 0x30, 0xfe,
	0x5a, 0x3d, 0xf7, 0x4a, 0x2b, 0x1c, 0xe5, 0x56, 0x59, 0x9a, 0x6a, 0x36, 0xb4, 0x4c, 0xbe, 0xf2,
	0xef, 0x3b, 0x7e, 0x81, 0xa3, 0x7c, 0xb1, 0xe7, 0x17, 0xb3, 0xe1, 0xa2, 0xe0, 0x7f, 0x56, 0x2f,
	0x4a, 0xd0, 0xb6, 0x71, 0x94, 0xdb, 0xf6, 0xbe, 0xc3, 0xb6, 0x67, 0x43, 0xdb, 0x2e, 0x7a, 0x3e,
	0x1d, 0x17, 0x83, 0xbd, 0xfb, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x31, 0x11, 0x32, 0x27, 0x03,
	0x00, 0x00,
}
