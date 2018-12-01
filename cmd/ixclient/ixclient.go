package main

import (
	"fmt"

	"github.com/innoxchain/ixstorage/pkg/apps/ixclient"
	"github.com/innoxchain/ixstorage/build"
)

func main() {
	fmt.Println(ixclient.Greet("world"))
	fmt.Println("Version: ", version.Version)
}