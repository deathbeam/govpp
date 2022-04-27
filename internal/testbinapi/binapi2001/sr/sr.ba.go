// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.5.0-dev
//  VPP:              20.01
// source: .vppapi/core/sr.api.json

// Package sr contains generated bindings for API file sr.api.
//
// Contents:
//   3 structs
//  20 messages
//
package sr

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
	APIFile    = "sr"
	APIVersion = "1.2.0"
	VersionCrc = 0xbf277f96
)

// SrIP6Address defines type 'sr_ip6_address'.
type SrIP6Address struct {
	Data []byte `binapi:"u8[16],name=data" json:"data,omitempty"`
}

// Srv6Sid defines type 'srv6_sid'.
type Srv6Sid struct {
	Addr []byte `binapi:"u8[16],name=addr" json:"addr,omitempty"`
}

// Srv6SidList defines type 'srv6_sid_list'.
type Srv6SidList struct {
	NumSids uint8       `binapi:"u8,name=num_sids" json:"num_sids,omitempty"`
	Weight  uint32      `binapi:"u32,name=weight" json:"weight,omitempty"`
	Sids    [16]Srv6Sid `binapi:"srv6_sid[16],name=sids" json:"sids,omitempty"`
}

// SrLocalsidAddDel defines message 'sr_localsid_add_del'.
type SrLocalsidAddDel struct {
	IsDel     uint8   `binapi:"u8,name=is_del" json:"is_del,omitempty"`
	Localsid  Srv6Sid `binapi:"srv6_sid,name=localsid" json:"localsid,omitempty"`
	EndPsp    uint8   `binapi:"u8,name=end_psp" json:"end_psp,omitempty"`
	Behavior  uint8   `binapi:"u8,name=behavior" json:"behavior,omitempty"`
	SwIfIndex uint32  `binapi:"u32,name=sw_if_index" json:"sw_if_index,omitempty"`
	VlanIndex uint32  `binapi:"u32,name=vlan_index" json:"vlan_index,omitempty"`
	FibTable  uint32  `binapi:"u32,name=fib_table" json:"fib_table,omitempty"`
	NhAddr6   []byte  `binapi:"u8[16],name=nh_addr6" json:"nh_addr6,omitempty"`
	NhAddr4   []byte  `binapi:"u8[4],name=nh_addr4" json:"nh_addr4,omitempty"`
}

func (m *SrLocalsidAddDel) Reset()               { *m = SrLocalsidAddDel{} }
func (*SrLocalsidAddDel) GetMessageName() string { return "sr_localsid_add_del" }
func (*SrLocalsidAddDel) GetCrcString() string   { return "b30489eb" }
func (*SrLocalsidAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SrLocalsidAddDel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1      // m.IsDel
	size += 1 * 16 // m.Localsid.Addr
	size += 1      // m.EndPsp
	size += 1      // m.Behavior
	size += 4      // m.SwIfIndex
	size += 4      // m.VlanIndex
	size += 4      // m.FibTable
	size += 1 * 16 // m.NhAddr6
	size += 1 * 4  // m.NhAddr4
	return size
}
func (m *SrLocalsidAddDel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(m.IsDel)
	buf.EncodeBytes(m.Localsid.Addr, 16)
	buf.EncodeUint8(m.EndPsp)
	buf.EncodeUint8(m.Behavior)
	buf.EncodeUint32(m.SwIfIndex)
	buf.EncodeUint32(m.VlanIndex)
	buf.EncodeUint32(m.FibTable)
	buf.EncodeBytes(m.NhAddr6, 16)
	buf.EncodeBytes(m.NhAddr4, 4)
	return buf.Bytes(), nil
}
func (m *SrLocalsidAddDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsDel = buf.DecodeUint8()
	m.Localsid.Addr = make([]byte, 16)
	copy(m.Localsid.Addr, buf.DecodeBytes(len(m.Localsid.Addr)))
	m.EndPsp = buf.DecodeUint8()
	m.Behavior = buf.DecodeUint8()
	m.SwIfIndex = buf.DecodeUint32()
	m.VlanIndex = buf.DecodeUint32()
	m.FibTable = buf.DecodeUint32()
	m.NhAddr6 = make([]byte, 16)
	copy(m.NhAddr6, buf.DecodeBytes(len(m.NhAddr6)))
	m.NhAddr4 = make([]byte, 4)
	copy(m.NhAddr4, buf.DecodeBytes(len(m.NhAddr4)))
	return nil
}

