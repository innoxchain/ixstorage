package main

import (
	"fmt"
	"os"

	"github.com/innoxchain/ixstorage/pkg/apps/ixclient"
	"github.com/innoxchain/ixstorage/build"

	"github.com/spf13/cobra"
)

func main() {

	ixclient.Greet("world")

	root := &cobra.Command{
		Use:   "ixclient cli",
		Short: "ixclient cli",
		Long:  "ixclient command line interface",
		Run: defaultStatus,
	}

	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Show detailed version information",
		Long:  "Show detailed information.",
		Run:   showVersion,
	}

	root.AddCommand(versionCmd)

	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func defaultStatus(cmd *cobra.Command, args []string) {
	fmt.Println("Run xclient version for detailed version information")
}

func showVersion(cmd *cobra.Command, args []string) {
	fmt.Println("ixclient")
	fmt.Println("\tVersion " + build.Version)
	fmt.Println("\tCommit " + build.Commit)
	fmt.Println("\tBranch " + build.Branch)
	fmt.Println("\tBuild Time " + build.BuildTime)
} 