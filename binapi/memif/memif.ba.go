// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.5.0-dev
//  VPP:              22.02-release
// source: /usr/share/vpp/api/plugins/memif.api.json

// Package memif contains generated bindings for API file memif.api.
//
// Contents:
//   2 enums
//  10 messages
//
package memif

import (
	"strconv"

	api "git.fd.io/govpp.git/api"
	ethernet_types "git.fd.io/govpp.git/binapi/ethernet_types"
	interface_types "git.fd.io/govpp.git/binapi/interface_types"
	codec "git.fd.io/govpp.git/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "memif"
	APIVersion = "3.0.0"
	VersionCrc = 0xd189f1e1
)

// MemifMode defines enum 'memif_mode'.
type MemifMode uint32

const (
	MEMIF_MODE_API_ETHERNET    MemifMode = 0
	MEMIF_MODE_API_IP          MemifMode = 1
	MEMIF_MODE_API_PUNT_INJECT MemifMode = 2
)

var (
	MemifMode_name = map[uint32]string{
		0: "MEMIF_MODE_API_ETHERNET",
		1: "MEMIF_MODE_API_IP",
		2: "MEMIF_MODE_API_PUNT_INJECT",
	}
	MemifMode_value = map[string]uint32{
		"MEMIF_MODE_API_ETHERNET":    0,
		"MEMIF_MODE_API_IP":          1,
		"MEMIF_MODE_API_PUNT_INJECT": 2,
	}
)

func (x MemifMode) String() string {
	s, ok := MemifMode_name[uint32(x)]
	if ok {
		return s
	}
	return "MemifMode(" + strconv.Itoa(int(x)) + ")"
}

// MemifRole defines enum 'memif_role'.
type MemifRole uint32

const (
	MEMIF_ROLE_API_MASTER MemifRole = 0
	MEMIF_ROLE_API_SLAVE  MemifRole = 1
)

var (
	MemifRole_name = map[uint32]string{
		0: "MEMIF_ROLE_API_MASTER",
		1: "MEMIF_ROLE_API_SLAVE",
	}
	MemifRole_value = map[string]uint32{
		"MEMIF_ROLE_API_MASTER": 0,
		"MEMIF_ROLE_API_SLAVE":  1,
	}
)

func (x MemifRole) String() string {
	s, ok := MemifRole_name[uint32(x)]
	if ok {
		return s
	}
	return "MemifRole(" + strconv.Itoa(int(x)) + ")"
}

// MemifCreate defines message 'memif_create'.
type MemifCreate struct {
	Role       MemifRole                 `binapi:"memif_role,name=role" json:"role,omitempty"`
	Mode       MemifMode                 `binapi:"memif_mode,name=mode" json:"mode,omitempty"`
	RxQueues   uint8                     `binapi:"u8,name=rx_queues" json:"rx_queues,omitempty"`
	TxQueues   uint8                     `binapi:"u8,name=tx_queues" json:"tx_queues,omitempty"`
	ID         uint32                    `binapi:"u32,name=id" json:"id,omitempty"`
	SocketID   uint32                    `binapi:"u32,name=socket_id" json:"socket_id,omitempty"`
	RingSize   uint32                    `binapi:"u32,name=ring_size" json:"ring_size,omitempty"`
	BufferSize uint16                    `binapi:"u16,name=buffer_size" json:"buffer_size,omitempty"`
	NoZeroCopy bool                      `binapi:"bool,name=no_zero_copy" json:"no_zero_copy,omitempty"`
	HwAddr     ethernet_types.MacAddress `binapi:"mac_address,name=hw_addr" json:"hw_addr,omitempty"`
	Secret     string                    `binapi:"string[24],name=secret" json:"secret,omitempty"`
}