// SrLocalsidAddDelReply defines message 'sr_localsid_add_del_reply'.
type SrLocalsidAddDelReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *SrLocalsidAddDelReply) Reset()               { *m = SrLocalsidAddDelReply{} }
func (*SrLocalsidAddDelReply) GetMessageName() string { return "sr_localsid_add_del_reply" }
func (*SrLocalsidAddDelReply) GetCrcString() string   { return "e8d4e804" }
func (*SrLocalsidAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SrLocalsidAddDelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *SrLocalsidAddDelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *SrLocalsidAddDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// SrLocalsidsDetails defines message 'sr_localsids_details'.
type SrLocalsidsDetails struct {
	Addr                    Srv6Sid `binapi:"srv6_sid,name=addr" json:"addr,omitempty"`
	EndPsp                  uint8   `binapi:"u8,name=end_psp" json:"end_psp,omitempty"`
	Behavior                uint16  `binapi:"u16,name=behavior" json:"behavior,omitempty"`
	FibTable                uint32  `binapi:"u32,name=fib_table" json:"fib_table,omitempty"`
	VlanIndex               uint32  `binapi:"u32,name=vlan_index" json:"vlan_index,omitempty"`
	XconnectNhAddr6         []byte  `binapi:"u8[16],name=xconnect_nh_addr6" json:"xconnect_nh_addr6,omitempty"`
	XconnectNhAddr4         []byte  `binapi:"u8[4],name=xconnect_nh_addr4" json:"xconnect_nh_addr4,omitempty"`
	XconnectIfaceOrVrfTable uint32  `binapi:"u32,name=xconnect_iface_or_vrf_table" json:"xconnect_iface_or_vrf_table,omitempty"`
}

func (m *SrLocalsidsDetails) Reset()               { *m = SrLocalsidsDetails{} }
func (*SrLocalsidsDetails) GetMessageName() string { return "sr_localsids_details" }
func (*SrLocalsidsDetails) GetCrcString() string   { return "0791babc" }
func (*SrLocalsidsDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SrLocalsidsDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 * 16 // m.Addr.Addr
	size += 1      // m.EndPsp
	size += 2      // m.Behavior
	size += 4      // m.FibTable
	size += 4      // m.VlanIndex
	size += 1 * 16 // m.XconnectNhAddr6
	size += 1 * 4  // m.XconnectNhAddr4
	size += 4      // m.XconnectIfaceOrVrfTable
	return size
}
func (m *SrLocalsidsDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBytes(m.Addr.Addr, 16)
	buf.EncodeUint8(m.EndPsp)
	buf.EncodeUint16(m.Behavior)
	buf.EncodeUint32(m.FibTable)
	buf.EncodeUint32(m.VlanIndex)
	buf.EncodeBytes(m.XconnectNhAddr6, 16)
	buf.EncodeBytes(m.XconnectNhAddr4, 4)
	buf.EncodeUint32(m.XconnectIfaceOrVrfTable)
	return buf.Bytes(), nil
}
func (m *SrLocalsidsDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Addr.Addr = make([]byte, 16)
	copy(m.Addr.Addr, buf.DecodeBytes(len(m.Addr.Addr)))
	m.EndPsp = buf.DecodeUint8()
	m.Behavior = buf.DecodeUint16()
	m.FibTable = buf.DecodeUint32()
	m.VlanIndex = buf.DecodeUint32()
	m.XconnectNhAddr6 = make([]byte, 16)
	copy(m.XconnectNhAddr6, buf.DecodeBytes(len(m.XconnectNhAddr6)))
	m.XconnectNhAddr4 = make([]byte, 4)
	copy(m.XconnectNhAddr4, buf.DecodeBytes(len(m.XconnectNhAddr4)))
	m.XconnectIfaceOrVrfTable = buf.DecodeUint32()
	return nil
}

// SrLocalsidsDump defines message 'sr_localsids_dump'.
type SrLocalsidsDump struct{}

func (m *SrLocalsidsDump) Reset()               { *m = SrLocalsidsDump{} }
func (*SrLocalsidsDump) GetMessageName() string { return "sr_localsids_dump" }
func (*SrLocalsidsDump) GetCrcString() string   { return "51077d14" }
func (*SrLocalsidsDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SrLocalsidsDump) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *SrLocalsidsDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *SrLocalsidsDump) Unmarshal(b []byte) error {
	return nil
}

