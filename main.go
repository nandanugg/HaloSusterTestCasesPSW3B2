package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/nandanugg/HaloSusterTestCasesPSW3B2/entity/pb"
	"google.golang.org/grpc"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type Config struct {
	DBName     string `env:"DB_NAME"`
	DBPort     string `env:"DB_PORT"`
	DBHost     string `env:"DB_HOST"`
	DBUsername string `env:"DB_USERNAME"`
	DBPassword string `env:"DB_PASSWORD"`
	DBParams   string `env:"DB_PARAMS"`
}

// Define a struct to implement the NIPServiceServer interface
type NipServiceServer struct {
	pb.UnsafeNIPServiceServer
	itNIPs        []uint64
	nurseNIPs     []uint64
	itNipIndex    *ItNipIndex
	nurseNipIndex *NurseNipIndex
}

// Implement the GetItNip method
func (s *NipServiceServer) GetItNip(ctx context.Context, req *emptypb.Empty) (*pb.GetNipResponse, error) {
	s.itNipIndex.Add(1)
	i := s.itNipIndex.Value()
	return &pb.GetNipResponse{
		Nip: s.itNIPs[i],
	}, nil
}

// Implement the GetNurseNip method
func (s *NipServiceServer) GetNurseNip(ctx context.Context, req *emptypb.Empty) (*pb.GetNipResponse, error) {
	s.nurseNipIndex.Add(1)
	i := s.nurseNipIndex.Value()
	return &pb.GetNipResponse{
		Nip: s.nurseNIPs[i],
	}, nil
}

type ItNipIndex struct {
	sync.Mutex
	val uint64
}

func (c *ItNipIndex) Add(uint64) {
	c.Lock()
	c.val++
	c.Unlock()
}

func (c *ItNipIndex) Value() uint64 {
	return c.val
}

type NurseNipIndex struct {
	sync.Mutex
	val uint64
}

func (c *NurseNipIndex) Add(uint64) {
	c.Lock()
	c.val++
	c.Unlock()
}

func (c *NurseNipIndex) Value() uint64 {
	return c.val
}

func main() {
	fmt.Println("Starting server...")
	fmt.Println("Generating NIPs...")
	itNIPs := generateNIPs("615")
	nurseNIPs := generateNIPs("303")
	fmt.Println("Generated NIPs!")

	server := grpc.NewServer()
	srv := &NipServiceServer{
		itNIPs:    itNIPs,
		nurseNIPs: nurseNIPs,
	}
	pb.RegisterNIPServiceServer(server, srv)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	fmt.Println("Server is running on port 50051")
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
				for randomDigits := 0; randomDigits <= 9; randomDigits++ {
					randomPart := strconv.Itoa(randomDigits)
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
