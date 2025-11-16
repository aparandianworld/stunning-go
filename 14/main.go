package main

import (
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
)

type Config struct {
	Host string
	Port int
}

func readConfig(path string) (*Config, error) {

	if path == "" {
		return nil, fmt.Errorf("path is empty")
	}

	fh, err := os.Open(path) // Pattern: open a resource, check for error, close the resource with defer
	if err != nil {
		return nil, errors.Wrap(err, "failed to open config file") // This is coming from pkg/errors
	}

	defer fh.Close()

	cfg := &Config{}
	// parse config file
	if _, err := toml.NewDecoder(fh).Decode(cfg); err != nil {
		return nil, errors.Wrap(err, "failed to decode config file")
	}

	return cfg, nil
}

func main() {

	path := "./config.toml"
	cfg, err := readConfig(path)

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read config: %s\n", err)
		log.Printf("error: %+v", err)
		os.Exit(1)
	}

	fmt.Println(*cfg)

}
