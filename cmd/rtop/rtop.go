package main

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/immanoj16/rtop/pkg/utils"
	ver "github.com/immanoj16/rtop/pkg/version"
	"github.com/spf13/pflag"
)

var (
	version     = ""
	currentUser *user.User
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
		versionFlag    bool
		helpFlag       bool
		privateKeyPath string
	)

	pflag.BoolVarP(&versionFlag, "version", "v", false, "Show rtop version")
	pflag.BoolVarP(&helpFlag, "help", "h", false, "Show rtop usage")
	pflag.StringVarP(&privateKeyPath, "input", "i", "", "Private key path")
	pflag.Parse()

	if versionFlag {
		fmt.Printf("rtop version: %s\n", ver.GetVersion(version))
		return
	}

	if helpFlag {
		pflag.Usage()
		return
	}

	var err error
	currentUser, err = user.Current()
	if err != nil {
		log.Print(err)
		return
	}

	username, addr := utils.ParseUserAndHost(pflag.Args(), currentUser)
	host, port := utils.ParseHostAndPort(addr)

	fmt.Println(username, "\t", host, "\t", port)
}
