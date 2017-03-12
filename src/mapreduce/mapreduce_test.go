package mapreduce

import (
	"testing"
	"fmt"
	"sync"
)

func Test(t *testing.T) {
	fmt.Println("Test mapreduce Start")
	var wg sync.WaitGroup
	var m Master
	var wk1, wk2, wk3 Worker
	//init the mapreduce struct
	m.Init(wg)
	wk1.Init("127.0.0.1", 20001)
	wk2.Init("127.0.0.1", 20002)
	wk3.Init("127.0.0.1", 20003)
	//start mapreduce
	m.Start()
	wg.Add(3)
	go wk1.Start()
	go wk2.Start()
	go wk3.Start()
	wg.Done()
	fmt.Println("Test mapreduce End")
}
