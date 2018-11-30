package main

import (
	"fmt"

	"github.com/innoxchain/ixstorage/pkg/apps/ixclient"
)

func main() {
	fmt.Println(ixclient.Greet("world"))
}