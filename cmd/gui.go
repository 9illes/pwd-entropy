package cmd

import (
	"pwd-entropy/gui"

	"github.com/spf13/cobra"
)

var guiCmd = &cobra.Command{
	Use:     "gui",
	Short:   "Launch GUI",
	Aliases: []string{"ui"},
	Run: func(cmd *cobra.Command, args []string) {
		gui.Run()
	},
}
