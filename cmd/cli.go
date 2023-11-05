package cmd

import (
	"bufio"
	"fmt"
	"os"
	"pwd-entropy/pkg/entropy"
	"strings"

	"github.com/spf13/cobra"
)

var cliCmd = &cobra.Command{
	Use:     "cli",
	Short:   "Calculate password strength",
	Aliases: []string{"c"},
	Run: func(cmd *cobra.Command, args []string) {

		pwd := Prompt("Enter password :")
		fmt.Printf("Entropy : %d bits\n", int(entropy.GetEntropy(entropy.NewNamespaceToEntropy(entropy.DefaultRules), pwd)))
	},
}

func Prompt(label string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, label+" ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}
