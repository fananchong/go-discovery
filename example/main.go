package main

import (
	"flag"
	"strconv"
	"strings"
	"time"

	"github.com/fananchong/go-discovery"
)

type MyNode struct {
	godiscovery.Node
}

func NewMyNode() *MyNode {
	return &MyNode{}
}

func (this *MyNode) OnNodeUpdate(nodeType int, id string, data string) {

}

func (this *MyNode) OnNodeJoin(nodeType int, id string, data string) {

}

func (this *MyNode) OnNodeLeave(nodeType int, id string) {

}

func (this *MyNode) GetPutData() string {
	return ""
}

func main() {

	hosts := ""
	flag.StringVar(&hosts, "hosts", "192.168.1.4:12379,192.168.1.4:22379,192.168.1.4:32379", "etcd hosts")
	nodeType := 0
	flag.IntVar(&nodeType, "nodeType", 1, "node type")
	watchNodeTypes := ""
	flag.StringVar(&watchNodeTypes, "watchNodeTypes", "1,2,3,4", "watch node type")
	putInterval := int64(0)
	flag.Int64Var(&putInterval, "putInterval", 5000, "put interval")

	flag.Parse()

	var wnt []int
	for _, val := range strings.Split(watchNodeTypes, ",") {
		v, _ := strconv.Atoi(val)
		wnt = append(wnt, v)
	}

	node := NewMyNode()
	node.SetLogger(NewGLog())
	node.Open(strings.Split(hosts, ","), nodeType, wnt, putInterval)

	for {
		time.Sleep(time.Minute)
	}
}
