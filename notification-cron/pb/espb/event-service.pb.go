// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: proto/event-service.proto

package espb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DateRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DateFrom string `protobuf:"bytes,1,opt,name=date_from,json=dateFrom,proto3" json:"date_from,omitempty"`
	DateTo   string `protobuf:"bytes,2,opt,name=date_to,json=dateTo,proto3" json:"date_to,omitempty"`
}

func (x *DateRange) Reset() {
	*x = DateRange{}
	mi := &file_proto_event_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DateRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DateRange) ProtoMessage() {}

func (x *DateRange) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DateRange.ProtoReflect.Descriptor instead.
func (*DateRange) Descriptor() ([]byte, []int) {
	return file_proto_event_service_proto_rawDescGZIP(), []int{0}
}

func (x *DateRange) GetDateFrom() string {
	if x != nil {
		return x.DateFrom
	}
	return ""
}

func (x *DateRange) GetDateTo() string {
	if x != nil {
		return x.DateTo
	}
	return ""
}

type ParticipantRequirements struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gender bool   `protobuf:"varint,1,opt,name=gender,proto3" json:"gender,omitempty"`
	MinAge *int32 `protobuf:"varint,2,opt,name=min_age,json=minAge,proto3,oneof" json:"min_age,omitempty"`
	MaxAge *int32 `protobuf:"varint,3,opt,name=max_age,json=maxAge,proto3,oneof" json:"max_age,omitempty"`
}

func (x *ParticipantRequirements) Reset() {
	*x = ParticipantRequirements{}
	mi := &file_proto_event_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ParticipantRequirements) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParticipantRequirements) ProtoMessage() {}

func (x *ParticipantRequirements) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParticipantRequirements.ProtoReflect.Descriptor instead.
func (*ParticipantRequirements) Descriptor() ([]byte, []int) {
	return file_proto_event_service_proto_rawDescGZIP(), []int{1}
}

func (x *ParticipantRequirements) GetGender() bool {
	if x != nil {
		return x.Gender
	}
	return false
}

func (x *ParticipantRequirements) GetMinAge() int32 {
	if x != nil && x.MinAge != nil {
		return *x.MinAge
	}
	return 0
}

func (x *ParticipantRequirements) GetMaxAge() int32 {
	if x != nil && x.MaxAge != nil {
		return *x.MaxAge
	}
	return 0
}

type SportType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SportType) Reset() {
	*x = SportType{}
	mi := &file_proto_event_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SportType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SportType) ProtoMessage() {}

func (x *SportType) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SportType.ProtoReflect.Descriptor instead.
func (*SportType) Descriptor() ([]byte, []int) {
	return file_proto_event_service_proto_rawDescGZIP(), []int{2}
}

func (x *SportType) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SportType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SportSubtype struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Parent *SportType `protobuf:"bytes,3,opt,name=parent,proto3" json:"parent,omitempty"`
}

func (x *SportSubtype) Reset() {
	*x = SportSubtype{}
	mi := &file_proto_event_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SportSubtype) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SportSubtype) ProtoMessage() {}

func (x *SportSubtype) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SportSubtype.ProtoReflect.Descriptor instead.
func (*SportSubtype) Descriptor() ([]byte, []int) {
	return file_proto_event_service_proto_rawDescGZIP(), []int{3}
}

func (x *SportSubtype) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SportSubtype) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SportSubtype) GetParent() *SportType {
	if x != nil {
		return x.Parent
	}
	return nil
}

type SportTypeWithSubtypes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Subtypes []*SportSubtype2 `protobuf:"bytes,3,rep,name=subtypes,proto3" json:"subtypes,omitempty"`
}

func (x *SportTypeWithSubtypes) Reset() {
	*x = SportTypeWithSubtypes{}
	mi := &file_proto_event_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SportTypeWithSubtypes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SportTypeWithSubtypes) ProtoMessage() {}

