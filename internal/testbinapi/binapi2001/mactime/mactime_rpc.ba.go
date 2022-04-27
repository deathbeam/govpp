// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package mactime

import (
	"context"
	"fmt"
	"io"

	api "git.fd.io/govpp.git/api"
	vpe "git.fd.io/govpp.git/internal/testbinapi/binapi2001/vpe"
)

// RPCService defines RPC service mactime.
type RPCService interface {
	MactimeAddDelRange(ctx context.Context, in *MactimeAddDelRange) (*MactimeAddDelRangeReply, error)
	MactimeDump(ctx context.Context, in *MactimeDump) (RPCService_MactimeDumpClient, error)
	MactimeEnableDisable(ctx context.Context, in *MactimeEnableDisable) (*MactimeEnableDisableReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) MactimeAddDelRange(ctx context.Context, in *MactimeAddDelRange) (*MactimeAddDelRangeReply, error) {
	out := new(MactimeAddDelRangeReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) MactimeDump(ctx context.Context, in *MactimeDump) (RPCService_MactimeDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_MactimeDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_MactimeDumpClient interface {
	Recv() (*MactimeDetails, error)
	api.Stream
}

type serviceClient_MactimeDumpClient struct {
	api.Stream
}

func (c *serviceClient_MactimeDumpClient) Recv() (*MactimeDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *MactimeDetails:
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

func (c *serviceClient) MactimeEnableDisable(ctx context.Context, in *MactimeEnableDisable) (*MactimeEnableDisableReply, error) {
	out := new(MactimeEnableDisableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}