// SrPoliciesDetails defines message 'sr_policies_details'.
type SrPoliciesDetails struct {
	Bsid        Srv6Sid       `binapi:"srv6_sid,name=bsid" json:"bsid,omitempty"`
	Type        uint8         `binapi:"u8,name=type" json:"type,omitempty"`
	IsEncap     uint8         `binapi:"u8,name=is_encap" json:"is_encap,omitempty"`
	FibTable    uint32        `binapi:"u32,name=fib_table" json:"fib_table,omitempty"`
	NumSidLists uint8         `binapi:"u8,name=num_sid_lists" json:"-"`
	SidLists    []Srv6SidList `binapi:"srv6_sid_list[num_sid_lists],name=sid_lists" json:"sid_lists,omitempty"`
}

func (m *SrPoliciesDetails) Reset()               { *m = SrPoliciesDetails{} }
func (*SrPoliciesDetails) GetMessageName() string { return "sr_policies_details" }
func (*SrPoliciesDetails) GetCrcString() string   { return "5087f460" }
func (*SrPoliciesDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SrPoliciesDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 * 16 // m.Bsid.Addr
	size += 1      // m.Type
	size += 1      // m.IsEncap
	size += 4      // m.FibTable
	size += 1      // m.NumSidLists
	for j1 := 0; j1 < len(m.SidLists); j1++ {
		var s1 Srv6SidList
		_ = s1
		if j1 < len(m.SidLists) {
			s1 = m.SidLists[j1]
		}
		size += 1 // s1.NumSids
		size += 4 // s1.Weight
		for j2 := 0; j2 < 16; j2++ {
			size += 1 * 16 // s1.Sids[j2].Addr
		}
	}
	return size
}
func (m *SrPoliciesDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBytes(m.Bsid.Addr, 16)
	buf.EncodeUint8(m.Type)
	buf.EncodeUint8(m.IsEncap)
	buf.EncodeUint32(m.FibTable)
	buf.EncodeUint8(uint8(len(m.SidLists)))
	for j0 := 0; j0 < len(m.SidLists); j0++ {
		var v0 Srv6SidList // SidLists
		if j0 < len(m.SidLists) {
			v0 = m.SidLists[j0]
		}
		buf.EncodeUint8(v0.NumSids)
		buf.EncodeUint32(v0.Weight)
		for j1 := 0; j1 < 16; j1++ {
			buf.EncodeBytes(v0.Sids[j1].Addr, 16)
		}
	}
	return buf.Bytes(), nil
}
func (m *SrPoliciesDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Bsid.Addr = make([]byte, 16)
	copy(m.Bsid.Addr, buf.DecodeBytes(len(m.Bsid.Addr)))
	m.Type = buf.DecodeUint8()
	m.IsEncap = buf.DecodeUint8()
	m.FibTable = buf.DecodeUint32()
	m.NumSidLists = buf.DecodeUint8()
	m.SidLists = make([]Srv6SidList, m.NumSidLists)
	for j0 := 0; j0 < len(m.SidLists); j0++ {
		m.SidLists[j0].NumSids = buf.DecodeUint8()
		m.SidLists[j0].Weight = buf.DecodeUint32()
		for j1 := 0; j1 < 16; j1++ {
			m.SidLists[j0].Sids[j1].Addr = make([]byte, 16)
			copy(m.SidLists[j0].Sids[j1].Addr, buf.DecodeBytes(len(m.SidLists[j0].Sids[j1].Addr)))
		}
	}
	return nil
}

// SrPoliciesDump defines message 'sr_policies_dump'.
type SrPoliciesDump struct{}

func (m *SrPoliciesDump) Reset()               { *m = SrPoliciesDump{} }
func (*SrPoliciesDump) GetMessageName() string { return "sr_policies_dump" }
func (*SrPoliciesDump) GetCrcString() string   { return "51077d14" }
func (*SrPoliciesDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SrPoliciesDump) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *SrPoliciesDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *SrPoliciesDump) Unmarshal(b []byte) error {
	return nil
}

