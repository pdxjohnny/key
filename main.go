package main

import (
	"runtime"

	"github.com/spf13/cobra"

	"github.com/pdxjohnny/key/commands"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var rootCmd = &cobra.Command{Use: "key"}
	rootCmd.AddCommand(commands.Commands...)
	rootCmd.Execute()
}
