// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package dhcp

import (
	"context"
	"fmt"
	"io"

	api "git.fd.io/govpp.git/api"
	vpe "git.fd.io/govpp.git/internal/testbinapi/binapi2001/vpe"
)

// RPCService defines RPC service dhcp.
type RPCService interface {
	DHCP6ClientsEnableDisable(ctx context.Context, in *DHCP6ClientsEnableDisable) (*DHCP6ClientsEnableDisableReply, error)
	DHCP6DuidLlSet(ctx context.Context, in *DHCP6DuidLlSet) (*DHCP6DuidLlSetReply, error)
	DHCP6PdSendClientMessage(ctx context.Context, in *DHCP6PdSendClientMessage) (*DHCP6PdSendClientMessageReply, error)
	DHCP6SendClientMessage(ctx context.Context, in *DHCP6SendClientMessage) (*DHCP6SendClientMessageReply, error)
	DHCPClientConfig(ctx context.Context, in *DHCPClientConfig) (*DHCPClientConfigReply, error)
	DHCPClientDump(ctx context.Context, in *DHCPClientDump) (RPCService_DHCPClientDumpClient, error)
	DHCPPluginControlPing(ctx context.Context, in *DHCPPluginControlPing) (*DHCPPluginControlPingReply, error)
	DHCPPluginGetVersion(ctx context.Context, in *DHCPPluginGetVersion) (*DHCPPluginGetVersionReply, error)
	DHCPProxyConfig(ctx context.Context, in *DHCPProxyConfig) (*DHCPProxyConfigReply, error)
	DHCPProxyDump(ctx context.Context, in *DHCPProxyDump) (RPCService_DHCPProxyDumpClient, error)
	DHCPProxySetVss(ctx context.Context, in *DHCPProxySetVss) (*DHCPProxySetVssReply, error)
	WantDHCP6PdReplyEvents(ctx context.Context, in *WantDHCP6PdReplyEvents) (*WantDHCP6PdReplyEventsReply, error)
	WantDHCP6ReplyEvents(ctx context.Context, in *WantDHCP6ReplyEvents) (*WantDHCP6ReplyEventsReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) DHCP6ClientsEnableDisable(ctx context.Context, in *DHCP6ClientsEnableDisable) (*DHCP6ClientsEnableDisableReply, error) {
	out := new(DHCP6ClientsEnableDisableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) DHCP6DuidLlSet(ctx context.Context, in *DHCP6DuidLlSet) (*DHCP6DuidLlSetReply, error) {
	out := new(DHCP6DuidLlSetReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) DHCP6PdSendClientMessage(ctx context.Context, in *DHCP6PdSendClientMessage) (*DHCP6PdSendClientMessageReply, error) {
	out := new(DHCP6PdSendClientMessageReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) DHCP6SendClientMessage(ctx context.Context, in *DHCP6SendClientMessage) (*DHCP6SendClientMessageReply, error) {
	out := new(DHCP6SendClientMessageReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) DHCPClientConfig(ctx context.Context, in *DHCPClientConfig) (*DHCPClientConfigReply, error) {
	out := new(DHCPClientConfigReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) DHCPClientDump(ctx context.Context, in *DHCPClientDump) (RPCService_DHCPClientDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_DHCPClientDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_DHCPClientDumpClient interface {
	Recv() (*DHCPClientDetails, error)
	api.Stream
}

type serviceClient_DHCPClientDumpClient struct {
	api.Stream
}

func (c *serviceClient_DHCPClientDumpClient) Recv() (*DHCPClientDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *DHCPClientDetails:
		return m, nil
	case *vpe.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) DHCPPluginControlPing(ctx context.Context, in *DHCPPluginControlPing) (*DHCPPluginControlPingReply, error) {
	out := new(DHCPPluginControlPingReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) DHCPPluginGetVersion(ctx context.Context, in *DHCPPluginGetVersion) (*DHCPPluginGetVersionReply, error) {
	out := new(DHCPPluginGetVersionReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DHCPProxyConfig(ctx context.Context, in *DHCPProxyConfig) (*DHCPProxyConfigReply, error) {
	out := new(DHCPProxyConfigReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) DHCPProxyDump(ctx context.Context, in *DHCPProxyDump) (RPCService_DHCPProxyDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_DHCPProxyDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_DHCPProxyDumpClient interface {
	Recv() (*DHCPProxyDetails, error)
	api.Stream
}

type serviceClient_DHCPProxyDumpClient struct {
	api.Stream
}

func (c *serviceClient_DHCPProxyDumpClient) Recv() (*DHCPProxyDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *DHCPProxyDetails:
		return m, nil
	case *vpe.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) DHCPProxySetVss(ctx context.Context, in *DHCPProxySetVss) (*DHCPProxySetVssReply, error) {
	out := new(DHCPProxySetVssReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) WantDHCP6PdReplyEvents(ctx context.Context, in *WantDHCP6PdReplyEvents) (*WantDHCP6PdReplyEventsReply, error) {
	out := new(WantDHCP6PdReplyEventsReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) WantDHCP6ReplyEvents(ctx context.Context, in *WantDHCP6ReplyEvents) (*WantDHCP6ReplyEventsReply, error) {
	out := new(WantDHCP6ReplyEventsReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}
