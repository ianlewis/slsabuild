package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/slsa-framework/slsa-github-generator/signing"
	"github.com/slsa-framework/slsa-github-generator/signing/sigstore"
	"github.com/spf13/cobra"

	"github.com/ianlewis/slsabuild/internal/cmd"
)

func checkExit(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func defaultSigningFunc(opts cmd.SigningOpts) (signing.Signer, signing.TransparencyLog) {
	return sigstore.NewDefaultFulcio(), sigstore.NewDefaultRekor()
}

func rootCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   "slsabuild",
		Short: "Generate SLSA provenance",
		Long: `Generate SLSA provenance.
For more information on SLSA, visit https://slsa.dev`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("expected command")
		},
	}
	c.AddCommand(cmd.VersionCmd())
	c.AddCommand(cmd.RunCmd(checkExit, defaultSigningFunc))
	c.AddCommand(cmd.AttestCmd(checkExit, defaultSigningFunc))
	return c
}

func main() {
	rootCmd().Execute()
}
