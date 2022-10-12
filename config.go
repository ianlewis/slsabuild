package main

import (
	"bytes"
	"os"
	"strings"
	"text/template"

	"github.com/ianlewis/slsabuild/internal/runner"
)

// Config defines the structure of the config file.
type Config struct {
	Commands []*runner.CommandStep
}

type tmplData struct {
	Env map[string]string
}

// Runner returns a command runner for the given config.
func (c Config) Runner() (*runner.CommandRunner, error) {
	var commands []*runner.CommandStep
	environ := getEnviron()
	for _, cmd := range c.Commands {
		var env []string
		for _, e := range cmd.Env {
			var buf bytes.Buffer
			tmpl, err := template.New("").Parse(e)
			if err != nil {
				return nil, err
			}
			if err := tmpl.Execute(&buf, tmplData{
				Env: environ,
			}); err != nil {
				return nil, err
			}

			env = append(env, buf.String())
		}

		commands = append(commands, &runner.CommandStep{
			Command:    cmd.Command,
			Env:        env,
			WorkingDir: cmd.WorkingDir,
		})
	}

	return &runner.CommandRunner{
		Steps: commands,
	}, nil
}

func getEnviron() map[string]string {
	env := make(map[string]string)
	for _, str := range os.Environ() {
		var k, v string
		parts := strings.SplitN(str, "=", 2)
		if len(parts) > 0 {
			k = parts[0]
			if len(parts) > 1 {
				v = parts[1]
			}
			env[k] = v
		}
	}
	return env
}