func (x *SportTypeWithSubtypes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SportTypeWithSubtypes.ProtoReflect.Descriptor instead.
func (*SportTypeWithSubtypes) Descriptor() ([]byte, []int) {
	return file_proto_event_service_proto_rawDescGZIP(), []int{4}
}

func (x *SportTypeWithSubtypes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SportTypeWithSubtypes) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SportTypeWithSubtypes) GetSubtypes() []*SportSubtype2 {
	if x != nil {
		return x.Subtypes
	}
	return nil
}

type SportSubtype2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SportSubtype2) Reset() {
	*x = SportSubtype2{}
	mi := &file_proto_event_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SportSubtype2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SportSubtype2) ProtoMessage() {}

func (x *SportSubtype2) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SportSubtype2.ProtoReflect.Descriptor instead.
func (*SportSubtype2) Descriptor() ([]byte, []int) {
	return file_proto_event_service_proto_rawDescGZIP(), []int{5}
}

func (x *SportSubtype2) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SportSubtype2) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type EventInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EkpId                   string                     `protobuf:"bytes,1,opt,name=ekpId,proto3" json:"ekpId,omitempty"`
	SportSubtype            *SportSubtype              `protobuf:"bytes,3,opt,name=sportSubtype,proto3" json:"sportSubtype,omitempty"`
	Name                    string                     `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Description             string                     `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Dates                   *DateRange                 `protobuf:"bytes,6,opt,name=dates,proto3" json:"dates,omitempty"`
	Location                string                     `protobuf:"bytes,7,opt,name=location,proto3" json:"location,omitempty"`
	Participants            int32                      `protobuf:"varint,8,opt,name=participants,proto3" json:"participants,omitempty"`
	ParticipantRequirements []*ParticipantRequirements `protobuf:"bytes,9,rep,name=participantRequirements,proto3" json:"participantRequirements,omitempty"`
	Id                      string                     `protobuf:"bytes,10,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *EventInfo) Reset() {
	*x = EventInfo{}
	mi := &file_proto_event_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventInfo) ProtoMessage() {}

func (x *EventInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventInfo.ProtoReflect.Descriptor instead.
func (*EventInfo) Descriptor() ([]byte, []int) {
	return file_proto_event_service_proto_rawDescGZIP(), []int{6}
}

func (x *EventInfo) GetEkpId() string {
	if x != nil {
		return x.EkpId
	}
	return ""
}

func (x *EventInfo) GetSportSubtype() *SportSubtype {
	if x != nil {
		return x.SportSubtype
	}
	return nil
}

func (x *EventInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EventInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *EventInfo) GetDates() *DateRange {
	if x != nil {
		return x.Dates
	}
	return nil
}

func (x *EventInfo) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *EventInfo) GetParticipants() int32 {
	if x != nil {
		return x.Participants
	}
	return 0
}

func (x *EventInfo) GetParticipantRequirements() []*ParticipantRequirements {
	if x != nil {
		return x.ParticipantRequirements
	}
	return nil
}

func (x *EventInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type LoadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EkpId                   string                     `protobuf:"bytes,1,opt,name=ekpId,proto3" json:"ekpId,omitempty"`
	SportType               string                     `protobuf:"bytes,2,opt,name=sportType,proto3" json:"sportType,omitempty"`
	SportSubtype            string                     `protobuf:"bytes,3,opt,name=sportSubtype,proto3" json:"sportSubtype,omitempty"`
	Name                    string                     `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Description             string                     `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Dates                   *DateRange                 `protobuf:"bytes,6,opt,name=dates,proto3" json:"dates,omitempty"`
	Location                string                     `protobuf:"bytes,7,opt,name=location,proto3" json:"location,omitempty"`
	Participants            int32                      `protobuf:"varint,8,opt,name=participants,proto3" json:"participants,omitempty"`
	ParticipantRequirements []*ParticipantRequirements `protobuf:"bytes,9,rep,name=participantRequirements,proto3" json:"participantRequirements,omitempty"`
}

func (x *LoadRequest) Reset() {
	*x = LoadRequest{}
	mi := &file_proto_event_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadRequest) ProtoMessage() {}

func (x *LoadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadRequest.ProtoReflect.Descriptor instead.
func (*LoadRequest) Descriptor() ([]byte, []int) {
	return file_proto_event_service_proto_rawDescGZIP(), []int{7}
}

func (x *LoadRequest) GetEkpId() string {
	if x != nil {
		return x.EkpId
	}
	return ""
}

func (x *LoadRequest) GetSportType() string {
	if x != nil {
		return x.SportType
	}
	return ""
}

func (x *LoadRequest) GetSportSubtype() string {
	if x != nil {
		return x.SportSubtype
	}
	return ""
}

func (x *LoadRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LoadRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *LoadRequest) GetDates() *DateRange {
	if x != nil {
		return x.Dates
	}
	return nil
}

func (x *LoadRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *LoadRequest) GetParticipants() int32 {
	if x != nil {
		return x.Participants
	}
	return 0
}

func (x *LoadRequest) GetParticipantRequirements() []*ParticipantRequirements {
	if x != nil {
		return x.ParticipantRequirements
	}
	return nil
}

type LoadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Saved int32 `protobuf:"varint,1,opt,name=saved,proto3" json:"saved,omitempty"`
}

func (x *LoadResponse) Reset() {
	*x = LoadResponse{}
	mi := &file_proto_event_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadResponse) ProtoMessage() {}

func (x *LoadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadResponse.ProtoReflect.Descriptor instead.
func (*LoadResponse) Descriptor() ([]byte, []int) {
	return file_proto_event_service_proto_rawDescGZIP(), []int{8}
}

func (x *LoadResponse) GetSaved() int32 {
	if x != nil {
		return x.Saved
	}
	return 0
}

type EventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *EventRequest) Reset() {
	*x = EventRequest{}
	mi := &file_proto_event_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventRequest) ProtoMessage() {}

func (x *EventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventRequest.ProtoReflect.Descriptor instead.
func (*EventRequest) Descriptor() ([]byte, []int) {
	return file_proto_event_service_proto_rawDescGZIP(), []int{9}
}

func (x *EventRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type EventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *EventInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *EventResponse) Reset() {
	*x = EventResponse{}
	mi := &file_proto_event_service_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventResponse) ProtoMessage() {}

func (x *EventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_service_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventResponse.ProtoReflect.Descriptor instead.
func (*EventResponse) Descriptor() ([]byte, []int) {
	return file_proto_event_service_proto_rawDescGZIP(), []int{10}
}

func (x *EventResponse) GetInfo() *EventInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type UpcomingEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event    *EventInfo `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	DaysLeft uint32     `protobuf:"varint,2,opt,name=daysLeft,proto3" json:"daysLeft,omitempty"` // Кол-во дней до эвента
}

