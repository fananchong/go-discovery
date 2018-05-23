package main

import (
	"flag"
	"fmt"
	"sync/atomic"
	"time"

	discovery "github.com/fananchong/go-discovery/serverlist"
)

var gCounter int64 = 1
var gAppID int64 = 0

type MyNode struct {
	discovery.Node
}

func NewMyNode() *MyNode {
	this := &MyNode{}
	this.Node.InitPolicy(discovery.RoundRobin)
	this.Node.Init(this)
	return this
}

func (this *MyNode) NewNodeId() (uint32, error) {
	temp := atomic.AddInt64(&gCounter, 1)
	return uint32(gAppID*10000 + temp), nil
}

func (this *MyNode) OnNodeJoin(nodeIP string, nodeType int, id uint32, data []byte) {
	this.Node.OnNodeJoin(nodeIP, nodeType, id, data)

	if showmsg != 0 {
		fmt.Println("[join] current node count:", this.Servers.Count(nodeType))
	}
}

func (this *MyNode) OnNodeLeave(nodeType int, id uint32) {
	this.Node.OnNodeLeave(nodeType, id)

	if showmsg != 0 {
		fmt.Println("[leave] current node count:", this.Servers.Count(nodeType))
	}
}

var showmsg = 0

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
	clientCount := 0
	flag.IntVar(&clientCount, "clientCount", 3, "client count")
	flag.IntVar(&showmsg, "showmsg", 0, "showmsg")
	flag.Int64Var(&gAppID, "appid", 1, "appid")
	flag.Parse()

	for i := 0; i < clientCount; i++ {
		go func() {
			node := NewMyNode()
			node.DisablePort()
			node.OpenByStr(hosts, whatsmyip, nodeType, watchNodeTypes, putInterval)
		}()
		time.Sleep(1 * time.Millisecond)
	}
	for {
		time.Sleep(10 * time.Second)
	}
}