// SrPolicyAdd defines message 'sr_policy_add'.
type SrPolicyAdd struct {
	BsidAddr []byte      `binapi:"u8[16],name=bsid_addr" json:"bsid_addr,omitempty"`
	Weight   uint32      `binapi:"u32,name=weight" json:"weight,omitempty"`
	IsEncap  uint8       `binapi:"u8,name=is_encap" json:"is_encap,omitempty"`
	Type     uint8       `binapi:"u8,name=type" json:"type,omitempty"`
	FibTable uint32      `binapi:"u32,name=fib_table" json:"fib_table,omitempty"`
	Sids     Srv6SidList `binapi:"srv6_sid_list,name=sids" json:"sids,omitempty"`
}

func (m *SrPolicyAdd) Reset()               { *m = SrPolicyAdd{} }
func (*SrPolicyAdd) GetMessageName() string { return "sr_policy_add" }
func (*SrPolicyAdd) GetCrcString() string   { return "4b6e2484" }
func (*SrPolicyAdd) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SrPolicyAdd) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 * 16 // m.BsidAddr
	size += 4      // m.Weight
	size += 1      // m.IsEncap
	size += 1      // m.Type
	size += 4      // m.FibTable
	size += 1      // m.Sids.NumSids
	size += 4      // m.Sids.Weight
	for j2 := 0; j2 < 16; j2++ {
		size += 1 * 16 // m.Sids.Sids[j2].Addr
	}
	return size
}
func (m *SrPolicyAdd) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBytes(m.BsidAddr, 16)
	buf.EncodeUint32(m.Weight)
	buf.EncodeUint8(m.IsEncap)
	buf.EncodeUint8(m.Type)
	buf.EncodeUint32(m.FibTable)
	buf.EncodeUint8(m.Sids.NumSids)
	buf.EncodeUint32(m.Sids.Weight)
	for j1 := 0; j1 < 16; j1++ {
		buf.EncodeBytes(m.Sids.Sids[j1].Addr, 16)
	}
	return buf.Bytes(), nil
}
func (m *SrPolicyAdd) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.BsidAddr = make([]byte, 16)
	copy(m.BsidAddr, buf.DecodeBytes(len(m.BsidAddr)))
	m.Weight = buf.DecodeUint32()
	m.IsEncap = buf.DecodeUint8()
	m.Type = buf.DecodeUint8()
	m.FibTable = buf.DecodeUint32()
	m.Sids.NumSids = buf.DecodeUint8()
	m.Sids.Weight = buf.DecodeUint32()
	for j1 := 0; j1 < 16; j1++ {
		m.Sids.Sids[j1].Addr = make([]byte, 16)
		copy(m.Sids.Sids[j1].Addr, buf.DecodeBytes(len(m.Sids.Sids[j1].Addr)))
	}
	return nil
}

