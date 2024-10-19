// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: protobufs/atak.proto

package protobufs

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

type Team int32

const (
	// Unspecifed
	Team_Unspecifed_Color Team = 0
	// White
	Team_White Team = 1
	// Yellow
	Team_Yellow Team = 2
	// Orange
	Team_Orange Team = 3
	// Magenta
	Team_Magenta Team = 4
	// Red
	Team_Red Team = 5
	// Maroon
	Team_Maroon Team = 6
	// Purple
	Team_Purple Team = 7
	// Dark Blue
	Team_Dark_Blue Team = 8
	// Blue
	Team_Blue Team = 9
	// Cyan
	Team_Cyan Team = 10
	// Teal
	Team_Teal Team = 11
	// Green
	Team_Green Team = 12
	// Dark Green
	Team_Dark_Green Team = 13
	// Brown
	Team_Brown Team = 14
)

// Enum value maps for Team.
var (
	Team_name = map[int32]string{
		0:  "Unspecifed_Color",
		1:  "White",
		2:  "Yellow",
		3:  "Orange",
		4:  "Magenta",
		5:  "Red",
		6:  "Maroon",
		7:  "Purple",
		8:  "Dark_Blue",
		9:  "Blue",
		10: "Cyan",
		11: "Teal",
		12: "Green",
		13: "Dark_Green",
		14: "Brown",
	}
	Team_value = map[string]int32{
		"Unspecifed_Color": 0,
		"White":            1,
		"Yellow":           2,
		"Orange":           3,
		"Magenta":          4,
		"Red":              5,
		"Maroon":           6,
		"Purple":           7,
		"Dark_Blue":        8,
		"Blue":             9,
		"Cyan":             10,
		"Teal":             11,
		"Green":            12,
		"Dark_Green":       13,
		"Brown":            14,
	}
)

func (x Team) Enum() *Team {
	p := new(Team)
	*p = x
	return p
}

func (x Team) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Team) Descriptor() protoreflect.EnumDescriptor {
	return file_protobufs_atak_proto_enumTypes[0].Descriptor()
}

func (Team) Type() protoreflect.EnumType {
	return &file_protobufs_atak_proto_enumTypes[0]
}

func (x Team) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Team.Descriptor instead.
func (Team) EnumDescriptor() ([]byte, []int) {
	return file_protobufs_atak_proto_rawDescGZIP(), []int{0}
}

// Role of the group member
type MemberRole int32

const (
	// Unspecifed
	MemberRole_Unspecifed MemberRole = 0
	// Team Member
	MemberRole_TeamMember MemberRole = 1
	// Team Lead
	MemberRole_TeamLead MemberRole = 2
	// Headquarters
	MemberRole_HQ MemberRole = 3
	// Airsoft enthusiast
	MemberRole_Sniper MemberRole = 4
	// Medic
	MemberRole_Medic MemberRole = 5
	// ForwardObserver
	MemberRole_ForwardObserver MemberRole = 6
	// Radio Telephone Operator
	MemberRole_RTO MemberRole = 7
	// Doggo
	MemberRole_K9 MemberRole = 8
)

// Enum value maps for MemberRole.
var (
	MemberRole_name = map[int32]string{
		0: "Unspecifed",
		1: "TeamMember",
		2: "TeamLead",
		3: "HQ",
		4: "Sniper",
		5: "Medic",
		6: "ForwardObserver",
		7: "RTO",
		8: "K9",
	}
	MemberRole_value = map[string]int32{
		"Unspecifed":      0,
		"TeamMember":      1,
		"TeamLead":        2,
		"HQ":              3,
		"Sniper":          4,
		"Medic":           5,
		"ForwardObserver": 6,
		"RTO":             7,
		"K9":              8,
	}
)

func (x MemberRole) Enum() *MemberRole {
	p := new(MemberRole)
	*p = x
	return p
}

func (x MemberRole) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MemberRole) Descriptor() protoreflect.EnumDescriptor {
	return file_protobufs_atak_proto_enumTypes[1].Descriptor()
}

func (MemberRole) Type() protoreflect.EnumType {
	return &file_protobufs_atak_proto_enumTypes[1]
}

func (x MemberRole) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MemberRole.Descriptor instead.
func (MemberRole) EnumDescriptor() ([]byte, []int) {
	return file_protobufs_atak_proto_rawDescGZIP(), []int{1}
}