func (x *UpcomingEventResponse) Reset() {
	*x = UpcomingEventResponse{}
	mi := &file_proto_event_service_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpcomingEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpcomingEventResponse) ProtoMessage() {}

func (x *UpcomingEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_service_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpcomingEventResponse.ProtoReflect.Descriptor instead.
func (*UpcomingEventResponse) Descriptor() ([]byte, []int) {
	return file_proto_event_service_proto_rawDescGZIP(), []int{11}
}

func (x *UpcomingEventResponse) GetEvent() *EventInfo {
	if x != nil {
		return x.Event
	}
	return nil
}

func (x *UpcomingEventResponse) GetDaysLeft() uint32 {
	if x != nil {
		return x.DaysLeft
	}
	return 0
}

type SportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SportRequest) Reset() {
	*x = SportRequest{}
	mi := &file_proto_event_service_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SportRequest) ProtoMessage() {}

func (x *SportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_service_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SportRequest.ProtoReflect.Descriptor instead.
func (*SportRequest) Descriptor() ([]byte, []int) {
	return file_proto_event_service_proto_rawDescGZIP(), []int{12}
}

func (x *SportRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SportType *SportTypeWithSubtypes `protobuf:"bytes,1,opt,name=sportType,proto3" json:"sportType,omitempty"`
}

func (x *SportResponse) Reset() {
	*x = SportResponse{}
	mi := &file_proto_event_service_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SportResponse) ProtoMessage() {}

func (x *SportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_service_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SportResponse.ProtoReflect.Descriptor instead.
func (*SportResponse) Descriptor() ([]byte, []int) {
	return file_proto_event_service_proto_rawDescGZIP(), []int{13}
}

func (x *SportResponse) GetSportType() *SportTypeWithSubtypes {
	if x != nil {
		return x.SportType
	}
	return nil
}

var File_proto_event_service_proto protoreflect.FileDescriptor

var file_proto_event_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x41, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x6f, 0x22, 0x85, 0x01, 0x0a, 0x17, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70,
	0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x07, 0x6d, 0x69, 0x6e, 0x5f, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x06, 0x6d, 0x69, 0x6e, 0x41,
	0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a, 0x07, 0x6d, 0x61, 0x78, 0x5f, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x06, 0x6d, 0x61, 0x78, 0x41, 0x67, 0x65,
	0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6d, 0x69, 0x6e, 0x5f, 0x61, 0x67, 0x65, 0x42,
	0x0a, 0x0a, 0x08, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x61, 0x67, 0x65, 0x22, 0x2f, 0x0a, 0x09, 0x53,
	0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x5d, 0x0a, 0x0c,
	0x53, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x75, 0x62, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x29, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x22, 0x6e, 0x0a, 0x15, 0x53,
	0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x57, 0x69, 0x74, 0x68, 0x53, 0x75, 0x62, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x75, 0x62, 0x74, 0x79, 0x70, 0x65,
	0x32, 0x52, 0x08, 0x73, 0x75, 0x62, 0x74, 0x79, 0x70, 0x65, 0x73, 0x22, 0x33, 0x0a, 0x0d, 0x53,
	0x70, 0x6f, 0x72, 0x74, 0x53, 0x75, 0x62, 0x74, 0x79, 0x70, 0x65, 0x32, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0xe5, 0x02, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6b, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6b, 0x70, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x0c, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x75, 0x62,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x75, 0x62, 0x74, 0x79, 0x70, 0x65,
	0x52, 0x0c, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x75, 0x62, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x05, 0x64, 0x61, 0x74, 0x65, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x44, 0x61, 0x74,
	0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x05, 0x64, 0x61, 0x74, 0x65, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x59, 0x0a,
	0x17, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70,
	0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x17, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xdf, 0x02, 0x0a, 0x0b, 0x4c, 0x6f, 0x61,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6b, 0x70, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6b, 0x70, 0x49, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0c,
	0x73, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x75, 0x62, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x75, 0x62, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x05, 0x64, 0x61, 0x74, 0x65, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x44,
	0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x05, 0x64, 0x61, 0x74, 0x65, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x12,
	0x59, 0x0a, 0x17, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x69, 0x70, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x17, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x24, 0x0a, 0x0c, 0x4c, 0x6f,
	0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x61,
	0x76, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x61, 0x76, 0x65, 0x64,
	0x22, 0x1e, 0x0a, 0x0c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x36, 0x0a, 0x0d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x25, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x5c, 0x0a, 0x15, 0x55, 0x70, 0x63, 0x6f,
	0x6d, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x27, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61,
	0x79, 0x73, 0x4c, 0x65, 0x66, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x64, 0x61,
	0x79, 0x73, 0x4c, 0x65, 0x66, 0x74, 0x22, 0x1e, 0x0a, 0x0c, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4c, 0x0a, 0x0d, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x09, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x57, 0x69, 0x74,
	0x68, 0x53, 0x75, 0x62, 0x74, 0x79, 0x70, 0x65, 0x73, 0x52, 0x09, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x32, 0xbd, 0x02, 0x0a, 0x0c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x4c, 0x6f, 0x61, 0x64, 0x12, 0x13, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x4c, 0x6f, 0x61, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x34, 0x0a, 0x05, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x14, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x39, 0x0a, 0x06, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x14, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x12, 0x39, 0x0a, 0x06, 0x53,
	0x70, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x14, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x53,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x12, 0x4c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x55, 0x70, 0x63,
	0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x1d, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x55, 0x70, 0x63,
	0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x30, 0x01, 0x42, 0x1f, 0x5a, 0x1d, 0x6d, 0x7a, 0x68, 0x6e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x3b, 0x65, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_event_service_proto_rawDescOnce sync.Once
	file_proto_event_service_proto_rawDescData = file_proto_event_service_proto_rawDesc
)

func file_proto_event_service_proto_rawDescGZIP() []byte {
	file_proto_event_service_proto_rawDescOnce.Do(func() {
		file_proto_event_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_event_service_proto_rawDescData)
	})
	return file_proto_event_service_proto_rawDescData
}

