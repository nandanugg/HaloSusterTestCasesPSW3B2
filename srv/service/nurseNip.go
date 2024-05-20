package service

func (c *NipService) IncrementNurseNipUsedIndex() {
	c.nurseIndexMutex.Lock()
	c.nurseUsedIndexNIP += uint64(GenerateRandomNumber(1, 100))
	c.nurseIndexMutex.Unlock()
}

func (c *NipService) GetNurseNip() uint64 {
	c.nurseIndexMutex.Lock()
	defer c.nurseIndexMutex.Unlock()
	return c.nurseNIPs[c.nurseUsedIndexNIP]
}

func (c *NipService) ResetNurseNipUsedIndex() {
	c.nurseIndexMutex.Lock()
	c.nurseUsedIndexNIP = 0
	defer c.nurseIndexMutex.Unlock()
}
