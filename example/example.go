package main

import (
	"fmt"
	"os"

	"github.com/funkygao/golib/debug"
	"github.com/funkygao/pastry"
)

func init() {
	debug.Debug = true
}

func createNodeId() pastry.NodeID {
	hostname, _ := os.Hostname()
	id, e := pastry.NodeIDFromBytes([]byte(hostname + " test server on mac"))
	if e != nil {
		panic(e)
	}

	return id
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover:", r)
		}
	}()

	id := createNodeId()
	debug.Debugf("%#v\n", id)

	node := pastry.NewNode(id, "localhost", "12.43.34.11", "home", 1090)
	debug.Debugf("%#v\n", node)

	credentials := pastry.Passphrase("we are here")
	cluster := pastry.NewCluster(node, credentials)
	debug.Debugf("%#v\n", cluster)

	cluster.Listen()
	defer cluster.Stop()

}
