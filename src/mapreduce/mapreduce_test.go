package mapreduce

import (
	"testing"
	"fmt"

)

func Test(t *testing.T) {
	fmt.Println("Test mapreduce Start")
	var m Master
	var wk1, wk2, wk3 Worker
	//init the mapreduce struct
	m.Init()
	wk1.Init("127.0.0.1", 20001)
	wk2.Init("127.0.0.1", 20002)
	wk3.Init("127.0.0.1", 20003)
	//start mapreduce
	m.Start()
	wk1.Start()
	wk2.Start()
	wk3.Start()
	fmt.Println("Test mapreduce End")
}