func (m *MemifCreate) Reset()               { *m = MemifCreate{} }
func (*MemifCreate) GetMessageName() string { return "memif_create" }
func (*MemifCreate) GetCrcString() string   { return "b1b25061" }
func (*MemifCreate) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *MemifCreate) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4     // m.Role
	size += 4     // m.Mode
	size += 1     // m.RxQueues
	size += 1     // m.TxQueues
	size += 4     // m.ID
	size += 4     // m.SocketID
	size += 4     // m.RingSize
	size += 2     // m.BufferSize
	size += 1     // m.NoZeroCopy
	size += 1 * 6 // m.HwAddr
	size += 24    // m.Secret
	return size
}
func (m *MemifCreate) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.Role))
	buf.EncodeUint32(uint32(m.Mode))
	buf.EncodeUint8(m.RxQueues)
	buf.EncodeUint8(m.TxQueues)
	buf.EncodeUint32(m.ID)
	buf.EncodeUint32(m.SocketID)
	buf.EncodeUint32(m.RingSize)
	buf.EncodeUint16(m.BufferSize)
	buf.EncodeBool(m.NoZeroCopy)
	buf.EncodeBytes(m.HwAddr[:], 6)
	buf.EncodeString(m.Secret, 24)
	return buf.Bytes(), nil
}
func (m *MemifCreate) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Role = MemifRole(buf.DecodeUint32())
	m.Mode = MemifMode(buf.DecodeUint32())
	m.RxQueues = buf.DecodeUint8()
	m.TxQueues = buf.DecodeUint8()
	m.ID = buf.DecodeUint32()
	m.SocketID = buf.DecodeUint32()
	m.RingSize = buf.DecodeUint32()
	m.BufferSize = buf.DecodeUint16()
	m.NoZeroCopy = buf.DecodeBool()
	copy(m.HwAddr[:], buf.DecodeBytes(6))
	m.Secret = buf.DecodeString(24)
	return nil
}

