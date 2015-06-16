package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/funkygao/pastry"
	"github.com/funkygao/peter"
)

var (
	opts struct {
		m string
	}

	topic = peter.Topic(strings.Repeat("topic", 10))
)

func init() {
	flag.StringVar(&opts.m, "m", "x", "run mode")
	flag.Parse()
}

func dieIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func pub() {
	id, _ := pastry.NodeIDFromBytes([]byte(strings.Repeat("pub", 20)))
	p := peter.New(id, "127.0.0.1", "127.0.0.1", "home", 1234)
	go p.Listen()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("pub> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		text = text[:len(text)-1] // strip EOL
		if text == "" {
			continue
		}
		if text == "exit" {
			break
		}

		if err := p.Broadcast(topic, []byte(text)); err != nil {
			fmt.Println(err)
		}

	}

}

func sub() {
	id, _ := pastry.NodeIDFromBytes([]byte(strings.Repeat("sub", 20)))
	p := peter.New(id, "127.0.0.1", "127.0.0.1", "home", 1235)
	go p.Listen()
	defer p.Stop()

	err := p.Join("127.0.0.1", 1234)
	dieIfError(err)
	err = p.Subscribe(topic)
	dieIfError(err)
	time.Sleep(time.Hour)

}

func main() {
	if opts.m == "pub" {
		pub()
	} else if opts.m == "sub" {
		sub()
	} else {
		flag.Usage()
	}

}
