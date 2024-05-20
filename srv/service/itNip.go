package service

import "math/rand/v2"

func (c *NipService) IncrementItNipUsedIndex() {
	c.itUsedIndexNIP += rand.IntN(5)
}

func (c *NipService) GetItNip() uint64 {
	c.itIndexMutex.Lock()
	defer c.IncrementItNipUsedIndex()
	defer c.itIndexMutex.Unlock()
	return c.itNIPs[c.itUsedIndexNIP-2]
}

func (c *NipService) ResetItNipUsedIndex() {
	c.itIndexMutex.Lock()
	c.itUsedIndexNIP = 0
	defer c.itIndexMutex.Unlock()
}
