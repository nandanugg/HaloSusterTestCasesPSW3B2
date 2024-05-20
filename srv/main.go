package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/nandanugg/HaloSusterTestCasesPSW3B2/entity/pb"
	"github.com/nandanugg/HaloSusterTestCasesPSW3B2/handler"
	"github.com/nandanugg/HaloSusterTestCasesPSW3B2/service"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting server...")
	fmt.Println("Generating NIPs...")
	itNIPs := generateNIPs("615")
	nurseNIPs := generateNIPs("303")
	fmt.Println("Generated NIPs!")

	server := grpc.NewServer()
	var nipMutex sync.Mutex
	var itMutex sync.Mutex
	var usedItMutex sync.RWMutex
	var usedNurseMutex sync.RWMutex

	handler := handler.NewRequestHandler(service.NewNipService(
		itNIPs,
		nurseNIPs,
		&itMutex,
		&nipMutex,
		&usedItMutex,
		&usedNurseMutex,
	))

	pb.RegisterNIPServiceServer(server, handler)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func generateNIPs(prefix string) []uint64 {
	currentYear := time.Now().Year()
	res := []uint64{}

	for year := 2000; year <= currentYear; year++ {
		for month := 1; month <= 12; month++ {
			yearStr := fmt.Sprintf("%d", year)
			monthStr := fmt.Sprintf("%02d", month)
			for gender := 1; gender <= 2; gender++ {
				genderStr := strconv.Itoa(gender)
				for randomDigits := 0; randomDigits <= 9999; randomDigits++ {
					randomPart := strconv.Itoa(randomDigits)
					if len(randomPart) < 3 {
						randomPart = fmt.Sprintf("%03s", randomPart)
					}
					nipStr := prefix + genderStr + yearStr + monthStr + randomPart
					nip, err := strconv.ParseUint(nipStr, 10, 64)
					handleErr(err)
					res = append(res, nip)
				}
			}
		}
	}
	return res
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
