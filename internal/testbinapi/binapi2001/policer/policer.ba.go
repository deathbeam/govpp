// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.5.0-dev
//  VPP:              20.01
// source: .vppapi/core/policer.api.json

// Package policer contains generated bindings for API file policer.api.
//
// Contents:
//   4 messages
//
package policer

import (
	api "git.fd.io/govpp.git/api"
	codec "git.fd.io/govpp.git/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "policer"
	APIVersion = "1.0.0"
	VersionCrc = 0x4218ed4f
)

// PolicerAddDel defines message 'policer_add_del'.
type PolicerAddDel struct {
	IsAdd             uint8  `binapi:"u8,name=is_add" json:"is_add,omitempty"`
	Name              []byte `binapi:"u8[64],name=name" json:"name,omitempty"`
	Cir               uint32 `binapi:"u32,name=cir" json:"cir,omitempty"`
	Eir               uint32 `binapi:"u32,name=eir" json:"eir,omitempty"`
	Cb                uint64 `binapi:"u64,name=cb" json:"cb,omitempty"`
	Eb                uint64 `binapi:"u64,name=eb" json:"eb,omitempty"`
	RateType          uint8  `binapi:"u8,name=rate_type" json:"rate_type,omitempty"`
	RoundType         uint8  `binapi:"u8,name=round_type" json:"round_type,omitempty"`
	Type              uint8  `binapi:"u8,name=type" json:"type,omitempty"`
	ColorAware        uint8  `binapi:"u8,name=color_aware" json:"color_aware,omitempty"`
	ConformActionType uint8  `binapi:"u8,name=conform_action_type" json:"conform_action_type,omitempty"`
	ConformDscp       uint8  `binapi:"u8,name=conform_dscp" json:"conform_dscp,omitempty"`
	ExceedActionType  uint8  `binapi:"u8,name=exceed_action_type" json:"exceed_action_type,omitempty"`
	ExceedDscp        uint8  `binapi:"u8,name=exceed_dscp" json:"exceed_dscp,omitempty"`
	ViolateActionType uint8  `binapi:"u8,name=violate_action_type" json:"violate_action_type,omitempty"`
	ViolateDscp       uint8  `binapi:"u8,name=violate_dscp" json:"violate_dscp,omitempty"`
}

func (m *PolicerAddDel) Reset()               { *m = PolicerAddDel{} }
func (*PolicerAddDel) GetMessageName() string { return "policer_add_del" }
func (*PolicerAddDel) GetCrcString() string   { return "dfea2be8" }
func (*PolicerAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *PolicerAddDel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1      // m.IsAdd
	size += 1 * 64 // m.Name
	size += 4      // m.Cir
	size += 4      // m.Eir
	size += 8      // m.Cb
	size += 8      // m.Eb
	size += 1      // m.RateType
	size += 1      // m.RoundType
	size += 1      // m.Type
	size += 1      // m.ColorAware
	size += 1      // m.ConformActionType
	size += 1      // m.ConformDscp
	size += 1      // m.ExceedActionType
	size += 1      // m.ExceedDscp
	size += 1      // m.ViolateActionType
	size += 1      // m.ViolateDscp
	return size
}
func (m *PolicerAddDel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(m.IsAdd)
	buf.EncodeBytes(m.Name, 64)
	buf.EncodeUint32(m.Cir)
	buf.EncodeUint32(m.Eir)
	buf.EncodeUint64(m.Cb)
	buf.EncodeUint64(m.Eb)
	buf.EncodeUint8(m.RateType)
	buf.EncodeUint8(m.RoundType)
	buf.EncodeUint8(m.Type)
	buf.EncodeUint8(m.ColorAware)
	buf.EncodeUint8(m.ConformActionType)
	buf.EncodeUint8(m.ConformDscp)
	buf.EncodeUint8(m.ExceedActionType)
	buf.EncodeUint8(m.ExceedDscp)
	buf.EncodeUint8(m.ViolateActionType)
	buf.EncodeUint8(m.ViolateDscp)
	return buf.Bytes(), nil
}
func (m *PolicerAddDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsAdd = buf.DecodeUint8()
	m.Name = make([]byte, 64)
	copy(m.Name, buf.DecodeBytes(len(m.Name)))
	m.Cir = buf.DecodeUint32()
	m.Eir = buf.DecodeUint32()
	m.Cb = buf.DecodeUint64()
	m.Eb = buf.DecodeUint64()
	m.RateType = buf.DecodeUint8()
	m.RoundType = buf.DecodeUint8()
	m.Type = buf.DecodeUint8()
	m.ColorAware = buf.DecodeUint8()
	m.ConformActionType = buf.DecodeUint8()
	m.ConformDscp = buf.DecodeUint8()
	m.ExceedActionType = buf.DecodeUint8()
	m.ExceedDscp = buf.DecodeUint8()
	m.ViolateActionType = buf.DecodeUint8()
	m.ViolateDscp = buf.DecodeUint8()
	return nil
}