// Packets for the official ATAK Plugin
type TAKPacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Are the payloads strings compressed for LoRA transport?
	IsCompressed bool `protobuf:"varint,1,opt,name=is_compressed,json=isCompressed,proto3" json:"is_compressed,omitempty"`
	// The contact / callsign for ATAK user
	Contact *Contact `protobuf:"bytes,2,opt,name=contact,proto3" json:"contact,omitempty"`
	// The group for ATAK user
	Group *Group `protobuf:"bytes,3,opt,name=group,proto3" json:"group,omitempty"`
	// The status of the ATAK EUD
	Status *Status `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	// The payload of the packet
	//
	// Types that are assignable to PayloadVariant:
	//
	//	*TAKPacket_Pli
	//	*TAKPacket_Chat
	//	*TAKPacket_Detail
	PayloadVariant isTAKPacket_PayloadVariant `protobuf_oneof:"payload_variant"`
}

func (x *TAKPacket) Reset() {
	*x = TAKPacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobufs_atak_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TAKPacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TAKPacket) ProtoMessage() {}

func (x *TAKPacket) ProtoReflect() protoreflect.Message {
	mi := &file_protobufs_atak_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TAKPacket.ProtoReflect.Descriptor instead.
func (*TAKPacket) Descriptor() ([]byte, []int) {
	return file_protobufs_atak_proto_rawDescGZIP(), []int{0}
}

func (x *TAKPacket) GetIsCompressed() bool {
	if x != nil {
		return x.IsCompressed
	}
	return false
}

func (x *TAKPacket) GetContact() *Contact {
	if x != nil {
		return x.Contact
	}
	return nil
}

func (x *TAKPacket) GetGroup() *Group {
	if x != nil {
		return x.Group
	}
	return nil
}

func (x *TAKPacket) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (m *TAKPacket) GetPayloadVariant() isTAKPacket_PayloadVariant {
	if m != nil {
		return m.PayloadVariant
	}
	return nil
}

func (x *TAKPacket) GetPli() *PLI {
	if x, ok := x.GetPayloadVariant().(*TAKPacket_Pli); ok {
		return x.Pli
	}
	return nil
}

func (x *TAKPacket) GetChat() *GeoChat {
	if x, ok := x.GetPayloadVariant().(*TAKPacket_Chat); ok {
		return x.Chat
	}
	return nil
}

func (x *TAKPacket) GetDetail() []byte {
	if x, ok := x.GetPayloadVariant().(*TAKPacket_Detail); ok {
		return x.Detail
	}
	return nil
}

type isTAKPacket_PayloadVariant interface {
	isTAKPacket_PayloadVariant()
}

type TAKPacket_Pli struct {
	// TAK position report
	Pli *PLI `protobuf:"bytes,5,opt,name=pli,proto3,oneof"`
}

type TAKPacket_Chat struct {
	// ATAK GeoChat message
	Chat *GeoChat `protobuf:"bytes,6,opt,name=chat,proto3,oneof"`
}

type TAKPacket_Detail struct {
	// Generic CoT detail XML
	// May be compressed / truncated by the sender (EUD)
	Detail []byte `protobuf:"bytes,7,opt,name=detail,proto3,oneof"`
}

func (*TAKPacket_Pli) isTAKPacket_PayloadVariant() {}

func (*TAKPacket_Chat) isTAKPacket_PayloadVariant() {}

func (*TAKPacket_Detail) isTAKPacket_PayloadVariant() {}

// ATAK GeoChat message
type GeoChat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The text message
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	// Uid recipient of the message
	To *string `protobuf:"bytes,2,opt,name=to,proto3,oneof" json:"to,omitempty"`
	// Callsign of the recipient for the message
	ToCallsign *string `protobuf:"bytes,3,opt,name=to_callsign,json=toCallsign,proto3,oneof" json:"to_callsign,omitempty"`
}

func (x *GeoChat) Reset() {
	*x = GeoChat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobufs_atak_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeoChat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoChat) ProtoMessage() {}

func (x *GeoChat) ProtoReflect() protoreflect.Message {
	mi := &file_protobufs_atak_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeoChat.ProtoReflect.Descriptor instead.
func (*GeoChat) Descriptor() ([]byte, []int) {
	return file_protobufs_atak_proto_rawDescGZIP(), []int{1}
}

func (x *GeoChat) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GeoChat) GetTo() string {
	if x != nil && x.To != nil {
		return *x.To
	}
	return ""
}

func (x *GeoChat) GetToCallsign() string {
	if x != nil && x.ToCallsign != nil {
		return *x.ToCallsign
	}
	return ""
}

// ATAK Group
// <__group role='Team Member' name='Cyan'/>
type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Role of the group member
	Role MemberRole `protobuf:"varint,1,opt,name=role,proto3,enum=meshtastic.MemberRole" json:"role,omitempty"`
	// Team (color)
	// Default Cyan
	Team Team `protobuf:"varint,2,opt,name=team,proto3,enum=meshtastic.Team" json:"team,omitempty"`
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobufs_atak_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_protobufs_atak_proto_msgTypes[2]
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
	return file_protobufs_atak_proto_rawDescGZIP(), []int{2}
}

func (x *Group) GetRole() MemberRole {
	if x != nil {
		return x.Role
	}
	return MemberRole_Unspecifed
}

func (x *Group) GetTeam() Team {
	if x != nil {
		return x.Team
	}
	return Team_Unspecifed_Color
}

// ATAK EUD Status
// <status battery='100' />
type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Battery level
	Battery uint32 `protobuf:"varint,1,opt,name=battery,proto3" json:"battery,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobufs_atak_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_protobufs_atak_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_protobufs_atak_proto_rawDescGZIP(), []int{3}
}

