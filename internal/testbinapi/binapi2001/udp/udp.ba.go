// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.5.0-dev
//  VPP:              20.01
// source: .vppapi/core/udp.api.json

// Package udp contains generated bindings for API file udp.api.
//
// Contents:
//   5 aliases
//   4 enums
//   7 structs
//   1 union
//   6 messages
//
package udp

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	api "git.fd.io/govpp.git/api"
	codec "git.fd.io/govpp.git/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "udp"
	APIVersion = "1.1.0"
	VersionCrc = 0xa6b4c475
)

// AddressFamily defines enum 'address_family'.
type AddressFamily uint32

const (
	ADDRESS_IP4 AddressFamily = 0
	ADDRESS_IP6 AddressFamily = 1
)

var (
	AddressFamily_name = map[uint32]string{
		0: "ADDRESS_IP4",
		1: "ADDRESS_IP6",
	}
	AddressFamily_value = map[string]uint32{
		"ADDRESS_IP4": 0,
		"ADDRESS_IP6": 1,
	}
)

func (x AddressFamily) String() string {
	s, ok := AddressFamily_name[uint32(x)]
	if ok {
		return s
	}
	return "AddressFamily(" + strconv.Itoa(int(x)) + ")"
}

// IPDscp defines enum 'ip_dscp'.
type IPDscp uint8

const (
	IP_API_DSCP_CS0  IPDscp = 0
	IP_API_DSCP_CS1  IPDscp = 8
	IP_API_DSCP_AF11 IPDscp = 10
	IP_API_DSCP_AF12 IPDscp = 12
	IP_API_DSCP_AF13 IPDscp = 14
	IP_API_DSCP_CS2  IPDscp = 16
	IP_API_DSCP_AF21 IPDscp = 18
	IP_API_DSCP_AF22 IPDscp = 20
	IP_API_DSCP_AF23 IPDscp = 22
	IP_API_DSCP_CS3  IPDscp = 24
	IP_API_DSCP_AF31 IPDscp = 26
	IP_API_DSCP_AF32 IPDscp = 28
	IP_API_DSCP_AF33 IPDscp = 30
	IP_API_DSCP_CS4  IPDscp = 32
	IP_API_DSCP_AF41 IPDscp = 34
	IP_API_DSCP_AF42 IPDscp = 36
	IP_API_DSCP_AF43 IPDscp = 38
	IP_API_DSCP_CS5  IPDscp = 40
	IP_API_DSCP_EF   IPDscp = 46
	IP_API_DSCP_CS6  IPDscp = 48
	IP_API_DSCP_CS7  IPDscp = 50
)

var (
	IPDscp_name = map[uint8]string{
		0:  "IP_API_DSCP_CS0",
		8:  "IP_API_DSCP_CS1",
		10: "IP_API_DSCP_AF11",
		12: "IP_API_DSCP_AF12",
		14: "IP_API_DSCP_AF13",
		16: "IP_API_DSCP_CS2",
		18: "IP_API_DSCP_AF21",
		20: "IP_API_DSCP_AF22",
		22: "IP_API_DSCP_AF23",
		24: "IP_API_DSCP_CS3",
		26: "IP_API_DSCP_AF31",
		28: "IP_API_DSCP_AF32",
		30: "IP_API_DSCP_AF33",
		32: "IP_API_DSCP_CS4",
		34: "IP_API_DSCP_AF41",
		36: "IP_API_DSCP_AF42",
		38: "IP_API_DSCP_AF43",
		40: "IP_API_DSCP_CS5",
		46: "IP_API_DSCP_EF",
		48: "IP_API_DSCP_CS6",
		50: "IP_API_DSCP_CS7",
	}
	IPDscp_value = map[string]uint8{
		"IP_API_DSCP_CS0":  0,
		"IP_API_DSCP_CS1":  8,
		"IP_API_DSCP_AF11": 10,
		"IP_API_DSCP_AF12": 12,
		"IP_API_DSCP_AF13": 14,
		"IP_API_DSCP_CS2":  16,
		"IP_API_DSCP_AF21": 18,
		"IP_API_DSCP_AF22": 20,
		"IP_API_DSCP_AF23": 22,
		"IP_API_DSCP_CS3":  24,
		"IP_API_DSCP_AF31": 26,
		"IP_API_DSCP_AF32": 28,
		"IP_API_DSCP_AF33": 30,
		"IP_API_DSCP_CS4":  32,
		"IP_API_DSCP_AF41": 34,
		"IP_API_DSCP_AF42": 36,
		"IP_API_DSCP_AF43": 38,
		"IP_API_DSCP_CS5":  40,
		"IP_API_DSCP_EF":   46,
		"IP_API_DSCP_CS6":  48,
		"IP_API_DSCP_CS7":  50,
	}
)