// PolicerAddDelReply defines message 'policer_add_del_reply'.
type PolicerAddDelReply struct {
	Retval       int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	PolicerIndex uint32 `binapi:"u32,name=policer_index" json:"policer_index,omitempty"`
}

func (m *PolicerAddDelReply) Reset()               { *m = PolicerAddDelReply{} }
func (*PolicerAddDelReply) GetMessageName() string { return "policer_add_del_reply" }
func (*PolicerAddDelReply) GetCrcString() string   { return "a177cef2" }
func (*PolicerAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *PolicerAddDelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.PolicerIndex
	return size
}
func (m *PolicerAddDelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(m.PolicerIndex)
	return buf.Bytes(), nil
}
func (m *PolicerAddDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.PolicerIndex = buf.DecodeUint32()
	return nil
}

// PolicerDetails defines message 'policer_details'.
type PolicerDetails struct {
	Name               []byte `binapi:"u8[64],name=name" json:"name,omitempty"`
	Cir                uint32 `binapi:"u32,name=cir" json:"cir,omitempty"`
	Eir                uint32 `binapi:"u32,name=eir" json:"eir,omitempty"`
	Cb                 uint64 `binapi:"u64,name=cb" json:"cb,omitempty"`
	Eb                 uint64 `binapi:"u64,name=eb" json:"eb,omitempty"`
	RateType           uint8  `binapi:"u8,name=rate_type" json:"rate_type,omitempty"`
	RoundType          uint8  `binapi:"u8,name=round_type" json:"round_type,omitempty"`
	Type               uint8  `binapi:"u8,name=type" json:"type,omitempty"`
	ConformActionType  uint8  `binapi:"u8,name=conform_action_type" json:"conform_action_type,omitempty"`
	ConformDscp        uint8  `binapi:"u8,name=conform_dscp" json:"conform_dscp,omitempty"`
	ExceedActionType   uint8  `binapi:"u8,name=exceed_action_type" json:"exceed_action_type,omitempty"`
	ExceedDscp         uint8  `binapi:"u8,name=exceed_dscp" json:"exceed_dscp,omitempty"`
	ViolateActionType  uint8  `binapi:"u8,name=violate_action_type" json:"violate_action_type,omitempty"`
	ViolateDscp        uint8  `binapi:"u8,name=violate_dscp" json:"violate_dscp,omitempty"`
	SingleRate         uint8  `binapi:"u8,name=single_rate" json:"single_rate,omitempty"`
	ColorAware         uint8  `binapi:"u8,name=color_aware" json:"color_aware,omitempty"`
	Scale              uint32 `binapi:"u32,name=scale" json:"scale,omitempty"`
	CirTokensPerPeriod uint32 `binapi:"u32,name=cir_tokens_per_period" json:"cir_tokens_per_period,omitempty"`
	PirTokensPerPeriod uint32 `binapi:"u32,name=pir_tokens_per_period" json:"pir_tokens_per_period,omitempty"`
	CurrentLimit       uint32 `binapi:"u32,name=current_limit" json:"current_limit,omitempty"`
	CurrentBucket      uint32 `binapi:"u32,name=current_bucket" json:"current_bucket,omitempty"`
	ExtendedLimit      uint32 `binapi:"u32,name=extended_limit" json:"extended_limit,omitempty"`
	ExtendedBucket     uint32 `binapi:"u32,name=extended_bucket" json:"extended_bucket,omitempty"`
	LastUpdateTime     uint64 `binapi:"u64,name=last_update_time" json:"last_update_time,omitempty"`
}

