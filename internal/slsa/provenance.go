package slsa

import (
	"fmt"

	intoto "github.com/in-toto/in-toto-golang/in_toto"
	slsacommon "github.com/in-toto/in-toto-golang/in_toto/slsa_provenance/common"
	slsa "github.com/in-toto/in-toto-golang/in_toto/slsa_provenance/v0.2"
	"sigs.k8s.io/release-utils/version"

	"github.com/ianlewis/slsabuild/internal/runner"
)

var builderID = "https://github.com/ianlewis/slsabuild@%s"
var buildType = "https://github.com/ianlewis/slsabuild/generic@v1"

type buildConfig struct {
	Version int                   `json:"version"`
	Steps   []*runner.CommandStep `json:"steps"`
}

// GenerateProvenance generates provenance for the given artifacts.
func GenerateProvenance(paths []string, steps []*runner.CommandStep) (*intoto.Statement, error) {
	var subjects []intoto.Subject
	for _, path := range paths {
		digest, err := Sha256(path)
		if err != nil {
			return nil, err
		}

		subjects = append(subjects, intoto.Subject{
			Name: path,
			Digest: slsacommon.DigestSet{
				"sha256": digest,
			},
		})
	}

	vInfo := version.GetVersionInfo()

	return &intoto.Statement{
		StatementHeader: intoto.StatementHeader{
			Type:          intoto.StatementInTotoV01,
			PredicateType: slsa.PredicateSLSAProvenance,
			Subject:       subjects,
		},
		Predicate: slsa.ProvenancePredicate{
			BuildType: buildType,
			Builder: slsacommon.ProvenanceBuilder{
				ID: fmt.Sprintf(builderID, vInfo.GitVersion),
			},
			// TODO: Set the invocation.
			// Invocation:  invocation,
			BuildConfig: buildConfig{
				Version: 1,
				Steps:   steps,
			},
			// TODO: Detect the materials from local git repo.
			// TODO: warn users when running in a dirty checkout.
			// TODO: warn users when running on a commit that isn't pushed.
			// Materials:   materials,
			// TODO: Set metadata.
			Metadata: &slsa.ProvenanceMetadata{},
		},
	}, nil
}
