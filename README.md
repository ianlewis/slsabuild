# slsabuild

Build projects and generate SLSA provenance

# Getting Started

Create a config file `slsabuild.yaml` like the following:

```yaml
# artifacts are outputs of the build.
artifacts: ["slsabuild"]

# commands are commands that are run to generate the artifacts.
commands:
  - command:
      [
        "go",
        "build",
        "-ldflags=-X sigs.k8s.io/release-utils/version.gitVersion={{ .Env.VERSION }}",
        ".",
      ]
    env:
      - "HOME={{ .Env.HOME }}"
      - "GOPATH={{ .Env.GOPATH }}"
```

Run `slsabuild` to build the project. This will run the commands to generate the
artifact(s), generate provenance and sign it using keyless signing with Fulcio.
The resulting certificate is uploaded to the public Rekor transparency log.

```
slsabuild
```

You should now see your artifact and `<artifact>.intoto.jsonl`.