func (m *PolicerDetails) Reset()               { *m = PolicerDetails{} }
func (*PolicerDetails) GetMessageName() string { return "policer_details" }
func (*PolicerDetails) GetCrcString() string   { return "ff2765f0" }
func (*PolicerDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *PolicerDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 * 64 // m.Name
	size += 4      // m.Cir
	size += 4      // m.Eir
	size += 8      // m.Cb
	size += 8      // m.Eb
	size += 1      // m.RateType
	size += 1      // m.RoundType
	size += 1      // m.Type
	size += 1      // m.ConformActionType
	size += 1      // m.ConformDscp
	size += 1      // m.ExceedActionType
	size += 1      // m.ExceedDscp
	size += 1      // m.ViolateActionType
	size += 1      // m.ViolateDscp
	size += 1      // m.SingleRate
	size += 1      // m.ColorAware
	size += 4      // m.Scale
	size += 4      // m.CirTokensPerPeriod
	size += 4      // m.PirTokensPerPeriod
	size += 4      // m.CurrentLimit
	size += 4      // m.CurrentBucket
	size += 4      // m.ExtendedLimit
	size += 4      // m.ExtendedBucket
	size += 8      // m.LastUpdateTime
	return size
}
func (m *PolicerDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBytes(m.Name, 64)
	buf.EncodeUint32(m.Cir)
	buf.EncodeUint32(m.Eir)
	buf.EncodeUint64(m.Cb)
	buf.EncodeUint64(m.Eb)
	buf.EncodeUint8(m.RateType)
	buf.EncodeUint8(m.RoundType)
	buf.EncodeUint8(m.Type)
	buf.EncodeUint8(m.ConformActionType)
	buf.EncodeUint8(m.ConformDscp)
	buf.EncodeUint8(m.ExceedActionType)
	buf.EncodeUint8(m.ExceedDscp)
	buf.EncodeUint8(m.ViolateActionType)
	buf.EncodeUint8(m.ViolateDscp)
	buf.EncodeUint8(m.SingleRate)
	buf.EncodeUint8(m.ColorAware)
	buf.EncodeUint32(m.Scale)
	buf.EncodeUint32(m.CirTokensPerPeriod)
	buf.EncodeUint32(m.PirTokensPerPeriod)
	buf.EncodeUint32(m.CurrentLimit)
	buf.EncodeUint32(m.CurrentBucket)
	buf.EncodeUint32(m.ExtendedLimit)
	buf.EncodeUint32(m.ExtendedBucket)
	buf.EncodeUint64(m.LastUpdateTime)
	return buf.Bytes(), nil
}
func (m *PolicerDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Name = make([]byte, 64)
	copy(m.Name, buf.DecodeBytes(len(m.Name)))
	m.Cir = buf.DecodeUint32()
	m.Eir = buf.DecodeUint32()
	m.Cb = buf.DecodeUint64()
	m.Eb = buf.DecodeUint64()
	m.RateType = buf.DecodeUint8()
	m.RoundType = buf.DecodeUint8()
	m.Type = buf.DecodeUint8()
	m.ConformActionType = buf.DecodeUint8()
	m.ConformDscp = buf.DecodeUint8()
	m.ExceedActionType = buf.DecodeUint8()
	m.ExceedDscp = buf.DecodeUint8()
	m.ViolateActionType = buf.DecodeUint8()
	m.ViolateDscp = buf.DecodeUint8()
	m.SingleRate = buf.DecodeUint8()
	m.ColorAware = buf.DecodeUint8()
	m.Scale = buf.DecodeUint32()
	m.CirTokensPerPeriod = buf.DecodeUint32()
	m.PirTokensPerPeriod = buf.DecodeUint32()
	m.CurrentLimit = buf.DecodeUint32()
	m.CurrentBucket = buf.DecodeUint32()
	m.ExtendedLimit = buf.DecodeUint32()
	m.ExtendedBucket = buf.DecodeUint32()
	m.LastUpdateTime = buf.DecodeUint64()
	return nil
}

// PolicerDump defines message 'policer_dump'.
type PolicerDump struct {
	MatchNameValid uint8  `binapi:"u8,name=match_name_valid" json:"match_name_valid,omitempty"`
	MatchName      []byte `binapi:"u8[64],name=match_name" json:"match_name,omitempty"`
}

func (m *PolicerDump) Reset()               { *m = PolicerDump{} }
func (*PolicerDump) GetMessageName() string { return "policer_dump" }
func (*PolicerDump) GetCrcString() string   { return "8be04d34" }
func (*PolicerDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *PolicerDump) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1      // m.MatchNameValid
	size += 1 * 64 // m.MatchName
	return size
}
func (m *PolicerDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(m.MatchNameValid)
	buf.EncodeBytes(m.MatchName, 64)
	return buf.Bytes(), nil
}
func (m *PolicerDump) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.MatchNameValid = buf.DecodeUint8()
	m.MatchName = make([]byte, 64)
	copy(m.MatchName, buf.DecodeBytes(len(m.MatchName)))
	return nil
}

func init() { file_policer_binapi_init() }
func file_policer_binapi_init() {
	api.RegisterMessage((*PolicerAddDel)(nil), "policer_add_del_dfea2be8")
	api.RegisterMessage((*PolicerAddDelReply)(nil), "policer_add_del_reply_a177cef2")
	api.RegisterMessage((*PolicerDetails)(nil), "policer_details_ff2765f0")
	api.RegisterMessage((*PolicerDump)(nil), "policer_dump_8be04d34")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*PolicerAddDel)(nil),
		(*PolicerAddDelReply)(nil),
		(*PolicerDetails)(nil),
		(*PolicerDump)(nil),
	}
}
