package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
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
	name := fmt.Sprintf("%d %d", port, port) + "on host:" + hostname
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

	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("key to route> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		text = text[:len(text)-1] // strip EOL
		if text == "" {
			continue
		}

		fmt.Printf("Your key: %s\n", text)

		key, err := pastry.NodeIDFromBytes([]byte(text))
		if err != nil {
            fmt.Printf("shit: %s", err)
			continue
		}

		msg := cluster.NewMessage(byte(19), key, []byte("we are here"))
		err = cluster.Send(msg)
        if err != nil {
            println(err)
        }
	}

}

func startListener(cluster *pastry.Cluster) {
	if err := cluster.Listen(); err != nil {
		panic(err)
	}
	defer cluster.Stop()
}

func inData() []byte {
	fmt.Print("input key to send>")
	data, _ := ioutil.ReadAll(os.Stdin)
	return data
}
