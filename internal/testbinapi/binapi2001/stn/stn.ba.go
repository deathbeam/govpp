// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.5.0-dev
//  VPP:              20.01
// source: .vppapi/plugins/stn.api.json

// Package stn contains generated bindings for API file stn.api.
//
// Contents:
//   6 aliases
//  10 enums
//   6 structs
//   1 union
//   4 messages
//
package stn

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
	APIFile    = "stn"
	APIVersion = "2.0.0"
	VersionCrc = 0x619d8f3
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

// IfStatusFlags defines enum 'if_status_flags'.
type IfStatusFlags uint32

const (
	IF_STATUS_API_FLAG_ADMIN_UP IfStatusFlags = 1
	IF_STATUS_API_FLAG_LINK_UP  IfStatusFlags = 2
)

var (
	IfStatusFlags_name = map[uint32]string{
		1: "IF_STATUS_API_FLAG_ADMIN_UP",
		2: "IF_STATUS_API_FLAG_LINK_UP",
	}
	IfStatusFlags_value = map[string]uint32{
		"IF_STATUS_API_FLAG_ADMIN_UP": 1,
		"IF_STATUS_API_FLAG_LINK_UP":  2,
	}
)

func (x IfStatusFlags) String() string {
	s, ok := IfStatusFlags_name[uint32(x)]
	if ok {
		return s
	}
	str := func(n uint32) string {
		s, ok := IfStatusFlags_name[uint32(n)]
		if ok {
			return s
		}
		return "IfStatusFlags(" + strconv.Itoa(int(n)) + ")"
	}
	for i := uint32(0); i <= 32; i++ {
		val := uint32(x)
		if val&(1<<i) != 0 {
			if s != "" {
				s += "|"
			}
			s += str(1 << i)
		}
	}
	if s == "" {
		return str(uint32(x))
	}
	return s
}

// IfType defines enum 'if_type'.
type IfType uint32

const (
	IF_API_TYPE_HARDWARE IfType = 1
	IF_API_TYPE_SUB      IfType = 2
	IF_API_TYPE_P2P      IfType = 3
	IF_API_TYPE_PIPE     IfType = 4
)

var (
	IfType_name = map[uint32]string{
		1: "IF_API_TYPE_HARDWARE",
		2: "IF_API_TYPE_SUB",
		3: "IF_API_TYPE_P2P",
		4: "IF_API_TYPE_PIPE",
	}
	IfType_value = map[string]uint32{
		"IF_API_TYPE_HARDWARE": 1,
		"IF_API_TYPE_SUB":      2,
		"IF_API_TYPE_P2P":      3,
		"IF_API_TYPE_PIPE":     4,
	}
)