func (x *Status) GetBattery() uint32 {
	if x != nil {
		return x.Battery
	}
	return 0
}

// ATAK Contact
// <contact endpoint='0.0.0.0:4242:tcp' phone='+12345678' callsign='FALKE'/>
type Contact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Callsign
	Callsign string `protobuf:"bytes,1,opt,name=callsign,proto3" json:"callsign,omitempty"`
	// Device callsign
	DeviceCallsign string `protobuf:"bytes,2,opt,name=device_callsign,json=deviceCallsign,proto3" json:"device_callsign,omitempty"`
}

func (x *Contact) Reset() {
	*x = Contact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobufs_atak_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contact) ProtoMessage() {}

func (x *Contact) ProtoReflect() protoreflect.Message {
	mi := &file_protobufs_atak_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contact.ProtoReflect.Descriptor instead.
func (*Contact) Descriptor() ([]byte, []int) {
	return file_protobufs_atak_proto_rawDescGZIP(), []int{4}
}

func (x *Contact) GetCallsign() string {
	if x != nil {
		return x.Callsign
	}
	return ""
}

func (x *Contact) GetDeviceCallsign() string {
	if x != nil {
		return x.DeviceCallsign
	}
	return ""
}

// Position Location Information from ATAK
type PLI struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The new preferred location encoding, multiply by 1e-7 to get degrees
	// in floating point
	LatitudeI int32 `protobuf:"fixed32,1,opt,name=latitude_i,json=latitudeI,proto3" json:"latitude_i,omitempty"`
	// The new preferred location encoding, multiply by 1e-7 to get degrees
	// in floating point
	LongitudeI int32 `protobuf:"fixed32,2,opt,name=longitude_i,json=longitudeI,proto3" json:"longitude_i,omitempty"`
	// Altitude (ATAK prefers HAE)
	Altitude int32 `protobuf:"varint,3,opt,name=altitude,proto3" json:"altitude,omitempty"`
	// Speed
	Speed uint32 `protobuf:"varint,4,opt,name=speed,proto3" json:"speed,omitempty"`
	// Course in degrees
	Course uint32 `protobuf:"varint,5,opt,name=course,proto3" json:"course,omitempty"`
}

