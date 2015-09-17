package commands

import (
	"github.com/spf13/cobra"

	"github.com/pdxjohnny/key/decrypt"
	"github.com/pdxjohnny/key/encrypt"
	"github.com/pdxjohnny/key/gen"
)

var Commands = []*cobra.Command{
	&cobra.Command{
		Use:   "gen",
		Short: "Generate keys",
		Run: func(cmd *cobra.Command, args []string) {
			ConfigBindFlags(cmd)
			gen.Run()
		},
	},
	&cobra.Command{
		Use:   "encrypt",
		Short: "Encrypt messages",
		Run: func(cmd *cobra.Command, args []string) {
			ConfigBindFlags(cmd)
			encrypt.Run()
		},
	},
	&cobra.Command{
		Use:   "decrypt",
		Short: "Decrypt messages",
		Run: func(cmd *cobra.Command, args []string) {
			ConfigBindFlags(cmd)
			decrypt.Run()
		},
	},
}

func init() {
	ConfigDefaults(Commands...)
}
