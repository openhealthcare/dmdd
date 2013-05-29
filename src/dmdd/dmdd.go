package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"launchpad.net/goyaml"
	"os"
)

type Config struct {
	Database string
	Server   struct {
		Bind string
		Port int ",flow"
	}
}

var config Config
var install string
var config_file string

func init() {
	flag.StringVar(&install, "install", "", "Specifies the zip file containing the DMD data to install")
	flag.StringVar(&config_file, "config", "config.yaml", "Points to the configuration file")
}

func usage() {
	fmt.Fprintf(os.Stderr, "DMD v0.1\n========\n")
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, USAGE)
	os.Exit(0)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	test_xml()

	content, err := ioutil.ReadFile(config_file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load config file, does it exist?\n")
		os.Exit(1)
	}

	// Read the config into the config var
	err = goyaml.Unmarshal([]byte(content), &config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse the config file, maybe it's malformed?\n")
		os.Exit(1)
	}
	fmt.Println(config.Server)

	// If the user has given us a zip file to install, we should try and
	// install it.
	if install != "" {
		//install_package(install)
		os.Exit(0)
	}

	// Haven't been told what to do, so let's serve some content
}

const USAGE = `

`