func (x IPDscp) String() string {
	s, ok := IPDscp_name[uint8(x)]
	if ok {
		return s
	}
	return "IPDscp(" + strconv.Itoa(int(x)) + ")"
}

// IPEcn defines enum 'ip_ecn'.
type IPEcn uint8

const (
	IP_API_ECN_NONE IPEcn = 0
	IP_API_ECN_ECT0 IPEcn = 1
	IP_API_ECN_ECT1 IPEcn = 2
	IP_API_ECN_CE   IPEcn = 3
)

var (
	IPEcn_name = map[uint8]string{
		0: "IP_API_ECN_NONE",
		1: "IP_API_ECN_ECT0",
		2: "IP_API_ECN_ECT1",
		3: "IP_API_ECN_CE",
	}
	IPEcn_value = map[string]uint8{
		"IP_API_ECN_NONE": 0,
		"IP_API_ECN_ECT0": 1,
		"IP_API_ECN_ECT1": 2,
		"IP_API_ECN_CE":   3,
	}
)

func (x IPEcn) String() string {
	s, ok := IPEcn_name[uint8(x)]
	if ok {
		return s
	}
	return "IPEcn(" + strconv.Itoa(int(x)) + ")"
}

// IPProto defines enum 'ip_proto'.
type IPProto uint32

const (
	IP_API_PROTO_HOPOPT   IPProto = 0
	IP_API_PROTO_ICMP     IPProto = 1
	IP_API_PROTO_IGMP     IPProto = 2
	IP_API_PROTO_TCP      IPProto = 6
	IP_API_PROTO_UDP      IPProto = 17
	IP_API_PROTO_GRE      IPProto = 47
	IP_API_PROTO_AH       IPProto = 50
	IP_API_PROTO_ESP      IPProto = 51
	IP_API_PROTO_EIGRP    IPProto = 88
	IP_API_PROTO_OSPF     IPProto = 89
	IP_API_PROTO_SCTP     IPProto = 132
	IP_API_PROTO_RESERVED IPProto = 255
)

var (
	IPProto_name = map[uint32]string{
		0:   "IP_API_PROTO_HOPOPT",
		1:   "IP_API_PROTO_ICMP",
		2:   "IP_API_PROTO_IGMP",
		6:   "IP_API_PROTO_TCP",
		17:  "IP_API_PROTO_UDP",
		47:  "IP_API_PROTO_GRE",
		50:  "IP_API_PROTO_AH",
		51:  "IP_API_PROTO_ESP",
		88:  "IP_API_PROTO_EIGRP",
		89:  "IP_API_PROTO_OSPF",
		132: "IP_API_PROTO_SCTP",
		255: "IP_API_PROTO_RESERVED",
	}
	IPProto_value = map[string]uint32{
		"IP_API_PROTO_HOPOPT":   0,
		"IP_API_PROTO_ICMP":     1,
		"IP_API_PROTO_IGMP":     2,
		"IP_API_PROTO_TCP":      6,
		"IP_API_PROTO_UDP":      17,
		"IP_API_PROTO_GRE":      47,
		"IP_API_PROTO_AH":       50,
		"IP_API_PROTO_ESP":      51,
		"IP_API_PROTO_EIGRP":    88,
		"IP_API_PROTO_OSPF":     89,
		"IP_API_PROTO_SCTP":     132,
		"IP_API_PROTO_RESERVED": 255,
	}
)

func (x IPProto) String() string {
	s, ok := IPProto_name[uint32(x)]
	if ok {
		return s
	}
	return "IPProto(" + strconv.Itoa(int(x)) + ")"
}

