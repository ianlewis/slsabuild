package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/slsa-framework/slsa-github-generator/signing"
	"github.com/spf13/cobra"

	"github.com/ianlewis/slsabuild/internal/config"
	"github.com/ianlewis/slsabuild/internal/slsa"
)

// SigningOpts defines options for signers.
type SigningOpts struct {
	Keyless bool
	KeyPath string
}

// SigningFunc returns a signer and transparency log.
type SigningFunc func(SigningOpts) (signing.Signer, signing.TransparencyLog, error)

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

			ctx := context.Background()

			steps, err := r.Run(ctx)
			check(err)

			// Generate SLSA provenance.
			p, err := slsa.GenerateProvenance(cfg.Artifacts, steps)
			check(err)

			// Sign SLSA provenance.
			// TODO: Support local keys.
			signer, tlog, err := f(SigningOpts{
				Keyless: true,
			})

			// Sign provenance. Need auth flow.
			att, err := signer.Sign(ctx, p)
			check(err)

			_, err = tlog.Upload(ctx, att)
			check(err)

			if attPath == "" {
				attPath = "multiple.intoto.jsonl"
				if len(cfg.Artifacts) == 1 {
					attPath = fmt.Sprintf("%s.intoto.jsonl", filepath.Base(cfg.Artifacts[0]))
				}
			}

			// Write signed attestation.
			check(os.WriteFile(attPath, att.Bytes(), 0600))
		},
	}

	c.Flags().StringVarP(&configPath, "config", "c", "slsabuild.yaml", "Path to the config file.")
	c.Flags().StringVarP(&attPath, "provenance", "p", "", "Path to write the signed provenance.")

	return c
}
