package main

import (
	"flag"
	"fmt"
	"time"

	discovery "github.com/fananchong/go-discovery/serverlist"
)

type MyNode struct {
	*discovery.Node
}

func NewMyNode() *MyNode {
	this := &MyNode{}
	this.Node = discovery.NewDefaultNode(this)
	return this
}

func (this *MyNode) OnNodeJoin(nodeIP string, nodeType int, id uint32, data []byte) {
	this.Node.OnNodeJoin(nodeIP, nodeType, id, data)

	fmt.Println("print current all node:")
	for t := 1; t <= 4; t++ {
		if lst, ok := this.Servers.GetAll(t); ok {
			for _, info := range lst {
				fmt.Printf("    id:%d type:%d addr:%s\n", info.GetId(), info.GetType(), info.GetExternalIp())
			}
		}
	}
}

func (this *MyNode) OnNodeLeave(nodeType int, id uint32) {
	this.Node.OnNodeLeave(nodeType, id)

	fmt.Println("print current all node:")
	for t := 1; t <= 4; t++ {
		if lst, ok := this.Servers.GetAll(t); ok {
			for _, info := range lst {
				fmt.Printf("    id:%d type:%d addr:%s\n", info.GetId(), info.GetType(), info.GetExternalIp())
			}
		}
	}
}

func main() {

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

	node := NewMyNode()
	node.OpenByStr(hosts, whatsmyip, nodeType, watchNodeTypes, putInterval)

	for {
		time.Sleep(10 * time.Second)
	}

	node.Close()
}