// AddressWithPrefix defines alias 'address_with_prefix'.
type AddressWithPrefix Prefix

func ParseAddressWithPrefix(s string) (AddressWithPrefix, error) {
	prefix, err := ParsePrefix(s)
	if err != nil {
		return AddressWithPrefix{}, err
	}
	return AddressWithPrefix(prefix), nil
}

func (x AddressWithPrefix) String() string {
	return Prefix(x).String()
}

func (x *AddressWithPrefix) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

func (x *AddressWithPrefix) UnmarshalText(text []byte) error {
	prefix, err := ParseAddressWithPrefix(string(text))
	if err != nil {
		return err
	}
	*x = prefix
	return nil
}

// IP4Address defines alias 'ip4_address'.
type IP4Address [4]uint8

func ParseIP4Address(s string) (IP4Address, error) {
	ip := net.ParseIP(s).To4()
	if ip == nil {
		return IP4Address{}, fmt.Errorf("invalid IP address: %s", s)
	}
	var ipaddr IP4Address
	copy(ipaddr[:], ip.To4())
	return ipaddr, nil
}

func (x IP4Address) ToIP() net.IP {
	return net.IP(x[:]).To4()
}

func (x IP4Address) String() string {
	return x.ToIP().String()
}

func (x *IP4Address) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

func (x *IP4Address) UnmarshalText(text []byte) error {
	ipaddr, err := ParseIP4Address(string(text))
	if err != nil {
		return err
	}
	*x = ipaddr
	return nil
}

// IP4AddressWithPrefix defines alias 'ip4_address_with_prefix'.
type IP4AddressWithPrefix IP4Prefix

// IP6Address defines alias 'ip6_address'.
type IP6Address [16]uint8

func ParseIP6Address(s string) (IP6Address, error) {
	ip := net.ParseIP(s).To16()
	if ip == nil {
		return IP6Address{}, fmt.Errorf("invalid IP address: %s", s)
	}
	var ipaddr IP6Address
	copy(ipaddr[:], ip.To16())
	return ipaddr, nil
}

func (x IP6Address) ToIP() net.IP {
	return net.IP(x[:]).To16()
}

func (x IP6Address) String() string {
	return x.ToIP().String()
}

func (x *IP6Address) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

func (x *IP6Address) UnmarshalText(text []byte) error {
	ipaddr, err := ParseIP6Address(string(text))
	if err != nil {
		return err
	}
	*x = ipaddr
	return nil
}

// IP6AddressWithPrefix defines alias 'ip6_address_with_prefix'.
type IP6AddressWithPrefix IP6Prefix

// Address defines type 'address'.
type Address struct {
	Af AddressFamily `binapi:"address_family,name=af" json:"af,omitempty"`
	Un AddressUnion  `binapi:"address_union,name=un" json:"un,omitempty"`
}

func ParseAddress(s string) (Address, error) {
	ip := net.ParseIP(s)
	if ip == nil {
		return Address{}, fmt.Errorf("invalid address: %s", s)
	}
	return AddressFromIP(ip), nil
}

func AddressFromIP(ip net.IP) Address {
	var addr Address
	if ip.To4() == nil {
		addr.Af = ADDRESS_IP6
		var ip6 IP6Address
		copy(ip6[:], ip.To16())
		addr.Un.SetIP6(ip6)
	} else {
		addr.Af = ADDRESS_IP4
		var ip4 IP4Address
		copy(ip4[:], ip.To4())
		addr.Un.SetIP4(ip4)
	}
	return addr
}

func (x Address) ToIP() net.IP {
	if x.Af == ADDRESS_IP6 {
		ip6 := x.Un.GetIP6()
		return net.IP(ip6[:]).To16()
	} else {
		ip4 := x.Un.GetIP4()
		return net.IP(ip4[:]).To4()
	}
}

func (x Address) String() string {
	return x.ToIP().String()
}

func (x *Address) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

func (x *Address) UnmarshalText(text []byte) error {
	addr, err := ParseAddress(string(text))
	if err != nil {
		return err
	}
	*x = addr
	return nil
}

