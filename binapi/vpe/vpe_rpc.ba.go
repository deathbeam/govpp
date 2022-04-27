// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package vpe

import (
	"context"
	"fmt"
	"io"

	api "git.fd.io/govpp.git/api"
	memclnt "git.fd.io/govpp.git/binapi/memclnt"
)

// RPCService defines RPC service vpe.
type RPCService interface {
	LogDump(ctx context.Context, in *LogDump) (RPCService_LogDumpClient, error)
	ShowVersion(ctx context.Context, in *ShowVersion) (*ShowVersionReply, error)
	ShowVpeSystemTime(ctx context.Context, in *ShowVpeSystemTime) (*ShowVpeSystemTimeReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) LogDump(ctx context.Context, in *LogDump) (RPCService_LogDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_LogDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_LogDumpClient interface {
	Recv() (*LogDetails, error)
	api.Stream
}

type serviceClient_LogDumpClient struct {
	api.Stream
}

func (c *serviceClient_LogDumpClient) Recv() (*LogDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *LogDetails:
		return m, nil
	case *memclnt.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) ShowVersion(ctx context.Context, in *ShowVersion) (*ShowVersionReply, error) {
	out := new(ShowVersionReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) ShowVpeSystemTime(ctx context.Context, in *ShowVpeSystemTime) (*ShowVpeSystemTimeReply, error) {
	out := new(ShowVpeSystemTimeReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}