// SrPolicyAddReply defines message 'sr_policy_add_reply'.
type SrPolicyAddReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *SrPolicyAddReply) Reset()               { *m = SrPolicyAddReply{} }
func (*SrPolicyAddReply) GetMessageName() string { return "sr_policy_add_reply" }
func (*SrPolicyAddReply) GetCrcString() string   { return "e8d4e804" }
func (*SrPolicyAddReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SrPolicyAddReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *SrPolicyAddReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *SrPolicyAddReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// SrPolicyDel defines message 'sr_policy_del'.
type SrPolicyDel struct {
	BsidAddr      Srv6Sid `binapi:"srv6_sid,name=bsid_addr" json:"bsid_addr,omitempty"`
	SrPolicyIndex uint32  `binapi:"u32,name=sr_policy_index" json:"sr_policy_index,omitempty"`
}

func (m *SrPolicyDel) Reset()               { *m = SrPolicyDel{} }
func (*SrPolicyDel) GetMessageName() string { return "sr_policy_del" }
func (*SrPolicyDel) GetCrcString() string   { return "e4133171" }
func (*SrPolicyDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SrPolicyDel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 * 16 // m.BsidAddr.Addr
	size += 4      // m.SrPolicyIndex
	return size
}
func (m *SrPolicyDel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBytes(m.BsidAddr.Addr, 16)
	buf.EncodeUint32(m.SrPolicyIndex)
	return buf.Bytes(), nil
}
func (m *SrPolicyDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.BsidAddr.Addr = make([]byte, 16)
	copy(m.BsidAddr.Addr, buf.DecodeBytes(len(m.BsidAddr.Addr)))
	m.SrPolicyIndex = buf.DecodeUint32()
	return nil
}

// SrPolicyDelReply defines message 'sr_policy_del_reply'.
type SrPolicyDelReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *SrPolicyDelReply) Reset()               { *m = SrPolicyDelReply{} }
func (*SrPolicyDelReply) GetMessageName() string { return "sr_policy_del_reply" }
func (*SrPolicyDelReply) GetCrcString() string   { return "e8d4e804" }
func (*SrPolicyDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SrPolicyDelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *SrPolicyDelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *SrPolicyDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// SrPolicyMod defines message 'sr_policy_mod'.
type SrPolicyMod struct {
	BsidAddr      []byte      `binapi:"u8[16],name=bsid_addr" json:"bsid_addr,omitempty"`
	SrPolicyIndex uint32      `binapi:"u32,name=sr_policy_index" json:"sr_policy_index,omitempty"`
	FibTable      uint32      `binapi:"u32,name=fib_table" json:"fib_table,omitempty"`
	Operation     uint8       `binapi:"u8,name=operation" json:"operation,omitempty"`
	SlIndex       uint32      `binapi:"u32,name=sl_index" json:"sl_index,omitempty"`
	Weight        uint32      `binapi:"u32,name=weight" json:"weight,omitempty"`
	Sids          Srv6SidList `binapi:"srv6_sid_list,name=sids" json:"sids,omitempty"`
}

func (m *SrPolicyMod) Reset()               { *m = SrPolicyMod{} }
func (*SrPolicyMod) GetMessageName() string { return "sr_policy_mod" }
func (*SrPolicyMod) GetCrcString() string   { return "c1dfaee0" }
func (*SrPolicyMod) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SrPolicyMod) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 * 16 // m.BsidAddr
	size += 4      // m.SrPolicyIndex
	size += 4      // m.FibTable
	size += 1      // m.Operation
	size += 4      // m.SlIndex
	size += 4      // m.Weight
	size += 1      // m.Sids.NumSids
	size += 4      // m.Sids.Weight
	for j2 := 0; j2 < 16; j2++ {
		size += 1 * 16 // m.Sids.Sids[j2].Addr
	}
	return size
}
func (m *SrPolicyMod) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBytes(m.BsidAddr, 16)
	buf.EncodeUint32(m.SrPolicyIndex)
	buf.EncodeUint32(m.FibTable)
	buf.EncodeUint8(m.Operation)
	buf.EncodeUint32(m.SlIndex)
	buf.EncodeUint32(m.Weight)
	buf.EncodeUint8(m.Sids.NumSids)
	buf.EncodeUint32(m.Sids.Weight)
	for j1 := 0; j1 < 16; j1++ {
		buf.EncodeBytes(m.Sids.Sids[j1].Addr, 16)
	}
	return buf.Bytes(), nil
}
func (m *SrPolicyMod) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.BsidAddr = make([]byte, 16)
	copy(m.BsidAddr, buf.DecodeBytes(len(m.BsidAddr)))
	m.SrPolicyIndex = buf.DecodeUint32()
	m.FibTable = buf.DecodeUint32()
	m.Operation = buf.DecodeUint8()
	m.SlIndex = buf.DecodeUint32()
	m.Weight = buf.DecodeUint32()
	m.Sids.NumSids = buf.DecodeUint8()
	m.Sids.Weight = buf.DecodeUint32()
	for j1 := 0; j1 < 16; j1++ {
		m.Sids.Sids[j1].Addr = make([]byte, 16)
		copy(m.Sids.Sids[j1].Addr, buf.DecodeBytes(len(m.Sids.Sids[j1].Addr)))
	}
	return nil
}

