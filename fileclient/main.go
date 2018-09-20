package main

import (
	"flag"
	"fmt"
	"github.com/quinor/treediff-go/diff"
	"gopkg.in/bblfsh/client-go.v2"
	"gopkg.in/bblfsh/sdk.v2/uast/nodes"
	"path/filepath"
)

var endpoint = flag.String("e", "localhost:9432", "endpoint of the babelfish server")
var filename = flag.String("f", "", "file to parse")

func getData(name string) (before, after nodes.Node) {
	client, err := bblfsh.NewClient(*endpoint)
	if err != nil {
		panic(err)
	}

	fn := func(name, suffix string) nodes.Node {
		pattern := fmt.Sprintf("/home/quinor/data/sourced/treediff/python-dataset/%v_%v_*.src", name, suffix)
		files, err := filepath.Glob(pattern)
		if err != nil || len(files) == 0 {
			panic(fmt.Sprintf("couldn't get a file from pattern %v", pattern))
		}
		fmt.Println(files[0])
		res, err := client.NewParseRequestV2().Language("python").ReadFile(files[0]).Do()
		if err != nil {
			panic(err)
		}
		node, err := res.Nodes()
		if err != nil {
			panic(err)
		}
		return node
	}
	return fn(name, "before"), fn(name, "after")
}

func main() {
	flag.Parse()
	if *filename == "" {
		fmt.Println("filename was not provided. Use the -f flag")
		return
	}
	src, dst := getData(*filename)
	changes := diff.Changes(src, dst)
	for _, ch := range changes {
		fmt.Printf("%T: %v\n", ch, ch)
	}
	fmt.Printf("changelist length is %v\n", len(changes))
	fmt.Printf("cost is %v\n", diff.Cost(src, dst))
	//https://github.com/bblfsh/sdk/blob/ccba81e734443cce192d2aef033ccc0e23751348/protocol/driver.go#L62
}
