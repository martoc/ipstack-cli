package core

import (
	"context"

	"github.com/martoc/ipstack-cli/adapter"
)

type GetCoordinatesCommandBuilder struct {
	IP        string
	AccessKey string
	ctx       context.Context //nolint:containedctx
}

func NewGetCoordinatesCommandBuilder() *GetCoordinatesCommandBuilder {
	return &GetCoordinatesCommandBuilder{}
}

func (b *GetCoordinatesCommandBuilder) SetIP(ip string) *GetCoordinatesCommandBuilder {
	b.IP = ip

	return b
}

func (b *GetCoordinatesCommandBuilder) SetAccessKey(accessKey string) *GetCoordinatesCommandBuilder {
	b.AccessKey = accessKey

	return b
}

func (b *GetCoordinatesCommandBuilder) SetContext(ctx context.Context) *GetCoordinatesCommandBuilder {
	b.ctx = ctx

	return b
}

func (b *GetCoordinatesCommandBuilder) Build() Command {
	return &GetCoordinatesCommandImpl{
		IP:        b.IP,
		AccessKey: b.AccessKey,
		Ctx:       b.ctx,
	}
}

type GetCoordinatesCommandImpl struct {
	Command
	IP        string
	AccessKey string
	Ctx       context.Context //nolint:containedctx
}

func (c *GetCoordinatesCommandImpl) Execute() (string, error) {
	return adapter.NewIPStackServiceAdapterBuilder().Build().GetCoordinates(c.Ctx, c.IP, c.AccessKey)
}
