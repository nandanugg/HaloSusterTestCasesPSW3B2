package service

import (
	"fmt"
	"math/rand/v2"

	"github.com/nandanugg/HaloSusterTestCasesPSW3B2/entity"
)

func (c *NipService) AddNurseUsedAccount(usr entity.UsedUser) {
	fmt.Println("AddNurseUsedAccount: ", usr)
	c.nurseUsedAccountMutex.Lock()
	c.nurseUsedAccount = append(c.nurseUsedAccount, usr)
	c.nurseUsedAccountMutex.Unlock()
}

func (c *NipService) GetNurseUsedAccount() entity.UsedUser {
	c.nurseUsedAccountMutex.RLock()
	defer c.nurseUsedAccountMutex.RUnlock()
	choosenIndex := rand.IntN(len(c.itUsedAccount))
	return c.nurseUsedAccount[choosenIndex]
}

func (c *NipService) ResetNurseUsedAccount() {
	c.nurseUsedAccountMutex.Lock()
	defer c.nurseUsedAccountMutex.Unlock()
	c.nurseUsedAccount = []entity.UsedUser{}
}