func (x IfType) String() string {
	s, ok := IfType_name[uint32(x)]
	if ok {
		return s
	}
	return "IfType(" + strconv.Itoa(int(x)) + ")"
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

// LinkDuplex defines enum 'link_duplex'.
type LinkDuplex uint32

const (
	LINK_DUPLEX_API_UNKNOWN LinkDuplex = 0
	LINK_DUPLEX_API_HALF    LinkDuplex = 1
	LINK_DUPLEX_API_FULL    LinkDuplex = 2
)

var (
	LinkDuplex_name = map[uint32]string{
		0: "LINK_DUPLEX_API_UNKNOWN",
		1: "LINK_DUPLEX_API_HALF",
		2: "LINK_DUPLEX_API_FULL",
	}
	LinkDuplex_value = map[string]uint32{
		"LINK_DUPLEX_API_UNKNOWN": 0,
		"LINK_DUPLEX_API_HALF":    1,
		"LINK_DUPLEX_API_FULL":    2,
	}
)

func (x LinkDuplex) String() string {
	s, ok := LinkDuplex_name[uint32(x)]
	if ok {
		return s
	}
	return "LinkDuplex(" + strconv.Itoa(int(x)) + ")"
}

// MtuProto defines enum 'mtu_proto'.
type MtuProto uint32

const (
	MTU_PROTO_API_L3   MtuProto = 1
	MTU_PROTO_API_IP4  MtuProto = 2
	MTU_PROTO_API_IP6  MtuProto = 3
	MTU_PROTO_API_MPLS MtuProto = 4
	MTU_PROTO_API_N    MtuProto = 5
)

var (
	MtuProto_name = map[uint32]string{
		1: "MTU_PROTO_API_L3",
		2: "MTU_PROTO_API_IP4",
		3: "MTU_PROTO_API_IP6",
		4: "MTU_PROTO_API_MPLS",
		5: "MTU_PROTO_API_N",
	}
	MtuProto_value = map[string]uint32{
		"MTU_PROTO_API_L3":   1,
		"MTU_PROTO_API_IP4":  2,
		"MTU_PROTO_API_IP6":  3,
		"MTU_PROTO_API_MPLS": 4,
		"MTU_PROTO_API_N":    5,
	}
)

func (x MtuProto) String() string {
	s, ok := MtuProto_name[uint32(x)]
	if ok {
		return s
	}
	return "MtuProto(" + strconv.Itoa(int(x)) + ")"
}

// RxMode defines enum 'rx_mode'.
type RxMode uint32

const (
	RX_MODE_API_UNKNOWN   RxMode = 0
	RX_MODE_API_POLLING   RxMode = 1
	RX_MODE_API_INTERRUPT RxMode = 2
	RX_MODE_API_ADAPTIVE  RxMode = 3
	RX_MODE_API_DEFAULT   RxMode = 4
)

var (
	RxMode_name = map[uint32]string{
		0: "RX_MODE_API_UNKNOWN",
		1: "RX_MODE_API_POLLING",
		2: "RX_MODE_API_INTERRUPT",
		3: "RX_MODE_API_ADAPTIVE",
		4: "RX_MODE_API_DEFAULT",
	}
	RxMode_value = map[string]uint32{
		"RX_MODE_API_UNKNOWN":   0,
		"RX_MODE_API_POLLING":   1,
		"RX_MODE_API_INTERRUPT": 2,
		"RX_MODE_API_ADAPTIVE":  3,
		"RX_MODE_API_DEFAULT":   4,
	}
)

func (x RxMode) String() string {
	s, ok := RxMode_name[uint32(x)]
	if ok {
		return s
	}
	return "RxMode(" + strconv.Itoa(int(x)) + ")"
}

// SubIfFlags defines enum 'sub_if_flags'.
type SubIfFlags uint32

const (
	SUB_IF_API_FLAG_NO_TAGS           SubIfFlags = 1
	SUB_IF_API_FLAG_ONE_TAG           SubIfFlags = 2
	SUB_IF_API_FLAG_TWO_TAGS          SubIfFlags = 4
	SUB_IF_API_FLAG_DOT1AD            SubIfFlags = 8
	SUB_IF_API_FLAG_EXACT_MATCH       SubIfFlags = 16
	SUB_IF_API_FLAG_DEFAULT           SubIfFlags = 32
	SUB_IF_API_FLAG_OUTER_VLAN_ID_ANY SubIfFlags = 64
	SUB_IF_API_FLAG_INNER_VLAN_ID_ANY SubIfFlags = 128
	SUB_IF_API_FLAG_MASK_VNET         SubIfFlags = 254
	SUB_IF_API_FLAG_DOT1AH            SubIfFlags = 256
)

var (
	SubIfFlags_name = map[uint32]string{
		1:   "SUB_IF_API_FLAG_NO_TAGS",
		2:   "SUB_IF_API_FLAG_ONE_TAG",
		4:   "SUB_IF_API_FLAG_TWO_TAGS",
		8:   "SUB_IF_API_FLAG_DOT1AD",
		16:  "SUB_IF_API_FLAG_EXACT_MATCH",
		32:  "SUB_IF_API_FLAG_DEFAULT",
		64:  "SUB_IF_API_FLAG_OUTER_VLAN_ID_ANY",
		128: "SUB_IF_API_FLAG_INNER_VLAN_ID_ANY",
		254: "SUB_IF_API_FLAG_MASK_VNET",
		256: "SUB_IF_API_FLAG_DOT1AH",
	}
	SubIfFlags_value = map[string]uint32{
		"SUB_IF_API_FLAG_NO_TAGS":           1,
		"SUB_IF_API_FLAG_ONE_TAG":           2,
		"SUB_IF_API_FLAG_TWO_TAGS":          4,
		"SUB_IF_API_FLAG_DOT1AD":            8,
		"SUB_IF_API_FLAG_EXACT_MATCH":       16,
		"SUB_IF_API_FLAG_DEFAULT":           32,
		"SUB_IF_API_FLAG_OUTER_VLAN_ID_ANY": 64,
		"SUB_IF_API_FLAG_INNER_VLAN_ID_ANY": 128,
		"SUB_IF_API_FLAG_MASK_VNET":         254,
		"SUB_IF_API_FLAG_DOT1AH":            256,
	}
)

func (x SubIfFlags) String() string {
	s, ok := SubIfFlags_name[uint32(x)]
	if ok {
		return s
	}
	str := func(n uint32) string {
		s, ok := SubIfFlags_name[uint32(n)]
		if ok {
			return s
		}
		return "SubIfFlags(" + strconv.Itoa(int(n)) + ")"
	}
	for i := uint32(0); i <= 32; i++ {
		val := uint32(x)
		if val&(1<<i) != 0 {
			if s != "" {
				s += "|"
			}
			s += str(1 << i)
		}
	}
	if s == "" {
		return str(uint32(x))
	}
	return s
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

// InterfaceIndex defines alias 'interface_index'.
type InterfaceIndex uint32

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

// StnAddDelRule defines message 'stn_add_del_rule'.
type StnAddDelRule struct {
	IPAddress Address        `binapi:"address,name=ip_address" json:"ip_address,omitempty"`
	SwIfIndex InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	IsAdd     bool           `binapi:"bool,name=is_add" json:"is_add,omitempty"`
}

func (m *StnAddDelRule) Reset()               { *m = StnAddDelRule{} }
func (*StnAddDelRule) GetMessageName() string { return "stn_add_del_rule" }
func (*StnAddDelRule) GetCrcString() string   { return "53f751e6" }
func (*StnAddDelRule) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *StnAddDelRule) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4      // m.IPAddress.Af
	size += 1 * 16 // m.IPAddress.Un
	size += 4      // m.SwIfIndex
	size += 1      // m.IsAdd
	return size
}
func (m *StnAddDelRule) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.IPAddress.Af))
	buf.EncodeBytes(m.IPAddress.Un.XXX_UnionData[:], 16)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeBool(m.IsAdd)
	return buf.Bytes(), nil
}
func (m *StnAddDelRule) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IPAddress.Af = AddressFamily(buf.DecodeUint32())
	copy(m.IPAddress.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.SwIfIndex = InterfaceIndex(buf.DecodeUint32())
	m.IsAdd = buf.DecodeBool()
	return nil
}

