package service

import (
	"context"

	"github.com/nandanugg/HaloSusterTestCasesPSW3B2/entity"
)

func (c *NipService) AddItUsedAccount(ctx context.Context, usr *entity.UsedUser) {
	c.repo.PostUsedITAccount(ctx, usr)
}

func (c *NipService) GetItUsedAccount(ctx context.Context) *entity.UsedUser {
	return c.repo.GetUsedITAccount(ctx)
}
