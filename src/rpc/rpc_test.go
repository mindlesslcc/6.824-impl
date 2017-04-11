package rpc

import (
	"errors"
	"testing"
)

var testServer *Server
var testClient *Client

func newTestServer() *Server {
	testServer = NewServer("tcp", "127.0.0.1:22222")
	go testServer.Start()

	return testServer
}

func newTestClient() *Client {
	testClient = NewClient("tcp", "127.0.0.1:22222", 10)

	return testClient
}

func test1(id int) (int, string, error) {
	return id * 10, "abc", nil
}

func test2(ids []int) ([]int, error) {
	if ids == nil || len(ids) == 0 {
		return nil, errors.New("nid ids")
	}

	if len(ids) >= 2 {
		return []int{}, nil
	}

	return []int{ids[0] * 10}, nil
}

func test3(id int) error {
	return errors.New("hello world")
}

func TestRpc(t *testing.T) {
	s := newTestServer()

	s.Register("rpc1", test1)
	s.Register("rpc2", test2)
	s.Register("rpc3", test3)

	c := newTestClient()

	var r1 func(int) (int, string, error)
	if err1 := c.MakeRpc("rpc1", &r1); err1 != nil {
		t.Fatal(err1)
	}

	var r2 func(ids []int) ([]int, error)
	c.MakeRpc("rpc2", &r2)

	var r3 func(int) error
	if err3 := c.MakeRpc("rpc3", &r3); err3 != nil {
		t.Fatal(err3)
	}
	//test rpc1
	a1, b1, e1 := r1(10)
	if e1 != nil {
		t.Fatal(e1)
	}

	if a1 != 100 || b1 != "abc" {
		t.Fatal(a1, b1)
	}
	//test rpc2
	a2, e2 := r2(nil)
	if e2 == nil {
		t.Fatal("must error")
	}

	a2, e2 = r2([]int{})
	if e2 == nil {
		t.Fatal("must error")
	}

	a2, e2 = r2([]int{1})
	if e2 != nil {
		t.Fatal(e2)
	} else if a2[0] != 10 {
		t.Fatal(a2[0])
	}

	a2, e2 = r2([]int{1, 2, 3})
	if e2 != nil {
		t.Fatal(e2)
	} else if len(a2) != 0 {
		t.Fatal("must 0")
	}
	//test rpc3
}


