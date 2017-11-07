package godiscovery

import (
	"sync"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/golang/glog"
)

type INode interface {
	GetClient() *clientv3.Client
	Close()
}

type Node struct {
	Watch
	Put
	client         *clientv3.Client
	hosts          []string
	nodeType       int
	watchNodeTypes []int
	putInterval    int64
	mutex          sync.Mutex
}

func (this *Node) Open(hosts []string, nodeType int, watchNodeTypes []int, putInterval int64) {
	this.hosts = hosts
	this.nodeType = nodeType
	this.watchNodeTypes = watchNodeTypes
	this.putInterval = putInterval
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   hosts,
		DialTimeout: 5 * time.Second,
	})
	this.client = cli
	if err != nil {
		glog.Errorln(err)
		go this.reopen()
		return
	}
	if len(watchNodeTypes) != 0 {
		this.Watch.Open(this, watchNodeTypes)
	}
	if nodeType != 0 {
		this.Put.Open(this, nodeType, putInterval)
	}
}

func (this *Node) SetLogger(log clientv3.Logger) {
	clientv3.SetLogger(log)
}

func (this *Node) Close() {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	if this.client != nil {
		this.client.Close()
		this.client = nil
	}
	this.Put.Close()
	this.Watch.Close()
}

func (this *Node) reopen() {
	glog.Infoln("reopen after 5 sec.")
	t := time.NewTimer(5 * time.Second)
	select {
	case <-t.C:
		this.Open(this.hosts, this.nodeType, this.watchNodeTypes, this.putInterval)
	}
}

func (this *Node) Id() string {
	return this.Put.nodeId
}

func (this *Node) GetClient() *clientv3.Client {
	return this.client
}

// 子类可以根据需要重载下面的方法
func (this *Node) OnNodeUpdate(nodeType int, id string, data string) {

}

func (this *Node) OnNodeJoin(nodeType int, id string, data string) {

}

func (this *Node) OnNodeLeave(nodeType int, id string) {

}

func (this *Node) GetPutData() string {
	return ""
}
