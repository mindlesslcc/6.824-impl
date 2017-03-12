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

type WorkerInfo struct {
	ip string
	port int
	state int
}

type Master struct {
	ip string
	port int
	nWorker int
	workerDone chan bool
	workers []*WorkerInfo
}

var m Master

func (m *Master)Init() {
	m.ip = MASTER_IP
	m.port = MASTER_PORT
	m.nWorker = 0
	m.workers = make([]*WorkerInfo, 10)
	fmt.Println("Master Init OK!")
}

func (m *Master) existWorker(worker *WorkerInfo) bool {
	for _, value := range m.workers {
		if value.ip == worker.ip && value.port == worker.port {
			return true
		}
	}
	return false
}

func (m *Master) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	w := &WorkerInfo{
		ip		:in.Ip,
		port	:int(in.Port),
		state	:WORKER_STATE_CONNECTED,
	}
	//check exist or not
	if m.existWorker(w) {
		return &pb.RegisterResponse{Result:false}, nil
	} else {
		m.workers = append(m.workers, w)
		return &pb.RegisterResponse{Result:true}, nil
	}
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
