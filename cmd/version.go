package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Returns the version information",
	Run: func(cmd *cobra.Command, args []string) {
		_version()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func _version() {
	fmt.Printf("build date: %s\n"+
		"git version: %s commit: %s\n"+
		"go version: %s os: %s arch: %s compiler: %s\n",
		date, version, commit, runtime.Version(), runtime.GOARCH, runtime.GOOS, runtime.Compiler)
}
