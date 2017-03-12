package mapreduce

import (
	"fmt"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	_"google.golang.org/grpc/reflection"

	pb "proto"
)

type Worker struct {
	masterIp	string
	masterPort	int
	ip string
	port int
}

// register to master
func (w *Worker)register(c pb.MasterClient) error {
	registerRequest := &pb.RegisterRequest{
		Ip			: w.ip,
		Port		: int32(w.port),
	}
	r, err := c.Register(context.Background(), registerRequest)
	if err != nil {
		fmt.Printf("could not register: %v", err)
		os.Exit(1)
	}
	fmt.Println(r)
	return err
}

func (w *Worker)Start() {
	//connect to master
	conn, err := grpc.Dial(MASTER_ADDRESS, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("did not connect: %v", err)
		os.Exit(1)
	}
	defer conn.Close()
	c := pb.NewMasterClient(conn)
	err = w.register(c)
	if err != nil {
		fmt.Printf("blockServer register error %s\n", err)
		os.Exit(1)
	}
}

func (w *Worker)Init(ip string, port int) {
	w.ip = ip
	w.port = port
}
