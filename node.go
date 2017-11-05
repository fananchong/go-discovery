package godiscovery

import (
	"github.com/coreos/etcd/clientv3"
	"github.com/golang/glog"
	"time"
)

type IEtcd interface {
	GetClient() *clientv3.Client
	Close()
}

type Node struct {
	Watch
	Put
	client *clientv3.Client
}

func (this *Node) Open(hosts []string, nodeType int, watchNodeTypes []int, putInterval int64) {
	clientv3.SetLogger(NewLogger())
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   hosts,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		glog.Errorln(err)
	}
	this.client = cli
	if len(watchNodeTypes) != 0 {
		this.Watch.Open(this, watchNodeTypes)
	}
	if nodeType != 0 {
		this.Put.Open(this, nodeType, putInterval)
	}
}

func (this *Node) Close() {
	if this.client != nil {
		this.client.Close()
		this.client = nil
	}
	this.Put.Close()
	this.Watch.Close()
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