func (x *PLI) Reset() {
	*x = PLI{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobufs_atak_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PLI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PLI) ProtoMessage() {}

func (x *PLI) ProtoReflect() protoreflect.Message {
	mi := &file_protobufs_atak_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PLI.ProtoReflect.Descriptor instead.
func (*PLI) Descriptor() ([]byte, []int) {
	return file_protobufs_atak_proto_rawDescGZIP(), []int{5}
}

func (x *PLI) GetLatitudeI() int32 {
	if x != nil {
		return x.LatitudeI
	}
	return 0
}

func (x *PLI) GetLongitudeI() int32 {
	if x != nil {
		return x.LongitudeI
	}
	return 0
}

func (x *PLI) GetAltitude() int32 {
	if x != nil {
		return x.Altitude
	}
	return 0
}

func (x *PLI) GetSpeed() uint32 {
	if x != nil {
		return x.Speed
	}
	return 0
}

func (x *PLI) GetCourse() uint32 {
	if x != nil {
		return x.Course
	}
	return 0
}

var File_protobufs_atak_proto protoreflect.FileDescriptor

var file_protobufs_atak_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x2f, 0x61, 0x74, 0x61, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74,
	0x69, 0x63, 0x22, 0xb1, 0x02, 0x0a, 0x09, 0x54, 0x41, 0x4b, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x12, 0x23, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x65, 0x64, 0x12, 0x2d, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x61, 0x73,
	0x74, 0x69, 0x63, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x12, 0x27, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63,
	0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x2a, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x6d, 0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x0a, 0x03, 0x70, 0x6c, 0x69,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x61, 0x73,
	0x74, 0x69, 0x63, 0x2e, 0x50, 0x4c, 0x49, 0x48, 0x00, 0x52, 0x03, 0x70, 0x6c, 0x69, 0x12, 0x29,
	0x0a, 0x04, 0x63, 0x68, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d,
	0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x47, 0x65, 0x6f, 0x43, 0x68, 0x61,
	0x74, 0x48, 0x00, 0x52, 0x04, 0x63, 0x68, 0x61, 0x74, 0x12, 0x18, 0x0a, 0x06, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x06, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x42, 0x11, 0x0a, 0x0f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x76,
	0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x22, 0x75, 0x0a, 0x07, 0x47, 0x65, 0x6f, 0x43, 0x68, 0x61,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x13, 0x0a, 0x02, 0x74,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x02, 0x74, 0x6f, 0x88, 0x01, 0x01,
	0x12, 0x24, 0x0a, 0x0b, 0x74, 0x6f, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x69, 0x67, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0a, 0x74, 0x6f, 0x43, 0x61, 0x6c, 0x6c, 0x73,
	0x69, 0x67, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x74, 0x6f, 0x42, 0x0e, 0x0a,
	0x0c, 0x5f, 0x74, 0x6f, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x69, 0x67, 0x6e, 0x22, 0x59, 0x0a,
	0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x2a, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69,
	0x63, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x10, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x54, 0x65,
	0x61, 0x6d, 0x52, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x22, 0x22, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x22, 0x4e, 0x0a, 0x07,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x73,
	0x69, 0x67, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x73,
	0x69, 0x67, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x61,
	0x6c, 0x6c, 0x73, 0x69, 0x67, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x73, 0x69, 0x67, 0x6e, 0x22, 0x8f, 0x01, 0x0a,
	0x03, 0x50, 0x4c, 0x49, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x5f, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0f, 0x52, 0x09, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x49, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x5f, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0f, 0x52, 0x0a, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x49, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x61, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2a, 0xc0,
	0x01, 0x0a, 0x04, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x6e, 0x73, 0x70, 0x65,
	0x63, 0x69, 0x66, 0x65, 0x64, 0x5f, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x10, 0x00, 0x12, 0x09, 0x0a,
	0x05, 0x57, 0x68, 0x69, 0x74, 0x65, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x59, 0x65, 0x6c, 0x6c,
	0x6f, 0x77, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x4f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x10, 0x03,
	0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x61, 0x10, 0x04, 0x12, 0x07, 0x0a,
	0x03, 0x52, 0x65, 0x64, 0x10, 0x05, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x61, 0x72, 0x6f, 0x6f, 0x6e,
	0x10, 0x06, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x75, 0x72, 0x70, 0x6c, 0x65, 0x10, 0x07, 0x12, 0x0d,
	0x0a, 0x09, 0x44, 0x61, 0x72, 0x6b, 0x5f, 0x42, 0x6c, 0x75, 0x65, 0x10, 0x08, 0x12, 0x08, 0x0a,
	0x04, 0x42, 0x6c, 0x75, 0x65, 0x10, 0x09, 0x12, 0x08, 0x0a, 0x04, 0x43, 0x79, 0x61, 0x6e, 0x10,
	0x0a, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x65, 0x61, 0x6c, 0x10, 0x0b, 0x12, 0x09, 0x0a, 0x05, 0x47,
	0x72, 0x65, 0x65, 0x6e, 0x10, 0x0c, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x61, 0x72, 0x6b, 0x5f, 0x47,
	0x72, 0x65, 0x65, 0x6e, 0x10, 0x0d, 0x12, 0x09, 0x0a, 0x05, 0x42, 0x72, 0x6f, 0x77, 0x6e, 0x10,
	0x0e, 0x2a, 0x7f, 0x0a, 0x0a, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x12,
	0x0e, 0x0a, 0x0a, 0x55, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x65, 0x64, 0x10, 0x00, 0x12,
	0x0e, 0x0a, 0x0a, 0x54, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x10, 0x01, 0x12,
	0x0c, 0x0a, 0x08, 0x54, 0x65, 0x61, 0x6d, 0x4c, 0x65, 0x61, 0x64, 0x10, 0x02, 0x12, 0x06, 0x0a,
	0x02, 0x48, 0x51, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x6e, 0x69, 0x70, 0x65, 0x72, 0x10,
	0x04, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x65, 0x64, 0x69, 0x63, 0x10, 0x05, 0x12, 0x13, 0x0a, 0x0f,
	0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x10,
	0x06, 0x12, 0x07, 0x0a, 0x03, 0x52, 0x54, 0x4f, 0x10, 0x07, 0x12, 0x06, 0x0a, 0x02, 0x4b, 0x39,
	0x10, 0x08, 0x42, 0x86, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x65, 0x65, 0x6b, 0x73,
	0x76, 0x69, 0x6c, 0x6c, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x42, 0x0a, 0x41, 0x54, 0x41, 0x4b,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x2e, 0x6a, 0x61, 0x6e, 0x6b,
	0x79, 0x2e, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6e, 0x6e,
	0x2f, 0x6d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74,
	0x69, 0x63, 0x2d, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2d, 0x67, 0x6f, 0x2f, 0x6d, 0x65, 0x73,
	0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x73, 0xaa, 0x02, 0x14, 0x4d, 0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0xba, 0x02, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_protobufs_atak_proto_rawDescOnce sync.Once
	file_protobufs_atak_proto_rawDescData = file_protobufs_atak_proto_rawDesc
)

func file_protobufs_atak_proto_rawDescGZIP() []byte {
	file_protobufs_atak_proto_rawDescOnce.Do(func() {
		file_protobufs_atak_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobufs_atak_proto_rawDescData)
	})
	return file_protobufs_atak_proto_rawDescData
}

