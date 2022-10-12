package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	configPath := flag.String("config", "slsabuild.yaml", "path to config file")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [FLAGS] -- [BUILD FLAGS]\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Println()
		fmt.Println("Run 'go help build' for details on build flags.")
	}

	flag.Parse()

	// Read the config
	cf, err := os.Open(*configPath)
	if err != nil {
		log.Fatalf("error: %v", err)

	}

	d := yaml.NewDecoder(cf)
	var config Config
	if err := d.Decode(&config); err != nil {
		log.Fatalf("error: %v", err)
	}

	// Run the commands.
	runner, err := config.Runner()
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	if _, err := runner.Run(context.Background()); err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Println("done")
}
