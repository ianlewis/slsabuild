package config

import (
	"bytes"
	"os"
	"strings"
	"text/template"

	"github.com/ianlewis/slsabuild/internal/runner"
	"gopkg.in/yaml.v2"
)

// ReadConfig reads the config from the given path.
func ReadConfig(path string) (*Config, error) {
	// Read the config
	cf, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	d := yaml.NewDecoder(cf)
	var cfg Config
	if err := d.Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

// Config defines the structure of the config file.
type Config struct {
	Commands []*runner.CommandStep
}

type tmplData struct {
	Env map[string]string
}

// Runner returns a command runner for the given config.
func (c *Config) Runner() (*runner.CommandRunner, error) {
	var commands []*runner.CommandStep
	environ := getEnviron()
	for _, cmd := range c.Commands {
		var args []string
		for _, arg := range cmd.Command {
			argVar, err := resolveTmpl(arg, tmplData{
				Env: environ,
			})
			if err != nil {
				return nil, err
			}

			args = append(args, argVar)
		}

		var env []string
		for _, e := range cmd.Env {
			envVar, err := resolveTmpl(e, tmplData{
				Env: environ,
			})
			if err != nil {
				return nil, err
			}

			env = append(env, envVar)
		}

		commands = append(commands, &runner.CommandStep{
			Command:    args,
			Env:        env,
			WorkingDir: cmd.WorkingDir,
		})
	}

	return &runner.CommandRunner{
		Steps: commands,
	}, nil
}

func resolveTmpl(str string, data interface{}) (string, error) {
	var buf bytes.Buffer
	tmpl, err := template.New("").Parse(str)
	if err != nil {
		return "", err
	}
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
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
