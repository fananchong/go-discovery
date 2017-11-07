package godiscovery

import (
	"context"
	"runtime/debug"
	"strconv"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/golang/glog"
)

type IWatch interface {
	INode
	OnNodeUpdate(nodeType int, id string, data string)
	OnNodeJoin(nodeType int, id string, data string)
	OnNodeLeave(nodeType int, id string)
}

type Watch struct {
	Derived IWatch
	nodes   map[int]map[string]string
}

func (this *Watch) Open(derived IWatch, watchNodeTypes []int) {
	this.Derived = derived
	for _, nodeType := range watchNodeTypes {
		go this.watch(nodeType)
	}
}

func (this *Watch) watch(nodeType int) {
	glog.Infoln("start watch node, node type =", nodeType)
	defer func() {
		if err := recover(); err != nil {
			glog.Errorln("[异常] ", err, "\n", string(debug.Stack()))
		}
		this.Derived.Close()
	}()
	prefix := strconv.Itoa(nodeType) + "-"
	rch := this.Derived.GetClient().Watch(context.Background(), prefix, clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			glog.Infof("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			if ev.Type == mvccpb.PUT {

			} else if ev.Type == mvccpb.DELETE {

			} else {
				panic("unknow error!")
			}
		}
	}
}

func (this *Watch) Close() {
}
