package service

import (
	"sync"

	"github.com/nandanugg/HaloSusterTestCasesPSW3B2/entity"
	"github.com/nandanugg/HaloSusterTestCasesPSW3B2/entity/pb"
)

// Define a struct to implement the NIPServiceServer interface
type NipService struct {
	pb.UnsafeNIPServiceServer
	itNIPs                []uint64
	nurseNIPs             []uint64
	itUsedIndexNIP        int
	nurseUsedIndexNIP     int
	itUsedAccountCount    int
	nurseUsedAccountCount int
	itUsedAccount         []entity.UsedUser
	nurseUsedAccount      []entity.UsedUser
	itIndexMutex          *sync.Mutex
	nurseIndexMutex       *sync.Mutex
	itUsedAccountMutex    *sync.RWMutex
	nurseUsedAccountMutex *sync.RWMutex
}

// Create a new NipServiceServer
func NewNipService(
	itNIPs, nurseNIPs []uint64,
	itIndexMutex *sync.Mutex,
	nurseIndexMutex *sync.Mutex,
	itUsedAccountMutex *sync.RWMutex,
	nurseUsedAccountMutex *sync.RWMutex,
) *NipService {
	return &NipService{
		itNIPs:                itNIPs,
		nurseNIPs:             nurseNIPs,
		itIndexMutex:          itIndexMutex,
		nurseIndexMutex:       nurseIndexMutex,
		itUsedAccountMutex:    itUsedAccountMutex,
		nurseUsedAccountMutex: nurseUsedAccountMutex,
	}
}
