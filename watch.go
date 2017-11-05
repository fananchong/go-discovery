package godiscovery

import (
	"github.com/coreos/etcd/clientv3"
	"github.com/golang/glog"
	"golang.org/x/net/context"
	"runtime/debug"
	"strconv"
)

type IWatch interface {
	IEtcd
	OnNodeUpdate(nodeType int, id string, data string)
	OnNodeJoin(nodeType int, id string, data string)
	OnNodeLeave(nodeType int, id string)
}

type Watch struct {
	Derived        IWatch
	watchNodeTypes []int
	nodes          map[int]map[string]string
}

func (this *Watch) Open(derived IWatch, watchNodeTypes []int) {
	this.Derived = derived
	this.watchNodeTypes = watchNodeTypes
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
		}
	}
}

func (this *Watch) Close() {
}
