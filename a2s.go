package main

import (
	"fmt"

	"github.com/rumblefrog/go-a2s"
)

func main() {
	client, err := a2s.NewClient("squaller:8211")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer client.Close()

	info, err := client.QueryInfo() // QueryInfo, QueryPlayer, QueryRules

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v", info)
}