// SrPolicyModReply defines message 'sr_policy_mod_reply'.
type SrPolicyModReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *SrPolicyModReply) Reset()               { *m = SrPolicyModReply{} }
func (*SrPolicyModReply) GetMessageName() string { return "sr_policy_mod_reply" }
func (*SrPolicyModReply) GetCrcString() string   { return "e8d4e804" }
func (*SrPolicyModReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SrPolicyModReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *SrPolicyModReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *SrPolicyModReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// SrSetEncapHopLimit defines message 'sr_set_encap_hop_limit'.
type SrSetEncapHopLimit struct {
	HopLimit uint8 `binapi:"u8,name=hop_limit" json:"hop_limit,omitempty"`
}

func (m *SrSetEncapHopLimit) Reset()               { *m = SrSetEncapHopLimit{} }
func (*SrSetEncapHopLimit) GetMessageName() string { return "sr_set_encap_hop_limit" }
func (*SrSetEncapHopLimit) GetCrcString() string   { return "aa75d7d0" }
func (*SrSetEncapHopLimit) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SrSetEncapHopLimit) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.HopLimit
	return size
}
func (m *SrSetEncapHopLimit) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(m.HopLimit)
	return buf.Bytes(), nil
}
func (m *SrSetEncapHopLimit) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.HopLimit = buf.DecodeUint8()
	return nil
}

// SrSetEncapHopLimitReply defines message 'sr_set_encap_hop_limit_reply'.
type SrSetEncapHopLimitReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *SrSetEncapHopLimitReply) Reset()               { *m = SrSetEncapHopLimitReply{} }
func (*SrSetEncapHopLimitReply) GetMessageName() string { return "sr_set_encap_hop_limit_reply" }
func (*SrSetEncapHopLimitReply) GetCrcString() string   { return "e8d4e804" }
func (*SrSetEncapHopLimitReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SrSetEncapHopLimitReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *SrSetEncapHopLimitReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *SrSetEncapHopLimitReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// SrSetEncapSource defines message 'sr_set_encap_source'.
type SrSetEncapSource struct {
	EncapsSource []byte `binapi:"u8[16],name=encaps_source" json:"encaps_source,omitempty"`
}

func (m *SrSetEncapSource) Reset()               { *m = SrSetEncapSource{} }
func (*SrSetEncapSource) GetMessageName() string { return "sr_set_encap_source" }
func (*SrSetEncapSource) GetCrcString() string   { return "d05bb4de" }
func (*SrSetEncapSource) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SrSetEncapSource) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 * 16 // m.EncapsSource
	return size
}
func (m *SrSetEncapSource) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBytes(m.EncapsSource, 16)
	return buf.Bytes(), nil
}
func (m *SrSetEncapSource) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.EncapsSource = make([]byte, 16)
	copy(m.EncapsSource, buf.DecodeBytes(len(m.EncapsSource)))
	return nil
}

