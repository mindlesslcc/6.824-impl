package mapreduce

import (
	"fmt"
	"os"
	"net"
	"sync"
	"time"

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
	wg sync.WaitGroup
}

var m Master

func (m *Master)Init(wg sync.WaitGroup) {
	m.ip = MASTER_IP
	m.port = MASTER_PORT
	m.nWorker = 0
	m.wg = wg
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
		fmt.Println("hehe")
		return &pb.RegisterResponse{Result:true}, nil
	}
}

func (m *Master) Reduce(ctx context.Context, in *pb.ReduceRequest) (*pb.ReduceResponse, error) {
	//wait for map report
	//print reduce result
	Reduce()
	m.wg.Add(-1)
	if len(m.workers) == 0 {
		//do reduce
		m.wg.Add(-1)
	}
	return &pb.ReduceResponse{}, nil
}

func masterPeriod() {
	fmt.Println("master period function")
}

func masterTimerFunc() {
	timer := time.NewTimer(time.Second * 10)
	for {
		select {
			case <- timer.C:
				masterPeriod()
				timer.Reset(time.Second * 10)
		}
	}
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
	go masterTimerFunc()
}

func (m *Master)Start() {
	fmt.Println("Master Start..")
	//start master rpc server
	go startMaster()
}
