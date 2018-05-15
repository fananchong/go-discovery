package main

import (
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"

	godiscovery "github.com/fananchong/go-discovery"
)

type MyNode struct {
	godiscovery.Node
}

func NewMyNode() *MyNode {
	this := &MyNode{}
	this.Node.Init(this)
	return this
}

func (this *MyNode) OnNodeUpdate(nodeIP string, nodeType int, id string, data []byte) {
	fmt.Println("OnNodeUpdate: nodeIP =", nodeIP, "nodeType =", nodeType, "id =", id, "data =", data)
}

func (this *MyNode) OnNodeJoin(nodeIP string, nodeType int, id string, data []byte) {
	fmt.Println("OnNodeJoin: nodeIP =", nodeIP, "nodeType =", nodeType, "id =", id, "data =", data)
}

func (this *MyNode) OnNodeLeave(nodeType int, id string) {
	fmt.Println("OnNodeLeave: nodeType =", nodeType, "id =", id)
}

func (this *MyNode) GetPutData() (string, error) {
	return string([]byte{1, 2, 3, 4}), nil
}

func main() {

	pprof_port := 0
	flag.IntVar(&pprof_port, "pprofPort", 3000, "pprof port")

	hosts := ""
	flag.StringVar(&hosts, "hosts", "101.132.47.70:12379,101.132.47.70:22379,101.132.47.70:32379", "etcd hosts")
	whatsmyip := ""
	flag.StringVar(&whatsmyip, "whatsmyip", "101.132.47.70:3000", "whatsmyip host")
	nodeType := 0
	flag.IntVar(&nodeType, "nodeType", 1, "node type")
	watchNodeTypes := ""
	flag.StringVar(&watchNodeTypes, "watchNodeTypes", "1,2,3,4", "watch node type")
	putInterval := int64(0)
	flag.Int64Var(&putInterval, "putInterval", 1, "put interval")

	flag.Parse()

	go http.ListenAndServe(fmt.Sprintf(":%d", pprof_port), nil)

	for {
		node := NewMyNode()
		node.OpenByStr(hosts, whatsmyip, nodeType, watchNodeTypes, putInterval)

		time.Sleep(10 * time.Second)
		fmt.Println("node close .....................")
		node.Close()
	}
}
