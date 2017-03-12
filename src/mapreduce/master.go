package mapreduce

import (
	"fmt"
	"os"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "proto"
)

type Master struct {
	ip string
	port int
	nWorker int
	workerDone chan bool
}

var m Master

func (m *Master)Init() {
	m.ip = MASTER_IP
	m.port = MASTER_PORT
	m.nWorker = 0
	fmt.Println("Master Init OK!")
}

func (m *Master) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return &pb.RegisterResponse{}, nil
}

func (m *Master) Reduce(ctx context.Context, in *pb.ReduceRequest) (*pb.ReduceResponse, error) {
	//wait for map report
	//print reduce result
	return &pb.ReduceResponse{}, nil
}

// startBlockServer start blockserver (rpc)
func startMaster() {
	lis, err := net.Listen("tcp", MASTER_ADDRESS)
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		os.Exit(1)
	}
	s := grpc.NewServer()
	pb.RegisterMasterServer(s, &m)
	//Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v", err)
		os.Exit(1)
	}
}

func (m *Master)Start() {
	fmt.Println("Master Start..")
	//start master rpc server
	go startMaster()
}
