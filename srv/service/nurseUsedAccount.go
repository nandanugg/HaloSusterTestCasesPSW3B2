package service

import "github.com/nandanugg/HaloSusterTestCasesPSW3B2/entity"

func (c *NipService) AddNurseUsedAccount(usr entity.UsedUser) {
	c.nurseUsedAccountMutex.Lock()
	c.nurseUsedAccount = append(c.nurseUsedAccount, usr)
	c.nurseUsedAccountMutex.Unlock()
}

func (c *NipService) GetNurseUsedAccount() entity.UsedUser {
	c.nurseUsedAccountMutex.RLocker()
	defer c.nurseUsedAccountMutex.RUnlock()
	return c.nurseUsedAccount[GenerateRandomNumber(0, len(c.nurseUsedAccount)-1)]
}

func (c *NipService) ResetNurseUsedAccount() {
	c.nurseUsedAccountMutex.Lock()
	defer c.nurseUsedAccountMutex.Unlock()
	c.nurseUsedAccount = []entity.UsedUser{}
}
