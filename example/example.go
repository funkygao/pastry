package main

import (
    "fmt"
    "os"
    "secondbit.org/pastry"
)

func createNodeId() pastry.NodeID {
    hostname, _ := os.Hostname()
    id, e := pastry.NodeIDFromBytes([]byte(hostname + " test server on mac"))
    if e != nil {
        panic(e)
    }

    return id
}

func debug(v ...interface{}) {
    for _, x := range v {
        fmt.Printf("%#v\n", x)
    }
}

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("recover:", r)
        }
    }()

    id := createNodeId()
    debug(id)

    node := pastry.NewNode(id, "localhost", "12.43.34.11", "home", 1090)
    debug(node)

    credentials := pastry.Passphrase("we are here")
    cluster := pastry.NewCluster(node, credentials)
    debug(cluster)

    cluster.Listen()
    defer cluster.Stop()

}
