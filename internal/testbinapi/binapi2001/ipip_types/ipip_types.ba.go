// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.4.0-dev
//  VPP:              20.01
// source: .vppapi/core/ipip_types.api.json

// Package ipip_types contains generated bindings for API file ipip_types.api.
//
// Contents:
//   1 enum
//
package ipip_types

import (
	api "git.fd.io/govpp.git/api"
	"strconv"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

// IpipTunnelFlags defines enum 'ipip_tunnel_flags'.
type IpipTunnelFlags uint8

const (
	IPIP_TUNNEL_API_FLAG_NONE            IpipTunnelFlags = 0
	IPIP_TUNNEL_API_FLAG_ENCAP_COPY_DF   IpipTunnelFlags = 1
	IPIP_TUNNEL_API_FLAG_ENCAP_SET_DF    IpipTunnelFlags = 2
	IPIP_TUNNEL_API_FLAG_ENCAP_COPY_DSCP IpipTunnelFlags = 4
	IPIP_TUNNEL_API_FLAG_ENCAP_COPY_ECN  IpipTunnelFlags = 8
	IPIP_TUNNEL_API_FLAG_DECAP_COPY_ECN  IpipTunnelFlags = 16
)

var (
	IpipTunnelFlags_name = map[uint8]string{
		0:  "IPIP_TUNNEL_API_FLAG_NONE",
		1:  "IPIP_TUNNEL_API_FLAG_ENCAP_COPY_DF",
		2:  "IPIP_TUNNEL_API_FLAG_ENCAP_SET_DF",
		4:  "IPIP_TUNNEL_API_FLAG_ENCAP_COPY_DSCP",
		8:  "IPIP_TUNNEL_API_FLAG_ENCAP_COPY_ECN",
		16: "IPIP_TUNNEL_API_FLAG_DECAP_COPY_ECN",
	}
	IpipTunnelFlags_value = map[string]uint8{
		"IPIP_TUNNEL_API_FLAG_NONE":            0,
		"IPIP_TUNNEL_API_FLAG_ENCAP_COPY_DF":   1,
		"IPIP_TUNNEL_API_FLAG_ENCAP_SET_DF":    2,
		"IPIP_TUNNEL_API_FLAG_ENCAP_COPY_DSCP": 4,
		"IPIP_TUNNEL_API_FLAG_ENCAP_COPY_ECN":  8,
		"IPIP_TUNNEL_API_FLAG_DECAP_COPY_ECN":  16,
	}
)

func (x IpipTunnelFlags) String() string {
	s, ok := IpipTunnelFlags_name[uint8(x)]
	if ok {
		return s
	}
	str := func(n uint8) string {
		s, ok := IpipTunnelFlags_name[uint8(n)]
		if ok {
			return s
		}
		return "IpipTunnelFlags(" + strconv.Itoa(int(n)) + ")"
	}
	for i := uint8(0); i <= 8; i++ {
		val := uint8(x)
		if val&(1<<i) != 0 {
			if s != "" {
				s += "|"
			}
			s += str(1 << i)
		}
	}
	if s == "" {
		return str(uint8(x))
	}
	return s
}
