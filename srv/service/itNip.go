package service

func (c *NipService) IncrementItNipUsedIndex() {
	c.itIndexMutex.Lock()
	c.itUsedIndexNIP += uint64(GenerateRandomNumber(1, 100))
	c.itIndexMutex.Unlock()
}

func (c *NipService) GetItNip() uint64 {
	c.itIndexMutex.Lock()
	defer c.itIndexMutex.Unlock()
	return c.itNIPs[c.itUsedIndexNIP]
}

func (c *NipService) ResetItNipUsedIndex() {
	c.itIndexMutex.Lock()
	c.itUsedIndexNIP = 0
	defer c.itIndexMutex.Unlock()
}
