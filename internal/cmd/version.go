package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"sigs.k8s.io/release-utils/version"
)

// VersionCmd returns a command that prints version info and exists.
func VersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print version info and exit",
		Run: func(cmd *cobra.Command, args []string) {
			vInfo := version.GetVersionInfo()
			fmt.Println((&vInfo).String())
		},
	}
}