// IP4Prefix defines type 'ip4_prefix'.
type IP4Prefix struct {
	Address IP4Address `binapi:"ip4_address,name=address" json:"address,omitempty"`
	Len     uint8      `binapi:"u8,name=len" json:"len,omitempty"`
}

func ParseIP4Prefix(s string) (prefix IP4Prefix, err error) {
	hasPrefix := strings.Contains(s, "/")
	if hasPrefix {
		ip, network, err := net.ParseCIDR(s)
		if err != nil {
			return IP4Prefix{}, fmt.Errorf("invalid IP %s: %s", s, err)
		}
		maskSize, _ := network.Mask.Size()
		prefix.Len = byte(maskSize)
		prefix.Address, err = ParseIP4Address(ip.String())
		if err != nil {
			return IP4Prefix{}, fmt.Errorf("invalid IP %s: %s", s, err)
		}
	} else {
		ip := net.ParseIP(s)
		defaultMaskSize, _ := net.CIDRMask(32, 32).Size()
		if ip.To4() == nil {
			defaultMaskSize, _ = net.CIDRMask(128, 128).Size()
		}
		prefix.Len = byte(defaultMaskSize)
		prefix.Address, err = ParseIP4Address(ip.String())
		if err != nil {
			return IP4Prefix{}, fmt.Errorf("invalid IP %s: %s", s, err)
		}
	}
	return prefix, nil
}

func (x IP4Prefix) ToIPNet() *net.IPNet {
	mask := net.CIDRMask(int(x.Len), 32)
	ipnet := &net.IPNet{IP: x.Address.ToIP(), Mask: mask}
	return ipnet
}

func (x IP4Prefix) String() string {
	ip := x.Address.String()
	return ip + "/" + strconv.Itoa(int(x.Len))
}

func (x *IP4Prefix) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

func (x *IP4Prefix) UnmarshalText(text []byte) error {
	prefix, err := ParseIP4Prefix(string(text))
	if err != nil {
		return err
	}
	*x = prefix
	return nil
}

// IP6Prefix defines type 'ip6_prefix'.
type IP6Prefix struct {
	Address IP6Address `binapi:"ip6_address,name=address" json:"address,omitempty"`
	Len     uint8      `binapi:"u8,name=len" json:"len,omitempty"`
}

func ParseIP6Prefix(s string) (prefix IP6Prefix, err error) {
	hasPrefix := strings.Contains(s, "/")
	if hasPrefix {
		ip, network, err := net.ParseCIDR(s)
		if err != nil {
			return IP6Prefix{}, fmt.Errorf("invalid IP %s: %s", s, err)
		}
		maskSize, _ := network.Mask.Size()
		prefix.Len = byte(maskSize)
		prefix.Address, err = ParseIP6Address(ip.String())
		if err != nil {
			return IP6Prefix{}, fmt.Errorf("invalid IP %s: %s", s, err)
		}
	} else {
		ip := net.ParseIP(s)
		defaultMaskSize, _ := net.CIDRMask(32, 32).Size()
		if ip.To4() == nil {
			defaultMaskSize, _ = net.CIDRMask(128, 128).Size()
		}
		prefix.Len = byte(defaultMaskSize)
		prefix.Address, err = ParseIP6Address(ip.String())
		if err != nil {
			return IP6Prefix{}, fmt.Errorf("invalid IP %s: %s", s, err)
		}
	}
	return prefix, nil
}

func (x IP6Prefix) ToIPNet() *net.IPNet {
	mask := net.CIDRMask(int(x.Len), 128)
	ipnet := &net.IPNet{IP: x.Address.ToIP(), Mask: mask}
	return ipnet
}

func (x IP6Prefix) String() string {
	ip := x.Address.String()
	return ip + "/" + strconv.Itoa(int(x.Len))
}

func (x *IP6Prefix) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

func (x *IP6Prefix) UnmarshalText(text []byte) error {
	prefix, err := ParseIP6Prefix(string(text))
	if err != nil {
		return err
	}
	*x = prefix
	return nil
}

