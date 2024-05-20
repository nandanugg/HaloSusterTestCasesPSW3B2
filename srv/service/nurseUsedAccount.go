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
	c.nurseUsedAccountCount++
	c.nurseUsedAccountMutex.Unlock()
}

func (c *NipService) GetNurseUsedAccount() entity.UsedUser {
	c.nurseUsedAccountMutex.Lock()
	defer c.nurseUsedAccountMutex.Unlock()
	choosenIndex := rand.IntN(c.nurseUsedAccountCount / 2)
	return c.nurseUsedAccount[choosenIndex]
}

func (c *NipService) ResetNurseUsedAccount() {
	c.nurseUsedAccountMutex.Lock()
	defer c.nurseUsedAccountMutex.Unlock()
	c.nurseUsedAccount = []entity.UsedUser{}
}
