package main

import (
	"flag"
	"fmt"
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

func (this *MyNode) OnNodeUpdate(nodeType int, id string, data []byte) {
	fmt.Println("OnNodeUpdate: nodeType =", nodeType, "id =", id, "data =", data)
}

func (this *MyNode) OnNodeJoin(nodeType int, id string, data []byte) {
	fmt.Println("OnNodeJoin: nodeType =", nodeType, "id =", id, "data =", data)
}

func (this *MyNode) OnNodeLeave(nodeType int, id string) {
	fmt.Println("OnNodeLeave: nodeType =", nodeType, "id =", id)
}

func (this *MyNode) GetPutData() (string, error) {
	return "", nil
}

func main() {

	hosts := ""
	flag.StringVar(&hosts, "hosts", "192.168.1.4:12379,192.168.1.4:22379,192.168.1.4:32379", "etcd hosts")
	nodeType := 0
	flag.IntVar(&nodeType, "nodeType", 1, "node type")
	watchNodeTypes := ""
	flag.StringVar(&watchNodeTypes, "watchNodeTypes", "1,2,3,4", "watch node type")
	putInterval := int64(0)
	flag.Int64Var(&putInterval, "putInterval", 1, "put interval")

	flag.Parse()

	node := NewMyNode()
	node.OpenByStr(hosts, nodeType, watchNodeTypes, putInterval)

	for {
		time.Sleep(time.Minute)
	}
}