var file_protobufs_atak_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_protobufs_atak_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protobufs_atak_proto_goTypes = []interface{}{
	(Team)(0),         // 0: meshtastic.Team
	(MemberRole)(0),   // 1: meshtastic.MemberRole
	(*TAKPacket)(nil), // 2: meshtastic.TAKPacket
	(*GeoChat)(nil),   // 3: meshtastic.GeoChat
	(*Group)(nil),     // 4: meshtastic.Group
	(*Status)(nil),    // 5: meshtastic.Status
	(*Contact)(nil),   // 6: meshtastic.Contact
	(*PLI)(nil),       // 7: meshtastic.PLI
}
var file_protobufs_atak_proto_depIdxs = []int32{
	6, // 0: meshtastic.TAKPacket.contact:type_name -> meshtastic.Contact
	4, // 1: meshtastic.TAKPacket.group:type_name -> meshtastic.Group
	5, // 2: meshtastic.TAKPacket.status:type_name -> meshtastic.Status
	7, // 3: meshtastic.TAKPacket.pli:type_name -> meshtastic.PLI
	3, // 4: meshtastic.TAKPacket.chat:type_name -> meshtastic.GeoChat
	1, // 5: meshtastic.Group.role:type_name -> meshtastic.MemberRole
	0, // 6: meshtastic.Group.team:type_name -> meshtastic.Team
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_protobufs_atak_proto_init() }
func file_protobufs_atak_proto_init() {
	if File_protobufs_atak_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobufs_atak_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TAKPacket); i {
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
		file_protobufs_atak_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeoChat); i {
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
		file_protobufs_atak_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_protobufs_atak_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_protobufs_atak_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Contact); i {
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
		file_protobufs_atak_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PLI); i {
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
	file_protobufs_atak_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*TAKPacket_Pli)(nil),
		(*TAKPacket_Chat)(nil),
		(*TAKPacket_Detail)(nil),
	}
	file_protobufs_atak_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protobufs_atak_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobufs_atak_proto_goTypes,
		DependencyIndexes: file_protobufs_atak_proto_depIdxs,
		EnumInfos:         file_protobufs_atak_proto_enumTypes,
		MessageInfos:      file_protobufs_atak_proto_msgTypes,
	}.Build()
	File_protobufs_atak_proto = out.File
	file_protobufs_atak_proto_rawDesc = nil
	file_protobufs_atak_proto_goTypes = nil
	file_protobufs_atak_proto_depIdxs = nil
}