// MemifCreateReply defines message 'memif_create_reply'.
type MemifCreateReply struct {
	Retval    int32                          `binapi:"i32,name=retval" json:"retval,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *MemifCreateReply) Reset()               { *m = MemifCreateReply{} }
func (*MemifCreateReply) GetMessageName() string { return "memif_create_reply" }
func (*MemifCreateReply) GetCrcString() string   { return "5383d31f" }
func (*MemifCreateReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *MemifCreateReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.SwIfIndex
	return size
}
func (m *MemifCreateReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *MemifCreateReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// MemifDelete defines message 'memif_delete'.
type MemifDelete struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *MemifDelete) Reset()               { *m = MemifDelete{} }
func (*MemifDelete) GetMessageName() string { return "memif_delete" }
func (*MemifDelete) GetCrcString() string   { return "f9e6675e" }
func (*MemifDelete) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *MemifDelete) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	return size
}
func (m *MemifDelete) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *MemifDelete) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// MemifDeleteReply defines message 'memif_delete_reply'.
type MemifDeleteReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *MemifDeleteReply) Reset()               { *m = MemifDeleteReply{} }
func (*MemifDeleteReply) GetMessageName() string { return "memif_delete_reply" }
func (*MemifDeleteReply) GetCrcString() string   { return "e8d4e804" }
func (*MemifDeleteReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *MemifDeleteReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *MemifDeleteReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *MemifDeleteReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// MemifDetails defines message 'memif_details'.
type MemifDetails struct {
	SwIfIndex  interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	HwAddr     ethernet_types.MacAddress      `binapi:"mac_address,name=hw_addr" json:"hw_addr,omitempty"`
	ID         uint32                         `binapi:"u32,name=id" json:"id,omitempty"`
	Role       MemifRole                      `binapi:"memif_role,name=role" json:"role,omitempty"`
	Mode       MemifMode                      `binapi:"memif_mode,name=mode" json:"mode,omitempty"`
	ZeroCopy   bool                           `binapi:"bool,name=zero_copy" json:"zero_copy,omitempty"`
	SocketID   uint32                         `binapi:"u32,name=socket_id" json:"socket_id,omitempty"`
	RingSize   uint32                         `binapi:"u32,name=ring_size" json:"ring_size,omitempty"`
	BufferSize uint16                         `binapi:"u16,name=buffer_size" json:"buffer_size,omitempty"`
	Flags      interface_types.IfStatusFlags  `binapi:"if_status_flags,name=flags" json:"flags,omitempty"`
	IfName     string                         `binapi:"string[64],name=if_name" json:"if_name,omitempty"`
}

func (m *MemifDetails) Reset()               { *m = MemifDetails{} }
func (*MemifDetails) GetMessageName() string { return "memif_details" }
func (*MemifDetails) GetCrcString() string   { return "da34feb9" }
func (*MemifDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *MemifDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4     // m.SwIfIndex
	size += 1 * 6 // m.HwAddr
	size += 4     // m.ID
	size += 4     // m.Role
	size += 4     // m.Mode
	size += 1     // m.ZeroCopy
	size += 4     // m.SocketID
	size += 4     // m.RingSize
	size += 2     // m.BufferSize
	size += 4     // m.Flags
	size += 64    // m.IfName
	return size
}
func (m *MemifDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeBytes(m.HwAddr[:], 6)
	buf.EncodeUint32(m.ID)
	buf.EncodeUint32(uint32(m.Role))
	buf.EncodeUint32(uint32(m.Mode))
	buf.EncodeBool(m.ZeroCopy)
	buf.EncodeUint32(m.SocketID)
	buf.EncodeUint32(m.RingSize)
	buf.EncodeUint16(m.BufferSize)
	buf.EncodeUint32(uint32(m.Flags))
	buf.EncodeString(m.IfName, 64)
	return buf.Bytes(), nil
}
func (m *MemifDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	copy(m.HwAddr[:], buf.DecodeBytes(6))
	m.ID = buf.DecodeUint32()
	m.Role = MemifRole(buf.DecodeUint32())
	m.Mode = MemifMode(buf.DecodeUint32())
	m.ZeroCopy = buf.DecodeBool()
	m.SocketID = buf.DecodeUint32()
	m.RingSize = buf.DecodeUint32()
	m.BufferSize = buf.DecodeUint16()
	m.Flags = interface_types.IfStatusFlags(buf.DecodeUint32())
	m.IfName = buf.DecodeString(64)
	return nil
}

// MemifDump defines message 'memif_dump'.
type MemifDump struct{}

func (m *MemifDump) Reset()               { *m = MemifDump{} }
func (*MemifDump) GetMessageName() string { return "memif_dump" }
func (*MemifDump) GetCrcString() string   { return "51077d14" }
func (*MemifDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *MemifDump) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *MemifDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *MemifDump) Unmarshal(b []byte) error {
	return nil
}

// MemifSocketFilenameAddDel defines message 'memif_socket_filename_add_del'.
type MemifSocketFilenameAddDel struct {
	IsAdd          bool   `binapi:"bool,name=is_add" json:"is_add,omitempty"`
	SocketID       uint32 `binapi:"u32,name=socket_id" json:"socket_id,omitempty"`
	SocketFilename string `binapi:"string[108],name=socket_filename" json:"socket_filename,omitempty"`
}

func (m *MemifSocketFilenameAddDel) Reset()               { *m = MemifSocketFilenameAddDel{} }
func (*MemifSocketFilenameAddDel) GetMessageName() string { return "memif_socket_filename_add_del" }
func (*MemifSocketFilenameAddDel) GetCrcString() string   { return "a2ce1a10" }
func (*MemifSocketFilenameAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *MemifSocketFilenameAddDel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1   // m.IsAdd
	size += 4   // m.SocketID
	size += 108 // m.SocketFilename
	return size
}
func (m *MemifSocketFilenameAddDel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.IsAdd)
	buf.EncodeUint32(m.SocketID)
	buf.EncodeString(m.SocketFilename, 108)
	return buf.Bytes(), nil
}
func (m *MemifSocketFilenameAddDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsAdd = buf.DecodeBool()
	m.SocketID = buf.DecodeUint32()
	m.SocketFilename = buf.DecodeString(108)
	return nil
}

// MemifSocketFilenameAddDelReply defines message 'memif_socket_filename_add_del_reply'.
type MemifSocketFilenameAddDelReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *MemifSocketFilenameAddDelReply) Reset() { *m = MemifSocketFilenameAddDelReply{} }
func (*MemifSocketFilenameAddDelReply) GetMessageName() string {
	return "memif_socket_filename_add_del_reply"
}
func (*MemifSocketFilenameAddDelReply) GetCrcString() string { return "e8d4e804" }
func (*MemifSocketFilenameAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *MemifSocketFilenameAddDelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *MemifSocketFilenameAddDelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *MemifSocketFilenameAddDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// MemifSocketFilenameDetails defines message 'memif_socket_filename_details'.
type MemifSocketFilenameDetails struct {
	SocketID       uint32 `binapi:"u32,name=socket_id" json:"socket_id,omitempty"`
	SocketFilename string `binapi:"string[108],name=socket_filename" json:"socket_filename,omitempty"`
}

func (m *MemifSocketFilenameDetails) Reset()               { *m = MemifSocketFilenameDetails{} }
func (*MemifSocketFilenameDetails) GetMessageName() string { return "memif_socket_filename_details" }
func (*MemifSocketFilenameDetails) GetCrcString() string   { return "7ff326f7" }
func (*MemifSocketFilenameDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *MemifSocketFilenameDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4   // m.SocketID
	size += 108 // m.SocketFilename
	return size
}
func (m *MemifSocketFilenameDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.SocketID)
	buf.EncodeString(m.SocketFilename, 108)
	return buf.Bytes(), nil
}
func (m *MemifSocketFilenameDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SocketID = buf.DecodeUint32()
	m.SocketFilename = buf.DecodeString(108)
	return nil
}

// MemifSocketFilenameDump defines message 'memif_socket_filename_dump'.
type MemifSocketFilenameDump struct{}

func (m *MemifSocketFilenameDump) Reset()               { *m = MemifSocketFilenameDump{} }
func (*MemifSocketFilenameDump) GetMessageName() string { return "memif_socket_filename_dump" }
func (*MemifSocketFilenameDump) GetCrcString() string   { return "51077d14" }
func (*MemifSocketFilenameDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *MemifSocketFilenameDump) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *MemifSocketFilenameDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *MemifSocketFilenameDump) Unmarshal(b []byte) error {
	return nil
}

func init() { file_memif_binapi_init() }
func file_memif_binapi_init() {
	api.RegisterMessage((*MemifCreate)(nil), "memif_create_b1b25061")
	api.RegisterMessage((*MemifCreateReply)(nil), "memif_create_reply_5383d31f")
	api.RegisterMessage((*MemifDelete)(nil), "memif_delete_f9e6675e")
	api.RegisterMessage((*MemifDeleteReply)(nil), "memif_delete_reply_e8d4e804")
	api.RegisterMessage((*MemifDetails)(nil), "memif_details_da34feb9")
	api.RegisterMessage((*MemifDump)(nil), "memif_dump_51077d14")
	api.RegisterMessage((*MemifSocketFilenameAddDel)(nil), "memif_socket_filename_add_del_a2ce1a10")
	api.RegisterMessage((*MemifSocketFilenameAddDelReply)(nil), "memif_socket_filename_add_del_reply_e8d4e804")
	api.RegisterMessage((*MemifSocketFilenameDetails)(nil), "memif_socket_filename_details_7ff326f7")
	api.RegisterMessage((*MemifSocketFilenameDump)(nil), "memif_socket_filename_dump_51077d14")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*MemifCreate)(nil),
		(*MemifCreateReply)(nil),
		(*MemifDelete)(nil),
		(*MemifDeleteReply)(nil),
		(*MemifDetails)(nil),
		(*MemifDump)(nil),
		(*MemifSocketFilenameAddDel)(nil),
		(*MemifSocketFilenameAddDelReply)(nil),
		(*MemifSocketFilenameDetails)(nil),
		(*MemifSocketFilenameDump)(nil),
	}
}
