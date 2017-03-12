package mapreduce

import (
	"fmt"
	_"net"

	_"golang.org/x/net/context"
	_"google.golang.org/grpc"
	_"google.golang.org/grpc/reflection"

	_ "proto"
)

type Master struct {
	ip string
	port int
}

func (m *Master)Init() {
	m.ip = MASTER_IP
	m.port = MASTER_PORT
	fmt.Println("Master Init OK!")
}

func (m *Master)Start() {
	fmt.Println("Master Start..")
	//start master rpc server
	
	//wait for map report
	//print reduce result
}
