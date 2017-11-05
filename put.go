package godiscovery

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/satori/go.uuid"
	"golang.org/x/net/context"
	"runtime/debug"
	"time"
)

type IPut interface {
	IEtcd
	GetPutData() string
}

type Put struct {
	Derived     IPut
	nodeType    int
	nodeId      string
	putInterval int64
	tick        *time.Ticker
}

func (this *Put) Open(derived IPut, nodeType int, putInterval int64) {
	this.Derived = derived
	this.nodeType = nodeType
	this.putInterval = putInterval
	this.nodeId = fmt.Sprintf("%d-%s", nodeType, uuid.NewV1().String())
	glog.Infoln("node id:", this.nodeId)
	go this.put(nodeType, putInterval)
}

func (this *Put) put(nodeType int, putInterval int64) {
	defer func() {
		if err := recover(); err != nil {
			glog.Errorln("[异常] ", err, "\n", string(debug.Stack()))
		}
		this.Derived.Close()
	}()

	this.tick = time.NewTicker(time.Duration(putInterval) * time.Millisecond)
	for {
		select {
		case <-this.tick.C:
			_, err := this.Derived.GetClient().Put(context.TODO(), this.nodeId, this.Derived.GetPutData())
			if err != nil {
				glog.Errorln(err)
			}
		}
	}
}

func (this *Put) Close() {
	if this.tick != nil {
		this.tick.Stop()
		this.tick = nil
	}
}