// SrSetEncapSourceReply defines message 'sr_set_encap_source_reply'.
type SrSetEncapSourceReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *SrSetEncapSourceReply) Reset()               { *m = SrSetEncapSourceReply{} }
func (*SrSetEncapSourceReply) GetMessageName() string { return "sr_set_encap_source_reply" }
func (*SrSetEncapSourceReply) GetCrcString() string   { return "e8d4e804" }
func (*SrSetEncapSourceReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SrSetEncapSourceReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *SrSetEncapSourceReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *SrSetEncapSourceReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// SrSteeringAddDel defines message 'sr_steering_add_del'.
type SrSteeringAddDel struct {
	IsDel         uint8  `binapi:"u8,name=is_del" json:"is_del,omitempty"`
	BsidAddr      []byte `binapi:"u8[16],name=bsid_addr" json:"bsid_addr,omitempty"`
	SrPolicyIndex uint32 `binapi:"u32,name=sr_policy_index" json:"sr_policy_index,omitempty"`
	TableID       uint32 `binapi:"u32,name=table_id" json:"table_id,omitempty"`
	PrefixAddr    []byte `binapi:"u8[16],name=prefix_addr" json:"prefix_addr,omitempty"`
	MaskWidth     uint32 `binapi:"u32,name=mask_width" json:"mask_width,omitempty"`
	SwIfIndex     uint32 `binapi:"u32,name=sw_if_index" json:"sw_if_index,omitempty"`
	TrafficType   uint8  `binapi:"u8,name=traffic_type" json:"traffic_type,omitempty"`
}

func (m *SrSteeringAddDel) Reset()               { *m = SrSteeringAddDel{} }
func (*SrSteeringAddDel) GetMessageName() string { return "sr_steering_add_del" }
func (*SrSteeringAddDel) GetCrcString() string   { return "28b5dcab" }
func (*SrSteeringAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SrSteeringAddDel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1      // m.IsDel
	size += 1 * 16 // m.BsidAddr
	size += 4      // m.SrPolicyIndex
	size += 4      // m.TableID
	size += 1 * 16 // m.PrefixAddr
	size += 4      // m.MaskWidth
	size += 4      // m.SwIfIndex
	size += 1      // m.TrafficType
	return size
}
func (m *SrSteeringAddDel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(m.IsDel)
	buf.EncodeBytes(m.BsidAddr, 16)
	buf.EncodeUint32(m.SrPolicyIndex)
	buf.EncodeUint32(m.TableID)
	buf.EncodeBytes(m.PrefixAddr, 16)
	buf.EncodeUint32(m.MaskWidth)
	buf.EncodeUint32(m.SwIfIndex)
	buf.EncodeUint8(m.TrafficType)
	return buf.Bytes(), nil
}
func (m *SrSteeringAddDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsDel = buf.DecodeUint8()
	m.BsidAddr = make([]byte, 16)
	copy(m.BsidAddr, buf.DecodeBytes(len(m.BsidAddr)))
	m.SrPolicyIndex = buf.DecodeUint32()
	m.TableID = buf.DecodeUint32()
	m.PrefixAddr = make([]byte, 16)
	copy(m.PrefixAddr, buf.DecodeBytes(len(m.PrefixAddr)))
	m.MaskWidth = buf.DecodeUint32()
	m.SwIfIndex = buf.DecodeUint32()
	m.TrafficType = buf.DecodeUint8()
	return nil
}

// SrSteeringAddDelReply defines message 'sr_steering_add_del_reply'.
type SrSteeringAddDelReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *SrSteeringAddDelReply) Reset()               { *m = SrSteeringAddDelReply{} }
func (*SrSteeringAddDelReply) GetMessageName() string { return "sr_steering_add_del_reply" }
func (*SrSteeringAddDelReply) GetCrcString() string   { return "e8d4e804" }
func (*SrSteeringAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SrSteeringAddDelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *SrSteeringAddDelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *SrSteeringAddDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// SrSteeringPolDetails defines message 'sr_steering_pol_details'.
type SrSteeringPolDetails struct {
	TrafficType uint8   `binapi:"u8,name=traffic_type" json:"traffic_type,omitempty"`
	FibTable    uint32  `binapi:"u32,name=fib_table" json:"fib_table,omitempty"`
	PrefixAddr  []byte  `binapi:"u8[16],name=prefix_addr" json:"prefix_addr,omitempty"`
	MaskWidth   uint32  `binapi:"u32,name=mask_width" json:"mask_width,omitempty"`
	SwIfIndex   uint32  `binapi:"u32,name=sw_if_index" json:"sw_if_index,omitempty"`
	Bsid        Srv6Sid `binapi:"srv6_sid,name=bsid" json:"bsid,omitempty"`
}

func (m *SrSteeringPolDetails) Reset()               { *m = SrSteeringPolDetails{} }
func (*SrSteeringPolDetails) GetMessageName() string { return "sr_steering_pol_details" }
func (*SrSteeringPolDetails) GetCrcString() string   { return "5627d483" }
func (*SrSteeringPolDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SrSteeringPolDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1      // m.TrafficType
	size += 4      // m.FibTable
	size += 1 * 16 // m.PrefixAddr
	size += 4      // m.MaskWidth
	size += 4      // m.SwIfIndex
	size += 1 * 16 // m.Bsid.Addr
	return size
}
func (m *SrSteeringPolDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(m.TrafficType)
	buf.EncodeUint32(m.FibTable)
	buf.EncodeBytes(m.PrefixAddr, 16)
	buf.EncodeUint32(m.MaskWidth)
	buf.EncodeUint32(m.SwIfIndex)
	buf.EncodeBytes(m.Bsid.Addr, 16)
	return buf.Bytes(), nil
}
func (m *SrSteeringPolDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.TrafficType = buf.DecodeUint8()
	m.FibTable = buf.DecodeUint32()
	m.PrefixAddr = make([]byte, 16)
	copy(m.PrefixAddr, buf.DecodeBytes(len(m.PrefixAddr)))
	m.MaskWidth = buf.DecodeUint32()
	m.SwIfIndex = buf.DecodeUint32()
	m.Bsid.Addr = make([]byte, 16)
	copy(m.Bsid.Addr, buf.DecodeBytes(len(m.Bsid.Addr)))
	return nil
}

// SrSteeringPolDump defines message 'sr_steering_pol_dump'.
type SrSteeringPolDump struct{}

func (m *SrSteeringPolDump) Reset()               { *m = SrSteeringPolDump{} }
func (*SrSteeringPolDump) GetMessageName() string { return "sr_steering_pol_dump" }
func (*SrSteeringPolDump) GetCrcString() string   { return "51077d14" }
func (*SrSteeringPolDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SrSteeringPolDump) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *SrSteeringPolDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *SrSteeringPolDump) Unmarshal(b []byte) error {
	return nil
}