// Mprefix defines type 'mprefix'.
type Mprefix struct {
	Af               AddressFamily `binapi:"address_family,name=af" json:"af,omitempty"`
	GrpAddressLength uint16        `binapi:"u16,name=grp_address_length" json:"grp_address_length,omitempty"`
	GrpAddress       AddressUnion  `binapi:"address_union,name=grp_address" json:"grp_address,omitempty"`
	SrcAddress       AddressUnion  `binapi:"address_union,name=src_address" json:"src_address,omitempty"`
}

// Prefix defines type 'prefix'.
type Prefix struct {
	Address Address `binapi:"address,name=address" json:"address,omitempty"`
	Len     uint8   `binapi:"u8,name=len" json:"len,omitempty"`
}

func ParsePrefix(ip string) (prefix Prefix, err error) {
	hasPrefix := strings.Contains(ip, "/")
	if hasPrefix {
		netIP, network, err := net.ParseCIDR(ip)
		if err != nil {
			return Prefix{}, fmt.Errorf("invalid IP %s: %s", ip, err)
		}
		maskSize, _ := network.Mask.Size()
		prefix.Len = byte(maskSize)
		prefix.Address, err = ParseAddress(netIP.String())
		if err != nil {
			return Prefix{}, fmt.Errorf("invalid IP %s: %s", ip, err)
		}
	} else {
		netIP := net.ParseIP(ip)
		defaultMaskSize, _ := net.CIDRMask(32, 32).Size()
		if netIP.To4() == nil {
			defaultMaskSize, _ = net.CIDRMask(128, 128).Size()
		}
		prefix.Len = byte(defaultMaskSize)
		prefix.Address, err = ParseAddress(netIP.String())
		if err != nil {
			return Prefix{}, fmt.Errorf("invalid IP %s: %s", ip, err)
		}
	}
	return prefix, nil
}

func (x Prefix) ToIPNet() *net.IPNet {
	var mask net.IPMask
	if x.Address.Af == ADDRESS_IP4 {
		mask = net.CIDRMask(int(x.Len), 32)
	} else {
		mask = net.CIDRMask(int(x.Len), 128)
	}
	ipnet := &net.IPNet{IP: x.Address.ToIP(), Mask: mask}
	return ipnet
}

func (x Prefix) String() string {
	ip := x.Address.String()
	return ip + "/" + strconv.Itoa(int(x.Len))
}

func (x *Prefix) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

func (x *Prefix) UnmarshalText(text []byte) error {
	prefix, err := ParsePrefix(string(text))
	if err != nil {
		return err
	}
	*x = prefix
	return nil
}

// PrefixMatcher defines type 'prefix_matcher'.
type PrefixMatcher struct {
	Le uint8 `binapi:"u8,name=le" json:"le,omitempty"`
	Ge uint8 `binapi:"u8,name=ge" json:"ge,omitempty"`
}

// UDPEncap defines type 'udp_encap'.
type UDPEncap struct {
	TableID uint32  `binapi:"u32,name=table_id" json:"table_id,omitempty"`
	SrcPort uint16  `binapi:"u16,name=src_port" json:"src_port,omitempty"`
	DstPort uint16  `binapi:"u16,name=dst_port" json:"dst_port,omitempty"`
	SrcIP   Address `binapi:"address,name=src_ip" json:"src_ip,omitempty"`
	DstIP   Address `binapi:"address,name=dst_ip" json:"dst_ip,omitempty"`
	ID      uint32  `binapi:"u32,name=id" json:"id,omitempty"`
}

// AddressUnion defines union 'address_union'.
type AddressUnion struct {
	// AddressUnion can be one of:
	// - IP4 *IP4Address
	// - IP6 *IP6Address
	XXX_UnionData [16]byte
}

func AddressUnionIP4(a IP4Address) (u AddressUnion) {
	u.SetIP4(a)
	return
}
func (u *AddressUnion) SetIP4(a IP4Address) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	buf.EncodeBytes(a[:], 4)
}
func (u *AddressUnion) GetIP4() (a IP4Address) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	copy(a[:], buf.DecodeBytes(4))
	return
}

