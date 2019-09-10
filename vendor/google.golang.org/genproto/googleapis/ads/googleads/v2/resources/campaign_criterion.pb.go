// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/campaign_criterion.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v2/common"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v2/enums"
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

// A campaign criterion.
type CampaignCriterion struct {
	// The resource name of the campaign criterion.
	// Campaign criterion resource names have the form:
	//
	// `customers/{customer_id}/campaignCriteria/{campaign_id}~{criterion_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The campaign to which the criterion belongs.
	Campaign *wrappers.StringValue `protobuf:"bytes,4,opt,name=campaign,proto3" json:"campaign,omitempty"`
	// The ID of the criterion.
	//
	// This field is ignored during mutate.
	CriterionId *wrappers.Int64Value `protobuf:"bytes,5,opt,name=criterion_id,json=criterionId,proto3" json:"criterion_id,omitempty"`
	// The modifier for the bids when the criterion matches. The modifier must be
	// in the range: 0.1 - 10.0. Most targetable criteria types support modifiers.
	// Use 0 to opt out of a Device type.
	BidModifier *wrappers.FloatValue `protobuf:"bytes,14,opt,name=bid_modifier,json=bidModifier,proto3" json:"bid_modifier,omitempty"`
	// Whether to target (`false`) or exclude (`true`) the criterion.
	Negative *wrappers.BoolValue `protobuf:"bytes,7,opt,name=negative,proto3" json:"negative,omitempty"`
	// The type of the criterion.
	Type enums.CriterionTypeEnum_CriterionType `protobuf:"varint,6,opt,name=type,proto3,enum=google.ads.googleads.v2.enums.CriterionTypeEnum_CriterionType" json:"type,omitempty"`
	// The status of the criterion.
	Status enums.CampaignCriterionStatusEnum_CampaignCriterionStatus `protobuf:"varint,35,opt,name=status,proto3,enum=google.ads.googleads.v2.enums.CampaignCriterionStatusEnum_CampaignCriterionStatus" json:"status,omitempty"`
	// The campaign criterion.
	//
	// Exactly one must be set.
	//
	// Types that are valid to be assigned to Criterion:
	//	*CampaignCriterion_Keyword
	//	*CampaignCriterion_Placement
	//	*CampaignCriterion_MobileAppCategory
	//	*CampaignCriterion_MobileApplication
	//	*CampaignCriterion_Location
	//	*CampaignCriterion_Device
	//	*CampaignCriterion_AdSchedule
	//	*CampaignCriterion_AgeRange
	//	*CampaignCriterion_Gender
	//	*CampaignCriterion_IncomeRange
	//	*CampaignCriterion_ParentalStatus
	//	*CampaignCriterion_UserList
	//	*CampaignCriterion_YoutubeVideo
	//	*CampaignCriterion_YoutubeChannel
	//	*CampaignCriterion_Proximity
	//	*CampaignCriterion_Topic
	//	*CampaignCriterion_ListingScope
	//	*CampaignCriterion_Language
	//	*CampaignCriterion_IpBlock
	//	*CampaignCriterion_ContentLabel
	//	*CampaignCriterion_Carrier
	//	*CampaignCriterion_UserInterest
	//	*CampaignCriterion_Webpage
	//	*CampaignCriterion_OperatingSystemVersion
	//	*CampaignCriterion_MobileDevice
	//	*CampaignCriterion_LocationGroup
	Criterion            isCampaignCriterion_Criterion `protobuf_oneof:"criterion"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *CampaignCriterion) Reset()         { *m = CampaignCriterion{} }
func (m *CampaignCriterion) String() string { return proto.CompactTextString(m) }
func (*CampaignCriterion) ProtoMessage()    {}
func (*CampaignCriterion) Descriptor() ([]byte, []int) {
	return fileDescriptor_395bdd66d6311d9f, []int{0}
}

func (m *CampaignCriterion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignCriterion.Unmarshal(m, b)
}
func (m *CampaignCriterion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignCriterion.Marshal(b, m, deterministic)
}
func (m *CampaignCriterion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignCriterion.Merge(m, src)
}
func (m *CampaignCriterion) XXX_Size() int {
	return xxx_messageInfo_CampaignCriterion.Size(m)
}
func (m *CampaignCriterion) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignCriterion.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignCriterion proto.InternalMessageInfo

func (m *CampaignCriterion) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CampaignCriterion) GetCampaign() *wrappers.StringValue {
	if m != nil {
		return m.Campaign
	}
	return nil
}

func (m *CampaignCriterion) GetCriterionId() *wrappers.Int64Value {
	if m != nil {
		return m.CriterionId
	}
	return nil
}

func (m *CampaignCriterion) GetBidModifier() *wrappers.FloatValue {
	if m != nil {
		return m.BidModifier
	}
	return nil
}

func (m *CampaignCriterion) GetNegative() *wrappers.BoolValue {
	if m != nil {
		return m.Negative
	}
	return nil
}

func (m *CampaignCriterion) GetType() enums.CriterionTypeEnum_CriterionType {
	if m != nil {
		return m.Type
	}
	return enums.CriterionTypeEnum_UNSPECIFIED
}

func (m *CampaignCriterion) GetStatus() enums.CampaignCriterionStatusEnum_CampaignCriterionStatus {
	if m != nil {
		return m.Status
	}
	return enums.CampaignCriterionStatusEnum_UNSPECIFIED
}

type isCampaignCriterion_Criterion interface {
	isCampaignCriterion_Criterion()
}

type CampaignCriterion_Keyword struct {
	Keyword *common.KeywordInfo `protobuf:"bytes,8,opt,name=keyword,proto3,oneof"`
}

type CampaignCriterion_Placement struct {
	Placement *common.PlacementInfo `protobuf:"bytes,9,opt,name=placement,proto3,oneof"`
}

type CampaignCriterion_MobileAppCategory struct {
	MobileAppCategory *common.MobileAppCategoryInfo `protobuf:"bytes,10,opt,name=mobile_app_category,json=mobileAppCategory,proto3,oneof"`
}

type CampaignCriterion_MobileApplication struct {
	MobileApplication *common.MobileApplicationInfo `protobuf:"bytes,11,opt,name=mobile_application,json=mobileApplication,proto3,oneof"`
}

type CampaignCriterion_Location struct {
	Location *common.LocationInfo `protobuf:"bytes,12,opt,name=location,proto3,oneof"`
}

type CampaignCriterion_Device struct {
	Device *common.DeviceInfo `protobuf:"bytes,13,opt,name=device,proto3,oneof"`
}

type CampaignCriterion_AdSchedule struct {
	AdSchedule *common.AdScheduleInfo `protobuf:"bytes,15,opt,name=ad_schedule,json=adSchedule,proto3,oneof"`
}

type CampaignCriterion_AgeRange struct {
	AgeRange *common.AgeRangeInfo `protobuf:"bytes,16,opt,name=age_range,json=ageRange,proto3,oneof"`
}

type CampaignCriterion_Gender struct {
	Gender *common.GenderInfo `protobuf:"bytes,17,opt,name=gender,proto3,oneof"`
}

type CampaignCriterion_IncomeRange struct {
	IncomeRange *common.IncomeRangeInfo `protobuf:"bytes,18,opt,name=income_range,json=incomeRange,proto3,oneof"`
}

type CampaignCriterion_ParentalStatus struct {
	ParentalStatus *common.ParentalStatusInfo `protobuf:"bytes,19,opt,name=parental_status,json=parentalStatus,proto3,oneof"`
}

type CampaignCriterion_UserList struct {
	UserList *common.UserListInfo `protobuf:"bytes,22,opt,name=user_list,json=userList,proto3,oneof"`
}

type CampaignCriterion_YoutubeVideo struct {
	YoutubeVideo *common.YouTubeVideoInfo `protobuf:"bytes,20,opt,name=youtube_video,json=youtubeVideo,proto3,oneof"`
}

type CampaignCriterion_YoutubeChannel struct {
	YoutubeChannel *common.YouTubeChannelInfo `protobuf:"bytes,21,opt,name=youtube_channel,json=youtubeChannel,proto3,oneof"`
}

type CampaignCriterion_Proximity struct {
	Proximity *common.ProximityInfo `protobuf:"bytes,23,opt,name=proximity,proto3,oneof"`
}

type CampaignCriterion_Topic struct {
	Topic *common.TopicInfo `protobuf:"bytes,24,opt,name=topic,proto3,oneof"`
}

type CampaignCriterion_ListingScope struct {
	ListingScope *common.ListingScopeInfo `protobuf:"bytes,25,opt,name=listing_scope,json=listingScope,proto3,oneof"`
}

type CampaignCriterion_Language struct {
	Language *common.LanguageInfo `protobuf:"bytes,26,opt,name=language,proto3,oneof"`
}

type CampaignCriterion_IpBlock struct {
	IpBlock *common.IpBlockInfo `protobuf:"bytes,27,opt,name=ip_block,json=ipBlock,proto3,oneof"`
}

type CampaignCriterion_ContentLabel struct {
	ContentLabel *common.ContentLabelInfo `protobuf:"bytes,28,opt,name=content_label,json=contentLabel,proto3,oneof"`
}

type CampaignCriterion_Carrier struct {
	Carrier *common.CarrierInfo `protobuf:"bytes,29,opt,name=carrier,proto3,oneof"`
}

type CampaignCriterion_UserInterest struct {
	UserInterest *common.UserInterestInfo `protobuf:"bytes,30,opt,name=user_interest,json=userInterest,proto3,oneof"`
}

type CampaignCriterion_Webpage struct {
	Webpage *common.WebpageInfo `protobuf:"bytes,31,opt,name=webpage,proto3,oneof"`
}

type CampaignCriterion_OperatingSystemVersion struct {
	OperatingSystemVersion *common.OperatingSystemVersionInfo `protobuf:"bytes,32,opt,name=operating_system_version,json=operatingSystemVersion,proto3,oneof"`
}

type CampaignCriterion_MobileDevice struct {
	MobileDevice *common.MobileDeviceInfo `protobuf:"bytes,33,opt,name=mobile_device,json=mobileDevice,proto3,oneof"`
}

type CampaignCriterion_LocationGroup struct {
	LocationGroup *common.LocationGroupInfo `protobuf:"bytes,34,opt,name=location_group,json=locationGroup,proto3,oneof"`
}

func (*CampaignCriterion_Keyword) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_Placement) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_MobileAppCategory) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_MobileApplication) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_Location) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_Device) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_AdSchedule) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_AgeRange) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_Gender) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_IncomeRange) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_ParentalStatus) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_UserList) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_YoutubeVideo) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_YoutubeChannel) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_Proximity) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_Topic) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_ListingScope) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_Language) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_IpBlock) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_ContentLabel) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_Carrier) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_UserInterest) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_Webpage) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_OperatingSystemVersion) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_MobileDevice) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_LocationGroup) isCampaignCriterion_Criterion() {}

func (m *CampaignCriterion) GetCriterion() isCampaignCriterion_Criterion {
	if m != nil {
		return m.Criterion
	}
	return nil
}

func (m *CampaignCriterion) GetKeyword() *common.KeywordInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Keyword); ok {
		return x.Keyword
	}
	return nil
}

func (m *CampaignCriterion) GetPlacement() *common.PlacementInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Placement); ok {
		return x.Placement
	}
	return nil
}

func (m *CampaignCriterion) GetMobileAppCategory() *common.MobileAppCategoryInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_MobileAppCategory); ok {
		return x.MobileAppCategory
	}
	return nil
}

func (m *CampaignCriterion) GetMobileApplication() *common.MobileApplicationInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_MobileApplication); ok {
		return x.MobileApplication
	}
	return nil
}

func (m *CampaignCriterion) GetLocation() *common.LocationInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Location); ok {
		return x.Location
	}
	return nil
}

func (m *CampaignCriterion) GetDevice() *common.DeviceInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Device); ok {
		return x.Device
	}
	return nil
}

func (m *CampaignCriterion) GetAdSchedule() *common.AdScheduleInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_AdSchedule); ok {
		return x.AdSchedule
	}
	return nil
}

func (m *CampaignCriterion) GetAgeRange() *common.AgeRangeInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_AgeRange); ok {
		return x.AgeRange
	}
	return nil
}

func (m *CampaignCriterion) GetGender() *common.GenderInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Gender); ok {
		return x.Gender
	}
	return nil
}

func (m *CampaignCriterion) GetIncomeRange() *common.IncomeRangeInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_IncomeRange); ok {
		return x.IncomeRange
	}
	return nil
}

func (m *CampaignCriterion) GetParentalStatus() *common.ParentalStatusInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_ParentalStatus); ok {
		return x.ParentalStatus
	}
	return nil
}

func (m *CampaignCriterion) GetUserList() *common.UserListInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_UserList); ok {
		return x.UserList
	}
	return nil
}

func (m *CampaignCriterion) GetYoutubeVideo() *common.YouTubeVideoInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_YoutubeVideo); ok {
		return x.YoutubeVideo
	}
	return nil
}

func (m *CampaignCriterion) GetYoutubeChannel() *common.YouTubeChannelInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_YoutubeChannel); ok {
		return x.YoutubeChannel
	}
	return nil
}

func (m *CampaignCriterion) GetProximity() *common.ProximityInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Proximity); ok {
		return x.Proximity
	}
	return nil
}

func (m *CampaignCriterion) GetTopic() *common.TopicInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Topic); ok {
		return x.Topic
	}
	return nil
}

func (m *CampaignCriterion) GetListingScope() *common.ListingScopeInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_ListingScope); ok {
		return x.ListingScope
	}
	return nil
}

func (m *CampaignCriterion) GetLanguage() *common.LanguageInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Language); ok {
		return x.Language
	}
	return nil
}

func (m *CampaignCriterion) GetIpBlock() *common.IpBlockInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_IpBlock); ok {
		return x.IpBlock
	}
	return nil
}

func (m *CampaignCriterion) GetContentLabel() *common.ContentLabelInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_ContentLabel); ok {
		return x.ContentLabel
	}
	return nil
}

func (m *CampaignCriterion) GetCarrier() *common.CarrierInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Carrier); ok {
		return x.Carrier
	}
	return nil
}

func (m *CampaignCriterion) GetUserInterest() *common.UserInterestInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_UserInterest); ok {
		return x.UserInterest
	}
	return nil
}

func (m *CampaignCriterion) GetWebpage() *common.WebpageInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Webpage); ok {
		return x.Webpage
	}
	return nil
}

func (m *CampaignCriterion) GetOperatingSystemVersion() *common.OperatingSystemVersionInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_OperatingSystemVersion); ok {
		return x.OperatingSystemVersion
	}
	return nil
}

func (m *CampaignCriterion) GetMobileDevice() *common.MobileDeviceInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_MobileDevice); ok {
		return x.MobileDevice
	}
	return nil
}

func (m *CampaignCriterion) GetLocationGroup() *common.LocationGroupInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_LocationGroup); ok {
		return x.LocationGroup
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CampaignCriterion) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CampaignCriterion_Keyword)(nil),
		(*CampaignCriterion_Placement)(nil),
		(*CampaignCriterion_MobileAppCategory)(nil),
		(*CampaignCriterion_MobileApplication)(nil),
		(*CampaignCriterion_Location)(nil),
		(*CampaignCriterion_Device)(nil),
		(*CampaignCriterion_AdSchedule)(nil),
		(*CampaignCriterion_AgeRange)(nil),
		(*CampaignCriterion_Gender)(nil),
		(*CampaignCriterion_IncomeRange)(nil),
		(*CampaignCriterion_ParentalStatus)(nil),
		(*CampaignCriterion_UserList)(nil),
		(*CampaignCriterion_YoutubeVideo)(nil),
		(*CampaignCriterion_YoutubeChannel)(nil),
		(*CampaignCriterion_Proximity)(nil),
		(*CampaignCriterion_Topic)(nil),
		(*CampaignCriterion_ListingScope)(nil),
		(*CampaignCriterion_Language)(nil),
		(*CampaignCriterion_IpBlock)(nil),
		(*CampaignCriterion_ContentLabel)(nil),
		(*CampaignCriterion_Carrier)(nil),
		(*CampaignCriterion_UserInterest)(nil),
		(*CampaignCriterion_Webpage)(nil),
		(*CampaignCriterion_OperatingSystemVersion)(nil),
		(*CampaignCriterion_MobileDevice)(nil),
		(*CampaignCriterion_LocationGroup)(nil),
	}
}

func init() {
	proto.RegisterType((*CampaignCriterion)(nil), "google.ads.googleads.v2.resources.CampaignCriterion")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/campaign_criterion.proto", fileDescriptor_395bdd66d6311d9f)
}

var fileDescriptor_395bdd66d6311d9f = []byte{
	// 1114 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x97, 0xed, 0x72, 0xdb, 0x44,
	0x17, 0x80, 0xdf, 0xe4, 0x6d, 0x53, 0x67, 0xed, 0xa4, 0x74, 0x4b, 0xcb, 0x92, 0x96, 0x92, 0xb6,
	0xd3, 0x99, 0xf2, 0x51, 0x19, 0x0c, 0x74, 0x18, 0x33, 0x74, 0xc6, 0x71, 0x21, 0x75, 0x9b, 0x42,
	0x70, 0x82, 0x3b, 0x74, 0xc2, 0x68, 0xd6, 0xd2, 0x89, 0xba, 0x54, 0xda, 0xd5, 0xac, 0x56, 0x0e,
	0xfe, 0x0d, 0x57, 0xc2, 0x4f, 0x2e, 0x85, 0x4b, 0xe1, 0x2a, 0x98, 0xfd, 0x52, 0x9c, 0x26, 0x41,
	0xe2, 0x9f, 0x74, 0x74, 0x9e, 0x67, 0xcf, 0x1e, 0x4b, 0xbb, 0x6b, 0xd4, 0x4f, 0x84, 0x48, 0x52,
	0xe8, 0xd2, 0xb8, 0xe8, 0xda, 0x4b, 0x7d, 0x35, 0xeb, 0x75, 0x25, 0x14, 0xa2, 0x94, 0x11, 0x14,
	0xdd, 0x88, 0x66, 0x39, 0x65, 0x09, 0x0f, 0x23, 0xc9, 0x14, 0x48, 0x26, 0x78, 0x90, 0x4b, 0xa1,
	0x04, 0xbe, 0x6d, 0x81, 0x80, 0xc6, 0x45, 0x50, 0xb1, 0xc1, 0xac, 0x17, 0x54, 0xec, 0xc6, 0x83,
	0xf3, 0xf4, 0x91, 0xc8, 0x32, 0xc1, 0xbb, 0x4e, 0x49, 0xad, 0x71, 0xe3, 0xeb, 0xf3, 0xd2, 0x81,
	0x97, 0xd9, 0x59, 0x95, 0x84, 0x85, 0xa2, 0xaa, 0x2c, 0x1c, 0xde, 0xab, 0xc1, 0x2b, 0x4a, 0xcd,
	0x73, 0x70, 0xcc, 0x2d, 0xc7, 0x98, 0xbb, 0x69, 0x79, 0xd8, 0x3d, 0x92, 0x34, 0xcf, 0x41, 0x7a,
	0xe7, 0x4d, 0xef, 0xcc, 0x59, 0x97, 0x72, 0x2e, 0x14, 0x55, 0x4c, 0x70, 0xf7, 0xf4, 0xce, 0xef,
	0xd7, 0xd0, 0x95, 0xa1, 0xab, 0x6a, 0xe8, 0xf5, 0xf8, 0x2e, 0x5a, 0xf3, 0x2d, 0x08, 0x39, 0xcd,
	0x80, 0x2c, 0x6d, 0x2e, 0xdd, 0x5f, 0x1d, 0x77, 0x7c, 0xf0, 0x3b, 0x9a, 0x01, 0xfe, 0x12, 0xb5,
	0xfc, 0x7c, 0xc8, 0x85, 0xcd, 0xa5, 0xfb, 0xed, 0xde, 0x4d, 0xd7, 0xc5, 0xc0, 0xd7, 0x12, 0xec,
	0x29, 0xc9, 0x78, 0x32, 0xa1, 0x69, 0x09, 0xe3, 0x2a, 0x1b, 0x3f, 0x42, 0x9d, 0xe3, 0xa9, 0xb0,
	0x98, 0x5c, 0x34, 0xf4, 0x8d, 0x53, 0xf4, 0x88, 0xab, 0x87, 0x9f, 0x5b, 0xb8, 0x5d, 0x01, 0xa3,
	0x58, 0xf3, 0x53, 0x16, 0x87, 0x99, 0x88, 0xd9, 0x21, 0x03, 0x49, 0xd6, 0xcf, 0xe1, 0xbf, 0x4d,
	0x05, 0x55, 0x8e, 0x9f, 0xb2, 0xf8, 0xb9, 0xcb, 0xc7, 0x0f, 0x51, 0x8b, 0x43, 0x42, 0x15, 0x9b,
	0x01, 0xb9, 0x64, 0xd8, 0x8d, 0x53, 0xec, 0x96, 0x10, 0xa9, 0xab, 0xdb, 0xe7, 0xe2, 0x31, 0xba,
	0xa0, 0x1b, 0x4f, 0x56, 0x36, 0x97, 0xee, 0xaf, 0xf7, 0x1e, 0x05, 0xe7, 0xbd, 0x3e, 0xe6, 0xd7,
	0x0a, 0xaa, 0x76, 0xee, 0xcf, 0x73, 0xf8, 0x86, 0x97, 0xd9, 0xc9, 0xc8, 0xd8, 0xb8, 0xf0, 0x2f,
	0x68, 0xc5, 0xbe, 0x02, 0xe4, 0xae, 0xb1, 0x8e, 0xeb, 0xac, 0x6f, 0xfe, 0x58, 0x7b, 0x86, 0xb6,
	0xfe, 0xb3, 0x9f, 0x8d, 0xdd, 0x08, 0x78, 0x1b, 0x5d, 0x7a, 0x0d, 0xf3, 0x23, 0x21, 0x63, 0xd2,
	0x32, 0xd3, 0xfe, 0xe8, 0xdc, 0xc1, 0xec, 0xeb, 0x1d, 0x3c, 0xb3, 0xe9, 0x23, 0x7e, 0x28, 0x9e,
	0xfc, 0x6f, 0xec, 0x69, 0xfc, 0x1c, 0xad, 0xe6, 0x29, 0x8d, 0x20, 0x03, 0xae, 0xc8, 0xaa, 0x51,
	0x3d, 0xa8, 0x53, 0xed, 0x7a, 0xc0, 0xc9, 0x8e, 0x0d, 0x38, 0x41, 0x57, 0x33, 0x31, 0x65, 0x29,
	0x84, 0x34, 0xcf, 0xc3, 0x88, 0x2a, 0x48, 0x84, 0x9c, 0x13, 0x64, 0xc4, 0x5f, 0xd4, 0x89, 0x9f,
	0x1b, 0x74, 0x90, 0xe7, 0x43, 0x07, 0xba, 0x01, 0xae, 0x64, 0x6f, 0x3e, 0xc0, 0x87, 0x08, 0x1f,
	0x0f, 0x94, 0xb2, 0xc8, 0x7c, 0x0a, 0xa4, 0xfd, 0x1f, 0xc7, 0xf1, 0xe0, 0xa9, 0x71, 0xfc, 0x03,
	0xfc, 0x14, 0xb5, 0x52, 0xe1, 0xec, 0x1d, 0x63, 0xff, 0xb8, 0xce, 0xbe, 0x23, 0x4e, 0x48, 0x2b,
	0x1e, 0x3f, 0x46, 0x2b, 0x31, 0xcc, 0x58, 0x04, 0x64, 0xcd, 0x98, 0x3e, 0xac, 0x33, 0x3d, 0x36,
	0xd9, 0xce, 0xe3, 0x58, 0xfc, 0x03, 0x6a, 0xd3, 0x38, 0x2c, 0xa2, 0x57, 0x10, 0x97, 0x29, 0x90,
	0xcb, 0x46, 0x15, 0xd4, 0xa9, 0x06, 0xf1, 0x9e, 0x23, 0x9c, 0x0e, 0xd1, 0x2a, 0x82, 0x9f, 0xa1,
	0x55, 0x9a, 0x40, 0x28, 0x29, 0x4f, 0x80, 0xbc, 0xd5, 0x6c, 0x96, 0x83, 0x04, 0xc6, 0x3a, 0xdf,
	0xcf, 0x92, 0xba, 0x7b, 0x3d, 0xcb, 0x04, 0x78, 0x0c, 0x92, 0x5c, 0x69, 0x36, 0xcb, 0x6d, 0x93,
	0xed, 0x67, 0x69, 0x59, 0xbc, 0x8f, 0x3a, 0x8c, 0x47, 0x22, 0xf3, 0x55, 0x61, 0xe3, 0xea, 0xd6,
	0xb9, 0x46, 0x86, 0x59, 0x2c, 0xac, 0xcd, 0x8e, 0x43, 0xf8, 0x67, 0x74, 0x39, 0xa7, 0x12, 0xb8,
	0xa2, 0xa9, 0x5b, 0xae, 0xc9, 0x55, 0x23, 0xee, 0xd5, 0xbe, 0xf3, 0x0e, 0xb3, 0xdf, 0xa1, 0x73,
	0xaf, 0xe7, 0x27, 0xa2, 0xba, 0x8f, 0x65, 0x01, 0x32, 0x4c, 0x59, 0xa1, 0xc8, 0xf5, 0x66, 0x7d,
	0xfc, 0xb1, 0x00, 0xb9, 0xc3, 0x0a, 0xff, 0x2d, 0xb5, 0x4a, 0x77, 0x8f, 0x5f, 0xa0, 0xb5, 0xb9,
	0x28, 0x55, 0x39, 0x85, 0x70, 0xc6, 0x62, 0x10, 0xe4, 0x6d, 0x23, 0xfc, 0xa4, 0x4e, 0xf8, 0x93,
	0x28, 0xf7, 0xcb, 0x29, 0x4c, 0x34, 0xe3, 0xa4, 0x1d, 0x27, 0x32, 0x31, 0xdd, 0x04, 0x2f, 0x8e,
	0x5e, 0x51, 0xce, 0x21, 0x25, 0xd7, 0x9a, 0x35, 0xc1, 0xa9, 0x87, 0x96, 0xf2, 0x4d, 0x70, 0x32,
	0x17, 0x35, 0x2b, 0x8a, 0x14, 0xbf, 0xb2, 0x8c, 0xa9, 0x39, 0x79, 0xa7, 0xe1, 0x8a, 0xe2, 0x81,
	0x6a, 0x45, 0xf1, 0x01, 0x3c, 0x40, 0x17, 0x95, 0xc8, 0x59, 0x44, 0x88, 0x51, 0x7d, 0x50, 0xa7,
	0xda, 0xd7, 0xc9, 0x4e, 0x63, 0x49, 0xdd, 0x49, 0xfd, 0x8b, 0x30, 0x9e, 0x84, 0x45, 0x24, 0x72,
	0x20, 0xef, 0x36, 0xeb, 0xe4, 0x8e, 0x85, 0xf6, 0x34, 0xe3, 0x3b, 0x99, 0x2e, 0xc4, 0xcc, 0xe2,
	0x40, 0x79, 0x52, 0xd2, 0x04, 0xc8, 0x46, 0xc3, 0xc5, 0xc1, 0xe5, 0x57, 0x8b, 0x83, 0xbb, 0xc7,
	0x4f, 0x50, 0x8b, 0xe5, 0xe1, 0x34, 0x15, 0xd1, 0x6b, 0x72, 0xa3, 0xd9, 0x92, 0x3e, 0xca, 0xb7,
	0x74, 0xba, 0x5f, 0xd2, 0x99, 0xbd, 0xd5, 0xd3, 0x8d, 0x04, 0x57, 0xc0, 0x55, 0x98, 0xd2, 0x29,
	0xa4, 0xe4, 0x66, 0xb3, 0xe9, 0x0e, 0x2d, 0xb4, 0xa3, 0x19, 0x3f, 0xdd, 0x68, 0x21, 0xa6, 0x37,
	0x9d, 0x88, 0x4a, 0xa9, 0xf7, 0xe9, 0xf7, 0x9a, 0x55, 0x38, 0xb4, 0xe9, 0xbe, 0x42, 0x47, 0xeb,
	0x0a, 0xcd, 0x77, 0xc2, 0xb8, 0x02, 0x09, 0x85, 0x22, 0xb7, 0x9a, 0x55, 0xa8, 0xbf, 0x95, 0x91,
	0x63, 0x7c, 0x85, 0xe5, 0x42, 0x4c, 0x57, 0x78, 0x04, 0xd3, 0x5c, 0xff, 0x1e, 0xef, 0x37, 0xab,
	0xf0, 0x85, 0x4d, 0xf7, 0x15, 0x3a, 0x1a, 0xcf, 0x10, 0x11, 0x39, 0x48, 0x6a, 0x5f, 0x9a, 0x79,
	0xa1, 0x20, 0x0b, 0x67, 0x20, 0x0b, 0xbd, 0x0d, 0x6c, 0x1a, 0x73, 0xbf, 0xce, 0xfc, 0xbd, 0xe7,
	0xf7, 0x0c, 0x3e, 0xb1, 0xb4, 0x1b, 0xe8, 0xba, 0x38, 0xf3, 0xa9, 0xee, 0x8c, 0xdb, 0xd6, 0xdc,
	0x4e, 0x71, 0xbb, 0x59, 0x67, 0xec, 0x8e, 0x76, 0x62, 0xbf, 0xe8, 0x64, 0x0b, 0x31, 0xfc, 0x12,
	0xad, 0xfb, 0x7d, 0x28, 0x4c, 0xa4, 0x28, 0x73, 0x72, 0xc7, 0x98, 0x3f, 0x6d, 0xba, 0x9b, 0x6d,
	0x6b, 0xc8, 0xa9, 0xd7, 0xd2, 0xc5, 0xe0, 0x56, 0x1b, 0xad, 0x56, 0x67, 0xba, 0xad, 0xdf, 0x96,
	0xd1, 0xbd, 0x48, 0x64, 0x41, 0xed, 0x81, 0x7c, 0xeb, 0xfa, 0xa9, 0x43, 0xce, 0xae, 0x3e, 0xb2,
	0xed, 0x2e, 0xbd, 0x7c, 0xea, 0xe0, 0x44, 0xe8, 0xcf, 0x23, 0x10, 0x32, 0xe9, 0x26, 0xc0, 0xcd,
	0x81, 0xce, 0x1f, 0xa6, 0x73, 0x56, 0xfc, 0xcb, 0x1f, 0x85, 0xaf, 0xaa, 0xab, 0x3f, 0x96, 0xff,
	0xbf, 0x3d, 0x18, 0xfc, 0xb9, 0x7c, 0x7b, 0xdb, 0x2a, 0x07, 0x71, 0x11, 0xd8, 0x4b, 0x7d, 0x35,
	0xe9, 0x05, 0x63, 0x9f, 0xf9, 0x97, 0xcf, 0x39, 0x18, 0xc4, 0xc5, 0x41, 0x95, 0x73, 0x30, 0xe9,
	0x1d, 0x54, 0x39, 0x7f, 0x2f, 0xdf, 0xb3, 0x0f, 0xfa, 0xfd, 0x41, 0x5c, 0xf4, 0xfb, 0x55, 0x56,
	0xbf, 0x3f, 0xe9, 0xf5, 0xfb, 0x55, 0xde, 0x74, 0xc5, 0x14, 0xfb, 0xd9, 0x3f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xb8, 0x4e, 0x73, 0x21, 0xd4, 0x0c, 0x00, 0x00,
}
