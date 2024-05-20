package service

import (
	"fmt"
	"math/rand/v2"

	"github.com/nandanugg/HaloSusterTestCasesPSW3B2/entity"
)

func (c *NipService) AddItUsedAccount(usr entity.UsedUser) {
	fmt.Println("AddItUsedAccount: ", usr)
	c.itUsedAccountMutex.Lock()
	c.itUsedAccount = append(c.itUsedAccount, usr)
	c.itUsedAccountMutex.Unlock()
}

func (c *NipService) GetItUsedAccount() entity.UsedUser {
	c.itUsedAccountMutex.RLock()
	defer c.itUsedAccountMutex.RUnlock()
	choosenIndex := rand.IntN(len(c.itUsedAccount))
	return c.itUsedAccount[choosenIndex]
}

func (c *NipService) ResetItUsedAccount() {
	c.itUsedAccountMutex.Lock()
	defer c.itUsedAccountMutex.Unlock()
	c.itUsedAccount = []entity.UsedUser{}
}
