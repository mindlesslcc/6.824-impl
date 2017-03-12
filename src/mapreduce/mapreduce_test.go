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
	wk1.Init()
	wk2.Init()
	wk3.Init()
	//start mapreduce
	m.Start()
	wk1.Start()
	wk2.Start()
	wk3.Start()
	fmt.Println("Test mapreduce End")
}
