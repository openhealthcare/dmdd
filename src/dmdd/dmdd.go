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
	Debug bool
}

var config Config
var install string
var config_file string
var cleanup bool

func init() {
	flag.StringVar(&install, "install", "", "Specifies the zip file containing the DMD data to install")
	flag.StringVar(&config_file, "config", "config.yaml", "Points to the configuration file")
	flag.BoolVar(&cleanup, "cleanup", false, "Cleans up the database")
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

	//test_xml_parsing()

	content, err := ioutil.ReadFile(config_file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Failed to load config file, does it exist?\n")
		os.Exit(1)
	}

	// Read the config into the config var
	err = goyaml.Unmarshal([]byte(content), &config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Failed to parse the config file, maybe it's malformed?\n")
		os.Exit(1)
	}

	err = db_init()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	defer db_close()

	if cleanup {
		fmt.Print("Cleaning up database...")
		db_clean()
		fmt.Println("...done")
		os.Exit(0)
	}

	// If the user has given us a zip file to install, we should try and
	// install it.
	if install != "" {
		fmt.Println("Looking for install package")
		install_package(install)
		fmt.Println("...done")
		os.Exit(0)
	}

	// Haven't been told what to do, so let's serve some content
}

const USAGE = `

`
