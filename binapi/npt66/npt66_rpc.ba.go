// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package npt66

import (
	"context"

	api "go.fd.io/govpp/api"
)

// RPCService defines RPC service npt66.
type RPCService interface {
	Npt66BindingAddDel(ctx context.Context, in *Npt66BindingAddDel) (*Npt66BindingAddDelReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) Npt66BindingAddDel(ctx context.Context, in *Npt66BindingAddDel) (*Npt66BindingAddDelReply, error) {
	out := new(Npt66BindingAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}