func AddressUnionIP6(a IP6Address) (u AddressUnion) {
	u.SetIP6(a)
	return
}
func (u *AddressUnion) SetIP6(a IP6Address) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	buf.EncodeBytes(a[:], 16)
}
func (u *AddressUnion) GetIP6() (a IP6Address) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	copy(a[:], buf.DecodeBytes(16))
	return
}

// UDPEncapAdd defines message 'udp_encap_add'.
type UDPEncapAdd struct {
	UDPEncap UDPEncap `binapi:"udp_encap,name=udp_encap" json:"udp_encap,omitempty"`
}

func (m *UDPEncapAdd) Reset()               { *m = UDPEncapAdd{} }
func (*UDPEncapAdd) GetMessageName() string { return "udp_encap_add" }
func (*UDPEncapAdd) GetCrcString() string   { return "61d5fc48" }
func (*UDPEncapAdd) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *UDPEncapAdd) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4      // m.UDPEncap.TableID
	size += 2      // m.UDPEncap.SrcPort
	size += 2      // m.UDPEncap.DstPort
	size += 4      // m.UDPEncap.SrcIP.Af
	size += 1 * 16 // m.UDPEncap.SrcIP.Un
	size += 4      // m.UDPEncap.DstIP.Af
	size += 1 * 16 // m.UDPEncap.DstIP.Un
	size += 4      // m.UDPEncap.ID
	return size
}
func (m *UDPEncapAdd) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.UDPEncap.TableID)
	buf.EncodeUint16(m.UDPEncap.SrcPort)
	buf.EncodeUint16(m.UDPEncap.DstPort)
	buf.EncodeUint32(uint32(m.UDPEncap.SrcIP.Af))
	buf.EncodeBytes(m.UDPEncap.SrcIP.Un.XXX_UnionData[:], 16)
	buf.EncodeUint32(uint32(m.UDPEncap.DstIP.Af))
	buf.EncodeBytes(m.UDPEncap.DstIP.Un.XXX_UnionData[:], 16)
	buf.EncodeUint32(m.UDPEncap.ID)
	return buf.Bytes(), nil
}
func (m *UDPEncapAdd) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.UDPEncap.TableID = buf.DecodeUint32()
	m.UDPEncap.SrcPort = buf.DecodeUint16()
	m.UDPEncap.DstPort = buf.DecodeUint16()
	m.UDPEncap.SrcIP.Af = AddressFamily(buf.DecodeUint32())
	copy(m.UDPEncap.SrcIP.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.UDPEncap.DstIP.Af = AddressFamily(buf.DecodeUint32())
	copy(m.UDPEncap.DstIP.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.UDPEncap.ID = buf.DecodeUint32()
	return nil
}

// UDPEncapAddReply defines message 'udp_encap_add_reply'.
type UDPEncapAddReply struct {
	Retval int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	ID     uint32 `binapi:"u32,name=id" json:"id,omitempty"`
}

func (m *UDPEncapAddReply) Reset()               { *m = UDPEncapAddReply{} }
func (*UDPEncapAddReply) GetMessageName() string { return "udp_encap_add_reply" }
func (*UDPEncapAddReply) GetCrcString() string   { return "e2fc8294" }
func (*UDPEncapAddReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *UDPEncapAddReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.ID
	return size
}
func (m *UDPEncapAddReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(m.ID)
	return buf.Bytes(), nil
}
func (m *UDPEncapAddReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.ID = buf.DecodeUint32()
	return nil
}

// UDPEncapDel defines message 'udp_encap_del'.
type UDPEncapDel struct {
	ID uint32 `binapi:"u32,name=id" json:"id,omitempty"`
}

func (m *UDPEncapDel) Reset()               { *m = UDPEncapDel{} }
func (*UDPEncapDel) GetMessageName() string { return "udp_encap_del" }
func (*UDPEncapDel) GetCrcString() string   { return "3a91bde5" }
func (*UDPEncapDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *UDPEncapDel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.ID
	return size
}
func (m *UDPEncapDel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.ID)
	return buf.Bytes(), nil
}
func (m *UDPEncapDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.ID = buf.DecodeUint32()
	return nil
}