func init() { file_sr_binapi_init() }
func file_sr_binapi_init() {
	api.RegisterMessage((*SrLocalsidAddDel)(nil), "sr_localsid_add_del_b30489eb")
	api.RegisterMessage((*SrLocalsidAddDelReply)(nil), "sr_localsid_add_del_reply_e8d4e804")
	api.RegisterMessage((*SrLocalsidsDetails)(nil), "sr_localsids_details_0791babc")
	api.RegisterMessage((*SrLocalsidsDump)(nil), "sr_localsids_dump_51077d14")
	api.RegisterMessage((*SrPoliciesDetails)(nil), "sr_policies_details_5087f460")
	api.RegisterMessage((*SrPoliciesDump)(nil), "sr_policies_dump_51077d14")
	api.RegisterMessage((*SrPolicyAdd)(nil), "sr_policy_add_4b6e2484")
	api.RegisterMessage((*SrPolicyAddReply)(nil), "sr_policy_add_reply_e8d4e804")
	api.RegisterMessage((*SrPolicyDel)(nil), "sr_policy_del_e4133171")
	api.RegisterMessage((*SrPolicyDelReply)(nil), "sr_policy_del_reply_e8d4e804")
	api.RegisterMessage((*SrPolicyMod)(nil), "sr_policy_mod_c1dfaee0")
	api.RegisterMessage((*SrPolicyModReply)(nil), "sr_policy_mod_reply_e8d4e804")
	api.RegisterMessage((*SrSetEncapHopLimit)(nil), "sr_set_encap_hop_limit_aa75d7d0")
	api.RegisterMessage((*SrSetEncapHopLimitReply)(nil), "sr_set_encap_hop_limit_reply_e8d4e804")
	api.RegisterMessage((*SrSetEncapSource)(nil), "sr_set_encap_source_d05bb4de")
	api.RegisterMessage((*SrSetEncapSourceReply)(nil), "sr_set_encap_source_reply_e8d4e804")
	api.RegisterMessage((*SrSteeringAddDel)(nil), "sr_steering_add_del_28b5dcab")
	api.RegisterMessage((*SrSteeringAddDelReply)(nil), "sr_steering_add_del_reply_e8d4e804")
	api.RegisterMessage((*SrSteeringPolDetails)(nil), "sr_steering_pol_details_5627d483")
	api.RegisterMessage((*SrSteeringPolDump)(nil), "sr_steering_pol_dump_51077d14")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*SrLocalsidAddDel)(nil),
		(*SrLocalsidAddDelReply)(nil),
		(*SrLocalsidsDetails)(nil),
		(*SrLocalsidsDump)(nil),
		(*SrPoliciesDetails)(nil),
		(*SrPoliciesDump)(nil),
		(*SrPolicyAdd)(nil),
		(*SrPolicyAddReply)(nil),
		(*SrPolicyDel)(nil),
		(*SrPolicyDelReply)(nil),
		(*SrPolicyMod)(nil),
		(*SrPolicyModReply)(nil),
		(*SrSetEncapHopLimit)(nil),
		(*SrSetEncapHopLimitReply)(nil),
		(*SrSetEncapSource)(nil),
		(*SrSetEncapSourceReply)(nil),
		(*SrSteeringAddDel)(nil),
		(*SrSteeringAddDelReply)(nil),
		(*SrSteeringPolDetails)(nil),
		(*SrSteeringPolDump)(nil),
	}
}
