// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.7.0
//  VPP:              22.10-release
// source: /usr/share/vpp/api/plugins/abf.api.json

// Package abf contains generated bindings for API file abf.api.
//
// Contents:
// -  2 structs
// - 10 messages
package abf

import (
	api "go.fd.io/govpp/api"
	fib_types "go.fd.io/govpp/binapi/fib_types"
	interface_types "go.fd.io/govpp/binapi/interface_types"
	_ "go.fd.io/govpp/binapi/ip_types"
	codec "go.fd.io/govpp/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "abf"
	APIVersion = "1.0.0"
	VersionCrc = 0xf2367b47
)

// AbfItfAttach defines type 'abf_itf_attach'.
type AbfItfAttach struct {
	PolicyID  uint32                         `binapi:"u32,name=policy_id" json:"policy_id,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	Priority  uint32                         `binapi:"u32,name=priority" json:"priority,omitempty"`
	IsIPv6    bool                           `binapi:"bool,name=is_ipv6" json:"is_ipv6,omitempty"`
}

// AbfPolicy defines type 'abf_policy'.
type AbfPolicy struct {
	PolicyID uint32              `binapi:"u32,name=policy_id" json:"policy_id,omitempty"`
	ACLIndex uint32              `binapi:"u32,name=acl_index" json:"acl_index,omitempty"`
	NPaths   uint8               `binapi:"u8,name=n_paths" json:"-"`
	Paths    []fib_types.FibPath `binapi:"fib_path[n_paths],name=paths" json:"paths,omitempty"`
}

// AbfItfAttachAddDel defines message 'abf_itf_attach_add_del'.
// InProgress: the message form may change in the future versions
type AbfItfAttachAddDel struct {
	IsAdd  bool         `binapi:"bool,name=is_add" json:"is_add,omitempty"`
	Attach AbfItfAttach `binapi:"abf_itf_attach,name=attach" json:"attach,omitempty"`
}

func (m *AbfItfAttachAddDel) Reset()               { *m = AbfItfAttachAddDel{} }
func (*AbfItfAttachAddDel) GetMessageName() string { return "abf_itf_attach_add_del" }
func (*AbfItfAttachAddDel) GetCrcString() string   { return "25c8621b" }
func (*AbfItfAttachAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *AbfItfAttachAddDel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.IsAdd
	size += 4 // m.Attach.PolicyID
	size += 4 // m.Attach.SwIfIndex
	size += 4 // m.Attach.Priority
	size += 1 // m.Attach.IsIPv6
	return size
}
func (m *AbfItfAttachAddDel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.IsAdd)
	buf.EncodeUint32(m.Attach.PolicyID)
	buf.EncodeUint32(uint32(m.Attach.SwIfIndex))
	buf.EncodeUint32(m.Attach.Priority)
	buf.EncodeBool(m.Attach.IsIPv6)
	return buf.Bytes(), nil
}
func (m *AbfItfAttachAddDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsAdd = buf.DecodeBool()
	m.Attach.PolicyID = buf.DecodeUint32()
	m.Attach.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.Attach.Priority = buf.DecodeUint32()
	m.Attach.IsIPv6 = buf.DecodeBool()
	return nil
}

// AbfItfAttachAddDelReply defines message 'abf_itf_attach_add_del_reply'.
// InProgress: the message form may change in the future versions
type AbfItfAttachAddDelReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *AbfItfAttachAddDelReply) Reset()               { *m = AbfItfAttachAddDelReply{} }
func (*AbfItfAttachAddDelReply) GetMessageName() string { return "abf_itf_attach_add_del_reply" }
func (*AbfItfAttachAddDelReply) GetCrcString() string   { return "e8d4e804" }
func (*AbfItfAttachAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *AbfItfAttachAddDelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *AbfItfAttachAddDelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *AbfItfAttachAddDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// AbfItfAttachDetails defines message 'abf_itf_attach_details'.
// InProgress: the message form may change in the future versions
type AbfItfAttachDetails struct {
	Attach AbfItfAttach `binapi:"abf_itf_attach,name=attach" json:"attach,omitempty"`
}

func (m *AbfItfAttachDetails) Reset()               { *m = AbfItfAttachDetails{} }
func (*AbfItfAttachDetails) GetMessageName() string { return "abf_itf_attach_details" }
func (*AbfItfAttachDetails) GetCrcString() string   { return "7819523e" }
func (*AbfItfAttachDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *AbfItfAttachDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Attach.PolicyID
	size += 4 // m.Attach.SwIfIndex
	size += 4 // m.Attach.Priority
	size += 1 // m.Attach.IsIPv6
	return size
}
func (m *AbfItfAttachDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.Attach.PolicyID)
	buf.EncodeUint32(uint32(m.Attach.SwIfIndex))
	buf.EncodeUint32(m.Attach.Priority)
	buf.EncodeBool(m.Attach.IsIPv6)
	return buf.Bytes(), nil
}
func (m *AbfItfAttachDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Attach.PolicyID = buf.DecodeUint32()
	m.Attach.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.Attach.Priority = buf.DecodeUint32()
	m.Attach.IsIPv6 = buf.DecodeBool()
	return nil
}

// AbfItfAttachDump defines message 'abf_itf_attach_dump'.
// InProgress: the message form may change in the future versions
type AbfItfAttachDump struct{}

func (m *AbfItfAttachDump) Reset()               { *m = AbfItfAttachDump{} }
func (*AbfItfAttachDump) GetMessageName() string { return "abf_itf_attach_dump" }
func (*AbfItfAttachDump) GetCrcString() string   { return "51077d14" }
func (*AbfItfAttachDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *AbfItfAttachDump) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *AbfItfAttachDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *AbfItfAttachDump) Unmarshal(b []byte) error {
	return nil
}

// AbfPluginGetVersion defines message 'abf_plugin_get_version'.
// InProgress: the message form may change in the future versions
type AbfPluginGetVersion struct{}

func (m *AbfPluginGetVersion) Reset()               { *m = AbfPluginGetVersion{} }
func (*AbfPluginGetVersion) GetMessageName() string { return "abf_plugin_get_version" }
func (*AbfPluginGetVersion) GetCrcString() string   { return "51077d14" }
func (*AbfPluginGetVersion) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *AbfPluginGetVersion) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *AbfPluginGetVersion) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *AbfPluginGetVersion) Unmarshal(b []byte) error {
	return nil
}

// AbfPluginGetVersionReply defines message 'abf_plugin_get_version_reply'.
// InProgress: the message form may change in the future versions
type AbfPluginGetVersionReply struct {
	Major uint32 `binapi:"u32,name=major" json:"major,omitempty"`
	Minor uint32 `binapi:"u32,name=minor" json:"minor,omitempty"`
}

func (m *AbfPluginGetVersionReply) Reset()               { *m = AbfPluginGetVersionReply{} }
func (*AbfPluginGetVersionReply) GetMessageName() string { return "abf_plugin_get_version_reply" }
func (*AbfPluginGetVersionReply) GetCrcString() string   { return "9b32cf86" }
func (*AbfPluginGetVersionReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *AbfPluginGetVersionReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Major
	size += 4 // m.Minor
	return size
}
func (m *AbfPluginGetVersionReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.Major)
	buf.EncodeUint32(m.Minor)
	return buf.Bytes(), nil
}
func (m *AbfPluginGetVersionReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Major = buf.DecodeUint32()
	m.Minor = buf.DecodeUint32()
	return nil
}

// AbfPolicyAddDel defines message 'abf_policy_add_del'.
// InProgress: the message form may change in the future versions
type AbfPolicyAddDel struct {
	IsAdd  bool      `binapi:"bool,name=is_add" json:"is_add,omitempty"`
	Policy AbfPolicy `binapi:"abf_policy,name=policy" json:"policy,omitempty"`
}

func (m *AbfPolicyAddDel) Reset()               { *m = AbfPolicyAddDel{} }
func (*AbfPolicyAddDel) GetMessageName() string { return "abf_policy_add_del" }
func (*AbfPolicyAddDel) GetCrcString() string   { return "c6131197" }
func (*AbfPolicyAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *AbfPolicyAddDel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.IsAdd
	size += 4 // m.Policy.PolicyID
	size += 4 // m.Policy.ACLIndex
	size += 1 // m.Policy.NPaths
	for j2 := 0; j2 < len(m.Policy.Paths); j2++ {
		var s2 fib_types.FibPath
		_ = s2
		if j2 < len(m.Policy.Paths) {
			s2 = m.Policy.Paths[j2]
		}
		size += 4      // s2.SwIfIndex
		size += 4      // s2.TableID
		size += 4      // s2.RpfID
		size += 1      // s2.Weight
		size += 1      // s2.Preference
		size += 4      // s2.Type
		size += 4      // s2.Flags
		size += 4      // s2.Proto
		size += 1 * 16 // s2.Nh.Address
		size += 4      // s2.Nh.ViaLabel
		size += 4      // s2.Nh.ObjID
		size += 4      // s2.Nh.ClassifyTableIndex
		size += 1      // s2.NLabels
		for j3 := 0; j3 < 16; j3++ {
			size += 1 // s2.LabelStack[j3].IsUniform
			size += 4 // s2.LabelStack[j3].Label
			size += 1 // s2.LabelStack[j3].TTL
			size += 1 // s2.LabelStack[j3].Exp
		}
	}
	return size
}
func (m *AbfPolicyAddDel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.IsAdd)
	buf.EncodeUint32(m.Policy.PolicyID)
	buf.EncodeUint32(m.Policy.ACLIndex)
	buf.EncodeUint8(uint8(len(m.Policy.Paths)))
	for j1 := 0; j1 < len(m.Policy.Paths); j1++ {
		var v1 fib_types.FibPath // Paths
		if j1 < len(m.Policy.Paths) {
			v1 = m.Policy.Paths[j1]
		}
		buf.EncodeUint32(v1.SwIfIndex)
		buf.EncodeUint32(v1.TableID)
		buf.EncodeUint32(v1.RpfID)
		buf.EncodeUint8(v1.Weight)
		buf.EncodeUint8(v1.Preference)
		buf.EncodeUint32(uint32(v1.Type))
		buf.EncodeUint32(uint32(v1.Flags))
		buf.EncodeUint32(uint32(v1.Proto))
		buf.EncodeBytes(v1.Nh.Address.XXX_UnionData[:], 16)
		buf.EncodeUint32(v1.Nh.ViaLabel)
		buf.EncodeUint32(v1.Nh.ObjID)
		buf.EncodeUint32(v1.Nh.ClassifyTableIndex)
		buf.EncodeUint8(v1.NLabels)
		for j2 := 0; j2 < 16; j2++ {
			buf.EncodeUint8(v1.LabelStack[j2].IsUniform)
			buf.EncodeUint32(v1.LabelStack[j2].Label)
			buf.EncodeUint8(v1.LabelStack[j2].TTL)
			buf.EncodeUint8(v1.LabelStack[j2].Exp)
		}
	}
	return buf.Bytes(), nil
}
func (m *AbfPolicyAddDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsAdd = buf.DecodeBool()
	m.Policy.PolicyID = buf.DecodeUint32()
	m.Policy.ACLIndex = buf.DecodeUint32()
	m.Policy.NPaths = buf.DecodeUint8()
	m.Policy.Paths = make([]fib_types.FibPath, m.Policy.NPaths)
	for j1 := 0; j1 < len(m.Policy.Paths); j1++ {
		m.Policy.Paths[j1].SwIfIndex = buf.DecodeUint32()
		m.Policy.Paths[j1].TableID = buf.DecodeUint32()
		m.Policy.Paths[j1].RpfID = buf.DecodeUint32()
		m.Policy.Paths[j1].Weight = buf.DecodeUint8()
		m.Policy.Paths[j1].Preference = buf.DecodeUint8()
		m.Policy.Paths[j1].Type = fib_types.FibPathType(buf.DecodeUint32())
		m.Policy.Paths[j1].Flags = fib_types.FibPathFlags(buf.DecodeUint32())
		m.Policy.Paths[j1].Proto = fib_types.FibPathNhProto(buf.DecodeUint32())
		copy(m.Policy.Paths[j1].Nh.Address.XXX_UnionData[:], buf.DecodeBytes(16))
		m.Policy.Paths[j1].Nh.ViaLabel = buf.DecodeUint32()
		m.Policy.Paths[j1].Nh.ObjID = buf.DecodeUint32()
		m.Policy.Paths[j1].Nh.ClassifyTableIndex = buf.DecodeUint32()
		m.Policy.Paths[j1].NLabels = buf.DecodeUint8()
		for j2 := 0; j2 < 16; j2++ {
			m.Policy.Paths[j1].LabelStack[j2].IsUniform = buf.DecodeUint8()
			m.Policy.Paths[j1].LabelStack[j2].Label = buf.DecodeUint32()
			m.Policy.Paths[j1].LabelStack[j2].TTL = buf.DecodeUint8()
			m.Policy.Paths[j1].LabelStack[j2].Exp = buf.DecodeUint8()
		}
	}
	return nil
}

// AbfPolicyAddDelReply defines message 'abf_policy_add_del_reply'.
// InProgress: the message form may change in the future versions
type AbfPolicyAddDelReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *AbfPolicyAddDelReply) Reset()               { *m = AbfPolicyAddDelReply{} }
func (*AbfPolicyAddDelReply) GetMessageName() string { return "abf_policy_add_del_reply" }
func (*AbfPolicyAddDelReply) GetCrcString() string   { return "e8d4e804" }
func (*AbfPolicyAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *AbfPolicyAddDelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *AbfPolicyAddDelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *AbfPolicyAddDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// AbfPolicyDetails defines message 'abf_policy_details'.
// InProgress: the message form may change in the future versions
type AbfPolicyDetails struct {
	Policy AbfPolicy `binapi:"abf_policy,name=policy" json:"policy,omitempty"`
}

func (m *AbfPolicyDetails) Reset()               { *m = AbfPolicyDetails{} }
func (*AbfPolicyDetails) GetMessageName() string { return "abf_policy_details" }
func (*AbfPolicyDetails) GetCrcString() string   { return "b7487fa4" }
func (*AbfPolicyDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *AbfPolicyDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Policy.PolicyID
	size += 4 // m.Policy.ACLIndex
	size += 1 // m.Policy.NPaths
	for j2 := 0; j2 < len(m.Policy.Paths); j2++ {
		var s2 fib_types.FibPath
		_ = s2
		if j2 < len(m.Policy.Paths) {
			s2 = m.Policy.Paths[j2]
		}
		size += 4      // s2.SwIfIndex
		size += 4      // s2.TableID
		size += 4      // s2.RpfID
		size += 1      // s2.Weight
		size += 1      // s2.Preference
		size += 4      // s2.Type
		size += 4      // s2.Flags
		size += 4      // s2.Proto
		size += 1 * 16 // s2.Nh.Address
		size += 4      // s2.Nh.ViaLabel
		size += 4      // s2.Nh.ObjID
		size += 4      // s2.Nh.ClassifyTableIndex
		size += 1      // s2.NLabels
		for j3 := 0; j3 < 16; j3++ {
			size += 1 // s2.LabelStack[j3].IsUniform
			size += 4 // s2.LabelStack[j3].Label
			size += 1 // s2.LabelStack[j3].TTL
			size += 1 // s2.LabelStack[j3].Exp
		}
	}
	return size
}
func (m *AbfPolicyDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.Policy.PolicyID)
	buf.EncodeUint32(m.Policy.ACLIndex)
	buf.EncodeUint8(uint8(len(m.Policy.Paths)))
	for j1 := 0; j1 < len(m.Policy.Paths); j1++ {
		var v1 fib_types.FibPath // Paths
		if j1 < len(m.Policy.Paths) {
			v1 = m.Policy.Paths[j1]
		}
		buf.EncodeUint32(v1.SwIfIndex)
		buf.EncodeUint32(v1.TableID)
		buf.EncodeUint32(v1.RpfID)
		buf.EncodeUint8(v1.Weight)
		buf.EncodeUint8(v1.Preference)
		buf.EncodeUint32(uint32(v1.Type))
		buf.EncodeUint32(uint32(v1.Flags))
		buf.EncodeUint32(uint32(v1.Proto))
		buf.EncodeBytes(v1.Nh.Address.XXX_UnionData[:], 16)
		buf.EncodeUint32(v1.Nh.ViaLabel)
		buf.EncodeUint32(v1.Nh.ObjID)
		buf.EncodeUint32(v1.Nh.ClassifyTableIndex)
		buf.EncodeUint8(v1.NLabels)
		for j2 := 0; j2 < 16; j2++ {
			buf.EncodeUint8(v1.LabelStack[j2].IsUniform)
			buf.EncodeUint32(v1.LabelStack[j2].Label)
			buf.EncodeUint8(v1.LabelStack[j2].TTL)
			buf.EncodeUint8(v1.LabelStack[j2].Exp)
		}
	}
	return buf.Bytes(), nil
}
func (m *AbfPolicyDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Policy.PolicyID = buf.DecodeUint32()
	m.Policy.ACLIndex = buf.DecodeUint32()
	m.Policy.NPaths = buf.DecodeUint8()
	m.Policy.Paths = make([]fib_types.FibPath, m.Policy.NPaths)
	for j1 := 0; j1 < len(m.Policy.Paths); j1++ {
		m.Policy.Paths[j1].SwIfIndex = buf.DecodeUint32()
		m.Policy.Paths[j1].TableID = buf.DecodeUint32()
		m.Policy.Paths[j1].RpfID = buf.DecodeUint32()
		m.Policy.Paths[j1].Weight = buf.DecodeUint8()
		m.Policy.Paths[j1].Preference = buf.DecodeUint8()
		m.Policy.Paths[j1].Type = fib_types.FibPathType(buf.DecodeUint32())
		m.Policy.Paths[j1].Flags = fib_types.FibPathFlags(buf.DecodeUint32())
		m.Policy.Paths[j1].Proto = fib_types.FibPathNhProto(buf.DecodeUint32())
		copy(m.Policy.Paths[j1].Nh.Address.XXX_UnionData[:], buf.DecodeBytes(16))
		m.Policy.Paths[j1].Nh.ViaLabel = buf.DecodeUint32()
		m.Policy.Paths[j1].Nh.ObjID = buf.DecodeUint32()
		m.Policy.Paths[j1].Nh.ClassifyTableIndex = buf.DecodeUint32()
		m.Policy.Paths[j1].NLabels = buf.DecodeUint8()
		for j2 := 0; j2 < 16; j2++ {
			m.Policy.Paths[j1].LabelStack[j2].IsUniform = buf.DecodeUint8()
			m.Policy.Paths[j1].LabelStack[j2].Label = buf.DecodeUint32()
			m.Policy.Paths[j1].LabelStack[j2].TTL = buf.DecodeUint8()
			m.Policy.Paths[j1].LabelStack[j2].Exp = buf.DecodeUint8()
		}
	}
	return nil
}

// AbfPolicyDump defines message 'abf_policy_dump'.
// InProgress: the message form may change in the future versions
type AbfPolicyDump struct{}

func (m *AbfPolicyDump) Reset()               { *m = AbfPolicyDump{} }
func (*AbfPolicyDump) GetMessageName() string { return "abf_policy_dump" }
func (*AbfPolicyDump) GetCrcString() string   { return "51077d14" }
func (*AbfPolicyDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *AbfPolicyDump) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *AbfPolicyDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *AbfPolicyDump) Unmarshal(b []byte) error {
	return nil
}

func init() { file_abf_binapi_init() }
func file_abf_binapi_init() {
	api.RegisterMessage((*AbfItfAttachAddDel)(nil), "abf_itf_attach_add_del_25c8621b")
	api.RegisterMessage((*AbfItfAttachAddDelReply)(nil), "abf_itf_attach_add_del_reply_e8d4e804")
	api.RegisterMessage((*AbfItfAttachDetails)(nil), "abf_itf_attach_details_7819523e")
	api.RegisterMessage((*AbfItfAttachDump)(nil), "abf_itf_attach_dump_51077d14")
	api.RegisterMessage((*AbfPluginGetVersion)(nil), "abf_plugin_get_version_51077d14")
	api.RegisterMessage((*AbfPluginGetVersionReply)(nil), "abf_plugin_get_version_reply_9b32cf86")
	api.RegisterMessage((*AbfPolicyAddDel)(nil), "abf_policy_add_del_c6131197")
	api.RegisterMessage((*AbfPolicyAddDelReply)(nil), "abf_policy_add_del_reply_e8d4e804")
	api.RegisterMessage((*AbfPolicyDetails)(nil), "abf_policy_details_b7487fa4")
	api.RegisterMessage((*AbfPolicyDump)(nil), "abf_policy_dump_51077d14")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*AbfItfAttachAddDel)(nil),
		(*AbfItfAttachAddDelReply)(nil),
		(*AbfItfAttachDetails)(nil),
		(*AbfItfAttachDump)(nil),
		(*AbfPluginGetVersion)(nil),
		(*AbfPluginGetVersionReply)(nil),
		(*AbfPolicyAddDel)(nil),
		(*AbfPolicyAddDelReply)(nil),
		(*AbfPolicyDetails)(nil),
		(*AbfPolicyDump)(nil),
	}
}
