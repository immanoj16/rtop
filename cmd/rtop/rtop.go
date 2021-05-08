package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"log"
	"os"
	"runtime/debug"
)

var (
	version = ""
)

const usage = `Usage: rtop [-i <private-key-path>] [--version] [--help]
Options:
`

func main() {
	log.SetFlags(0)
	log.SetOutput(os.Stderr)

	pflag.Usage = func() {
		log.Print(usage)
		pflag.PrintDefaults()
	}

	var (
		versionFlag bool
		helpFlag bool
		privateKeyPath string
	)

	pflag.BoolVarP(&versionFlag, "version", "v", false, "Show rtop version")
	pflag.BoolVarP(&helpFlag, "help", "h", false, "Show rtop usage")
	pflag.StringVarP(&privateKeyPath, "input", "i", "", "Private key path")
	pflag.Parse()

	if versionFlag {
		fmt.Printf("Task version: %s\n", getVersion())
		return
	}

	if helpFlag {
		pflag.Usage()
		return
	}
}

func getVersion() string {
	if version != "" {
		return version
	}

	info, ok := debug.ReadBuildInfo()
	if !ok || info.Main.Version == "" {
		return "unknown"
	}

	version = info.Main.Version
	if info.Main.Sum != "" {
		version += fmt.Sprintf(" (%s)", info.Main.Sum)
	}

	return version
}
