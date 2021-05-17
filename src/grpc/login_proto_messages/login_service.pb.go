// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.0
// source: login_service.proto

package login_proto_messages

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

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_login_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_login_service_proto_rawDescGZIP(), []int{0}
}

func (x *Error) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Type     string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_login_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_login_service_proto_rawDescGZIP(), []int{1}
}

func (x *LoginRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *LoginRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayData *PlayData `protobuf:"bytes,1,opt,name=play_data,json=playData,proto3" json:"play_data,omitempty"`
	Session  *Session  `protobuf:"bytes,2,opt,name=session,proto3" json:"session,omitempty"`
	Error    *Error    `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_login_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_login_service_proto_rawDescGZIP(), []int{2}
}

func (x *LoginResponse) GetPlayData() *PlayData {
	if x != nil {
		return x.PlayData
	}
	return nil
}

func (x *LoginResponse) GetSession() *Session {
	if x != nil {
		return x.Session
	}
	return nil
}

func (x *LoginResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type PlayData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Characters []*Character `protobuf:"bytes,1,rep,name=characters,proto3" json:"characters,omitempty"`
	Worlds     []*World     `protobuf:"bytes,2,rep,name=worlds,proto3" json:"worlds,omitempty"`
}

func (x *PlayData) Reset() {
	*x = PlayData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayData) ProtoMessage() {}

func (x *PlayData) ProtoReflect() protoreflect.Message {
	mi := &file_login_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayData.ProtoReflect.Descriptor instead.
func (*PlayData) Descriptor() ([]byte, []int) {
	return file_login_service_proto_rawDescGZIP(), []int{3}
}

func (x *PlayData) GetCharacters() []*Character {
	if x != nil {
		return x.Characters
	}
	return nil
}

func (x *PlayData) GetWorlds() []*World {
	if x != nil {
		return x.Worlds
	}
	return nil
}

type Session struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsPremium    bool   `protobuf:"varint,1,opt,name=is_premium,json=isPremium,proto3" json:"is_premium,omitempty"`
	PremiumUntil uint32 `protobuf:"varint,2,opt,name=premium_until,json=premiumUntil,proto3" json:"premium_until,omitempty"`
	SessionKey   string `protobuf:"bytes,3,opt,name=session_key,json=sessionKey,proto3" json:"session_key,omitempty"`
	LastLogin    uint32 `protobuf:"varint,4,opt,name=last_login,json=lastLogin,proto3" json:"last_login,omitempty"`
}

func (x *Session) Reset() {
	*x = Session{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
	mi := &file_login_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session.ProtoReflect.Descriptor instead.
func (*Session) Descriptor() ([]byte, []int) {
	return file_login_service_proto_rawDescGZIP(), []int{4}
}

func (x *Session) GetIsPremium() bool {
	if x != nil {
		return x.IsPremium
	}
	return false
}

func (x *Session) GetPremiumUntil() uint32 {
	if x != nil {
		return x.PremiumUntil
	}
	return 0
}

func (x *Session) GetSessionKey() string {
	if x != nil {
		return x.SessionKey
	}
	return ""
}

func (x *Session) GetLastLogin() uint32 {
	if x != nil {
		return x.LastLogin
	}
	return 0
}

type Character struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorldId uint32           `protobuf:"varint,1,opt,name=world_id,json=worldId,proto3" json:"world_id,omitempty"`
	Info    *CharacterInfo   `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	Outfit  *CharacterOutfit `protobuf:"bytes,3,opt,name=outfit,proto3" json:"outfit,omitempty"`
}

func (x *Character) Reset() {
	*x = Character{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Character) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Character) ProtoMessage() {}

func (x *Character) ProtoReflect() protoreflect.Message {
	mi := &file_login_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Character.ProtoReflect.Descriptor instead.
func (*Character) Descriptor() ([]byte, []int) {
	return file_login_service_proto_rawDescGZIP(), []int{5}
}

func (x *Character) GetWorldId() uint32 {
	if x != nil {
		return x.WorldId
	}
	return 0
}

func (x *Character) GetInfo() *CharacterInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *Character) GetOutfit() *CharacterOutfit {
	if x != nil {
		return x.Outfit
	}
	return nil
}

type CharacterInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	LastLogin uint32 `protobuf:"varint,2,opt,name=last_login,json=lastLogin,proto3" json:"last_login,omitempty"`
	Level     uint32 `protobuf:"varint,3,opt,name=level,proto3" json:"level,omitempty"`
	Vocation  string `protobuf:"bytes,4,opt,name=vocation,proto3" json:"vocation,omitempty"`
	Sex       uint32 `protobuf:"varint,5,opt,name=sex,proto3" json:"sex,omitempty"`
}

func (x *CharacterInfo) Reset() {
	*x = CharacterInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CharacterInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CharacterInfo) ProtoMessage() {}

func (x *CharacterInfo) ProtoReflect() protoreflect.Message {
	mi := &file_login_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CharacterInfo.ProtoReflect.Descriptor instead.
func (*CharacterInfo) Descriptor() ([]byte, []int) {
	return file_login_service_proto_rawDescGZIP(), []int{6}
}

func (x *CharacterInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CharacterInfo) GetLastLogin() uint32 {
	if x != nil {
		return x.LastLogin
	}
	return 0
}

func (x *CharacterInfo) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *CharacterInfo) GetVocation() string {
	if x != nil {
		return x.Vocation
	}
	return ""
}

func (x *CharacterInfo) GetSex() uint32 {
	if x != nil {
		return x.Sex
	}
	return 0
}

type CharacterOutfit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LookType uint32 `protobuf:"varint,1,opt,name=look_type,json=lookType,proto3" json:"look_type,omitempty"`
	LookHead uint32 `protobuf:"varint,2,opt,name=look_head,json=lookHead,proto3" json:"look_head,omitempty"`
	LookBody uint32 `protobuf:"varint,3,opt,name=look_body,json=lookBody,proto3" json:"look_body,omitempty"`
	LookLegs uint32 `protobuf:"varint,4,opt,name=look_legs,json=lookLegs,proto3" json:"look_legs,omitempty"`
	LookFeet uint32 `protobuf:"varint,5,opt,name=look_feet,json=lookFeet,proto3" json:"look_feet,omitempty"`
	Addons   uint32 `protobuf:"varint,6,opt,name=addons,proto3" json:"addons,omitempty"`
}

func (x *CharacterOutfit) Reset() {
	*x = CharacterOutfit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CharacterOutfit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CharacterOutfit) ProtoMessage() {}

