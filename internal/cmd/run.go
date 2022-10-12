package cmd

import (
	"context"

	"github.com/ianlewis/slsabuild/internal/config"
	"github.com/slsa-framework/slsa-github-generator/signing"
	"github.com/spf13/cobra"
)

// SigningOpts defines options for signers.
type SigningOpts struct {
	Keyless bool
	KeyPath string
}

// SigningFunc returns a signer and transparency log.
type SigningFunc func(SigningOpts) (signing.Signer, signing.TransparencyLog)

// RunCmd returns the 'run' command.
func RunCmd(check func(error), f SigningFunc) *cobra.Command {
	var attPath string
	var configPath string

	c := &cobra.Command{
		Use:   "run",
		Short: "Build artifact(s) and generate SLSA provenance",
		Long:  `Builds a set of artifacts, generates and signs SLSA provenance.`,

		Run: func(cmd *cobra.Command, args []string) {
			cfg, err := config.ReadConfig(configPath)
			check(err)

			r, err := cfg.Runner()
			check(err)

			_, err = r.Run(context.Background())
			check(err)

			// TODO: Generate SLSA provenance.
			// TODO: Sign SLSA provenance.
		},
	}

	c.Flags().StringVarP(&configPath, "config", "c", "slsabuild.yaml", "Path to the config file.")
	c.Flags().StringVarP(&attPath, "provenance", "p", "", "Path to write the signed provenance.")

	return c
}
