package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/funkygao/golib/debug"
	"github.com/funkygao/pastry"
	"github.com/funkygao/pretty"
)

var port int

func init() {
	debug.Debug = true

	flag.IntVar(&port, "p", 1090, "port")
	flag.Parse()
}

func format(v interface{}) interface{} {
	if false {
		return pretty.Formatter(v)
	}
	return v
}

func createNodeId() pastry.NodeID {
	hostname, _ := os.Hostname()
	name := hostname + fmt.Sprintf("%d %d", port, port)
	debug.Debugf("%s\n", name)
	id, e := pastry.NodeIDFromBytes([]byte(name))
	if e != nil {
		panic(e)
	}

	return id
}

type app struct{}

func (this *app) OnError(err error) {
	debug.Debugf("%s", err)
}

func (this *app) OnDeliver(msg pastry.Message) {
	// msg sent out to cluster
	debug.Debugf("%s", msg.String())
}

func (this *app) OnForward(msg *pastry.Message, nextId pastry.NodeID) bool {
	return false
}

func (this *app) OnNewLeaves(leafset []*pastry.Node) {
	debug.Debugf("%+v", leafset[0])
}

func (this *app) OnNodeJoin(node pastry.Node) {
	debug.Debugf("%+v", node)
}

func (this *app) OnNodeExit(node pastry.Node) {
	debug.Debugf("%+v", node)
}

func (this *app) OnHeartbeat(node pastry.Node) {
	debug.Debugf("%+v", node)
}

func main() {
	id := createNodeId()

	self := pastry.NewNode(id, "localhost", "12.43.34.11", "home", port)

	credentials := pastry.Passphrase("we are here")
	cluster := pastry.NewCluster(self, credentials)
	app := &app{}
	cluster.RegisterCallback(app)
	switch port {
	case 1091:
		cluster.SetColor("blue")
	case 1092:
		cluster.SetColor("green")
	}

	go startListener(cluster)

	go func() {
		ticker := time.NewTicker(time.Second * 10)
		for _ = range ticker.C {
			debug.Debugf("%s", cluster.LRM())
		}

	}()

	if port != 1090 {
		if err := cluster.Join("localhost", 1090); err != nil {
			panic(err)
		}

		/*

			key, _ := pastry.NodeIDFromBytes([]byte("adfasf"))
			node, err := cluster.Route(key)
			debug.Debugf("%#v %#v\n", node, err)*/
	}

	select {}

}

func startListener(cluster *pastry.Cluster) {
	if err := cluster.Listen(); err != nil {
		panic(err)
	}
	defer cluster.Stop()
}
