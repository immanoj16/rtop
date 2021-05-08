package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"runtime/debug"
	"strconv"
	"strings"

	"github.com/spf13/pflag"
)

var (
	version = ""
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
		fmt.Printf("Task version: %s\n", getVersion())
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

	username, addr := parseUserAndHost(pflag.Args())
	host, port := parseHostAndPort(addr)

	fmt.Println(username, "\t", host, "\t", port)
}

func parseUserAndHost(args []string) (username string, host string) {
	if len(args) == 1 {
		argHost := args[0]
		if i := strings.Index(argHost, "@"); i != 1 {
			username = argHost[:i]
			if i+1 >= len(argHost) {
				pflag.Usage()
			}
			host = argHost[i+1:]
		}
	} else {
		pflag.Usage()
	}
	if len(username) == 0 {
		username = currentUser.Username
	}
	return
}

func parseHostAndPort(addr string) (host string, port int) {
	if p := strings.Split(addr, ":"); len(p) == 2 {
		host = p[0]
		var err error
		if port, err = strconv.Atoi(p[1]); err != nil {
			log.Printf("bad port: %v", err)
			pflag.Usage()
		}
		if port <= 0 || port >= 65536 {
			log.Printf("bad port: %v", err)
			pflag.Usage()
		}
	} else {
		host = addr
	}

	if port == 0 {
		port = 22
	}
	return
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
