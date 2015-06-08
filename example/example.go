package main

import (
	"flag"
	"fmt"
	"os"

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
	debug.Debugf("%# v\n", format(id))

	node := pastry.NewNode(id, "localhost", "12.43.34.11", "home", port)
	debug.Debugf("%# v\n", format(node))

	credentials := pastry.Passphrase("we are here")
	cluster := pastry.NewCluster(node, credentials)
	debug.Debugf("%# v\n", format(cluster))

	if port != 1090 {
		if err := cluster.Join("localhost", 1090); err != nil {
			panic(err)
		}
	}

	if err := cluster.Listen(); err != nil {
		panic(err)
	}
	defer cluster.Stop()

}
