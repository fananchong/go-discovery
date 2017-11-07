package godiscovery

import (
	"context"
	"fmt"
	"runtime/debug"
	"time"

	"github.com/golang/glog"
	uuid "github.com/satori/go.uuid"
)

type IPut interface {
	INode
	GetPutData() string
}

type Put struct {
	Derived  IPut
	nodeId   string
	tick     *time.Ticker
	chanStop chan int
}

func (this *Put) Open(derived IPut, nodeType int, putInterval int64) {
	this.Derived = derived
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
			if this.Derived.GetClient() == nil {
				return
			}
			_, err := this.Derived.GetClient().Put(context.TODO(), this.nodeId, this.Derived.GetPutData())
			if err != nil {
				glog.Errorln(err)
			}
		case <-this.chanStop:
			return
		}
	}
}

func (this *Put) Close() {
	if this.tick != nil {
		this.tick.Stop()
		this.tick = nil
	}
	this.chanStop <- 1
}