// UDPEncapDelReply defines message 'udp_encap_del_reply'.
type UDPEncapDelReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *UDPEncapDelReply) Reset()               { *m = UDPEncapDelReply{} }
func (*UDPEncapDelReply) GetMessageName() string { return "udp_encap_del_reply" }
func (*UDPEncapDelReply) GetCrcString() string   { return "e8d4e804" }
func (*UDPEncapDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *UDPEncapDelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *UDPEncapDelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *UDPEncapDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// UDPEncapDetails defines message 'udp_encap_details'.
type UDPEncapDetails struct {
	UDPEncap UDPEncap `binapi:"udp_encap,name=udp_encap" json:"udp_encap,omitempty"`
}

func (m *UDPEncapDetails) Reset()               { *m = UDPEncapDetails{} }
func (*UDPEncapDetails) GetMessageName() string { return "udp_encap_details" }
func (*UDPEncapDetails) GetCrcString() string   { return "87c82821" }
func (*UDPEncapDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *UDPEncapDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4      // m.UDPEncap.TableID
	size += 2      // m.UDPEncap.SrcPort
	size += 2      // m.UDPEncap.DstPort
	size += 4      // m.UDPEncap.SrcIP.Af
	size += 1 * 16 // m.UDPEncap.SrcIP.Un
	size += 4      // m.UDPEncap.DstIP.Af
	size += 1 * 16 // m.UDPEncap.DstIP.Un
	size += 4      // m.UDPEncap.ID
	return size
}
func (m *UDPEncapDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.UDPEncap.TableID)
	buf.EncodeUint16(m.UDPEncap.SrcPort)
	buf.EncodeUint16(m.UDPEncap.DstPort)
	buf.EncodeUint32(uint32(m.UDPEncap.SrcIP.Af))
	buf.EncodeBytes(m.UDPEncap.SrcIP.Un.XXX_UnionData[:], 16)
	buf.EncodeUint32(uint32(m.UDPEncap.DstIP.Af))
	buf.EncodeBytes(m.UDPEncap.DstIP.Un.XXX_UnionData[:], 16)
	buf.EncodeUint32(m.UDPEncap.ID)
	return buf.Bytes(), nil
}
func (m *UDPEncapDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.UDPEncap.TableID = buf.DecodeUint32()
	m.UDPEncap.SrcPort = buf.DecodeUint16()
	m.UDPEncap.DstPort = buf.DecodeUint16()
	m.UDPEncap.SrcIP.Af = AddressFamily(buf.DecodeUint32())
	copy(m.UDPEncap.SrcIP.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.UDPEncap.DstIP.Af = AddressFamily(buf.DecodeUint32())
	copy(m.UDPEncap.DstIP.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.UDPEncap.ID = buf.DecodeUint32()
	return nil
}

// UDPEncapDump defines message 'udp_encap_dump'.
type UDPEncapDump struct{}

func (m *UDPEncapDump) Reset()               { *m = UDPEncapDump{} }
func (*UDPEncapDump) GetMessageName() string { return "udp_encap_dump" }
func (*UDPEncapDump) GetCrcString() string   { return "51077d14" }
func (*UDPEncapDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *UDPEncapDump) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *UDPEncapDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *UDPEncapDump) Unmarshal(b []byte) error {
	return nil
}

func init() { file_udp_binapi_init() }
func file_udp_binapi_init() {
	api.RegisterMessage((*UDPEncapAdd)(nil), "udp_encap_add_61d5fc48")
	api.RegisterMessage((*UDPEncapAddReply)(nil), "udp_encap_add_reply_e2fc8294")
	api.RegisterMessage((*UDPEncapDel)(nil), "udp_encap_del_3a91bde5")
	api.RegisterMessage((*UDPEncapDelReply)(nil), "udp_encap_del_reply_e8d4e804")
	api.RegisterMessage((*UDPEncapDetails)(nil), "udp_encap_details_87c82821")
	api.RegisterMessage((*UDPEncapDump)(nil), "udp_encap_dump_51077d14")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*UDPEncapAdd)(nil),
		(*UDPEncapAddReply)(nil),
		(*UDPEncapDel)(nil),
		(*UDPEncapDelReply)(nil),
		(*UDPEncapDetails)(nil),
		(*UDPEncapDump)(nil),
	}
}