var file_proto_event_service_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_proto_event_service_proto_goTypes = []any{
	(*DateRange)(nil),               // 0: events.DateRange
	(*ParticipantRequirements)(nil), // 1: events.ParticipantRequirements
	(*SportType)(nil),               // 2: events.SportType
	(*SportSubtype)(nil),            // 3: events.SportSubtype
	(*SportTypeWithSubtypes)(nil),   // 4: events.SportTypeWithSubtypes
	(*SportSubtype2)(nil),           // 5: events.SportSubtype2
	(*EventInfo)(nil),               // 6: events.EventInfo
	(*LoadRequest)(nil),             // 7: events.LoadRequest
	(*LoadResponse)(nil),            // 8: events.LoadResponse
	(*EventRequest)(nil),            // 9: events.EventRequest
	(*EventResponse)(nil),           // 10: events.EventResponse
	(*UpcomingEventResponse)(nil),   // 11: events.UpcomingEventResponse
	(*SportRequest)(nil),            // 12: events.SportRequest
	(*SportResponse)(nil),           // 13: events.SportResponse
	(*emptypb.Empty)(nil),           // 14: google.protobuf.Empty
}
var file_proto_event_service_proto_depIdxs = []int32{
	2,  // 0: events.SportSubtype.parent:type_name -> events.SportType
	5,  // 1: events.SportTypeWithSubtypes.subtypes:type_name -> events.SportSubtype2
	3,  // 2: events.EventInfo.sportSubtype:type_name -> events.SportSubtype
	0,  // 3: events.EventInfo.dates:type_name -> events.DateRange
	1,  // 4: events.EventInfo.participantRequirements:type_name -> events.ParticipantRequirements
	0,  // 5: events.LoadRequest.dates:type_name -> events.DateRange
	1,  // 6: events.LoadRequest.participantRequirements:type_name -> events.ParticipantRequirements
	6,  // 7: events.EventResponse.info:type_name -> events.EventInfo
	6,  // 8: events.UpcomingEventResponse.event:type_name -> events.EventInfo
	4,  // 9: events.SportResponse.sportType:type_name -> events.SportTypeWithSubtypes
	7,  // 10: events.EventService.Load:input_type -> events.LoadRequest
	9,  // 11: events.EventService.Event:input_type -> events.EventRequest
	9,  // 12: events.EventService.Events:input_type -> events.EventRequest
	12, // 13: events.EventService.Sports:input_type -> events.SportRequest
	14, // 14: events.EventService.GetUpcomingEvents:input_type -> google.protobuf.Empty
	8,  // 15: events.EventService.Load:output_type -> events.LoadResponse
	10, // 16: events.EventService.Event:output_type -> events.EventResponse
	10, // 17: events.EventService.Events:output_type -> events.EventResponse
	13, // 18: events.EventService.Sports:output_type -> events.SportResponse
	11, // 19: events.EventService.GetUpcomingEvents:output_type -> events.UpcomingEventResponse
	15, // [15:20] is the sub-list for method output_type
	10, // [10:15] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_proto_event_service_proto_init() }
func file_proto_event_service_proto_init() {
	if File_proto_event_service_proto != nil {
		return
	}
	file_proto_event_service_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_event_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_event_service_proto_goTypes,
		DependencyIndexes: file_proto_event_service_proto_depIdxs,
		MessageInfos:      file_proto_event_service_proto_msgTypes,
	}.Build()
	File_proto_event_service_proto = out.File
	file_proto_event_service_proto_rawDesc = nil
	file_proto_event_service_proto_goTypes = nil
	file_proto_event_service_proto_depIdxs = nil
}