func (x *CharacterOutfit) ProtoReflect() protoreflect.Message {
	mi := &file_login_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CharacterOutfit.ProtoReflect.Descriptor instead.
func (*CharacterOutfit) Descriptor() ([]byte, []int) {
	return file_login_service_proto_rawDescGZIP(), []int{7}
}

func (x *CharacterOutfit) GetLookType() uint32 {
	if x != nil {
		return x.LookType
	}
	return 0
}

func (x *CharacterOutfit) GetLookHead() uint32 {
	if x != nil {
		return x.LookHead
	}
	return 0
}

func (x *CharacterOutfit) GetLookBody() uint32 {
	if x != nil {
		return x.LookBody
	}
	return 0
}

func (x *CharacterOutfit) GetLookLegs() uint32 {
	if x != nil {
		return x.LookLegs
	}
	return 0
}

func (x *CharacterOutfit) GetLookFeet() uint32 {
	if x != nil {
		return x.LookFeet
	}
	return 0
}

func (x *CharacterOutfit) GetAddons() uint32 {
	if x != nil {
		return x.Addons
	}
	return 0
}

type World struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                         uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ExternalAddress            string `protobuf:"bytes,3,opt,name=external_address,json=externalAddress,proto3" json:"external_address,omitempty"`
	ExternalAddressProtected   string `protobuf:"bytes,4,opt,name=external_address_protected,json=externalAddressProtected,proto3" json:"external_address_protected,omitempty"`
	ExternalAddressUnprotected string `protobuf:"bytes,5,opt,name=external_address_unprotected,json=externalAddressUnprotected,proto3" json:"external_address_unprotected,omitempty"`
	ExternalPort               uint32 `protobuf:"varint,6,opt,name=external_port,json=externalPort,proto3" json:"external_port,omitempty"`
	ExternalPortProtected      uint32 `protobuf:"varint,7,opt,name=external_port_protected,json=externalPortProtected,proto3" json:"external_port_protected,omitempty"`
	ExternalPortUnprotected    uint32 `protobuf:"varint,8,opt,name=external_port_unprotected,json=externalPortUnprotected,proto3" json:"external_port_unprotected,omitempty"`
	Location                   string `protobuf:"bytes,9,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *World) Reset() {
	*x = World{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *World) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*World) ProtoMessage() {}

func (x *World) ProtoReflect() protoreflect.Message {
	mi := &file_login_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use World.ProtoReflect.Descriptor instead.
func (*World) Descriptor() ([]byte, []int) {
	return file_login_service_proto_rawDescGZIP(), []int{8}
}

func (x *World) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *World) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *World) GetExternalAddress() string {
	if x != nil {
		return x.ExternalAddress
	}
	return ""
}

func (x *World) GetExternalAddressProtected() string {
	if x != nil {
		return x.ExternalAddressProtected
	}
	return ""
}

func (x *World) GetExternalAddressUnprotected() string {
	if x != nil {
		return x.ExternalAddressUnprotected
	}
	return ""
}

func (x *World) GetExternalPort() uint32 {
	if x != nil {
		return x.ExternalPort
	}
	return 0
}

func (x *World) GetExternalPortProtected() uint32 {
	if x != nil {
		return x.ExternalPortProtected
	}
	return 0
}

func (x *World) GetExternalPortUnprotected() uint32 {
	if x != nil {
		return x.ExternalPortUnprotected
	}
	return 0
}

func (x *World) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

var File_login_service_proto protoreflect.FileDescriptor

var file_login_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x35, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x54, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0xaf, 0x01, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50,
	0x6c, 0x61, 0x79, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x34, 0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x7a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x3c, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x61, 0x72,
	0x61, 0x63, 0x74, 0x65, 0x72, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72,
	0x73, 0x12, 0x30, 0x0a, 0x06, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x06, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x73, 0x22, 0x8d, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x65, 0x6d, 0x69, 0x75, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x50, 0x72, 0x65, 0x6d, 0x69, 0x75, 0x6d, 0x12, 0x23,
	0x0a, 0x0d, 0x70, 0x72, 0x65, 0x6d, 0x69, 0x75, 0x6d, 0x5f, 0x75, 0x6e, 0x74, 0x69, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x70, 0x72, 0x65, 0x6d, 0x69, 0x75, 0x6d, 0x55, 0x6e,
	0x74, 0x69, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x4b, 0x65, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x22, 0x98, 0x01, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65,
	0x72, 0x12, 0x19, 0x0a, 0x08, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x43,
	0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x12, 0x3a, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x66, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72,
	0x4f, 0x75, 0x74, 0x66, 0x69, 0x74, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x66, 0x69, 0x74, 0x22, 0x86,
	0x01, 0x0a, 0x0d, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x76, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x03, 0x73, 0x65, 0x78, 0x22, 0xba, 0x01, 0x0a, 0x0f, 0x43, 0x68, 0x61, 0x72,
	0x61, 0x63, 0x74, 0x65, 0x72, 0x4f, 0x75, 0x74, 0x66, 0x69, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6c,
	0x6f, 0x6f, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x6c, 0x6f, 0x6f, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x6f, 0x6f, 0x6b,
	0x5f, 0x68, 0x65, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6c, 0x6f, 0x6f,
	0x6b, 0x48, 0x65, 0x61, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x6f, 0x6f, 0x6b, 0x5f, 0x62, 0x6f,
	0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6c, 0x6f, 0x6f, 0x6b, 0x42, 0x6f,
	0x64, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x6f, 0x6f, 0x6b, 0x5f, 0x6c, 0x65, 0x67, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6c, 0x6f, 0x6f, 0x6b, 0x4c, 0x65, 0x67, 0x73, 0x12,
	0x1b, 0x0a, 0x09, 0x6c, 0x6f, 0x6f, 0x6b, 0x5f, 0x66, 0x65, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x6c, 0x6f, 0x6f, 0x6b, 0x46, 0x65, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x64, 0x64, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x61, 0x64,
	0x64, 0x6f, 0x6e, 0x73, 0x22, 0x8b, 0x03, 0x0a, 0x05, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x3c, 0x0a,
	0x1a, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x18, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x40, 0x0a, 0x1c, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f,
	0x75, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x1a, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x55, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x23, 0x0a,
	0x0d, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x50, 0x6f,
	0x72, 0x74, 0x12, 0x36, 0x0a, 0x17, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x70,
	0x6f, 0x72, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x15, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x50, 0x6f, 0x72,
	0x74, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x3a, 0x0a, 0x19, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x17, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x50, 0x6f, 0x72, 0x74, 0x55, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x32, 0x5a, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x4a, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1f, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x13,
	0x5a, 0x11, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x64, 0x65,
	0x66, 0x73, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_login_service_proto_rawDescOnce sync.Once
	file_login_service_proto_rawDescData = file_login_service_proto_rawDesc
)

func file_login_service_proto_rawDescGZIP() []byte {
	file_login_service_proto_rawDescOnce.Do(func() {
		file_login_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_login_service_proto_rawDescData)
	})
	return file_login_service_proto_rawDescData
}

var file_login_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_login_service_proto_goTypes = []interface{}{
	(*Error)(nil),           // 0: grpc.login_server.Error
	(*LoginRequest)(nil),    // 1: grpc.login_server.LoginRequest
	(*LoginResponse)(nil),   // 2: grpc.login_server.LoginResponse
	(*PlayData)(nil),        // 3: grpc.login_server.PlayData
	(*Session)(nil),         // 4: grpc.login_server.Session
	(*Character)(nil),       // 5: grpc.login_server.Character
	(*CharacterInfo)(nil),   // 6: grpc.login_server.CharacterInfo
	(*CharacterOutfit)(nil), // 7: grpc.login_server.CharacterOutfit
	(*World)(nil),           // 8: grpc.login_server.World
}
var file_login_service_proto_depIdxs = []int32{
	3, // 0: grpc.login_server.LoginResponse.play_data:type_name -> grpc.login_server.PlayData
	4, // 1: grpc.login_server.LoginResponse.session:type_name -> grpc.login_server.Session
	0, // 2: grpc.login_server.LoginResponse.error:type_name -> grpc.login_server.Error
	5, // 3: grpc.login_server.PlayData.characters:type_name -> grpc.login_server.Character
	8, // 4: grpc.login_server.PlayData.worlds:type_name -> grpc.login_server.World
	6, // 5: grpc.login_server.Character.info:type_name -> grpc.login_server.CharacterInfo
	7, // 6: grpc.login_server.Character.outfit:type_name -> grpc.login_server.CharacterOutfit
	1, // 7: grpc.login_server.LoginService.Login:input_type -> grpc.login_server.LoginRequest
	2, // 8: grpc.login_server.LoginService.Login:output_type -> grpc.login_server.LoginResponse
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_login_service_proto_init() }
func file_login_service_proto_init() {
	if File_login_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_login_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_login_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_login_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
		file_login_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayData); i {
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
		file_login_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Session); i {
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
		file_login_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Character); i {
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
		file_login_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CharacterInfo); i {
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
		file_login_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CharacterOutfit); i {
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
		file_login_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*World); i {
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
			RawDescriptor: file_login_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_login_service_proto_goTypes,
		DependencyIndexes: file_login_service_proto_depIdxs,
		MessageInfos:      file_login_service_proto_msgTypes,
	}.Build()
	File_login_service_proto = out.File
	file_login_service_proto_rawDesc = nil
	file_login_service_proto_goTypes = nil
	file_login_service_proto_depIdxs = nil
}
