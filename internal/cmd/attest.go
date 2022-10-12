package cmd

import (
	"github.com/spf13/cobra"
)

// AttestCmd returns the 'attest' command.
func AttestCmd(check func(error), f SigningFunc) *cobra.Command {
	var attPath string

	c := &cobra.Command{
		Use:   "attest",
		Short: "Generate a SLSA provenance attestation from existing files",
		Long:  `Generates and signs SLSA provenance from an existing set of files.`,

		Run: func(cmd *cobra.Command, args []string) {
			// TODO
		},
	}

	c.Flags().StringVarP(&attPath, "provenance", "p", "", "Path to write the signed provenance.")

	return c
}