// StnAddDelRuleReply defines message 'stn_add_del_rule_reply'.
type StnAddDelRuleReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *StnAddDelRuleReply) Reset()               { *m = StnAddDelRuleReply{} }
func (*StnAddDelRuleReply) GetMessageName() string { return "stn_add_del_rule_reply" }
func (*StnAddDelRuleReply) GetCrcString() string   { return "e8d4e804" }
func (*StnAddDelRuleReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *StnAddDelRuleReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *StnAddDelRuleReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *StnAddDelRuleReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// StnRulesDetails defines message 'stn_rules_details'.
type StnRulesDetails struct {
	IPAddress Address        `binapi:"address,name=ip_address" json:"ip_address,omitempty"`
	SwIfIndex InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *StnRulesDetails) Reset()               { *m = StnRulesDetails{} }
func (*StnRulesDetails) GetMessageName() string { return "stn_rules_details" }
func (*StnRulesDetails) GetCrcString() string   { return "b0f6606c" }
func (*StnRulesDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *StnRulesDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4      // m.IPAddress.Af
	size += 1 * 16 // m.IPAddress.Un
	size += 4      // m.SwIfIndex
	return size
}
func (m *StnRulesDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.IPAddress.Af))
	buf.EncodeBytes(m.IPAddress.Un.XXX_UnionData[:], 16)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *StnRulesDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IPAddress.Af = AddressFamily(buf.DecodeUint32())
	copy(m.IPAddress.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.SwIfIndex = InterfaceIndex(buf.DecodeUint32())
	return nil
}

// StnRulesDump defines message 'stn_rules_dump'.
type StnRulesDump struct{}

func (m *StnRulesDump) Reset()               { *m = StnRulesDump{} }
func (*StnRulesDump) GetMessageName() string { return "stn_rules_dump" }
func (*StnRulesDump) GetCrcString() string   { return "51077d14" }
func (*StnRulesDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *StnRulesDump) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *StnRulesDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *StnRulesDump) Unmarshal(b []byte) error {
	return nil
}

func init() { file_stn_binapi_init() }
func file_stn_binapi_init() {
	api.RegisterMessage((*StnAddDelRule)(nil), "stn_add_del_rule_53f751e6")
	api.RegisterMessage((*StnAddDelRuleReply)(nil), "stn_add_del_rule_reply_e8d4e804")
	api.RegisterMessage((*StnRulesDetails)(nil), "stn_rules_details_b0f6606c")
	api.RegisterMessage((*StnRulesDump)(nil), "stn_rules_dump_51077d14")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*StnAddDelRule)(nil),
		(*StnAddDelRuleReply)(nil),
		(*StnRulesDetails)(nil),
		(*StnRulesDump)(nil),
	}
}
