package service

import "math/rand/v2"

func (c *NipService) IncrementNurseNipUsedIndex() {
	c.nurseUsedIndexNIP += rand.IntN(5)
}

func (c *NipService) GetNurseNip() uint64 {
	c.nurseIndexMutex.Lock()
	defer c.IncrementNurseNipUsedIndex()
	defer c.nurseIndexMutex.Unlock()
	return c.nurseNIPs[c.nurseUsedIndexNIP]
}

func (c *NipService) ResetNurseNipUsedIndex() {
	c.nurseIndexMutex.Lock()
	c.nurseUsedIndexNIP = 0
	defer c.nurseIndexMutex.Unlock()
}
